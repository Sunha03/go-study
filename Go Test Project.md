# Go Test Project

* GOPATH에 폴더 추가

  1. bin, src, pkg 폴더 3개 추가

     ```
     mkdir bin
     mkdir src
     mkdir pkg
     ```

  2. /src에 프로젝트 폴더 생성

     ```
     mkdir test
     ```

  3. 프로젝트 폴더 안에 main.go 파일 생성

     ```
     touch main.go
     ```

  4. 프로젝트 폴더 안에서 vscode 실행

     (터미널 > 프로젝트 폴더 위치 > "code ." 실행)

     ```
     code .
     ```

  5. vscode 오른쪽 하단에 팝업 install

     ![image-20201123011907655](/Users/sunhapark/Library/Application Support/typora-user-images/image-20201123011907655.png)

     ![image-20201123011959750](/Users/sunhapark/Library/Application Support/typora-user-images/image-20201123011959750.png)

     ![image-20201123012047390](/Users/sunhapark/Library/Application Support/typora-user-images/image-20201123012047390.png)

  6. main.go 파일에 코드 입력

     ```
     #main.go
     package main
     
     import "fmt"
     
     func main() {
         fmt.Plintln("Hello World")
     }
     ```

  7. vscode Terminal에서 main.go 실행

     ```
     #main.go 실행
     go run main.go
     ```

     ![image-20201123012345480](/Users/sunhapark/Library/Application Support/typora-user-images/image-20201123012345480.png)

     