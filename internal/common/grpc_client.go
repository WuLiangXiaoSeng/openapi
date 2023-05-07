package common

// import (
// 	"reflect"
// 	"strconv"
// 	"wuliangxiaoseng/errcode"
// 	"wuliangxiaoseng/openapi/internal/logger"
// )

// /*
// 	约束条件：
// 	1、left和right内的变量名称要一致，次序可以不相同，类型可以不相同，但必须是同一层次且是可转换的：基础类型、struct和切片
// 		bool -> string(false为"false"，true为"true")、int(false为2，true为1，转换为proto文件中CommonBoolean的true和false)
// 		string -> int(可转换为整形)、bool(不为"enable"是false，为"enable"是true)
// 		int -> string(字符串)、bool(不为1是false，为1是true)
// 	2、left是指针类型
// 	3、需要实现left中的方法, 名称为:Allocate+slice变量名称
// 		例如:
// 		type show struct {
// 			Name string
// 			Inst2VlansMap []*ShowMap
// 			Test []*showStruct
// 		}

// 		func (self *show) AllocateInst2VlansMap(n int) {
// 			for i := 0; i < n; i++ {
// 				self.Inst2VlansMap = append(self.Inst2VlansMap, &ShowMap{})
// 			}
// 		}

// 		func (self *show) AllocateTest(n int) {
// 			for i := 0; i < n; i++ {
// 				self.Test = append(self.Test, &showStruct{})
// 			}
// 		}
// */
// func AssignValue(left interface{}, right interface{}) bool {

// 	rightReflectValue := reflect.ValueOf(right)
// 	leftReflectValue := reflect.ValueOf(left)

// 	if leftReflectValue.Kind() != reflect.Ptr {
// 		logger.Errorf("common", errcode.RESULT_ERROR_COMMON, "left is not ptr")
// 		return false
// 	}

// 	leftReflectValue = leftReflectValue.Elem()
// 	if rightReflectValue.Kind() == reflect.Ptr {
// 		rightReflectValue = rightReflectValue.Elem()
// 	}

// 	return reflectValue(leftReflectValue, rightReflectValue)
// }

// func reflectValue(leftReflectValue reflect.Value, rightReflectValue reflect.Value) bool {
// 	rightFieldNum := rightReflectValue.NumField()
// 	leftFieldNum := leftReflectValue.NumField()

// 	for i := 0; i < leftFieldNum; i++ {
// 		leftFieldValue := leftReflectValue.Field(i)
// 		leftFieldType := leftReflectValue.Type().Field(i)
// 		leftFieldTypeKind := leftFieldValue.Kind()
// 		leftName := leftFieldType.Name

// 		for j := 0; j < rightFieldNum; j++ {
// 			rightFieldValue := rightReflectValue.Field(j)
// 			rightFieldType := rightReflectValue.Type().Field(j)
// 			rightFieldTypeKind := rightFieldValue.Kind()
// 			rightName := rightFieldType.Name

// 			if rightName != leftName {
// 				continue
// 			}

// 			switch rightFieldTypeKind {
// 			case reflect.String:
// 				fallthrough
// 			case reflect.Bool:
// 				fallthrough
// 			case reflect.Int32:
// 				fallthrough
// 			case reflect.Int64:
// 				fallthrough
// 			case reflect.Uint32:
// 				fallthrough
// 			case reflect.Uint64:
// 				if !setReflectValue(leftFieldValue, rightFieldValue) {
// 					logger.Errorf("common", errcode.RESULT_ERROR_COMMON, "assignment failed")
// 					return false
// 				}
// 			case reflect.Slice:
// 				if leftFieldTypeKind != reflect.Slice {
// 					return true
// 				}
// 				funName := "Allocate" + rightName
// 				fun := leftReflectValue.Addr().MethodByName(funName)
// 				if !fun.IsValid() {
// 					logger.Errorf("common", errcode.RESULT_ERROR_COMMON, "no implementation method:"+funName)
// 					return false
// 				}
// 				funPara := []reflect.Value{reflect.ValueOf(rightFieldValue.Len())}
// 				fun.Call(funPara)
// 				for j := 0; j < rightFieldValue.Len(); j++ {
// 					rightSliceValue := rightFieldValue.Index(j)
// 					leftSliceValue := leftFieldValue.Index(j)
// 					if rightSliceValue.Kind() == reflect.Ptr {
// 						rightSliceValue = rightSliceValue.Elem()
// 						leftSliceValue = leftSliceValue.Elem()
// 					}
// 					switch rightSliceValue.Kind() {
// 					case reflect.String:
// 						fallthrough
// 					case reflect.Bool:
// 						fallthrough
// 					case reflect.Int32:
// 						fallthrough
// 					case reflect.Int64:
// 						fallthrough
// 					case reflect.Uint32:
// 						fallthrough
// 					case reflect.Uint64:
// 						if !setReflectValue(leftSliceValue, rightSliceValue) {
// 							logger.Errorf("common", errcode.RESULT_ERROR_COMMON, "slice assignment failed")
// 							return false
// 						}
// 					case reflect.Struct:
// 						reflectValue(leftSliceValue, rightSliceValue)
// 					}
// 				}
// 			case reflect.Struct:
// 				reflectValue(leftFieldValue, rightFieldValue)
// 			default:
// 				logger.Errorf("common", errcode.RESULT_ERROR_COMMON, "type is wrong")
// 				return false
// 			}
// 		}
// 	}
// 	return true
// }

// func setReflectValue(leftFieldValue reflect.Value, rightFieldValue reflect.Value) bool {
// 	leftFieldTypeKind := leftFieldValue.Kind()
// 	rightFieldTypeKind := rightFieldValue.Kind()
// 	switch rightFieldTypeKind {
// 	case reflect.String:
// 		value := rightFieldValue.String()
// 		switch leftFieldTypeKind {
// 		case reflect.String:
// 			leftFieldValue.SetString(rightFieldValue.String())
// 		case reflect.Bool:
// 			leftFieldValue.SetBool(false)
// 			if value == "enable" {
// 				leftFieldValue.SetBool(true)
// 			}
// 		case reflect.Int32:
// 			fallthrough
// 		case reflect.Int64:
// 			tmp, err := strconv.ParseInt(value, 10, 64)
// 			if err != nil {
// 				logger.Errorf("common", errcode.RESULT_ERROR_COMMON, "Failed to convert type(%v,%v)", leftFieldTypeKind, rightFieldTypeKind)
// 				return false
// 			}
// 			leftFieldValue.SetInt(tmp)
// 		case reflect.Uint32:
// 			fallthrough
// 		case reflect.Uint64:
// 			tmp, err := strconv.ParseUint(value, 10, 64)
// 			if err != nil {
// 				logger.Errorf("common", errcode.RESULT_ERROR_COMMON, "Failed to convert type(%v,%v)", leftFieldTypeKind, rightFieldTypeKind)
// 				return false
// 			}
// 			leftFieldValue.SetUint(tmp)
// 		default:
// 			logger.Errorf("common", errcode.RESULT_ERROR_COMMON, "Type(%v) conversion is not supported", leftFieldTypeKind)
// 			return false
// 		}
// 	case reflect.Bool:
// 		value := rightFieldValue.Bool()
// 		switch leftFieldTypeKind {
// 		case reflect.String:
// 			leftFieldValue.SetString("disable")
// 			if value {
// 				leftFieldValue.SetString("enable")
// 			}
// 		case reflect.Bool:
// 			leftFieldValue.SetBool(value)
// 		case reflect.Int32:
// 			fallthrough
// 		case reflect.Int64:
// 			leftFieldValue.SetInt(2)
// 			if value {
// 				leftFieldValue.SetInt(1)
// 			}
// 		case reflect.Uint32:
// 			fallthrough
// 		case reflect.Uint64:
// 			leftFieldValue.SetInt(2)
// 			if value {
// 				leftFieldValue.SetUint(1)
// 			}
// 		default:
// 			logger.Errorf("common", errcode.RESULT_ERROR_COMMON, "Type(%v) conversion is not supported", leftFieldTypeKind)
// 			return false
// 		}
// 	case reflect.Int32:
// 		fallthrough
// 	case reflect.Int64:
// 		value := rightFieldValue.Int()
// 		switch leftFieldTypeKind {
// 		case reflect.String:
// 			leftFieldValue.SetString(strconv.Itoa(int(value)))
// 		case reflect.Bool:
// 			leftFieldValue.SetBool(false)
// 			if value == 1 {
// 				leftFieldValue.SetBool(true)
// 			}
// 		case reflect.Int32:
// 			fallthrough
// 		case reflect.Int64:
// 			leftFieldValue.SetInt(value)
// 		case reflect.Uint32:
// 			fallthrough
// 		case reflect.Uint64:
// 			leftFieldValue.SetUint(uint64(value))
// 		default:
// 			logger.Errorf("common", errcode.RESULT_ERROR_COMMON, "Type(%v) conversion is not supported", leftFieldTypeKind)
// 			return false
// 		}
// 	case reflect.Uint32:
// 		fallthrough
// 	case reflect.Uint64:
// 		value := rightFieldValue.Uint()
// 		switch leftFieldTypeKind {
// 		case reflect.String:
// 			leftFieldValue.SetString(strconv.Itoa(int(value)))
// 		case reflect.Bool:
// 			leftFieldValue.SetBool(false)
// 			if value == 1 {
// 				leftFieldValue.SetBool(true)
// 			}
// 		case reflect.Int32:
// 			fallthrough
// 		case reflect.Int64:
// 			leftFieldValue.SetInt(int64(value))
// 		case reflect.Uint32:
// 			fallthrough
// 		case reflect.Uint64:
// 			leftFieldValue.SetUint(value)
// 		default:
// 			logger.Errorf("common", errcode.RESULT_ERROR_COMMON, "Type(%v) conversion is not supported", leftFieldTypeKind)
// 			return false
// 		}
// 	}
// 	return true
// }
