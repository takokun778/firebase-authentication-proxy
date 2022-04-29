package firebase_test

import (
	"reflect"
	"testing"

	"github.com/google/uuid"
	"github.com/takokun778/firebase-authentication-proxy/domain/model/firebase"
)

func TestNewUID(t *testing.T) {
	t.Parallel()

	type args struct {
		value string
	}

	id := uuid.New()

	tests := []struct {
		name    string
		args    args
		want    firebase.UID
		wantErr bool
	}{
		{
			name: "UIDが正常に作成されることを確認",
			args: args{
				value: id.String(),
			},
			want:    firebase.UID(id),
			wantErr: false,
		},
		{
			name: "uuidでない文字列でエラーが発生することを確認",
			args: args{
				value: "uid",
			},
			want:    firebase.UID(uuid.UUID{}),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := firebase.NewUID(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewUID() error = %v, wantErr %v", err, tt.wantErr)

				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUID() = %v, want %v", got, tt.want)
			}
		})
	}
}
