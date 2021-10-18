// Copyright 2021 by William Jones. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Sinple commin hashing functions
package hashing

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"hash"
	"io"
	"os"
)

const (
	MD5    = 1
	SHA256 = 2
	SHA512 = 3
)

func FileHash(hashtype int, fn string) (string, error) {
	f, err := os.Open(fn)
	if err != nil {
		return "", err
	}
	defer f.Close()

	var h hash.Hash

	switch hashtype {
	case MD5:
		h = md5.New()
	case SHA256:
		h = sha256.New()
	case SHA512:
		h = sha512.New()
	}

	if _, err := io.Copy(h, f); err != nil {
		if err != nil {
			return "", err
		}
	}
	sum := h.Sum(nil)
	return fmt.Sprintf("%x", sum), nil
}

func StringHash(hashtype int, s string) string {

	switch hashtype {
	case MD5:
		sum := md5.Sum([]byte(s))
		return fmt.Sprintf("%x", sum)
	case SHA256:
		sum := sha256.Sum256([]byte(s))
		return fmt.Sprintf("%x", sum)
	case SHA512:
		sum := sha512.Sum512([]byte(s))
		return fmt.Sprintf("%x", sum)
	default:
		return ""
	}

}
