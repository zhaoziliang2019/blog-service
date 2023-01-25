package app

import (
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type ValidError struct {
	Key     string
	Message string
}
type ValidErrors []*ValidError

func (v *ValidError) Error() string {
	return v.Message
}
func (v ValidErrors) Errors() []string {
	var errs []string
	for _, err := range v {
		errs = append(errs, err.Error())
	}
	return errs
}
func BindAndValid(ctx *gin.Context, v interface{}) (bool, ValidErrors) {
	var errs ValidErrors
	err := ctx.ShouldBind(v)
	if err != nil {
		v := ctx.Value("trans")
		trans, _ := v.(ut.Translator)
		verrs, ok := err.(validator.ValidationErrors)
		if !ok {
			return true, nil
		}
		for key, value := range verrs.Translate(trans) {
			errs = append(errs, &ValidError{
				Key:     key,
				Message: value,
			})
		}
		return true, errs
	}
	return false, nil
}
