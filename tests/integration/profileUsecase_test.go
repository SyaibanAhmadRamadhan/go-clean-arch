package integration

import (
	"context"
	"github.com/SyaibanAhmadRamadhan/go-clean-arch/domain/dto"
	"testing"
	"time"

	"github.com/SyaibanAhmadRamadhan/go-clean-arch/internal/usecase"
	"github.com/stretchr/testify/assert"
)

func ProfileUsecase(t *testing.T) {
	timeOut := 2 * time.Second
	profile := usecase.NewProfileUsecaseImpl(ProfileRepo, UserRepo, timeOut)

	t.Run("SUCCESS_StoreProfile", func(t *testing.T) {
		profile, err := profile.StoreProfile(context.Background(), &dto.StoreProfileReq{
			UserID: userID_2,
		})
		assert.NoError(t, err)
		assert.NotNil(t, profile)
	})

	t.Run("SUCCESS_GetProfileByID", func(t *testing.T) {
		profile, err := profile.GetProfileByID(context.Background(), &dto.GetProfileReq{
			UserID: userID_2,
		})
		assert.NoError(t, err)
		assert.NotNil(t, profile)
	})

}
