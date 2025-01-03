package upload

import (
	"io"
	"mime/multipart"
	"net/url"
	"os"
	"path"
	"strings"
	"time"

	"github.com/haierkeys/obsidian-image-api-gateway/global"
)

type FileType int

const TypeImage FileType = iota + 1

func GetFileName(name string) string {
	ext := GetFileExt(name)
	fileName := strings.TrimSuffix(name, ext)
	// fileName = util.EncodeMD5(fileName)

	return fileName + ext
}

func FileToMultipart(file *os.File) (multipart.File, *multipart.FileHeader, error) {

	// 将 *os.File 对象转换为 multipart.File 类型
	fileInfo, _ := file.Stat()
	return file, &multipart.FileHeader{
		Filename: fileInfo.Name(),
		Size:     fileInfo.Size(),
		// ModTime:  fileInfo.ModTime(),
		// 如果还需要其他属性，可以根据实际情况进行设置
	}, nil
}

func GetFileExt(name string) string {
	return path.Ext(name)
}

func GetSavePath() string {
	return global.Config.LocalFS.SavePath
}

func GetTempPath() string {
	return global.Config.App.TempPath
}

func GetSavePreDirPath() string {

	getYearMonth := time.Now().Format("200601")
	getDay := time.Now().Format("02")
	return getYearMonth + "/" + getDay + "/"
}

func UrlEscape(fileKey string) string {

	if strings.Contains(fileKey, "/") {

		i := strings.LastIndex(fileKey, "/")

		fileKey = fileKey[:i+1] + url.PathEscape(fileKey[i+1:])
	} else {
		fileKey = url.PathEscape(fileKey)
	}

	return fileKey
}

func GetServerUrl() string {
	return global.Config.App.UploadUrlPre
}

func CheckContainExt(t FileType, name string) bool {
	ext := GetFileExt(name)
	ext = strings.ToUpper(ext)
	switch t {
	case TypeImage:
		for _, allowExt := range global.Config.App.UploadAllowExts {
			if strings.ToUpper(allowExt) == ext {
				return true
			}
		}
	}
	return false
}

func CheckMaxSize(t FileType, f multipart.File) bool {
	content, _ := io.ReadAll(f)
	size := len(content)
	switch t {
	case TypeImage:
		if size >= global.Config.App.UploadMaxSize*1024*1024 {
			return true
		}
	}

	return false
}

func CheckPermission(dst string) bool {
	_, err := os.Stat(dst)

	return os.IsPermission(err)
}

func CheckPath(dst string) bool {
	_, err := os.Stat(dst)
	return os.IsNotExist(err)
}

func CreatePath(dst string, perm os.FileMode) error {
	err := os.MkdirAll(dst, perm)
	if err != nil {
		return err
	}

	return nil
}

func SaveFile(file multipart.File, dst string) error {

	err := os.MkdirAll(path.Dir(dst), os.ModePerm)
	if err != nil {
		return err
	}

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	file.Seek(0, 0)
	_, err = io.Copy(out, file)
	return err
}
