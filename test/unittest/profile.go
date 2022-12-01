package unittest

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"runtime/pprof"
)

// CpuProfile  工具形应用主要使用 runtime/pprof 库，将画像数据写入文件中。
func CpuProfile() {
	//CPU Profile
	f, err := os.Create("./cpuprofile")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	//Memory Profile
	fm, err := os.Create("./memoryprofile")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer fm.Close()
	_ = pprof.WriteHeapProfile(fm)

	//for i := 0; i < 100; i++ {
	//	fmt.Println("程序员麻辣烫")
	//}

}

//ReadFile 将文件整个读入内存
func ReadFile(path string) {
	file, _ := os.Open(path)
	defer file.Close()
	fs, _ := ioutil.ReadAll(file)
	fmt.Printf("%s\n", string(fs))
}

//ReadIoFile 🌟将文件整个读入内存，效率比较高，占用内存也最高。
func ReadIoFile(path string) {
	file, _ := ioutil.ReadFile(path)
	fmt.Printf("%s\n", string(file))
}

//checkFileIsExist 检查文件是否存在
func checkFileIsExist(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}

//BufferWFile TODO:if fileExist execute create else append
//	content:[]string
//	filename:AbsolutePath or RelativePath
//go:generate stringer -type=string
func BufferWFile(content []string, filename string) {
	var (
		f *os.File
		err1 error
	)
	if checkFileIsExist(filename) { //如果文件存在
		f, err1 = os.OpenFile(filename, os.O_APPEND |os.O_WRONLY, 0666) //打开文件
		fmt.Println("fileExistsAppend")
	} else {
		f, err1 = os.Create(filename) //创建文件
		fmt.Println("fileDoesNotExist")
	}
	defer f.Close()
	if err1 != nil {
		panic(err1)
	}
	w := bufio.NewWriter(f) //创建新的 Writer 对象

	for _,u := range content{
		w.WriteString(u)
	}

	w.Flush()
}
