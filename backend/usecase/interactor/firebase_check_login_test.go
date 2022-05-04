package interactor_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/takokun778/firebase-authentication-proxy/domain/model/firebase"
	"github.com/takokun778/firebase-authentication-proxy/usecase/interactor"
	"github.com/takokun778/firebase-authentication-proxy/usecase/mock"
	"github.com/takokun778/firebase-authentication-proxy/usecase/port"
)

type MockFirebaseCheckLoginOutputPort struct {
	Assert      func(output port.FirebaseCheckLoginOutput)
	ErrorAssert func(err error)
}

func (m *MockFirebaseCheckLoginOutputPort) Render(ctx context.Context, output port.FirebaseCheckLoginOutput) {
	m.Assert(output)
}

func (m *MockFirebaseCheckLoginOutputPort) ErrorRender(ctx context.Context, err error) {
	m.ErrorAssert(err)
}

func TestFirebaseCheckLoginInteractorExecute(t *testing.T) {
	t.Parallel()

	type fields struct {
		firebaseRepository port.FirebaseRepository
		output             port.FirebaseCheckLoginOutputPort
	}

	type args struct {
		ctx   context.Context
		input port.FirebaseCheckLoginInput
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
					WantVerifyError: false,
				},
				output: &MockFirebaseCheckLoginOutputPort{
					Assert: func(output port.FirebaseCheckLoginOutput) {},
					ErrorAssert: func(err error) {
						t.Errorf("fail")
					},
				},
			},
			args: args{
				ctx: context.Background(),
				input: port.FirebaseCheckLoginInput{
					AccessToken:  firebase.AccessToken(mock.CreateTestJwt(uuid.New().String(), "test@example.com")),
					RefreshToken: firebase.RefreshToken("refresh"),
				},
			},
		},
		{
			name: "RepositoryにてErrorが発生することを確認",
			fields: fields{
				firebaseRepository: &mock.FirebaseRepository{
					WantVerifyError: true,
				},
				output: &MockFirebaseCheckLoginOutputPort{
					Assert: func(output port.FirebaseCheckLoginOutput) {
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
				input: port.FirebaseCheckLoginInput{
					AccessToken:  firebase.AccessToken(mock.CreateTestJwt(uuid.New().String(), "test@example.com")),
					RefreshToken: firebase.RefreshToken(""),
				},
			},
		},
		{
			name: "inputが空文字のためErrorが発生することを確認",
			fields: fields{
				firebaseRepository: &mock.FirebaseRepository{
					WantVerifyError: true,
				},
				output: &MockFirebaseCheckLoginOutputPort{
					Assert: func(output port.FirebaseCheckLoginOutput) {
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
				input: port.FirebaseCheckLoginInput{
					AccessToken:  firebase.AccessToken(""),
					RefreshToken: firebase.RefreshToken(""),
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			i := interactor.NewFirebaseCheckLoginInteractor(
				tt.fields.firebaseRepository,
				tt.fields.output,
			)
			i.Execute(tt.args.ctx, tt.args.input)
		})
	}
}
