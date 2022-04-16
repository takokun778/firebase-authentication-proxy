package injector

import (
	"firebase-authentication/adapter/controller"
	"firebase-authentication/adapter/gateway"
	"firebase-authentication/adapter/presenter"
	"firebase-authentication/driver/firebase"
	"firebase-authentication/usecase"
)

type Injector struct {
	Firebase *controller.FirebaseController
	Key      *controller.KeyController
}

func NewInjector() *Injector {
	fg := gateway.NewFirebaseGateway(firebase.AdminClient, firebase.ApiClient)
	ug := gateway.NewUserInMemory()

	fp := presenter.NewFirebasePresenter()
	kp := presenter.NewKeyPresenter()

	fi := usecase.NewFirebaseInteractor(fg, fp, ug)
	ki := usecase.NewKeyInteractor(kp)

	fc := controller.NewFirebaseController(fi)
	kc := controller.NewKeyController(ki)

	return &Injector{
		Firebase: fc,
		Key:      kc,
	}
}
