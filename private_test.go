package bexs

import (
	"fmt"
	"testing"
)

func TestBexs_OrderStatus(t *testing.T) {
	client := New("bitrecife", "", "", true)
	result, err := client.OrderStatus("123")
	fmt.Println(result, err)
}
