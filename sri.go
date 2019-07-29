// package sri provides the means to compute sub-resource-integrity sha sums for a list of files
package sri

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"hash"
	"io"
	"io/ioutil"
	"os"
)

// Resource is a named file with sub-resource-integrity hashes combined
type Resource struct {
	Path   string    `json:"path"`
	Sha256 hash.Hash `json:"sha_256"`
	Sha384 hash.Hash `json:"sha_384"`
	Sha512 hash.Hash `json:"sha_512"`
}

// Reset resets the hashes of a Resource and zero-s out the filename
func (r *Resource) Reset() {
	r.Path = ""
	r.Sha256.Reset()
	r.Sha384.Reset()
	r.Sha512.Reset()
}

func (r *Resource) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Path   string `json:"path"`
		Sha256 string `json:"sha_256"`
		Sha384 string `json:"sha_384"`
		Sha512 string `json:"sha_512"`
	}{
		Path:   r.Path,
		Sha256: hex.EncodeToString(r.Sha256.Sum(nil)),
		Sha384: hex.EncodeToString(r.Sha384.Sum(nil)),
		Sha512: hex.EncodeToString(r.Sha512.Sum(nil)),
	})
}

// Hash computes the SHA256/384/512 sum for a list of files
// and writes each line as NDJSON to the given writer
// If an error occurs, Hash wont run to completion and the
// error will be returned immediately
func Hash(src io.Reader, dest io.Writer) error {
	out := json.NewEncoder(dest)
	s := bufio.NewScanner(src)

	r := &Resource{
		Path:   "",
		Sha256: sha256.New(),
		Sha384: sha512.New384(),
		Sha512: sha512.New(),
	}

	for s.Scan() {
		f, err := os.Open(s.Text())
		if err != nil {
			f.Close()
			return err
		}

		r.Path = s.Text()

		b, err := ioutil.ReadAll(f)
		if err != nil {
			f.Close()
			r.Reset()
			return err
		}

		if _, err := r.Sha256.Write(b); err != nil {
			f.Close()
			r.Reset()
			return err
		}

		if _, err := r.Sha384.Write(b); err != nil {
			f.Close()
			r.Reset()
			return err
		}

		if _, err := r.Sha512.Write(b); err != nil {
			f.Close()
			r.Reset()
			return err
		}

		if err := out.Encode(r); err != nil {
			f.Close()
			r.Reset()
			return err
		}

		f.Close()
		r.Reset()
	}

	return nil
}
