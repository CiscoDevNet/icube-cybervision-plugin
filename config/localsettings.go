package config

import (
	_ "bufio"
	_ "fmt"
	_ "io"
	"io/ioutil"
	"os"
	"time"
	"encoding/json"
)


type LocalSetting struct {
	LastSyncUp int64 `config:"last_syncup"`
	SettingFile string   `default:"/tmp/cvsettings.json"`
}

func (ls *LocalSetting) GetDefaultStartTime() time.Time {
	time_out,err := time.Parse(time.RFC3339, "2000-01-01T00:00:00:000Z")
	if err ==nil{
		return time.Now()
	}else {
		return time_out
	}
}
func (ls *LocalSetting) Load( file_name string) (bool, error){
	jsonFile, err := os.Open(file_name)

	if err !=nil {
		ls.LastSyncUp = ls.GetDefaultStartTime().Unix()
		return true,err
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	err_m := json.Unmarshal(byteValue, &ls)
	if err_m !=nil{
		ls.LastSyncUp = ls.GetDefaultStartTime().Unix()
	}

	return true,err_m
}

func (ls *LocalSetting) Save() (int,error){
	jsonString, _ := json.Marshal(ls)
	err :=ioutil.WriteFile(ls.SettingFile, jsonString, os.ModePerm)
	return 0,err
}

func (ls *LocalSetting) Update(to_time time.Time) (bool,error){
	ls.LastSyncUp = to_time.Unix()
	_, err :=ls.Save()
	if err !=nil{
		return false,err
	}
	return true,nil
}



