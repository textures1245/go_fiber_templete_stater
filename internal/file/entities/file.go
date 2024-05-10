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
	"net/http"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/textures1245/go-template/pkg/apperror"
)

type File struct {
	Id        int64  `db:"id"`
	FileName  string `db:"file_name"`
	FileData  []byte `db:"file_data"`
	FileType  string `db:"file_type"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}

func (f *File) Base64toPng(c *fiber.Ctx) (*string, *string, error) {

	if len(f.FileData) == 0 || f.FileType != "PNG" {
		return nil, nil, errors.New("Invalid file data or file type, expected PNG file type but got " + f.FileType)
	}

	fileData := base64.StdEncoding.EncodeToString(f.FileData)
	hasher := sha256.New()
	hasher.Write([]byte(fileData))
	hash := hex.EncodeToString(hasher.Sum(nil))

	path := "public/image/"
	pngFilename := path + hash + ".png"

	if _, err := os.Stat(pngFilename); os.IsNotExist(err) {
		reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(fileData))
		m, _, err := image.Decode(reader)
		if err != nil {
			return nil, nil, err
		}
		// bounds := m.Bounds()
		// fmt.Println(bounds, formatString)

		osFile, errOnOpenFIle := os.OpenFile(pngFilename, os.O_WRONLY|os.O_CREATE, 0777)
		if errOnOpenFIle != nil {
			return nil, nil, err
		}
		err = png.Encode(osFile, m)
		if err != nil {
			return nil, nil, err
		}
		buffer := new(bytes.Buffer)
		errWhileEncoding := png.Encode(buffer, m) // img is your image.Image
		if errWhileEncoding != nil {
			return nil, nil, err
		}
		base64url := fmt.Sprintf("data:image/png;base64,%s", base64.StdEncoding.EncodeToString(buffer.Bytes()))
		filePathData := fmt.Sprintf("%s/%s", c.Hostname(), pngFilename)
		log.Info("Create new PNG file name: ", pngFilename, "as the output")

		return &base64url, &filePathData, nil
	}

	data, err := os.ReadFile(pngFilename)
	if err != nil {
		return nil, nil, err
	}

	base64url := "data:image/png;base64," + base64.StdEncoding.EncodeToString(data)
	filePathData := fmt.Sprintf("%s/%s", c.Hostname(), pngFilename)
	log.Info("Reusing exist PNG file name: ", pngFilename, "as the output")

	return &base64url, &filePathData, nil

}

func (f *File) Base64toJpg(c *fiber.Ctx) (*string, *string, error) {

	if len(f.FileData) == 0 || f.FileType != "JPG" {
		return nil, nil, errors.New("Invalid file data or file type, expected PNG file type but got " + f.FileType)
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
			return nil, nil, err
		}
		bounds := m.Bounds()
		fmt.Println("base64toJpg", bounds, formatString)

		osFile, err := os.OpenFile(jpgFilename, os.O_WRONLY|os.O_CREATE, 0777)
		if err != nil {
			return nil, nil, err
		}

		err = jpeg.Encode(osFile, m, &jpeg.Options{Quality: 75})
		if err != nil {
			return nil, nil, err
		}

		buffer := new(bytes.Buffer)
		errWhileEncoding := jpeg.Encode(buffer, m, nil) // img is your image.Image
		if errWhileEncoding != nil {
			log.Fatal(errWhileEncoding)
		}
		base64url := fmt.Sprintf("data:image/jpeg;base64,%s", base64.StdEncoding.EncodeToString(buffer.Bytes()))
		filePathData := fmt.Sprintf("%s/%s", c.Hostname(), jpgFilename)
		log.Info("Create new JPG file name: ", jpgFilename, "as the output")

		return &base64url, &filePathData, nil
	}

	data, err := os.ReadFile(jpgFilename)

	if err != nil {
		return nil, nil, err
	}

	base64url := "data:image/jpeg;base64," + base64.StdEncoding.EncodeToString(data)
	filePathData := fmt.Sprintf("%s/%s", c.Hostname(), jpgFilename)

	log.Info("Reusing exist JPG file name: ", jpgFilename, "as the output")

	return &base64url, &filePathData, nil
}

func (f *File) Base64toFile(c *fiber.Ctx, includeDomain bool) (*string, *string, error) {
	if len(f.FileData) == 0 || f.FileType != "PDF" {
		return nil, nil, errors.New("Invalid file data or file type, expected PDF file type but got " + f.FileType)
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
			return nil, nil, err
		}

		err = os.WriteFile(fileName, data, 0644)
		if err != nil {
			return nil, nil, err
		}

		srcFile := fmt.Sprintf("data:file/%s;base64,%s", strings.ToLower(f.FileType), base64.StdEncoding.EncodeToString(data))

		log.Info("Reusing exist ", f.FileType, " file name: ", fileName, "as the output")

		filePathData := fileName
		if includeDomain {
			filePathData = fmt.Sprintf("%s/%s", c.Hostname(), fileName)
		}
		return &srcFile, &filePathData, nil
	}

	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, nil, err
	}

	srcFile := fmt.Sprintf("data:file/%s;base64,%s", strings.ToLower(f.FileType), base64.StdEncoding.EncodeToString(data))
	filePathData := fileName
	if includeDomain {
		filePathData = fmt.Sprintf("%s/%s", c.Hostname(), fileName)
	}
	log.Info("Reusing exist ", f.FileType, " file name: ", fileName, "as the output")

	return &srcFile, &filePathData, nil

}

func (file *File) DecodeBlobToFile(c *fiber.Ctx, domainIncludeOnFile bool) (*string, *string, int, *apperror.CErr) {
	var (
		base64urlRes string
		fPathDatRes  string
	)
	switch file.FileType {
	case "PNG":
		base64url, fPathDat, err := file.Base64toPng(c)
		if err != nil {
			status, cErr := apperror.CustomSqlExecuteHandler("File", err)
			return nil, nil, status, cErr
		}
		base64urlRes = *base64url
		fPathDatRes = *fPathDat
	case "JPG":
		base64url, fPathDat, err := file.Base64toJpg(c)
		if err != nil {
			status, cErr := apperror.CustomSqlExecuteHandler("File", err)
			return nil, nil, status, cErr
		}
		base64urlRes = *base64url
		fPathDatRes = *fPathDat
	case "PDF":
		base64url, fPathDat, err := file.Base64toFile(c, domainIncludeOnFile)
		if err != nil {
			status, cErr := apperror.CustomSqlExecuteHandler("File", err)
			return nil, nil, status, cErr
		}
		base64urlRes = *base64url
		fPathDatRes = *fPathDat
	default:
		return nil, nil, http.StatusBadRequest, apperror.NewCErr(errors.New("Only except for PNG and JPG for now"), errors.ErrUnsupported)
	}

	return &base64urlRes, &fPathDatRes, http.StatusOK, nil
}

// func working() {
// 	files, err := os.ReadDir(path)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// Determine the number of goroutines to use
// 	numGoroutines := 2

// 	// Calculate the size of each group
// 	groupSize := (len(files) + numGoroutines - 1) / numGoroutines

// 	var wg sync.WaitGroup

// 	// Start a goroutine for each group
// 	for i := 0; i < len(files); i += groupSize {
// 		end := i + groupSize
// 		if end > len(files) {
// 			end = len(files)
// 		}

// 		wg.Add(1)
// 		go func(files []os.DirEntry) {
// 			defer wg.Done()
// 			doJob := func(f []os.DirEntry) (*string, *string, error) {
// 				for _, file := range f {
// 					_, err := os.Stat(file.Name())
// 					if os.IsNotExist(err) {
// 						reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(fileData))
// 						m, _, err := image.Decode(reader)
// 						if err != nil {
// 							return nil, nil, err
// 						}
// 						// bounds := m.Bounds()
// 						// fmt.Println(bounds, formatString)

// 						osFile, errOnOpenFIle := os.OpenFile(pngFilename, os.O_WRONLY|os.O_CREATE, 0777)
// 						if errOnOpenFIle != nil {
// 							return nil, nil, err
// 						}
// 						err = png.Encode(osFile, m)
// 						if err != nil {
// 							return nil, nil, err
// 						}
// 						buffer := new(bytes.Buffer)
// 						errWhileEncoding := png.Encode(buffer, m) // img is your image.Image
// 						if errWhileEncoding != nil {
// 							return nil, nil, err
// 						}
// 						base64url := fmt.Sprintf("data:image/png;base64,%s", base64.StdEncoding.EncodeToString(buffer.Bytes()))
// 						filePathData := fmt.Sprintf("%s/%s", c.Hostname(), pngFilename)
// 						log.Info("Create new PNG file name: ", pngFilename, "as the output")

// 						return &base64url, &filePathData, nil
// 					}
// 				}
// 				return nil, nil, nil
// 			}
// 			base64url, filePathData, err := doJob(files)
// 		}(files[i:end])
// 	}
// 	wg.Wait()
// }
