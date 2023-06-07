package generate_method

import (
	"log"
	"os"
	"strings"
	"unicode"
)

//去掉下划线，  首字母不大写 .   demo: extYidunConfig
func RemoveUnderscoreLowercase(str string) string {
	var (
		strs []string
		str2 string
	)

	strs = strings.Split(str, "_")

	for index, value := range strs {
		if index == 0 {
			str2 += value
		} else {
			str2 += string(unicode.ToUpper(rune(value[0]))) + value[1:]
		}
	}
	return str2
}

//去掉下划线，  首字母大写 . ExtYidunConfig
func RemoveUnderscoreCapitalize(str string) string {
	var (
		strs []string
		str2 string
	)
	strs = strings.Split(str, "_")
	for _, value := range strs {
		str2 += string(unicode.ToUpper(rune(value[0]))) + value[1:]
	}
	return str2
}

//
func writeFile(dir string, filename string, byteSlice []byte, isCover bool) {
	var (
		fil = dir + filename
	)

	log.Println(`file : `, fil)

	if !isCover {
		fileInfo, _ := os.Stat(fil)
		//if err != nil {
		//	log.Fatal(err)
		//	return
		//}

		if fileInfo != nil {
			return
		}

		//if fileInfo.Size() != 0 {
		//	log.Println(`file exist , return `)
		//	return
		//}
	}

	file, err := os.Create(fil)
	if err != nil {
		log.Println(err)
		return
	}

	defer file.Close()
	_, err = file.Write(byteSlice)
	if err != nil {
		log.Println(err)
		return
	}
	//log.Printf("Wrote %d bytes.\n", bytesWritten)

}
