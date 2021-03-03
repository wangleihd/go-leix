package models

import (
	"database/sql"
	"fmt"
	"reflect"
	"time"

	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql" // 导入数据库驱动
)

var (
	NewsLists map[string]*H1
)

const (
	dbType   string = "mysql"
	dbHost   string = "sql427.main-hosting.eu"
	dbPort   string = "3306"
	dbName   string = "u637214094_spider"
	dbUser   string = "u637214094_spider"
	dbPasswd string = "Aasdfgh12@"
	dbParams string = "charset=utf8mb4&parseTime=true"
)

type DBInit struct {
	Db *sql.DB
}

type H1 struct {
	Id    int    `orm:"auto"`
	Title string `orm:"size(500)"`
}

type Infor struct {
	Gender  string
	Age     int
	Address string
	Email   string
}

func init() {
	// dbInit().initDatabase()
	orm.RegisterDriver("mysql", orm.DRMySQL)
	sqlConn := fmt.Sprintf("%s:%s@(%s:%s)/%s?%s", dbUser, dbPasswd, dbHost, dbPort, dbName, dbParams)
	orm.RegisterDataBase("default", "mysql", sqlConn)

	// orm.RegisterModel(new(H1))

	orm.DefaultTimeLoc = time.UTC
	// create table
	orm.RunSyncdb("default", false, true)

	NewsLists = make(map[string]*H1)
	info := H1{11122233, "astaxie"}
	NewsLists["user_11111"] = &info
}

func GetAllNewsEn() map[int]*H1 {
	var maps []orm.Params
	o := orm.NewOrm()
	// var news []*H1

	num, err := o.Raw("SELECT * FROM nytimes").Values(&maps)
	// o.QueryTable("nytimes").OrderBy("-Ctime").All(&users)
	// num, err := o.Raw("SELECT * FROM nytimes").QueryRows(&news)
	if err != nil {
		fmt.Println("user nums: ", num, err)
	}
	for _, term := range maps {
		fmt.Println(term["id"], ":", term["title"])
	}

	fmt.Print(num, reflect.TypeOf(maps))

	list := make(map[int]*H1, num)
	for k, v := range maps {
		
		list[k] = &H1{v["Id"], v["Title"]}
		fmt.Print(k, v)
	}

	return list
}
