package interactor_test

import (
	"context"
	"testing"

	"github.com/takokun778/firebase-authentication-proxy/domain/model/firebase"
	"github.com/takokun778/firebase-authentication-proxy/usecase/interactor"
	"github.com/takokun778/firebase-authentication-proxy/usecase/mock"
	"github.com/takokun778/firebase-authentication-proxy/usecase/port"
)

type MockFirebaseRegisterOutputPort struct {
	Assert      func(output port.FirebaseRegisterOutput)
	ErrorAssert func(err error)
}

func (m *MockFirebaseRegisterOutputPort) Render(ctx context.Context, output port.FirebaseRegisterOutput) {
	m.Assert(output)
}

func (m *MockFirebaseRegisterOutputPort) ErrorRender(ctx context.Context, err error) {
	m.ErrorAssert(err)
}

func TestFirebaseRegisterInteractorExecute(t *testing.T) {
	t.Parallel()

	type fields struct {
		firebaseRepository port.FirebaseRepository
		userRepository     port.UserRepository
		output             port.FirebaseRegisterOutputPort
	}

	type args struct {
		ctx   context.Context
		input port.FirebaseRegisterInput
	}

	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "正常に動作することを確認",
			fields: fields{
				firebaseRepository: &mock.FirebaseRepository{
					WantSaveError: false,
				},
				userRepository: &mock.UserRepository{
					WantSaveError: false,
				},
				output: &MockFirebaseRegisterOutputPort{
					Assert: func(output port.FirebaseRegisterOutput) {},
					ErrorAssert: func(err error) {
						t.Errorf("fail")
					},
				},
			},
			args: args{
				ctx: context.Background(),
				input: port.FirebaseRegisterInput{
					Email:    firebase.Email("test@example.com"),
					Password: firebase.Password("password"),
				},
			},
		},
		{
			name: "FirebaseRepositoryにてErrorが発生することを確認",
			fields: fields{
				firebaseRepository: &mock.FirebaseRepository{
					WantSaveError: true,
				},
				userRepository: &mock.UserRepository{
					WantSaveError: false,
				},
				output: &MockFirebaseRegisterOutputPort{
					Assert: func(output port.FirebaseRegisterOutput) {
						t.Errorf("fail")
					},
					ErrorAssert: func(err error) {
						if err == nil {
							t.Errorf("error is nil")
						}
					},
				},
			},
			args: args{
				ctx: context.Background(),
				input: port.FirebaseRegisterInput{
					Email:    firebase.Email("test@example.com"),
					Password: firebase.Password("password"),
				},
			},
		},
		{
			name: "UserRepositoryにてErrorが発生することを確認",
			fields: fields{
				firebaseRepository: &mock.FirebaseRepository{
					WantSaveError: false,
				},
				userRepository: &mock.UserRepository{
					WantSaveError: true,
				},
				output: &MockFirebaseRegisterOutputPort{
					Assert: func(output port.FirebaseRegisterOutput) {
						t.Errorf("fail")
					},
					ErrorAssert: func(err error) {
						if err == nil {
							t.Errorf("error is nil")
						}
					},
				},
			},
			args: args{
				ctx: context.Background(),
				input: port.FirebaseRegisterInput{
					Email:    firebase.Email("test@example.com"),
					Password: firebase.Password("password"),
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			i := interactor.NewFirebaseRegisterInteractor(
				tt.fields.firebaseRepository,
				tt.fields.userRepository,
				tt.fields.output,
			)
			i.Execute(tt.args.ctx, tt.args.input)
		})
	}
}
