package entities

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2/log"
)

type File struct {
	Id        int64  `db:"id"`
	FileName  string `db:"file_name"`
	FileData  []byte `db:"file_data"`
	FileType  string `db:"file_type"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}

func (f *File) Base64toPng() (string, error) {

	if len(f.FileData) == 0 || f.FileType != "PNG" {
		return "", errors.New("Invalid file data or file type, expected PNG file type but got " + f.FileType)
	}

	fileData := base64.StdEncoding.EncodeToString(f.FileData)
	log.Info("File data", fileData)

	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(fileData))
	m, formatString, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	bounds := m.Bounds()
	fmt.Println(bounds, formatString)

	//Encode from image format to writer
	timestamp := time.Now().Format("20060102150405")
	pngFilename := "public/image/img_" + timestamp + ".png"

	osFile, errOnOpenFIle := os.OpenFile(pngFilename, os.O_WRONLY|os.O_CREATE, 0777)
	if errOnOpenFIle != nil {
		log.Fatal(errOnOpenFIle)
	}

	err = png.Encode(osFile, m)
	if err != nil {
		log.Fatal(err)
	}

	buffer := new(bytes.Buffer)
	errWhileEncoding := png.Encode(buffer, m) // img is your image.Image
	if errWhileEncoding != nil {
		log.Fatal(errWhileEncoding)
	}
	srcImg := fmt.Sprintf("data:image/png;base64,%s", base64.StdEncoding.EncodeToString(buffer.Bytes()))
	log.Info("Png file", pngFilename, "created")

	return srcImg, nil

}

func (f *File) Base64toJpg() (string, error) {

	if len(f.FileData) == 0 || f.FileType != "JPG" {
		return "", errors.New("Invalid file data or file type, expected PNG file type but got " + f.FileType)
	}

	fileData := base64.StdEncoding.EncodeToString(f.FileData)

	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(fileData))
	m, formatString, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	bounds := m.Bounds()
	fmt.Println("base64toJpg", bounds, formatString)

	//Encode from image format to writer
	timestamp := time.Now().Format("20060102150405")
	jpgFilename := "public/image/img_" + timestamp + ".jpg"

	osFile, err := os.OpenFile(jpgFilename, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		log.Fatal(err)
	}

	err = jpeg.Encode(osFile, m, &jpeg.Options{Quality: 75})
	if err != nil {
		log.Fatal(err)
	}

	buffer := new(bytes.Buffer)
	errWhileEncoding := jpeg.Encode(buffer, m, nil) // img is your image.Image
	if errWhileEncoding != nil {
		log.Fatal(errWhileEncoding)
	}
	srcImg := fmt.Sprintf("data:image/jpeg;base64,%s", base64.StdEncoding.EncodeToString(buffer.Bytes()))
	log.Info("Png file", jpgFilename, "created")

	return srcImg, nil

}
