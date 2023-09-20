package pkg

import (
	"time"

	"github.com/bwmarrin/snowflake"
)

var node *snowflake.Node

/*
InitSnowflake

	startTime 时间因子~服务开启时间！
	machineID 机器 id 唯一节点~
*/
func InitSnowflake(startTime string, machineID int64) (err error) {
	var st time.Time
	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		return
	}
	snowflake.Epoch = st.UnixNano() / 1000000
	node, err = snowflake.NewNode(machineID)
	return
}
func GenSnowflakeID() int64 {
	return node.Generate().Int64()
}
