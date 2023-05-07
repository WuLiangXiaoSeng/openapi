package logger

const (
	WEB_ALARM_BEGIN = 2000
	//web模块告警
	WEB_ALARM_LOGIN_SUCCESS  = WEB_ALARM_BEGIN + 1
	WEB_ALARM_LOGIN_FAIL     = WEB_ALARM_BEGIN + 2
	WEB_ALARM_LOGOUT_SUCCESS = WEB_ALARM_BEGIN + 3
	WEB_ALARM_LOGOUT_FAIL    = WEB_ALARM_BEGIN + 4
	WEB_ALARM_USER_MODIFY_SUCCESS    = WEB_ALARM_BEGIN + 5
	WEB_ALARM_USER_MODIFY_FAIL    = WEB_ALARM_BEGIN + 6
	WEB_ALARM_USER_DELETE_SUCCESS    = WEB_ALARM_BEGIN + 7
	WEB_ALARM_USER_DELETE_FAIL    = WEB_ALARM_BEGIN + 8
	WEB_ALARM_NTP_TIMECHANGE = WEB_ALARM_BEGIN + 9
	WEB_ALARM_SYS_REBOOT     = WEB_ALARM_BEGIN + 10
)
