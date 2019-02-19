package mov

import (
	"log"
	"mp4info/comm"
)

// ftyp 文件类型的描述
type MovFtypBox struct {
	MovBaseBox
	MajorBrand       string
	MinorVersion     int
	CompatibleBrands []string
}

// 申请一个ftyp类型的box
func NewFtypBox(head MovBaseBox) (ftyp *MovFtypBox) {
	ftyp = new(MovFtypBox)
	ftyp.BoxType = head.BoxType
	ftyp.BoxSize = head.BoxSize
	return ftyp
}

// 解析ftyp的box
func (ftyp *MovFtypBox) Parse(buf []byte) (err error) {
	log.Println("ftyp")
	ftyp.MajorBrand = string(buf[:FTYP_MAJORBRAND_SIZE])
	buf = buf[FTYP_MAJORBRAND_SIZE:]

	ftyp.MinorVersion, err = comm.BytesToInt(buf[:FTYP_MINORVERSION_SIZE])
	if err != nil {
		return err
	}
	buf = buf[FTYP_MINORVERSION_SIZE:]

	for {
		if len(buf) < FTYP_COMPATIBLE_BRANDS_SIZE {
			break
		}
		ftyp.CompatibleBrands = append(ftyp.CompatibleBrands, string(buf[:FTYP_COMPATIBLE_BRANDS_SIZE]))
		buf = buf[FTYP_COMPATIBLE_BRANDS_SIZE:]
	}

	ftyp.Show()
	return nil
}

func (ftyp *MovFtypBox)Show(){
	log.Printf("	Major Band:%s\n" , ftyp.MajorBrand)
	log.Printf("	Monor Version:%d\n" , ftyp.MinorVersion)
	log.Printf("	Compatible Brands:%s\n" , ftyp.CompatibleBrands)
 }
