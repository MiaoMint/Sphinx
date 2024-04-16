package util

import "strings"

// 解析hosts文件
func ParseHosts(hosts string) map[string]string {
	hostsMap := make(map[string]string)

	hostsLines := strings.Split(hosts, "\n")

	for _, line := range hostsLines {
		line = strings.TrimSpace(line)

		if strings.HasPrefix(line, "#") {
			continue
		}

		fields := strings.Fields(line)

		if len(fields) < 2 {
			continue
		}

		ip := fields[0]
		for _, host := range fields[1:] {
			hostsMap[host] = ip
		}
	}

	return hostsMap
}
