package utils

import (
	"encoding/base64"
	"fmt"
	"unicode"
)

/**
 * @author       weimenghua
 * @time         2023-08-22 14:05
 * @description  编码解码
 * 参考资料：https://gobyexample-cn.github.io/base64-encoding
 */

func base64Example(data string) {
	fmt.Println("Original string: ", data)

	sEnc := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println("Base64 Encode string: ", sEnc)

	sDec, _ := base64.StdEncoding.DecodeString(sEnc)
	fmt.Println("Base64 Decode string: ", string(sDec))

	fmt.Println("\n")
}

func UnicodeExample(data string) {
	fmt.Println("Original string:", data)

	str := ChineseToUnicodeString(data)
	fmt.Println("chinese to unicode:", str)

	str2 := UnicodeToChineseString(str)
	fmt.Println("unicode to chinese:", str2)

	fmt.Println("\n")
}

// ChineseToUnicodeString 将中文字符串转换为 Unicode 字符串
func ChineseToUnicodeString(chineseStr string) string {
	var unicodeStr string
	for _, r := range chineseStr {
		if r > 127 {
			// 如果是中文字符,将其转换为 Unicode 字符
			unicodeStr += fmt.Sprintf("\\u%04X", r)
		} else {
			// 如果是非中文字符,则保留原样
			unicodeStr += string(r)
		}
	}
	return unicodeStr
}

// UnicodeToChineseString 将 Unicode 字符串转换为中文字符串 @todo 没有转换为中文
func UnicodeToChineseString(unicodeStr string) string {
	var chineseRunes []rune
	for _, r := range unicodeStr {
		if unicode.Is(unicode.Han, r) {
			// 如果是中文字符,添加到结果切片中
			chineseRunes = append(chineseRunes, r)
		} else {
			// 如果是非中文字符,则保留原样
			chineseRunes = append(chineseRunes, r)
		}
	}
	return string(chineseRunes)
}

func main() {
	base64Example("你好，世界")
	UnicodeExample("你好，世界")
}
