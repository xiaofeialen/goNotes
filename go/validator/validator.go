package validator


import (
	"reflect"
	"strings"
	"sync"

	"github.com/go-playground/locales/zh"
	translator "github.com/go-playground/universal-translator"
	translations "github.com/go-playground/validator/translations/zh"
	validator "gopkg.in/go-playground/validator.v9"
)

var validate *validator.Validate
var trans translator.Translator
var once sync.Once

// InitValidator init validate
func InitValidator() {
	once.Do(func() {
		zh := zh.New()
		ut := translator.New(zh, zh)
		trans, _ = ut.GetTranslator("zh")

		validate = validator.New()
		validate.RegisterTagNameFunc(func(field reflect.StructField) string {
			label := field.Tag.Get("label")
			if label == "" {
				return field.Name
			}
			return label
		})
		translations.RegisterDefaultTranslations(validate, trans)
	})
}

// generateErrMsg generate error massage
func generateErrMsg(errs validator.ValidationErrors) string {
	var errList []string
	for _, e := range errs {
		errList = append(errList, e.Translate(trans))
	}
	return strings.Join(errList, "\n")
}

// GetValidate  use customize massage must get this
func GetValidate() *validator.Validate {
	if validate == nil {
		InitValidator()
	}
	return validate
}

// CheckStructParam check struct param
func CheckStructParam(structData interface{}) (ok bool, errStr string) {
	err := validate.Struct(structData)
	if err == nil {
		return true, ""
	}
	if errs := err.(validator.ValidationErrors); errs != nil {
		errStr = generateErrMsg(errs)
		return
	}
	return true, ""

}

