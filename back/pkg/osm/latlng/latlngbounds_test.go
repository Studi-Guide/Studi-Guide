package latlng

import (
	"reflect"
	"testing"
)

func TestNewLatLngBounds(t *testing.T) {
	southWest, err := NewLatLngLiteral(49.4126, 11.0111)
	if err != nil {
		t.Error(err)
	}
	northEast, err := NewLatLngLiteral(49.5118, 11.2167)
	if err != nil {
		t.Error(err)
	}

	bounds, err := NewLatLngBounds(southWest, northEast)
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(bounds.SouthWest, southWest) || !reflect.DeepEqual(bounds.NorthEast, northEast) {
		t.Error("expected equality")
	}

	if _, err := NewLatLngBounds(northEast, southWest); err == nil {
		t.Error("expected error")
	}

	northEast, err = NewLatLngLiteral(49.5118, 11.00111)
	if err != nil {
		t.Error(err)
	}

	if _, err := NewLatLngBounds(southWest, northEast); err == nil {
		t.Error("expected error")
	}

}

func TestLatLngBounds_IncludeLiteral(t *testing.T) {

	southWest, err := NewLatLngLiteral(49.4126, 11.0111)
	if err != nil {
		t.Error(err)
	}
	northEast, err := NewLatLngLiteral(49.5118, 11.2167)
	if err != nil {
		t.Error(err)
	}

	bounds, err := NewLatLngBounds(southWest, northEast)
	if err != nil {
		t.Error(err)
	}

	testLiteral, err := NewLatLngLiteral(49.4461, 11.082)
	if err != nil {
		t.Error(err)
	}

	if !bounds.IncludeLiteral(testLiteral) {
		t.Error("expected literal to be included")
	}

	testLiteral, err = NewLatLngLiteral(48.13761, 11.5799)

	if bounds.IncludeLiteral(testLiteral) {
		t.Error("expected literal not to be included")
	}
}
