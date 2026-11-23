package hosting

import (
	"context"
	"errors"

	"github.com/hostinger/api-cli/site"
)

// Reasons a remote deploy (Options.Remote) failed, as FailureReason names them.
const (
	ReasonWouldOverwrite  = "would_overwrite"  // refused: the website has content and Yes was not given
	ReasonUnsupportedType = "unsupported_type" // refused: a WordPress import, or static files onto a Node.js app
	ReasonUploadFailed    = "upload_failed"
	ReasonBuildFailed     = "build_failed"
	ReasonDeployFailed    = "deploy_failed"
	ReasonCancelled       = "cancelled"
)

// RefusedError is a remote deploy refused before anything changed.
type RefusedError struct {
	Reason  string // ReasonWouldOverwrite or ReasonUnsupportedType
	Message string
}

func (e *RefusedError) Error() string { return e.Message }

// uploadError is a failed upload, which FailureReason tells apart. Its
// message is the cause's.
type uploadError struct{ err error }

func (e *uploadError) Error() string { return e.err.Error() }
func (e *uploadError) Unwrap() error { return e.err }

// FailureReason names the step of a deploy that err ended.
func FailureReason(err error) string {
	var refused *RefusedError
	var build *BuildError
	var upload *uploadError
	switch {
	case errors.As(err, &refused):
		return refused.Reason
	case errors.Is(err, context.Canceled):
		return ReasonCancelled
	case errors.As(err, &build):
		return ReasonBuildFailed
	case errors.As(err, &upload):
		return ReasonUploadFailed
	}
	return ReasonDeployFailed
}

// LinkedType is the type Run deploys a linked project's folder src as when
// nobody is asked: a headless website takes Node.js; else the remembered
// type, unless the folder now looks like the other of static and nodejs;
// else what Detect makes of the folder.
func LinkedType(project *site.Project, src string) (Type, error) {
	if project.Headless {
		return NodeJS, nil
	}
	for _, t := range DeployableTypes {
		if string(t) == project.Type {
			if other := otherKind(src, t); other != "" {
				return other, nil
			}
			return t, nil
		}
	}
	t, _, err := Detect(src)
	return t, err
}
