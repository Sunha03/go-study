package gorm

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Student struct {
	gorm.Model
	Name string
	Age  int
}

func ConnMySQL() {
	// Create the database handle, confirm driver is present
	db, err := sql.Open("mysql", "sunhapark:root@tcp(127.0.0.1:3306)/")
	defer db.Close() // Close the DB
	if err != nil || db.Ping() != nil {
		panic(err.Error())
	}

	// Connect and check the server version
	var version string
	db.QueryRow("SELECT VERSION()").Scan(&version)
	fmt.Println("Connected to : ", version)
	//(Outputs) Connected to :  10.5.8-MariaDB
}

func CreateDB() {
	db, _ := gorm.Open("mysql", "sunhapark:root@tcp(127.0.0.1:3306)/")
	defer db.Close()

	dbName := "students" // 생성할 DB 이름

	// Raw SQL로 DB 생성
	db.Exec("CREATE DATABASE IF NOT EXISTS " + dbName)
	db.Exec("commit;")
}

func CRUD() {
	db, err := gorm.Open("mysql", "sunhapark:root@tcp(127.0.0.1:3306)/students?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect DB")
	}
	defer db.Close()

	// 스키마를 마이그레이트(테이블 생성)
	db.AutoMigrate(&Student{})

	db.Create(&Student{Name: "Harry", Age: 20})
	db.Create(&Student{Name: "Ron", Age: 20})
	db.Create(&Student{Name: "Hermione", Age: 20})

	// Read : 읽기
	var student Student
	db.First(&student, 1)                   // id가 1인 student 찾기
	db.First(&student, "name = ?", "Harry") // name이 Harry인 학생 찾기

	// Update : 수정
	db.Model(&student).Update("Age", 21)

	// Delete : 삭제
	//db.Delete(&student)
}
