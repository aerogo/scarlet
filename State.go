package scarlet

// State represents the whole compiler state.
type State struct {
	Variables map[string]string
	Constants map[string]string
	Mixins    map[string]*Mixin
}

// NewState creates a new compiler state.
func NewState() *State {
	return &State{
		Variables: make(map[string]string),
		Constants: make(map[string]string),
		Mixins:    make(map[string]*Mixin),
	}
}
