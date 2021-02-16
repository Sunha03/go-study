

# GORM

## ORM

### ORM(=Object Relational Mapping, 객체-관계 매핑)

* 객체와 관계형 DB의 데이터를 자동으로 매핑(연결)해주는 것

  -> Object(객체)와 Relation(관계)를 Mapping(연결) 해주는 개념

  -> 객체 간의 관계를 바탕으로 SQL문을 자동 생성하여 모델 간의 불일치를 해결함

  -> Java의 JPA, Node의 Sequelize가 대표적

  -> 객체 지행 프로그래밍은 클래스를 사용하고, 관계형 DB는 테이블을 사용함

  ->  DB 데이터 <-- (매핑) --> Object 필드

  ##### 1. ORM을 사용하는 이유?

  * 종속성
    -> DBMS에 대한 종속성이 적음
    -> SQL 구문에 대한 종속성이 적음
  * 생산성
    -> 객체별로 코드를 작성하여 가독성이 높음
    -> 객체 지향적 코드로 인해 비즈니스 로직에 집중이 가능함
    -> 선언문, 할당문, CRUD 등 부수적인 코드가 줄어 작성하는 코드의 양이 30% 정도 절감됨
  * 유지보수
    -> 매핑정보가 명확하여 ERD에 대한 의존도가 늦음
    -> ORM은 객체가 독립적으로 작성되어 있어 유지보수가 좋음

  ##### 2. ORM의 장점

  ​	-> 객체 지향적인 코드로 인해 더 직관적이고 비즈니스 로직에 더 집중할 수 있게 도와줌
  ​	-> 재사용 및 유지보수의 편리성이 증가함
  ​	-> DBMS에 대한 종속성이 줄어듦

  ##### 3. ORM의 단점

  ​	-> 완벽한 ORM으로만 서비스를 구현하기 어려움
  ​	-> 프로시저가 많은 시스템에선 ORM의 객체 지향적인 장점을 활용하기 어려움

## GORM

### GORM

* 개발자 친화적인 것을 목표로 하는 Go 언어 ORM 라이브러리

  ##### 1. 설치

  ```shell
  go get -u gorm.io/gorm
  go get -u gorm.io/driver/sqlite
  ```

  ##### 2. 구조 작성

  * 모델 선언

    * 모델 : 기본 Go 유형, 포인터/별칭 or Scanner / Valuer 인터페이스를 구현하는 사용자 정의 유형을 갖춘 일반적인 구조체

    *   ex)

      ```go
      type User struct {
        ID           uint
        Name         string
        Email        *string
        Age          uint8
        Birthday     *time.Time
        MemberNumber sql.NullString
        ActivatedAt  sql.NullTime
        CreatedAt    time.Time
        UpdatedAt    time.Time
      }
      ```

  * Convention

    -> ID : 기본키로 사용
    -> 구조체/변수 이름 : snake_cases화 한 것을 테이블/열 이름으로 사용
    -> CreatedAt, UpdatedAt : 생성/업데이트 시간을 추적
    
    -> gorm에서 사용하는 convention을 따르는 경우, 구성/코드를 거의 작성하지 않아도 됨
    
  * gorm.Model

    -> gorm은 ID, CreatedAt, UpdatedAt, DeletedAt 필드를 포함하는 gorm.Model 구조체를 정의함

    ```go
    // gorm.Model definition
    type Model struct {
      ID        uint           `gorm:"primaryKey"`
      CreatedAt time.Time
      UpdatedAt time.Time
      DeletedAt gorm.DeletedAt `gorm:"index"`
    }
    ```

  * Advanced

    * Field-Level Permission
      -> 일반적인 필드는 모든 권한을 가지지만 태그를 사용하여 필드의 권한을 read-only / write-only / create-only / update-only 또는 ignored로 설정할 수 있음

      (단, gorm migrator를 사용해 테이블을 만들 때에는 ignored  필드는 생성되지 않음)

      ```go
      type User struct {
        Name string `gorm:"<-:create"` // allow read and create
        Name string `gorm:"<-:update"` // allow read and update
        Name string `gorm:"<-"`        // allow read and write (create and update)
        Name string `gorm:"<-:false"`  // allow read, disable write permission
        Name string `gorm:"->"`        // readonly (disable write permission unless it configured )
        Name string `gorm:"->;<-:create"` // allow read and create
        Name string `gorm:"->:false;<-:create"` // createonly (disabled read from db)
        Name string `gorm:"-"`  // ignore this field when write and read with struct
      }
      ```

    * Creating/Updating Time/Unix (Milli/Nano) Seconds Tracking
      -> gorm은 CreatedAt, UpdatedAt을 사용하여 작성/업데이트 시간을 추적하고, gorm은 필드가 정의된 경우 작성/업데이트 시 현재 시간을 설정함

      ```go
      type User struct {
        CreatedAt time.Time // Set to current time if it is zero on creating
        UpdatedAt int       // Set to current unix seconds on updaing or if it is zero on creating
        Updated   int64 `gorm:"autoUpdateTime:nano"` // Use unix nano seconds as updating time
        Updated   int64 `gorm:"autoUpdateTime:milli"`// Use unix milli seconds as updating time
        Created   int64 `gorm:"autoCreateTime"`      // Use unix seconds as creating time
      }
      ```

    * Embedded Struct
      -> 익명 필드의 경우, gorm은 다음과 같이  필드를 상위 구조에 포함함

      ```go
      type User struct {
        gorm.Model // include ID, CreatedAt, UpdatedAt, DeletedAt fields
        Name string
      }
      // equals
      type User struct {
        ID        uint           `gorm:"primaryKey"`
        CreatedAt time.Time
        UpdatedAt time.Time
        DeletedAt gorm.DeletedAt `gorm:"index"`
        Name string
      }
      ```

      -> 일반 필드의 경우, 다음과 같이 "embeded" 태그를 포함할 수 있음

      ```go
      type Author struct {
        Name  string
        Email string
      }
      
      type Blog struct {
        ID      int
        Author  Author `gorm:"embedded"` // include Name, Email fields
        Upvotes int32
      }
      // equals
      type Blog struct {
        ID    int64
        Name  string
        Email string
        Upvotes  int32
      }
      ```

      -> "embededPrefix" 태그를 사용하여 내장된 필드의 DB 이름에 prefix를 추가할 수 있음

      ```go
      type Blog struct {
        ID      int
        Author  Author `gorm:"embedded;embeddedPrefix:author_"` // include AuthorName, AuthorEmail fields
        Upvotes int32
      }
      // equals
      type Blog struct {
        ID          int64
          AuthorName  string
          AuthorEmail string
        Upvotes     int32
      }
      ```

    * Fields Tags

      * Tags : 모델을 선언할 때 사용할 수 있는 옵션. camelCase를 선호함.

        | tag name               | description                                                  |
        | ---------------------- | ------------------------------------------------------------ |
        | column                 | column 명                                                    |
        | type                   | column의 데이터 유형<br />-> ex) bool / int / uint / float / string / time / bytes / ...<br />-> not null, size, autoIncrement와 같은 태그들과 함께 사용 가능함 |
        | size                   | column 크기/길이<br />-> ex) size:256                        |
        | primaryKey             | column을 primary key로 지정                                  |
        | unique                 | column을 unique로 지정                                       |
        | default                | column의 기본값을 지정                                       |
        | precision              | column의 자릿수 설정                                         |
        | scale                  | column 크기 지정                                             |
        | not null               | column을 NOT NULL로 지정                                     |
        | autoIncrement          | column을 auto incrementable로 지정                           |
        | autoIncrementIncrement | 자동 increment step, 연속(successive) 열 값 사이의 간격을 제어 |
        | embedded               | 항목을 넣음                                                  |
        | embeddedPrefix         | 항목에 prefix를 붙여서 넣음                                  |
        | autoCreateTime         | track current time 생성 시 int 필드의 경우 unix seconds를 추적함. nano/milli 값을 사용하여 unix nano/milli 초를 추적함.<br />-> ex) autoCreateTime:nano |
        | autoUpdateTime         | track current time 업데이트 시 int 필드의 경우 unix seconds를 추적함. nano/milli 값을 사용하여 unix nano/milli 초를 추적함.<br />-> ex) autoCreateTime:milli |
        | index                  | 옵션을 사용하여 인덱스를 생성<br />-> 여러 필드에 동일한 이름을 사용하여 복합 인덱스 생성 |
        | uniqueIndex            | 유일한 index를 생성함                                        |
        | check                  | check 제약조건을 생성함                                      |
        | <-                     | 필드의 쓰기 권한을 추가<br />ex) <-:create : create-only 필드 / <-:update : update-only 필드 / <-:false : 쓰기 권한 없음 / <- : create, update 권한 |
        | ->                     | 필드의 읽기 권한을 추가<br />ex) ->:false : 읽기 권한 없음   |
        | -                      | 해당 필드 ignore<br />ex) - : 읽기/쓰기 권한 없음            |
        | comment                | migration할 때 필드에 comment 추가                           |

