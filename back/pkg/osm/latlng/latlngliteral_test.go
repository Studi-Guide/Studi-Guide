package latlng

import "testing"

func TestNewLatLngLiteral(t *testing.T) {
	var lat float64 = 90.0
	var lng float64 = 180.0

	if _, err := NewLatLngLiteral(lat, lng); err != nil {
		t.Error(err)
	}

	lat = 90.00000001
	lng = 180.00000001

	if _, err := NewLatLngLiteral(lat, lng); err == nil {
		t.Error("expected error because of out of range Lat/Lng")
	}

	lat = -90.0
	lng = -180.0

	if _, err := NewLatLngLiteral(lat, lng); err != nil {
		t.Error(err)
	}

	lat = -90.0000000000001
	lng = -180.000000000000001

	if _, err := NewLatLngLiteral(lat, lng); err == nil {
		t.Error("expected error because of out of range Lat/Lng")
	}
}

func TestParseLatLngLiteral(t *testing.T) {
	latPositive1 := "49.5118"
	lngPositive1 := "11.2167"

	latNegative1 := "49.4126,";
	lngNegative1 := "xyz11.0111";

	latPositive2 := "49"
	lngPositive2 := "-11"

	latNegative2 := "-90.01"
	lngNegative2 := "180.01"

	if _, err := ParseLatLngLiteral(latPositive1, lngPositive1); err != nil {
		t.Error(err)
	}

	if _, err := ParseLatLngLiteral(latNegative1, latPositive2); err == nil {
		t.Error("expected error")
	}

	if _, err := ParseLatLngLiteral(latPositive2, lngPositive2); err != nil {
		t.Error(err)
	}

	if _, err := ParseLatLngLiteral(latPositive2, lngNegative1); err == nil {
		t.Error("expected error")
	}

	if _, err := ParseLatLngLiteral(latNegative2, lngPositive1); err == nil {
		t.Error("expected error")
	}

	if _, err := ParseLatLngLiteral(latPositive1, lngNegative2); err == nil {
		t.Error("expected error")
	}

}

func TestLatLngLiteral_LatLng(t *testing.T) {
	latPositive1 := "49.5118"
	lngPositive1 := "11.2167"

	literal, err := ParseLatLngLiteral(latPositive1, lngPositive1)
	if err != nil {
		t.Error(err)
	}

	if literal.LatStr() != latPositive1 || literal.LngStr() != lngPositive1 {
		t.Error("strings not equal")
	}

	if literal.Lat != 49.5118 || literal.Lng != 11.2167 {
		t.Error("conversion failed")
	}
}