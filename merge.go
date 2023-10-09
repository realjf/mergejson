// #############################################################################
// # File: merge.go                                                            #
// # Project: mergejson                                                        #
// # Created Date: 2023/10/09 09:42:14                                         #
// # Author: realjf                                                            #
// # -----                                                                     #
// # Last Modified: 2023/10/09 10:56:36                                        #
// # Modified By: realjf                                                       #
// # -----                                                                     #
// # Copyright (c) 2023 realjf                                                 #
// #############################################################################
package mergejson

import (
	"encoding/json"
	"reflect"
)

func MergeJson(src, dst []byte) (res []byte, err error) {
	var srcStruct, srcStruct2 map[string]interface{}
	err = json.Unmarshal(src, &srcStruct)
	if err != nil {
		return
	}
	err = json.Unmarshal(dst, &srcStruct2)
	if err != nil {
		return
	}

	resMap := map[string]interface{}{}
	for key, val := range srcStruct {
		if val2, ok := srcStruct2[key]; ok {
			// 存在，需要合并
			valDst, err := mergeJson(val, val2)
			if err != nil {
				return nil, err
			}
			resMap[key] = valDst
		} else {
			// 不存在，保留
			resMap[key] = val
		}
	}

	for key, val := range srcStruct2 {
		if val2, ok := resMap[key]; ok {
			// 存在，需要合并
			valDst, err := mergeJson(val2, val)
			if err != nil {
				return nil, err
			}
			resMap[key] = valDst
		} else {
			// 不存在，保留
			resMap[key] = val
		}
	}

	res, err = json.Marshal(resMap)
	return
}

func mergeJson(val interface{}, val2 interface{}) (valDst interface{}, err error) {
	valRef := reflect.ValueOf(val)
	valRef2 := reflect.ValueOf(val2)
	if !valRef2.IsValid() {
		return val, nil
	}
	if !valRef.IsValid() {
		return val2, nil
	}
	switch valRef2.Kind() {
	case reflect.Map, reflect.Struct:
		if kindEqual(valRef, valRef2) {
			// 合并
			valDstMap := make(map[string]interface{}, 0)
			for _, k := range valRef.MapKeys() {
				if _, exist := valDstMap[k.String()]; exist {
					continue
				}
				if valRef2.IsValid() && valRef2.MapIndex(k).IsValid() && valRef2.MapIndex(k).CanInterface() {
					// 存在
					v, err := mergeJson(valRef.MapIndex(k).Interface(), valRef2.MapIndex(k).Interface())
					if err != nil {
						return nil, err
					}
					valDstMap[k.String()] = v
				} else {
					// 不存在，保留
					valDstMap[k.String()] = valRef.MapIndex(k).Interface()
				}
			}
			return valDstMap, nil
		} else {
			valDst = val2
			return
		}
	case reflect.Slice:
		return val2, nil
	default:
		return val2, nil
	}
}

func kindEqual(val reflect.Value, val2 reflect.Value) bool {
	return val.Kind() == val2.Kind()
}
