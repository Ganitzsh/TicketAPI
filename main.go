package main

import (
	"fmt"

	"github.com/SQLApi/builder"
	"github.com/SQLApi/handlers"
)

func main() {
	fmt.Println(builder.NewMySQLQuery("user", "id").GetAll(nil))
	fmt.Println(builder.NewMySQLQuery("user", "id").GetAll([]string{"col1", "col2"}))
	fmt.Println(builder.NewMySQLQuery("user", "id").GetByID(3, nil))
	fmt.Println(builder.NewMySQLQuery("user", "id").GetByID(42, []string{"col1", "col2", "col3"}))
	fmt.Println(builder.NewMySQLQuery("user", "id").GetWhere(nil, nil, nil))
	fmt.Println(builder.NewMySQLQuery("user", "id").GetWhere(nil, []string{"col1", "col2", "col3"}, []string{"bonjour", "bonsoir", "4"}))
	fmt.Println(builder.NewMySQLQuery("user", "id").GetInnerJoin(nil, "col1", "col2", "col3"))
	fmt.Println(builder.NewMySQLQuery("user", "id").GetInnerJoin([]string{"col1", "col2"}, "table2", "col1", "col2"))
	api, err := NewAPIFromFile("./config.json")
	if err != nil {
		panic(err)
	}
	h := api.DBHandler.(*handlers.MySQLHandler)
	res, err := h.GetAll("Categories")
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
	res, err = h.GetBy("Categories", []string{"CategoryName", "Description"}, []string{"CategoryID"}, []string{"1"})
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
	api.Run()
}
