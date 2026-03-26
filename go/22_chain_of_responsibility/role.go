package chainofresponsibility

type RoleHandle struct {
	BaseHandler
}

func (r *RoleHandle) Handle(req *Request) string {
	if req.Role != "admin" {
		return "no permision"
	}
	return r.Next(req)
}