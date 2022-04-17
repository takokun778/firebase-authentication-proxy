# firebase-authentication

## ディレクトリ構成

```bash
.
├── main.go
├── domain/model
│   ├── errors
│   │   └── errors.go
│   ├── firebase
│   │   ├── value.go
│   │   └── repository.go
│   ├── key
│   │   └── key.go
│   ├── user
│   │   ├── value.go
│   │   └── repository.go
│   ├── user.go
│   └── firebase.go
├── usecase
│   ├── firebase_input_port.go
│   ├── firebase_output_port.go
│   ├── firebase_interactor.go
│   ├── key_input_port.go
│   ├── key_output_port.go
│   └── key_interactor.go
├── adapter
│   ├── controller
│   │   ├── firebase.go
│   │   └── key.go
│   ├── gateway
│   │   ├── firebase.go
│   │   └── user.go
│   └── presenter
│       ├── firebase.go
│       ├── key.go
│       └── error.go
└── driver
    ├── server
    │   ├── middleware.go
    │   ├── rooter.go
    │   └── server.go
    ├── firebase
    │   └── client.go
    ├── injector
    │   └── injector.go
    ├── log
    │   └── log.go
    ├── env
    │   └── env.go
    ├── db
    │   └── client.go
    └── rest
        └── client.go
```
