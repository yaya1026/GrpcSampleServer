# GrpcSampleServer

Goで書いたgRPCのサンプル

protoは[このリポジトリ](https://github.com/yaya1026/ProtocolBuffersSample)をsubmodule化しています。


## protobufのインストールについて
[Protocol Buffers](https://developers.google.com/protocol-buffers/)

MacでのインストールはHomebrew経由
https://formulae.brew.sh/formula/protobuf
```
$ brew install protobuf
```
その他のOSは[リリースページ](https://github.com/protocolbuffers/protobuf/releases/tag/v3.11.2)からお願いします。


[こちら](https://grpc.io/docs/languages/go/quickstart/)を参照しました。

## protoから更新する手法
scriptで生成。
```
./generated.sh
```


## 起動

通信には[gRPCurl](https://github.com/fullstorydev/grpcurl)で確認

```
$ go run main.go

// 別Terminalで呼び出し
$ grpcurl -plaintext localhost:8080 BookService.GetBooks
```

```json
{
  "books": [
    {
      "id": "1",
      "title": "Go",
      "author": "Yaya",
      "isbn": "978-4-86501-422-8",
      "overview": "いいよ"
    }
  ]
}
```


//
protoc-gen-doc使えないなう