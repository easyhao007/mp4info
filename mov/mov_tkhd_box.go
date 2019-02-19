package mov

import (
	"log"
	"mp4info/comm"
)

type MovTkhdBox struct {
	MovBaseBox
	Version uint8 // 1 box版本，0或1，一般为0。（以下字节数均按version = 0）
	//按位或操作结果值，预定义如下;
	//	0x000001 track_enabled，否则该track不被播放；
	//	0x000002 track_in_movie，表示该track在播放中被引用；
	//	0x000004 track_in_preview，表示该track在预览时被引用。
	//	一般该值为7，如果一个媒体所有track均未设置track_in_movie和track_in_preview,将被理解为所有track均设置了这两项;
	//	对于hint track，该值为0;
	Flags            int     //3
	CreationTime     int     //4 创建时间（相对于UTC时间1904 - 01 - 01零点的秒数）
	ModificationTime int     //4 修改时间
	TrackID          int     //4 id号 不能重复且不能为0
	Reserved1        []byte  //4 保留位
	Duration         int     //4 track的时间长度
	Reserved2        []byte  //8 保留位
	Layer            int     //2 视频层，默认为0，值小的在上层
	AlternateGroup   int     //2 track分组信息，默认为0表示该track未与其他track有群组关系
	Volume           float32 //2 [8.8] 格式，如果为音频track，1.0（0x0100）表示最大音量；否则为0
	Reserved3        []byte  //2 保留位
	Matrix           []byte  //36 视频变换矩阵
	Width            float32 //4 宽
	Height           float32 //4 高，均为[16.16] 格式值 与sample描述中的实际画面大小比值，用于播放时的展示宽高
}

// 新建一个tkdh box
func NewTkhdBox(head MovBaseBox) (tkhd *MovTkhdBox) {
	tkhd = new(MovTkhdBox)
	tkhd.BoxSize = head.BoxSize
	tkhd.BoxType = head.BoxType

	return tkhd
}

func (tkhd *MovTkhdBox) Parse(buf []byte) (err error) {
	log.Println("		tkhd")

	tkhd.Version = uint8(buf[0])
	buf = buf[TKHD_VERSION_SIZE:]

	tkhd.Flags, err = comm.BytesToInt(buf[:TKHD_FLAGS_SIZE])
	if err != nil {
		return err
	}
	buf = buf[TKHD_FLAGS_SIZE:]

	tkhd.CreationTime, err = comm.BytesToInt(buf[:TKHD_CREATIONTIME_SIZE])
	if err != nil {
		return err
	}
	buf = buf[TKHD_CREATIONTIME_SIZE:]

	tkhd.ModificationTime , err = comm.BytesToInt(buf[:TKHD_MODIFYTIME_SIZE])
	if err != nil{
		return err
	}
	buf = buf[TKHD_MODIFYTIME_SIZE:]

	tkhd.TrackID , err = comm.BytesToInt(buf[:MVHD_TRACKID_SIZE])
	if err != nil{
		return err
	}
	buf = buf[MVHD_TRACKID_SIZE:]

	tkhd.Reserved1 = buf[:TKHD_RESERVED1_SIZE]
	buf = buf[TKHD_RESERVED1_SIZE:]

	tkhd.Duration , err = comm.BytesToInt(buf[:MVHD_DURATION_SIZE])
	if err != nil{
		return err
	}
	buf = buf[MVHD_DURATION_SIZE:]

	tkhd.Reserved2 = buf[:TKHD_RESERVED2_SIZE]
	buf = buf[TKHD_RESERVED2_SIZE:]

	tkhd.Layer , err = comm.BytesToInt(buf[:TKHD_LAYER_SIZE])
	if err != nil{
		return err
	}
	buf = buf[TKHD_LAYER_SIZE:]

	tkhd.AlternateGroup , err = comm.BytesToInt(buf[:TKHD_ALTERGROUP_SIZE])
	if err != nil{
		return err
	}
	buf = buf[TKHD_ALTERGROUP_SIZE:]

	tkhd.Volume , err = comm.BytesToFloat32Ex(buf[:TKHD_VOLUME_SIZE/2] , buf[TKHD_VOLUME_SIZE/2:TKHD_VOLUME_SIZE])
	if err != nil{
		return err
	}
	buf = buf[TKHD_VOLUME_SIZE:]

	tkhd.Reserved3 = buf[:TKHD_RESERVED3_SIZE]
	buf = buf[TKHD_RESERVED3_SIZE:]

	tkhd.Matrix = buf[:TKHD_MATRIX_SIZE]
	buf = buf[TKHD_MATRIX_SIZE:]

	tkhd.Width , err = comm.BytesToFloat32Ex(buf[:TKHD_WIDTH_SIZE/2] , buf[TKHD_WIDTH_SIZE/2:TKHD_WIDTH_SIZE])
	if err != nil{
		return err
	}
	buf = buf[TKHD_WIDTH_SIZE:]

	tkhd.Height , err = comm.BytesToFloat32Ex(buf[:TKHD_HEIGHT_SIZE/2] , buf[TKHD_HEIGHT_SIZE/2:TKHD_HEIGHT_SIZE])
	if err != nil{
		return err
	}
	buf = buf[TKHD_HEIGHT_SIZE:]

	tkhd.Show()
	return nil
}

func (tkhd* MovTkhdBox)Show(){
	log.Printf("			Version:%d\n" , tkhd.Version)
	log.Printf("			Flags:%d\n" , tkhd.Flags)
	log.Printf("			Creation Time:%d\n" , tkhd.CreationTime)
	log.Printf("			Modification Time:%d\n" , tkhd.ModificationTime)
	log.Printf("			Track ID:%d\n" , tkhd.TrackID)
	log.Printf("			Duration:%d\n" , tkhd.Duration)
	log.Printf("			Layer:%d\n" , tkhd.Layer)
	log.Printf("			Alternate Group:%d\n" , tkhd.AlternateGroup)
	log.Printf("			Volume:%f\n" , tkhd.Volume)
	log.Printf("			Width:%f\n" , tkhd.Width)
	log.Printf("			Height:%f\n" , tkhd.Height)
}