package controller_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/takokun778/firebase-authentication-proxy/adapter/controller"
	"github.com/takokun778/firebase-authentication-proxy/adapter/mock"
	"github.com/takokun778/firebase-authentication-proxy/usecase/port"
)

type MockFirebaseWithdrawInputPort struct {
	Assert func(input port.FirebaseWithdrawInput)
}

func (m *MockFirebaseWithdrawInputPort) Execute(ctx context.Context, input port.FirebaseWithdrawInput) {
	m.Assert(input)
}

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

func TestFirebaseWithdrawControllerPost(t *testing.T) {
	t.Parallel()

	type fields struct {
		input  port.FirebaseWithdrawInputPort
		output port.FirebaseWithdrawOutputPort
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
				input: &MockFirebaseWithdrawInputPort{
					Assert: func(input port.FirebaseWithdrawInput) {},
				},
				output: &MockFirebaseWithdrawOutputPort{
					Assert: func(output port.FirebaseWithdrawOutput) {},
					ErrorAssert: func(err error) {
						if err != nil {
							t.Errorf("out put error render is called")
						}
					},
				},
			},
			args: args{
				w: &mock.ResponseWriter{},
				r: mock.WithToken(&http.Request{
					Method: http.MethodPost,
					Header: http.Header{},
					Body: mock.CreateFirebaseWithdrawPostBody(controller.FirebaseWithdrawPostBody{
						Password: mock.Encrypt("password"),
					}),
				}),
			},
		},
		{
			name: "PostMethod以外のリクエストでエラーが発生することを確認",
			fields: fields{
				input: &MockFirebaseWithdrawInputPort{
					Assert: func(input port.FirebaseWithdrawInput) {
						t.Errorf("input port execute is called")
					},
				},
				output: &MockFirebaseWithdrawOutputPort{
					Assert: func(output port.FirebaseWithdrawOutput) {
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
					Method: http.MethodPut,
					Header: http.Header{},
					Body: mock.CreateFirebaseWithdrawPostBody(controller.FirebaseWithdrawPostBody{
						Password: mock.Encrypt("password"),
					}),
				}),
			},
		},
		{
			name: "認証Tokenがない場合にエラーが発生することを確認",
			fields: fields{
				input: &MockFirebaseWithdrawInputPort{
					Assert: func(input port.FirebaseWithdrawInput) {
						t.Errorf("input port execute is called")
					},
				},
				output: &MockFirebaseWithdrawOutputPort{
					Assert: func(output port.FirebaseWithdrawOutput) {
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
					Header: http.Header{},
					Body: mock.CreateFirebaseWithdrawPostBody(controller.FirebaseWithdrawPostBody{
						Password: mock.Encrypt("password"),
					}),
				},
			},
		},
		{
			name: "公開鍵で暗号化していないパスワードがリクエストされるとエラーが発生することを確認",
			fields: fields{
				input: &MockFirebaseWithdrawInputPort{
					Assert: func(input port.FirebaseWithdrawInput) {
						t.Errorf("input port execute is called")
					},
				},
				output: &MockFirebaseWithdrawOutputPort{
					Assert: func(output port.FirebaseWithdrawOutput) {
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
					Method: http.MethodPut,
					Header: http.Header{},
					Body: mock.CreateFirebaseWithdrawPostBody(controller.FirebaseWithdrawPostBody{
						Password: "password",
					}),
				}),
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := controller.NewFirebaseWithdrawController(
				tt.fields.input,
				tt.fields.output,
			)
			c.Post(tt.args.w, tt.args.r)
		})
	}
}
