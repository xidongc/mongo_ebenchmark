package proxy

import (
	"github.com/xidongc-wish/mgo"
	"github.com/xidongc-wish/mongoproxy/mprpc"
	"time"
)

type Amplifier func() *AmplifyOptions

type Config struct {
	ServerIp 		string	`short:"s" long:"server" default:"10.88.30.82" description:"server ip address"`
	Port			int		`short:"p" long:"port" default:"31024" description:"rpc server port"`
	Insecure		bool	`short:"p" long:"port" default:"31024" description:"rpc server port"`
	RpcTimeout  	int64
	BatchSize   	int64
	ReadPref    	int32
	AllowPartial 	bool
}

type AmplifyOptions struct {
	Connections     uint
	Concurrency		uint
	TotalRequest    uint
	QPS				uint
	Timeout			time.Duration
	CPUs            uint
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

