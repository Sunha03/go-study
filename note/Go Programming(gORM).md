# Go Programming(gORM)

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

## gORM

* ORM(Object Relational Mapping, 객체-관계 매핑)

   : 객체와 관계형 DB의 데이터를 자동으로 매핑(연결)해주는 것

  - 객체 지향 프로그래밍은 클래스를 사용하고, 관계형 DB는 테이블을 사용함
  - 객체 모델과 관계형 모델 간에 불일치가 존재함
  - ORM을 통해 깩체 간의 관계를 바탕으로 SQL을 자동으로 생성하여 불일치를 해결함

* gORM : go언어에서 사용 가능한 ORM(Object Relation Mapping) 라이브러리([*http://gorm.io)*](http://gorm.io/)

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

### 2. gORM 설치

* gORM 설치

  ```shell
  $ go get -u github.com/jinzhu/gorm
  ```

  -> https://github.com/jinzhu/gorm 사이트에 접속 > 소스 다운로드 > 로컬 컴퓨터의 GOPATH/src/디렉토리 아래에 동일한 디렉토리 경로 구조로 저장함(ex. GOPATH/src/github.com/jinzhu/gorm에 clone함)

  -> -u 옵션 : gorm이 내부적으로 사용하는 패키지(의존성)는 자동으로 다운로드하여 GOPATH/src/디렉토리 아래에 소스를 다운로드하여 저장 > 저장이 완료되면 각 소스를 빌ㄹ드하여 라이브러리(패키지)를 생성

  ** 라이브러리는 정적(static) 라이브러리로 생성되고, GOPATH/pkg/디렉토리 아래에 라이브리러 확장자인 *.a 형태로 저장됨

  ** 참고로, 다운로드한 패키지가 실행파일(유틸리티, 도구)인 경우에는 실행파일을 만들어 이를 GOPATH/bin 이하에 생성/저장

### 3. gORM용 MySQL 드라이버 설치

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
  
  $ tail -f /(경로)/Sunha-MacBookPro.log	# 로그 파일 확인
  ```

  ![image-20210112192847752](/Users/sunhapark/project/GoStudy/note/images/image-20210112192847752.png)

  ![image-20210112194043910](/Users/sunhapark/project/GoStudy/note/images/image-20210112194043910.png)



[참고] [https://medium.com/@amoebamach/go%EC%96%B8%EC%96%B4%EC%97%90%EC%84%9C-orm-gorm-83ab33ecdc98]