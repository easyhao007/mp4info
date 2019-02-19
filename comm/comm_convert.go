package comm

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"log"
	"strconv"
)

func BytesToInt(buf []byte) (int, error) {
	if len(buf) == 0 {
		return -1, errors.New("invalid buf to convert int")
	} else if len(buf) == 3 {
		buf = append([]byte{0}, buf...)
	}

	bbuf := bytes.NewBuffer(buf)

	switch len(buf) {
	case 1:
		var tmp uint8
		err := binary.Read(bbuf, binary.BigEndian, &tmp)
		return int(tmp), err
	case 2:
		var tmp uint16
		err := binary.Read(bbuf, binary.BigEndian, &tmp)
		return int(tmp), err
	case 4:
		var tmp uint32
		err := binary.Read(bbuf, binary.BigEndian, &tmp)
		return int(tmp), err
	default:
		return 0, errors.New("bytes to int buf len is invalid")
	}
}

func BytesToFloat32(buf []byte) (float32, error) {
	if len(buf) == 0 {
		return -1, errors.New("invalid buf to convert int")
	}
	if len(buf) != 4{
		switch len(buf) {
		case 1:
			buf = append([]byte{0, 0, 0}, buf...)
		case 2:
			buf = append([]byte{0, 0}, buf...)
		case 3:
			buf = append([]byte{0}, buf...)
		default:
			return 0, errors.New("bytes to float32 buf len is invalid")
		}
	}

	bbuf := bytes.NewBuffer(buf)
	var tmp float32
	err := binary.Read(bbuf, binary.LittleEndian, &tmp)
	if err != nil {
		log.Println(err.Error())
		return 0.0, err
	}

	return tmp , nil
}

func BytesToFloat32Ex(bufa , bufb []byte ) (float32, error) {

	front ,err := BytesToInt(bufa)
	if err != nil{
		return 0.0 , err
	}
	end , err := BytesToInt(bufb)
	if err != nil{
		return 0.0 , err
	}
	s := fmt.Sprintf("%d.%d" , front , end)
	tmp , err := strconv.ParseFloat(s , 32)
	if err != nil{
		return 0.0, err
	}

	return float32(tmp) , nil
}

func BytesToFloat64(buf []byte) (float64, error) {
	if len(buf) == 0 {
		return -1, errors.New("invalid buf to convert int")
	}
	bbuf := bytes.NewBuffer(buf)
	var ret32 float64
	err := binary.Read(bbuf, binary.BigEndian, &ret32)
	if err != nil {
		log.Println(err.Error())
		return 0.0, err
	}
	return ret32, nil
}
