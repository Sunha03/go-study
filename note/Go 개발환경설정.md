# Go 개발환경설정

## Go 설치

1. 패키지 파일로 설치

   https://golang.org/dl

   -> MacOS로 다운 & 설치

   ```bash
   # go version
   $ go version
   >> go version go1.15.5 darwin/amd64
   ```

   - GO 환경설정

   ```bash
   # go 환경설정 확인
   $ go env
   >> --
   GO111MODULE=""
   GOARCH="amd64"
   GOBIN=""
   GOCACHE="/Users/sunhapark/Library/Caches/go-build"
   GOENV="/Users/sunhapark/Library/Application Support/go/env"
   GOEXE=""
   GOFLAGS=""
   GOHOSTARCH="amd64"
   GOHOSTOS="darwin"
   GOINSECURE=""
   GOMODCACHE="/Users/sunhapark/go/pkg/mod"
   GONOPROXY=""
   GONOSUMDB=""
   GOOS="darwin"
   GOPATH="/Users/sunhapark/go"
   GOPRIVATE=""
   GOPROXY="https://proxy.golang.org,direct"
   GOROOT="/usr/local/go"
   GOSUMDB="sum.golang.org"
   GOTMPDIR=""
   GOTOOLDIR="/usr/local/go/pkg/tool/darwin_amd64"
   GCCGO="gccgo"
   AR="ar"
   CC="clang"
   CXX="clang++"
   CGO_ENABLED="1"
   GOMOD=""
   CGO_CFLAGS="-g -O2"
   CGO_CPPFLAGS=""
   CGO_CXXFLAGS="-g -O2"
   CGO_FFLAGS="-g -O2"
   CGO_LDFLAGS="-g -O2"
   PKG_CONFIG="pkg-config"
   GOGCCFLAGS="-fPIC -m64 -pthread -fno-caret-diagnostics -Qunused-arguments -fmessage-length=0 -fdebug-prefix-map=/var/folders/n_/9hrsjfpd1bngz7c3nyhpfnqr0000gn/T/go-build978047049=/tmp/go-build -gno-record-gcc-switches -fno-common"
   --
   ```

   - GOROOT 설정

   ```bash
   # GOROOT 확인
   $ go env GOROOT
   >> /usr/local/go
   ```

   - GOPATH 설정

   ```bash
   # GOPATH 확인
   $ go env GOPATH
   >> /Users/sunhapark/go
   ```

   - GOPATH 수정방법

     1) .zshrc or bash_profile에 환경변수 직접 선언

     ```bash
     
     ```
   # terminal
     $ vi $HOME/.bash_profile
     export GOPATH=$HOME/GoStudy
     $ source ~/.bash_profile
     ```
     
     2) vscode 내 언어 설정(setting.json)에 gopath 지정
     
     
     ```

2. homebrew로 설치

   ```bash
   # GO install
   $ brew install go
   
   # GO version
   $ go version
   >> go version go1.15.5 darwin/amd64
   ```



## vscode 설치

1. vscode 설치

   ```bash
   # vscode 설치
   $ brew cask install visual-studio-code
   ```

2. Go extension 설치

   -> Extension > "go extension" 검색 후 go 설치

![image-20201123010055990](/Users/sunhapark/Library/Application Support/typora-user-images/image-20201123010055990.png?raw=true)

3. shell command 설치

   1) 검색(command+shift+p) > "shell command" 검색 > 'Install 'code' command in PATH' 클릭 후 설치

   2) 원하는 경로에 들어가서 "code ."를 입력하면 바로 vscode 실행



## Go 실행 테스트

1. Go가 설치된 폴더 - bin폴더(생성) - test.go 파일 생성

   ```go
   package main
   func main() {
       println("Test")
   }
   ```

2. go 컴파일 및 실행 (.exe 생성 안함)

   ```bash
   # terminal
   $ go run test.go
   ```

3. go build (.exe 생성)

   ```bash
   # terminal
   $ go build test.go
   ```

 ** 실행 후 test.go 삭제



## Workspace 폴더

- 일반적으로 Workspace 폴더(작업 폴더)를 생성함. Workspace 폴더 안에는 3개의 서브 디렉토리(bin, pkg, src) 생성함.

  -> Workspace : /Users/sunhapark/go

  - bin : 실행 가능한 커맨트 파일. 소스 파일(패키지)를 컴파일하여 실행 파일(바이너리)이 생성되는 디렉토리
  - pkg : 패키지 객체들. 패키지를 컴파일하여 라이브러리 파일이 생성되는 디렉토리
  - src : Go 소스파일들의 집합(일반적으로 src 하위 1개 폴더 = 1개의 프로젝트). 소스 파일이 저장되는 디렉토리

- Workspace 폴더 생성 후, GOPATH 환경변수에 이 Workspace 폴더 경로를 추가(GOPATH 설정 방법은 위의 내용 참고)

- GO 환경변수

  1. GOROOT : Go가 설치된 디렉토리(/bin : Go 실행파일, /pkg : Go 표준 패키지). Go 설치 시, 자동으로 시스템 환경변수로 설정됨.

  2. GOPATH : 표준 패키지 이외의 3rd Party 패키지나 사용자 정의 패키지가 여기에 존재. 복수 개의 경로를 지정한 경우, 3rd Party 패키지는 처음 경로에 설치됨

     ** GOPATH는 일종의 작업 공간(workspace)임. 따라서 다른 프로젝트를 작업해야 할 때에는 GOPATH를 바꿔줘야 함.(-> GOPATH 수정은 위의 내용 참고)