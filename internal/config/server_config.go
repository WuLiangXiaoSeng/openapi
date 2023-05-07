package config

type serverConfig struct {
	AppName     string `json:"app_name"`
	Env         string `json:"env"`
	ListenPort  int    `json:"listen_port"`
	LogPort     int    `json:"log_port"`
	LogLevel    string `json:"log_level"`
	Title       string `json:"title"`
	Keywords    string `json:"keywords"`
	Description string `json:"description"`
	Icp         string `json:"icp"`
	TimeStamp   int    `json:"time_stamp"`
}
