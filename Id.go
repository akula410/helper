/**
Helper.ID.Conn = func() *sql.DB{
	return db.MySql.Connect()
}

Helper.ID.ConnClose = func(){
	db.MySql.Close()
}

id := Helper.ID.GetKey("go_filters", "filter_id", 1)
 */
package Helper

import (
	"database/sql"
	"fmt"
	"strconv"
	"sync"
)

type id struct {
	sync.Mutex
	fields map[string]map[string]int
	Conn func()*sql.DB
	ConnClose func()
}

var ID id

func (a *id)GetKey(table string, field string, count int)int{
	var ok bool
	var id int
	ID.Lock()
	if ID.fields == nil {
		ID.fields = make(map[string]map[string]int, 0)
	}

	_, ok = ID.fields[table]
	if !ok {
		ID.fields[table] = make(map[string]int, 0)
	}


	id, ok = ID.fields[table][field]

	if !ok {
		id = ID.lastId(table, field)

	}
	ID.fields[table][field] = id+count

	ID.Unlock()

	return id
}

func (a *id)lastId(table string, field string)int{
	textSql := fmt.Sprintf("SELECT %s FROM %s ORDER BY %s DESC LIMIT 1", field, table, field)
	rows, err := ID.Conn().Query(textSql)
	if err != nil {
		panic(err)
	}

	columns, err := rows.Columns()

	if err != nil {
		panic(err.Error())
	}

	count := len(columns)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	fields := make(map[string]interface{})

	for rows.Next() {
		for i := range columns {
			valuePtrs[i] = &values[i]
		}

		err := rows.Scan(valuePtrs...)
		if err != nil {
			panic(err.Error())
		}

		for i, col := range columns {

			var v interface{}

			val := values[i]

			b, ok := val.([]byte)

			if ok {
				v = string(b)
			} else {
				v = val
			}

			fields[col] = v
		}
		break
	}

	ID.ConnClose()

	result, ok := fields[field]

	if ok {
		id, err := strconv.Atoi(fmt.Sprintf("%v", result))
		if err != nil {
			panic(err)
		}
		return id
	}else{
		return 0
	}

}