package data

import (
	"errors"
	"testing"

	"connectrpc.com/connect/v2"
)

func TestStoreFailed(t *testing.T) {
	err := StoreFailed(t.Context(), "identity.data", "read the identity", errors.New("down"))

	if got := connect.CodeOf(err); got != connect.CodeInternal {
		t.Errorf("code = %v, want %v", got, connect.CodeInternal)
	}

	if got := err.Error(); got != "internal: failed to read the identity" {
		t.Errorf("message = %q, want the action named", got)
	}
}
