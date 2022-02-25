package did

import (
	"regexp"

	. "github.com/getzion/relay/utils"
)

const (
	DIDMethodName  = "did:zion"
	PCT_ENCODED    = `(?:%[0-9a-fA-F]{2})`
	ID_CHAR        = `(?:[a-zA-Z0-9._-]|` + PCT_ENCODED
	METHOD         = `([a-z0-9]+)`
	METHOD_ID      = `((?:` + ID_CHAR + `*:)*(` + ID_CHAR + `+))`
	METHOD_ID_TEMP = ID_CHAR + `*:)*(` + ID_CHAR + `+))`
	PARAM_CHAR     = `[a-zA-Z0-9_.:%-]`
	PARAM          = `;` + PARAM_CHAR + `+=` + PARAM_CHAR + `*`
	PARAMS         = `((` + PARAM + `)*)`
	PATH           = `(/[^#?]*)?`
	QUERY          = `([?][^#]*)?`
	FRAGMENT       = `(#.*)?`
	DID_MATCHER    = `^did:` + METHOD + `:` + METHOD_ID_TEMP + `$`
)

type ParsedDID struct {
	did    string
	didUrl string
	method string
	id     string
}

func Parse(didUrl string) *ParsedDID {
	Log.Info().Str("DID_MATCHER", DID_MATCHER).Msg("Parse")
	match, err := regexp.MatchString(DID_MATCHER, didUrl)

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
