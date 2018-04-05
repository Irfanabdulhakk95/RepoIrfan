package server

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	addSupplier                = `insert into Supplier(kodesupplier, namasupplier,keterangan, status, createby, createon)values (?,?,?,?,?,?)`
	selectSupplier             = `select kodesupplier, namasupplier, keterangan, status, createby from Supplier Where status ='1'`
	updateSupplier             = `update Supplier set namasupplier=?, keterangan=?, status=?, updateby=?, updateon=? where kodesupplier=?`
	selectSupplierByKeterangan = `select kodesupplier, namasupplier, status, createby, keterangan from Supplier where keterangan LIKE ?`
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

func (rw *dbReadWriter) AddSupplier(supplier Supplier) error {
	fmt.Println("menambahkan data")
	tx, err := rw.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(addSupplier, supplier.KodeSupplier, supplier.NamaSupplier, supplier.Keterangan, OnAdd, supplier.CreateBy, time.Now())
	//fmt.Println(err)
	if err != nil {
		return err

	}
	return tx.Commit()
}

func (rw *dbReadWriter) ReadSupplierByKeterangan(keterangan string) (Suppliers, error) {
	fmt.Println("show by keterangan")
	supplier := Suppliers{}
	rows, _ := rw.db.Query(selectSupplierByKeterangan, keterangan)
	defer rows.Close()
	for rows.Next() {
		var u Supplier
		err := rows.Scan(&u.KodeSupplier, &u.NamaSupplier,
			&u.Status, &u.CreateBy, &u.Keterangan)
		if err != nil {
			fmt.Println("error query:", err)
			return supplier, err
		}
		supplier = append(supplier, u)
	}
	//fmt.Println("db nya:", mahasiswa)
	return supplier, nil

}
func (rw *dbReadWriter) ReadSupplier() (Suppliers, error) {
	fmt.Println("show all")
	supplier := Suppliers{}
	rows, _ := rw.db.Query(selectSupplier)
	defer rows.Close()
	for rows.Next() {
		var k Supplier
		err := rows.Scan(&k.KodeSupplier, &k.NamaSupplier, &k.Keterangan,
			&k.Status, &k.CreateBy)
		if err != nil {
			fmt.Println("error query:", err)
			return supplier, err
		}
		supplier = append(supplier, k)
	}
	//fmt.Println("db nya:", mahasiswa)
	return supplier, nil
}

func (rw *dbReadWriter) UpdateSupplier(kar Supplier) error {
	fmt.Println("update successfuly")
	tx, err := rw.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(updateSupplier, kar.NamaSupplier, kar.Keterangan, kar.Status, kar.UpdateBy, time.Now(), kar.KodeSupplier)

	fmt.Println("name:", kar.NamaSupplier, kar.Keterangan, kar.Status)
	if err != nil {
		return err
	}

	return tx.Commit()
}
