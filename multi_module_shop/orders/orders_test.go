package orders

import "testing"

// Тестирование структуры Order
func TestOrder(t *testing.T) {
    order := Order{ID: 1, Product: "Laptop", Quantity: 1}

    if order.ID != 1 {
        t.Errorf("Expected ID 1, got %d", order.ID)
    }

    if order.Product != "Laptop" {
        t.Errorf("Expected Product 'Laptop', got %s", order.Product)
    }

    if order.Quantity != 1 {
        t.Errorf("Expected Quantity 1, got %d", order.Quantity)
    }
}
