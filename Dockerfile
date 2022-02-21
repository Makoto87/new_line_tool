# goのイメージを取得、バイナリファイル作るまで必要
FROM golang:1.17 as builder
# アップデートとgitのインストール
# RUN apk update && apk add git

# クロスコンパイル用に環境変数の設定
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
# ワーキングディレクトリの設定・作成
WORKDIR /go/go_linetool

# ホストのファイルをコンテナ(goディレクトリ)に移行
COPY go.mod go.sum ./
RUN go mod download
# キャッシュのコピー?
COPY . .
# ビルド。オプションoでファイル名を指定(ファイル名: app)
RUN go build -o app

# バイナリファイルを保存する環境
FROM alpine
# no-cacheで不要なファイルを削除する
# addでca-certificatesパッケージを取得
RUN apk add --no-cache ca-certificates

# builderからコピー、appディレクトリにファイルを入れる
COPY --from=builder /go/go_linetool/app /app
CMD ["/app"]