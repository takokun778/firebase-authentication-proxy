# firebase-authentication

## ディレクトリ構成

```bash
.
├── main.go
├── domain/model
│   ├── error
│   │   └── error.go
│   └── firebase
│       └── firebase.go
├── usecase
│   ├── firebase_repository.go
│   ├── firebase_input_port.go
│   ├── firebase_output_port.go
│   └── firebase_interactor.go
├── adapter
│   ├── controller
│   │   ├── firebase_controller.go
│   │   └── firebase_controller.go
│   ├── gateway
│   │   ├── firebase_entity.go
│   │   ├── firebase_gateway.go
│   │   └── firebase_sql.go
│   └── presenter
│       └── firebase_presenter.go
└── driver
    ├── server
    │   ├── server.go
    │   ├── rooter.go
    │   └── middleware.go
    ├── db
    │   └── db.go
    └── rest
        └── client.go
```
