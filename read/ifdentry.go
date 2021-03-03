package read

import (
	"encoding/binary"
	"fmt"
	"strconv"
)

var (
	tagIds = map[int]string{
		271: "Exif.Image.Make",
		272: "Exif.Image.Model",
		256: "Exif.Image.ImageWidth",
		257: "Exif.Image.ImageLength",
		40962: "Exif.Photo.PixelXDimension",
		40963: "Exif.Photo.PixelYDimension",
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

	// string.
	case 2:
		start := ier.getTagValue()
		end := start + ier.getTagCount()
		bytes := (*ier.data)[start:end]

		infoValue = string(bytes)

	// ushort, ulong
	case 3, 4, 8, 9:
		infoValue = strconv.Itoa(ier.getTagValue())

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
