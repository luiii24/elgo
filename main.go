package main

import (
    "bufio"
    "context"
    "encoding/base64"
    //"encoding/json"
    "fmt"
    "log"
    "os"
    "os/signal"
    "syscall"
    "strings"
    "bot.lui/go/jangkep"
	 _ "bot.lui/go/kidung"
	_ "bot.lui/go/kidung/tes"
	_ "bot.lui/go/kidung/games"
	_ "bot.lui/go/kidung/downloader"	
	_ "bot.lui/go/kidung/maker"	

    _ "github.com/mattn/go-sqlite3"
    "github.com/mdp/qrterminal"
    //"github.com/probandula/figlet4go"
    //"github.com/joho/godotenv"
	"github.com/manifoldco/promptui"

    "github.com/amiruldev20/waSocket"
    //waProto "github.com/amiruldev20/waSocket/binary/proto"
    //"github.com/amiruldev20/waSocket/store"
    "github.com/amiruldev20/waSocket/store/sqlstore"
    "github.com/amiruldev20/waSocket/types/events"
    "github.com/amiruldev20/waSocket/types"
    waLog "github.com/amiruldev20/waSocket/util/log"

    //"google.golang.org/protobuf/proto"
)

 func InputPrompt(label string) string {
    var s string
    r := bufio.NewReader(os.Stdin)
    for {
        fmt.Fprint(os.Stderr, label+" ")
        s, _ = r.ReadString('\n')
        if s != "" {
            break
        }
    }
    return strings.TrimSpace(s)
} 

func main() {
    dbLog := waLog.Stdout("Database", "ERROR", true)
    dxz,
    err := base64.StdEncoding.DecodeString("TVlXQSBCT1Q=")
    if err != nil {
        panic("malformed input")
        log.Println(dxz)
    }
    container,
    err := sqlstore.New("sqlite3", "file:elgo.db?_foreign_keys=on", dbLog)
    if err != nil {
        panic(err)
    }

    deviceStore,
    err := container.GetFirstDevice()
    if err != nil {
        panic(err)
    }
    clientLog := waLog.Stdout("Client", "ERROR", true)

    /* setting env */
    /* client */
        sock := waSocket.NewClient(deviceStore, clientLog)
    eventHandler := registerHandler(sock)
    sock.AddEventHandler(eventHandler)
    if sock.Store.ID == nil {
	prompt := promptui.Select{
		Label: "Pilih tipe login",
		Items: []string{"KODE", "QR"},
	}

	_, typeLogin, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}    
        if typeLogin == "KODE" {
        numberBot := InputPrompt("your bot number (628xxx):")
            fmt.Println("You login with pairing code")
            err = sock.Connect()
            if err != nil {
                panic(err)
            }

            /* don't edit */
            code, err := sock.PairPhone(numberBot, true, waSocket.PairClientChrome, "Chrome (Linux)")

            if err != nil {
                fmt.Println(err)
                return
            }
            log.Println("Your Code: " + code)

        } else {
            qrChan,
            _ := sock.GetQRChannel(context.Background())

                err = sock.Connect()
            if err != nil {
                panic(err)
            }

            for evt := range qrChan {
                if evt.Event == "code" {
                    qrterminal.GenerateHalfBlock(evt.Code, qrterminal.L, os.Stdout)
                    log.Println("Please scan this QR...")
                } else {
                    log.Println("Login successfully!!")
                }
            }
        }
    } else {

        /* Already logged in, just connect */
        err = sock.Connect()
        log.Println("Berhasil terhubung!!")
        if err != nil {
            panic(err)
        }
    }

    c := make(chan os.Signal)
    signal.Notify(c, os.Interrupt, syscall.SIGTERM)
<-c

    sock.Disconnect()
}

func registerHandler(sock * waSocket.Client) func(evt interface {}) {
    return func(evt interface {}) {
        switch v := evt.(type) {
            case *events.Message:
                if strings.HasPrefix(v.Info.ID, "BAE5") {
                    return
                }
                if v.Info.Chat.String() == "status@broadcast" {
                    sock.MarkRead([] types.MessageID {v.Info.ID}, v.Info.Timestamp, v.Info.Chat, v.Info.Sender)
                    fmt.Println("Berhasil melihat status", v.Info.PushName)
                }
                m := jangkep.NewSimp(sock, v)
                pesan := m.Pesan()
                sender := v.Info.Sender.String()
                formattedString := fmt.Sprintf("%v", "@"+sender+"/~># "+pesan)
                 fmt.Println(formattedString)
                go jangkep.Buto(sock, m)
                break
            case *events.GroupInfo:
            //go message.Gc(sock, v)
            return     
                break
        }
    }
}
