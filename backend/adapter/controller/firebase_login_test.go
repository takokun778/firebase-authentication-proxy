package controller_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/takokun778/firebase-authentication-proxy/adapter/controller"
	"github.com/takokun778/firebase-authentication-proxy/adapter/mock"
	"github.com/takokun778/firebase-authentication-proxy/usecase/port"
)

type MockFirebaseLoginInputPort struct {
	Assert func(input port.FirebaseLoginInput)
}

func (m *MockFirebaseLoginInputPort) Execute(ctx context.Context, input port.FirebaseLoginInput) {
	m.Assert(input)
}

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

func TestFirebaseLoginControllerPost(t *testing.T) {
	t.Parallel()

	type fields struct {
		input  port.FirebaseLoginInputPort
		output port.FirebaseLoginOutputPort
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
				input: &MockFirebaseLoginInputPort{
					Assert: func(input port.FirebaseLoginInput) {},
				},
				output: &MockFirebaseLoginOutputPort{
					Assert: func(output port.FirebaseLoginOutput) {},
					ErrorAssert: func(err error) {
						t.Errorf("out put error render is called")
					},
				},
			},
			args: args{
				w: &mock.ResponseWriter{},
				r: &http.Request{
					Method: http.MethodPost,
					Body: mock.CreateFirebaseLoginPostBody(controller.FirebaseLoginPostBody{
						Email:    "test@example.com",
						Password: mock.Encrypt("password"),
					}),
				},
			},
		},
		{
			name: "PostMethod以外のリクエストでエラーが発生することを確認",
			fields: fields{
				input: &MockFirebaseLoginInputPort{
					Assert: func(input port.FirebaseLoginInput) {
						t.Errorf("input port execute is called")
					},
				},
				output: &MockFirebaseLoginOutputPort{
					Assert: func(output port.FirebaseLoginOutput) {
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
					Method: http.MethodPut,
					Body: mock.CreateFirebaseLoginPostBody(controller.FirebaseLoginPostBody{
						Email:    "test@example.com",
						Password: mock.Encrypt("password"),
					}),
				},
			},
		},
		{
			name: "公開鍵で暗号化していないパスワードがリクエストされるとエラーが発生することを確認",
			fields: fields{
				input: &MockFirebaseLoginInputPort{
					Assert: func(input port.FirebaseLoginInput) {
						t.Errorf("input port execute is called")
					},
				},
				output: &MockFirebaseLoginOutputPort{
					Assert: func(output port.FirebaseLoginOutput) {
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
					Body: mock.CreateFirebaseLoginPostBody(controller.FirebaseLoginPostBody{
						Email:    "test@example.com",
						Password: "password",
					}),
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := controller.NewFirebaseLoginController(
				tt.fields.input,
				tt.fields.output,
			)
			c.Post(tt.args.w, tt.args.r)
		})
	}
}
