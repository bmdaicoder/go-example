package state

type State interface {
	Enter()
	Update(sm *StateMachine)
}

type StateMachine struct {
	currentState State
	states       map[string]State
}

func NewStateMachine(initialState State) *StateMachine {
	sm := &StateMachine{
		currentState: initialState,
		states:       make(map[string]State),
	}
	sm.currentState.Enter()
	return sm
}

func (sm *StateMachine) setState(s State) {
	sm.currentState = s
	sm.currentState.Enter()
}

func (sm *StateMachine) Transition() {
	sm.currentState.Update(sm)
}
