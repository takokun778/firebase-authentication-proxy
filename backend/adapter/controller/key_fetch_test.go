package controller_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/takokun778/firebase-authentication-proxy/adapter/controller"
	"github.com/takokun778/firebase-authentication-proxy/adapter/mock"
	"github.com/takokun778/firebase-authentication-proxy/usecase/port"
)

type MockKeyFetchInputPort struct {
	Assert func()
}

func (m *MockKeyFetchInputPort) Execute(ctx context.Context) {
	m.Assert()
}

type MockKeyFetchOutputPort struct {
	Assert      func(output port.KeyFetchOutput)
	ErrorAssert func(err error)
}

func (m *MockKeyFetchOutputPort) Render(ctx context.Context, output port.KeyFetchOutput) {
	m.Assert(output)
}

func (m *MockKeyFetchOutputPort) ErrorRender(ctx context.Context, err error) {
	m.ErrorAssert(err)
}

func TestKeyFetchControllerGet(t *testing.T) {
	t.Parallel()

	type fields struct {
		input  port.KeyFetchInputPort
		output port.KeyFetchOutputPort
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
				input: &MockKeyFetchInputPort{
					Assert: func() {},
				},
				output: &MockKeyFetchOutputPort{
					Assert: func(output port.KeyFetchOutput) {},
					ErrorAssert: func(err error) {
						t.Errorf("out put error render is called")
					},
				},
			},
			args: args{
				w: &mock.ResponseWriter{},
				r: &http.Request{
					Method: http.MethodGet,
				},
			},
		},
		{
			name: "GetMethod以外のリクエストでエラーが発生することを確認",
			fields: fields{
				input: &MockKeyFetchInputPort{
					Assert: func() {
						t.Errorf("input port execute is called")
					},
				},
				output: &MockKeyFetchOutputPort{
					Assert: func(output port.KeyFetchOutput) {
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
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := controller.NewKeyFetchController(
				tt.fields.input,
				tt.fields.output,
			)
			c.Get(tt.args.w, tt.args.r)
		})
	}
}
