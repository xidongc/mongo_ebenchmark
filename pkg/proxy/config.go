package proxy

import (
	"github.com/xidongc-wish/mgo"
	"github.com/xidongc-wish/mongoproxy/mprpc"
	"time"
)

type Amplifier func() *AmplifyOptions

type Config struct {
	ServerIp 		string			`short:"s" long:"storage" default:"127.0.0.1" description:"storage server address"`
	Port			int				`short:"p" long:"port" default:"50051" description:"storage server port"`
	Insecure		bool			`long:"insecure" description:"storage server connection secure"`
	RpcTimeout  	int64			`short:"t" long:"timeout" default:"25000" description:"request timeout"`
	BatchSize   	int64			`short:"b" long:"batch" default:"10000" description:"batch size"`
	ReadPref    	int32			`short:"r" long:"readpref" default:"2" description:"read preference"`
	AllowPartial 	bool			`long:"partial" description:"allow partial"`
}

type AmplifyOptions struct {
	Connections     uint			`long:"connections" default:"1" description:"storage server address"`
	Concurrency		uint			`long:"concurrency" default:"5" description:"storage server address"`
	TotalRequest    uint			`long:"requests" default:"20" description:"storage server address"`
	QPS				uint			`long:"qps" description:"storage server address"`
	Timeout			time.Duration	`long:"timeout" description:"storage server address"`
	CPUs            uint			`long:"cpu" default:"1" description:"storage server address"`
}

// Create default config
func DefaultConfig() (config *Config) {
	config = &Config{
		ServerIp: "127.0.0.1",
		Port: 50051,
		Insecure: true,
		RpcTimeout: 25000,
		BatchSize: 10000,
		ReadPref: int32(mgo.Primary),
		AllowPartial: false,
	}
	return
}

func MicroAmplifier() (amplifier *AmplifyOptions) {
	amplifier = &AmplifyOptions{
		Connections: 1,
		Concurrency: 5,
		TotalRequest: 20,
	}
	return
}

func StressAmplifier() (amplifier *AmplifyOptions) {
	amplifier = & AmplifyOptions{
		Connections: 10,
		Concurrency: 100,
		TotalRequest: 200,
	}
	return
}

// merge fsync and J: https://jira.mongodb.org/browse/SERVER-11399
func getTurboWriteOptions() (wOptions *mprpc.WriteOptions) {
	wOptions = &mprpc.WriteOptions{
		Writeconcern: 1,
		Writetimeout: 0,
		Writemode: "",
		J: false,
	}
	return
}

func getSafeWriteOptions() (wOptions *mprpc.WriteOptions) {
	wOptions = &mprpc.WriteOptions{
		Writetimeout: 0,
		Writemode: "majority",
		J: true,
	}
	return
}

