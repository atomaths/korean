korean [![Build Status](https://travis-ci.org/atomaths/korean.png?branch=master)](https://travis-ci.org/atomaths/korean)
======
Package korean provides Korean encodings such as EUC-KR and CP949. It is a wrapper of code.google.com/p/go.text/encoding/korean package for easy to use.

Installation
============

```bash
$ go get github.com/atomaths/korean
```

Usage
=====

```Go
import "github.com/atomaths/korean"
```

```Go
src := "\xb9\xe6\xb0\xa1\xb9\xe6\xb0\xa1\x20\xb0\xed\xc6\xdb"
dst, _ := korean.UTF8([]byte(s))
fmt.Println(string(dst)) // Output: 방가방가 고퍼
```

```Go
src := "방가방가 고퍼"
dst, _ = korean.EUCKR([]byte(src))
for _, v := range dst {
        fmt.Printf("\\x%x", v)
        // Output: \xb9\xe6\xb0\xa1\xb9\xe6\xb0\xa1\x20\xb0\xed\xc6\xdb
}
```
