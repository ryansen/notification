package services

import (
	"golang.org/x/net/context"
	"openpitrix.io/logger"
	notification "openpitrix.io/notification/pkg/pb"
	"openpitrix.io/notification/pkg/util/pbutil"
	"testing"
	"time"
)

func TestNewServer(t *testing.T) {
	logger.SetLevelByString("debug")
	InitGlobelSetting()
	s, _ := NewServer()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	s.SayHello(ctx, &notification.HelloRequest{Name: "unit_test2"})
}

func TestSayHello(t *testing.T) {
	InitGlobelSetting()
	s, _ := NewServer()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	s.SayHello(ctx, &notification.HelloRequest{Name: "unit_test2"})
}

func TestCreateNfWaddrs(t *testing.T) {
	InitGlobelSetting()
	s, _ := NewServer()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	testAddrsStr:="huojiao2006@163.com;huojiao2006@163.com"
	var req = &notification.CreateNfWaddrsRequest{
		NfPostType:   pbutil.ToProtoString("Information"),
		NotifyType:   pbutil.ToProtoString("Email"),
		AddrsStr:     pbutil.ToProtoString(testAddrsStr),
		Title:        pbutil.ToProtoString("Run case"),
		Content:      pbutil.ToProtoString("Run case content"),
		ShortContent: pbutil.ToProtoString("Run case ShortContent"),
		ExpiredDays:  pbutil.ToProtoString("7"),
		Owner:        pbutil.ToProtoString("HuoJiao"),
		Status:       pbutil.ToProtoString("New"),
	}
	s.CreateNfWaddrs(ctx, req)
}
