package action

type Action struct {
	Args   []string
	Action func()
}
