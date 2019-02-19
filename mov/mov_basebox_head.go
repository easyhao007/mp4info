package mov

import (
	"errors"
	"mp4info/comm"
)

// 基础的 box 头
type MovBaseBox struct {
	BoxSize int
	BoxType string
}

// 解析基础包头
func (basebox *MovBaseBox) Parse (buf []byte)(err error){
	if len(buf) < BASEBOX_HEAD_LEN {
		return errors.New("invalid buf to parse base box header")
	}

	basebox.BoxSize , err = comm.BytesToInt(buf[:BASEBOX_HEAD_SIZE_LEN])
	if err != nil{
		return err
	}

	basebox.BoxType = string(buf[BASEBOX_HEAD_SIZE_LEN:BASEBOX_HEAD_LEN])
	return nil
}

