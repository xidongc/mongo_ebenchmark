package proxy

import (
	"github.com/xidongc/mongodb_ebenchmark/model/order/orderpb"
	"testing"
)

// Test undo insert
func TestUndoInsert(t *testing.T) {
	item1 := orderpb.Item{
		ProductId: "pid12345",
		Name: "TestItem",
		Quantity: 4,
		Amount: 5,
		Description: "test description",
	}
	item2 := orderpb.Item{
		ProductId: "pid12345",
		Name: "TestItem",
		Quantity: 4,
		Amount: 5,
		Description: "test description",
	}
	items := make([]interface{}, 2)
	items[0] = item1
	items[1] = item2
	insertParam := InsertParam{
		Docs: items,
		Amp: MicroAmplifier(),
	}
	removeParam := UndoInsert(&insertParam)
	if len(removeParam) != 2 {
		t.Error("error")
	}
	t.Log(*removeParam[0])
	t.Log(*removeParam[1])
}
