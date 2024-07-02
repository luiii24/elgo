package kidung

import (
    //"encoding/json"
	//"fmt"
	x "bot.lui/go/jangkep"
	//"bot.lui/go/db"
    //"io/ioutil"
    //"math/rand"
	//"net/http"
	"regexp"
	//"strings"
    //"time"
	
    "github.com/amiruldev20/waSocket"
)

func init() {
	x.Rahwana(&x.Gace{
		Name:   "tes",
		Cmd:    []string{"tes", "tos", "tas"},
		Tags:   "tes",
		//IsQuery:  true,
		//ValueQ: ".ai siap kamu?",
		Exec: func(sock *waSocket.Client, m *x.Raden, q string) {
			m.React("⏰")
    regx := regexp.MustCompile(`^[\.//\/#\~\%\!\?]`)
    prefix := regx.FindString(q)
    m.Reply(prefix)	  
    m.React("✅")
    return
		},
	})
}







