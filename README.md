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


```bash
 go build -o dns_resolver dns_resolver.go
 ./dns_resolver example.com A

------or---------------
go run dns_resolver.go example.com A cloudflare

```


## Features

- Resolve various DNS record types: [ A, AAAA, CNAME, MX, TXT, SRV ].
- Configurable DNS servers [Google (default), cloudflare,dnswatch,quad9].
- Simple command-line interface.


## DNS TYPE
- A: Returns the IPv4 address associated with the domain.
- AAAA: Returns the IPv6 address associated with the domain.
- CNAME: Returns the canonical name (alias) of the domain.
- MX: Returns mail exchange server records.
- TXT: Returns text records associated with the domain.
- SRV: Returns service records
