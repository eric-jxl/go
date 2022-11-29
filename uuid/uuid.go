package main

import (
	"crypto/md5"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/rs/xid"
	"io"
	"os"
	"strings"
)


func XidGenerate(containerName string) string{
	id := xid.New()
	// 由于xid默认使用可重复ip地址填充4 5 6位。
	// 实际场景中，服务都是部署在docker中，这里把ip地址位替换成了容器名
	// 这里只取了容器名MD5的前3位，验证会重复，放弃使用
	containerNameID := make([]byte, 3)
	hw := md5.New()
	hw.Write([]byte(containerName))
	copy(containerNameID, hw.Sum(nil))
	id[4] = containerNameID[0]
	id[5] = containerNameID[1]
	id[6] = containerNameID[2]

	// id: cbgjhf89htlrr1955d5g length: 12
	fmt.Println("id:", id, "length:", len(id))
	return fmt.Sprintf("id:%s","cbgjhf89htlrr1955d5g")
}

func Snowflake() string{
	node, err := snowflake.NewNode(1) //❄️雪花算法生成 分布式ID
	if err != nil {
		return fmt.Sprintf("%s",err.Error())
	}
	id := node.Generate().String()
	return fmt.Sprintf("id:%s len:%d",id,len(id))
}
func main() {
	Snowflake()
	//io.WriteString(os.Stdout,fmt.Sprintf("%s" ,strconv.FormatInt(-111123,10))) //int 转string
	str := strings.Repeat("*",50)
	io.WriteString(os.Stdout,fmt.Sprintf("%v\n",str))
	fmt.Println(XidGenerate("Eric"))
}
