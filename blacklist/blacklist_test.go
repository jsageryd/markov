package blacklist

import (
	"bytes"
	"testing"
)

func TestBlacklist_AddHas(t *testing.T) {
	b := New()

	items := []string{"foo", "bar", "baz"}

	for n := range items {
		if got, want := b.Has([]byte(items[n])), false; got != want {
			t.Errorf("[before add] b.Has(%q) = %t, want %t", items[n], got, want)
		}

		b.Add([]byte(items[n]))

		if got, want := b.Has([]byte(items[n])), true; got != want {
			t.Errorf("[after add] b.Has(%q) = %t, want %t", items[n], got, want)
		}
	}
}

func TestBlacklist_ImportExportState(t *testing.T) {
	items := []string{"foo", "bar", "baz"}

	var buf bytes.Buffer

	exported := New()

	for n := range items {
		exported.Add([]byte(items[n]))
	}

	if err := exported.ExportState(&buf); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	imported := New()

	if err := imported.ImportState(&buf); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	for n := range items {
		if got, want := imported.Has([]byte(items[n])), true; got != want {
			t.Errorf("imported.Has(%q) = %t, want %t", items[n], got, want)
		}
	}
}
