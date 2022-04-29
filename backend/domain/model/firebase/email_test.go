package firebase_test

import (
	"testing"

	"github.com/takokun778/firebase-authentication-proxy/domain/model/firebase"
)

func TestNewEmail(t *testing.T) {
	t.Parallel()

	type args struct {
		value string
	}

	tests := []struct {
		name    string
		args    args
		want    firebase.Email
		wantErr bool
	}{
		{
			name: "正常にEmailが作成できることを確認",
			args: args{
				value: "sample@example.com",
			},
			want:    firebase.Email("sample@example.com"),
			wantErr: false,
		},
		{
			name: "Emailでない文字列でエラーが発生することを確認",
			args: args{
				value: "sampleexample.com",
			},
			want:    firebase.Email(""),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := firebase.NewEmail(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewEmail() error = %v, wantErr %v", err, tt.wantErr)

				return
			}
			if got != tt.want {
				t.Errorf("NewEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}
