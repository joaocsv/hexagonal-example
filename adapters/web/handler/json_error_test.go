package handler

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHandler_errorJson(t *testing.T) {
	msg := "Hello Json"

	result := errorJson(msg)

	require.Equal(t, []byte(`{"message":"Hello Json"}`), result)
}
