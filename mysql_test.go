package mysql_test

import (
	"os"

	"github.com/comame/mysql-go"
)

func Example() {
	os.Setenv("MYSQL_USER", "user")
	os.Setenv("MYSQL_PASSWORD", "pass")
	os.Setenv("MYSQL_DATABASE", "db")
	os.Setenv("MYSQL_HOST", "localhost")

	mysql.Initialize()
	db := mysql.Conn()

	db.Begin()
}
