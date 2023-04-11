package main

import (
	"fmt"
	"reflect"
)

type order struct {
	ordId      int
	customerId int
	username   string
	money      string
}

func genSql(o interface{}) string {
	var ot = reflect.TypeOf(o)
	var ov = reflect.ValueOf(o)
	if ot.Kind() != reflect.Struct {
		return "wrong type, the type of input must be reflect.Struct"
	}
	sql := fmt.Sprintf("insert into %s values(", ot.Name())

	filedNum := ov.NumField()

	for i := 0; i < filedNum; i++ {
		switch ov.Field(i).Kind() {
		case reflect.String:
			if i == 0 {
				sql = fmt.Sprintf("%s\"%s\"", sql, ov.Field(i).String())
			} else {
				sql = fmt.Sprintf("%s \"%s\"", sql, ov.Field(i).String())
			}
		case reflect.Int:
			if i == 0 {
				sql = fmt.Sprintf("%s%d", sql, ov.Field(i).Int())
			} else {
				sql = fmt.Sprintf("%s %d", sql, ov.Field(i).Int())
			}
		default:
			fmt.Println("wrong type")
			return ""
		}
	}
	sql = fmt.Sprintf("%s)", sql)
	return sql
}

func main() {
	order := order{
		ordId:      456,
		customerId: 56,
		username:   "dsds",
		money:      "2000",
	}

	fmt.Println(genSql(order))
}
