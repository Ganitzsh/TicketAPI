package main

import (
	"fmt"

	"github.com/SQLApi/builder"
	"github.com/SQLApi/handlers"
)

type TestType struct {
	ID        int    `db:"_id"`
	FirstName string `db:"first_name"`
	Size      float64
}

func main() {
	fmt.Println(builder.NewMySQLQuery("user", "id").GetAll(nil))
	fmt.Println(builder.NewMySQLQuery("user", "id").GetAll([]string{"col1", "col2"}))
	fmt.Println(builder.NewMySQLQuery("user", "id").GetByID(3, nil))
	fmt.Println(builder.NewMySQLQuery("user", "id").GetByID(42, []string{"col1", "col2", "col3"}))
	fmt.Println(builder.NewMySQLQuery("user", "id").GetWhere(nil, nil, nil))
	fmt.Println(builder.NewMySQLQuery("user", "id").GetWhere(nil, []string{"col1", "col2", "col3"}, []string{"bonjour", "bonsoir", "4"}))
	fmt.Println(builder.NewMySQLQuery("user", "id").GetInnerJoin(nil, "col1", "col2", "col3"))
	fmt.Println(builder.NewMySQLQuery("user", "id").GetInnerJoin([]string{"col1", "col2"}, "table2", "col1", "col2"))

	tmp := TestType{
		42,
		"Julien",
		1.54,
	}

	newTmp := TestType{
		42,
		"Jacky",
		1.60,
	}

	fmt.Println(builder.NewMySQLQuery("user", "id").Insert(tmp))
	fmt.Println(builder.NewMySQLQuery("user", "id").Delete(tmp))
	fmt.Println(builder.NewMySQLQuery("user", "id").Update(newTmp, tmp))
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
