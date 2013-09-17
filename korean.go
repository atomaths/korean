// Copyright 2013 Jongmin Kim. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package korean wraps EUC-KR encoder, decoder in
// code.google.com/p/go.text/encoding/korean package.
package korean

import (
	"reflect"

	"code.google.com/p/go.text/encoding/korean"
	"code.google.com/p/go.text/transform"
)

const (
	defaultBufSize = 4096
)

// Transform writes to dst the transformed bytes read from src, and
// returns dst bytes slice written and src bytes read.
func Transform(t transform.Transformer, src []byte) (dst []byte, err error) {
	switch reflect.TypeOf(t).Name() {
	case "eucKRDecoder":
		dst = make([]byte, len(src)+len(src)/2)
	case "eucKREncoder":
		dst = make([]byte, len(src))
	}

	for {
		nDst, _, err := t.Transform(dst, src, true)
		if err != nil {
			// Destination buffer was too short to receive
			// all of the transformed bytes.
			if err == transform.ErrShortDst {
				dst = make([]byte, len(dst)+defaultBufSize)
				continue
			} else {
				return nil, err
			}
		}
		return dst[:nDst], nil
	}
}

// UTF8 converts from EUC-KR bytes to UTF-8 bytes.
func UTF8(src []byte) (dst []byte, err error) {
	return Transform(korean.EUCKR.NewDecoder(), src)
}

// EUCKR converts from UTF-8 bytes to EUC-KR bytes.
func EUCKR(src []byte) (dst []byte, err error) {
	return Transform(korean.EUCKR.NewEncoder(), src)
}
