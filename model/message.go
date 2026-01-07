package model

type UserIdentity struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type BMessage struct {
	Sender   UserIdentity `json:"sender"`
	Receiver UserIdentity `json:"receiver"`
	Message  string       `json:"message"`
}
