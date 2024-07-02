package jangkep

import (
    "bytes"
    "context"
    "encoding/json"
    "encoding/hex"
    "errors"
    "fmt"
    "strconv"
    "strings"
    "regexp"
    "io/ioutil"
    "net/http"
    "log"
    "os"
    "os/exec"
    "math/rand"
    "image"
    "image/jpeg"
    "image/png"
    "golang.org/x/image/webp"
    "time"
    "bot.lui/go/db"

    "github.com/nfnt/Resize"
    "github.com/codedius/imagekit-go"
    "github.com/amiruldev20/waSocket"
    waProto "github.com/amiruldev20/waSocket/binary/proto"
    "github.com/amiruldev20/waSocket/types"
    "github.com/amiruldev20/waSocket/types/events"

    "google.golang.org/protobuf/proto"
)

var (
   PublicKey = "public_UQ6h5vB0OH3i2Z3kIxMV/tY1qk0="
   PrivateKey = "private_g+VylqvqKv/0UtK0EcznKUnJlTk="
)

func NewSimp(Cli * waSocket.Client, m * events.Message) * Raden {
    return &Raden {
        sock: Cli,
        Msg: m,
    }
}

/* parse jid */
func(m * Raden) ParseJID(arg string)(types.JID, bool) {
    if arg[0] == '+' {
        arg = arg[1: ]
    }
    if !strings.ContainsRune(arg, '@') {
        return types.NewJID(arg, types.DefaultUserServer), true
    } else {
        recipient,
        err := types.ParseJID(arg)
        if err != nil {
        fmt.Println("Invalid JID %s: %v", arg, err)
            return recipient, false
        } else if recipient.User == "" {
            fmt.Println("Invalid JID %s: no server specified", arg)
            return recipient, false
        }
        return recipient,
        true
    }
}

/* send react */
func(m * Raden) React(react string) {
    _,
    err := m.sock.SendMessage(context.Background(), m.Msg.Info.Chat, m.sock.BuildReaction(m.Msg.Info.Chat, m.Msg.Info.Sender, m.Msg.Info.ID, react))
    if err != nil {
        return
    }
}

/* send message */
func(m * Raden) SendMsg(jid types.JID, teks string) {
    _,
    err := m.sock.SendMessage(context.Background(), jid, & waProto.Message {Conversation: proto.String(teks)})
    if err != nil {
        return
    }
}

/* send reply */
func(m * Raden) Reply(teks string) waSocket.SendResponse {
    yes, _ := m.sock.SendMessage(context.Background(), m.Msg.Info.Chat, & waProto.Message {
        ExtendedTextMessage: & waProto.ExtendedTextMessage {
            Text: proto.String(teks),
            ContextInfo: & waProto.ContextInfo {
                Expiration: proto.Uint32(0),
                StanzaId: & m.Msg.Info.ID,
                Participant: proto.String(m.Msg.Info.Sender.String()),
                QuotedMessage: m.Msg.Message,
                MentionedJid: []string{m.Msg.Info.Sender.String()},
                ForwardingScore: proto.Uint32(99),
                IsForwarded: proto.Bool(true),
                BusinessMessageForwardInfo: &waProto.ContextInfo_BusinessMessageForwardInfo {
                BusinessOwnerJid: proto.String("6282146092695@s.whatsapp.net"),
                },
                ForwardedNewsletterMessageInfo: &waProto.ForwardedNewsletterMessageInfo {
                    NewsletterJid: proto.String("120363183703283232@newsletter"),
                    ServerMessageId: proto.Int32(100),
                    NewsletterName: proto.String("Whatsapp Bot Go. By lui, Inc."),                
                },                    
            },
        },
    })
    return yes
}

/* send adReply */
func(m * Raden) ReplyAd(teks string, option *SendModifyParams) {
    if option.Text == "" {
    option.Text = ""
    }
    if option.Title == "" {
    option.Title = "Whatsapp Bot Go."
    }    
    if option.Body == "" {
    option.Body = "di buat oleh lui, Inc."
    } 
    if option.Thumbnail == "" {
    option.Thumbnail = "https://wallpaperaccess.com/full/5750703.jpg"
    }
    if option.SourceUrl == "" {
    option.SourceUrl = "https://chat.whatsapp.com/8KLSiApeV9YLuZixNepwJX"
    }
    var isImage = waProto.ContextInfo_ExternalAdReplyInfo_IMAGE
    _, err := m.sock.SendMessage(context.Background(), m.Msg.Info.Chat, & waProto.Message {
        ExtendedTextMessage: & waProto.ExtendedTextMessage {
            Text: proto.String(teks),
            ContextInfo: & waProto.ContextInfo {
                ExternalAdReply: & waProto.ContextInfo_ExternalAdReplyInfo {
                    Title: proto.String(option.Title),
                    Body: proto.String(option.Body),
                    MediaType: & isImage,
                    Thumbnail: GetByte(option.Thumbnail),
                    ThumbnailUrl: proto.String(option.Thumbnail),
                    MediaUrl: proto.String("https://telegra.ph/?id="+ Random()),
                    SourceUrl: proto.String(option.SourceUrl),
                    ShowAdAttribution: proto.Bool(true),
                    RenderLargerThumbnail: proto.Bool(true),
                },
                Expiration: proto.Uint32(0),
                StanzaId: & m.Msg.Info.ID,
                Participant: proto.String(m.Msg.Info.Sender.String()),
                QuotedMessage: m.Msg.Message,
                MentionedJid: option.Mentions,
                ForwardingScore: proto.Uint32(199),
                IsForwarded: proto.Bool(true),
                BusinessMessageForwardInfo: &waProto.ContextInfo_BusinessMessageForwardInfo {
                    BusinessOwnerJid: proto.String("6282146092695@s.whatsapp.net"),
                },    
                ForwardedNewsletterMessageInfo: &waProto.ForwardedNewsletterMessageInfo {
                NewsletterJid: proto.String("120363183703283232@newsletter"),
                ServerMessageId: proto.Int32(100),
                NewsletterName: proto.String("Whatsapp Bot Go. By lui, Inc."),                
                },                
            },
        },
    })
    if err != nil {
        return
    }
}
/* send img */
func (m *Raden) SendImg(jid types.JID, imgurl string, capt string) {
    var uploadImg waSocket.UploadResponse
    image := GetByte(imgurl)
    thumb := GenThumb(image)
        uploadImg, error := m.sock.Upload(context.Background(), image, waSocket.MediaImage)
                if error != nil {
                        log.Println("Failed to upload file:", error)
                        return
                }
        _, err := m.sock.SendMessage(context.Background(), m.Msg.Info.Chat, &waProto.Message{
                ImageMessage: &waProto.ImageMessage{
                        Caption:           proto.String(capt),
                        Url:           proto.String(uploadImg.URL),
                        DirectPath:    proto.String(uploadImg.DirectPath),
                        MediaKey:      uploadImg.MediaKey,
                        Mimetype:      proto.String(http.DetectContentType(image)),
                        FileEncSha256: uploadImg.FileEncSHA256,
                        FileSha256:    uploadImg.FileSHA256,
                        FileLength:    proto.Uint64(uint64(len(image))),
                        JpegThumbnail: thumb,
                        ThumbnailDirectPath: proto.String(uploadImg.DirectPath),
                        ThumbnailSha256: uploadImg.FileSHA256,
                        ThumbnailEncSha256: uploadImg.FileEncSHA256,
                },
        })
        if err != nil {
                log.Println(err)
                return
        }
}

func (m *Raden) SendVid(jid types.JID, url string, capt string) {
    var uploadVid waSocket.UploadResponse
    video := GetByte(url)
    thumb := GenThumbv(video)
        uploadVid, error := m.sock.Upload(context.Background(), video, waSocket.MediaVideo)
                if error != nil {
                        log.Println("Failed to upload file:", error)
                        return
                }
        _, err := m.sock.SendMessage(context.Background(), m.Msg.Info.Chat, &waProto.Message{
                VideoMessage: &waProto.VideoMessage{
                        Caption:           proto.String(capt),
                        Url:           proto.String(uploadVid.URL),
                        DirectPath:    proto.String(uploadVid.DirectPath),
                        MediaKey:      uploadVid.MediaKey,
                        Mimetype:      proto.String(http.DetectContentType(video)),
                        FileEncSha256: uploadVid.FileEncSHA256,
                        FileSha256:    uploadVid.FileSHA256,
                        FileLength:    proto.Uint64(uint64(len(video))),
                        JpegThumbnail: thumb,
                        ThumbnailDirectPath: proto.String(uploadVid.DirectPath),
                        ThumbnailSha256: uploadVid.FileSHA256,
                        ThumbnailEncSha256: uploadVid.FileEncSHA256,
                        ContextInfo: & waProto.ContextInfo {
                        StanzaId: & m.Msg.Info.ID,
                        Participant: proto.String(m.Msg.Info.Sender.String()),
                        QuotedMessage: m.Msg.Message,
                   },
                },
        })
        if err != nil {
                log.Println(err)
                return
        }
}

func (m *Raden) Lainnya() string {
  return strings.Repeat(string(8206), 4001)
} 

func (m *Raden) IsQuoted() *Qt {
    extended := m.Msg.Message.GetExtendedTextMessage()
    quotedMsg := extended.GetContextInfo().GetQuotedMessage()
    if quotedMsg != nil && quotedMsg.GetConversation() != "" {
    return &Qt{
    Status: true,
    Type: "text",
    }    
    } else if quotedMsg != nil {
	do, _ := m.sock.DownloadAny(quotedMsg)
    return &Qt{
    Status: true,
    Type: http.DetectContentType(do),
    }
    } 
    return &Qt{
    Status: false,
    Type: "",
    }
}
func (m *Raden) IsGroup() bool {
    if strings.Split(m.Msg.Info.Chat.String(), "@")[1] != "g.us" {
    return false
    }
    return true
}

func (m *Raden) IsAdmin() bool {
    if m.GetGroupAdmin(m.Msg.Info.Chat, m.Msg.Info.Sender.String()) == false {
    return false
    }
    return true
}

func (m *Raden) IsBotAdmin() bool {
    if m.GetGroupAdmin(m.Msg.Info.Chat, "6282324180431@s.whatsapp.net") == false {
    return false
    }       
    return true
}

func (m *Raden) IsOwner() bool {
database, _ := db.ReadDB("db/db.json")
if strings.Contains(strings.Join(database.Owner, " "), strings.Split(m.Msg.Info.Sender.String(), "@")[0]) {
return true
}
return false
}

func (m *Raden) Games(user string) bool {
    database, _ := db.ReadDB("db/db.json")
    datanya := database.Game[user]
    q := m.IsQuoted()
    if q.Status {
    idnya := m.Msg.Message.ExtendedTextMessage.ContextInfo.GetStanzaId()
    jwbnya := m.Pesan()
    if idnya == datanya.Id {
    if strings.ToLower(jwbnya) == strings.ToLower(datanya.Answer) {
    m.Reply("BENAR!!")
    db.DelGM(user)
    return true
    } else {
    m.Reply("SALAH! COBA LAGI!!!")
    }
    }
    }
    if strings.Contains(strings.ToLower(strings.Split(m.Pesan(), " ")[0]), "hint") {
    raegex := regexp.MustCompile(`[bcdfghjklmnpqrstvwxyz]`)
      oker := strings.ToLower(datanya.Answer)
        m.Reply("```"+raegex.ReplaceAllString(oker, "_")+"```")
    return true
    } else if strings.Contains(strings.ToLower(strings.Split(m.Pesan(), " ")[0]), "nyerah") {
        m.Reply("YAH NYERAH")
        db.DelGM(user)
        return true
        }
    return false
}

func (m *Raden) SendVn(jid types.JID, url string) {
    var uploadVn waSocket.UploadResponse
    audio := GetByte(url)
	RawPath := fmt.Sprintf("tmp/%s%s", m.Msg.Info.ID, "." + strings.Split(http.DetectContentType(audio), "/")[1])
	ConvertedPath := fmt.Sprintf("tmp/%s%s", m.Msg.Info.ID, ".ogg")
	err := os.WriteFile(RawPath, audio, 0600)
	if err != nil {
		fmt.Printf("Failed to save image: %v", err)
	}
	RawPth := m.Msg.Info.ID + "." + strings.Split(http.DetectContentType(audio), "/")[1]
	ConvertedPth := m.Msg.Info.ID + ".ogg"
	exc := exec.Command("bash", "-c", "cd tmp && ffmpeg -i "+RawPth+" -acodec libopus -b:a 128000 "+ConvertedPth)
  	_ = exc.Run()
    aud, err := os.ReadFile(ConvertedPath)
	if err != nil {
		fmt.Printf("Failed to read %s: %s\n", ConvertedPath, err)
	}
	out, err := exec.Command("bash", "-c", "cd tmp && ffprobe "+ConvertedPth+" 2>&1 | grep -A1 Duration:").Output()
    if err != nil {
    fmt.Sprintf("%v", err)
    return
    }
    scnd, _ := strconv.ParseUint(strings.Split(strings.Split(string(out), ":")[3], ".")[0], 10, 32)
        uploadVn, error := m.sock.Upload(context.Background(), aud, waSocket.MediaAudio)
                if error != nil {
                        log.Println("Failed to upload file:", error)
                        return
                }
        _, err = m.sock.SendMessage(context.Background(), m.Msg.Info.Chat, &waProto.Message{
                AudioMessage: &waProto.AudioMessage{
                        Url:           proto.String(uploadVn.URL),
                        DirectPath:    proto.String(uploadVn.DirectPath),
                        MediaKey:      uploadVn.MediaKey,
                        Mimetype:      proto.String("audio/ogg; codecs=opus"),
                        FileEncSha256: uploadVn.FileEncSHA256,
                        FileSha256:    uploadVn.FileSHA256,
                        FileLength:    proto.Uint64(uint64(len(audio))),
                        Ptt:           proto.Bool(true),     
                        Seconds:       proto.Uint32(uint32(scnd)),	
                        //Waveform: proto.String("AAAAAAAAAAAAIB4LCAwYFhQZIS46NCUeGBQSEQoWGhQRDAgGCwILDwoRCwcGCAcGCAAABAsNBQIFAAIFAAAAAA=="),
                },
        })
        if err != nil {
                log.Println(err)
                return
        }
        _ = os.Remove(RawPath)  
        _ = os.Remove(ConvertedPath)         
}

/* send contact */
func(m * Raden) SendContact(jid types.JID, number string, nama string) {
    _, err := m.sock.SendMessage(context.Background(), jid, & waProto.Message {
        ContactMessage: & waProto.ContactMessage {
            DisplayName: proto.String(nama),
            Vcard: proto.String(fmt.Sprintf("BEGIN:VCARD\nVERSION:3.0\nN:%s;;;\nFN:%s\nitem1.TEL;waid=%s:+%s\nitem1.X-ABLabel:Mobile\nX-WA-BIZ-DESCRIPTION:Create By lui. Inc.\nX-WA-BIZ-NAME:Lui, Inc.\nEND:VCARD", nama, nama, number, number)),
            ContextInfo: & waProto.ContextInfo {
                StanzaId: & m.Msg.Info.ID,
                Participant: proto.String(m.Msg.Info.Sender.String()),
                QuotedMessage: m.Msg.Message,
            },
        },
    })
    if err != nil {
        return
    }
}
/*send document*/
func(m * Raden) SendDoc(jid types.JID, fileurl string, option *SendDocParams) {
    var uploaddoc waSocket.UploadResponse
        resp, _ := http.Get(fileurl)
    defer resp.Body.Close()
    fileByte, _ := ioutil.ReadAll(resp.Body)
        uploaddoc, error := m.sock.Upload(context.Background(), fileByte, waSocket.MediaDocument)
                if error != nil {
                        log.Println("Failed to upload file:", error)
                        return
                }
                if option.Capt == "" {
                option.Capt = "Create by Lui, Inc."
                }
                if option.Mime == "" {
                option.Mime = http.DetectContentType(fileByte)
                }
                if option.FileName == "" {
                option.FileName = "Lui, Inc.unknow"
                }                
                if option.Thumb == "" {
                option.Thumb = "https://www.nicepng.com/png/detail/264-2641184_111-kb-png-golang-logo.png"
                }
          image := GetByte(option.Thumb)
          thumb := GenThumb(image)
        _, err := m.sock.SendMessage(context.Background(), m.Msg.Info.Chat, & waProto.Message{
                DocumentMessage: & waProto.DocumentMessage{
                        Caption:           proto.String(option.Capt),
                        Title:           proto.String(option.FileName),
                        Url:           proto.String(uploaddoc.URL),
                        FileName:           proto.String(option.FileName),
                        DirectPath:    proto.String(uploaddoc.DirectPath),
                        MediaKey:      uploaddoc.MediaKey,
                        Mimetype:      proto.String(option.Mime),
                        FileEncSha256: uploaddoc.FileEncSHA256,
                        FileSha256:    uploaddoc.FileSHA256,
                        FileLength:    proto.Uint64(uint64(len(fileByte))),
                        JpegThumbnail: thumb,
                        ContextInfo: & waProto.ContextInfo {
                        StanzaId: & m.Msg.Info.ID,
                        Participant: proto.String(m.Msg.Info.Sender.String()),
                        QuotedMessage: m.Msg.Message,
                   },

                },
        })
        if err != nil {
                log.Println(err)
                return
        }
}

func (m * Raden) CreateStickerIMG(data []byte) waSocket.SendResponse {
	RawPath := fmt.Sprintf("tmp/%s%s", m.Msg.Info.ID, "." + strings.Split(http.DetectContentType(data), "/")[1])
	ConvertedPath := fmt.Sprintf("tmp/%s%s", m.Msg.Info.ID, ".webp")
	err := os.WriteFile(RawPath, data, 0600)
	if err != nil {
		fmt.Printf("Failed to save image: %v", err)
	}
	RawPth := m.Msg.Info.ID + "." + strings.Split(http.DetectContentType(data), "/")[1]
	ConvertedPth := m.Msg.Info.ID + ".webp"
	exc := exec.Command("bash", "-c", "cd tmp && ffmpeg -i "+ RawPth +" -vcodec libwebp -filter:v fps=fps=20 -lossless 1 -loop 0 -preset default -an -vsync 0 -s 512:512 "+ ConvertedPth)
  	_ = exc.Run()
    exc = exec.Command("bash", "-c", "cd tmp && webpmux -set exif exif/lui.exif "+ ConvertedPth +" -o " + ConvertedPth) 
    _ = exc.Run()
    stc, err := os.ReadFile(ConvertedPath)
	if err != nil {
		fmt.Printf("Failed to read %s: %s\n", ConvertedPath, err)
	}
	uploaded, err := m.sock.Upload(context.Background(), stc, waSocket.MediaImage)
	if err != nil {
		fmt.Printf("Failed to upload file: %v\n", err)
	}
    yes, eror := m.sock.SendMessage(context.Background(), m.Msg.Info.Chat, & waProto.Message{	
		StickerMessage: &waProto.StickerMessage{
			Url:           proto.String(uploaded.URL),
			DirectPath:    proto.String(uploaded.DirectPath),
			MediaKey:      uploaded.MediaKey,
			Mimetype:      proto.String(http.DetectContentType(stc)),
			FileEncSha256: uploaded.FileEncSHA256,
			FileSha256:    uploaded.FileSHA256,
			FileLength:    proto.Uint64(uint64(len(data))),
			IsAvatar:    proto.Bool(true),
			IsAnimated: proto.Bool(true),
			ContextInfo: &waProto.ContextInfo{
				StanzaId:      &m.Msg.Info.ID,
				Participant:   proto.String(m.Msg.Info.Sender.String()),
				QuotedMessage: m.Msg.Message,
			},
		},
        })
        if eror != nil {
                log.Println(err)
        }
    _ = os.Remove(RawPath)  
    _ = os.Remove(ConvertedPath)
    return yes         
	}

/* create channel */
func(m * Raden) createChannel(params[] string) {
    _,
    err := m.sock.CreateNewsletter(waSocket.CreateNewsletterParams {
        Name: strings.Join(params, " "),
    })
    if err != nil {
        return
    }
}

func(m *Raden) MediaOn(sd string) ([]byte, error) {
    duos := strings.Split(sd, `, `)
    for _, d := range duos {
    woila := m.IsQuoted()
    dor, _ := m.sock.DownloadAny(m.Msg.Message)
    if woila.Status {
    extended := m.Msg.Message.GetExtendedTextMessage()
    quotedMsg := extended.GetContextInfo().GetQuotedMessage()
    if d == "gambar" && strings.Split(woila.Type, "/")[0] == "image" {
	dor, _ = m.sock.DownloadAny(quotedMsg)
	return dor, nil
    } else if d == "video" && strings.Split(woila.Type, "/")[0] == "video" {
	dor, _ = m.sock.DownloadAny(quotedMsg)
	return dor, nil    
    } else if d == "audio" && strings.Split(woila.Type, "/")[0] == "audio" {
	dor, _ = m.sock.DownloadAny(quotedMsg)
	return dor, nil    
    }} else if d == "gambar" && strings.Split(http.DetectContentType(dor), "/")[0] == "image" {
	dor, _ = m.sock.DownloadAny(m.Msg.Message)
	return dor, nil    
    } else if d == "video" && strings.Split(http.DetectContentType(dor), "/")[0] == "video" {
	dor, _ = m.sock.DownloadAny(m.Msg.Message)
	return dor, nil        
    } else if d == "audio" && strings.Split(http.DetectContentType(dor), "/")[0] == "audio" {
	dor, _ = m.sock.DownloadAny(m.Msg.Message)
	return dor, nil        
    }
    }
    return []byte{}, errors.New("Reply/send Media plzz!")    
}

/* fetch group admin */
func(m * Raden) FetchGroupAdmin(Jid types.JID)([] string, error) {
    var Admin[] string
    resp, err := m.sock.GetGroupInfo(Jid)
    if err != nil {
        return Admin, err
    } else {
        for _, group := range resp.Participants {
            if group.IsAdmin || group.IsSuperAdmin {
                Admin = append(Admin, group.JID.String())
            }
        }
    }
    return Admin, nil
}

/* get group admin */
func(m * Raden) GetGroupAdmin(jid types.JID, sender string) bool {
        if !m.Msg.Info.IsGroup {
            return false
        }
        admin, err := m.FetchGroupAdmin(jid)
        if err != nil {
            return false
        }
        for _, v := range admin {
            if v == sender {
                return true
            }
        }
        return false
    }

/* get link group */
func(m * Raden) LinkGc(Jid types.JID, reset bool) string {
    link,
    err := m.sock.GetGroupInviteLink(Jid, reset)

    if err != nil {
        panic(err)
    }
    return link
}


func(m * Raden) Query() string {
    return strings.Join(strings.Split(m.Pesan(), ` `)[1: ], ``)
}


func(m * Raden) Pesan() string {
    var yh map[string]interface{}
    ok, _ := json.Marshal(m.Msg.Message)
    err := json.Unmarshal(ok, &yh)
	if err != nil {
       return fmt.Sprintf("%v", err)
	}  
	if yh["conversation"] != nil {
	return yh["conversation"].(string)
	}
	for k := range yh {
	if woi := yh[k].(map[string]interface{})["text"]; woi != nil {
	return yh[k].(map[string]interface{})["text"].(string)
	} else if woi := yh[k].(map[string]interface{})["caption"]; woi != nil {
	return yh[k].(map[string]interface{})["caption"].(string)
	} else if woi := yh[k].(map[string]interface{})["name"]; woi != nil {
	return yh[k].(map[string]interface{})["name"].(string)
	} else if woi := yh[k].(map[string]interface{})["InteractiveResponseMessage"]; woi != nil {
	return strings.Split(yh[k].(map[string]interface{})["InteractiveResponseMessage"].(map[string]interface{})["NativeFlowResponseMessage"].(map[string]interface{})["paramsJson"].(string), `"`)[3]
	}
	}	  
	return ""
}


func Upload(x []byte, n string) (*imagekit.UploadResponse, error) {
     a := imagekit.Options{
      PublicKey:  PublicKey,
      PrivateKey: PrivateKey,
     }   

     b, _ := imagekit.NewClient(&a)
     c := imagekit.UploadRequest{
	File:              x, // []byte OR *url.URL OR url.URL OR base64 string
	FileName:          n,
	UseUniqueFileName: false,
	Tags:              []string{"luibot", "lui"},
	Folder:            "/",
	IsPrivateFile:     false,
	CustomCoordinates: "",
	ResponseFields:    nil,
	}
	   
	ctx := context.Background()
	   
	d, _ := b.Upload.ServerUpload(ctx, &c)
       return d, nil
	}

func GetByte(source string) []byte {
	resp, _ := http.Get(source)
	defer resp.Body.Close()
	bytes, _ := ioutil.ReadAll(resp.Body)
       return bytes
}

func GenThumbv(video []byte) []byte {
      vidPath := "tmp/"+ Random() + ".mp4"
      thumbPath := Random() + ".jpg"
       _ = ByteToFile(video, vidPath) 
      cmd := exec.Command("ffmpeg", "-i", vidPath, "-ss", "00:00:01", "-vframes", "1", "-f", "image2", thumbPath)
      err := cmd.Run()
      if err  != nil { 
         panic("err") 
     }
      file, _ := os.Open(thumbPath)
    defer file.Close()
    stat, err  := file.Stat() 
    if err  != nil { 
        panic("err") 
    } 
       bs := make([]byte, stat.Size()) 
       _, err = file.Read(bs)
      if err != nil {
         panic("err")
      } 
       result := GenThumb(bs)
      os.Remove(vidPath)
      os.Remove(thumbPath)
      return result
}
      

func GenThumb(img []byte) []byte {
 var percentage int
  oker := GetMime(img)
       img2, _, _ := image.Decode(bytes.NewReader(img))
 if oker == "image/png" { 
	img2, _ = png.Decode(bytes.NewReader(img))
 } else if oker == "image/webp" {
	img2, _ = webp.Decode(bytes.NewReader(img))
 }
       w := img2.Bounds().Dx()
       h := img2.Bounds().Dy()
    if w > h {
       percentage = 150 * 100 / w
    } else {
       percentage = 150 * 100 / h
    }
 w = (w * percentage) / 100
 h = (h * percentage) / 100
       newimg := resize.Resize(uint(w), uint(h), img2, resize.Lanczos3) 
       buf := new(bytes.Buffer)
       _ = jpeg.Encode(buf, newimg, nil)     
        Converted := buf.Bytes()
return Converted
}

func Random() string {
   a := rand.Int63()
return fmt.Sprintf("%v", a)
}

func ByteToFile(data []byte, filename string) error {
       file, err := os.Create(filename)
    if err != nil {
       return err
    }
   defer file.Close()

   _, err = file.Write(data)
   return err
}

func CreateExif(fileName string, packname string, author string) *string {

	jsonData := map[string]interface{}{
		"sticker-pack-id":        "lui.bot",
		"sticker-pack-name":      packname,
		"sticker-pack-publisher": author,
		"android-app-store-link": "https://play.google.com/store/apps/details?id=com.dts.freefireth",
		"ios-app-store-link":     "https://apps.apple.com/app/id1300146617",
		"emojis": []string{
			"👋"},
	}

	jsonBytes,
		err := json.Marshal(jsonData)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	littleEndian := []byte{
		0x49,
		0x49,
		0x2a,
		0x00,
		0x08,
		0x00,
		0x00,
		0x00,
		0x01,
		0x00,
		0x41,
		0x57,
		0x07,
		0x00,
	}

	bytes := []byte{
		0x00,
		0x00,
		0x16,
		0x00,
		0x00,
		0x00}

	len := len(jsonBytes)
	var last string

	if len > 256 {
		len = len - 256
		bytes = append([]byte{
			0x01}, bytes...)
	} else {
		bytes = append([]byte{
			0x00}, bytes...)
	}

	if len < 16 {
		last = fmt.Sprintf("0%x", len)
	} else {
		last = fmt.Sprintf("%x", len)
	}

	buf2,
		err := hex.DecodeString(last)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	buf3 := bytes
	buf4 := jsonBytes

	buffer := append(littleEndian, buf2...)
	buffer = append(buffer, buf3...)
	buffer = append(buffer, buf4...)

	err = os.WriteFile("tmp/exif/"+fileName, buffer, 0644)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	return &fileName
}

func GetMime(file []byte) string {
	hsl := http.DetectContentType(file)
  return hsl
}

func Slamet(inpt time.Duration) string {
jam := int(inpt.Hours())
menit := int(inpt.Minutes()) % 60 
detik := int(inpt.Seconds()) % 60
hari := jam / 24
if menit == 00 {
 return strconv.Itoa(detik) + " detik."
} else if jam == 00 {
 return strconv.Itoa(menit) + " menit, " + strconv.Itoa(detik) + " detik."
} else if hari == 0 {
 return strconv.Itoa(jam) + " jam, " + strconv.Itoa(menit) + " menit, " + strconv.Itoa(detik) + " detik."
} else {
  return strconv.Itoa(hari) + " hari, " + strconv.Itoa(jam % 24) + " jam, " + strconv.Itoa(menit) + " menit, " + strconv.Itoa(detik) + " detik."
}
return ""
}

func Solikin(jid string, list []string) bool {
	for _, a := range list {
	 if strings.Contains(a, jid) {
	  return true
	 }
	}
	return false
	}

	func Yani(s []string, r string) []string {
		for i, v := range s {
		    if v == r {
			 return append(s[:i], s[i+1:]...)
		    }
		}
		return s
	   }
	   
func RandomStr(length int) string {
  rand.Seed(time.Now().UnixNano())
  randomFloat := rand.Float64()
  randomString := fmt.Sprintf("%x", int64(randomFloat * 1e16)) // 1e16 for better precision
  result := randomString[2:]
  return result
}

