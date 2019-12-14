package model

type SendInvite struct {
	FromName  string `json:"from_name"`
	FromEmail string `json:"from_email"`
	To        []*To  `json:"to,omitempty"`
	Subject   string `json:"subject"`
	Message   string `json:"message"`
}

type To struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}
