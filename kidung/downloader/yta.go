package cmd

import (
    "encoding/json"
	"fmt"
    "regexp"
	x "bot.lui/go/jangkep"
	z "bot.lui/go/scrap"

    "github.com/amiruldev20/waSocket"
)

func init() {
	x.Rahwana(&x.Gace{
		Name:   "yta",
		Cmd:    []string{"yta", "ytmp3", "ytaudio", "youtubemp3"},
		Tags:   "downloader",
		IsQuery:  true,
		Value: ".yta https://youtu.be/ONiQqw0bsHo?si=VZUBKMVVeeeN-k69",
		Exec: func(sock *waSocket.Client, m *x.Raden, q string) {
			m.React("⏰")
        regex := regexp.MustCompile(`^(?:https?:\/\/)?(?:www\.|m\.|music\.)?youtu\.?be(?:\.com)?\/?.*(?:watch|embed)?(?:.*v=|v\/|\/)([\w\-_]+)\&?`)
        if !regex.MatchString(q) {
        m.Reply("itu bukan link youtube")
        m.React("❌")
        return
        }
    type Inp struct {
	Localized            struct {
		Description string `json:"description"`
		Title       string `json:"title"`
	} `json:"localized"`
	Thumbnails  struct {
		Maxres struct {
			Height int    `json:"height"`
			URL    string `json:"url"`
			Width  int    `json:"width"`
		} `json:"maxres"`
	Title string `json:"title"`
    }
    }    
     resp, err := z.Ytdl(q)
    if err != nil {
    m.Reply("terjadi kesalahan coba dengan url yang lain!")    
    return
    }
    urldl := resp["mp3"].(map[string]map[string]interface{})["128kbps"]["url"].(string)
     var yt Inp
    bite, _ := json.Marshal(resp["information"])     
    if err := json.Unmarshal([]byte(bite), &yt); err != nil {
        fmt.Println("url yg anda masukan salah!", err)
        return
    }
    m.ReplyAd("tunggu sebentar file media akan segera dikirim!\n\n"+ m.Lainnya()+ "\n"+ yt.Localized.Description, &x.SendModifyParams{Title: yt.Localized.Title, Body: yt.Localized.Description, SourceUrl: q, Thumbnail: yt.Thumbnails.Maxres.URL})
    m.SendDoc(m.Msg.Info.Chat, urldl, &x.SendDocParams{Capt: "", Mime: "audio/mpeg", FileName: yt.Localized.Title +".mp3", Thumb: yt.Thumbnails.Maxres.URL})        
			m.React("✅")
		},
	})
}
