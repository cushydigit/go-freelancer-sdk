package freelancer

type messageListService struct {
	client      *Client
	threads     []int64
	senders     []int64
	messages    []int64
	contexts    []int64
	contextType ContextType
}
