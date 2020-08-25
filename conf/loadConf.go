package conf

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"os"
	"yeeep.cn/yepgo/model"
)

func LoadDBConf() model.DatabaseConf {
	file, _ := os.Open("conf/database.json")
	defer file.Close()

	decoder := json.NewDecoder(file)
	conf := model.DatabaseConf{}
	err := decoder.Decode(&conf)
	if err != nil {
		logrus.Error(err.Error())
		panic("failed to read database conf, err:" + err.Error())
	}
	logrus.Debug(conf)
	return conf
}
