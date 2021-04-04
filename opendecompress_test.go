package opendecompress

import (
	"bytes"
	"io/ioutil"
	"path"
	"testing"
)

func TestDecompress(t *testing.T) {
	expected, err := ioutil.ReadFile("./test_data/test.txt")
	if err != nil {
		t.Fatalf("Failed to read expected data: %v", err)
	}

	files, err := ioutil.ReadDir("./test_data/")
	if err != nil {
		t.Fatalf("Failed to list test data: %v", err)
	}

	for _, f := range files {
		fn := f.Name()
		t.Run(fn, func(t *testing.T) {
			t.Logf("Testing %s", fn)
			df, err := Open(path.Join("test_data", fn))
			if err != nil {
				t.Fatalf("Failed to open %v", fn)
			}
			contents, err := ioutil.ReadAll(df)
			if err != nil {
				t.Fatalf("Failed to read %v", fn)
			}
			if bytes.Compare(contents, expected) != 0 {
				t.Fatalf("contents != expected: %s != %s", contents, expected)
			}
			err = df.Close()
			if err != nil {
				t.Fatalf("Failed to Close: %v", fn)
			}
		})
	}
}
