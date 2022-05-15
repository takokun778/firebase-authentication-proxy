package controller_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/takokun778/firebase-authentication-proxy/adapter/controller"
	"github.com/takokun778/firebase-authentication-proxy/adapter/mock"
	"github.com/takokun778/firebase-authentication-proxy/usecase/port"
)

type MockFirebaseAuthorizeInputPort struct {
	Assert func(input port.FirebaseAuthorizeInput)
}

func (m *MockFirebaseAuthorizeInputPort) Execute(ctx context.Context, input port.FirebaseAuthorizeInput) {
	m.Assert(input)
}

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

func TestFirebaseAuthorizeControllerPost(t *testing.T) {
	t.Parallel()

	type fields struct {
		input  port.FirebaseAuthorizeInputPort
		output port.FirebaseAuthorizeOutputPort
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
				input: &MockFirebaseAuthorizeInputPort{
					Assert: func(input port.FirebaseAuthorizeInput) {},
				},
				output: &MockFirebaseAuthorizeOutputPort{
					Assert: func(output port.FirebaseAuthorizeOutput) {},
					ErrorAssert: func(err error) {
						if err != nil {
							t.Errorf("out put error render is called")
						}
					},
				},
			},
			args: args{
				w: &mock.ResponseWriter{},
				r: &http.Request{
					Method: http.MethodPost,
					Header: mock.CreateHeaderWithToken(),
				},
			},
		},
		{
			name: "PostMethod以外のリクエストでエラーが発生することを確認",
			fields: fields{
				input: &MockFirebaseAuthorizeInputPort{
					Assert: func(input port.FirebaseAuthorizeInput) {
						t.Errorf("input port execute is called")
					},
				},
				output: &MockFirebaseAuthorizeOutputPort{
					Assert: func(output port.FirebaseAuthorizeOutput) {
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
					Header: mock.CreateHeaderWithToken(),
				},
			},
		},
		{
			name: "Authorizationヘッダーがない場合エラーが発生することを確認",
			fields: fields{
				input: &MockFirebaseAuthorizeInputPort{
					Assert: func(input port.FirebaseAuthorizeInput) {
						t.Errorf("input port execute is called")
					},
				},
				output: &MockFirebaseAuthorizeOutputPort{
					Assert: func(output port.FirebaseAuthorizeOutput) {
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
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := controller.NewFirebaseAuthorizeController(
				tt.fields.input,
				tt.fields.output,
			)
			c.Post(tt.args.w, tt.args.r)
		})
	}
}
