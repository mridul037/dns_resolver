package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/miekg/dns"
)

var dnsServers = map[string]string{
	"google":    "8.8.8.8:53",
	"cloudflare": "1.1.1.1:53",
	"dnswatch":  "84.200.69.80:53",
	"quad9":     "9.9.9.9:53",
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: dnsclient <domain> <record_type> [dns_server...]")
		fmt.Println("Available DNS servers: google, cloudflare, dnswatch, quad9")
		fmt.Println("Default DNS server is google (8.8.8.8:53)")
		return
	}

	domain := os.Args[1]
	recordType := strings.ToUpper(os.Args[2])

	var dnsServer string
	if len(os.Args) > 3 {
		// Use the first provided DNS server from the arguments
		serverName := strings.ToLower(os.Args[3])
		var found bool
		dnsServer, found = dnsServers[serverName]
		if !found {
			fmt.Printf("Unknown DNS server: %s. Using default (Google DNS)\n", serverName)
			dnsServer = dnsServers["google"]
		}
	} else {
		dnsServer = dnsServers["google"] // Default to Google DNS
	}

	// Create a DNS client
	client := new(dns.Client)
	msg := new(dns.Msg)
	msg.SetQuestion(dns.Fqdn(domain), getDNSRecordType(recordType))

	// Perform the DNS query
	response, _, err := client.Exchange(msg, dnsServer)
	if err != nil {
		fmt.Printf("Failed to lookup %s for %s using %s: %v\n", recordType, domain, dnsServer, err)
		return
	}

	// Print the results
	printDNSRecords(response, recordType)
}

func getDNSRecordType(recordType string) uint16 {
	switch recordType {
	case "A":
		return dns.TypeA
	case "AAAA":
		return dns.TypeAAAA
	case "CNAME":
		return dns.TypeCNAME
	case "MX":
		return dns.TypeMX
	case "TXT":
		return dns.TypeTXT
	case "SRV":
		return dns.TypeSRV
	default:
		fmt.Printf("Unsupported record type: %s\n", recordType)
		os.Exit(1)
		return 0
	}
}

func printDNSRecords(response *dns.Msg, recordType string) {
	for _, answer := range response.Answer {
		switch recordType {
		case "A":
			if a, ok := answer.(*dns.A); ok {
				fmt.Println("A Record:", a.A)
			}
		case "AAAA":
			if aaaa, ok := answer.(*dns.AAAA); ok {
				fmt.Println("AAAA Record:", aaaa.AAAA)
			}
		case "CNAME":
			if cname, ok := answer.(*dns.CNAME); ok {
				fmt.Println("CNAME Record:", cname.Target)
			}
		case "MX":
			if mx, ok := answer.(*dns.MX); ok {
				fmt.Printf("MX Record: %d %s\n", mx.Preference, mx.Mx)
			}
		case "TXT":
			if txt, ok := answer.(*dns.TXT); ok {
				fmt.Println("TXT Record:", txt.Txt)
			}
		case "SRV":
			if srv, ok := answer.(*dns.SRV); ok {
				fmt.Printf("SRV Record: %d %d %d %s\n", srv.Priority, srv.Weight, srv.Port, srv.Target)
			}
		}
	}
}
