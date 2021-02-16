# Go Programming(GORM)

## * Mac MariaDB 설치

* Homebrew로 MariaDB 설치

  ```sh
  $ brew install mariadb
  ```

  ![image-20210112171612015](/Users/sunhapark/project/GoStudy/note/images/image-20210112171612015.png)

* MariaDB server 시작 / 종료 / 상태확인

  ```shell
  $ mysql.server start	# db 서버 시작
  $ mysql.server stop		# db 서버 종료
  $ mysql.server status	# db 서버 상태확인
  ```

  ![image-20210112172116418](/Users/sunhapark/project/GoStudy/note/images/image-20210112172116418.png)

* MariaDB server를 자동으로 시작하는 Homebrew 서비스 기능

  ```shell
  $ brew services start mariadb
  ```

* MariaDb server 시작 후 로그인

  ```shell
  $ mysql		# 사용자 계정으로 로그인
  $ sudo mysql -u root	# root 계정으로 로그인
  ```

  ```shell
  $ MariaDB > show databases;	# db 확인
  $ MariaDB > use mysql		# mysql 사용
  $ MariaDB > show tables like 'user';	# user 테이블 확인
  $ MariaDB > select host, user, password from user;	# user 테이블의 host, user, password 확인
  
  $ MariaDB > update user set password=password('new-password') where user='root';	# root 계정 패스워드 설정
  => 실행오류(ERROR 1356 (HY000): View 'mysql.user' references invalid table(s) or column(s) or function(s) or definer/invoker of view lack rights to use them)
  
  $ MariaDB > set password = password("new_password"); # 패스워드 설정1
  $ MariaDB > ALTER USER 'root'@'localhost' IDENTIFIED BY '(new_password)'	# 패스워드 설정2
  $ MariaDB > flush privileges;	# 변경사항 반영
  $ MariaDB > select host, user, password from user;	# user 테이블의 host, user, password 확인
  
  $ MariaDB > quit	# 종료
  ```

  ![image-20210112172437444](/Users/sunhapark/project/GoStudy/note/images/image-20210112172437444.png)

  -> 초기에는 비밀번호가 없어서 바로 접속 가능

  ![image-20210112173653372](/Users/sunhapark/project/GoStudy/note/images/image-20210112173653372.png)

  -> user 테이블 확인

  ![image-20210112180957779](/Users/sunhapark/project/GoStudy/note/images/image-20210112180957779.png)

  ![image-20210112181823669](/Users/sunhapark/project/GoStudy/note/images/image-20210112181823669.png)

  -> 패스워드 변경

## MariaDB 설정 및 GORM 설치

### 0. 전제조건

​	1) mariaDB 설치
​	2) golang(go) 개발환경 설치

### 1. MariaDB에 사용자 생성 및 권한부여

 1. bash에서 root 사용자로 mysql root 사용자 접속

    ```shell
    $ sudo mysql -u root	# root 계정으로 로그인
    ```

	2. mysql에서 사용자 생성 및 권한 부여

    ```shell
    $ MariaDB > CREATE user 'testuser'@'%' IDENTIFIED BY 'testpass';	# id:testuser, pw:testpass 계정 생성
    $ MariaDB > CREATE user 'testuser'@'localhost' IDENTIFIED BY 'testpass';	# id:testuser, pw:testpass 계정 생성
    $ MariaDB > GRANT ALL PRIVILEGES ON *.* TO 'testuser'@'&';	# 권한 부여
    $ MariaDB > GRANT ALL PRIVILEGES ON *.* TO 'testuser'@'localhost';	# 권한 부여
    $ MariaDB > FLUSH PRIVILEGES;		# 변경사항 저장
    $ MariaDB > select * from mysql.user where user='testuser'\G		# 계정 권한 확인 가능
    ```

    ![image-20210112191424782](/Users/sunhapark/project/GoStudy/note/images/image-20210112191424782.png)

### 2. GORM 설치

* gORM 설치

  ```shell
  $ go get -u github.com/jinzhu/gorm
  ```

  -> https://github.com/jinzhu/gorm 사이트에 접속 > 소스 다운로드 > 로컬 컴퓨터의 GOPATH/src/디렉토리 아래에 동일한 디렉토리 경로 구조로 저장함(ex. GOPATH/src/github.com/jinzhu/gorm에 clone함)

  -> -u 옵션 : gorm이 내부적으로 사용하는 패키지(의존성)는 자동으로 다운로드하여 GOPATH/src/디렉토리 아래에 소스를 다운로드하여 저장 > 저장이 완료되면 각 소스를 빌ㄹ드하여 라이브러리(패키지)를 생성

  ** 라이브러리는 정적(static) 라이브러리로 생성되고, GOPATH/pkg/디렉토리 아래에 라이브리러 확장자인 *.a 형태로 저장됨

  ** 참고로, 다운로드한 패키지가 실행파일(유틸리티, 도구)인 경우에는 실행파일을 만들어 이를 GOPATH/bin 이하에 생성/저장

### 3. GORM용 MySQL 드라이버 설치

* mysql, postgresql 등 DB별로 별도 설치가 필요함

  ```shell
  $ go get github.com/go-sql-driver/mysql
  ```

### 4. MySQL 쿼리 로그 설정

* mysql 쿼리를 추적하기 위해, MySQL 수준에서 수행되는 쿼리를 로그 파일로 출력하도록 설정

* gORM은 go언어에 치우친 DB 처리 구현을 보여주는데, 실제 DB에서 실행되는 명령어를 확인하기 위해 해당 작업이 필요함. 혹은 순수 DB 수준에서 쿼리를 튜닝하고 싶은 경우에도 수행되는 SQL 쿼리 및 수행시간을 추적하여, 분석/튜닝하는 용도로 사용할 수 있음

  ```shell
  $ MariaDB > SET GLOBAL general_log = 1;
  $ MariaDB > SELECT @@log_output, @@general_log, @@general_log_file;
  
  $ tail -f /usr/local/var/mysql/Sunha-MacBookPro.log	# 로그 파일 확인
  ```

  ![image-20210112192847752](/Users/sunhapark/project/GoStudy/note/images/image-20210112192847752.png)

  ![image-20210112194043910](/Users/sunhapark/project/GoStudy/note/images/image-20210112194043910.png)



## GORM

* ORM(Object Relational Mapping, 객체-관계 매핑)

   : 객체와 관계형 DB의 데이터를 자동으로 매핑(연결)해주는 것

  - 객체 지향 프로그래밍은 클래스를 사용하고, 관계형 DB는 테이블을 사용함
  - 객체 모델과 관계형 모델 간에 불일치가 존재함
  - ORM을 통해 깩체 간의 관계를 바탕으로 SQL을 자동으로 생성하여 불일치를 해결함

* gORM : go언어에서 사용 가능한 ORM(Object Relation Mapping) 라이브러리([*http://gorm.io)*](http://gorm.io/)

### Package

* 표준 패키지 database/sql을 사용함. database/sql 패키지는 관계형 DB들에게 공통적으로 사용되는 인터페이스들을 제공하고 있음

* database/sql 패키지는 여러 종류의 SQL DB를 지원하는데, 각각의 DB Driver와 함께 사용됨

  * MySQL : https://github.com/go-sql-driver/mysql
  * MSSQL : https://github.com/denisenkom/go-mssqldb
  * Oracle : https://github.com/rana/ora
  * Postgress : https://gitjib.com/lib/pq
  * SQLite : https://github.com/mattn/go-sqlite3
  * DB2 : https://bitbucket.org/phiggins/db2cli

  ```go
  import (
    _ "github.com/lib/pq"	// '_' : 드리이버를 사용하지 않음. 사용하진 않지만 연결을 위해 명시적으로 드라이버를 지정해야함. 아니면 에러 발생.
  	"database/sql"
  )
  ```

### Connect

* sql.DB - database/sql 패키지에서 가장 중요한 Type. 일반적으로 sql.Open()을 사용하여 sql.DB 객체를 얻음

* sql.Open(드라이버, Connection) : 어떤 DB를 사용할 것인지, 해당 DB 연결 정보를 제공하면, 결과로 sql.DB 객체를 얻음

* sql.DB 객체 얻기 >  sql.DB의 여러 메소드를 사용하여 쿼리 > SQL문 실행

* sql.Open()은 실제 DB Connection을 Open하지 않음. sql.DB는 드라이버 종류와 Connection 정보를 갖고 있지만, 실제 DB에 연결하지 않으며, 많은 경우 Connection 정보 조차 체크하지 않음. 실제 DB Connection은 Query 등과 같이 실제 DB 연결이 필요한 시점에 이루어짐

  ```go
  db, err := sql.Open(드라이버명, ₩host=db의ip port=db포트번호 user=접속할db유저 password=접속할db패스워드 dbname=접속할db명 sslmode=disable`)
  
  // sql.DB 객체 생성
  db, err := sql.Open("mysql", `host=localhost port=3306 user=sunhapark password=XXXX dbname=mysql sslmode=disable`)
  
  // db 사용 후 close
  defer db.Close()
  
  // error 출력
  if err != nil || db.Ping() != nil {
  	panic(err.Error())
  }
  ```

### Create

* db.Exec() : Raw SQL 실행

  ```go
  func  main() {
  	db, _  := gorm.Open("mysql", "testuser:testpass@tcp(127.0.0.1:3306)/")
  	defer db.Close() // 스콥(scope)을 벗어나면, db.CLose()실행, 이 경우는 main()함수 종료시
  	dbName  :=  "testdb"  // 생성할 데이터베이스 이름
  	// Raw SQL로 데이터베이스 생성
  	db.Exec("CREATE DATABASE IF NOT EXISTS "  + dbName)
  	db.Exec("commit;")
  }
  ```

  -> MySQL의 쿼리 로그로 결과 확인

  ![image-20210113004510818](/Users/sunhapark/project/GoStudy/note/images/image-20210113004510818.png)

### CRUD

* 테이블 생성 및 레코드 CRUD 연산 수행

  ```go
  type Student struct {
  	gorm.Model
  	Id string
  	Name string
  	Age int
  }
  
  func  main() {
  	db, err  := gorm.Open("mysql", "testuser:testpass@tcp(127.0.0.1:3306)/")
  	if err != nil {
  		panic("failed to connect DB")
  	}
  	defer db.Close()
  	
  	// 스키마를 마이그레이트(테이블 생성)
  	db.AutoMigrate(&Student{})
  	
  	db.Create(&Student{Id: "id001", Name: "Harry", Age: 20})
  
  	// Read : 읽기
  	var student Student
  	db.First(&student, 1)	// id가 1인 student 찾기
  	db.First(&student, "name = ?", "Harry") // name이 Harry인 학생 찾기
  	
  	// Update : 수정
  	db.Model(&student).Update("Age", 21)
  	
  	// Delete : 삭제
  	db.Delete(&student)
  }
  ```

  -> 실행 결과 로그

  ![image-20210113010113076](/Users/sunhapark/project/GoStudy/note/images/image-20210113010113076.png)

### Select

* 조회 함수 2개 - QueryRow(), Query()

  * QueryRow() : 1개의 Row만 리턴 or 1개의 Row만 리턴이 될 것을 예상한 경우

    -> Scan() : 1개의 Row에서 실제 데이터를 읽어 로컬 변수에 할당하기 위해 사용

    ```go
    var name string
    err := db.QueryRow("SELECT name FROM test WHERE id = 10").Scan(&name)
    if err != nil {
    	panic(err.Error())
    }
    fmt.Println(name)
    ```

    ![image-20210113013006260](/Users/sunhapark/project/GoStudy/note/images/image-20210113013006260.png)

  * Query() : 복수 개의 Row를 리턴

    -> Next() : 복수 Row에서 다음 Row로 이동하기 위해 사용
  
    ```go
    var id int
    var name string
    rows, err := db.Query("SELECT id, name FROM test1 where id >= $1", 1)
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
    ```

    ![image-20210113013313684](/Users/sunhapark/project/GoStudy/note/images/image-20210113013313684.png)

    -> Parameterized Query :  SQL Injection과 같은 문제를 방지하기 위해 파라미터를 문자열 결합이 아닌 별도의 파라미터로 대입시키는 방식

    -> Placeholder $1에는 1이 대입됨
  
    -> Placeholder는 DB 종류에 따라 다르게 사용됨.
    
    ​	MySQL - ? / Oracle - :val1, :val2 / PostgreSQL - $1, $2

### Insert, Delete, Update

* db.Exec() : INSERT, DELETE, UPDATE와 같이 DML Operation을 사용하기 위한 함수

* SELECT할 때, 리턴되는 데이터가 없는 경우에도 db.Exec() 사용함

* db.Exec()는 sql.Result와 error 객체를 리턴하며, sql.Result 객체로부터 갱신된 레코드수(RowsAffected())와 새로 추가된 Id(LastInsertId())를 구할 수 있음

  ```go
  // INSERT문 실행
  result, err := db.Exec("INSERT INTO test1 VALUES (?, ?)", 11, "Jack")
  if err != nil {
  	log.Fatal(err)
  }
  
  // sql.Result.RowsAffected() 체크
  n, err := result.RowsAffected()
  if n == 1 {
  	fmt.Println("1 row insered")
  }
  ```

  ![image-20210114021933900](/Users/sunhapark/project/GoStudy/note/images/image-20210114021933900.png)

### Prepared Statement

* Prepared Statement : DB 서버에 Placeholder를 가진 SQL문을 미리 준비시키는 것으로, 차후 해당 Statement를 호출할 때 준비된 SQL문을 빠르게 실행하도록 하는 기법

* db.Prepare() : Placeholder를 가진 SQL문을 미리 준비시키고, sql.Stmt 객체를 리턴

* 나중에 이 sql.Stmt 객체의 Exec() / Query() / QueryRow()를 이용하여 준비된 SQL문을 실행함

  ```go
    // Prepared Statement 생성
    stmt, err := db.Prepare("UPDATE test1 SET name=? WHERE id=?")
    checkError(err)
    defer stmt.Close()
  
    // Prepared Statement 실행
    _, err = stmt.Exec("Tom", 1)	// Placeholder 파라미터 순서대로 전달
    checkError(err)
    _, err = stmt.Exec("Jack", 2)
    checkError(err)
    _, err = stmt.Exec("Shawn", 3)
    checkError(err)
  }
  
  func checkError(err error) {
  	if(err != nil) {
  		panic(err)
  	}
  }
  ```

  ![image-20210114021955610](/Users/sunhapark/project/GoStudy/note/images/image-20210114021955610.png)

### 트랜잭션

* 트랜잭션 - 복수 개의 SQL 문을 실행하다 중간에 어떤 한 SQL 문에서라도 에러가 발생하면 전체 SQL문을 취소하고(롤백), 모두 성공적으로 실행되어야 전체를 커밋하게 됨

* db.Begin() : 복수 개의 SQL 문을 하나의 트랜잭션으로 묶음

* sql.Tx 타입의 Begin()는 sql.Tx 객체를 리턴하는데, 마지막에 최종 Commit을 위해 Tx.Commit()을, 롤백을 위해 Tx.Rollback()를 호출함.

  ```go
  // 트랜잭션 시작
  tx, err := db.Begin()
  checkError(err)
  defer tx.Rollback()		// 중간에 에러 발생 시 롤백
  
  // INSERT 문 실행
  _, err = db.Exec("INSERT INTO test1 VALUES (?, ?)", 15, "Jack")
  checkError(err)
  
  _, err = db.Exec("INSERT INTO test2 VALUES (?, ?)", 15, "Data")
  checkError(err)
  
  // 트랜잭션 커밋
  err = tx.Commit()
  checkError(err)
  ```

  ![image-20210114023421041](/Users/sunhapark/project/GoStudy/note/images/image-20210114023421041.png)





[참고] [https://medium.com/@amoebamach/go%EC%96%B8%EC%96%B4%EC%97%90%EC%84%9C-orm-gorm-83ab33ecdc98]

[참고] [https://brownbears.tistory.com/186]