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
	log "github.com/sirupsen/logrus"
	"github.com/xidongc-wish/mgo/bson"
	"reflect"
	"strings"
)

// UndoInsert generate removeParam based on given insert param
// Currently nested struct is not supported
//
// Please refer proxy.InsertParam, proxy.RemoveParam for more details
func UndoInsert(param *InsertParam) (params []*RemoveParam) {
	if param == nil {
		log.Error("received nil input")
		return
	}
	for _, doc := range param.Docs {
		var removeFilter = bson.M{}
		objType := reflect.TypeOf(doc)
		objVal := reflect.ValueOf(doc)
		for i := 0; i < objVal.NumField(); i++ {
			if objVal.Field(i).CanInterface() {
				if objVal.Field(i).Kind() == reflect.Struct {
					// TODO support nested struct
					log.Warning("by pass nested struct")
				} else {
					lowerName := strings.ToLower(objType.Field(i).Name)
					removeFilter[lowerName] = objVal.Field(i).Interface()
				}
			}
		}

		params = append(params, &RemoveParam{
			Filter: removeFilter,
			Amp:    param.Amp,
		})
	}
	return
}
