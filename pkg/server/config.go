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

package server

import "time"

type Config struct {
	ServerAddr 		string			`short:"s" long:"storage" default:"127.0.0.1" description:"storage server address"`
	ServerPort		int				`long:"storage-port" default:"50053" description:"storage server port"`
	ProxyPort		int				`long:"port" default:"50051" description:"storage server port"`
	Turbo			bool 			`long:"turbo" description:"enable turbo mode for performance"`
	Insecure		bool			`long:"insecure" description:"storage server connection secure"`
	RpcTimeout  	int64			`long:"rpc-timeout" default:"25000" description:"request timeout"`
	BatchSize   	int64			`short:"b" long:"batch" default:"10000" description:"batch size"`
	ReadPref    	int32			`short:"r" long:"readpref" default:"2" description:"read preference"`
	AllowPartial 	bool			`long:"partial" description:"allow partial"`
	Connections     uint			`long:"connections" default:"1" description:"storage server address"`
	Concurrency		uint			`long:"concurrency" default:"5" description:"storage server address"`
	TotalRequest    uint			`long:"requests" default:"20" description:"storage server address"`
	QPS				uint			`long:"qps" description:"storage server address"`
	Timeout			time.Duration	`long:"timeout" description:"storage server address"`
	CPUs            uint			`long:"cpu" default:"1" description:"storage server address"`
}
