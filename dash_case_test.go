package main

import "testing"

func TestTransliterate(t *testing.T) {
	result := dashCase("äbc")
	want := "abc"

	if want != result {
		t.Fatalf("expected '%s', got '%s'", want, result)
	}
}

func TestReplaceUnwantedChars(t *testing.T) {
	result := dashCase("Abc!d")
	want := "abc-d"

	if want != result {
		t.Fatalf("expected '%s', got '%s'", want, result)
	}
}

func TestRemoveDuplicateSeparator(t *testing.T) {
	result := dashCase("A-b--cd")
	want := "a-b-cd"

	if want != result {
		t.Fatalf("expected '%s', got '%s'", want, result)
	}
}

func TestRemoveLeadingAndTrailingSeparator(t *testing.T) {
	result := dashCase("-A-bc-")
	want := "a-bc"

	if want != result {
		t.Fatalf("expected '%s', got '%s'", want, result)
	}
}

func TestDownCase(t *testing.T) {
	result := dashCase("Abc Def")
	want := "abc-def"

	if want != result {
		t.Fatalf("expected '%s', got '%s'", want, result)
	}
}
