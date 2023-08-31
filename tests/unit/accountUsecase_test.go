package unit

import (
	"bytes"
	"context"
	"database/sql"
	"github.com/SyaibanAhmadRamadhan/go-clean-arch/internal/helpers"
	"mime/multipart"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/SyaibanAhmadRamadhan/go-clean-arch/domain/dto"
	"github.com/SyaibanAhmadRamadhan/go-clean-arch/domain/mocks"
	"github.com/SyaibanAhmadRamadhan/go-clean-arch/domain/model"
	"github.com/SyaibanAhmadRamadhan/go-clean-arch/internal/usecase"
	"github.com/stretchr/testify/assert"
)

func multipartFileHeader() *multipart.FileHeader {
	fileContent := []byte("Contoh isi file")
	fileHeader := &multipart.FileHeader{
		Filename: "example.png",
		Size:     int64(len(fileContent)),
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", fileHeader.Filename)
	if err != nil {
		panic(err)
	}
	_, _ = part.Write(fileContent)
	_ = writer.Close()

	req := httptest.NewRequest("POST", "/upload", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	file, fileHeader, err := req.FormFile("file")
	if err != nil {
		panic(err)
	}
	defer func(file multipart.File) {
		_ = file.Close()
	}(file)
	return fileHeader
}

func TestAccountUpdateUsecase(t *testing.T) {
	uow := &mocks.FakeUnitOfWork{}
	profileRepoMock := &mocks.FakeProfileRepo{}
	userRepoMock := &mocks.FakeUserRepo{}
	minioRepoMock := &mocks.FakeMinioRepo{}
	timeOutCtx := 3 * time.Second
	ctx := context.Background()
	image := "user-images/public/asd.png"
	accountUsecase := usecase.NewAccountUsecaseImpl(profileRepoMock, userRepoMock, minioRepoMock, timeOutCtx)

	profile := model.Profile{
		ProfileID: "profileid_1",
		UserID:    "userid_1",
		Quote:     sql.NullString{String: ""},
		CreatedAt: time.Now().Unix(),
		CreatedBy: "profileid_1",
		UpdatedAt: time.Now().Unix(),
		UpdatedBy: sql.NullString{},
		DeletedAt: sql.NullInt64{},
		DeletedBy: sql.NullString{},
	}

	user := model.User{
		ID:              "userid_1",
		FullName:        "rama_1",
		Gender:          "undefinied",
		Image:           "default-male.png",
		Username:        "ibanrmaa_1",
		Email:           "1_ibanrama29@gmail.com",
		Password:        "123456",
		PhoneNumber:     sql.NullString{},
		EmailVerifiedAt: true,
		CreatedAt:       time.Now().Unix(),
		CreatedBy:       "userid_1",
		UpdatedAt:       time.Now().Unix(),
		UpdatedBy:       sql.NullString{},
		DeletedAt:       sql.NullInt64{},
		DeletedBy:       sql.NullString{},
	}

	req := dto.UpdateAccountReq{
		UserID:      "userid_1",
		FullName:    "rama_update_1",
		Gender:      "male",
		Image:       multipartFileHeader(),
		PhoneNumber: "1234567890",
		Quote:       "semangat_update_1",
	}

	profileRepoMock.GetProfileByID(ctx, "profileid_1")
	profileRepoMock.GetProfileByIDReturns(profile, nil)

	userRepoMock.GetUserByID(ctx, "userid_1")
	userRepoMock.GetUserByIDReturns(user, nil)

	profileConv, userConv := helpers.UpdateAccountToModel(&req, user.Image)
	profileRepoMock.UpdateProfile(ctx, profileConv)
	profileRepoMock.UpdateProfileReturns(profile, nil)

	uow.GetTx()
	uow.GetTxReturns(&sql.Tx{}, nil)

	userRepoMock.UpdateUser(ctx, userConv)
	userRepoMock.UpdateUserReturns(user, nil)

	minioRepoMock.GenerateFileName(multipartFileHeader(), "user-images/public/", "")
	minioRepoMock.GenerateFileNameReturns(image)

	minioRepoMock.UploadFile(ctx, multipartFileHeader(), image, "files")
	minioRepoMock.UploadFileReturns(nil)

	uow.EndTx(nil)
	uow.EndTxReturns(nil)

	profileRes, userRes, err := accountUsecase.UpdateAccount(ctx, &req)
	assert.NoError(t, err)
	assert.NotNil(t, profileRes)
	assert.NotNil(t, userRes)
}

func TestAccounUpdateWithDeleteFileUsecase(t *testing.T) {
	profileRepoMock := &mocks.FakeProfileRepo{}
	userRepoMock := &mocks.FakeUserRepo{}
	minioRepoMock := &mocks.FakeMinioRepo{}
	timeOutCtx := 3 * time.Second
	ctx := context.Background()
	image := "/files/user-images/public/asd.png"

	accountUsecase := usecase.NewAccountUsecaseImpl(profileRepoMock, userRepoMock, minioRepoMock, timeOutCtx)

	profile := model.Profile{
		ProfileID: "profileid_1",
		UserID:    "userid_1",
		Quote:     sql.NullString{String: ""},
		CreatedAt: time.Now().Unix(),
		CreatedBy: "profileid_1",
		UpdatedAt: time.Now().Unix(),
		UpdatedBy: sql.NullString{},
		DeletedAt: sql.NullInt64{},
		DeletedBy: sql.NullString{},
	}

	user := model.User{
		ID:              "userid_1",
		FullName:        "rama_1",
		Gender:          "undefinied",
		Image:           image,
		Username:        "ibanrmaa_1",
		Email:           "1_ibanrama29@gmail.com",
		Password:        "123456",
		PhoneNumber:     sql.NullString{},
		EmailVerifiedAt: true,
		CreatedAt:       time.Now().Unix(),
		CreatedBy:       "userid_1",
		UpdatedAt:       time.Now().Unix(),
		UpdatedBy:       sql.NullString{},
		DeletedAt:       sql.NullInt64{},
		DeletedBy:       sql.NullString{},
	}

	req := dto.UpdateAccountReq{
		ProfileID:   "profileid_1",
		UserID:      "userid_1",
		FullName:    "rama_update_1",
		Gender:      "male",
		Image:       multipartFileHeader(),
		PhoneNumber: "1234567890",
		Quote:       "semangat_update_1",
	}

	profileRepoMock.GetProfileByID(ctx, "profileid_1")
	profileRepoMock.GetProfileByIDReturns(profile, nil)

	userRepoMock.GetUserByID(ctx, "userid_1")
	userRepoMock.GetUserByIDReturns(user, nil)

	profileConv, userConv := helpers.UpdateAccountToModel(&req, user.Image)
	profileRepoMock.UpdateProfile(ctx, profileConv)
	profileRepoMock.UpdateProfileReturns(profile, nil)

	userRepoMock.UpdateUser(ctx, userConv)
	userRepoMock.UpdateUserReturns(user, nil)

	minioRepoMock.GenerateFileName(multipartFileHeader(), "user-images/public/", "")
	minioRepoMock.GenerateFileNameReturns("user-images/public/asd.png")

	minioRepoMock.UploadFile(ctx, multipartFileHeader(), image, "files")
	minioRepoMock.UploadFileReturns(nil)

	minioRepoMock.DeleteFile(ctx, image, "files")
	minioRepoMock.DeleteFileReturns(nil)

	profileRes, userRes, err := accountUsecase.UpdateAccount(ctx, &req)
	assert.NoError(t, err)
	assert.NotNil(t, profileRes)
	assert.NotNil(t, userRes)
}
