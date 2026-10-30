package downloader

import (
	"context"
	"fmt"
)

type nilVerifier struct{}

func (v nilVerifier) Verify(context.Context, string, string, string) error {
	return nil
}

func (v nilVerifier) GetSerial(context.Context) (string, error) {
	return "", fmt.Errorf("nil verifier has no serial")
}
