package locales

import "testing"

func TestGetBestSupportedLocale(t *testing.T) {
	french := "fr"
	german := "de"
	swissGerman := "gsw"
	americanEnglish := "en-US"
	english := "en"

	if GetBestSupportedLocale(french) != "en-US" {
		t.Error("fallback should be american english")
	}

	if GetBestSupportedLocale(german) != "de" {
		t.Error("german locale should be supported")
	}

	if GetBestSupportedLocale(swissGerman) != "de" {
		t.Error("fallback for swiss german should be german")
	}

	if GetBestSupportedLocale(americanEnglish) != "en-US" {
		t.Error("language for american english should be english")
	}

	if GetBestSupportedLocale(english) != "en-US" {
		t.Error("fallback for english should be american english")
	}

}