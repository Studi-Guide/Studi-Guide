package locales

import (
	"golang.org/x/text/language"
)

var supportedLocales = []language.Tag{
	language.AmericanEnglish,
	language.German,
}

func GetSupportedLocales() []language.Tag {
	return supportedLocales
}

func GetBestSupportedLocale(locale string) string {
	matcher := language.NewMatcher(supportedLocales)

	lang := language.Make(locale)

	tag, _, _ := matcher.Match(lang)
	base, _ := tag.Base()

	return base.String()
}
