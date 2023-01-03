package models

import (
	"encoding/json"
	"time"
	"video/services/mq"

	"github.com/astaxie/beego/orm"
)

type Message struct {
	Id      int
	Content string
	AddTime int64
}

type MessageUser struct {
	Id        int
	MessageId int64
	AddTime   int64
	Status    int
	UserId    int
}

func init() {
	orm.RegisterModel(new(Message), new(MessageUser))
}

//保存通知消息
func SendMessageDo(content string) (int64, error) {
	o := orm.NewOrm()
	var message Message
	message.Content = content
	message.AddTime = time.Now().Unix()
	messageId, err := o.Insert(&message)
	return messageId, err
}

//保存消息接收人到队列中
func SendMessageUserMq(userId int, messageId int64) {
	//把数据转换成json字符串
	type Data struct {
		UserId    int
		MessageId int64
	}
	var data Data
	data.UserId = userId
	data.MessageId = messageId
	dataJson, _ := json.Marshal(data)
	mq.Publish("", "video_send_message_user", string(dataJson))
}
