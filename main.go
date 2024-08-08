package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

// 定义一个类型来实现flag.Value接口
type stringSlice []string

func (s *stringSlice) String() string {
	return fmt.Sprint(*s)
}

func (s *stringSlice) Set(value string) error {
	*s = append(*s, value)
	return nil
}
func main() {
	var hostName string
	var dnsServer string
	var dnsTypes []string
	var showHelp bool
	var showVerbose bool

	flag.StringVar(&dnsServer, "s", "https://1.1.1.1/dns-query", "Specify the DNS server URL")
	flag.StringVar(&dnsServer, "server", "https://1.1.1.1/dns-query", "Specify the DNS server URL")
	flag.Var((*stringSlice)(&dnsTypes), "t", "Specify the DNS query types (comma-separated, e.g., A,AAAA)")
	flag.Var((*stringSlice)(&dnsTypes), "type", "Specify the DNS query types (comma-separated, e.g., A,AAAA)")
	flag.BoolVar(&showVerbose, "v", false, "Display verbose output")
	flag.BoolVar(&showVerbose, "verbose", false, "Display verbose output")
	flag.BoolVar(&showHelp, "h", false, "Display help message")
	flag.BoolVar(&showHelp, "help", false, "Display help message")

	flag.Parse()
	flag.Usage = func() {
		fmt.Println("Usage:")
		fmt.Println("  -s, --server   Specify the DNS server URL (default: https://1.1.1.1/dns-query)")
		fmt.Println("  -t, --type     Specify the DNS query type (A/AAAA/CNAME...) (default: A,AAAA)")
		fmt.Println("                 More -t options to query multiple types. eg -t A -t CNAME ...")
		fmt.Println("  -v, --verbose  Display verbose output")
		fmt.Println("  -h, --help     Display this help message")
		fmt.Println("  <hostname>     The hostname to query")
	}

	if showHelp {
		flag.Usage()
		os.Exit(0)
	}

	if len(flag.Args()) != 1 {
		flag.Usage()
		os.Exit(1)
	}
	if len(dnsTypes) == 0 {
		dnsTypes = []string{"A", "AAAA"}
	}

	hostName = flag.Args()[0]
	fmt.Println("Server: ", dnsServer)
	if showVerbose {
		fmt.Println("Query Type: ", strings.Join(dnsTypes, ", "))
	}
	fmt.Println()
	for k, dnsType := range dnsTypes {
		res, err := DohQuery(hostName, dnsServer, dnsType)
		if err != nil {
			log.Println("error: ", err)
			os.Exit(1)
		}
		if k == 0 {
			fmt.Println("Answer: ")
		}
		if showVerbose {
			fmt.Println(res.String())
		} else {
			for _, answer := range res.Answer {
				fmt.Println(answer.String())
			}
		}
	}
}
