package util

import (
	"encoding/json"
	"reflect"
)

func BindJSONToStruct(data interface{}, targetStruct interface{}) error {
	// 解析 JSON 数据到 map[string]interface{} 对象
	var dataMap map[string]interface{}
	dataMap = data.(map[string]interface{})

	// 使用反射遍历目标结构体的字段
	targetValue := reflect.ValueOf(targetStruct).Elem()
	targetType := targetValue.Type()
	for i := 0; i < targetType.NumField(); i++ {
		field := targetType.Field(i)
		tagValue := field.Tag.Get("json")

		// 跳过没有 json 标签的字段
		if tagValue == "" {
			continue
		}

		// 获取字段的值并进行类型转换
		fieldValue := dataMap[tagValue]
		fieldValueRef := reflect.ValueOf(fieldValue)

		// 处理嵌套结构体
		if field.Type.Kind() == reflect.Struct && fieldValueRef.Kind() == reflect.Map {
			nestedStruct := reflect.New(field.Type).Interface()
			err := BindJSONToStruct(convertMapToJSON(fieldValueRef.Interface()), nestedStruct)
			if err != nil {
				return err
			}
			targetValue.Field(i).Set(reflect.ValueOf(nestedStruct).Elem())
		} else {
			if fieldValueRef.IsValid() && fieldValueRef.Type().ConvertibleTo(field.Type) {
				targetField := targetValue.Field(i)
				targetField.Set(fieldValueRef.Convert(field.Type))
			}
		}
	}

	return nil
}

// 辅助函数：将 interface{} 类型的 map 转换为 JSON 字节数组
func convertMapToJSON(data interface{}) []byte {
	jsonData, _ := json.Marshal(data)
	return jsonData
}
