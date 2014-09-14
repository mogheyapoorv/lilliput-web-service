package webservice

import (
	"code.google.com/p/gorest"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"net/url"
)

type Servers struct {
	Id         int    `xorm:"not null pk autoincr INT(10)"`
	MacId      string `xorm:"not null unique VARCHAR(255)"`
	RegisterId string `xorm:"not null unique VARCHAR(2)"`
}

type RegisterService struct {
	gorest.RestService `root:"/" consumes:"application/json" produce:"application/json"`

	registerMachine gorest.EndPoint `method:"GET" path:"/register/{macId:string}"  output:"Servers"`
}

var engine *xorm.Engine
var wConn string

var Id map[int]string

func init() {
	wConn = ConnString("db.master")
	db := Orm()
	db.CreateTables(&Servers{})
	Id = make(map[int]string)
	Id[1] = "a"
	Id[2] = "b"
	Id[3] = "c"
	Id[4] = "d"
	Id[5] = "e"
	Id[6] = "f"
	Id[7] = "g"
	Id[8] = "h"
	Id[9] = "i"
	Id[10] = "j"
	Id[11] = "k"
	Id[12] = "l"
	Id[13] = "m"
	Id[14] = "n"
	Id[15] = "o"
	Id[16] = "p"
	Id[17] = "q"
	Id[18] = "r"
	Id[15] = "s"
	Id[16] = "t"
	Id[17] = "u"
	Id[18] = "v"
	Id[19] = "w"
	Id[20] = "x"
	Id[21] = "y"
	Id[22] = "z"
}

func ConnString(iniContainer string) string {
	tmp := Get(iniContainer, nil)
	if tmp != nil {
		return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=%s",
			Get(iniContainer+".username", ""),
			Get(iniContainer+".password", ""),
			Get(iniContainer+".server", ""),
			Get(iniContainer+".port", ""),
			Get(iniContainer+".dbname", ""),
			url.QueryEscape(Get(iniContainer+".timezone", "").(string)),
		)
	}
	return ""
}

func Orm() *xorm.Engine {
	if engine == nil {
		var err error
		engine, err = xorm.NewEngine("mysql", wConn)
		if err != nil {
			panic(err)
		}
	}
	return engine
}

func (serv RegisterService) RegisterMachine(macId string) Servers {
	db := Orm()
	var eServer Servers
	//check machine already exist
	eSql := `SELECT * FROM servers WHERE mac_id = ?`
	eRows, _ := db.DB().Query(eSql, macId)
	defer eRows.Close()

	for eRows.Next() {
		err := eRows.ScanStructByIndex(&eServer)
		if err != nil {
			fmt.Println(err)
		}
	}

	var idMachine Servers
	// get last register machine
	idSql := `SELECT * FROM servers ORDER BY id DESC LIMIT 1`
	idRows, _ := db.DB().Query(idSql)
	defer idRows.Close()

	for idRows.Next() {
		err := idRows.ScanStructByIndex(&idMachine)
		if err != nil {
			fmt.Println(err)
		}
	}

	if eServer.Id == 0 {
		machine := new(Servers)
		if idMachine.Id != 0 {
			id := idMachine.Id
			id++
			machine.RegisterId = Id[id]
		} else {
			machine.RegisterId = Id[1]
		}
		machine.MacId = macId
		_, err := db.Insert(machine)
		if err != nil {
			panic(err)
		}
		return *machine
	}
	return eServer
}
