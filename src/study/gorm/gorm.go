package gorm

import (
	"database/sql"
	"fmt"
	"log"

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

func Select() {
	db, err := sql.Open("mysql", "sunhapark:root@tcp(127.0.0.1:3306)/students?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect DB")
	}
	defer db.Close()

	// QueryRow() : 1개의 row만 리턴
	var name string
	errQR := db.QueryRow("SELECT name FROM students WHERE id = 2").Scan(&name)
	if errQR != nil {
		panic(errQR.Error())
	}
	fmt.Println(name)
	//(Outputs) Harry

	// Query() : 여러 개의 row 리턴
	var id int
	rows, err := db.Query("SELECT id, name FROM students where id >= ?", 1)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, name)
	}
	//(Outputs) 1 Harry / 2 Harry / 3 Ron / 4 Hermione
}

func Insert() {
	db, err := sql.Open("mysql", "sunhapark:root@tcp(127.0.0.1:3306)/students?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect DB")
	}
	defer db.Close()

	// INSERT문 실행
	result, err := db.Exec("INSERT INTO students(name, age) VALUES (?, ?)", "Malfoy", 20)
	if err != nil {
		log.Fatal(err)
	}

	// sql.Result.RowsAffected() 체크
	n, err := result.RowsAffected()
	if n == 1 {
		fmt.Println("1 row insered")
	}
	//(Outputs) 1 row insered
}

func PreparedStatement() {
	db, err := sql.Open("mysql", "sunhapark:root@tcp(127.0.0.1:3306)/students?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect DB")
	}
	defer db.Close()

	// Prepared Statement 생성
	stmt, err := db.Prepare("UPDATE students SET name=? WHERE age=?")
	checkError(err)
	defer stmt.Close()

	// Prepared Statement 실행
	_, err = stmt.Exec("Harry", 23) // Placeholder 파라미터 순서대로 전달
	checkError(err)
	_, err = stmt.Exec("Ron", 23)
	checkError(err)
	_, err = stmt.Exec("Harmione", 23)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
