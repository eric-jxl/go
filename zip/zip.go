/*
	用Go语言编写的自动压缩文件工具，将指定目录下的所有文件压缩成一个以"Eric_"为前缀的zip文件

	An automatic file compression tool written in Go language that compresses all files in a specified directory
	into a zip file with the prefix "Eric_
 */
package zip

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// CreateZip
// sourceDir :指定要压缩的目录
// targetFile:指定压缩后的文件名
func CreateZip(sourceDir,targetFile string) {

	// 创建 zip 文件
	zipFile, err := os.Create(targetFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer zipFile.Close()

	// 创建 zip writer
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// 遍历目录，并添加文件到 zip
	filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			// 打开文件
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()

			// 添加文件到 zip
			header, err := zip.FileInfoHeader(info)
			if err != nil {
				return err
			}
			header.Name = "Eric_" + header.Name // 文件名加上前缀
			writer, err := zipWriter.CreateHeader(header)
			if err != nil {
				return err
			}
			_, err = io.Copy(writer, file)
			if err != nil {
				return err
			}
		}
		return nil
	})
}
