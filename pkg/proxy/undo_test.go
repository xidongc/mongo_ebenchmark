/*
 * mongodb_ebenchmark - Mongodb grpc proxy benchmark for e-commerce workload (still in dev)
 * Copyright (c) 2020 - Chen, Xidong <chenxidong2009@hotmail.com>
 *
 * All rights reserved.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 *
 */

package proxy

import (
	"github.com/xidongc/mongo_ebenchmark/model/order/orderpb"
	"github.com/xidongc/mongo_ebenchmark/pkg/cfg"
	"testing"
)

// Test undo insert
func TestUndoInsert(t *testing.T) {
	item1 := orderpb.Item{
		ProductId:   "pid12345",
		Name:        "TestItem",
		Quantity:    4,
		Amount:      5,
		Description: "test description",
	}
	item2 := orderpb.Item{
		ProductId:   "pid12345",
		Name:        "TestItem",
		Quantity:    4,
		Amount:      5,
		Description: "test description",
	}
	items := make([]interface{}, 2)
	items[0] = item1
	items[1] = item2
	insertParam := InsertParam{
		Docs: items,
		Amp:  cfg.MicroAmplifier(),
	}
	removeParam := UndoInsert(&insertParam)
	if len(removeParam) != 2 {
		t.Error("error")
	}
	t.Log(*removeParam[0])
	t.Log(*removeParam[1])
}
