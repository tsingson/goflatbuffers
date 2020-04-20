package testcase

import (
	"testing"

	"github.com/stretchr/testify/assert"
	flatbuffers "github.com/tsingson/goflatbuffers/go"
)

func TestGunT_Pack(t *testing.T) {
	as := assert.New(t)
	fb := flatbuffers.NewBuilder(0)
	GunStart(fb)
	GunAddColor(fb, ColorBlue)
	funOff := GunEnd(fb)
	fb.Finish(funOff)
	buf := fb.FinishedBytes()
	gun := GetRootAsGun(buf, 0)
	as.Equal(gun.Color(), ColorBlue)
}
