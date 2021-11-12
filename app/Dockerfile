FROM golang:1.17-alpine

ENV GO111MODULE=on

#アップデートとgitのインストール
RUN apk update && apk add git

#コンテナ上でappディレクトリの作成
RUN mkdir /go/src/app
#ワーキングディレクトリの設定
WORKDIR /go/src/app
#ホストのファイルをコンテナの作業ディレクトリに移行
ADD . /go/src/app

RUN go get -u github.com/go-sql-driver/mysql

#realizeするときはGO111MODULE=offにする
RUN GO111MODULE=off go get -u github.com/oxequa/realize

CMD ["realize", "start"]