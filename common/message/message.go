package message

const (
	LoginSuccess      = 0
	LoginUserNotExist = 1
)

const (
	LoginType    = "LoginMessage"
	LoginRspType = "LoginRspMesssage"
)

type Message struct {
	Type string    `json:"type"`
	Content string `json:"content"`
}

type LoginMessage struct {
	Account string  `json:"account"`
	Password string `json:"password"`
}

type LoginRspMesssage struct {
	Code int      `json:"code"`
	Status string `json:"status"`
}