package mov

import (
	"log"
	"mp4info/comm"
)

// media header box -> mdhd
type MovMdhdBox struct {
	// 基本头信息
	MovBaseBox

	Version         int // 1 版本信息
	Flags           int // 3
	CreationTime    int // 4 创建时间，（相对于UTC时间1904-01-01零点的秒数）
	ModificaionTime int // 4 修改时间
	TimeScale       int // 4 文件媒体在1秒内的刻度值，用duration和time_scale 值可以计算track时长
	Duration        int // 4 该track的时间长度，用duration和time scale值可以计算track时长
	Language        int // 2 媒体语言码。最高位为0，后面15位为3个字符（见ISO 639-2/T标准中定义）
	PreDefined      int // 2
}


func NewMdhdBox(head MovBaseBox)(mdhd *MovMdhdBox){
	mdhd = new(MovMdhdBox)
	mdhd.BoxType = head.BoxType
	mdhd.BoxSize = head.BoxSize
	return mdhd
}

func (mdhd *MovMdhdBox)Parse(buf []byte) (err error){
	mdhd.Version = int(buf[0])
	buf = buf[MDHD_VERSION_SIZE:]

	mdhd.Flags , err = comm.BytesToInt(buf[:MDHD_FLAGS_SIZE])
	if err != nil{
		return err
	}
	buf = buf[MDHD_FLAGS_SIZE:]

	mdhd.CreationTime , err = comm.BytesToInt(buf[:MDHD_CREATIONTIME_SIZE])
	if err != nil{
		return err
	}
	buf = buf[MDHD_CREATIONTIME_SIZE:]

	mdhd.ModificaionTime , err = comm.BytesToInt(buf[:MDHD_MODIFYTIME_SIZE])
	if err != nil{
		return err
	}
	buf = buf[MDHD_MODIFYTIME_SIZE:]

	mdhd.TimeScale , err = comm.BytesToInt(buf[:MDHD_TIMESCALE_SIZE])
	if err != nil{
		return err
	}
	buf = buf[MDHD_TIMESCALE_SIZE:]

	mdhd.Duration , err = comm.BytesToInt(buf[:MDHD_DURATION_SIZE])
	if err != nil{
		return err
	}
	buf = buf[MDHD_DURATION_SIZE:]

	mdhd.Language , err = comm.BytesToInt(buf[:MDHD_LANGUAGE_SIZE])
	if err != nil{
		return err
	}
	buf = buf[MDHD_LANGUAGE_SIZE:]

	mdhd.PreDefined , err = comm.BytesToInt(buf[:MDHD_PREDEFINED_SIZE])
	if err != nil{
		return err
	}
	buf = buf[MDHD_PREDEFINED_SIZE:]

	mdhd.Show()
	return nil
}

func (mdhd *MovMdhdBox)Show(){
	log.Printf("			mdhd\n")
	log.Printf("				Version:%d\n" , mdhd.Version)
	log.Printf("				Flags:%d\n" , mdhd.Flags)
	log.Printf("				CreationTime:%d\n" , mdhd.CreationTime)
	log.Printf("				ModificaionTime:%d\n" , mdhd.ModificaionTime)
	log.Printf("				TimeScale:%d\n" , mdhd.TimeScale)
	log.Printf("				Duration:%d\n" , mdhd.Duration)
	log.Printf("				Language:%d\n" , mdhd.Language)
	log.Printf("				PreDefined:%d\n" , mdhd.PreDefined)
}


