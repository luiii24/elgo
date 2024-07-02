package kidung

import (
    "context"
    "encoding/json"
    "strings"
	//"fmt"
	x "bot.lui/go/jangkep"
	//"net/url"
	
    "google.golang.org/protobuf/proto"
    waProto "github.com/amiruldev20/waSocket/binary/proto"
    "github.com/amiruldev20/waSocket"
)


type Jsonstruk struct {
			Title       string `json:"title"`
			Description string `json:"description"`
			ID          string `json:"id"`
}

func removeDuplicates(Nun []Jsonstruk) []Jsonstruk {
	uniqueMap := make(map[string]Jsonstruk)
	var uniqueNun []Jsonstruk

	for _, json := range Nun {
		if _, ok := uniqueMap[json.ID];!ok {
			uniqueMap[json.ID] = json
			uniqueNun = append(uniqueNun, json)
		}
	}

	return uniqueNun
}

func List(dom *[]Jsonstruk, sock *waSocket.Client, m *x.Raden) {
drm, _ := json.Marshal(dom)
sock.SendMessage(context.Background(), m.Msg.Info.Chat, &waProto.Message{
    ViewOnceMessage: &waProto.FutureProofMessage{
    Message: &waProto.Message{
      MessageContextInfo: &waProto.MessageContextInfo{
      DeviceListMetadata:    &waProto.DeviceListMetadata{},
      DeviceListMetadataVersion: proto.Int32(2),
      },
      InteractiveMessage: &waProto.InteractiveMessage{
      Header: &waProto.InteractiveMessage_Header{
        Title:        proto.String("Simple bot yang di buat oleh lui, Sedjak 1990."),
        Subtitle:       proto.String(""),
        HasMediaAttachment: proto.Bool(false),
      },
      Body: &waProto.InteractiveMessage_Body{
        Text: proto.String(""),
      },
      Footer: &waProto.InteractiveMessage_Footer{
        Text: proto.String("©Create by Lui, 2024."),
      },
      InteractiveMessage: &waProto.InteractiveMessage_NativeFlowMessage_{
        NativeFlowMessage: &waProto.InteractiveMessage_NativeFlowMessage{
        Buttons: []*waProto.InteractiveMessage_NativeFlowMessage_NativeFlowButton{
          {
          Name: proto.String("cta_url"),
          ButtonParamsJson: proto.String(`{
            "display_text": "error? Chat Me!",
            "url": "https://wa.me/6282146092695",
            "merchant_url":"https://wa.me/6282146092695"
          }`),
          },
          {
          Name: proto.String("single_select"),
          ButtonParamsJson: proto.String(`{
    "title": "Click Here!",
    "sections": [
        {
            "title": "berikut ini list menu bot saya.",
            "highlight_label": "Main di QQ24 Di jamin langsung WD!",
            "rows": `+string(drm)+`
        }
    ]
}`), 
          },                             
        },
        },
      },
      },
    },
    },
   })    
 return
}

func init() {
	x.Rahwana(&x.Gace{
		Name:   "menu",
		Cmd:    []string{"menu", "cmd", "listmenu", "help", "menuall"},
		Tags:   "",
		//IsQuery:  true,
		//ValueQ: ".ai siap kamu?",
		Exec: func(sock *waSocket.Client, m *x.Raden, q string) {
			m.React("⏱️")
        var Nun []Jsonstruk
	    Todo := x.Sinta()
	    Duos := strings.Split(m.Pesan(), ` `)
	    if len(Duos) == 2 {
	    var exam string
	     for _, cmd := range Todo {
	     if cmd.Value == "" {
	     exam = "."+cmd.Cmd[0]
	     } else {
	     exam = cmd.Value 
	     }
	     if cmd.Tags == Duos[1] {
	     ruko := &Jsonstruk {
	     Title: strings.ToUpper(cmd.Cmd[0]),
	     Description: "Contoh: "+exam,
	     ID: "."+ cmd.Cmd[0],
	     }
         Nun = append(Nun, *ruko)
	     }}
	     List(&Nun, sock, m)
			m.React("✅")
         return
	    } else {
	     for _, cmd := range Todo {
	     ruko := &Jsonstruk {
	     Title: strings.ToUpper(cmd.Tags)+" MENU",
	     Description: "",
	     ID: ".menu "+cmd.Tags,
	     }
         Nun = append(Nun, *ruko)
	     }}
	     Noel := removeDuplicates(Nun)
	     List(&Noel, sock, m)
			m.React("✅")
		},
	})
}

