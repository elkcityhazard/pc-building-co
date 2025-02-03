package models

import "testing"

func Test_NewNavItem(t *testing.T) {

	ni := NewNavItem("Home", "/", 0)

	if ni.Name != "Home" {
		t.Errorf("expected Name to be Home but got %s", ni.Name)
	}

	if ni.Url != "/" {
		t.Errorf("expected Url to be / but got %s", ni.Url)
	}

	if ni.Weight != 0 {
		t.Errorf("Expected weight to be 0 but got %d", ni.Weight)
	}
}
