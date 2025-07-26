package autolykos2

import (
	"bytes"
	"testing"

	"github.com/snipa22/powkit/support/common/testutil"
)

func TestCalcN(t *testing.T) {
	tests := []struct {
		nBase  uint32
		height uint64
		value  uint32
	}{
		{
			nBase:  1 << 26,
			height: 500000,
			value:  67108864,
		},
		{
			nBase:  1 << 26,
			height: 600000,
			value:  67108864,
		},
		{
			nBase:  1 << 26,
			height: 614400,
			value:  70464240,
		},
		{
			nBase:  1 << 26,
			height: 665600,
			value:  73987410,
		},
		{
			nBase:  1 << 26,
			height: 700000,
			value:  73987410,
		},
		{
			nBase:  1 << 26,
			height: 788400,
			value:  81571035,
		},
		{
			nBase:  1 << 26,
			height: 1051200,
			value:  104107290,
		},
		{
			nBase:  1 << 26,
			height: 4198400,
			value:  2143944600,
		},
		{
			nBase:  1 << 26,
			height: 41984000,
			value:  2143944600,
		},
	}

	for i, tt := range tests {
		value := calcN(tt.nBase, tt.height)
		if value != tt.value {
			t.Errorf("failed on %d: have %d, want %d", i, value, tt.value)
		}
	}
}

func TestComputeErgo(t *testing.T) {
	tests := []struct {
		msg    []byte
		nonce  uint64
		height uint64
		result []byte
	}{
		{
			msg:    testutil.MustDecodeHex("548c3e602a8f36f8f2738f5f643b02425038044d98543a51cabaa9785e7e864f"),
			nonce:  0x3105,
			height: 614400,
			result: testutil.MustDecodeHex("0002fcb113fe65e5754959872dfdbffea0489bf830beb4961ddc0e9e66a1412a"),
		},
		{
			msg:    testutil.MustDecodeHex("8e26ba46cd2516cce5c0573dc92c6de8f1b7f110bee9aca4d196e5e0e391d029"),
			nonce:  0x062360e36e133e4d,
			height: 771223,
			result: testutil.MustDecodeHex("000000006b216892578fad42f928c9a19638a2abb667a5a8113311393dcca017"),
		},
		{
			msg:    testutil.MustDecodeHex("8e26ba46cd2516cce5c0573dc92c6de8f1b7f110bee9aca4d196e5e0e391d029"),
			nonce:  0x062360e385ffa5ce,
			height: 771223,
			result: testutil.MustDecodeHex("0000000005353a7890377419c1ddfdec0e185fa445886597a4ed69e376190aa1"),
		},
	}

	for i, tt := range tests {
		result, err := NewErgo().Compute(tt.msg, tt.height, tt.nonce)
		if err != nil {
			t.Errorf("failed on %d: %v", i, err)
		} else if bytes.Compare(result, tt.result) != 0 {
			t.Errorf("failed on %d: have %x, want %x", i, result, tt.result)
		}
	}
}
