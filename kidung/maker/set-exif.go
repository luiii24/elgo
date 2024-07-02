package kidung

import (
      "strings"

	x "bot.lui/go/jangkep"	
      "github.com/amiruldev20/waSocket"
)

func init() {
	x.Rahwana(&x.Gace{
		Name:   "exif",
		Cmd:    []string{"exif"},
		Tags:   "maker",
		//IsMedia: true,
		IsOwner: true,
		//IsQuery:  true,
		//ValueQ: ".ai siap kamu?",
		Exec: func(sock *waSocket.Client, m *x.Raden, q string) {
			m.React("⏰")
      input := strings.Split(q, ",")
    if len(input) != 2 {
     m.Reply("coba lagi")
    return
    }
    x.CreateExif("lui.exif", input[0], input[1])
    m.React("✅")
    return
		},
	})
}







