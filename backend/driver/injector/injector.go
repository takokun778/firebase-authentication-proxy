package injector

import (
	"github.com/takokun778/firebase-authentication-proxy/adapter/controller"
	"github.com/takokun778/firebase-authentication-proxy/adapter/gateway"
	"github.com/takokun778/firebase-authentication-proxy/adapter/presenter"
	"github.com/takokun778/firebase-authentication-proxy/driver/firebase"
	"github.com/takokun778/firebase-authentication-proxy/usecase"
)

type Injector struct {
	Firebase *controller.FirebaseController
	Key      *controller.KeyController
}

func NewInjector() *Injector {
	fg := gateway.NewFirebaseGateway(firebase.NewFirebaseClient())
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
