# syntax=docker/dockerfile:1
FROM golang:1.19.4-alpine3.17

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

RUN mkdir -p /pkg/backend/sql
RUN mkdir -p /pkg/frontend/ui
RUN mkdir -p /cmd/suggestApp

#後ほどcmd/main.goからcmd/suggestApp/main.goにディレクトリ構成を変更
COPY cmd/main.go cmd/suggestApp 

#COPY pkg/backend /pkg/backend
#COPY pkg/frontend /pkg/frontend (フロントエンド実装時にコメントアウト解除)

#実行ファイルを/app/cmd/suggestAppディレクトリに配置する
RUN go build -o /cmd/suggestApp

#リッスンポートは8080を使用する
EXPOSE 8080






