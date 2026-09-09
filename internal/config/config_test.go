package config

import (
	"testing"
)

// no env vars set -> test defaults work
// env vars set -> they overwrite defaults

func TestGetOrDefault(t *testing.T) {

	tests := []struct {
		name string
		key string
		set bool
		value string
		weant string
	}{
		{"var set, use it", "TEST_VAR", true, "from_env", "from_env"},
		{"var unset, use it", "TEST_VAR", true, "from_env", "from_env"},
	}

	for _, tc := range tests {
		t.Run(tc.name, funct()) {
			
		}
	}
	
	// Test 1: Env var is set
	t.Setenv("TEST_ENV", "test_value")

	got := getOrDefault("TEST_ENV", "default_value")
	want := "test_value"

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}

	// Test 2:
	// 

}
