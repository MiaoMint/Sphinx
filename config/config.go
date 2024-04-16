package config

import (
	"flag"
	"log"
	"os"
	"runtime"

	"github.com/MiaoMint/Sphinx/util"
)

var (
	HostsPath string
	LocalIP   string
	Domain    string
	hostsMap  map[string]string
	Listen80  bool
	DevMode   bool
)

func init() {
	var defaultHostsPath string

	if runtime.GOOS == "windows" {
		defaultHostsPath = "C:\\Windows\\System32\\drivers\\etc\\hosts"
	} else {
		defaultHostsPath = "/etc/hosts"
	}

	flag.StringVar(&HostsPath, "hosts_path", defaultHostsPath, "Hosts file path")
	flag.StringVar(&LocalIP, "local_ip", "127.0.0.1", "Local IP address")
	flag.StringVar(&Domain, "domain", "sphinx.lan", "Domain name")
	flag.BoolVar(&Listen80, "listen80", false, "Listen on port 80")
	flag.BoolVar(&DevMode, "dev", false, "Development mode")

	flag.Parse()

	LoadHosts()
	AddHosts(Domain)
}

func LoadHosts() {
	hosts, err := os.ReadFile(HostsPath)
	if err != nil {
		log.Fatalln("Failed to read hosts file:", err)
	}
	hostsMap = util.ParseHosts(string(hosts))
}

func SaveHosts() {
	hosts := ""
	for domain, ip := range hostsMap {
		hosts += ip + " " + domain + "\n"
	}

	err := os.WriteFile(HostsPath, []byte(hosts), 0644)
	if err != nil {
		log.Fatalln("Failed to write hosts file:", err)
	}
}

func GetHosts(domain ...string) any {
	if len(domain) == 0 {
		return hostsMap
	}
	return hostsMap[domain[0]]
}

func AddHosts(domain string) {
	hostsMap[domain] = LocalIP
	SaveHosts()
}

func RemoveHosts(domain string) {
	delete(hostsMap, domain)
	SaveHosts()
}
