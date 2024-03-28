package setting

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"strconv"
)

var appConf map[string]any

type AppInfo struct {
	Name string
	Port int
}

type CorsConfInfo struct {
	Debug bool
}

func init() {
	confFile := "./conf/app.yaml"
	buff, err := os.ReadFile(confFile)
	if err != nil {
		log.Println(err)
		return
	}

	if err := yaml.Unmarshal(buff, &appConf); err != nil {
		log.Println(err)
		return
	}
}

func GetLocalStorePath() string {
	s3Conf := getMapInfo(appConf, "store")
	return convertString(s3Conf["path"])
}

func GetAppInfo() *AppInfo {
	v, ok := appConf["app"]
	if !ok {
		return nil
	}

	mp := v.(map[string]any)
	return &AppInfo{
		Name: GetFromMapString(mp, "name"),
		Port: getFromMapInt(mp, "port"),
	}
}

func GetCorsConf() *CorsConfInfo {
	v, ok := appConf["cors"]
	if !ok {
		return nil
	}

	mp := v.(map[string]any)
	debug := false
	debugValue, ok := mp["debug"]
	if ok {
		debug = debugValue.(bool)
	}

	return &CorsConfInfo{
		Debug: debug,
	}
}

func GetFromMapString(mpValue map[string]any, key string) string {
	v, ok := mpValue[key]
	if !ok {
		return ""
	}
	return v.(string)
}

func getFromMapInt(mpValue map[string]any, key string) int {
	v, ok := mpValue[key]
	if !ok {
		return 0
	}
	return v.(int)
}

func getMapInfo(conf map[string]any, key string) map[string]any {
	v, ok := conf[key]
	if !ok {
		return nil
	}
	return v.(map[string]any)
}

func convertString(raw any) string {
	v, ok := raw.(int)
	if ok {
		return strconv.Itoa(v)
	}

	return raw.(string)
}
