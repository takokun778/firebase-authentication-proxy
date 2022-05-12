package controller_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/takokun778/firebase-authentication-proxy/adapter/controller"
	"github.com/takokun778/firebase-authentication-proxy/adapter/mock"
	"github.com/takokun778/firebase-authentication-proxy/usecase/port"
)

type MockFirebaseChangePasswordInputPort struct {
	Assert func(input port.FirebaseChangePasswordInput)
}

func (m *MockFirebaseChangePasswordInputPort) Execute(ctx context.Context, input port.FirebaseChangePasswordInput) {
	m.Assert(input)
}

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

func TestFirebaseChangePasswordControllerPut(t *testing.T) {
	t.Parallel()

	type fields struct {
		input  port.FirebaseChangePasswordInputPort
		output port.FirebaseChangePasswordOutputPort
	}

	type args struct {
		w http.ResponseWriter
		r *http.Request
	}

	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "正常に動作することを確認",
			fields: fields{
				input: &MockFirebaseChangePasswordInputPort{
					Assert: func(input port.FirebaseChangePasswordInput) {},
				},
				output: &MockFirebaseChangePasswordOutputPort{
					Assert: func(output port.FirebaseChangePasswordOutput) {},
					ErrorAssert: func(err error) {
						t.Errorf("out put error render is called")
					},
				},
			},
			args: args{
				w: &mock.ResponseWriter{},
				r: mock.WithToken(&http.Request{
					Method: http.MethodPut,
					Header: http.Header{},
					Body: mock.CreateFirebaseChangePasswordPutBody(controller.FirebaseChangePasswordPutBody{
						OldPassword: mock.Encrypt("oldpassword"),
						NewPassword: mock.Encrypt("newpassword"),
					}),
				}),
			},
		},
		{
			name: "PutMethod以外のリクエストでエラーが発生することを確認",
			fields: fields{
				input: &MockFirebaseChangePasswordInputPort{
					Assert: func(input port.FirebaseChangePasswordInput) {
						t.Errorf("input port execute is called")
					},
				},
				output: &MockFirebaseChangePasswordOutputPort{
					Assert: func(output port.FirebaseChangePasswordOutput) {
						t.Errorf("out port render is called")
					},
					ErrorAssert: func(err error) {
						if err == nil {
							t.Errorf("error is nil")
						}
					},
				},
			},
			args: args{
				w: &mock.ResponseWriter{},
				r: mock.WithToken(&http.Request{
					Method: http.MethodPost,
					Header: http.Header{},
					Body: mock.CreateFirebaseChangePasswordPutBody(controller.FirebaseChangePasswordPutBody{
						OldPassword: "oldpassword",
						NewPassword: "newpassword",
					}),
				}),
			},
		},
		{
			name: "認証Tokenがない場合にエラーが発生することを確認",
			fields: fields{
				input: &MockFirebaseChangePasswordInputPort{
					Assert: func(input port.FirebaseChangePasswordInput) {
						t.Errorf("input port execute is called")
					},
				},
				output: &MockFirebaseChangePasswordOutputPort{
					Assert: func(output port.FirebaseChangePasswordOutput) {
						t.Errorf("out port render is called")
					},
					ErrorAssert: func(err error) {
						if err == nil {
							t.Errorf("error is nil")
						}
					},
				},
			},
			args: args{
				w: &mock.ResponseWriter{},
				r: &http.Request{
					Method: http.MethodPost,
					Header: http.Header{},
					Body: mock.CreateFirebaseChangePasswordPutBody(controller.FirebaseChangePasswordPutBody{
						OldPassword: "oldpassword",
						NewPassword: "newpassword",
					}),
				},
			},
		},
		{
			name: "公開鍵で暗号化していないパスワードがリクエストされるとエラーが発生することを確認",
			fields: fields{
				input: &MockFirebaseChangePasswordInputPort{
					Assert: func(input port.FirebaseChangePasswordInput) {
						t.Errorf("input port execute is called")
					},
				},
				output: &MockFirebaseChangePasswordOutputPort{
					Assert: func(output port.FirebaseChangePasswordOutput) {
						t.Errorf("out port render is called")
					},
					ErrorAssert: func(err error) {
						if err == nil {
							t.Errorf("error is nil")
						}
					},
				},
			},
			args: args{
				w: &mock.ResponseWriter{},
				r: mock.WithToken(&http.Request{
					Method: http.MethodPost,
					Header: http.Header{},
					Body: mock.CreateFirebaseChangePasswordPutBody(controller.FirebaseChangePasswordPutBody{
						OldPassword: mock.Encrypt("oldpassword"),
						NewPassword: mock.Encrypt("newpassword"),
					}),
				}),
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := controller.NewFirebaseChangePasswordController(
				tt.fields.input,
				tt.fields.output,
			)
			c.Put(tt.args.w, tt.args.r)
		})
	}
}
