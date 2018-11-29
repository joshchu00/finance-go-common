package logger

import (
	"testing"
)

func TestLogger(t *testing.T) {
	Init(".", "test")

	Info("test text")
}
