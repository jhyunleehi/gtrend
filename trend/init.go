package trend

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"gopkg.in/yaml.v2"

	log "github.com/sirupsen/logrus"
)

type Logging struct {
	Level     string `yaml:"level"`
	Formatter string `yaml:"formatter"`
	Writer    string `yaml:"writer"`
	Directory string `yaml:"directory"`
	Tracing   string `yaml:"caller_tracing"`
	Coloring  string `yaml:"coloring"`
}

type Config struct {
	Collect time.Duration `yaml:"collect"`
	Logging Logging       `yaml:"logging"`
}

var conf Config

func init() {

	filename, err := filepath.Abs("./config.yml")
	if err != nil {
		fmt.Printf("%s", err.Error())
		return
	}
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("%s", err.Error())
		return
	}

	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		panic(err)
	}

	level, err := log.ParseLevel(conf.Logging.Level)
	if err != nil {
		log.Errorf("%v", err)
		return
	}
	//logging
	log.SetLevel(level)
	log.SetReportCaller(true)
	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
	logdir := conf.Logging.Directory
	finfo, err := os.Stat(logdir)
	if os.IsNotExist(err) {
		err := os.MkdirAll(logdir, 0755)
		if err != nil {
			log.Error(err)
		}
		log.Debugf("%v", finfo)
	}
	logFile, err := os.OpenFile(logdir+"////trend_"+time.Now().Format("2006-01-02")+".log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Error(err.Error())
	}
	log.Debug(logFile)
	
	log.SetOutput(logFile)
	
	//log.SetOutput(os.Stdout)

}
