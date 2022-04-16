package firebase

import "testing"

func TestNewPassword(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name    string
		args    args
		want    Password
		wantErr bool
	}{
		{
			name: "正常にパスワードが作成できることを確認",
			args: args{
				value: "password",
			},
			want:    Password("password"),
			wantErr: false,
		},
		{
			name: "6文字のパスワードを作成できることを確認",
			args: args{
				value: "aaaaaa",
			},
			want:    Password("aaaaaa"),
			wantErr: false,
		},
		{
			name: "5文字のパスワードを作成してエラーが発生することを確認",
			args: args{
				value: "aaaaa",
			},
			want:    Password(""),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewPassword(tt.args.value)
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
