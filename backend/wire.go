//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/takokun778/firebase-authentication-proxy/adapter/controller"
	"github.com/takokun778/firebase-authentication-proxy/adapter/gateway"
	"github.com/takokun778/firebase-authentication-proxy/adapter/presenter"
	"github.com/takokun778/firebase-authentication-proxy/driver/firebase"
	"github.com/takokun778/firebase-authentication-proxy/driver/server"
	"github.com/takokun778/firebase-authentication-proxy/usecase/interactor"
)

var gatewaySet = wire.NewSet(
	firebase.NewClient,
	gateway.NewFirebaseGateway,
	gateway.NewUserInMemory,
)

var presenterSet = wire.NewSet(
	presenter.NewKeyGetPublicPresenter,
	presenter.NewFirebaseAuthorizePresenter,
	presenter.NewFirebaseChangePasswordPresenter,
	presenter.NewFirebaseCheckLoginPresenter,
	presenter.NewFirebaseLoginPresenter,
	presenter.NewFirebaseLogoutPresenter,
	presenter.NewFirebaseRegisterPresenter,
	presenter.NewFirebaseWithdrawPresenter,
)

var interactorSet = wire.NewSet(
	interactor.NewKeyFetchPublicInteractor,
	interactor.NewFirebaseAuthorizeInteractor,
	interactor.NewFirebaseChangePasswordInteractor,
	interactor.NewFirebaseCheckLoginInteractor,
	interactor.NewFirebaseLoginInteractor,
	interactor.NewFirebaseLogoutInteractor,
	interactor.NewFirebaseRegisterInteractor,
	interactor.NewFirebaseWithdrawInteractor,
)

var controllerSet = wire.NewSet(
	controller.NewKeyFetchPublicController,
	controller.NewFirebaseAuthorizeController,
	controller.NewFirebaseChangePasswordController,
	controller.NewFirebaseCheckLoginController,
	controller.NewFirebaseLoginController,
	controller.NewFirebaseLogoutController,
	controller.NewFirebaseRegisterController,
	controller.NewFirebaseWithdrawController,
)

func InitializeHTTPServer() *server.HTTPServer {
	wire.Build(
		gatewaySet,
		presenterSet,
		interactorSet,
		controllerSet,
		server.NewHTTPServer,
	)

	return nil
}
