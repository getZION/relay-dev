package did

import (
	"regexp"

	. "github.com/getzion/relay/utils"
)

const (
	DIDMethodName = "did:zion"
)

type DID struct {
	ID string `json:"id" form:"id" query:"id" validate:"required"`
}

type ParsedDID struct {
	did string
}

func Parse(didUrl string) *ParsedDID {
	match, err := regexp.MatchString("^did:([a-z0-9]+):blah", didUrl)

	Log.Info().Bool("match", match).Str("didUrl", didUrl).Msg("Parse")

	if !match {
		return nil
	}

	d := &ParsedDID{}
	d.did = "Hello"

	if err != nil {
		return nil
	} else {
		return d
	}
}
