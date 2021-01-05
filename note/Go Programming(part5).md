# Go Programming(part5)

## 1. Go defer

* defer : 특정 문장 or 함수를 나중에 (deferf를 호출하는 함수가 리턴하기 직전에) 실행하게 함

* 일반적으로는 C#, Java 같은 언어에서의 finally 블럭처럼 마지막에 Clean-up 작업을 위해 사용됨

  ```go
  import "os"
  
  func main() {
  	f, err := os.Open("1.txt")
  	if err != nil {
  		panic(err)
  	}
  	
  	defer f.Close()	// main() 마지막에 파일 close 실행
  	
  	bytes := make([]byte, 1024)	// 파일 읽기
  	f.Read(bytes)
  	println(len(bytes))
  }
  ```

  -> 파일을 Open한 후 항상 파일을 close 함

