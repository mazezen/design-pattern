package chainofresponsibility

type FinalHandler struct {
	BaseHandler
}

func (f *FinalHandler) Handle(req *Request) string {
	return "success"
}