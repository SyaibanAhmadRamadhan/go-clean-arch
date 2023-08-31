package integration

//goland:noinspection ALL
import (
	"bytes"
	"context"
	"mime/multipart"
	"net/http/httptest"
	"testing"

	"github.com/SyaibanAhmadRamadhan/go-clean-arch/infrastructures/repository"
	"github.com/minio/minio-go/v7"
	"github.com/stretchr/testify/assert"
)

func Minio(t *testing.T) {
	fileContent := []byte("Contoh isi file")
	fileHeader := &multipart.FileHeader{
		Filename: "example.png",
		Size:     int64(len(fileContent)),
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", fileHeader.Filename)
	assert.NoError(t, err)
	_, _ = part.Write(fileContent)

	_ = writer.Close()

	req := httptest.NewRequest("POST", "/upload", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	file, fileHeader, err := req.FormFile("file")
	assert.NoError(t, err)
	defer func() {
		_ = file.Close()
	}()

	err = minioClient.MakeBucket(context.Background(), "files", minio.MakeBucketOptions{})
	assert.NoError(t, err)

	minioIMPL := repository.NewMinioImpl(minioClient)

	filename := minioIMPL.GenerateFileName(fileHeader, "user-images/public/", "")
	t.Run("SUCCESS_Upload", func(t *testing.T) {
		err = minioIMPL.UploadFile(context.Background(), fileHeader, filename, "files")
		assert.NoError(t, err)
	})

	t.Run("SUCCESS_Delete", func(t *testing.T) {
		err = minioIMPL.DeleteFile(context.Background(), filename, "files")
		assert.NoError(t, err)
	})
}
