package chainofresponsibility

type ParamHandler struct {
	BaseHandler
}

func (p *ParamHandler) Handle(req  *Request) string {
	if req.Param == "" {
		return "invalid param"
	}
	return p.Next(req)
}