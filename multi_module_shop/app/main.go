package main

import (
    "fmt"
    "github.com/ojlogic/GO/multi_module_shop/orders"
    "github.com/ojlogic/GO/multi_module_shop/users"
)

func main() {
    fmt.Println("Main App - Multi Module Shop")

    // Добавляем пользователя
    users.AddUser("John Doe")

    // Обрабатываем заказ
    orders.ProcessOrder(101)
}