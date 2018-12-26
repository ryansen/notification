package task

import (
	"openpitrix.io/logger"
	"openpitrix.io/notification/pkg/config"
	"strconv"
)

type Handler interface {
	ExtractTasks() error
	HandleTask(handlerNum string) error
	ServeTask() error
}


type handler struct {
	tasksc Service
}

func NewHandler(tasksc Service) Handler {
	return &handler{
		tasksc: tasksc,
	}
}

func (h *handler) ExtractTasks() error {
	logger.Infof(nil, "ExtractTasks Starts....")
	h.tasksc.ExtractTasks()
	return nil
}

func (h *handler) HandleTask(handlerNum string) error {
	logger.Infof(nil, "*******HandleTask Starts,Num："+handlerNum+"....")
	err:=h.tasksc.HandleTask(handlerNum)
	if err != nil {
		logger.Warnf(nil, "%+v", err)
		return  err
	}
	return nil
}

func (h *handler) ServeTask() error {
	logger.Infof(nil, "Call handlerImpl.ServeTask")
	go h.ExtractTasks()

	MaxWorkingTasks:=config.GetInstance().App.Maxtasks

	for i := 0; i < MaxWorkingTasks; i++ {
		go h.HandleTask(strconv.Itoa(i))
	}
	return nil
}