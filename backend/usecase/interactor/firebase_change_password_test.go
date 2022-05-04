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

type MockFirebaseChangePasswordOutputPort struct {
	Assert      func(output port.FirebaseChangePasswordOutput)
	ErrorAssert func(err error)
}

func (m *MockFirebaseChangePasswordOutputPort) Render(ctx context.Context, output port.FirebaseChangePasswordOutput) {
	m.Assert(output)
}

func (m *MockFirebaseChangePasswordOutputPort) ErrorRender(ctx context.Context, err error) {
	m.ErrorAssert(err)
}

func TestFirebaseChangePasswordInteractorExecute(t *testing.T) {
	t.Parallel()

	type fields struct {
		firebaseRepository port.FirebaseRepository
		output             port.FirebaseChangePasswordOutputPort
	}

	type args struct {
		ctx   context.Context
		input port.FirebaseChangePasswordInput
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
					WantLoginError:          false,
					WantChangePasswordError: false,
				},
				output: &MockFirebaseChangePasswordOutputPort{
					Assert: func(output port.FirebaseChangePasswordOutput) {},
					ErrorAssert: func(err error) {
						t.Errorf("fail")
					},
				},
			},
			args: args{
				ctx: context.Background(),
				input: port.FirebaseChangePasswordInput{
					AccessToken: firebase.AccessToken(mock.CreateTestJwt(uuid.New().String(), "test@example.com")),
					OldPassword: firebase.Password("oldpassword"),
					NewPassword: firebase.Password("newpassword"),
				},
			},
		},
		{
			name: "FirebaseRepositoryにてLoginErrorが発生することを確認",
			fields: fields{
				firebaseRepository: &mock.FirebaseRepository{
					WantLoginError:          true,
					WantChangePasswordError: false,
				},
				output: &MockFirebaseChangePasswordOutputPort{
					Assert: func(output port.FirebaseChangePasswordOutput) {
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
				input: port.FirebaseChangePasswordInput{
					AccessToken: firebase.AccessToken(mock.CreateTestJwt(uuid.New().String(), "test@example.com")),
					OldPassword: firebase.Password("oldpassword"),
					NewPassword: firebase.Password("newpassword"),
				},
			},
		},
		{
			name: "FirebaseRepositoryにてChangePasswordErrorが発生することを確認",
			fields: fields{
				firebaseRepository: &mock.FirebaseRepository{
					WantLoginError:          false,
					WantChangePasswordError: true,
				},
				output: &MockFirebaseChangePasswordOutputPort{
					Assert: func(output port.FirebaseChangePasswordOutput) {
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
				input: port.FirebaseChangePasswordInput{
					AccessToken: firebase.AccessToken(mock.CreateTestJwt(uuid.New().String(), "test@example.com")),
					OldPassword: firebase.Password("oldpassword"),
					NewPassword: firebase.Password("newpassword"),
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			i := interactor.NewFirebaseChangePasswordInteractor(
				tt.fields.firebaseRepository,
				tt.fields.output,
			)
			i.Execute(tt.args.ctx, tt.args.input)
		})
	}
}
