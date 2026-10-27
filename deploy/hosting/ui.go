package hosting

// UI is how a deploy asks questions and shows progress.
type UI interface {
	// Interactive reports whether Select, Input and Confirm may be called.
	Interactive() bool
	Select(title string, options []string, def int) (int, error)
	// Input returns def for an empty answer.
	Input(title, def string) (string, error)
	Confirm(title string, def bool) (bool, error)
	Info(msg string)
	// Start begins a long step, ended by Done or Fail.
	Start(msg string) Step
}

type Step interface {
	Update(msg string)
	Done(msg string)
	Fail(msg string)
}

const RelinkHint = "Run `hostinger link` to change it."
