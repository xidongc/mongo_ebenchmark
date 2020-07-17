package proxy

import "time"

type Amplifier func() *AmplifyOptions

type Config struct {
	ServerIp 	string	`short:"s" long:"server" default:"10.88.30.82" description:"server ip address"`
	Port		int		`short:"p" long:"port" default:"31024" description:"rpc server port"`
	Insecure	bool	`short:"p" long:"port" default:"31024" description:"rpc server port"`
	RpcTimeout  int64
	BatchSize   int64
	ReadPref    int32
}

type AmplifyOptions struct {
	Connections     uint
	Concurrency		uint
	TotalRequest    uint
	QPS				uint
	Timeout			time.Duration
	CPUs            uint
}

func DefaultConfig() (config *Config) {
	config = &Config{
		ServerIp: "127.0.0.1",
		Port: 50051,
		Insecure: true,
		RpcTimeout: 25000,
		BatchSize: 10000,
		ReadPref: 5,
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


