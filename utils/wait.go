package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"slices"
	"time"

	"github.com/spf13/cobra"
)

// WaitSpec describes how to recognise the end state of an asynchronous resource.
type WaitSpec struct {
	IDField    string
	StateField string
	Done       []string
	Failed     []string
}

// ErrActionFailed is returned by Wait when the resource ends in a failed state.
var ErrActionFailed = errors.New("action failed")

// Poll intervals and the progress stream are package vars for tests.
var (
	waitFirstInterval = 2 * time.Second
	waitMaxInterval   = 10 * time.Second
)

// RegisterWaitFlags adds --wait and --wait-timeout to cmd.
func RegisterWaitFlags(cmd *cobra.Command) {
	cmd.Flags().Bool("wait", false, "Wait until the action finishes")
	cmd.Flags().Duration("wait-timeout", 10*time.Minute, "Maximum time to wait with --wait")
}

// Wait polls an asynchronous resource until its state is one of spec.Done or
// spec.Failed. body is the response that started it; poll fetches the
// resource by the ID read from body. It returns the last fetched body; when
// the final state is a failure the error is ErrActionFailed (wrapped).
func Wait(cmd *cobra.Command, body []byte, spec WaitSpec, poll func(ctx context.Context, id string) ([]byte, int, error)) ([]byte, int, error) {
	timeout, _ := cmd.Flags().GetDuration("wait-timeout")
	ctx, cancel := context.WithTimeout(cmd.Context(), timeout)
	defer cancel()

	id, state := fieldString(body, spec.IDField), fieldString(body, spec.StateField)
	if id == "" {
		return body, 200, fmt.Errorf("--wait: response has no %q field", spec.IDField)
	}
	status := 200
	interval := waitFirstInterval
	last := ""
	for {
		if state != last {
			fmt.Fprintf(promptOut, "Action %s: %s\n", id, state)
			last = state
		}
		if slices.Contains(spec.Done, state) {
			return body, status, nil
		}
		if slices.Contains(spec.Failed, state) {
			return body, status, fmt.Errorf("%w: action %s ended in state %q", ErrActionFailed, id, state)
		}
		select {
		case <-ctx.Done():
			return body, status, waitStopped(ctx, timeout, state)
		case <-time.After(interval):
		}
		interval = min(interval*2, waitMaxInterval)

		b, st, err := poll(ctx, id)
		if err != nil {
			if ctx.Err() != nil {
				return body, status, waitStopped(ctx, timeout, state)
			}
			return nil, 0, err
		}
		if st >= 300 {
			return b, st, nil
		}
		body, status, state = b, st, fieldString(b, spec.StateField)
	}
}

// waitStopped explains why waiting ended early, keeping the last known state.
func waitStopped(ctx context.Context, timeout time.Duration, state string) error {
	if errors.Is(ctx.Err(), context.Canceled) {
		return fmt.Errorf("--wait: interrupted; last state %q", state)
	}
	return fmt.Errorf("--wait: gave up after %s; last state %q", timeout, state)
}

// fieldString reads a top-level field of a JSON object (or of its "data"
// object) as a string; numbers keep their exact digits.
func fieldString(body []byte, field string) string {
	dec := json.NewDecoder(bytes.NewReader(body))
	dec.UseNumber()
	var obj map[string]any
	if err := dec.Decode(&obj); err != nil {
		return ""
	}
	if data, ok := obj["data"].(map[string]any); ok {
		obj = data
	}
	v, ok := obj[field]
	if !ok || v == nil {
		return ""
	}
	return fmt.Sprintf("%v", v)
}
