package pool

// Executable is the user standard function
type Executable interface {
	Exec()
	Finished()
}
