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
	Type string
	Content string
}

type LoginMessage struct {
	Account string
	Password string
}

type LoginRspMesssage struct {
	Code int
	Status string
}