package utils

import (
	"net"
)

func GetFeePort()(int,error)  {
	
	addr,eer:=net.ResolveTCPAddr("tcp","localhost:0")

	if eer != nil {
		return 0,nil
	}

	l,err:=net.ListenTCP("tcp",addr)
	if err != nil {
		return 0, err
	}
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port,nil
}

