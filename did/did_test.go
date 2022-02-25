package did

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseDID(t *testing.T) {
	// Returns nil if non compliant
	did := Parse("did:")
	assert.Nil(t, did)

	did = Parse("did:zion")
	assert.Nil(t, did)
}
