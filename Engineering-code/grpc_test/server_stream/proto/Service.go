package proto

import (
	"strconv"
	"time"
)

type StreamServer struct {}

// ListValue 实现ListValue方法
func (ss *StreamServer)ListValue(req *StreamRequest, stream StreamServer_ListValueServer) error {
	for i := 0; i < 10; i++ {
		// 向流中发送消息， 默认每次send送消息最大长度为`math.MaxInt32`bytes
		err := stream.Send(&StreamResponse{
			StreamValue: req.Data + strconv.Itoa(i),
		})
		if err != nil {
			return err
		}
		time.Sleep(1 * time.Second)
	}
	return nil
}