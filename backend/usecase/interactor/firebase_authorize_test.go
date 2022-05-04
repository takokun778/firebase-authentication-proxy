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

type MockFirebaseAuthorizeOutputPort struct {
	Assert      func(output port.FirebaseAuthorizeOutput)
	ErrorAssert func(err error)
}

func (m *MockFirebaseAuthorizeOutputPort) Render(ctx context.Context, output port.FirebaseAuthorizeOutput) {
	m.Assert(output)
}

func (m *MockFirebaseAuthorizeOutputPort) ErrorRender(ctx context.Context, err error) {
	m.ErrorAssert(err)
}

func TestFirebaseAuthorizeInteractorExecute(t *testing.T) {
	t.Parallel()

	type fields struct {
		firebaseRepository port.FirebaseRepository
		output             port.FirebaseAuthorizeOutputPort
	}

	type args struct {
		ctx   context.Context
		input port.FirebaseAuthorizeInput
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
				output: &MockFirebaseAuthorizeOutputPort{
					Assert: func(output port.FirebaseAuthorizeOutput) {},
					ErrorAssert: func(err error) {
						t.Errorf("fail")
					},
				},
			},
			args: args{
				ctx: context.Background(),
				input: port.FirebaseAuthorizeInput{
					AccessToken: firebase.AccessToken(mock.CreateTestJwt(uuid.New().String(), "test@example.com")),
				},
			},
		},
		{
			name: "RepositoryにてErrorが発生することを確認",
			fields: fields{
				firebaseRepository: &mock.FirebaseRepository{
					WantVerifyError: true,
				},
				output: &MockFirebaseAuthorizeOutputPort{
					Assert: func(output port.FirebaseAuthorizeOutput) {
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
				input: port.FirebaseAuthorizeInput{
					AccessToken: firebase.AccessToken(mock.CreateTestJwt(uuid.New().String(), "test@example.com")),
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			i := interactor.NewFirebaseAuthorizeInteractor(
				tt.fields.firebaseRepository,
				tt.fields.output,
			)
			i.Execute(tt.args.ctx, tt.args.input)
		})
	}
}
