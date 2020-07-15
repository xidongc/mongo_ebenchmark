package proxy

type Config struct {
	ServerIp 	string	`short:"s" long:"server" default:"10.88.30.82" description:"server ip address"`
	Port		int		`short:"p" long:"port" default:"31024" description:"rpc server port"`
}
