package validation

import (
	"errors"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	go_playground "github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"nashtanet-backend-go/common/types"
	"reflect"
)

type goPlayground struct {
	validator *go_playground.Validate
	translate ut.Translator
	err error
	msg []string
}

func NewGoPlayground() (Validator, error) {
	var (
		language = en.New()
		uni = ut.New(language, language)
		translate, found = uni.GetTranslator("en")
	)

	if !found {
		return nil, errors.New("Translator not found")
	}

	validator := go_playground.New()
	if err := en_translations.RegisterDefaultTranslations(validator, translate); err != nil {
		return nil, errors.New("Translator not found")
	}

	return &goPlayground{validator: validator, translate: translate}, nil
}

func (g *goPlayground) Validate(i interface{}) error {
	if len(g.msg) > 0 {
		g.msg = nil
	}

	g.err = g.validator.Struct(i)
	if g.err != nil {
		return g.err
	}

	return nil
}

func (g *goPlayground) ValidatePartial(i interface{}, fields ...string) error {
	if len(g.msg) > 0 {
		g.msg = nil
	}

	g.err = g.validator.StructPartial(i, fields...)
	if g.err != nil {
		return g.err
	}

	return nil
}

func (g *goPlayground) ValidatePointValue(point types.Point) error {
	if point.X != "" || point.Y != "" {
		g.err = g.Validate(point)
		if g.err != nil {
			return g.err
		}
	}

	return nil
}

func (g *goPlayground) ValidateEnumString(enum []interface{}, value interface{}) bool {
	for _, enumValue := range enum {
		if enumValue.(string) == value.(string) {
			return true
		}
	}
	return false
}

func (g *goPlayground) GetFieldValue(v interface{}, field string) string {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)
	return f.String()
}

func (g *goPlayground) Messages() []string {
	if g.err != nil {
		for _, err := range g.err.(go_playground.ValidationErrors) {
			g.msg = append(g.msg, err.Translate(g.translate))
		}
	}

	return g.msg
}

