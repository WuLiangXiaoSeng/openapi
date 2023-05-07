package consts

type ReturnCode int32

const (
	Ok       ReturnCode = 0
	Err      ReturnCode = 1
	CfgErr   ReturnCode = 2
	QueryErr ReturnCode = 3
)

const (
	QueryCfgOk  string = "查询配置成功"
	QueryCfgErr string = "查询配置失败"
	SetCfgOk    string = "下发配置成功"
	SetCfgErr   string = "下发配置失败"
	LogoutOk    string = "登出成功"
)
