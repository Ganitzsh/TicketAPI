package main

import (
	"fmt"

	"github.com/SQLApi/builder"
	"github.com/SQLApi/handlers"
	"github.com/SQLApi/models"
)

type TestType struct {
	ID        int    `db:"_id"`
	FirstName string `db:"first_name"`
	Size      float64
}

type Customer struct {
	CustomerID   string
	CompanyName  string
	ContactName  string
	ContactTitle string
	Address      string
	City         string
	PostalCode   string
	Country      string
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

	// c := Customer{
	// 	CustomerID:  "BLABL",
	// 	CompanyName: "Les Beaufs",
	// 	ContactName: "Maurice LeBlanc",
	// 	Address:     "3 Rue des plaines",
	// 	City:        "Nice",
	// 	PostalCode:  "89780",
	// 	Country:     "France",
	// }

	fmt.Println(builder.NewMySQLQuery("user", "id").Insert(tmp))
	fmt.Println(builder.NewMySQLQuery("user", "id").Delete(tmp))
	fmt.Println(builder.NewMySQLQuery("user", "id").Update(newTmp, tmp))
	api, err := NewAPIFromFile("./config.json")
	if err != nil {
		panic(err)
	}
	h := api.DBHandler.(*handlers.MySQLHandler)
	// res, err := h.GetAll("Categories")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(res)
	// res, err = h.GetBy("Categories", []string{"CategoryName", "Description"}, []string{"CategoryID"}, []string{"1"})
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(res)
	// res, err = h.Delete("Customers", c)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(res)
	// res, err = h.Insert("Customers", c)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(res)
	// new := c
	// new.PostalCode = "75018"
	// res, err = h.Update("Customers", new, c)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(res)
	res, err := h.GetAll(models.Ticket{}, "ticket")
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
	// api.Run()
}
