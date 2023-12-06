# docker compose
## docker composeの操作方法
docker compose build
docker compose up -d
docker compose down
docker compose down -v
docker compose down --rmi all --volumes --remove-orphans

# migrate
## マイグレーションを実施
docker compose up -d
docker compose exec web migrate -path /app/db/migration -database "postgres://postgres:Passw0rd@db:5432/db?sslmode=disable" up

## データベースへの接続
docker compose exec db psql -U postgres -d db
\dt
\d users

## マイグレーションファイルの確認
docker-compose exec web ls /app/db/migration
