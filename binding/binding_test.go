package binding

import (
	"os"
	"testing"
)

func TestError(t *testing.T) {
	_, err := NewServiceBinding()
	if err == nil {
		t.Error("expected error")
	}
}

func TestAllBindings(t *testing.T) {
	td := t.TempDir()
	os.Setenv("SERVICE_BINDING_ROOT", td)
	sb, err := NewServiceBinding()
	if err != nil {
		t.Error("expected no error")
	}
	sb.AllBindings()
}
