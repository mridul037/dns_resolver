# DNS Resolver

A simple command-line DNS resolver written in Go.

## Installation

You can install the DNS resolver using Go:

```bash
go install github.com/mridul037/dns-resolver@latest

dns_resolver <domain> <record_type> [dns_server]

Example: dns_resolver google.com A cloudflare
