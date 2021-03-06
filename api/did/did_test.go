package did

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseDID(t *testing.T) {
	// Returns parts
	did := Parse("did:zion:2nQtiQG6Cgm1GYTBaaKAgr76uY7iSexUkqX")
	assert.NotNil(t, did)
	assert.Equal(t, &ParsedDID{
		did:    "did:zion:2nQtiQG6Cgm1GYTBaaKAgr76uY7iSexUkqX",
		didUrl: "did:zion:2nQtiQG6Cgm1GYTBaaKAgr76uY7iSexUkqX",
		id:     "2nQtiQG6Cgm1GYTBaaKAgr76uY7iSexUkqX",
		method: "zion",
	}, did)

	did = Parse("did:zion:2nQtiQG6Cgm1GYTBaaKAgr76uY7iSexUkqX?service=IdentityHub&queries=W3sgIm1ldGhvZCI6ICJDb2xsZWN0aW9uc1F1ZXJ5IiwgInNjaGVtYSI6ICJodHRwczovL3NjaGVtYS5vcmcvU29jaWFsTWVkaWFQb3N0aW5nIiB9XQ==")
	assert.NotNil(t, did)
	assert.Equal(t, &ParsedDID{
		did:     "did:zion:2nQtiQG6Cgm1GYTBaaKAgr76uY7iSexUkqX",
		didUrl:  "did:zion:2nQtiQG6Cgm1GYTBaaKAgr76uY7iSexUkqX?service=IdentityHub&queries=W3sgIm1ldGhvZCI6ICJDb2xsZWN0aW9uc1F1ZXJ5IiwgInNjaGVtYSI6ICJodHRwczovL3NjaGVtYS5vcmcvU29jaWFsTWVkaWFQb3N0aW5nIiB9XQ==",
		id:      "2nQtiQG6Cgm1GYTBaaKAgr76uY7iSexUkqX",
		method:  "zion",
		service: "IdentityHub",
		queries: "[{ \"method\": \"CollectionsQuery\", \"schema\": \"https://schema.org/SocialMediaPosting\" }]",
	}, did)

	// Returns nil if non compliant
	did = Parse("")
	assert.Nil(t, did)

	did = Parse("did:")
	assert.Nil(t, did)

	did = Parse("did:zion")
	assert.Nil(t, did)

	did = Parse("did:zion:")
	assert.Nil(t, did)

	did = Parse("did:zion:1234_12313***")
	assert.Nil(t, did)

	did = Parse("2nQtiQG6Cgm1GYTBaaKAgr76uY7iSexUkqX")
	assert.Nil(t, did)

	did = Parse("did:method:%12%1")
	assert.Nil(t, did)

	did = Parse("did:method:%1233%Ay")
	assert.Nil(t, did)

	did = Parse("did:CAP:id")
	assert.Nil(t, did)

	did = Parse("did:method:id::anotherid%r9")
	assert.Nil(t, did)
}
