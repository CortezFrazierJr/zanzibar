// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package main

import (
	"crypto/md5"
	"encoding/base64"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/mailru/easyjson/parser"
	// Reference the gen package to be friendly to vendoring tools,
	// as it is an indirect dependency.
	// (The temporary bootstrapping code uses it.)
	"github.com/mailru/easyjson/bootstrap"
	_ "github.com/mailru/easyjson/gen"
)

var allStructs = flag.Bool(
	"all",
	false,
	"generate marshaler/unmarshalers for all structs in a file",
)

var checksumPrefix = "// Checksum : "
var prefixBytes = []byte(
	"// Code generated by zanzibar\n" +
		"// @generated\n",
)

func generate(fname string) error {
	fInfo, err := os.Stat(fname)
	if err != nil {
		return err
	}

	p := parser.Parser{AllStructs: *allStructs}
	if err := p.Parse(fname, fInfo.IsDir()); err != nil {
		return fmt.Errorf("Error parsing %v: %v", fname, err)
	}

	var outName string
	if fInfo.IsDir() {
		outName = filepath.Join(fname, p.PkgName+"_easyjson.go")
	} else {
		s := strings.TrimSuffix(fname, ".go")
		if s == fname {
			return errors.New("Filename must end in '.go'")
		}
		outName = s + "_easyjson.go"
	}

	g := bootstrap.Generator{
		BuildTags:       "",
		PkgPath:         p.PkgPath,
		PkgName:         p.PkgName,
		Types:           p.StructNames,
		SnakeCase:       false,
		NoStdMarshalers: false,
		OmitEmpty:       false,
		LeaveTemps:      false,
		OutName:         outName,
		StubsOnly:       false,
		NoFormat:        false,
	}

	if err := g.Run(); err != nil {
		return fmt.Errorf("Bootstrap failed: %v", err)
	}

	return nil
}

func getOldChecksum(easyJSONFile string) string {
	oldEasyJSONBytes, err := ioutil.ReadFile(easyJSONFile)
	if err == nil {
		sliceStart := len(prefixBytes) + len(checksumPrefix)
		sliceEnd := len(prefixBytes) + len(checksumPrefix) + 24
		return string(oldEasyJSONBytes[sliceStart:sliceEnd])
	}

	return ""
}

func getNewChecksum(file string) string {
	fileBytes, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
		return ""
	}

	checksum := md5.Sum(fileBytes)
	return base64.StdEncoding.EncodeToString(checksum[:])
}

func main() {
	flag.Parse()

	files := flag.Args()
	if len(files) == 0 {
		flag.Usage()
		os.Exit(1)
		return
	}

	// Do not paralleliz, not worth it.
	for _, file := range files {
		easyJSONFile := file[0:len(file)-3] + "_easyjson.go"
		oldChecksum := getOldChecksum(easyJSONFile)
		newChecksum := getNewChecksum(file)

		// If we have an checksum in easyjson file check it.
		if oldChecksum != "" && oldChecksum == newChecksum {
			continue
		}

		if err := generate(file); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
			return
		}

		bytes, err := ioutil.ReadFile(easyJSONFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
			return
		}

		checksumLine := checksumPrefix + newChecksum + "\n"
		newLength := len(bytes) + len(prefixBytes) + len(checksumLine)

		newBytes := make([]byte, newLength)
		copy(newBytes, prefixBytes)
		copy(newBytes[len(prefixBytes):], []byte(checksumLine))
		copy(newBytes[len(prefixBytes)+len(checksumLine):], bytes)

		err = ioutil.WriteFile(easyJSONFile, newBytes, 0644)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
			return
		}
	}
}
