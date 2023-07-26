package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/comame/readenv-go"
	_ "github.com/go-sql-driver/mysql"
)

type envType struct {
	User     string `env:"MYSQL_USER"`
	Password string `env:"MYSQL_PASSWORD"`
	Database string `env:"MYSQL_DATABASE"`
	Host     string `env:"MYSQL_HOST"`
}

var connections = make(map[string]*sql.DB)

func Initialize() (err error) {
	defer func() {
		if recovered := recover(); recovered != nil {
			switch v := recovered.(type) {
			case string:
				err = errors.New(v)
			case error:
				err = v
			default:
				err = errors.New("Panicked")
			}
		}
	}()

	var env envType
	readenv.Read(&env)

	if err := InitializeWithCustomParameter("default", env.User, env.Password, env.Host, env.Database); err != nil {
		return err
	}

	return nil
}

func Conn() *sql.DB {
	return GetConnection("default")
}

func InitializeWithCustomParameter(connectionName, user, password, host, database string) error {
	dataSourceName := fmt.Sprintf(
		"%s:%s@(%s)/%s",
		user,
		password,
		host,
		database,
	)

	conn, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return err
	}

	conn.SetConnMaxLifetime(3 * time.Minute)
	conn.SetMaxOpenConns(10)
	conn.SetMaxIdleConns(10)

	connections[connectionName] = conn

	return nil
}

func GetConnection(connectionName string) *sql.DB {
	conn, ok := connections[connectionName]
	if !ok {
		panic("Call mysql.Initialize() first.")
	}
	return conn
}
