package mov

import (
	"log"
	"mp4info/comm"
)

type MovMvhdBox struct {
	MovBaseBox

	Version         int     // 1 box版本，0/1，一般为0
	Flags           int     // 3
	CreationTime    int     // 4 创建时间（相对于UTC时间1904 - 01 - 01零点的秒数）
	ModificaionTime int     // 4 修改时间
	TimeScale       int     // 4 文件媒体在1秒内的刻度值，用duration和time_scale 值可以计算track时长
	Duration        int     // 4 该track的时间长度，用duration和time scale值可以计算track时长
	Rate            float32 // 4 推荐播放速率，高16位和低16位分别为小数点整数部分和小数部分，即[16.16] 格式.该值为1.0（0x00010000）表示正常前向播放
	Volume          float32 // 2 与rate类似，[8.8] 格式，1.0（0x0100）表示最大音量
	Reserved        string  // 10 保留位
	Matrix          string  // 36 视频变换矩阵
	PreDefined      string  // 24
	NextTrackID     int     // 4 下一个track使用的ID
}

func NewMvhdBox(head MovBaseBox) (mvhd *MovMvhdBox) {
	mvhd = new(MovMvhdBox)
	mvhd.BoxSize = head.BoxSize
	mvhd.BoxType = head.BoxType
	return mvhd
}

func (mvhd *MovMvhdBox) Parse(buf []byte) (err error) {
	log.Printf("	%s\n", mvhd.BoxType)
	mvhd.Version = int(buf[0])
	buf = buf[MVHD_VERSION_SIZE:]

	mvhd.Flags, err = comm.BytesToInt(buf[:MVHD_FLAGS_SIZE])
	if err != nil {
		return err
	}
	buf = buf[MVHD_FLAGS_SIZE:]

	mvhd.CreationTime, err = comm.BytesToInt(buf[:MVHD_CREATIONTIME_SIZE])
	if err != nil {
		return err
	}
	buf = buf[MVHD_CREATIONTIME_SIZE:]

	mvhd.ModificaionTime, err = comm.BytesToInt(buf[:MVHD_MODIFYTIME_SIZE])
	if err != nil {
		return err
	}
	buf = buf[MVHD_MODIFYTIME_SIZE:]

	mvhd.TimeScale, err = comm.BytesToInt(buf[:MVHD_TIMESCALE_SIZE])
	if err != nil {
		return err
	}
	buf = buf[MVHD_TIMESCALE_SIZE:]

	mvhd.Duration, err = comm.BytesToInt(buf[:MVHD_DURATION_SIZE])
	if err != nil {
		return err
	}
	buf = buf[MVHD_DURATION_SIZE:]

	mvhd.Rate, err = comm.BytesToFloat32Ex(buf[:MVHD_RATE_SIZE/2] , buf[MVHD_RATE_SIZE/2:MVHD_RATE_SIZE])
	if err != nil {
		return err
	}
	buf = buf[MVHD_RATE_SIZE:]

	mvhd.Volume, err = comm.BytesToFloat32Ex(buf[:MVHD_VOLUME_SIZE/2] , buf[MVHD_VOLUME_SIZE/2:MVHD_VOLUME_SIZE])
	if err != nil {
		return err
	}
	buf = buf[MVHD_VOLUME_SIZE:]

	//Reserved
	buf = buf[MVHD_RESERVED_SIZE:]

	mvhd.Matrix = string(buf[:MVHD_MATRIX_SIZE])
	buf = buf[MVHD_MATRIX_SIZE:]

	mvhd.PreDefined = string(buf[:MVHD_PREDEFINFO_SIZE])
	buf = buf[MVHD_PREDEFINFO_SIZE:]

	mvhd.NextTrackID, err = comm.BytesToInt(buf[:MVHD_NEXTTRACKID_SIZE])
	if err != nil {
		return err
	}

	mvhd.Show()
	return nil
}

func (mvhd *MovMvhdBox)Show(){
	log.Printf("		Version:%d\n" , mvhd.Version)
	log.Printf("		Flags:%d\n" , mvhd.Flags)
	log.Printf("		Creation Time:%d\n" , mvhd.CreationTime)
	log.Printf("		Modificaion Time:%d\n" , mvhd.ModificaionTime)
	log.Printf("		Time Scale:%d\n" , mvhd.TimeScale)
	log.Printf("		Duration:%d\n" , mvhd.Duration)
	log.Printf("		Rate:%f\n" , mvhd.Rate)
	log.Printf("		Volume:%f\n" , mvhd.Volume)
	log.Printf("		Matrix:%s\n" , mvhd.Matrix)
	log.Printf("		PreDefined:%s\n" , mvhd.PreDefined)
	log.Printf("		Next Track ID:%d\n" , mvhd.NextTrackID)
}