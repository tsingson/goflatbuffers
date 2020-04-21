package main

import (
	"fmt"
	"github.com/tsingson/goflatbuffers/go/extras/flexbuffers/types"
)

func main() {

	fbt := uint8(36)
	btw := uint8(3)

	flexT := types.SetFBType(fbt, btw)
	fmt.Println(flexT)

	bitWidth := types.GetBitWidth(flexT)
	fmt.Println("width > ", bitWidth)

	fbType := types.GetFBType(flexT)
	fmt.Println("type > ", fbType)
}
