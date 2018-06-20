package pandorabox

// Message is response message on http requests
type Message struct {
	Content string
	Status  string
	Body    interface{}
}
