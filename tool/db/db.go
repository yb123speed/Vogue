package db

import(
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"github.com/spf13/viper"
	"fmt"
)

type DbConn struct{
	Db *sql.DB
}

func InitDbConn() (*DbConn, error) {
	dbconn := viper.GetString("Database.DbConn")
	db, err := sql.Open("mysql", dbconn)
	if err != nil {
		return nil, fmt.Errorf("get db connection error")
	}
	maxOpenConns := viper.GetInt("Database.MaxOpenConn")
	maxIdleConns := viper.GetInt("Database.MaxIdleConn")
	maxLifetime := viper.GetInt("Database.MaxLifetime")

	// TODO Configure DB Service

	err = db.Ping()

	if err != nil {
		return nil, fmt.Errorf("connect to db error")
	} else {
		dbconn := &DbConn{db}
		return dbconn, nil
	}
}

func (c *DbConn) Insert(sql string, args ...interface{}) (lastInsertId int64, err error) {
	res, err := c.Db.Exec(sql, args...)
	if err != nil {
		return
	}
	return res.LastInsertId()
}

// TODO  code other db actions