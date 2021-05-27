// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hex_test

import (
	"fmt"
	"log"

	hex "github.com/AlekSi/modhex"
)

func ExampleEncode() {
	src := []byte("Hello Gopher!")

	dst := make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(dst, src)

	fmt.Printf("%s\n", dst)

	// Output:
	// fjhghrhrhvdcfihvichjhgiddb
}

func ExampleDecode() {
	src := []byte("fjhghrhrhvdcfihvichjhgiddb")

	dst := make([]byte, hex.DecodedLen(len(src)))
	n, err := hex.Decode(dst, src)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", dst[:n])

	// Output:
	// Hello Gopher!
}

func ExampleDecodeString() {
	const s = "fjhghrhrhvdcfihvichjhgiddb"
	decoded, err := hex.DecodeString(s)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", decoded)

	// Output:
	// Hello Gopher!
}
func ExampleEncodeToString() {
	src := []byte("Hello")
	encodedStr := hex.EncodeToString(src)

	fmt.Printf("%s\n", encodedStr)

	// Output:
	// fjhghrhrhv
}
