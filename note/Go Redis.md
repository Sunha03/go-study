# Redis

* Redis : Key-Value 기반의 in-memory 데이터 저장소(=in-memory data structure store) >> NoSQL

  -> 디스크 or SSD에 데이터를 기록하는 다른 DB와 달리 메모리에 있음 >> 디스크 엑세스 시간이 필요 없음. 속도가 빠름.

  -> Key-Value 기반이기 때문에 쿼리를 따로 할 필요 없이 결과를 바로 가져올 수 있음

  -> 보통 DB, Cache, 메세지 브로커, 대기열, 인증 토큰 저장, Ranking Board 등에 사용됨

  -> 비슷한 기능으로 Memcached가 있음

  -> 기본 port 번호 = 6379 >> 이 설정은 기본사항으로 redis.conf라는 파일에 기록됨

  -> 접근 허용 ip(bind)와 password 등을 설정하여 유효한 사용자만 접근을 허락할 수 있음

  * 다양한 데이터 구조(Collection) 제공
    1. String : 단순 Key-Value 매핑 구조
    2. Lists : Array 형식의 데이터 구조. List를 사용하면 처음과 끝에 데이터를 넣고 빼는 것은 속도가 빠르지만 중간에 데이터를 삽입할 때는 어려움.
    3. Sets : 순서가 없는 Strings 데이터 집합. Sets에서는 중복된 데이터는 하나로 처리함 >> 중복에 대한 걱정을 할 필요가 없음.
    4. Sorted Sets : Sets와 같은 구조이지만, Score을 통해 순서를 정할 수 있음. Sorted Sets를 사용하면 Leaderboard와 같은 기능을 손쉽게 구현할 수 있음
    5. Hashes : Key-Value의 구조를 여러 개 가진 object 타입을 저장하기 좋은 구조임.



# Go에서 Redis 사용하기

* Client

  1. go-redis : https://github.com/go-redis/redis

     -> 장점 : 문서화가 잘 되어 있음. 쓰기 편함

     -> 단점 : 확장성이 낮음

     ```go
     package main
     
     import (
         "fmt"
     
         redis "github.com/go-redis/redis"
     )
     
     func initializeRedisClient() (*redis.Client, error) {
     
         client := redis.NewClient(&redis.Options{
             Addr:     "localhost:6379", // 접근 url 및 port
             Password: "",               // password ""값은 없다는 뜻
             DB:       0,                // 기본 DB 사용
         })
     
         _, err := client.Ping().Result()
     
         return client, err
     }
     
     func executeSomething() {
     
         client, err := initializeRedisClient()
     
         if nil != err {
         	panic(err)
         }
     
         err = client.Set("value", "2", 0).Err() // 마지막 인자는 만료 시간 Redis 데이터의 만료 시간을 지정할 수 있다.
         if err != nil {
         	panic(err)
         }
     
         val, err := client.Get("value").Result()
         if err != nil {
         	panic(err)
         }
         fmt.Println("value : ", val)
     }
     
     func main() {
     	executeSomething()
     }
     ```

     

  2. Redigo : https://github.com/gomodule/redigo

     -> 장점 : 확장성이 높음

     -> 단점 : 문서화가 잘 안되어 있음. 타입 변환 등 불편한 점이 존재함





[참고] https://brunch.co.kr/@jehovah/20
[참고] https://hodongman.github.io/2020/04/16/Go-Databases-Redis.html