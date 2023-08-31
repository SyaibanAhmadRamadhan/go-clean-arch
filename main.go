package main

import (
	"github.com/SyaibanAhmadRamadhan/go-clean-arch/delivery/restapi"
	cusmiddleware "github.com/SyaibanAhmadRamadhan/go-clean-arch/delivery/restapi/middleware"
	"github.com/SyaibanAhmadRamadhan/go-clean-arch/infrastructures/config"
	"github.com/SyaibanAhmadRamadhan/go-clean-arch/infrastructures/repository"
	"github.com/SyaibanAhmadRamadhan/go-clean-arch/internal/usecase"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog/log"
	"net/http"
	"time"
)

func main() {
	config.LogInit()
	config.EnvInit()
	pgConn := config.NewPgConn()
	defer func() {
		err := pgConn.Close()
		if err != nil {
			log.Err(err).Msg("ERROR CLOSE DB")
		}
	}()

	redisConn := config.NewRedisConn()
	defer func() {
		err := redisConn.Client.Close()
		if err != nil {
			log.Err(err).Msg("ERROR CLOSE REDIS CONN")
		}
	}()

	minioConn, err := config.NewMinioConn(config.MinIoEndpoint, config.MinIoID, config.MinIoSecretKey, config.MinIoSSL)
	if err != nil {
		panic(err)
	}

	uow := repository.NewUnitOfWorkImpl(pgConn)
	profileRepoCfg := repository.NewProfileCfgRepoImpl(uow)
	profileRepo := repository.NewProfileRepoImpl(uow)
	userRepo := repository.NewUserRepoImpl(uow)
	minioRepo := repository.NewMinioImpl(minioConn)

	accountUsecase := usecase.NewAccountUsecaseImpl(profileRepo, userRepo, minioRepo, 2*time.Second)
	profileUsecase := usecase.NewProfileUsecaseImpl(profileRepo, userRepo, 2*time.Second)
	profileCfgUsecase := usecase.NewProfileCfgUsecaseImpl(profileRepo, profileRepoCfg, 2*time.Second)
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(cusmiddleware.IPMiddleware)

	accountHandler := restapi.NewAccountHandler(accountUsecase)
	profileHandler := restapi.NewProfileHandler(profileUsecase)
	profileCfgHandler := restapi.NewProfileCfgHandler(profileCfgUsecase)

	r.Put("/account/{profile-id}", accountHandler.UpdateAccount)
	r.Get("/profile", profileHandler.GetProfileByID)
	r.Post("/profile", profileHandler.StoreProfile)
	r.Post("/profile-config/{profile-id}", profileCfgHandler.CreateProfileCfg)
	r.Get("/profile-config/{profile-id}/{config-name}", profileCfgHandler.GetProfileCfgByNameAndID)
	r.Put("/profile-config/{profile-id}/{config-name}", profileCfgHandler.UpdateProfileCfg)

	err = http.ListenAndServe(config.AppPort, r)
	if err != nil {
		panic(err)
	}
}
