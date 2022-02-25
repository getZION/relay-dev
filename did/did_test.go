package did

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseDID(t *testing.T) {
	did := Parse()
	assert.NotEmpty(t, did)
}
