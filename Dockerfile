# 使用するGoのバージョンを指定
FROM golang:1.18 AS builder

# ワーキングディレクトリを設定
WORKDIR /app

# モジュールファイルをコピー
COPY go_example/go.mod go_example/go.sum ./

# モジュールをダウンロード
RUN go mod download

# ソースコードをコピー
COPY go_example/. .

# アプリケーションをビルド
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# 実行ステージ
FROM alpine:latest  
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# ビルドステージで作成したバイナリをコピー
COPY --from=builder /app/main .
COPY go_example/templates ./templates
COPY go_example/css ./css

# ポート8040を開放
EXPOSE 8040

# アプリケーションを実行
CMD ["./main"]
