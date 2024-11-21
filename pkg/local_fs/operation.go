package local_fs

import (
	"bytes"
	"errors"
	"io"
	"mime/multipart"
	"os"
	"path"
	"strings"

	"github.com/haierspi/golang-image-upload-service/global"
	"github.com/haierspi/golang-image-upload-service/pkg/upload"
)

type LocalFS struct {
	IsCheckSave bool
}

func (p LocalFS) CheckSave() error {

	savePath := p.getSavePath()

	if CheckPath(savePath) {
		if err := upload.CreatePath(savePath, os.ModePerm); err != nil {
			return errors.New("failed to create the save-path directory")
		}
	}
	if Permission(savePath) {
		return errors.New("no permission to upload the save path directory")
	}
	p.IsCheckSave = true
	return nil
}

func (p LocalFS) getSavePath() string {
	if strings.HasSuffix(global.Config.LocalFS.SavePath, "/") {
		return global.Config.LocalFS.SavePath
	} else {
		return global.Config.LocalFS.SavePath + "/"
	}
}

// SendByFile 上传文件
func (p *LocalFS) SendFile(fileKey string, file multipart.File, h *multipart.FileHeader) (string, error) {
	if !p.IsCheckSave {
		if err := p.CheckSave(); err != nil {
			return "", err
		}
	}

	dstFileKey := p.getSavePath() + fileKey

	err := os.MkdirAll(path.Dir(dstFileKey), os.ModePerm)
	if err != nil {
		return "", err
	}

	out, err := os.Create(dstFileKey)
	if err != nil {
		return "", err
	}
	defer out.Close()

	file.Seek(0, 0)
	_, err = io.Copy(out, file)
	if err != nil {
		return "", err
	} else {
		return dstFileKey, nil
	}
}

func (p *LocalFS) SendContent(fileKey string, content []byte) (string, error) {

	if !p.IsCheckSave {
		if err := p.CheckSave(); err != nil {
			return "", err
		}
	}

	dstFileKey := p.getSavePath() + fileKey

	out, err := os.Create(dstFileKey)
	if err != nil {
		return "", err
	}
	defer out.Close()

	_, err = io.Copy(out, bytes.NewReader(content))
	if err != nil {
		return "", err
	} else {
		return dstFileKey, nil
	}
}
