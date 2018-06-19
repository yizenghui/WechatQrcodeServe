package repository

import (
	"fmt"

	"github.com/chanxuehong/wechat.v2/mp/core"
	"github.com/chanxuehong/wechat.v2/mp/qrcode"
	"github.com/yizenghui/WechatQrcodeServe/orm"
)

//Task struct
type Task struct {
	ID     string
	Ticket string
	URL    string
}

//CreateTempQrcode 创建临时二维码
func CreateTempQrcode(id int32) (*qrcode.TempQrcode, error) {
	ats := core.NewDefaultAccessTokenServer("wx702b93aef72f3549", "8b69f45fc737a938cbaaffc05b192394", nil)
	clt := core.NewClient(ats, nil)
	return qrcode.CreateTempQrcode(clt, id, 7200)
}

// GetBindQrcode 站点获取签名任务
func GetBindQrcode(taskID int) (url string, err error) {
	qrcode, err := CreateTempQrcode(int32(taskID))
	if err != nil {
		return "", err
	}
	url = fmt.Sprintf("https://mp.weixin.qq.com/cgi-bin/showqrcode?ticket=%v", qrcode.Ticket)
	return url, nil
}

// NewQrcodeTask 获取签名任务二维码
func NewQrcodeTask() (task Task, err error) {
	var t orm.Task
	t.NewTask()

	qrcode, err := CreateTempQrcode(int32(t.ID))
	if err != nil {
		return task, err
	}
	task.URL = fmt.Sprintf("https://mp.weixin.qq.com/cgi-bin/showqrcode?ticket=%v", qrcode.Ticket)

	// 返回加密的任务ID
	id := int(t.ID)
	task.ID = Encode([]int{id})

	task.Ticket = qrcode.Ticket

	return task, nil
}

// CheckQrcodeTask 获取签名任务二维码
func CheckQrcodeTask(token string) (t orm.Task, err error) {
	ids := Decode(token)
	if id := ids[0]; id > 0 {
		t.GetTaskByID(id)
		return t, nil
	}
	return t, nil
}
