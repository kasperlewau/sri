# sri [![GoDoc](https://godoc.org/github.com/golang/gddo?status.svg)](https://godoc.org/github.com/kasperlewau/sri) [![Build Status](https://travis-ci.org/kasperlewau/sri.svg?branch=master)](https://travis-ci.org/kasperlewau/sri) [![Go Report Card](https://goreportcard.com/badge/github.com/kasperlewau/sri)](https://goreportcard.com/report/github.com/kasperlewau/sri)
sri computes the SHA256/384/512 sum for a list of files, writing each line as NDJSON
to a given io.Writer

## install
```sh
go get github.com/kasperlewau/base62
```

## usage
#### cli
```sh
$ find . -maxdepth 1 -type f | sri > shasums.json
```

#### library
```go
import "github.com/kasperlewau/sri"

func main() {
	b := []byte(`myfile.txt`)
	if err := sri.Hash(bytes.NewReader(b), os.Stdout); err != nil {
		panic(err)
	}
}
```

## License
MIT
