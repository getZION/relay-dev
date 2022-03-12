package permissions

import (
	"fmt"
	"net/url"

	"github.com/getzion/relay/api/identityhub/errors"
	"github.com/getzion/relay/api/identityhub/handler"
)

func PermissionsGrant(context *handler.RequestContext) ([]string, *errors.MessageLevelError) {

	var err error
	var schema *url.URL

	if schema, err = url.ParseRequestURI(context.Message.Descriptor_.Schema); err != nil {
		return nil, errors.NewMessageLevelError(400, errors.ImproperlyConstructedErrorMessage, err)
	}

	fmt.Printf("request -> schema: %s", schema.String())
	return nil, nil
}
