package server

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	addKendaraan               = `insert into Kendaraan(kodekendaraan, jeniskendaraan, nomorplat, status, createby, createon)values (?,?,?,?,?,?)`
	selectKendaraan            = `select kodekendaraan, jeniskendaraan, nomorplat, status, createby from Kendaraan Where status ='1'`
	updateKendaraan            = `update Kendaraan set jeniskendaraan=?, nomorplat=?, status=?, updateby=?, updateon=? where kodekendaraan=?`
	selectKendaraanByNomorPlat = `select kodekendaraan, jeniskendaraan, nomorplat, status, createon from Kendaraan where nomorplat=?`
	selectKendaraanByNoPlat    = `select kodekendaraan, jeniskendaraan, nomorplat, status, createon from Kendaraan where nomorplat LIKE`
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

func (rw *dbReadWriter) AddKendaraan(kendaraan Kendaraan) error {
	fmt.Println("menambahkan data")
	tx, err := rw.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(addKendaraan, kendaraan.KodeKendaraan, kendaraan.JenisKendaraan, kendaraan.NomorPlat, OnAdd, kendaraan.CreateBy, time.Now())
	//fmt.Println(err)
	if err != nil {
		return err

	}
	return tx.Commit()
}

func (rw *dbReadWriter) ReadKendaraanByNomorPlat(nomorplat string) (Kendaraan, error) {
	fmt.Println("show by nomor")
	kendaraan := Kendaraan{NomorPlat: nomorplat}
	err := rw.db.QueryRow(selectKendaraanByNomorPlat, nomorplat).Scan(&kendaraan.KodeKendaraan, &kendaraan.JenisKendaraan, &kendaraan.NomorPlat,
		&kendaraan.Status, &kendaraan.CreateBy)

	if err != nil {
		return Kendaraan{}, err
	}

	return kendaraan, nil
}

func (rw *dbReadWriter) ReadKendaraan() (Kendaraans, error) {
	fmt.Println("show all")
	kendaraan := Kendaraans{}
	rows, _ := rw.db.Query(selectKendaraan)
	defer rows.Close()
	for rows.Next() {
		var k Kendaraan
		err := rows.Scan(&k.KodeKendaraan, &k.JenisKendaraan, &k.NomorPlat,
			&k.Status, &k.CreateBy)
		if err != nil {
			fmt.Println("error query:", err)
			return kendaraan, err
		}
		kendaraan = append(kendaraan, k)
	}
	//fmt.Println("db nya:", mahasiswa)
	return kendaraan, nil
}

func (rw *dbReadWriter) ReadKendaraanByNoPlat() (Kendaraans, error) {
	fmt.Println("show all")
	kendaraan := Kendaraans{}
	rows, _ := rw.db.Query(selectKendaraan)
	defer rows.Close()
	for rows.Next() {
		var k Kendaraan
		err := rows.Scan(&k.KodeKendaraan, &k.JenisKendaraan, &k.NomorPlat,
			&k.Status, &k.CreateBy)
		if err != nil {
			fmt.Println("error query:", err)
			return kendaraan, err
		}
		kendaraan = append(kendaraan, k)
	}
	//fmt.Println("db nya:", mahasiswa)
	return kendaraan, nil
}
func (rw *dbReadWriter) UpdateKendaraan(kar Kendaraan) error {
	fmt.Println("update successfuly")
	tx, err := rw.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(updateKendaraan, kar.JenisKendaraan, kar.NomorPlat, kar.Status, kar.UpdateBy, time.Now(), kar.KodeKendaraan)

	fmt.Println("name:", kar.JenisKendaraan, kar.Status)
	if err != nil {
		return err
	}

	return tx.Commit()
}
