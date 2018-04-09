package scarlet

// State represents the whole compiler state.
type State struct {
	Variables map[string]string
	Constants map[string]string
	Mixins    map[string]*Mixin

	// To preserve the order of appearance, save the names in correct order
	VariableNames []string
}

// NewState creates a new compiler state.
func NewState() *State {
	return &State{
		Variables:     map[string]string{},
		Constants:     map[string]string{},
		Mixins:        map[string]*Mixin{},
		VariableNames: []string{},
	}
}
