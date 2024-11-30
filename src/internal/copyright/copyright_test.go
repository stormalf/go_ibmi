// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package copyright

import (
	"bytes"
	"internal/testenv"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"testing"
)

var copyright = []byte("Copyright")

var permitted = [][]byte{
	[]byte("// Code generated by "),
	[]byte("// Code generated from "),
	[]byte("// Created by cgo -cdefs"),
	[]byte("// DO NOT EDIT\n// generated by:"),
	[]byte("// Empty assembly file"),
	[]byte("// Generated using cgo"),
	[]byte("// Original source:\n//\thttp://www.zorinaq.com/papers/md5-amd64.html"), // public domain crypto/md5
	[]byte("// created by cgo -cdefs"),
	[]byte("// go run mkasm.go"),
	[]byte("// mkerrors"),
	[]byte("// mksys"),
	[]byte("// run\n// Code generated by"), // cmd/compile/internal/test/constFold_test.go
}

func TestCopyright(t *testing.T) {
	buf := make([]byte, 2048)
	filepath.WalkDir(filepath.Join(testenv.GOROOT(t), "src"), func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() && (d.Name() == "testdata" || d.Name() == "vendor") {
			return filepath.SkipDir
		}
		switch filepath.Ext(d.Name()) {
		default:
			return nil
		case ".s", ".go":
			// check
		}

		f, err := os.Open(path)
		if err != nil {
			t.Error(err)
			return nil
		}
		defer f.Close()
		n, err := f.Read(buf)
		if err != nil && err != io.EOF {
			t.Error(err)
			return nil
		}
		b := buf[:n]
		if bytes.Contains(b, copyright) {
			return nil
		}
		for _, ok := range permitted {
			if bytes.HasPrefix(b, ok) {
				return nil
			}
		}
		t.Errorf("%s: missing copyright notice", path)
		return nil
	})
}
