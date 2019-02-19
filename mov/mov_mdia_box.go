package mov

import "log"

type MovMdiaBox struct{
	MovBaseBox
}

func NewMdiaBox(head MovBaseBox) (mdia *MovMdiaBox){
	mdia = new(MovMdiaBox)
	mdia.BoxSize = head.BoxSize
	mdia.BoxType = head.BoxType

	return mdia
}

func (mdia *MovMdiaBox)Parse(buf []byte) error{
	log.Println("		mdia")
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

