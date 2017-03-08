package spack

import (
	"encoding/binary"
)

func Pack(str string) []byte {
	byteBuf := []byte(str)
	length := len(byteBuf)
	sendBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(sendBytes, uint32(length))
	sendBytes = append(sendBytes, byteBuf...)
	return sendBytes
}

type Buffer struct {
	buffer []byte
}

func NewBuffer() *Buffer {
	return &Buffer{
		buffer: []byte{},
	}
}

func (this *Buffer) Unpack(buf []byte) []string {
	this.buffer = append(this.buffer, buf...)
	return this.getMessages()
}

func (this *Buffer) getMessages() []string {
	result := []string{}
	for len(this.buffer) > 4 {
		length := binary.BigEndian.Uint32(this.buffer[0:4])
		readToPtr := length + 4
		if uint32(len(this.buffer)) < readToPtr {
			break
		}
		strBuf := this.buffer[4:readToPtr]
		this.buffer = this.buffer[readToPtr:]
		result = append(result, string(strBuf))
	}
	return result
}
