package controller_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/takokun778/firebase-authentication-proxy/adapter/controller"
	"github.com/takokun778/firebase-authentication-proxy/adapter/mock"
	"github.com/takokun778/firebase-authentication-proxy/usecase/port"
)

type MockFirebaseCheckLoginInputPort struct {
	Assert func(input port.FirebaseCheckLoginInput)
}

func (m *MockFirebaseCheckLoginInputPort) Execute(ctx context.Context, input port.FirebaseCheckLoginInput) {
	m.Assert(input)
}

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

func TestFirebaseCheckLoginControllerPost(t *testing.T) {
	t.Parallel()

	type fields struct {
		input  port.FirebaseCheckLoginInputPort
		output port.FirebaseCheckLoginOutputPort
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
				input: &MockFirebaseCheckLoginInputPort{
					Assert: func(input port.FirebaseCheckLoginInput) {},
				},
				output: &MockFirebaseCheckLoginOutputPort{
					Assert: func(output port.FirebaseCheckLoginOutput) {},
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
				input: &MockFirebaseCheckLoginInputPort{
					Assert: func(input port.FirebaseCheckLoginInput) {
						t.Errorf("input port execute is called")
					},
				},
				output: &MockFirebaseCheckLoginOutputPort{
					Assert: func(output port.FirebaseCheckLoginOutput) {
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
				}),
			},
		},
		{
			name: "認証Tokenがない場合にエラーが発生することを確認",
			fields: fields{
				input: &MockFirebaseCheckLoginInputPort{
					Assert: func(input port.FirebaseCheckLoginInput) {
						t.Errorf("input port execute is called")
					},
				},
				output: &MockFirebaseCheckLoginOutputPort{
					Assert: func(output port.FirebaseCheckLoginOutput) {
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
			c := controller.NewFirebaseCheckLoginController(
				tt.fields.input,
				tt.fields.output,
			)
			c.Post(tt.args.w, tt.args.r)
		})
	}
}
