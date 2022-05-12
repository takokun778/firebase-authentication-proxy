package controller_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/takokun778/firebase-authentication-proxy/adapter/controller"
	"github.com/takokun778/firebase-authentication-proxy/adapter/mock"
	"github.com/takokun778/firebase-authentication-proxy/usecase/port"
)

type MockFirebaseLogoutInputPort struct {
	Assert func(input port.FirebaseLogoutInput)
}

func (m *MockFirebaseLogoutInputPort) Execute(ctx context.Context, input port.FirebaseLogoutInput) {
	m.Assert(input)
}

type MockFirebaseLogoutOutputPort struct {
	Assert      func(output port.FirebaseLogoutOutput)
	ErrorAssert func(err error)
}

func (m *MockFirebaseLogoutOutputPort) Render(ctx context.Context, output port.FirebaseLogoutOutput) {
	m.Assert(output)
}

func (m *MockFirebaseLogoutOutputPort) ErrorRender(ctx context.Context, err error) {
	m.ErrorAssert(err)
}

func TestFirebaseLogoutControllerPost(t *testing.T) {
	t.Parallel()

	type fields struct {
		input  port.FirebaseLogoutInputPort
		output port.FirebaseLogoutOutputPort
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
				input: &MockFirebaseLogoutInputPort{
					Assert: func(input port.FirebaseLogoutInput) {},
				},
				output: &MockFirebaseLogoutOutputPort{
					Assert: func(output port.FirebaseLogoutOutput) {},
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
				}),
			},
		},
		{
			name: "PostMethod以外のリクエストでエラーが発生することを確認",
			fields: fields{
				input: &MockFirebaseLogoutInputPort{
					Assert: func(input port.FirebaseLogoutInput) {
						t.Errorf("input port execute is called")
					},
				},
				output: &MockFirebaseLogoutOutputPort{
					Assert: func(output port.FirebaseLogoutOutput) {
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
					Method: http.MethodGet,
					Header: http.Header{},
				}),
			},
		},
		{
			name: "認証Tokenがない場合にエラーが発生することを確認",
			fields: fields{
				input: &MockFirebaseLogoutInputPort{
					Assert: func(input port.FirebaseLogoutInput) {
						t.Errorf("input port execute is called")
					},
				},
				output: &MockFirebaseLogoutOutputPort{
					Assert: func(output port.FirebaseLogoutOutput) {
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
					Method: http.MethodGet,
					Header: http.Header{},
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := controller.NewFirebaseLogoutController(
				tt.fields.input,
				tt.fields.output,
			)
			c.Post(tt.args.w, tt.args.r)
		})
	}
}
