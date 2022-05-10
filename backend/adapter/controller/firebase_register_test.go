package controller_test

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"net/http"
	"testing"

	"github.com/takokun778/firebase-authentication-proxy/adapter/controller"
	"github.com/takokun778/firebase-authentication-proxy/adapter/mock"
	"github.com/takokun778/firebase-authentication-proxy/domain/model/key"
	"github.com/takokun778/firebase-authentication-proxy/usecase/port"
)

type MockFirebaseRegisterInputPort struct {
	Assert func(input port.FirebaseRegisterInput)
}

func (m *MockFirebaseRegisterInputPort) Execute(ctx context.Context, input port.FirebaseRegisterInput) {
	m.Assert(input)
}

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

func TestFirebaseRegisterControllerPost(t *testing.T) {
	t.Parallel()

	type fields struct {
		input  port.FirebaseRegisterInputPort
		output port.FirebaseRegisterOutputPort
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
				input: &MockFirebaseRegisterInputPort{
					Assert: func(input port.FirebaseRegisterInput) {},
				},
				output: &MockFirebaseRegisterOutputPort{
					Assert: func(output port.FirebaseRegisterOutput) {},
					ErrorAssert: func(err error) {
						t.Errorf("out put error render is called")
					},
				},
			},
			args: args{
				w: &mock.ResponseWriter{},
				r: &http.Request{
					Method: http.MethodPost,
					Body: mock.CreateFirebaseRegisterPostBody(controller.FirebaseRegisterPostBody{
						Email:    "test@example.com",
						Password: encrypt("password"),
					}),
				},
			},
		},
		{
			name: "PostMethod以外のリクエストでエラーが発生することを確認",
			fields: fields{
				input: &MockFirebaseRegisterInputPort{
					Assert: func(input port.FirebaseRegisterInput) {
						t.Errorf("input port execute is called")
					},
				},
				output: &MockFirebaseRegisterOutputPort{
					Assert: func(output port.FirebaseRegisterOutput) {
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
					Body: mock.CreateFirebaseRegisterPostBody(controller.FirebaseRegisterPostBody{
						Email:    "test@example.com",
						Password: encrypt("password"),
					}),
				},
			},
		},
		{
			name: "公開鍵で暗号化していないパスワードがリクエストされるとエラーが発生することを確認",
			fields: fields{
				input: &MockFirebaseRegisterInputPort{
					Assert: func(input port.FirebaseRegisterInput) {
						t.Errorf("input port execute is called")
					},
				},
				output: &MockFirebaseRegisterOutputPort{
					Assert: func(output port.FirebaseRegisterOutput) {
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
					Body: mock.CreateFirebaseRegisterPostBody(controller.FirebaseRegisterPostBody{
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
			c := controller.NewFirebaseRegisterController(
				tt.fields.input,
				tt.fields.output,
			)
			c.Post(tt.args.w, tt.args.r)
		})
	}
}

func encrypt(value string) string {
	enc, _ := rsa.EncryptOAEP(sha256.New(), rand.Reader, key.GetPublic(), []byte(value), nil)
	return base64.StdEncoding.EncodeToString(enc)
}
