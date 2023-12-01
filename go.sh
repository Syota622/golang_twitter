# ターミナルで実行したもの

# go
go mod init golang_twitter
go mod tidy

# docker
docker compose build
docker compose up

# 確認
curl http://localhost:8080/health_check
