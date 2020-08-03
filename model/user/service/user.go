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
	"errors"
	"github.com/mitchellh/mapstructure"
	log "github.com/sirupsen/logrus"
	"github.com/xidongc-wish/mgo/bson"
	"github.com/xidongc/mongodb_ebenchmark/model/user/userpb"
	"github.com/xidongc/mongodb_ebenchmark/pkg/proxy"
)

const ns = "user"

type Service struct {
	Storage   proxy.Client
	Amplifier proxy.Amplifier
}

// Create User
func (s Service) New(ctx context.Context, req *userpb.NewRequest) (*userpb.User, error) {
	user := userpb.User{
		Name:        req.GetName(),
		Metadata:    req.GetMetadata(),
		Active:      req.GetActive(),
	}

	var docs []interface{}
	docs = append(docs, user)

	param := &proxy.InsertParam{
		Docs: docs,
		Amp:  s.Amplifier,
	}

	if err := s.Storage.Insert(ctx, param); err != nil {
		log.Error(err)
		return nil, err
	}
	return &user, nil
}

func (s Service) Get(ctx context.Context, req *userpb.GetRequest) (user *userpb.User, err error) {
	param := &proxy.QueryParam{
		Filter:      bson.M{"_id": req.Id},
		FindOne:     true,
		Amp:         s.Amplifier,
	}

	results, err := s.Storage.Find(ctx, param)

	if err != nil || len(results) > 1 {
		log.Error(err)
		return
	} else if len(results) == 0 {
		return user, errors.New("no result found")
	}

	err = mapstructure.Decode(results[0], user)
	if err != nil {
		log.Fatal(err)
	}
	return user, nil
}

func (s Service) Update(ctx context.Context, req *userpb.UpdateRequest) (user *userpb.User, err error) {
	var updateParams bson.M
	if err = mapstructure.Decode(req, &updateParams); err != nil {
		log.Error(err)
		return
	}
	updateQuery := &proxy.UpdateParam{
		Filter: bson.M{"_id": req.Id},
		Update: updateParams,
		Upsert: false,
		Multi:  true,
		Amp:    s.Amplifier,
	}
	_, err = s.Storage.Update(ctx, updateQuery)
	if err != nil {
		return
	}
	return
}

func (s Service) Delete(ctx context.Context, req *userpb.DeleteRequest) (*userpb.Empty, error) {
	removeQuery := &proxy.RemoveParam{
		Filter: bson.M{"_id": req.Id},
		Amp:    s.Amplifier,
	}
	_, err := s.Storage.Remove(ctx, removeQuery)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &userpb.Empty{}, nil
}
