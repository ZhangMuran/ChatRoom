package message

const (
	LoginSuccess    = 0
	LoginFail       = 1
	RegisterSuccess = 10
	RegisterFail    = 11
)

const (
	LoginType       = "LoginMessage"
	LoginRspType    = "LoginRspMesssage"
	RegisterType    = "RegisterMessage"
	RegisterRspType = "RegisterRspMessage"
)

type Message struct {
	Type string    `json:"type"`
	Content string `json:"content"`
}

type LoginMessage struct {
	Account string  `json:"account"`
	Password string `json:"password"`
}

type LoginRspMessage struct {
	Code int      `json:"code"`
	Status string `json:"status"`
}

type RegisterMessage struct {
	User User
}

type RegisterRspMessage struct {
	Code int      `json:"code"`
	Status string `json:"status"`
}