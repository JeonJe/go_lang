package main

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_uploadsHandler(t *testing.T) {
	
	// TODO: Add test cases.
	assert := assert.New(t)
	path := "/Users/premise/go/src/github.com/JeonJe/goTuckerweb/web4-1/uploads/test.png"
	file, _  := os.Open(path)
	defer file.Close()

	os.RemoveAll("./uploads")
	buf := &bytes.Buffer{}
	writer := multipart.NewWriter(buf)
	//form 파일을 만듬
	multi, err := writer.CreateFormFile("upload_file", filepath.Base(path))
	assert.NoError(err)

	io.Copy(multi, file)
	writer.Close()

	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/uploads", buf)
	//form data 형태인 것을 알려줘야 함 
	req.Header.Set("Content-Type", writer.FormDataContentType())
	uploadsHandler(res, req)
	assert.Equal(http.StatusOK, res.Code)

	uploadFilePath := "./uploads/" + filepath.Base(path)
	_, err_stat := os.Stat(uploadFilePath)
	assert.NoError(err_stat)

	uploadFile, _ := os.Open(uploadFilePath)
	originFile, _ := os.Open(path)

	defer uploadFile.Close()
	defer originFile.Close()
	uploadData := []byte{}
	originData := []byte{}

	uploadFile.Read(uploadData) 
	originFile.Read(originData)

	assert.Equal(uploadData, originData)		
}
