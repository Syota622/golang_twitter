# go
go mod init golang_twitter
go mod tidy

# docker compose
## docker composeの操作方法
docker compose build
docker compose up -d
docker compose down
docker compose down -v
docker compose down --rmi all --volumes --remove-orphans

# migrate
## マイグレーションを実行
docker compose run --rm web migrate -path /app/db/migration -database "postgres://postgres:Passw0rd@db:5432/db?sslmode=disable" up
docker compose run --rm web migrate -path /app/db/migration -database "postgres://postgres:Passw0rd@db:5432/db?sslmode=disable" down

## データベースへの接続
docker compose exec db psql -U postgres -d db
\dt
\d users

## マイグレーション生成
### sqlcを実行することで、migrationファイルからGoのコードを生成する
docker compose run --rm web sqlc generate
docker-compose exec web sqlc genera
docker-compose exec web ls /app/db/migration

## ライブラリ追加
go get github.com/lib/pq

# Redis
## go get
go get github.com/gin-contrib/sessions
go get github.com/gin-contrib/sessions/redis
go get github.com/google/uuid

## redis-cli
docker-compose exec redis redis-cli
keys *
get "key"
del "key"
