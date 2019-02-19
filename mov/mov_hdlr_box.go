package mov

import (
	"log"
	"mp4info/comm"
)

type MovHdlrBox struct {
	MovBaseBox

	Version     int    //1 box版本，0或1，一般为0。（以下字节数均按version=0）
	Flags       int    //3
	PreDefined  int    //4
	HandlerType []byte //4 在media box中，该值为4个字符：vide”— video track “soun”— audio track “hint”— hint track
	Reserved    []byte //12
	Name        []byte //不定 track type name，以‘\0’结尾的字符串
}

func NewHdlrBox(head MovBaseBox)(hdlr *MovHdlrBox){
	hdlr = new(MovHdlrBox)
	hdlr.BoxSize = head.BoxSize
	hdlr.BoxType = head.BoxType

	return hdlr
}

func (hdlr* MovHdlrBox)Parse(buf []byte) (err error){
	hdlr.Version = int(buf[0])
	buf = buf[HDLR_VERSION_SIZE:]

	hdlr.Flags , err = comm.BytesToInt(buf[:HDLR_FLAGS_SIZE])
	if err != nil{
		return err
	}
	buf = buf[HDLR_FLAGS_SIZE:]

	hdlr.PreDefined , err = comm.BytesToInt(buf[:HDLR_PREDEFINED_SIZE])
	if err != nil{
		return err
	}
	buf = buf[HDLR_PREDEFINED_SIZE:]

	hdlr.HandlerType = buf[:HDLR_HANDLERTYPE_SIZE]
	buf = buf[HDLR_HANDLERTYPE_SIZE:]

	hdlr.Reserved = buf[:HDLR_RESERVED_SIZE]
	buf = buf[HDLR_RESERVED_SIZE:]

	hdlr.Name = buf

	hdlr.Show()
	return nil
}

func (hdlr* MovHdlrBox)Show(){
	log.Printf("			hdlr\n")
	log.Printf("				Version:%d\n" , hdlr.Version)
	log.Printf("				Flags:%d\n" , hdlr.Flags)
	log.Printf("				PreDefined:%d\n" , hdlr.PreDefined)
	log.Printf("				HandlerType:%s\n" , string(hdlr.HandlerType))
	log.Printf("				name:%s\n" , string(hdlr.Name))
}