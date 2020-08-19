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

package provider

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/smartwalle/alipay/v3"
	"github.com/xidongc/mongo_ebenchmark/model/payment/paymentpb"
	"strconv"
	"time"
)

type AliPay struct {
	client alipay.Client
}

func (*AliPay) ProviderId() (paymentId paymentpb.PaymentProviderId) {
	return
}

func (a *AliPay) Charge(req *paymentpb.ChargeRequest) (charge *paymentpb.Charge, err error) {
	pay := alipay.TradeWapPay{}
	if req == nil {
		return
	}
	pay.TotalAmount = strconv.Itoa(int(req.GetAmount()))
	pay.Subject = req.GetStatement()
	pay.OutTradeNo = req.GetUserId()
	pay.ProductCode = req.GetUserId() // TODO add order id

	charge = &paymentpb.Charge{
		Id:           fmt.Sprintf("%s", uuid.New()),
		Currency:     req.Currency,
		Paid:         true,
		ChargeAmount: req.Amount,
		UserId:       req.GetUserId(),
	}
	charge.Created = time.Now().UnixNano()
	charge.Updated = charge.Created
	return
}

func (*AliPay) Refund(chargeId string, amount uint64, currency paymentpb.Currency, reason paymentpb.RefundReason) (refund *paymentpb.Refund, err error) {
	return
}

func (*AliPay) SupportedCards() (cardType []paymentpb.CardType) {
	return
}
