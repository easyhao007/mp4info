package mov

import "log"

// 解析box的整体调用
func ParseBox(buf []byte)(boxLen int , err error){
	var baseBoxHead MovBaseBox
	err = baseBoxHead.Parse(buf)
	if err != nil{
		return 0 , err
	}
	buf = buf[BASEBOX_HEAD_LEN:baseBoxHead.BoxSize]
	return baseBoxHead.BoxSize , ParseBoxAttribute(baseBoxHead , buf)
}



// 根据类型解析box的属性，并dump信息
func ParseBoxAttribute(head MovBaseBox , buf []byte) (err error){
	switch head.BoxType {
	case "ftyp":
		ftyp := NewFtypBox(head)
		err = ftyp.Parse(buf)
		return err
	case "moov":
		moov := NewMoovBox(head)
		err = moov.Parse(buf)
		return err
	case "mvhd":
		mvhd := NewMvhdBox(head)
		err = mvhd.Parse(buf)
		return err
	case "iods":
		log.Println("	iods")
	case "trak":
		trak := NewTrakBox(head)
		err = trak.Parse(buf)
		return err
	case "tkhd":
		tkhd := NewTkhdBox(head)
		err = tkhd.Parse(buf)
		return err
	case "mdia":
		mdia := NewMdiaBox(head)
		err = mdia.Parse(buf)
		return err
	case "mdhd":
		mdhd := NewMdhdBox(head)
		err = mdhd.Parse(buf)
		return err
	case "hdlr":
		hdlr := NewHdlrBox(head)
		err = hdlr.Parse(buf)
		return err
	case "edts":
		log.Println("		edts")

	}

	return err
}
