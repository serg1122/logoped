package logoped

import (
	"encoding/json"
	"time"
)

type ISerializable interface {
	Serialize() []byte
}

type IReport interface{}

type Report struct {
	IReport
	ISerializable

	messageInfo IMessageInfo
}

func CreateReport(messageInfo IMessageInfo) *Report {

	return &Report{
		messageInfo: messageInfo,
	}
}

func (r Report) Serialize() []byte {

	return r.messageInfo.Serialize()
}

type IMessageInfo interface {
	ISerializable
}

type MessageInfo struct {
	bytes    []byte    `json:"message"`
	logLevel ILogLevel `json:"logLevel`
	time     time.Time `json:"time"`
}

func CreateMessageInfo(time time.Time, logLevel ILogLevel, bytes []byte) *MessageInfo {

	return &MessageInfo{
		bytes:    bytes,
		logLevel: logLevel,
		time:     time,
	}
}

func (m *MessageInfo) SerializeJSON() ([]byte, error) {

	return json.Marshal(struct {
		bytes    string    `json:"bytes"`
		logLevel ILogLevel `json:"logLevel"`
		time     string    `json:"time"`
	}{
		bytes:    string(m.bytes),
		logLevel: m.logLevel,
		time:     m.time.Format("2006-01-02 15:04:05.000000Z"),
	})
}

func (m *MessageInfo) Serialize() []byte {
	return []byte("bytes here...")
}
