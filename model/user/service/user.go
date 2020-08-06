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

package service

import (
	"context"
	"errors"
	"github.com/mitchellh/mapstructure"
	log "github.com/sirupsen/logrus"
	"github.com/xidongc-wish/mgo/bson"
	"github.com/xidongc/mongo_ebenchmark/model/user/userpb"
	"github.com/xidongc/mongo_ebenchmark/pkg/proxy"
)

const ns = "user"

type Service struct {
	Storage   proxy.Client
	Amplifier proxy.Amplifier
}

// Create User
func (s Service) New(ctx context.Context, req *userpb.NewRequest) (user *userpb.User, err error) {
	 reqUser := userpb.User{
		Name:     req.GetName(),
		Active:   req.GetActive(),
		Nickname: req.GetNickname(),
		Email:	  req.GetEmail(),
		Balance:  req.GetBalance(),
		Currency: req.GetCurrency(),
		Image: 	  req.GetImage(),
		Pwd:	  req.GetPwd(),
		Metadata: req.GetMetadata(),
	}

	var desired bson.M
    if err = mapstructure.Decode(reqUser, &desired); err != nil {
	    log.Error(err)
    }

	param := proxy.FindModifyParam{
		Filter:   bson.M{"Nickname": req.Nickname},
		Desired:  desired,
		Mode:     proxy.FindAndUpsert,
		SortRule: nil,
		Fields:   nil,
		Amp:      s.Amplifier,
	}

	result, err := s.Storage.FindAndModify(ctx, &param)
	if err != nil {
		log.Error(err)
	}
	err = mapstructure.Decode(result, &user)
	return
}

// Get User
func (s Service) Get(ctx context.Context, req *userpb.GetRequest) (user *userpb.User, err error) {
	param := &proxy.QueryParam{
		Filter:  bson.M{"Nickname": req.GetNickname()},
		FindOne: true,
		Amp:     s.Amplifier,
	}

	results, err := s.Storage.Find(ctx, param)

	if err != nil || len(results) > 1 {
		log.Error(err)
		return
	} else if len(results) == 0 {
		return user, errors.New("no result found")
	}

	err = mapstructure.Decode(results[0], &user)
	if err != nil {
		log.Fatal(err)
	}
	return user, nil
}

// Deactivate User
func (s Service) Deactivate(ctx context.Context, req *userpb.DeleteRequest) (user *userpb.User, err error) {
	user, err = s.Get(ctx, &userpb.GetRequest{Nickname: req.Nickname})
	if err != nil {
		log.Error(err)
	}
	user.Active = false

	var desired bson.M

	if err = mapstructure.Decode(user, &desired); err != nil {
		log.Error(err)
	}

	params := &proxy.FindModifyParam{
		Filter:   bson.M{"Nickname": req.GetNickname()},
		Desired:  desired,
		Mode:     proxy.FindAndUpdate,
		SortRule: nil,
		Fields:   nil,
		Amp:      s.Amplifier,
	}

	result, err := s.Storage.FindAndModify(ctx, params)
	if err != nil {
		log.Error(err)
	}
	err = mapstructure.Decode(result, &user)
	return
}

// Create User Service client
func NewClient(config *proxy.Config, cancel context.CancelFunc) (client *proxy.Client) {
	client, _ = proxy.NewClient(config, ns, cancel)
	return
}
