# sri

sri computes the SHA256/384/512 sum for a list of files, writing each line as NDJSON
to a given io.Writer

## usage
### from the cli
```sh
$ find . -maxdepth 1 -type f | sri > shasums.json
```

### as a library
```go
import "github.com/kasperlewau/sri"

func main() {
	b := []byte(`myfile.txt`)
	if err := sri.Hash(bytes.NewReader(b), os.Stdout); err != nil {
		panic(err)
	}
}
```
