# ビルドステージ
FROM golang:alpine3.19 as builder

# ソースコードをコピー
COPY ./main.go /workspace/main.go

# ワークディレクトリを設定
WORKDIR /workspace

# main.goをビルドして実行ファイルを生成
RUN go build -o main ./main.go

# 最終ステージ
FROM alpine:3.19

# 必要なパッケージのみをインストール
RUN apk --update --no-cache add bash=5.2.21-r0 stress-ng=0.17.01-r0

# ビルドステージから実行ファイルをコピー
COPY --from=builder /workspace/main /main

# コンテナ起動時に実行されるコマンド
CMD ["./main"]
