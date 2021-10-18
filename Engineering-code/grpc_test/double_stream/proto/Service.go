package proto

import (
	"io"
	"log"
	"strconv"
)

type BothStream struct{}

func (bs *BothStream) Conversations(stream Stream_ConversationsServer) error {
	//  用于计数请求
	n := 1

	//  开始读写数据
	for {
		//  读取数据，读取问题
		recv, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		//  写入数据，回答问题
		err = stream.Send(&StreamResponse{
			Answer: "from client stream , server anwser : the " + strconv.Itoa(n) + " question is " + recv.Question,
		})
		if err != nil {
			return err
		}

		n ++
		log.Printf("from stream client question is: %s", recv.Question)
	}

}