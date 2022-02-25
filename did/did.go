package did

import (
	"regexp"

	. "github.com/getzion/relay/utils"
)

const (
	DIDMethodName   = "did:zion"
	PCT_ENCODED     = `(?:%[0-9a-fA-F]{2})`
	ID_CHAR         = `(?:[a-zA-Z0-9._-]|` + PCT_ENCODED
	METHOD          = `([a-z0-9]+)`
	METHOD_ID       = `((?:` + ID_CHAR + `*:)*(` + ID_CHAR + `+))`
	METHOD_ID_TEMP  = ID_CHAR + `*:)*(` + ID_CHAR + `+))`
	METHOD_ID_TEMP2 = `(\w+)`
	PARAM_CHAR      = `[a-zA-Z0-9_.:%-]`
	PARAM           = `;` + PARAM_CHAR + `+=` + PARAM_CHAR + `*`
	PARAMS          = `((` + PARAM + `)*)`
	PATH            = `(/[^#?]*)?`
	QUERY           = `([?][^#]*)?`
	FRAGMENT        = `(#.*)?`
	DID_MATCHER     = `^did:` + METHOD + `:` + METHOD_ID_TEMP2 + `$`
)

type ParsedDID struct {
	did    string
	didUrl string
	method string
	id     string
}

func Parse(didUrl string) *ParsedDID {
	r := regexp.MustCompile(DID_MATCHER)
	match := r.MatchString(didUrl)

	Log.Info().Bool("match", match).Str("didUrl", didUrl).Msg("Parse")

	if !match {
		return nil
	}

	matches := r.FindAllStringSubmatch(didUrl, -1)

	d := &ParsedDID{}
	d.did = matches[0][0]
	d.didUrl = matches[0][0]
	d.method = matches[0][1]
	d.id = matches[0][2]
	return d
}
