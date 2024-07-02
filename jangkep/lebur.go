package jangkep

import (
	"fmt"
	"regexp"
	"strings"
	"time"
	"bot.lui/go/db"
	
    "github.com/amiruldev20/waSocket"
    "github.com/amiruldev20/waSocket/types"	
)

var Semar []Gace
var ExCrot map[string]Ex = make(map[string]Ex)
var ExM map[string]Bd = make(map[string]Bd)

func Rahwana(cmd *Gace) {
	Semar = append(Semar, *cmd)
}

func AppendCrt(Snd string, Gle * Gace) {
    ExCrot[Snd] = Ex{Cmd: *Gle}
}

func Diley(Snd string) {
    ExM[Snd] = Bd{Bs: "lui dongker."}
}

func Sinta() []Gace {
	return Semar
}

func Buto(c *waSocket.Client, m *Raden) {    
    var command string
    database, _ := db.ReadDB("db/db.json")
    pushName := m.Msg.Info.PushName
    from := m.Msg.Info.Chat
    froms := m.Msg.Info.Chat.String()    
    if db.CekUsr(froms) == false && strings.Split(froms, "@")[1] == "s.whatsapp.net" { err := db.AddUsr(froms, pushName); if err != nil {fmt.Println("eror menambahkan user ke database")}; fmt.Println("sukses menambahkan user!") }
    if db.CekGc(froms) == false && strings.Split(froms, "@")[1] == "g.us" { infogc, _ := c.GetGroupInfo(from); db.AddGc(froms, infogc.Name) }
    if _, ok := database.Game[froms]; ok { m.Games(froms) }
    if !m.IsOwner() && database.Self { return }
    regx := regexp.MustCompile(`^[\.//\/#\~\%\!\?]`)
    args := strings.Split(m.Pesan(), " ")
    query := strings.Join(args[1: ], ` `)
    cmd := strings.ToLower(args[0])
    prefix := regx.FindString(cmd)        
    if prefix == "" {
    if _, ok := ExCrot[m.Msg.Info.Sender.String()]; ok {
    Duos := ExCrot[m.Msg.Info.Sender.String()]
    delete(ExCrot, m.Msg.Info.Sender.String())    
    Duos.Cmd.Exec(c, m, m.Pesan())
    }
    return
    }
    if _, ok := ExM[m.Msg.Info.Sender.String()]; ok {m.Reply("anda telah mencapai batas per command tunggu 4 detik dan ekskusi kembali. Atau berlangganan premium untuk akses yg lebih.");delete(ExM, m.Msg.Info.Sender.String());return}    
    command = regx.ReplaceAllString(cmd, "")
    if cmd == prefix && len(args) > 1 {
    command = regx.ReplaceAllString(strings.ToLower(args[1]), "")
    query = strings.Join(args[2: ], ` `)
    }
	for _, cmd := range Semar {
	ada := false
    for i := 0; i < len(cmd.Cmd); i++ {
        if cmd.Cmd[i] == command {
            ada = true
            break
        }
    }
	if ada {
	Diley(m.Msg.Info.Sender.String())
			//Checking
            if database.AutoLiat {
            c.MarkRead([] types.MessageID {m.Msg.Info.ID}, m.Msg.Info.Timestamp, m.Msg.Info.Chat, m.Msg.Info.Sender)
            }

            if database.AutoKetik {
            c.SendChatPresence(from, types.ChatPresenceComposing, types.ChatPresenceMediaText)
            }

			if cmd.IsOwner && m.IsOwner() {
				m.Reply("Fitur ini hanya untuk owner!!")
				continue
			}

			/** if cmd.IsMedia && m.Media == nil {
				m.Reply("Silahkan reply / input media!!")
				continue
			} **/

			if cmd.IsQuery && query == "" {
				m.Reply("kirim teks atau link sebagai input!")
				AppendCrt(m.Msg.Info.Sender.String(), &cmd)
				continue
			}
			if cmd.IsGroup && !m.IsGroup() {
				m.Reply("Fitur ini hanya dapat digunakan didalam grup!!")
				continue
			}

			if (m.IsGroup() && cmd.IsAdmin) && !m.IsAdmin() {
				m.Reply("Fitur ini hanya untuk admin grup!!")
				continue
			}

			if m.IsBotAdmin() && cmd.IsBotAdmin {
				m.Reply("Untuk menggunakan fitur ini, bot harus menjadi admin!!")
				continue
			}

			cmd.Exec(c, m, query)
         time.Sleep(4 * time.Second)
         delete(ExM, m.Msg.Info.Sender.String())
		}
	}
}
