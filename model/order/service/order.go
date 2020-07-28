/*
 * mongodb_ebenchmark - Mongodb grpc proxy benchmark for e-commerce workload (still in dev)
 *
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
 */

package service

import (
	"context"
	"github.com/xidongc/mongodb_ebenchmark/model/order/orderpb"
	"github.com/xidongc/mongodb_ebenchmark/pkg/proxy"
)

const ns = "order"

type Service struct {
	Storage proxy.Client
	Amplifier proxy.Amplifier
}

func (s Service) New(context.Context, *orderpb.NewRequest) (*orderpb.Order, error) {
	panic("implement me")
}

func (s Service) Get(context.Context, *orderpb.GetRequest) (*orderpb.Order, error) {
	panic("implement me")
}

func (s Service) Pay(context.Context, *orderpb.PayRequest) (*orderpb.Order, error) {
	panic("implement me")
}

func (s Service) Return(context.Context, *orderpb.ReturnRequest) (*orderpb.Order, error) {
	panic("implement me")
}


