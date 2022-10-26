package gen

import (
	"errors"
	"fmt"
	"github.com/MenciusCheng/go-text-template/parse"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func GenByDirConfig(dirPath string) error {
	dirFileInfos, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return fmt.Errorf("ReadDir err: %w", err)
	}
	if len(dirFileInfos) == 0 {
		return errors.New("dir files is empty")
	}

	// 读取输入
	dataPath := fmt.Sprintf("%s/data.json", dirPath)
	dataMap, err := ReadJsonFileToMap(dataPath)
	if err != nil {
		return fmt.Errorf("ReadJsonFileToMap err: %w", err)
	}

	// 输出目录
	outputPath := fmt.Sprintf("%s/out", dirPath)
	outputDir, err := os.Stat(outputPath)
	if os.IsNotExist(err) {
		// os.MkdirAll("a/b/c/d", os.ModePerm)
		if err := os.Mkdir(outputPath, os.ModePerm); err != nil {
			return fmt.Errorf("os.Mkdir err: %w", err)
		}
	} else if err != nil {
		return fmt.Errorf("os.Stat err: %w", err)
	} else if !outputDir.IsDir() {
		return errors.New("outputDir is not dir\n")
	}

	// 根据模版目录生成文件
	tmplPath := fmt.Sprintf("%s/tmpl", dirPath)

	// 只生成当前目录
	//err = GenSubFile(dataMap, tmplPath, outputPath)
	//if err != nil {
	//	return fmt.Errorf("GenSubFile err: %w", err)
	//}

	// 递归生成子文件
	err = GenSubFile2(dataMap, tmplPath, outputPath)
	if err != nil {
		return fmt.Errorf("GenSubFile2 err: %w", err)
	}

	return nil
}

// 只生成当前目录配置
func GenSubFile(dataMap map[string]interface{}, tmplPath, outputPath string) error {
	tmplFiles, err := ioutil.ReadDir(tmplPath)
	if err != nil {
		fmt.Printf("ReadDir tmplPath err: %s\n", err)
		return err
	}
	if len(tmplFiles) == 0 {
		fmt.Printf("tmplFiles is empty\n")
		return err
	}
	for _, item := range tmplFiles {
		if item.IsDir() {
			continue
		}
		if err := GenFile(dataMap, tmplPath, item.Name(), outputPath); err != nil {
			fmt.Printf("GenFile err: %s\n", err)
			return err
		}
	}
	return nil
}

func GenFile(dataMap map[string]interface{}, tmplPath, tmplName, outputPath string) error {
	tmplFilePath := fmt.Sprintf("%s/%s", tmplPath, tmplName)
	tmplBytes, err := ioutil.ReadFile(tmplFilePath)
	if err != nil {
		return fmt.Errorf("ReadFile err: %w", err)
	}

	res, err := parse.JsonMap(string(tmplBytes), dataMap)
	if err != nil {
		return fmt.Errorf("parse.JsonMap err: %w", err)
	}

	outputName := tmplName
	if strings.HasSuffix(outputName, ".tmpl") {
		outputName = outputName[:len(outputName)-5]
	}
	outputFilePath := fmt.Sprintf("%s/%s", outputPath, outputName)

	if err := ioutil.WriteFile(outputFilePath, []byte(res), os.ModePerm); err != nil {
		return fmt.Errorf("WriteFile err: %w", err)
	}

	return nil
}

func ReadJsonFileToMap(filePath string) (map[string]interface{}, error) {
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("ReadFile err: %w", err)
	}

	sMap, err := parse.StringToMap(string(bytes))
	if err != nil {
		return nil, fmt.Errorf("StringToMap err: %w", err)
	}

	return sMap, nil
}

// 递归生成子文件
func GenSubFile2(dataMap map[string]interface{}, tmplPath, outputPath string) error {
	err := filepath.Walk(tmplPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("filepath.Walk handle err: %s", err)
			return err
		}
		// 跳过目录
		if info.IsDir() {
			return nil
		}

		var subDir string // 子目录
		if len(path) > len(tmplPath)+len(info.Name())+1 {
			subDir = path[len(tmplPath)+1 : len(path)-len(info.Name())-1]
		}
		//fmt.Printf("subDir: %s\n", subDir)

		if err := GenFile2(dataMap, tmplPath, subDir, info.Name(), outputPath); err != nil {
			fmt.Printf("GenFile err: %s\n", err)
			return err
		}
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk err: %s", err)
		return err
	}

	return nil
}

func GenFile2(dataMap map[string]interface{}, tmplPath, subDir, tmplName, outputPath string) error {
	tmplFilePath := fmt.Sprintf("%s/%s", tmplPath, tmplName)
	if subDir != "" {
		tmplFilePath = fmt.Sprintf("%s/%s/%s", tmplPath, subDir, tmplName)
	}

	tmplBytes, err := ioutil.ReadFile(tmplFilePath)
	if err != nil {
		return fmt.Errorf("ReadFile err: %w", err)
	}

	res, err := parse.JsonMap(string(tmplBytes), dataMap)
	if err != nil {
		return fmt.Errorf("parse.JsonMap err: %w", err)
	}

	outputName := tmplName
	if strings.HasSuffix(outputName, ".tmpl") {
		outputName = outputName[:len(outputName)-5]
	}
	outputFilePath := fmt.Sprintf("%s/%s", outputPath, outputName)
	if subDir != "" {
		if err := os.MkdirAll(fmt.Sprintf("%s/%s", outputPath, subDir), os.ModePerm); err != nil {
			fmt.Printf("os.Mkdir err: %s\n", err)
			return err
		}
		outputFilePath = fmt.Sprintf("%s/%s/%s", outputPath, subDir, outputName)
	}

	if err := ioutil.WriteFile(outputFilePath, []byte(res), os.ModePerm); err != nil {
		return fmt.Errorf("WriteFile err: %w", err)
	}

	return nil
}
