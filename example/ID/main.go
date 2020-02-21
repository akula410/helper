package main

import (
	"database/sql"
	"fmt"
	"github.com/akula410/helper"
)

import (
	c "github.com/akula410/connect"
)

//To reserve table identifiers
/**
id - current last identifier
amount - number of records to be reserved
*/
func main() {
	var id int
	id = exampleID("TABLE_NAME", "FIELD_NAME", 1)
	fmt.Println("last id = ", id)

	id = exampleID("ws_dir_manifest_ews", "ew_id", 1250)
	fmt.Println("last id = ", id)
}



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
