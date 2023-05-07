package logger

// import (

// 	// "ctccmler/cml"
// 	"errors"
// 	"fmt"
// 	"wuliangxiaoseng/errcode"

// 	// "hstcmler/xlog"
// 	"time"

// 	"wuliangxiaoseng/openapi/internal/config"
// 	"wuliangxiaoseng/openapi/internal/consts"
// 	// grpc "equantum.com/httpserver/internal/grpc_conn"
// )

// var logger *xlog.XLogger

// // var syslog_client cml.LogServiceClient

// const (
// 	ALARM_LEVEL_ERROR    = 3
// 	ALARM_LEVEL_WARNNING = 4
// 	ALARM_LEVEL_NOTICE   = 5
// )

// func setLogLevel(level int) xlog.LOG_LEVEL {
// 	switch level {
// 	case consts.LOG_LEVEL_EMERGENCY:
// 		return xlog.LOG_LEVEL_EMERGENCY
// 	case consts.LOG_LEVEL_ALERT:
// 		return xlog.LOG_LEVEL_ALERT
// 	case consts.LOG_LEVEL_CRITICAL:
// 		return xlog.LOG_LEVEL_CRITICAL
// 	case consts.LOG_LEVEL_ERROR:
// 		return xlog.LOG_LEVEL_ERROR
// 	case consts.LOG_LEVEL_WARNING:
// 		return xlog.LOG_LEVEL_WARNING
// 	case consts.LOG_LEVEL_NOTICE:
// 		return xlog.LOG_LEVEL_NOTICE
// 	case consts.LOG_LEVEL_INFO:
// 		return xlog.LOG_LEVEL_INFO
// 	case consts.LOG_LEVEL_DEBUG:
// 		return xlog.LOG_LEVEL_DEBUG
// 	default:
// 		return xlog.LOG_LEVEL_EMERGENCY
// 	}
// }

// type LogStru struct {
// 	Type      cml.WebLogReq_LOG_TYPE `json:"type"`
// 	Level     cml.LOG_LEVEL          `json:"level"`
// 	IEDName   string                 `json:"iedname"`
// 	DEVName   string                 `json:"devname"`
// 	TimeStamp uint64                 `json:"time_stamp"`
// 	Position  string                 `json:"position"`
// 	Desc      string                 `json:"desc"`
// }

// type AlarmStru struct {
// 	Type           cml.WebLogReq_LOG_TYPE `json:"type"`
// 	Level          cml.LOG_LEVEL          `json:"level"`
// 	IEDName        string                 `json:"iedname"`
// 	DEVName        string                 `json:"devname"`
// 	RaiseTimeStamp uint64                 `json:"raise_time_stamp"`
// 	ClearTimeStamp uint64                 `json:"clear_time_stamp"`
// 	Position       string                 `json:"position"`
// 	Desc           string                 `json:"desc"`
// }

// type WebLogReq struct {
// 	Type       cml.WebLogReq_LOG_TYPE `json:"type"`
// 	ModuleName string                 `json:"module_name"`
// 	Level      cml.LOG_LEVEL          `json:"level"`
// 	TimeStamp  uint64                 `json:"time_stamp"`
// 	Position   string                 `json:"position"`
// 	ErrCode    int32                  `json:"err_code"`
// 	Desc       string                 `json:"desc"`
// 	Content    string                 `json:"content"`
// }

// func Oplog(desc string, v ...interface{}) {
// 	logger.OpLog(consts.LOG_MODEL_NAME, desc, v)
// }

// func Alertf(position string, err errcode.RESULT, desc string, v ...interface{}) {
// 	logger.Alertf(position, err, desc, v)
// }

// func Errorf(position string, err errcode.RESULT, desc string, v ...interface{}) {
// 	logger.Errorf(position, err, desc, v)
// }

// func Warningf(position string, err errcode.RESULT, desc string, v ...interface{}) {
// 	logger.Warningf(position, err, desc, v)
// }

// func Infof(position string, err errcode.RESULT, desc string, v ...interface{}) {
// 	logger.Infof(position, err, desc, v)
// }

// func Debugf(position string, err errcode.RESULT, desc string, v ...interface{}) {
// 	logger.Debugf(position, err, desc, v)
// }

// //交换机接口

// // func generateAlarmorLog(reqlog WebLogReq) {

// // 	syslog_client.SetWebOprLog(context.Background(), &cml.WebLogReq{
// // 		Type:       reqlog.Type,
// // 		ModuleName: reqlog.ModuleName,
// // 		Level:      reqlog.Level,
// // 		Position:   reqlog.Position,
// // 		ErrCode:    reqlog.ErrCode,
// // 		Desc:       reqlog.Desc,
// // 		Content:    reqlog.Content,
// // 	})

// // }

// //1. module name: 用户名
// //2. position：ip:web (web写死吧
// func OplogV2(position, descrip string, v ...interface{}) {

// 	desc := fmt.Sprintf(descrip, v...)
// 	generateAlarmorLog(WebLogReq{
// 		Type:       1,
// 		ModuleName: consts.LOG_MODEL_NAME,
// 		Position:   position,
// 		//Level:      level,
// 		//ErrCode:   errCode,
// 		Desc: desc,
// 		//Content:    content,
// 	})

// }

// func AlertfV2(position string, err errcode.RESULT, logDesc string, v ...interface{}) {

// 	desc := fmt.Sprintf(logDesc, v...)
// 	generateAlarmorLog(WebLogReq{
// 		Type:       cml.WebLogReq_SYSTEM,
// 		Level:      cml.LOG_LEVEL_LOG_LEVEL_ALERT,
// 		ModuleName: consts.LOG_MODEL_NAME,
// 		Position:   position,
// 		ErrCode:    int32(err),
// 		Desc:       desc,
// 	})

// }

// func ErrorfV2(position string, err errcode.RESULT, logDesc string, v ...interface{}) {

// 	desc := fmt.Sprintf(logDesc, v...)
// 	generateAlarmorLog(WebLogReq{
// 		Type:       cml.WebLogReq_SYSTEM,
// 		ModuleName: consts.LOG_MODEL_NAME,
// 		Level:      cml.LOG_LEVEL_LOG_LEVEL_ERROR,
// 		Position:   position,
// 		ErrCode:    int32(err),
// 		Desc:       desc,
// 	})
// }

// func WarningfV2(position string, err errcode.RESULT, logDesc string, v ...interface{}) {

// 	desc := fmt.Sprintf(logDesc, v...)
// 	generateAlarmorLog(WebLogReq{
// 		Type:       cml.WebLogReq_SYSTEM,
// 		ModuleName: consts.LOG_MODEL_NAME,
// 		Level:      cml.LOG_LEVEL_LOG_LEVEL_WARNING,
// 		Position:   position,
// 		ErrCode:    int32(err),
// 		Desc:       desc,
// 	})
// }

// func InfofV2(position string, err errcode.RESULT, logDesc string, v ...interface{}) {

// 	desc := fmt.Sprintf(logDesc, v...)
// 	generateAlarmorLog(WebLogReq{
// 		Type:       cml.WebLogReq_SYSTEM,
// 		ModuleName: consts.LOG_MODEL_NAME,
// 		Level:      cml.LOG_LEVEL_LOG_LEVEL_INFO,
// 		Position:   position,
// 		ErrCode:    int32(err),
// 		Desc:       desc,
// 	})
// }

// func DebugfV2(position string, err errcode.RESULT, logDesc string, v ...interface{}) {

// 	desc := fmt.Sprintf(logDesc, v...)
// 	generateAlarmorLog(WebLogReq{
// 		Type:       cml.WebLogReq_SYSTEM,
// 		ModuleName: consts.LOG_MODEL_NAME,
// 		Level:      cml.LOG_LEVEL_LOG_LEVEL_DEBUG,
// 		Position:   position,
// 		ErrCode:    int32(err),
// 		Desc:       desc,
// 	})
// }

// func AlarmfV2(position, loginuser string, alarmcode, level int32, logDesc string, v ...interface{}) {

// 	desc := fmt.Sprintf(logDesc, v...)
// 	generateAlarmorLog(WebLogReq{
// 		Type:       cml.WebLogReq_ALARM,
// 		ModuleName: loginuser,
// 		Level:      cml.LOG_LEVEL(level),
// 		Position:   position,
// 		ErrCode:    alarmcode,
// 		Desc:       desc,
// 	})
// }

// func OPerAlarm(position, loginuser string, alarmcode, level int32, logDesc string, v ...interface{}) {

// 	desc := fmt.Sprintf(logDesc, v...)
// 	generateAlarmorLog(WebLogReq{
// 		Type:       cml.WebLogReq_OPERATION,
// 		ModuleName: loginuser,
// 		Level:      cml.LOG_LEVEL(level),
// 		Position:   position,
// 		ErrCode:    alarmcode,
// 		Desc:       desc,
// 	})
// }

// func OperLogFormate(log LogStru) (string, error) {
// 	if log.Type != cml.WebLogReq_OPERATION {
// 		return "", errors.New("type wrong")
// 	}

// 	var levelstr = "NOTICE"
// 	timestring := time.Unix((int64)(log.TimeStamp), 0).Format("2006-01-02 15:04:05")

// 	// if int(log.Level) <= 3 {
// 	// 	levelstr = "ERROR"
// 	// } else if int(log.Level) == 4 {
// 	// 	levelstr = "WARNNING"
// 	// } else {
// 	// 	levelstr = "NOTICE"

// 	// }
// 	onelogstring := fmt.Sprintf("%%%s|%s|%s|%s|position:%s,%s", levelstr, timestring, log.IEDName, log.DEVName, log.Position, log.Desc)
// 	return onelogstring, nil
// }

// func AlarmFormate(alarm AlarmStru) (string, error) {
// 	if alarm.Type != cml.WebLogReq_ALARM {
// 		return "", errors.New("type wrong")
// 	}
// 	if alarm.Desc == "" {
// 		return "", errors.New("no description")
// 	}
// 	var levelstr, raiseOrclear, timestring string
// 	if alarm.ClearTimeStamp > alarm.RaiseTimeStamp {
// 		raiseOrclear = "Cleared:"
// 		//fmt.Printf("Clear: clear:alarm.ClearTimeStamp:%d  alarm.RaiseTimeStamp:%d\n", alarm.ClearTimeStamp, alarm.RaiseTimeStamp)
// 		timestring = time.Unix((int64)(alarm.ClearTimeStamp), 0).Format("2006-01-02 15:04:05")
// 	} else {
// 		raiseOrclear = "Raised:"
// 		//fmt.Printf("Raise: clear:alarm.ClearTimeStamp:%d  alarm.RaiseTimeStamp:%d\n", alarm.ClearTimeStamp, alarm.RaiseTimeStamp)
// 		timestring = time.Unix((int64)(alarm.RaiseTimeStamp), 0).Format("2006-01-02 15:04:05")
// 	}

// 	if int(alarm.Level) <= 3 {
// 		levelstr = "ERROR"
// 	} else if int(alarm.Level) == 4 {
// 		levelstr = "WARNNING"
// 	} else {
// 		levelstr = "NOTICE"

// 	}
// 	onelogstring := fmt.Sprintf("%%%s|%s|%s|%s|%sposition:%s,%s", levelstr, timestring, alarm.IEDName, alarm.DEVName, raiseOrclear, alarm.Position, alarm.Desc)
// 	return onelogstring, nil
// }

// func Init() consts.ReturnCode {
// 	logger = xlog.NewXLogger(consts.LOG_MODEL_NAME, consts.HOST_LOOPBACK_ADDR, 0, consts.HOST_LOOPBACK_ADDR, config.ServerConfig.LogPort)

// 	if logger == nil {
// 		return consts.Err
// 	}
// 	// syslog_client = cml.NewLogServiceClient(grpc.Getconn())

// 	return consts.Ok
// }
