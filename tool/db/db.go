package db

import(
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"github.com/spf13/viper"
	"fmt"
)

type Struct DbConn{
	Db *sql.DB
}

func InitDbConn() (*DbConn, error) {
	
}