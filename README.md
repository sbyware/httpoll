# httpoll

easy http endpoint polling, built with Go

## Installation

```bash
go install github.com/sbyware/httpoll
```
### Manual Installation

```bash
git clone https://github.com/sbyware/httpoll.git
cd httpoll
go build -o build/httpoll
cp build/httpoll /usr/local/bin
```

## Usage

```bash
httpoll -u http://localhost:8080 -i 1000 -c 5
```

## Options

```bash
httpoll -h
```

## License

MIT