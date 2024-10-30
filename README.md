# DNS Resolver

A simple command-line DNS resolver written in Go.

## Installation

You can install the DNS resolver using Go:

```bash
go install github.com/mridul037/dns-resolver@latest
```

```bash

dns_resolver <domain> <record_type> [dns_server]

Example: dns_resolver google.com A cloudflare
```
dns_resolver: This is the name of your DNS resolver executable. It’s the program that performs the DNS query.

google.com: This is the domain you want to query. In this case, you’re looking for DNS records associated with google.com.

A: This specifies the type of DNS record you want to retrieve. An "A" record maps a domain to its corresponding IPv4 address. So, when you specify A, you’re asking for the IPv4 address of google.com.

cloudflare: This is the DNS server you want to use for the query. In this case, you're specifying Cloudflare's DNS server. If your resolver supports it, it might default to Cloudflare's public DNS server (1.1.1.1) or another server you can configure.
