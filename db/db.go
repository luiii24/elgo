package db

import (
     "encoding/json"
     "os"
     "strings"
     "time"
     "io/ioutil"
     //"bot.lui/go/jangkep"
)

func ReadDB(path string) (DB, error) {
	var db DB
	file, err := os.Open(path)
	if err != nil {
	    return db, err
	}
	defer file.Close()
   
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&db)
	if err != nil {
	    return db, err
	}
   
	return db, nil
}

func UpdateDB(input DB, filePath string) error {
	jsonData, err := json.MarshalIndent(input, "", " ")
	if err != nil {
	    return err
	}
   
	err = ioutil.WriteFile(filePath, jsonData, 0644)
	if err != nil {
	    return err
	}
   
	return nil
}

func Self(self bool) error {
	db, err := ReadDB("db/db.json")
	if err != nil {
	     return err
	}
	 db.Self = self
	 err =  UpdateDB(db, "db/db.json")
	if err != nil {
	     return err
	}
  return nil
}

func AddUsr(jid string, name string) error {
    db, err := ReadDB("db/db.json")
  if err != nil {
	return err
  }
   db.User[jid] = UserSet{Jid: jid, Name: name, Limit: 10, Premium: false, Afk: AfkSet{Aksi: false, Alasan: "", Time: time.Now()},}
   err = UpdateDB(db, "db/db.json")
  if err != nil {
	return err
  }
  return nil
}

func CekUsr(jid string) bool {
	db, err := ReadDB("db/db.json")
	if err != nil {
	     return false
	}
	if _, ok := db.User[jid]; ok {
         return true
	} else {
         return false
	}
   return false   
}

func UpdateUsr(jid string, opt *UserSet) error {
	db, err := ReadDB("db/db.json")
	dbo := db.User[jid]
	if err != nil {
	      return err
      }
      if opt.Name == "" {
      opt.Name = dbo.Name
      }
	db.User[jid] = UserSet{Name: opt.Name, Limit: opt.Limit, Premium: opt.Premium, Afk: AfkSet{Aksi: opt.Afk.Aksi, Alasan: opt.Afk.Alasan, Time: opt.Afk.Time,},}
	err = UpdateDB(db, "db/db.json")
	 if err != nil {
	      return err
	 }
	 return nil
      }

func AddGc(jid string, name string) error {
	db, err := ReadDB("db/db.json")
     if err != nil {
	   return err
     }
      db.Group[jid] = GroupSet{Jid: jid, Name: name, Mute: false, AntiLink: false,}
      err = UpdateDB(db, "db/db.json")
     if err != nil {
	   return err
     }
     return nil
   }
   
func CekGc(jid string) bool {
	   db, err := ReadDB("db/db.json")
	   if err != nil {
		 return false
	   }
	   if _, ok := db.Group[jid]; ok {
	     return true
	   } else {
	     return false
	   }
      return false   
   }
   
func UpdateGc(jid string, opt *GroupSet) error {
    db, err := ReadDB("db/db.json")
    dbo := db.Group[jid]
    if err != nil {
	   return err
   }
   if opt.Name == "" {
   opt.Name = dbo.Name
   }
    db.Group[jid] = GroupSet{Name: opt.Name, Mute: opt.Mute, AntiLink: opt.AntiLink,}
    err = UpdateDB(db, "db/db.json")
     if err != nil {
	   return err
     }
     return nil
   }

   func AddGM(jid string, id string, jwb string) error {
	 db, err := ReadDB("db/db.json")
     if err != nil {
	   return err
     }
      db.Game[jid] = GameSet{Id: id, Answer: jwb,}
      err = UpdateDB(db, "db/db.json")
     if err != nil {
	   return err
     }
     return nil
   }
   
   func DelGM(jid string) error {
   db, err := ReadDB("db/db.json")
     if err != nil {
	   return err
     }
   delete(db.Game, jid)
   err = UpdateDB(db, "db/db.json")
   if err != nil {
	 return err
   }
   return nil
   }
   
   func AddMf(jid string, partner string, psn string, gd string) error {
    db, err := ReadDB("db/db.json")
       if err != nil {
       return err
       }
        db.Menfes[jid] = MenfeSet{Partner: partner, Pesan: psn, Grade: gd, Status: ""}
        err = UpdateDB(db, "db/db.json")
       if err != nil {
       return err
       }
       return nil
     }
     
     func DelMf(jid string) error {
     db, err := ReadDB("db/db.json")
       if err != nil {
       return err
       }
     delete(db.Menfes, jid)
     err = UpdateDB(db, "db/db.json")
     if err != nil {
     return err
     }
     return nil
     }
  
   func AddSm(jid string) error {
    db, err := ReadDB("db/db.json")
       if err != nil {
       return err
       }
    db.Simi = append(db.Simi, jid)
    err = UpdateDB(db, "db/db.json")
    if err != nil {
    return err
    }
    return nil
    }
  
  func CekSm(jid string) bool {
   db, err := ReadDB("db/db.json")
   if err != nil {
    return false
   }
   for _, a := range db.Simi {
    if strings.Contains(a, jid) {
     return true
    }
   }
   return false
   }
    
   func DelSm(jid string) error {
    db, err := ReadDB("db/db.json")
    if err != nil {
    return err
    }
    db.Simi = Yani(db.Simi, jid)
    err = UpdateDB(db, "db/db.json")
    if err != nil {
    return err
    }
    return nil
   }

   func AddAnon(jid string, partner string) error {
    db, err := ReadDB("db/db.json")
       if err != nil {
       return err
       }
        db.Anonim[jid] = AnoSet{Partner: partner}
        err = UpdateDB(db, "db/db.json")
       if err != nil {
       return err
       }
       return nil
     }
     
     func DelAnon(jid string) error {
     db, err := ReadDB("db/db.json")
       if err != nil {
       return err
       }
     delete(db.Anonim, jid)
     err = UpdateDB(db, "db/db.json")
     if err != nil {
     return err
     }
     return nil
     }

	func Yani(s []string, r string) []string {
		for i, v := range s {
		    if v == r {
			 return append(s[:i], s[i+1:]...)
		    }
		}
		return s
	   }

     
//func DefaultDB() eror {
//}