# GO 환경설정

## Go 설치

- 패키지 파일로 설치

  https://golang.org/dl

  -> MacOS로 다운 & 설치

  ```
  #go version
  go version
  -- go version go1.15.5 darwin/amd64
  ```

   - GO 환경설정

     ```
     #go 환경설정 확인
     --
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

     ```
     #GOROOT 확인
     go env GOROOT
     -- /usr/local/go
     ```

  - GOPATH 설정

    ```
    #GOPATH 확인
    go env GOPATH
    -- /Users/sunhapark/go
    ```

    - GOPATH 수정방법
      1. .zshrc or bash_profile에 환경변수 직접 선언
      2. vscode 내 언어 설정(setting.json)에 gopath 지정

- homebrew로 설치

```
#GO install
brew install go

#GO version
go version
-- go version go1.15.5 darwin/amd64
```



## vscode 설치

* vscode 설치

  ```
  #vscode 설치
  brew cask install visual-studio-code
  ```

* Go extension 설치

  -> Extension > "go extension" 검색 후 go 설치

  ![image-20201123010055990](/Users/sunhapark/Library/Application Support/typora-user-images/image-20201123010055990.png)

* shell command 설치

  -> 검색(command+shift+p) > "shell command" 검색 > 'Install 'code' command in PATH' 클릭 후 설치

  => 원하는 경로에 들어가서 "code ."를 입력하면 바로 vscode 실행