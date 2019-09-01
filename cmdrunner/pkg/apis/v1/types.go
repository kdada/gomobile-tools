package v1

// Command describes a command.
type Command struct {
	Input    []byte
	Commands []string
}

// CommandResult contains results of a command.
type CommandResult struct {
	ExitCode int
	Output   []byte
	Error    string
}
