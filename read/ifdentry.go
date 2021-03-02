package read

import (
	"encoding/binary"
	"fmt"
)

var (
	tagIds = map[int]string{
		271: "Exif.Image.Make",
		272: "Exif.Image.Model",
	}
)

type ifdEntryReader struct {
	offset int
	data *[]byte
	entryData *[]byte
}

func newIfdEntryReader(offset int, data *[]byte) *ifdEntryReader {

	entryData := (*data)[offset:offset + entryLength]

	ier := &ifdEntryReader{
		offset: offset,
		data: data,
		entryData: &entryData,
	}

	return ier
}

func (ier ifdEntryReader) GetEntryInfo() (string, string, error) {

	var infoName, infoValue string
	tagId := ier.getTagId()
	tagType := ier.getTagType()

	// Try retrieve tag name.
	value, found := tagIds[tagId]
	if !found {
		return infoName,
			infoValue,
			fmt.Errorf("Tag %d is not supported yet.", tagId)
	}
	infoName = value

	// Try retrieve tag value.
	switch tagType {

	// String.
	case 2:
		start := ier.getTagValue()
		end := start + ier.getTagCount()
		bytes := (*ier.data)[start:end]

		infoValue = string(bytes)
	default:
		return infoName,
			infoValue,
			fmt.Errorf("Type %d is not supported yet.", tagType)
	}

	return infoName,
		infoValue,
		nil
}

func (ier ifdEntryReader) getTagId() int {
	return int(binary.LittleEndian.Uint16((*ier.entryData)[:2]))
}

func (ier ifdEntryReader) getTagType() int {
	return int(binary.LittleEndian.Uint16((*ier.entryData)[2:4]))
}

func (ier ifdEntryReader) getTagCount() int {
	return int(binary.LittleEndian.Uint32((*ier.entryData)[4:8]))
}

func (ier ifdEntryReader) getTagValue() int {
	return int(binary.LittleEndian.Uint32((*ier.entryData)[8:]))
}

func (ier ifdEntryReader) actualOffset(offset int) int {
	return offset + ier.offset
}
