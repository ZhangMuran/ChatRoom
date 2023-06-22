package message

type User struct {
	Account  string `json:"account"`
	Password string `json:"password"`
	Status   int    `json:"status"`
}

const (
	OnLine  = 1
	OffLine = 2
)
