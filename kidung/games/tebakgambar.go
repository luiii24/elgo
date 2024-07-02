package kidung

import (
    "encoding/json"
	"fmt"
	x "bot.lui/go/jangkep"
	"bot.lui/go/db"
    "io/ioutil"
    "math/rand"
	"net/http"
    "time"
	
    "github.com/amiruldev20/waSocket"
)

func init() {
	x.Rahwana(&x.Gace{
		Name:   "tebakgambar",
		Cmd:    []string{"tebakgambar"},
		Tags:   "games",
		//IsQuery:  true,
		//ValueQ: ".ai siap kamu?",
		Exec: func(sock *waSocket.Client, m *x.Raden, q string) {
			m.React("⏰")
    type A struct {
      Gambar string `json:"img"`
      Jwb    string `json:"jawaban"`
      Desc   string `json:"deskripsi"`
    }
    froms := m.Msg.Info.Chat.String()
    database, _ := db.ReadDB("db/db.json")
    if _, ok := database.Game[froms]; ok {
    m.Reply("Masih ada soal yg belum terjawab! ketik *nyerah* untuk menyerah.")
    m.React("❌")
    return
    }
    resp, err := http.Get("https://raw.githubusercontent.com/ramadhankukuh/database/master/src/games/tebakgambar.json")
    if err != nil {
    }
    defer resp.Body.Close()
    bite, _ := ioutil.ReadAll(resp.Body)
    var B []A
    if err := json.Unmarshal([]byte(bite), &B); err != nil {
    fmt.Println("url yg anda masukan salah!", err)
    return
    }
    rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(B))
    data := B[randomIndex]
    resp, err = http.Get(data.Gambar)
    if err != nil {
	// handle error
    }
    defer resp.Body.Close()
    bite, _ = ioutil.ReadAll(resp.Body)
    m.Reply("Reply stiker di bawah ini untuk menjawab!\n\nClue: "+ data.Desc + "\n\n*NB:*\n```TIMEOUT``` *30 DETIK* ```ketik``` *hint* ```untuk mendapatkan kunci jawabannya! dan ketik``` *nyerah* ```untuk menyerah dan berganti soal yg baru!```")
    pkr := m.CreateStickerIMG(bite)
    err = db.AddGM(froms, pkr.ID, data.Jwb)
    if err != nil {
    fmt.Println("eror add in db")
    return
    }
    time.Sleep(30 * time.Second)
    database, _ = db.ReadDB("db/db.json")
    if _, ok := database.Game[froms]; ok {
    m.Reply("Waktu habis!!\n\nJawabannya adalah *"+data.Jwb+"*")
    db.DelGM(froms)
    m.React("✅")
    return
    }
		},
	})
}







