package models


import(
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Article struct {
	Id int
	Title string
	Content string
}



func AllArticles() ([]Article, error) {
	var artiles []Article
	db, _ := GetDb()
	err := db.Find(&artiles).Error
	db.Close()
	return artiles,err
}

func (article *Article) Insert() error {
	db, _ := GetDb()
	db.NewRecord(article)
	err := db.Create(&article).Error
	db.Close()
	return err
}

func (article *Article) Delete() error {
	db,_ := GetDb()
	err := db.Delete(&article).Error
	db.Close()
	return err
}

func (article *Article) Update() error  {
	db,_ := GetDb()
	err := db.Save(&article).Error
	db.Close()
	return err
}

func (article *Article) Find() error  {
	db,_ := GetDb()
	err := db.First(&article).Error
	db.Close()
	return err
}

func FindArticles(key string) ([]Article, error) {
	var artiles []Article
	db, _ := GetDb()
	err := db.Where("content LIKE ?", "%"+key+"%").Find(&artiles).Error
	db.Close()
	return artiles,err
}


func GetDb() (*gorm.DB, error) {
	return gorm.Open("mysql", "root:123456@/study?charset=utf8&parseTime=True&loc=Local")
}

