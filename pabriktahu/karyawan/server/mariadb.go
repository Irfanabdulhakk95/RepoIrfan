package server

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	addKaryawan          = `insert into karyawan(namakaryawan, idagama, alamat, jeniskelamin, status, createby, createon)values (?,?,?,?,?,?,?)`
	selectKaryawan       = `select namakaryawan, idagama, alamat, jeniskelamin, status, createby from karyawan Where status ='1'`
	updateKaryawan       = `update karyawan set idagama=?, alamat=?, jeniskelamin=?, status=?, updateby=?, updateon=? where namakaryawan=?`
	selectKaryawanByNama = `select namakaryawan, idagama, alamat, jeniskelamin, status, createon from karyawan where namakaryawan=?`
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

func (rw *dbReadWriter) AddKaryawan(karyawan Karyawan) error {
	fmt.Println("insert")
	tx, err := rw.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(addKaryawan, karyawan.NamaKaryawan, karyawan.IdAgama, karyawan.Alamat, karyawan.Jeniskelamin, OnAdd, karyawan.CreateBy, time.Now())
	//fmt.Println(err)
	if err != nil {
		return err

	}
	return tx.Commit()
}

func (rw *dbReadWriter) ReadKaryawanByNama(namakaryawan string) (Karyawan, error) {
	fmt.Println("show by nama")
	karyawan := Karyawan{NamaKaryawan: namakaryawan}
	err := rw.db.QueryRow(selectKaryawanByNama, namakaryawan).Scan(&karyawan.NamaKaryawan, &karyawan.IdAgama, &karyawan.Alamat,
		&karyawan.Jeniskelamin, &karyawan.Status, &karyawan.CreateBy)

	if err != nil {
		return Karyawan{}, err
	}

	return karyawan, nil
}

func (rw *dbReadWriter) ReadKaryawan() (Karyawans, error) {
	fmt.Println("show all")
	karyawan := Karyawans{}
	rows, _ := rw.db.Query(selectKaryawan)
	defer rows.Close()
	for rows.Next() {
		var k Karyawan
		err := rows.Scan(&k.NamaKaryawan, &k.IdAgama, &k.Alamat,
			&k.Jeniskelamin, &k.Status, &k.CreateBy)
		if err != nil {
			fmt.Println("error query:", err)
			return karyawan, err
		}
		karyawan = append(karyawan, k)
	}
	//fmt.Println("db nya:", mahasiswa)
	return karyawan, nil
}

func (rw *dbReadWriter) UpdateKaryawan(kar Karyawan) error {
	fmt.Println("update successfuly")
	tx, err := rw.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(updateKaryawan, kar.IdAgama, kar.Alamat, kar.Jeniskelamin, kar.Status, kar.UpdateBy, time.Now(), kar.NamaKaryawan)

	fmt.Println("name:", kar.NamaKaryawan, kar.Status)
	if err != nil {
		return err
	}

	return tx.Commit()
}
