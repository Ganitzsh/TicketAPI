package main

import (
	"fmt"

	"github.com/SQLApi/builder"
)

func main() {
	fmt.Println(builder.NewQuery("user", "id").GetAll(nil))
	fmt.Println(builder.NewQuery("user", "id").GetAll([]string{"col1", "col2"}))
	fmt.Println(builder.NewQuery("user", "id").GetByID(3, nil))
	fmt.Println(builder.NewQuery("user", "id").GetByID(42, []string{"col1", "col2", "col3"}))
	fmt.Println(builder.NewQuery("user", "id").GetWhere(nil, nil, nil))
	fmt.Println(builder.NewQuery("user", "id").GetWhere(nil, []string{"col1", "col2", "col3"}, []string{"bonjour", "bonsoir", "4"}))
	fmt.Println(builder.NewQuery("user", "id").GetInnerJoin(nil, "col1", "col2", "col3"))
	fmt.Println(builder.NewQuery("user", "id").GetInnerJoin([]string{"col1", "col2"}, "table2", "col1", "col2"))
	api, err := NewAPIFromFile("./config.json")
	if err != nil {
		panic(err)
	}
	api.Run()
}
