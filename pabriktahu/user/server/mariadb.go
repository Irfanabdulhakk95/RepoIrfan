package server

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	addUser         = `insert into User(Username, password,IDKaryawan, status, createdby, createdon)values (?,?,?,?,?,?)`
	selectUserByKet = `select Username, password,IDKaryawan, status, createdby, keterangan from User where keterangan like ? order by IDKaryawan asc`
	selectUser      = `select Username, password,IDKaryawan, status, createdby from User order by IDKaryawan asc`
	updateUser      = `update User set username=?,password=?, status=?, updatedby=?, updatedon=? where IDKaryawan=?`
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

func (rw *dbReadWriter) AddUser(user User) error {
	fmt.Println("insert")
	tx, err := rw.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(addUser, user.Username, user.Password, user.IDKaryawan, OnAdd, user.CreatedBy, time.Now())
	//fmt.Println(err)
	if err != nil {
		return err

	}
	return tx.Commit()
}

func (rw *dbReadWriter) ReadUserByKet(keterangan string) (Users, error) {
	fmt.Println("show all")
	user := Users{}
	rows, _ := rw.db.Query(selectUserByKet, keterangan)
	defer rows.Close()
	for rows.Next() {
		var u User
		err := rows.Scan(&u.Username, &u.Password,
			&u.IDKaryawan, &u.Status, &u.CreatedBy, &u.Keterangan)
		if err != nil {
			fmt.Println("error query:", err)
			return user, err
		}
		user = append(user, u)
	}
	//fmt.Println("db nya:", mahasiswa)
	return user, nil
}

func (rw *dbReadWriter) ReadUser() (Users, error) {
	fmt.Println("show all")
	user := Users{}
	rows, _ := rw.db.Query(selectUser)
	defer rows.Close()
	for rows.Next() {
		var u User
		err := rows.Scan(&u.Username, &u.Password,
			&u.IDKaryawan, &u.Status, &u.CreatedBy)
		if err != nil {
			fmt.Println("error query:", err)
			return user, err
		}
		user = append(user, u)
	}
	//fmt.Println("db nya:", mahasiswa)
	return user, nil
}

func (rw *dbReadWriter) UpdateUser(us User) error {
	fmt.Println("update successfuly")
	tx, err := rw.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(updateUser, us.Username, us.Password, us.Status, us.UpdatedBy, time.Now(), us.IDKaryawan)

	//fmt.Println("kode:", us.Username, us.IDKaryawan)
	if err != nil {
		return err
	}

	return tx.Commit()
}
