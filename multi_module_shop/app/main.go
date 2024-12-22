package main

import (
    "fmt"
    "github.com/ojlogic/GO/multi_module_shop/orders"
    "github.com/ojlogic/GO/multi_module_shop/users"
)

func main() {
    fmt.Println("Multi-module Shop")

    order := orders.Order{ID: 1, Product: "Laptop", Quantity: 1}
    user := users.User{ID: 1, Name: "John Doe", Email: "john@example.com"}

    fmt.Println("Order: ", order)
    fmt.Println("User: ", user)
}