package models

import (
	"database/sql"
	"fmt"
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
	Id          int       `orm:"auto"`
	Title       string    `orm:"size(500)"`
	Href        string    `orm:"size(1000)"`
	ImgUrl      string    `orm:"szie(1000)"`
	Description string    `orm:"szie(1000)"`
	Source      string    `orm:"szie(255)"`
	Country     string    `orm:"szie(255)"`
	Src         string    `orm:"size(500)"`
	Content     string    `orm:"column(content)"`
	CreateTime  time.Time `orm:"type(datetime)"`
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
}

func GetAllNewsEn(page int) []H1 {
	var lists []H1
	size := 20
	o := orm.NewOrm()
	// var news []*H1
	mysql := fmt.Sprintf("select * from nytimes ny, nytimes_details nyd where ny.id = nyd.nytimes_id order by ny.create_time desc limit %d,20;", (page-1)*size)
	// mysql := "select * from nytimes by create_time desc limit 20;"

	num, err := o.Raw(mysql).QueryRows(&lists)
	// o.QueryTable("nytimes").OrderBy("-Ctime").All(&users)
	// num, err := o.Raw("SELECT * FROM nytimes").QueryRows(&news)
	if err != nil {
		fmt.Println("user nums: ", num, err)
	}
	Lists := make([]H1, 0)
	for _, item := range lists {
		Lists = append(lists, H1{item.Id, item.Title, item.Src, item.Href, item.ImgUrl, item.Description, item.Country, item.Source, item.Content, item.CreateTime})
	}

	// for _, term := range maps {
	// 	fmt.Println(term["id"], ":", term["title"])
	// }

	// fmt.Print(num, reflect.TypeOf(maps))

	// list := make(map[int]*H1, num)
	// for k, v := range maps {

	// 	list[k] = &H1{v["Id"], v["Title"]}
	// 	fmt.Print(k, v)
	// }

	return Lists
}
