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
	customValidator.RegisterStructValidation(communityORMValidatorFunc, v1.CommunityORM{})
	customValidator.RegisterStructValidation(conversationORMValidatorFunc, v1.ConversationORM{})
	customValidator.RegisterStructValidation(userORMValidatorFunc, v1.UserORM{})
}

func ValidateStruct(s interface{}) error {
	return customValidator.Struct(s)
}

func communityORMValidatorFunc(sl validator.StructLevel) {
	s := sl.Current().Interface().(v1.CommunityORM)
	v := sl.Validator()
	if err := v.Var(s.Name, "required,max=150"); err != nil {
		sl.ReportValidationErrors("Name", "Name", err.(validator.ValidationErrors))
	}
	if err := v.Var(s.Description, "max=250"); err != nil {
		sl.ReportValidationErrors("Description", "Description", err.(validator.ValidationErrors))
	}
	if err := v.Var(s.EscrowAmount, "gte=0,lt=100000"); err != nil {
		sl.ReportValidationErrors("EscrowAmount", "EscrowAmount", err.(validator.ValidationErrors))
	}
	if err := v.Var(s.OwnerDid, "required"); err != nil {
		sl.ReportValidationErrors("OwnerDid", "OwnerDid", err.(validator.ValidationErrors))
	}
	if err := v.Var(s.OwnerUsername, "required"); err != nil {
		sl.ReportValidationErrors("OwnerUsername", "OwnerUsername", err.(validator.ValidationErrors))
	}
	if err := v.Var(s.PricePerMessage, "gte=0,lt=100000"); err != nil {
		sl.ReportValidationErrors("PricePerMessage", "PricePerMessage", err.(validator.ValidationErrors))
	}
	if err := v.Var(s.PriceToJoin, "gte=0,lt=100000"); err != nil {
		sl.ReportValidationErrors("PriceToJoin", "PriceToJoin", err.(validator.ValidationErrors))
	}
}

func conversationORMValidatorFunc(sl validator.StructLevel) {
	s := sl.Current().Interface().(v1.ConversationORM)
	v := sl.Validator()
	if err := v.Var(s.Zid, "required"); err != nil {
		sl.ReportValidationErrors("Zid", "Zid", err.(validator.ValidationErrors))
	}
}

func userORMValidatorFunc(sl validator.StructLevel) {
	s := sl.Current().Interface().(v1.UserORM)
	v := sl.Validator()
	if err := v.Var(s.Name, "required"); err != nil {
		sl.ReportValidationErrors("Name", "Name", err.(validator.ValidationErrors))
	}
	if err := v.Var(s.Email, "email"); err != nil {
		sl.ReportValidationErrors("Email", "Email", err.(validator.ValidationErrors))
	}
	if err := v.Var(s.Username, "required,username,min=6,max=16"); err != nil {
		sl.ReportValidationErrors("Username", "Username", err.(validator.ValidationErrors))
	}
}

func usernameValidator(fl validator.FieldLevel) bool {
	return usernameRegex.MatchString(fl.Field().String())
}
