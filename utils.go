package election
import (
	"net"
	"encoding/binary"
	"strconv"
	"strings"
)

func supportedInterfaces() []*net.Interface {

	var ifis []*net.Interface
	interfaces, err := net.Interfaces()

	if err != nil {
		return nil
	}

	for _, ifi := range interfaces {

		if ((ifi.Flags & net.FlagUp) != 0) && ((ifi.Flags & net.FlagMulticast) != 0){
			ifis = append(ifis, &net.Interface{Index:ifi.Index,MTU:ifi.MTU,Name:ifi.Name,HardwareAddr:ifi.HardwareAddr,Flags:ifi.Flags})
		}
	}

	return ifis;
}

func getLocalInterfaceIpAddress(ifi *net.Interface) (string, uint32) {

	addrs, err := ifi.Addrs()
	if err != nil {
		return "<nil>", 0
	}

	for _, add := range addrs {
		if ipnet, ok := add.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), binary.BigEndian.Uint32(ipnet.IP.To4())
			}
		}
	}

	return "<nil>",0
}

func parseMessage(data []string) *mMessage {

	num,_ := strconv.ParseUint(data[1],10,32)
	num2,_ := strconv.ParseInt(data[2],10,32)
	i := strings.Index(data[3],"###")
	ip := data[3][:i]

	return &mMessage{message:data[0], ipNumber:uint32(num), processId:int(num2), ipAddr:ip}
}
