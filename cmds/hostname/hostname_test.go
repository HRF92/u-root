// Copyright 2015-2017 the u-root Authors. All rights reserved
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"os"
	"testing"
)

func Test_hostname(t *testing.T) {
	var buf bytes.Buffer
	var host string
	var err error

	if err = hostname(&buf); err != nil {
		t.Errorf("%v", err)
	}

	if host, err = os.Hostname(); err != nil {
		t.Errorf("%v", err)
	}

	if bytes.Compare(buf.Bytes(), []byte(host)) != 0 {
		t.Fatalf("want %v, got %v", []byte(host), buf)
	}
}
