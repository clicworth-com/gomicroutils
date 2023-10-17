package grpcclient

type AuthInfo struct {
	Authorised  bool
	Name        string
	EmailId     string
	PhoneNumber string
	Role        string
}
