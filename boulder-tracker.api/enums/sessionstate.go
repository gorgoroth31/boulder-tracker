package SessionState

type SessionState int

const (
	Live SessionState = iota
	InProgress
	Finished
)

var stateName = map[SessionState]string{
	Live:       "live",
	InProgress: "inProgress",
	Finished:   "Finished",
}

func (state SessionState) String() string {
	return stateName[state]
}
