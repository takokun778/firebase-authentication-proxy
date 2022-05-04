package interactor_test

import (
	"context"
	"testing"

	"github.com/takokun778/firebase-authentication-proxy/domain/model/firebase"
	"github.com/takokun778/firebase-authentication-proxy/usecase/interactor"
	"github.com/takokun778/firebase-authentication-proxy/usecase/mock"
	"github.com/takokun778/firebase-authentication-proxy/usecase/port"
)

type MockFirebaseLoginOutputPort struct {
	Assert      func(output port.FirebaseLoginOutput)
	ErrorAssert func(err error)
}

func (m *MockFirebaseLoginOutputPort) Render(ctx context.Context, output port.FirebaseLoginOutput) {
	m.Assert(output)
}

func (m *MockFirebaseLoginOutputPort) ErrorRender(ctx context.Context, err error) {
	m.ErrorAssert(err)
}

func TestFirebaseLoginInteractorExecute(t *testing.T) {
	t.Parallel()

	type fields struct {
		firebaseRepository port.FirebaseRepository
		output             port.FirebaseLoginOutputPort
	}

	type args struct {
		ctx   context.Context
		input port.FirebaseLoginInput
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
					WantLoginError: false,
				},
				output: &MockFirebaseLoginOutputPort{
					Assert: func(output port.FirebaseLoginOutput) {},
					ErrorAssert: func(err error) {
						t.Errorf("fail")
					},
				},
			},
			args: args{
				ctx: context.Background(),
				input: port.FirebaseLoginInput{
					Email:    firebase.Email("test@example.com"),
					Password: firebase.Password("password"),
				},
			},
		},
		{
			name: "RepositoryにてErrorが発生することを確認",
			fields: fields{
				firebaseRepository: &mock.FirebaseRepository{
					WantLoginError: true,
				},
				output: &MockFirebaseLoginOutputPort{
					Assert: func(output port.FirebaseLoginOutput) {
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
				input: port.FirebaseLoginInput{
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
			i := interactor.NewFirebaseLoginInteractor(
				tt.fields.firebaseRepository,
				tt.fields.output,
			)
			i.Execute(tt.args.ctx, tt.args.input)
		})
	}
}
