package testcase

import (
	"testing"

	"github.com/tsingson/goflatbuffers/go"

	"github.com/stretchr/testify/assert"
)

func TestNewVirtualTable(t *testing.T) {
	as := assert.New(t)
	v := flatbuffers.NewVirtualTable()
	v.StartTable(5)
	v.EndTable().Finish()
	as.False(v.AllSet())
	as.Equal(v.ObjectSize(), 0)
}
