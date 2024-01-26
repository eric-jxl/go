package main

import (
	"fmt"

	"github.com/bwmarrin/snowflake"
	"github.com/rs/xid"
)

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
