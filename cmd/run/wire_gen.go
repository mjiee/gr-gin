// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package run

import (
	"github.com/mjiee/grf-gin/app/api"
	"github.com/mjiee/grf-gin/app/lib"
	"github.com/mjiee/grf-gin/app/pkg/conf"
	"github.com/mjiee/grf-gin/app/pkg/db"
	"github.com/mjiee/grf-gin/app/pkg/zlog"
	"github.com/mjiee/grf-gin/app/router"
)

// Injectors from wire.go:

func initApp(confFile2 string) (*app, func(), error) {
	config, err := conf.NewConfig(confFile2)
	if err != nil {
		return nil, nil, err
	}
	logger := zlog.NewLogger(config)
	client, cleanup := db.NewRedis(config)
	jwtService := lib.NewJwtService(config, client)
	gormDB, cleanup2, err := db.NewMysql(config)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	userService := lib.NewUserService(gormDB)
	authHandler := api.NewAuthHandler(config, jwtService, userService)
	noAuthApi := router.NewNoAuthApi(authHandler)
	userHandler := api.NewUserHandler(userService)
	bucket, err := db.NewOssBucket(config)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	ossService := lib.NewOssService(config, client, bucket)
	ossHandler := api.NewOssHandler(ossService)
	authApi := router.NewAuthApi(config, jwtService, userHandler, ossHandler)
	engine := router.NewRouter(config, logger, noAuthApi, authApi)
	runApp := newApp(config, logger, engine)
	return runApp, func() {
		cleanup2()
		cleanup()
	}, nil
}
