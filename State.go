package scarlet

// State represents the whole compiler state.
type State struct {
	Variables map[string]string
}

// NewState creates a new compiler state.
func NewState() *State {
	return &State{
		Variables: make(map[string]string),
	}
}
