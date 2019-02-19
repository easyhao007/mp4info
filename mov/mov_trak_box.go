package mov

import "log"

type MovTrakBox struct {
	MovBaseBox
}

func NewTrakBox(head MovBaseBox) (trak *MovTrakBox) {
	trak = new(MovTrakBox)
	trak.BoxType = head.BoxType
	trak.BoxSize = head.BoxSize

	return trak
}

func (trak *MovTrakBox)Parse(buf []byte) error{
	log.Printf("	trak\n")

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
