# DoH Query Tool

A simple command-line tool for querying DNS records using a specified DoH server.

## Features

- Specify DNS server URL
- Query multiple DNS types (A, AAAA, CNAME, etc.)
- Verbose output for detailed information
- Help message for usage instructions

## Usage
```bash
# source code with golang
go run main.go [options] <hostname>
# or compilated
dohlookup [options] <hostname>
```

## Compilation

To compile the program for different platforms, use the following commands:
```shell
# Compile for Windows
GOOS=windows GOARCH=amd64 go build -o dohlookup.exe main.go
# Compile for macOS
GOOS=darwin GOARCH=amd64 go build -o dohlookup main.go
# Compile for Linux
GOOS=linux GOARCH=amd64 go build -o dohlookup main.go
```

## Options

- `-s, --server <url>`: Specify the DNS server URL (default: `https://1.1.1.1/dns-query`)
- `-t, --type <type>`: Specify the DNS query types (comma-separated, e.g., A,AAAA). Use multiple `-t` options to query multiple types.
- `-v, --verbose`: Display verbose output.
- `-h, --help`: Display this help message.

## Example

To query the A and AAAA records for `example.com` using the default DNS server:
```shell
dohlookup -t A -t AAAA example.com
```

To use a custom DNS server:
```shell
dohlookup --server https://8.8.8.8/dns-query -t A example.com
```

## License

This project is licensed under the MIT License.