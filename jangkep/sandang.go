package jangkep
import (
    "github.com/amiruldev20/waSocket"
    "github.com/amiruldev20/waSocket/types/events"
)

type Ex struct {
    Cmd  Gace
}

type Bd struct {
    Bs string
}

type Qt struct {
    Status        bool
    Type        string
}
type Raden struct {
    sock * waSocket.Client
    Msg * events.Message
}

type SendModifyParams struct {
    Text        string
    Title        string
    Body       string
    Thumbnail string
    SourceUrl  string
    Mentions  []string
}

type SendDocParams struct {
    Thumb string
    Capt string
    Mime string
    FileName string
}

type Gace struct {
	Name     string
	Cmd      []string
	Desc     string
	Tags     string
	Prefix   bool
	IsOwner  bool
	IsMedia  bool
	IsQuery  bool
	Value   string
	IsGroup  bool
	IsAdmin  bool
	IsBotAdmin bool
	Exec     func(client *waSocket.Client, m *Raden, q string)
}

