package validator

import (
	"regexp"

	v1 "github.com/getzion/relay/gen/proto/zion/v1"
	"github.com/go-playground/validator/v10"
)

var (
	customValidator *validator.Validate
	usernameRegex   = regexp.MustCompile("^[A-Za-z0-9][A-Za-z0-9_-]*$")
)

func InitValidator() {
	customValidator = validator.New()
	customValidator.RegisterValidation("username", usernameValidator)
	customValidator.RegisterStructValidation(conversationORMValidatorFunc, v1.ConversationORM{})
}

func Struct(s interface{}) error {
	return customValidator.Struct(s)
}

func ValidateStruct(s interface{}) error {
	return customValidator.Struct(s)
}

func conversationORMValidatorFunc(sl validator.StructLevel) {
	s := sl.Current().Interface().(v1.ConversationORM)
	v := sl.Validator()
	if err := v.Var(s.Zid, "required"); err != nil {
		sl.ReportValidationErrors("Zid", "Zid", err.(validator.ValidationErrors))
	}
}

func usernameValidator(fl validator.FieldLevel) bool {
	return usernameRegex.MatchString(fl.Field().String())
}
