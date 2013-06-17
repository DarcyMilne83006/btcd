// Copyright (c) 2013 Conformal Systems LLC.
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package btcec_test

import (
	"github.com/conformal/btcec"
	"testing"
)

type signatureTest struct {
	name    string
	sig     []byte
	isValid bool
}

var signatureTests = []signatureTest{
	// signatures from bitcoin blockchain tx
	// 0437cd7f8525ceed2324359c2d0ba26006d92d85
	signatureTest{
		name: "valid signature.",
		sig: []byte{0x30, 0x44, 0x02, 0x20, 0x4e, 0x45, 0xe1, 0x69,
			0x32, 0xb8, 0xaf, 0x51, 0x49, 0x61, 0xa1, 0xd3, 0xa1,
			0xa2, 0x5f, 0xdf, 0x3f, 0x4f, 0x77, 0x32, 0xe9, 0xd6,
			0x24, 0xc6, 0xc6, 0x15, 0x48, 0xab, 0x5f, 0xb8, 0xcd,
			0x41, 0x02, 0x20, 0x18, 0x15, 0x22, 0xec, 0x8e, 0xca,
			0x07, 0xde, 0x48, 0x60, 0xa4, 0xac, 0xdd, 0x12, 0x90,
			0x9d, 0x83, 0x1c, 0xc5, 0x6c, 0xbb, 0xac, 0x46, 0x22,
			0x08, 0x22, 0x21, 0xa8, 0x76, 0x8d, 0x1d, 0x09,
		},
		isValid: true,
	},
	signatureTest{
		name:    "empty.",
		sig:     []byte{},
		isValid: false,
	},
	signatureTest{
		name: "bad magic.",
		sig: []byte{0x31, 0x44, 0x02, 0x20, 0x4e, 0x45, 0xe1, 0x69,
			0x32, 0xb8, 0xaf, 0x51, 0x49, 0x61, 0xa1, 0xd3, 0xa1,
			0xa2, 0x5f, 0xdf, 0x3f, 0x4f, 0x77, 0x32, 0xe9, 0xd6,
			0x24, 0xc6, 0xc6, 0x15, 0x48, 0xab, 0x5f, 0xb8, 0xcd,
			0x41, 0x02, 0x20, 0x18, 0x15, 0x22, 0xec, 0x8e, 0xca,
			0x07, 0xde, 0x48, 0x60, 0xa4, 0xac, 0xdd, 0x12, 0x90,
			0x9d, 0x83, 0x1c, 0xc5, 0x6c, 0xbb, 0xac, 0x46, 0x22,
			0x08, 0x22, 0x21, 0xa8, 0x76, 0x8d, 0x1d, 0x09,
		},
		isValid: false,
	},
	signatureTest{
		name: "bad 1st int marker magic.",
		sig: []byte{0x30, 0x44, 0x03, 0x20, 0x4e, 0x45, 0xe1, 0x69,
			0x32, 0xb8, 0xaf, 0x51, 0x49, 0x61, 0xa1, 0xd3, 0xa1,
			0xa2, 0x5f, 0xdf, 0x3f, 0x4f, 0x77, 0x32, 0xe9, 0xd6,
			0x24, 0xc6, 0xc6, 0x15, 0x48, 0xab, 0x5f, 0xb8, 0xcd,
			0x41, 0x02, 0x20, 0x18, 0x15, 0x22, 0xec, 0x8e, 0xca,
			0x07, 0xde, 0x48, 0x60, 0xa4, 0xac, 0xdd, 0x12, 0x90,
			0x9d, 0x83, 0x1c, 0xc5, 0x6c, 0xbb, 0xac, 0x46, 0x22,
			0x08, 0x22, 0x21, 0xa8, 0x76, 0x8d, 0x1d, 0x09,
		},
		isValid: false,
	},
	signatureTest{
		name: "bad 2nd int marker.",
		sig: []byte{0x30, 0x44, 0x02, 0x20, 0x4e, 0x45, 0xe1, 0x69,
			0x32, 0xb8, 0xaf, 0x51, 0x49, 0x61, 0xa1, 0xd3, 0xa1,
			0xa2, 0x5f, 0xdf, 0x3f, 0x4f, 0x77, 0x32, 0xe9, 0xd6,
			0x24, 0xc6, 0xc6, 0x15, 0x48, 0xab, 0x5f, 0xb8, 0xcd,
			0x41, 0x03, 0x20, 0x18, 0x15, 0x22, 0xec, 0x8e, 0xca,
			0x07, 0xde, 0x48, 0x60, 0xa4, 0xac, 0xdd, 0x12, 0x90,
			0x9d, 0x83, 0x1c, 0xc5, 0x6c, 0xbb, 0xac, 0x46, 0x22,
			0x08, 0x22, 0x21, 0xa8, 0x76, 0x8d, 0x1d, 0x09,
		},
		isValid: false,
	},
	signatureTest{
		name: "short len",
		sig: []byte{0x30, 0x43, 0x02, 0x20, 0x4e, 0x45, 0xe1, 0x69,
			0x32, 0xb8, 0xaf, 0x51, 0x49, 0x61, 0xa1, 0xd3, 0xa1,
			0xa2, 0x5f, 0xdf, 0x3f, 0x4f, 0x77, 0x32, 0xe9, 0xd6,
			0x24, 0xc6, 0xc6, 0x15, 0x48, 0xab, 0x5f, 0xb8, 0xcd,
			0x41, 0x02, 0x20, 0x18, 0x15, 0x22, 0xec, 0x8e, 0xca,
			0x07, 0xde, 0x48, 0x60, 0xa4, 0xac, 0xdd, 0x12, 0x90,
			0x9d, 0x83, 0x1c, 0xc5, 0x6c, 0xbb, 0xac, 0x46, 0x22,
			0x08, 0x22, 0x21, 0xa8, 0x76, 0x8d, 0x1d, 0x09,
		},
		isValid: false,
	},
	signatureTest{
		name: "long len",
		sig: []byte{0x30, 0x45, 0x02, 0x20, 0x4e, 0x45, 0xe1, 0x69,
			0x32, 0xb8, 0xaf, 0x51, 0x49, 0x61, 0xa1, 0xd3, 0xa1,
			0xa2, 0x5f, 0xdf, 0x3f, 0x4f, 0x77, 0x32, 0xe9, 0xd6,
			0x24, 0xc6, 0xc6, 0x15, 0x48, 0xab, 0x5f, 0xb8, 0xcd,
			0x41, 0x02, 0x20, 0x18, 0x15, 0x22, 0xec, 0x8e, 0xca,
			0x07, 0xde, 0x48, 0x60, 0xa4, 0xac, 0xdd, 0x12, 0x90,
			0x9d, 0x83, 0x1c, 0xc5, 0x6c, 0xbb, 0xac, 0x46, 0x22,
			0x08, 0x22, 0x21, 0xa8, 0x76, 0x8d, 0x1d, 0x09,
		},
		isValid: false,
	},
	signatureTest{
		name: "long X",
		sig: []byte{0x30, 0x44, 0x02, 0x42, 0x4e, 0x45, 0xe1, 0x69,
			0x32, 0xb8, 0xaf, 0x51, 0x49, 0x61, 0xa1, 0xd3, 0xa1,
			0xa2, 0x5f, 0xdf, 0x3f, 0x4f, 0x77, 0x32, 0xe9, 0xd6,
			0x24, 0xc6, 0xc6, 0x15, 0x48, 0xab, 0x5f, 0xb8, 0xcd,
			0x41, 0x02, 0x20, 0x18, 0x15, 0x22, 0xec, 0x8e, 0xca,
			0x07, 0xde, 0x48, 0x60, 0xa4, 0xac, 0xdd, 0x12, 0x90,
			0x9d, 0x83, 0x1c, 0xc5, 0x6c, 0xbb, 0xac, 0x46, 0x22,
			0x08, 0x22, 0x21, 0xa8, 0x76, 0x8d, 0x1d, 0x09,
		},
		isValid: false,
	},
	signatureTest{
		name: "long Y",
		sig: []byte{0x30, 0x44, 0x02, 0x20, 0x4e, 0x45, 0xe1, 0x69,
			0x32, 0xb8, 0xaf, 0x51, 0x49, 0x61, 0xa1, 0xd3, 0xa1,
			0xa2, 0x5f, 0xdf, 0x3f, 0x4f, 0x77, 0x32, 0xe9, 0xd6,
			0x24, 0xc6, 0xc6, 0x15, 0x48, 0xab, 0x5f, 0xb8, 0xcd,
			0x41, 0x02, 0x21, 0x18, 0x15, 0x22, 0xec, 0x8e, 0xca,
			0x07, 0xde, 0x48, 0x60, 0xa4, 0xac, 0xdd, 0x12, 0x90,
			0x9d, 0x83, 0x1c, 0xc5, 0x6c, 0xbb, 0xac, 0x46, 0x22,
			0x08, 0x22, 0x21, 0xa8, 0x76, 0x8d, 0x1d, 0x09,
		},
		isValid: false,
	},
	signatureTest{
		name: "short Y",
		sig: []byte{0x30, 0x44, 0x02, 0x20, 0x4e, 0x45, 0xe1, 0x69,
			0x32, 0xb8, 0xaf, 0x51, 0x49, 0x61, 0xa1, 0xd3, 0xa1,
			0xa2, 0x5f, 0xdf, 0x3f, 0x4f, 0x77, 0x32, 0xe9, 0xd6,
			0x24, 0xc6, 0xc6, 0x15, 0x48, 0xab, 0x5f, 0xb8, 0xcd,
			0x41, 0x02, 0x19, 0x18, 0x15, 0x22, 0xec, 0x8e, 0xca,
			0x07, 0xde, 0x48, 0x60, 0xa4, 0xac, 0xdd, 0x12, 0x90,
			0x9d, 0x83, 0x1c, 0xc5, 0x6c, 0xbb, 0xac, 0x46, 0x22,
			0x08, 0x22, 0x21, 0xa8, 0x76, 0x8d, 0x1d, 0x09,
		},
		isValid: false,
	},
	signatureTest{
		name: "trailing crap.",
		sig: []byte{0x30, 0x44, 0x02, 0x20, 0x4e, 0x45, 0xe1, 0x69,
			0x32, 0xb8, 0xaf, 0x51, 0x49, 0x61, 0xa1, 0xd3, 0xa1,
			0xa2, 0x5f, 0xdf, 0x3f, 0x4f, 0x77, 0x32, 0xe9, 0xd6,
			0x24, 0xc6, 0xc6, 0x15, 0x48, 0xab, 0x5f, 0xb8, 0xcd,
			0x41, 0x02, 0x20, 0x18, 0x15, 0x22, 0xec, 0x8e, 0xca,
			0x07, 0xde, 0x48, 0x60, 0xa4, 0xac, 0xdd, 0x12, 0x90,
			0x9d, 0x83, 0x1c, 0xc5, 0x6c, 0xbb, 0xac, 0x46, 0x22,
			0x08, 0x22, 0x21, 0xa8, 0x76, 0x8d, 0x1d, 0x09, 0x01,
		},
		isValid: false,
	},
}

func TestSignatures(t *testing.T) {
	for _, test := range signatureTests {
		_, err := btcec.ParseSignature(test.sig, btcec.S256())
		if err != nil {
			if test.isValid {
				t.Errorf("%s signature failed when shouldn't %v",
					test.name, err)
			} /* else {
				t.Errorf("%s got error %v", test.name, err)
			} */
			continue
		}
		if !test.isValid {
			t.Errorf("%s counted as valid when it should fail",
				test.name)
		}
	}
}
