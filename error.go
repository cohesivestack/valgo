package valgo

type Error struct {
	Value    interface{}
	Name     string
	Title    string
	Messages []string
}
