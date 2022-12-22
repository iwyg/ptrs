package ptrs

import (
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAny(t *testing.T) {
	v := Any[*int](nil)

	assert.Equal(t, "**int", fmt.Sprintf("%T", v))
	t.Logf("%T", v)

	vv := Any(12)
	assert.Equal(t, "*int", fmt.Sprintf("%T", vv))

	vvv := Any[uint64](12)
	assert.Equal(t, "*uint64", fmt.Sprintf("%T", vvv))

	var nn *uint64
	xn := Any[uint64](0)
	log.Printf("nn %#v \nxn %#v", nn, xn)
}
