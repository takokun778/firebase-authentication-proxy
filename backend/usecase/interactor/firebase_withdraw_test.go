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

type MockFirebaseWithdrawOutputPort struct {
	Assert      func(output port.FirebaseWithdrawOutput)
	ErrorAssert func(err error)
}

func (m *MockFirebaseWithdrawOutputPort) Render(ctx context.Context, output port.FirebaseWithdrawOutput) {
	m.Assert(output)
}

func (m *MockFirebaseWithdrawOutputPort) ErrorRender(ctx context.Context, err error) {
	m.ErrorAssert(err)
}

func TestFirebaseWithdrawInteractorExecute(t *testing.T) {
	t.Parallel()

	type fields struct {
		firebaseRepository port.FirebaseRepository
		output             port.FirebaseWithdrawOutputPort
	}

	type args struct {
		ctx   context.Context
		input port.FirebaseWithdrawInput
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
					WantLoginError:  false,
				},
				output: &MockFirebaseWithdrawOutputPort{
					Assert: func(output port.FirebaseWithdrawOutput) {},
					ErrorAssert: func(err error) {
						t.Errorf("fail")
					},
				},
			},
			args: args{
				ctx: context.Background(),
				input: port.FirebaseWithdrawInput{
					AccessToken: firebase.AccessToken(mock.CreateTestJwt(uuid.New().String(), "test@example.com")),
					Password:    firebase.Password("password"),
				},
			},
		},
		{
			name: "RepositoryにてVerifyErrorが発生することを確認",
			fields: fields{
				firebaseRepository: &mock.FirebaseRepository{
					WantVerifyError: true,
					WantLoginError:  false,
				},
				output: &MockFirebaseWithdrawOutputPort{
					Assert: func(output port.FirebaseWithdrawOutput) {
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
				input: port.FirebaseWithdrawInput{
					AccessToken: firebase.AccessToken(mock.CreateTestJwt(uuid.New().String(), "test@example.com")),
					Password:    firebase.Password("password"),
				},
			},
		},
		{
			name: "RepositoryにてLoginErrorが発生することを確認",
			fields: fields{
				firebaseRepository: &mock.FirebaseRepository{
					WantVerifyError: false,
					WantLoginError:  true,
				},
				output: &MockFirebaseWithdrawOutputPort{
					Assert: func(output port.FirebaseWithdrawOutput) {
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
				input: port.FirebaseWithdrawInput{
					AccessToken: firebase.AccessToken(mock.CreateTestJwt(uuid.New().String(), "test@example.com")),
					Password:    firebase.Password("password"),
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			i := interactor.NewFirebaseWithdrawInteractor(
				tt.fields.firebaseRepository,
				tt.fields.output,
			)
			i.Execute(tt.args.ctx, tt.args.input)
		})
	}
}
