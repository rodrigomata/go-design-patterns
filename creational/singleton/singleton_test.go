package singleton

import "testing"

// TestGetInstance test GetInstance() method
func TestGetInstance(t *testing.T) {
	test := GetInstance()
	if test == nil {
		t.Error("expected pointer to Singleton after calling GetInstance(), not nil")
	}
}
