package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"unicode/utf8"
)

var LearnNumOneDay int = 10

type configData struct {
	Grpc   grpcConfig   `json:"grpc"`
	Server serverConfig `json:"server"`
}

func initPath() {
	sep := string(os.PathSeparator)
	//root := filepath.Dir(os.Args[0])
	//ExecPath, _ = filepath.Abs(root)
	ExecPath, _ = os.Getwd()
	length := utf8.RuneCountInString(ExecPath)
	lastChar := ExecPath[length-1:]
	if lastChar != sep {
		ExecPath = ExecPath + sep
	}
}

var ConfigPath string = "/anos/scripts/httpserver/"

func initConfigJSON() {
	rawConfig, err := ioutil.ReadFile(ConfigPath + "config.json")
	if err != nil {
		//未初始化
		rawConfig = []byte("{\"grpc\":{\"server_port\":50001},\"server\":{\"app_name\":\"123\",\"env\": \"development\",\"listen_port\": 80,\"log_port\": 30001,\"log_level\":\"debug\",\"time_stamp\": 30}}")
	}
	if err := json.Unmarshal(rawConfig, &JsonData); err != nil {
		fmt.Println("Invalid Config: ", err.Error())
		os.Exit(-1)
	}
}

func initServer() {
	ServerConfig = JsonData.Server
	GrpcConfig = JsonData.Grpc
	sep := string(os.PathSeparator)
	//root := filepath.Dir(os.Args[0])
	//ExecPath, _ = filepath.Abs(root)
	ExecPath, _ = os.Getwd()
	length := utf8.RuneCountInString(ExecPath)
	lastChar := ExecPath[length-1:]
	if lastChar != sep {
		ExecPath = ExecPath + sep
	}
}

var ExecPath string
var JsonData configData
var ServerConfig serverConfig
var GrpcConfig grpcConfig

//var GrpcConnect grpcConnect

func init() {
	initPath()
	initConfigJSON()
	initServer()
	/*
		initGrpcConnect()
	*/
}

func WriteConfig() error {
	//将现有配置写回文件
	configFile, err := os.OpenFile(fmt.Sprintf("%sconfig.json", ExecPath), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}

	defer configFile.Close()

	buff := &bytes.Buffer{}

	buf, err := json.MarshalIndent(JsonData, "", "\t")
	if err != nil {
		return err
	}
	buff.Write(buf)

	_, err = io.Copy(configFile, buff)
	if err != nil {
		return err
	}

	return nil
}
