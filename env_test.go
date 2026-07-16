package main

import "testing"

func TestParseEnvDefaults(t *testing.T) {
	// Blank out both vars so values from the host shell can't leak in.
	// t.Setenv restores the originals when the test ends.
	t.Setenv("ENVIRONMENT", "")
	t.Setenv("PORT", "")

	got, err := parseEnv()

	if err != nil {
		t.Fatalf("parseEnv() error = %v", err)
	}

	if got.Env() != EnvDev {
		t.Errorf("Env() = %q, want %q", got.Env(), EnvDev)
	}

	if got.Port() != 8080 {
		t.Errorf("Port() = %d, want 8080", got.Port())
	}

}

func TestParseEnvInvalidEnvironment(t *testing.T) {
	t.Setenv("ENVIRONMENT", "staging")
	if _, err := parseEnv(); err == nil {
		t.Fatal("expected error for invalid ENVIRONMENT, got nil")
	}
}
