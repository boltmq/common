package codec

import "hash/crc32"

// Crc32 获得循环冗余校验码
// Author: rongzhihong
// Since: 2017/10/26
func Crc32(p []byte) (int32, error) {
	ieee := crc32.NewIEEE()
	_, err := ieee.Write(p)
	if err != nil {
		return 0, err
	}
	crc := ieee.Sum32()
	return int32(crc), nil
}
