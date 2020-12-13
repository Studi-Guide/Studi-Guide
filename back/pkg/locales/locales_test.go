package locales

import "testing"

func TestGetBestSupportedLocale(t *testing.T) {
	french := "fr"
	german := "de"
	swissGerman := "gsw"
	americanEnglish := "en-US"
	english := "en"
	germanGerman := "de-DE"

	if GetBestSupportedLocale(french) != "en" {
		t.Error("fallback should be american english")
	}

	if GetBestSupportedLocale(german) != "de" {
		t.Error("german locale should be supported")
	}

	if GetBestSupportedLocale(swissGerman) != "de" {
		t.Error("fallback for swiss german should be german")
	}

	if GetBestSupportedLocale(americanEnglish) != "en" {
		t.Error("language for american english should be english")
	}

	if GetBestSupportedLocale(english) != "en" {
		t.Error("fallback for english should be american english")
	}

	if GetBestSupportedLocale("") != "en" {
		t.Error("fallback for no locale should be american english")
	}

	if GetBestSupportedLocale(germanGerman) != "de" {
		t.Error("expected german for de-de")
	}

}