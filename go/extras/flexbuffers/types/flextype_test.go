package types

import (
	"fmt"
	"testing"
)

func TestGetFBType(t *testing.T) {
	fbt := uint8(3)
	btw := uint8(3)

	fbType := SetFBType(fbt, btw)

	bitWidth := GetBitWidth(fbType)

	if bitWidth != btw {
		t.Error("fbt error")
	}

	fbType1 := GetFBType(fbType)

	if fbType1 != fbt {
		t.Error("fbt error")
	}
	if IsInline(FBT( fbt )) {
		fmt.Println("--------------------- is inline")
	}
}

func TestGetBitWidth(t *testing.T) {
	fbt := uint8(1)
	btw := uint8(0)

	fbType := SetFBType(fbt, btw)

	bitWidth := GetBitWidth(fbType)

	if bitWidth != btw {
		t.Error("fbt error")
	}

	fbType1 := GetFBType(fbType)

	if fbType1 != fbt {
		t.Error("fbt error")
	}
}
