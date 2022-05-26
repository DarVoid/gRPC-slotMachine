package conSQLServer

import (
	"database/sql"
	"fmt"
	"net/url"

	_ "github.com/denisenkom/go-mssqldb"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func main() {

	query := url.Values{}
	u := &url.URL{
		Scheme: "sqlserver",
		User:   url.UserPassword("<user>", "<password>"),
		Host:   fmt.Sprintf("%s:%s", "<hostIp>", "<hostport>"),
		// Path:  instance, // if connecting to an instance instead of a port
		RawQuery: query.Encode(),
	}
	sqlDB, err := sql.Open("sqlserver", u.String())
	if err != nil {
		fmt.Println(err)
	}
	gormDB, err := gorm.Open(sqlserver.New(sqlserver.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	var result uRecord
	gormDB.Table("<tablename>").Select("*").First(&result) //.Scan(&result)
	fmt.Println(result)
}

type uRecord struct {
	ID     int    `gorm:"column:Id"`
	UserId string `gorm:"column:AspNetUserId"`
}
