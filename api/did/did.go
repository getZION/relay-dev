package did

import (
	"encoding/base64"
	"regexp"
	"strings"

	"github.com/sirupsen/logrus"
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
	DID_MATCHER     = `^did:` + METHOD + `:` + METHOD_ID_TEMP2 + QUERY + FRAGMENT + `$`
)

type ParsedDID struct {
	did     string
	didUrl  string
	method  string
	id      string
	service string
	queries string
}

func Parse(didUrl string) *ParsedDID {
	r := regexp.MustCompile(DID_MATCHER)
	match := r.MatchString(didUrl)

	logrus.Infof("match: %t, didUrl: %s", match, didUrl)

	if !match {
		return nil
	}

	matches := r.FindAllStringSubmatch(didUrl, -1)

	d := &ParsedDID{}
	d.did = "did:" + matches[0][1] + ":" + matches[0][2]
	d.didUrl = didUrl
	d.method = matches[0][1]
	d.id = matches[0][2]

	if strings.Contains(didUrl, "?") {
		chop1 := strings.Split(matches[0][3], "?service=")
		chop2 := strings.Split(chop1[1], "&")
		service := chop2[0]
		d.service = service

		chop3 := strings.Split(matches[0][3], "&queries=")
		queries := chop3[1]

		wat, _ := base64.StdEncoding.DecodeString(queries)
		d.queries = string(wat)
	}

	return d
}
