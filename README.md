# Go tutorial with database
I first found this `sqlc` stuff at [[Backend #4] Generate CRUD Golang code from SQL | Compare db/sql, gorm, sqlx & sqlc](https://youtu.be/prh0hTyI1sU) (hence the repository appellation). But since it only covered `sqlc` for `PostgreSQL`, I later opted to [Getting started with MySQL](https://docs.sqlc.dev/en/stable/tutorials/getting-started-mysql.html) - the official tutorial of `sqlc` on `MySQL`.

## Build and run
First, you need to login to `MySQL`, creating a database named `tutorial`, and call the script `./tutorial/schema.sql`. After that, do
```
go run app.go
```
to execute the `Go` app.
