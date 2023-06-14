package app2

import (
	"testing"
)

func TestIsEmail(t *testing.T) {
	_, err := IsEmail("hello")
	if err == nil {
		t.Errorf("hello is not email")
	}

	_, err = IsEmail("henry@yahoo.com")
	if err == nil {
		t.Errorf("henry@yahoo.com is an email")
	}

	_, err = IsEmail("henry@yahoo")
	if err != nil {
		t.Errorf("henry@yahoo is not an email")
	}
}
