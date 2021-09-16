package main

import(
  "context"
  "database/sql"
  "log"
  "os"
  "reflect"
  "github.com/phogbinh/go-db-prh0hTyI1sU/tutorial"
  _ "github.com/go-sql-driver/mysql"
)

func listAllAuthors(ctx context.Context, queries *tutorial.Queries) error {
  authors, err := queries.ListAuthors(ctx)
  if err != nil {
    return err
  }
  log.Println(authors)
  return nil
}

func createAnAuthor(ctx context.Context, queries *tutorial.Queries) (int64, error) {
  author := tutorial.CreateAuthorParams {
    Name: "Brian Kernighan",
    Bio: sql.NullString {
      String: "Co-author of The C Programming Language and The Go Programming Language",
      Valid: true,
    },
  }
  result, err := queries.CreateAuthor(ctx, author)
  if err != nil {
    return 0, err
  }
  insertedAuthorId, err := result.LastInsertId()
  if err != nil {
    return 0, err
  }
  log.Println(insertedAuthorId)
  return insertedAuthorId, nil
}

func run() error {
  ctx := context.Background() // [c]on[t]e[x]t
  db, err := sql.Open("mysql", os.Getenv("MYSQL_USER") + ":" + os.Getenv("MYSQL_PASSWORD") + "@/" + os.Getenv("MYSQL_DATABASE_NAME")) // [d]ata[b]ase, [err]or
  if err != nil {
    return err
  }
  queries := tutorial.New(db)
  err = listAllAuthors(ctx, queries)
  if err != nil {
    return err
  }
  insertedAuthorId, err := createAnAuthor(ctx, queries)
  if err != nil {
    return err
  }
  fetchedAuthor, err := queries.GetAuthor(ctx, insertedAuthorId)
  if err != nil {
    return err
  }
  log.Println(reflect.DeepEqual(insertedAuthorId, fetchedAuthor.ID))
  return nil
}

func main() {
  if err := run(); err != nil {
    log.Fatal(err)
  }
}
