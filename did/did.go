package did

const (
	DIDMethodName = "did:zion"
)

type DID struct {
	ID string `json:"id" form:"id" query:"id" validate:"required"`
}

func Parse() *DID {
	d := &DID{}
	d.ID = "Hello"
	return d
}
