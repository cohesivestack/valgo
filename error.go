package valgo

type Errors struct {
	message  string
	Field    string   `json:"field"`
	Messages []string `json:"messages"`
}

func (this *Errors) Error() string {
	return this.message
}

func (this *Errors) Add(message string) {
	this.Messages = append(this.Messages, message)
}
