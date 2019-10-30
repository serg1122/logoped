package logoped

import (
	"os"
	"time"
)

type IReport interface{}

type Report struct {
	systemInfo      ISysInfo
	contextInfo     IContextInfo
	messageInfo     IMessageInfo
	postmessageInfo IPostMessage
}

func CreateReport(
	systemInfo ISysInfo,
	contextInfo IContextInfo,
	messageInfo IMessageInfo,
	postmessageInfo IPostMessage,
) *Report {

	return &Report{
		systemInfo:      systemInfo,
		contextInfo:     contextInfo,
		messageInfo:     messageInfo,
		postmessageInfo: postmessageInfo,
	}
}

type ISysInfo interface {
	GetTime() time.Time
	GetGid() int
	GetPid() int
	GetPpid() int
}

type SysInfo struct {
	time time.Time
	gid  int
	pid  int
	ppod int
}

func CreateSysInfo() *SysInfo {
	return &SysInfo{
		time: time.Now(),
		gid:  os.Getgid(),
		pid:  os.Getpid(),
		ppid: os.Getppid(),
	}
}
