package interactor_test

import (
	"context"
	"testing"

	"github.com/takokun778/firebase-authentication-proxy/usecase/interactor"
	"github.com/takokun778/firebase-authentication-proxy/usecase/port"
)

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

func TestKeyFetchInteractorExecute(t *testing.T) {
	t.Parallel()

	type fields struct {
		output port.KeyFetchOutputPort
	}

	type args struct {
		ctx context.Context
	}

	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "正常に動作することを確認",
			fields: fields{
				output: &MockKeyFetchOutputPort{
					Assert: func(output port.KeyFetchOutput) {
						if output.PublicKey == nil {
							t.Errorf("fail")
						}
					},
					ErrorAssert: func(err error) {
						t.Errorf("fail")
					},
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			i := interactor.NewKeyFetchInteractor(
				tt.fields.output,
			)
			i.Execute(tt.args.ctx)
		})
	}
}
