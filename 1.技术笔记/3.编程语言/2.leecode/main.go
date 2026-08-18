package main

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"sort"
	"strings"
)

func generateSign(
	params map[string]string,
	appSecret string,
	signMethod string,
) (string, error) {
	// 1. 对参数名排序，排除 sign
	keys := make([]string, 0, len(params))

	for key := range params {
		if key != "sign" {
			keys = append(keys, key)
		}
	}

	sort.Strings(keys)

	// 2. 按“参数名 + 参数值”依次拼接
	var builder strings.Builder

	for _, key := range keys {
		builder.WriteString(key)
		builder.WriteString(params[key])
	}

	content := builder.String()
	var result []byte

	// 3. 根据 sign_method 计算摘要
	switch signMethod {
	case "md5":
		sum := md5.Sum([]byte(appSecret + content + appSecret))
		result = sum[:]

	case "hmac":
		mac := hmac.New(md5.New, []byte(appSecret))
		_, _ = mac.Write([]byte(content))
		result = mac.Sum(nil)

	case "hmac-sha256":
		mac := hmac.New(sha256.New, []byte(appSecret))
		_, _ = mac.Write([]byte(content))
		result = mac.Sum(nil)

	default:
		return "", fmt.Errorf("unsupported sign method: %s", signMethod)
	}
	fmt.Printf("content=%q\n", content)

	// 4. 转成大写十六进制
	return strings.ToUpper(hex.EncodeToString(result)), nil
}

func main() {
	params := map[string]string{
		"app_key":             "35326009",
		"format":              "json",
		"method":              "taobao.pc.sem.item.query",
		"sign_method":         "hmac",
		"timestamp":           "2026-08-06 11:09:00",
		"v":                   "2.0",
		"sem_product_request": `{"source":"baidu_ai","keyword":"男装"}`,
	}

	appSecret := "092040d60ddeedff07c94c6356f3e00a"

	sign, err := generateSign(params, appSecret, params["sign_method"])
	if err != nil {
		panic(err)
	}
	fmt.Printf("sign=%s\n", sign)
	fmt.Printf("sign length=%d\n", len(sign))

	params["sign"] = sign
	fmt.Println(sign)
}
