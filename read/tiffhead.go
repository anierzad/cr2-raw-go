package read

import "encoding/binary"

type tiffHeadReader struct {
	data *[]byte
}

func NewTiffHeadReader(data *[]byte) *tiffHeadReader {
	return &tiffHeadReader{
		data: data,
	}
}

func (thr tiffHeadReader) Endianness() string {
	bytes := (*thr.data)[:2]

	return string(bytes)
}

func (thr tiffHeadReader) MagicNumber() int {
	bytes := (*thr.data)[2:4]

	return int(binary.LittleEndian.Uint16(bytes))
}

func (thr tiffHeadReader) FirstIfdOffset() int {
	bytes := (*thr.data)[4:8]

	return int(binary.LittleEndian.Uint32(bytes))
}
