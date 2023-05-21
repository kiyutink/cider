# Cider
Cider is a small command line tool that given a valid CIDR range outputs the number of IP addresses in that range.

## Installation
> Make sure your `$GOPATH/bin` is added in your `$PATH`
```
go install github.com/kiyutink/cider@latest
```

## Usage
```
cider <CIDR>
```
### Example
```
~ cider 10.0.0.0/30
Total IP addresses: 4
10.0.0.0
10.0.0.3
```
### Flags
| -l | list all the ip addresses instead of only start and end of the block |
|----|----------------------------------------------------------------------|
