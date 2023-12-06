package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"reflect"
	"regexp"
	"strings"
)

var regex = `(\s|'|;|"|--)+(union|sleep|alter|insert|drop|truncate|update|from|grant|exec|where|select|and|or|count|chr|mid|like|limit)\s+`

func Waf() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		flag, data := checkWafParams(ctx)
		if flag {
			ctx.AbortWithStatus(403)
		} else {
			ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))
			ctx.Next()
		}
	}
}

func checkWafParams(ctx *gin.Context) (bool, []byte) {
	data, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("in checkWafParams, api read param error:", err)
	}
	rs := strings.Split(regex, "\n")
	return check(data, rs)
}

// 检测每个参数是否都匹配正则，只有当body是json时才继续判断
func check(data []byte, rs []string) (bool, []byte) {
	v := make(map[string]interface{})
	if err := json.Unmarshal(data, &v); err != nil || len(v) == 0 {
		return false, data
	}
	for _, r := range rs {
		if flag := checkValue(reflect.ValueOf(v), r); flag {
			return true, data
		}
	}
	return false, data
}

func checkValue(v reflect.Value, r string) bool {
	var flag bool
	switch v.Kind() {
	case reflect.Interface:
		return checkValue(reflect.ValueOf(v.Interface()), r)
	case reflect.Array:
		for i := 0; i < v.Len(); i++ {
			flag = flag || checkValue(v.Index(i), r)
		}
		return flag
	case reflect.Map:
		for _, key := range v.MapKeys() {
			flag = flag || checkValue(v.MapIndex(key), r)
		}
		return flag
	case reflect.String:
		return match(v.Interface().(string), r)
	default:
		return flag
	}
}

func match(v string, r string) bool {
	v = strings.ToLower(v)
	flag, _ := regexp.MatchString(r, v)
	if flag {
		log.Println("--------------------------------match:", v, r)
	}
	return flag
}
