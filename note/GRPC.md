# GRPC

### RPC (=Remote Procedure Call)

* 원래 분산환경 시스템에서 서로 다른 컴퓨터 프로그램들이 서로 다른 주소에서 서로를 호출하지만 마치 같은 주소에서 호출하는 것처럼 작동하게 하는 원격 프로시져 프로토콜

  -> 즉, 프로그램들은 서로가 누구인지 알 필요 없이 정해진 방식대로 단순히 함수 호출만 하면 됨

## GRPC

* 구글이 최초로 개발한 protobuf(a.k.a protocol buffer)라는 방식을 사용해 RPC라는 프로토콜 데이터를 주고 받는 플랫폼

* 특징

  -> 전송을 위해 HTTP/2를, 인터페이스 정의 언어로 프로토콜 버퍼를 사용함
  -> 인증, 양방향 스트리밍 및 흐름 제어, 차단 및 비차단 바인딩, 취소 및 타임아웃 등의 기능을 제공함

* 장점

  1) 빠르다
  	-> protocol buffer라는 방식은 XML과 같이 구조화된 데이터를 직렬화(serialize)하는 방식인데 압축을 해서 훨씬 빠르고, 사용법도 간단하고, 데이터의 크기도 작음. JSON 직렬화보다 최대 8배 더 빠를 수 있음

  2) 엄격한 사양 덕분에 생기는 가이드의 존재 유무 (IDL)
  	-> 



[참고] https://grpc.io/docs/languages/go/
[참고] https://devjin-blog.com/golang-grpc-server-1/
[참고] https://velog.io/@kyusung/grpc-web-example
