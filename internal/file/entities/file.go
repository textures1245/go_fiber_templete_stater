package entities

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"strings"

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
	hasher := sha256.New()
	hasher.Write([]byte(fileData))
	hash := hex.EncodeToString(hasher.Sum(nil))

	pngFilename := "public/image/" + hash + ".png"

	if _, err := os.Stat(pngFilename); os.IsNotExist(err) {
		reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(fileData))
		m, _, err := image.Decode(reader)
		if err != nil {
			return "", err
		}
		// bounds := m.Bounds()
		// fmt.Println(bounds, formatString)

		osFile, errOnOpenFIle := os.OpenFile(pngFilename, os.O_WRONLY|os.O_CREATE, 0777)
		if errOnOpenFIle != nil {
			return "", err
		}
		err = png.Encode(osFile, m)
		if err != nil {
			return "", err
		}
		buffer := new(bytes.Buffer)
		errWhileEncoding := png.Encode(buffer, m) // img is your image.Image
		if errWhileEncoding != nil {
			return "", err
		}
		srcImg := fmt.Sprintf("data:image/png;base64,%s", base64.StdEncoding.EncodeToString(buffer.Bytes()))
		log.Info("Create new PNG file name: ", pngFilename, "as the output")

		return srcImg, nil
	}

	data, err := os.ReadFile(pngFilename)
	if err != nil {
		return "", err
	}

	srcImg := "data:image/png;base64," + base64.StdEncoding.EncodeToString(data)
	log.Info("Reusing exist PNG file name: ", pngFilename, "as the output")

	return srcImg, nil

}

func (f *File) Base64toJpg() (string, error) {

	if len(f.FileData) == 0 || f.FileType != "JPG" {
		return "", errors.New("Invalid file data or file type, expected PNG file type but got " + f.FileType)
	}

	fileData := base64.StdEncoding.EncodeToString(f.FileData)
	hasher := sha256.New()
	hasher.Write([]byte(fileData))
	hash := hex.EncodeToString(hasher.Sum(nil))

	jpgFilename := "public/image/" + hash + ".jpg"

	if _, err := os.Stat(jpgFilename); os.IsNotExist(err) {
		reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(fileData))
		m, formatString, err := image.Decode(reader)
		if err != nil {
			return "", err
		}
		bounds := m.Bounds()
		fmt.Println("base64toJpg", bounds, formatString)

		osFile, err := os.OpenFile(jpgFilename, os.O_WRONLY|os.O_CREATE, 0777)
		if err != nil {
			return "", err
		}

		err = jpeg.Encode(osFile, m, &jpeg.Options{Quality: 75})
		if err != nil {
			return "", err
		}

		buffer := new(bytes.Buffer)
		errWhileEncoding := jpeg.Encode(buffer, m, nil) // img is your image.Image
		if errWhileEncoding != nil {
			log.Fatal(errWhileEncoding)
		}
		srcImg := fmt.Sprintf("data:image/jpeg;base64,%s", base64.StdEncoding.EncodeToString(buffer.Bytes()))
		log.Info("Create new JPG file name: ", jpgFilename, "as the output")

		return srcImg, nil
	}

	data, err := os.ReadFile(jpgFilename)
	if err != nil {
		return "", err
	}

	srcImg := "data:image/jpeg;base64," + base64.StdEncoding.EncodeToString(data)
	log.Info("Reusing exist JPG file name: ", jpgFilename, "as the output")

	return srcImg, nil
}

func (f *File) Base64toFile() (string, error) {
	if len(f.FileData) == 0 || f.FileType != "PDF" {
		return "", errors.New("Invalid file data or file type, expected PDF file type but got " + f.FileType)
	}

	// encode blob to string
	fileData := base64.StdEncoding.EncodeToString(f.FileData)
	hasher := sha256.New()
	hasher.Write([]byte(fileData))
	hash := hex.EncodeToString(hasher.Sum(nil))

	fileName := "public/file/" + hash + ".pdf"
	if _, err := os.Stat(fileName); os.IsNotExist(err) {

		data, err := base64.StdEncoding.DecodeString(fileData)
		if err != nil {
			return "", err
		}

		err = os.WriteFile(fileName, data, 0644)
		if err != nil {
			return "", err
		}

		srcFile := fmt.Sprintf("data:file/%s;base64,%s", strings.ToLower(f.FileType), base64.StdEncoding.EncodeToString(data))
		log.Info("Reusing exist ", f.FileType, " file name: ", fileName, "as the output")

		return srcFile, nil
	}

	data, err := os.ReadFile(fileName)
	if err != nil {
		return "", err
	}

	srcFile := fmt.Sprintf("data:file/%s;base64,%s", strings.ToLower(f.FileType), base64.StdEncoding.EncodeToString(data))
	log.Info("Reusing exist ", f.FileType, " file name: ", fileName, "as the output")

	return srcFile, nil

}
