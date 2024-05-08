package utils

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"

	"log"
	"os"
)

// Take an existing jpg srcFileName and decode/encode it
func CreateJpg() {

	srcFileName := "flower.jpg"
	dstFileName := "newFlower.jpg"
	// Decode the JPEG data. If reading from file, create a reader with
	reader, err := os.Open(srcFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()
	//Decode from reader to image format
	m, formatString, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Got format String", formatString)
	fmt.Println(m.Bounds())

	//Encode from image format to writer
	f, err := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		log.Fatal(err)
		return
	}

	err = jpeg.Encode(f, m, &jpeg.Options{Quality: 75})
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Jpg file", dstFileName, "created")

}

// Take an existing png srcFileName and decode/encode it
func CreatePng() {
	srcFileName := "mouse.png"
	dstFileName := "newMouse.png"
	reader, err := os.Open(srcFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()
	//Decode from reader to image format
	m, formatString, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Got format String", formatString)
	fmt.Println(m.Bounds())

	//Encode from image format to writer
	f, err := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		log.Fatal(err)
		return
	}

	err = png.Encode(f, m)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Png file", dstFileName, "created")

}

// Gets base64 string of an existing JPEG file
func GetJPEGbase64(fileName string) string {

	imgFile, err := os.Open(fileName)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer imgFile.Close()

	// create a new buffer base on file size
	fInfo, _ := imgFile.Stat()
	var size = fInfo.Size()
	buf := make([]byte, size)

	// read file content into buffer
	fReader := bufio.NewReader(imgFile)
	fReader.Read(buf)

	imgBase64Str := base64.StdEncoding.EncodeToString(buf)
	//fmt.Println("Base64 string is:", imgBase64Str)
	return imgBase64Str

}
