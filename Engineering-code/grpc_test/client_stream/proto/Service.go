package proto

import (
	"io"
	"log"
)

type StreamClient struct {}

// server 端 实现 service 中的方法 RouteList
func (sc StreamClient) RouteList(stream StreamClient_RouteListServer) error  {
	//  服务端，拿到客户端提供的 数据流， 可以读取其中的数据
	for {
		// 从数据流中读取数据
		res, err := stream.Recv()
		if err == io.EOF {
			//  如果 读取完成，则发送 响应 后 ，关闭 数据流
			return stream.SendAndClose(&SimpleResponse{
				Status: 200,
				Value: "OK",
			})
		} else if err != nil {
			return err
		}

		//  打印从数据流中读取的数据
		log.Println(res.StreamData)
	}
}