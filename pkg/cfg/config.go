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

package cfg

import (
	"github.com/xidongc-wish/mgo"
	"github.com/xidongc/mongo_ebenchmark/mprpc"
	"time"
)

// Server Config
type Config struct {
	ProxyConfig
	AmplifyOptions
	ServerPort int  `long:"server-port" default:"50053" description:" api server port"`
	Turbo      bool `long:"turbo" description:"enable turbo mode"`
}

// AmplifyOptions
type Amplifier *AmplifyOptions

// Proxy client cfg
type ProxyConfig struct {
	ProxyAddr    string `long:"proxy-addr" default:"127.0.0.1" description:"storage address"`
	ProxyPort    int    `long:"proxy-port" default:"50051" description:"storage port"`
	Secure       bool   `long:"https" description:"use tls to connect proxy backend"`
	RpcTimeout   int64  `long:"rpc-timeout" default:"25000" description:"storage request timeout"`
	BatchSize    int64  `short:"b" long:"batch" default:"10000" description:"batch size"`
	ReadPref     int32  `short:"r" long:"read-pref" default:"2" description:"read preference"`
	AllowPartial bool   `long:"partial" description:"allow partial"`
}

// AmplifyOptions for amp
type AmplifyOptions struct {
	Connections  uint          `long:"connections" default:"1" description:"request connections for amp"`
	Concurrency  uint          `long:"concurrency" default:"5" description:"request concurrency for amp"`
	TotalRequest uint          `long:"requests" default:"20" description:"total perf requests sent for amp"`
	QPS          uint          `long:"qps" description:"qps used for amp"`
	Timeout      time.Duration `long:"timeout" description:"timeout for amp request to backend"`
	CPUs         uint          `long:"cpu" default:"1" description:"cpus used for amp"`
}

// Create default cfg
func DefaultConfig() (config *ProxyConfig) {
	config = &ProxyConfig{
		ProxyAddr:    "127.0.0.1",
		ProxyPort:    50051,
		Secure:       false,
		RpcTimeout:   25000,
		BatchSize:    10000,
		ReadPref:     int32(mgo.Primary),
		AllowPartial: false,
	}
	return
}

// MicroAmplifier generate AmplifyOptions for light weight workload
func MicroAmplifier() (amplifier *AmplifyOptions) {
	amplifier = &AmplifyOptions{
		Connections:  1,
		Concurrency:  5,
		TotalRequest: 20,
	}
	return
}

// StressAmplifier generate AmplifyOptions for stress workload
func StressAmplifier() (amplifier *AmplifyOptions) {
	amplifier = &AmplifyOptions{
		Connections:  10,
		Concurrency:  100,
		TotalRequest: 200,
	}
	return
}

// Get AmplifyOptions based on given param
func Amplifer(options AmplifyOptions) *AmplifyOptions {
	return &options
}

// getTurboWriteOptions returns a eventual consistency write options
//
// Relevant documentation:
//
// merge fsync and J: https://jira.mongodb.org/browse/SERVER-11399
//
func GetTurboWriteOptions() (wOptions *mprpc.WriteOptions) {
	wOptions = &mprpc.WriteOptions{
		Writeconcern: 1,
		Writetimeout: 0,
		Writemode:    "",
		J:            false,
	}
	return
}

// getSafeWriteOptions returns a strictly consistency write options
//
// Relevant documentation:
//
// merge fsync and J: https://jira.mongodb.org/browse/SERVER-11399
//
func GetSafeWriteOptions() (wOptions *mprpc.WriteOptions) {
	wOptions = &mprpc.WriteOptions{
		Writetimeout: 0,
		Writemode:    "majority",
		J:            true,
	}
	return
}
