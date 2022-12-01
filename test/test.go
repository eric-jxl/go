package main

/*
	int sum(int a, int b) { return a+b; };
	int abs(int a){
		if (a <0) return -a ;
		else return a;
	};
*/
import "C"
import (
	Ca "Workspace/gocache"
	unit "Workspace/test/unittest"
	"crypto/md5"
	"fmt"
	"io"
	"sync"
	"time"
)

const (
	name = "Eric"
	// DefaultExpiration _    = iota
	//KB   = 1 << (10 * iota)
	//MB   = 1 << (10 * iota)
	//GB   = 1 << (10 * iota)
	//TB   = 1 << (10 * iota)
	//PB   = 1 << (10 * iota)
	//x    = iota // x == 0
	//y    = iota // y == 1
	//z    = iota // z == 2
	//w
	DefaultExpiration time.Duration = 0
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Addr string `json:"addr,omitempty"`
}

//New returns a new hash.Hash computing the MD5 checksum. The Hash also
//implements encoding.BinaryMarshaler and encoding.BinaryUnmarshaler to
//marshal and unmarshal the internal state of the hash.
func md5encrypt() {
	//假设用户名 abc，密码 123456
	h := md5.New()
	var pwd string
	fmt.Println("输入密码:")
	_, err := fmt.Scanln(&pwd)
	if err != nil {
		return
	}
	_, err = io.WriteString(h, pwd)
	if err != nil {
		return
	}

	// pwmd5 等于 e10adc3949ba59abbe56e057f20f883e
	pwmd5 := fmt.Sprintf("%x", h.Sum(nil))

	// 指定两个 salt： salt1 = @#$%   salt2 = ^&*()
	salt1 := "@#$%"
	salt2 := "^&*()"

	// salt1 + 用户名 + salt2 + MD5 拼接
	_, _ = io.WriteString(h, salt1)
	_, _ = io.WriteString(h, "abc")
	_, _ = io.WriteString(h, salt2)
	_, _ = io.WriteString(h, pwmd5)

	last := fmt.Sprintf("%x", h.Sum(nil))
	fmt.Println(last)
}

func setCache() {
	tc := Ca.New(DefaultExpiration, 1)
	s := struct {
		name string
		age  int
	}{"nike", 11}
	tc.Set("a", s, 0)
	a, ok := tc.Get("a")
	if ok {
		fmt.Println(a)
	} else {
		fmt.Println("Key is not Exists!")
	}
	if err := tc.Increment("tint8", 2); err != nil {
		e := fmt.Errorf("implementError:(%s) ", err)
		fmt.Println(e.Error())
	}
}

func routine() {
	wg := sync.WaitGroup{}
	var mutex sync.Mutex
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			mutex.Lock()
			//fmt.Printf("第%d次 Locked\n", i)
			setCache()
			//time.Sleep(1 * time.Second)
			//fmt.Print("UnLock\n")
			mutex.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
}
func producer(c chan<- int) {
	for {
		time.Sleep(1 * time.Second)
		c <- 1
	}
}

func consumer(c <-chan int) {
	for {
		fmt.Println(<-c)
	}
}

func testCase() {
	c := make(chan int, 10)
	go producer(c)
	go consumer(c)
	select {}
}

func fibTest(a ...interface{}) {
	start := time.Now()
	result := unit.Fib(10)
	end := time.Now()
	timeDiff := end.Sub(start).Seconds()
	fmt.Printf("斐切拉波数列第40项的值为：%d,用时:%.2f \n", result, timeDiff)
	h := time.Now()
	unit.SyncMux()
	fmt.Printf("APi调用时长%.2f\n", time.Since(h).Seconds())
}

func main() {

	//c := [...]int{4, 5, 6}
	//d := `hello
	//world`
	//fmt.Printf("%v,%v,%s,%d", name, c, d, len(c))
	//
	//slice := []byte{'a', 'b', 'c', 'd'}
	//fmt.Println(len(slice),string(slice))

	//fmt.Println(KB);fmt.Println(MB);fmt.Println(GB);fmt.Println(TB);fmt.Println(PB)
	//md5encrypt()

	//go routine()
	//testCase()
	//md5encrypt()

	//fmt.Println(C.sum(1, 2))
	//fmt.Println(C.abs(-2))
	//fibTest()
	//pt := new(Person)
	////s := make(chan int,1)
	////s <- 1
	//str := make([]string, 0, 10)
	//str = append(str, "22222", "2355", name)
	//fmt.Printf("%+v\n", str)
	//ptr := unsafe.Sizeof(str)
	//pts := unsafe.Alignof(str)
	//fmt.Println(ptr, pts)
	//pt.Name, pt.Age = name, 17
	//
	//obstruct := make([]string, len(str), 10)
	//copy(obstruct, str)
	//fmt.Println(reflect.DeepEqual(obstruct, str))
	//fmt.Println(obstruct)
	//fmt.Println(unit.ConnectMysql())
	//unit.CpuProfile()
	//unit.BufferWFile("Eric\n","a.txt")
	var ct = []string{"AAAAA\n","AAAAA\n","AAAAA\n"}
	unit.BufferWFile(ct,"a.txt")
}
