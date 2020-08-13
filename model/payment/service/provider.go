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

import "github.com/xidongc/mongo_ebenchmark/model/payment/paymentpb"

type Provider interface {
	ProviderId() paymentpb.PaymentProviderId
	Charge(req *paymentpb.ChargeRequest) (*paymentpb.Charge, error)
	Refund(chargeId string, amount uint64, currency paymentpb.Currency, reason paymentpb.RefundReason) (*paymentpb.Refund, error)
	SupportedCards() []paymentpb.CardType
}
