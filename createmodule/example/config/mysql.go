package config

import (
	"fmt"
)

type Mysql struct {
	DataBase   string
	UserName   string
	Password   string
	Parameters string
	//MaxIdle    int
	//MaxOpen    int
	//Debug      bool
}

//[username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
//"root:123456@/study?charset=utf8&parseTime=True&loc=Local"
func (conf *Mysql) String() string {
	return fmt.Sprintf("%s:%s@/%s?%s",conf.UserName,conf.Password,conf.DataBase,conf.Parameters)
}
//
//func (conf *Mysql) ToJSON() []byte {
//	data, _ := json.Marshal(conf)
//	return data
//}
//
//func (conf *Mysql) FromJSON(data []byte) error {
//	return json.Unmarshal(data, conf)
//}
//
//func (conf *Mysql) Save(file string) error {
//	data := conf.ToJSON()
//	return ioutil.WriteFile(file, data, 0644)
//}
//
//func (conf *Mysql) Load(file string) error {
//
//}