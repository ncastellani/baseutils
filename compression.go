package baseutils

import (
	"bytes"
	"compress/gzip"
	"io"
)

// GzipData
// take an content as a bytes pointer and Gzip those data, returning the output as bytes
func GzipData(input *[]byte) (output *[]byte, err error) {

	// generate a byte buffer to print out the compressed data
	var b bytes.Buffer
	gz := gzip.NewWriter(&b)

	// write the Gzip file
	_, err = gz.Write(*input)
	if err != nil {
		return
	}

	if err = gz.Flush(); err != nil {
		return
	}

	if err = gz.Close(); err != nil {
		return
	}

	// get the bytes of the buffer
	compressedData := b.Bytes()

	return &compressedData, err
}

// UnGzipData
// take a Gzip-ped content as a bytes pointer and undo the Gzip, returning the uncompressed data as bytes
func UnGzipData(input *[]byte) (output *[]byte, err error) {

	// gzip-read the bytes data as a buffer
	b := bytes.NewBuffer(*input)

	var r io.Reader
	r, err = gzip.NewReader(b)
	if err != nil {
		return
	}

	var resB bytes.Buffer
	_, err = resB.ReadFrom(r)
	if err != nil {
		return
	}

	// get the ungzipped data
	result := resB.Bytes()

	return &result, err
}
