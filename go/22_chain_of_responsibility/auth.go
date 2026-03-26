package chainofresponsibility

type AuthHandler struct {
	BaseHandler
}

func (a *AuthHandler) Handle(req *Request) string {
	if !req.IsLogin {
		return "not login"
	}
	return a.Next(req)
}