package infra

import (
	"database/sql"
	"os"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

type DB bun.DB

var (
	Host string
	Port string
	Name string
	User string
	Pass string
)

func init() {
	Host = os.Getenv("DB_HOST")
	Port = os.Getenv("DB_PORT")
	Name = os.Getenv("DB_NAME")
	User = os.Getenv("DB_USER")
	Pass = os.Getenv("DB_PASSWORD")
	if Host == "" || Port == "" || Name == "" || User == "" || Pass == "" {
		panic("DB connection information is not set")
	}
}

func NewDB() *bun.DB {
	db := bun.NewDB(Conn(), pgdialect.New())
	return db
}

func Conn() *sql.DB {
	conn, err := sql.Open("postgres", "postgres://"+User+":"+Pass+"@"+Host+":"+Port+"/"+Name+"?sslmode=disable")
	if err != nil {
		panic(err)
	}
	return conn
}
