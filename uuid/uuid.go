package main

import (
	"fmt"

	"github.com/bwmarrin/snowflake"
	"github.com/rs/xid"
)

//func XidGenerate(containerName string) string {
//	id := xid.New()
//	// 由于xid默认使用可重复ip地址填充4 5 6位。
//	// 实际场景中，服务都是部署在docker中，这里把ip地址位替换成了容器名
//	// 这里只取了容器名MD5的前3位，验证会重复，放弃使用
//	containerNameID := make([]byte, 3)
//	hw := md5.New()
//	hw.Write([]byte(containerName))
//	copy(containerNameID, hw.Sum(nil))
//	id[4] = containerNameID[0]
//	id[5] = containerNameID[1]
//	id[6] = containerNameID[2]
//
//	return fmt.Sprintf("id:%s", "cbgjhf89htlrr1955d5g")
//}

func GenerateXid(containerName string) string {
	generator := xid.New()
	id := generator.String()
	return fmt.Sprintf("id: %s", id)
}

func Snowflake() string {
	node, err := snowflake.NewNode(1) //❄️雪花算法生成 分布式ID
	if err != nil {
		return fmt.Sprintf("%s", err.Error())
	}
	id := node.Generate().String()
	return fmt.Sprintf("id:%s len:%d", id, len(id))
}
func main() {
	Snowflake()
	//_, err := io.WriteString(os.Stdout, fmt.Sprintf("%s\n", strconv.FormatInt(-111123, 10)))
	//str := strings.Repeat("*", 50)
	//_, err = io.WriteString(os.Stdout, fmt.Sprintf("%v\n", str))

	fmt.Println(GenerateXid("Eric"))
}
