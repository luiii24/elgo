package db

import (
         "time"
        )


type DB struct {
  Uptime    time.Time           `json:"Uptime"`
	Self      bool                `json:"Self"`
	Owner     []string              `json:"Owner"`
	AutoLiat  bool                `json:"AutoLiat"`
	AutoKetik  bool                `json:"AutoKetik"`
	WithPrefix  bool                `json:"WithPrefix"`
  Ban       []string            `json:"Ban"` 
  Simi      []string            `json:"Simi"`
  Game      map[string]GameSet  `json:"Game"`
 	User      map[string]UserSet  `json:"User"`
	Group     map[string]GroupSet `json:"Group"`
  Menfes    map[string]MenfeSet `json:"Menfes"`
  Anonim      map[string]AnoSet   `json:"Anonim"`    
  Cmd         map[string]CdSet    `json:"Cmd"`       
}

type GameSet struct {
   Id    string  `json:"id"`
   Answer string `json:"answer"`  
}

type UserSet struct {
    Jid     string `json:"jid"`
    Name    string `json:"name"`
    Limit   int    `json:"limit"`
    Premium bool   `json:"premium"`
    Afk     AfkSet `json:"afk"`
} 

type AfkSet struct {
    Aksi bool `json:"aksi"`
    Alasan string `json:"alasan"`
    Time time.Time `json:"time"`
}

type GroupSet struct {
   Jid      string `json:"jid"`
   Name     string `json:"name"`
   Mute     bool   `json:"mute"`
   AntiLink bool   `json:"antilink"`
}

type MenfeSet struct {
  Partner string `json:"partner"`
  Pesan   string `json:"pesan"`
  Grade   string `json:"grade"`
  Status  string `json:"status"`
}

type AnoSet struct {
 Partner string `json:"partner"`
}

type CdSet struct {
 Query string `json:"query"`
 Cmd string `json:"cmd"`
}