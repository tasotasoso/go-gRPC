# gRPCサンプルコード

## 実行方法
イメージの作成
```
docker image build -t client ./client
docker image build -t server ./server
```
コンテナの起動
```
docker-compose up -d
```
申し込み画面へのアクセス
```
http://localhost:50050
```

## ライセンス
このコードは下記コードを参考に作成しています。
MITライセンスに保護されているため、規約を確認してください。
https://github.com/vvatanabe/go-grpc-basics