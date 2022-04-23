package firebase_test

import (
	"testing"

	"github.com/takokun778/firebase-authentication-proxy/domain/model/firebase"
)

func TestNewPassword(t *testing.T) {
	t.Parallel()

	type args struct {
		value string
	}

	tests := []struct {
		name    string
		args    args
		want    firebase.Password
		wantErr bool
	}{
		{
			name: "正常にパスワードが作成できることを確認",
			args: args{
				value: "password",
			},
			want:    firebase.Password("password"),
			wantErr: false,
		},
		{
			name: "6文字のパスワードを作成できることを確認",
			args: args{
				value: "aaaaaa",
			},
			want:    firebase.Password("aaaaaa"),
			wantErr: false,
		},
		{
			name: "5文字のパスワードを作成してエラーが発生することを確認",
			args: args{
				value: "aaaaa",
			},
			want:    firebase.Password(""),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := firebase.NewPassword(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewPassword() error = %v, wantErr %v", err, tt.wantErr)

				return
			}
			if got != tt.want {
				t.Errorf("NewPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}
