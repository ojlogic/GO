package orders

import "testing"

func TestOrder(t *testing.T) {
    order := Order{ID: 1, Product: "Laptop", Quantity: 1}
    if order.ID != 1 {
        t.Errorf("Expected ID 1, got %d", order.ID)
    }
}
