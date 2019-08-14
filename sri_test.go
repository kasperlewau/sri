package sri

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
)

func TestHash(t *testing.T) {
	files, err := ioutil.ReadDir("testdata")
	if err != nil {
		t.Fatal(err)
	}

	var src, dest bytes.Buffer

	for _, fi := range files {
		if _, err := src.WriteString("testdata/" + fi.Name() + "\n"); err != nil {
			t.Fatal(err)
		}
	}

	if err := Hash(&src, &dest); err != nil {
		t.Fatal(err)
	}
}

func TestHashMissingFile(t *testing.T) {
	b := []byte(`testdata/Missing.File`)
	err := Hash(bytes.NewReader(b), ioutil.Discard)
	if err, ok := err.(*os.PathError); !ok {
		t.Fatalf("want '*os.pathError'. got = %T", err)
	}
}
