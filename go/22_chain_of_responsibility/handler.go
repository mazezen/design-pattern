package chainofresponsibility

type Handler interface {
	SetNext(handler Handler)
	Handle(req *Request) string
}

type BaseHandler struct {
	next Handler
}

func (b *BaseHandler) SetNext(handler Handler) {
	b.next = handler
}

func (b *BaseHandler) Next(req *Request) string {
	if b.next != nil {
		return b.next.Handle(req)
	}

	return "ok"
}
