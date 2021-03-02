package read

import (
	"encoding/binary"
)

const entryLength int = 12

type ifdReader struct {
	offset int
	data *[]byte
}

func NewIfdReader(offset int, data *[]byte) *ifdReader {
	return &ifdReader{
		offset: offset,
		data: data,
	}
}

func (ir ifdReader) Count() int {
	start := ir.actualOffset(0)
	end := ir.actualOffset(2)

	bytes := (*ir.data)[start:end]

	return int(binary.LittleEndian.Uint16(bytes))
}

func (ir ifdReader) GetIfdEntries() (map[string]string, error) {
	entries := make(map[string]string)

	for i := 0; i < ir.Count(); i++ {
		offset := ir.actualOffset(2 + (entryLength * i))

		entry := newIfdEntryReader(offset, ir.data)

		tag, value, err := entry.GetEntryInfo()
		if err == nil {
			entries[tag] = value
		}
	}

	return entries, nil
}

func (ir ifdReader) NextIfdOffset() int {
	start := ir.actualOffset(2 + (ir.Count() * entryLength))
	end := ir.actualOffset(start + 4)

	bytes := (*ir.data)[start:end]

	return int(binary.LittleEndian.Uint32(bytes))
}

func (ir ifdReader) actualOffset(offset int) int {
	return offset + ir.offset
}
