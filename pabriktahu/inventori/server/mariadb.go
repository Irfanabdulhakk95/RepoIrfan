package server

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	addInventori                   = `insert into Inventori(kodeinventori, namainventori, jumlah, status, createby, createon)values (?,?,?,?,?,?)`
	selectInventori                = `select kodeinventori, namainventori, jumlah, status, createby from Inventori Where status ='1'`
	updateInventori                = `update Inventori set namainventori=?, jumlah=?, status=?, updateby=?, updateon=? where kodeinventori=?`
	selectInventoriByNamaInventori = `select kodeinventori, namainventori, jumlah, status, createon from Inventori where namainventori=?`
)

//langkah 4
type dbReadWriter struct {
	db *sql.DB
}

func NewDBReadWriter(url string, schema string, user string, password string) ReadWriter {
	schemaURL := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, url, schema)
	db, err := sql.Open("mysql", schemaURL)
	if err != nil {
		panic(err)
	}
	return &dbReadWriter{db: db}
}

func (rw *dbReadWriter) AddInventori(inventori Inventori) error {
	fmt.Println("menambahkan data")
	tx, err := rw.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(addInventori, inventori.KodeInventori, inventori.NamaInventori, inventori.Jumlah, OnAdd, inventori.CreateBy, time.Now())
	//fmt.Println(err)
	if err != nil {
		return err

	}
	return tx.Commit()
}

func (rw *dbReadWriter) ReadInventoriByNamaInventori(namainventori string) (Inventori, error) {
	fmt.Println("show by nama")
	inventori := Inventori{NamaInventori: namainventori}
	err := rw.db.QueryRow(selectInventoriByNamaInventori, namainventori).Scan(&inventori.KodeInventori, &inventori.NamaInventori, &inventori.Jumlah,
		&inventori.Status, &inventori.CreateBy)

	if err != nil {
		return Inventori{}, err
	}

	return inventori, nil
}

func (rw *dbReadWriter) ReadInventori() (Inventoris, error) {
	fmt.Println("show all")
	inventori := Inventoris{}
	rows, _ := rw.db.Query(selectInventori)
	defer rows.Close()
	for rows.Next() {
		var k Inventori
		err := rows.Scan(&k.NamaInventori, &k.NamaInventori, &k.Jumlah,
			&k.Status, &k.CreateBy)
		if err != nil {
			fmt.Println("error query:", err)
			return inventori, err
		}
		inventori = append(inventori, k)
	}
	//fmt.Println("db nya:", mahasiswa)
	return inventori, nil
}

func (rw *dbReadWriter) UpdateInventori(kar Inventori) error {
	fmt.Println("update successfuly")
	tx, err := rw.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(updateInventori, kar.NamaInventori, kar.Jumlah, kar.Status, kar.UpdateBy, time.Now(), kar.KodeInventori)

	fmt.Println("name:", kar.NamaInventori, kar.Status)
	if err != nil {
		return err
	}

	return tx.Commit()
}
