package goutil

import (
	"gopkg.in/olahol/melody.v1"
	"net"
	"strings"
)

// ClientWebsocketRemoteAddress ws获取远程地址(不标准:proxyIp+:+localPort)
func ClientWebsocketRemoteAddress(session *melody.Session) string {
	ip := ClientWebsocketIP(session)
	strs := strings.Split(session.Request.RemoteAddr, ":")
	if len(strs) < 1 {
		return ip
	}
	return ip + ":" + strs[len(strs)-1]
}

// ClientWebsocketIP ws设备ip
func ClientWebsocketIP(session *melody.Session) string {
	//获取反响代理ip
	for headerName := range session.Request.Header {
		ip, valid := validateHeader(session.Request.Header.Get(headerName))
		if valid {
			return ip
		}
	}

	//获取主机ip
	ip, _, err := net.SplitHostPort(strings.TrimSpace(session.Request.RemoteAddr))
	if err != nil {
		return ""
	}
	remoteIP := net.ParseIP(ip)
	if remoteIP == nil {
		return ""
	}
	return remoteIP.String()
}

// validateHeader will parse X-Forwarded-For header and return the trusted client IP address
func validateHeader(header string) (clientIP string, valid bool) {
	if header == "" {
		return "", false
	}
	items := strings.Split(header, ",")
	for i := len(items) - 1; i >= 0; i-- {
		ipStr := strings.TrimSpace(items[i])
		ip := net.ParseIP(ipStr)
		if ip == nil {
			break
		}

		// X-Forwarded-For is appended by proxy
		// Check IPs in reverse order and stop when find untrusted proxy
		if i == 0 {
			return ipStr, true
		}
	}
	return "", false
}
