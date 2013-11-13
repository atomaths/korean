// Copyright 2013 Jongmin Kim. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package korean

import (
	"bytes"
	"fmt"
	"testing"

	"code.google.com/p/go.text/encoding/korean"
)

var basicTestCases = []struct {
	euckr string
	utf8  string
}{
	{
		euckr: "\xb9\xe6\xb0\xa1\xb9\xe6\xb0\xa1\x20\xb0\xed\xc6\xdb",
		utf8:  "방가방가 고퍼",
	},
}

type transFunc func([]byte) ([]byte, error)

func TestBasics(t *testing.T) {
	for _, tc := range basicTestCases {
		for _, direction := range []string{"UTF8", "EUCKR"} {
			newTransformer, want, src := (transFunc)(nil), "", ""
			if direction == "UTF8" {
				newTransformer, want, src = UTF8, tc.utf8, tc.euckr
			} else {
				newTransformer, want, src = EUCKR, tc.euckr, tc.utf8
			}

			result, err := newTransformer([]byte(src))
			if err != nil {
				t.Errorf("%s didn't match: %v", direction, err)
				continue
			}
			if !bytes.Equal(result, []byte(want)) {
				t.Errorf("%s didn't match: %v", direction, err)
				continue
			}
		}
	}
}

func TestTrans(t *testing.T) {
	trans(korean.EUCKR.NewDecoder(), []byte("고퍼"))
}

// TODO(atomaths): fill correct benchmark test case.
func BenchmarkTrans(b *testing.B) {
	for i := 0; i < b.N; i++ {
	}
}

func ExampleTest() {
	dst, _ := UTF8([]byte("\xb9\xe6\xb0\xa1\xb9\xe6\xb0\xa1\x20\xb0\xed\xc6\xdb"))
	fmt.Printf("%s", dst)
	// Output:
	// 방가방가 고퍼
}
