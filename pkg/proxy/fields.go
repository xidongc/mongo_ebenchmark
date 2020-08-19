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
	"github.com/xidongc-wish/mgo/bson"
	"github.com/xidongc/mongo_ebenchmark/pkg/cfg"
)

type Empty struct{}

type FindAndModifyMode int

// Query param for upper services
type QueryParam struct {
	Filter      bson.M
	Fields      bson.M
	Limit       int64
	Skip        int64
	Sort        []string
	Distinctkey string
	FindOne     bool
	UsingIndex  []string
	Amp         cfg.Amplifier
}

// Insert param for upper services
type InsertParam struct {
	Docs []interface{}
	Amp  cfg.Amplifier
}

// Remove param for upper services
type RemoveParam struct {
	Filter bson.M
	Amp    cfg.Amplifier
}

// Update param for upper services
type UpdateParam struct {
	Filter bson.M
	Update bson.M
	Upsert bool
	Multi  bool
	Amp    cfg.Amplifier
}

// FindAndModify param for upper services
type FindModifyParam struct {
	Filter   bson.M
	Desired  bson.M
	Mode     FindAndModifyMode
	SortRule []string
	Fields   bson.M
	Amp      cfg.Amplifier
}

// Aggregate param for upper services
type AggregateParam struct {
	Pipeline bson.M
}
