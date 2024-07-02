package cmd

import (
	"fmt"
    "regexp"
    "strings"
    "net/http"
    "net/url"
    "io/ioutil"
    "os/exec"
	x "bot.lui/go/jangkep"

    "github.com/amiruldev20/waSocket"
    "github.com/antchfx/htmlquery"
)

func init() {
	x.Rahwana(&x.Gace{
		Name:   "tiktok",
		Cmd:    []string{"tiktok", "tiktokmp3", "tiktokmp4", "tiktokdl", "tiktokdownloader", "tt"},
		Tags:   "downloader",
		IsQuery:  true,
		Value: ".tiktok https://www.tiktok.com/@suzimartinezlocutora/video/7351427095035268358?is_from_webapp=1&sender_device=pc",
		Exec: func(sock *waSocket.Client, m *x.Raden, q string) {
			m.React("⏰")
        regex := regexp.MustCompile(`^(?:https?:\/\/)?(?:www\.|vt\.|vm\.|t\.)?(?:tiktok\.com\/)(?:\S+)?$`)
        if !regex.MatchString(q) {
        m.Reply("itu bukan link tiktok")
        m.React("❌")
        return
        }
    resp, err := http.Get("https://ssstik.io/")
    if err != nil {
	m.React("❌")
	return
    }
    defer resp.Body.Close()
    bite, _ := ioutil.ReadAll(resp.Body)
    strr := string(bite)
    regex = regexp.MustCompile(`s_tt\s*=\s*'(.*?)'`)
    regex2 := regexp.MustCompile(`<p class="maintext">([^<]*)</p>`)
    regex3 := regexp.MustCompile(`<h2 class="text-shadow--black">([^<]*)</h2>`)
    match := regex.FindStringSubmatch(strr)    
    token := string(match[1]) 
    out, err := exec.Command("bash", "-c", `curl 'https://ssstik.io/abc?url=dl'  -H 'HX-Request: true'  -H 'HX-Trigger: _gcaptcha_pt'  -H 'HX-Target: target'  -H 'HX-Current-URL: https://ssstik.io/en'  -H 'Content-Type: application/x-www-form-urlencoded; charset=UTF-8'  -H 'User-Agent: Mozilla/5.0 (Linux; Android 10; K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Mobile Safari/537.36'  -H 'Referer: https://ssstik.io/en'  --data-raw 'id=`+url.QueryEscape(q)+`&locale=en&tt=`+token+`'  --compressed`).Output()
    if err != nil {
    m.Reply(fmt.Sprintf("%v", err))
    return
    }
    mentah := string(out)
    doc, err := htmlquery.Parse(strings.NewReader(mentah))
    if err != nil {
        fmt.Println("Error parsing HTML:", err)
        return
    }
    links := htmlquery.Find(doc, "//a")
    title := regex2.FindStringSubmatch(mentah)
    if fmt.Sprintf("%v", title) == "[]" {
    title = regex3.FindStringSubmatch(mentah)
    }
    for _, link := range links {
        href := htmlquery.SelectAttr(link, "href")
        if href != "" && strings.Contains(fmt.Sprintf("%v", link), "without_watermark") { 
          m.SendVid(m.Msg.Info.Chat, href, title[1])
            fmt.Println("Extracted link:", href)
        } else if href != "" && strings.Contains(fmt.Sprintf("%v", link), "music") {
        m.SendDoc(m.Msg.Info.Chat, href, &x.SendDocParams{Capt: "", Mime: "audio/mpeg", FileName: title[1] +".mp3"})          
            fmt.Println("Extracted link:", href)
        }  else if href != "" && strings.Contains(fmt.Sprintf("%v", link), "slide") {
        m.SendImg(m.Msg.Info.Chat, href, title[1])        
        }  
        } 
			m.React("✅")
		},
	})
}
