package mov

import "log"

type MovMoovBox struct{
	MovBaseBox
}

func NewMoovBox(head MovBaseBox)(moov* MovMoovBox){
	moov = new(MovMoovBox)
	moov.BoxType = head.BoxType
	moov.BoxSize = head.BoxSize
	return moov
}

// 解析moov box
func (moov* MovMoovBox)Parse(buf []byte) (err error){
	log.Printf("%s\n" , moov.BoxType)

	for{
		boxLen , err := ParseBox(buf)
		if err != nil{
			return err
		}
		buf = buf[boxLen:]
		if len(buf) == 0{
			break
		}
	}
	return nil
}
