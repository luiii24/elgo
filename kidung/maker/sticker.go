package kidung

import (
	x "bot.lui/go/jangkep"
	
    "github.com/amiruldev20/waSocket"
)

func init() {
	x.Rahwana(&x.Gace{
		Name:   "stiker",
		Cmd:    []string{"stiker", "stiker", "s", "stc"},
		Tags:   "maker",
		IsMedia: true,
		//IsQuery:  true,
		//ValueQ: ".ai siap kamu?",
		Exec: func(sock *waSocket.Client, m *x.Raden, q string) {
			m.React("⏰")
		woo, err := m.MediaOn("gambar, video")
		if err != nil {
		m.Reply("Reply/send gambar/video!")
		m.React("❌")
		return
		}
    m.CreateStickerIMG(woo)    		
    m.React("✅")
    return
		},
	})
}







