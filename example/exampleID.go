package main

import (
	"database/sql"
	"github.com/akula410/helper"
)

import (
	c "github.com/akula410/connect"
)

var MySql c.MySql

func init(){
	MySql.DBName = "DB_NAME"
	MySql.Host = "localhost"
	MySql.User = "root"
	MySql.Password = ""
	MySql.Port = "3306"
	MySql.Charset = "utf8"
	MySql.InterpolateParams = true
	MySql.MaxOpenCoons = 10
}

func exampleID(table string, field string, amount int) int {
	helper.ID.Conn = func() *sql.DB {
		return MySql.Connect()
	}

	helper.ID.ConnClose = func() {
		MySql.Close()
	}

	return helper.ID.GetKey(table, field, amount)
}
