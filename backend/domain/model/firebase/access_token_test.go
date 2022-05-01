package firebase_test

import (
	"log"
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/takokun778/firebase-authentication-proxy/domain/model/firebase"
)

func TestNewAccessToken(t *testing.T) {
	t.Parallel()

	type args struct {
		value string
	}

	uid := uuid.New().String()
	email := "test@example.com"
	jwt := CreateTestJwt(uid, email)

	tests := []struct {
		name    string
		args    args
		want    firebase.AccessToken
		wantErr bool
	}{
		{
			name: "正常に作成できることを確認",
			args: args{
				value: jwt,
			},
			want:    firebase.AccessToken(jwt),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := firebase.NewAccessToken(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewAccessToken() error = %v, wantErr %v", err, tt.wantErr)

				return
			}
			if got != tt.want {
				t.Errorf("NewAccessToken() = %v, want %v", got, tt.want)
			}

			gotUID, err := got.GetUID()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUID() error = %v, wantErr %v", err, tt.wantErr)

				return
			}
			if gotUID.String() != uid {
				t.Errorf("GetUID() = %v, want %v", uid, gotUID.String())

				return
			}

			gotEmail, err := got.GetEmail()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetEmail() error = %v, wantErr %v", err, tt.wantErr)

				return
			}
			if gotEmail.Value() != email {
				t.Errorf("GetEmail() = %v, want %v", email, gotEmail.Value())

				return
			}

			gotExpires, err := got.CalcExpires()
			if (err != nil) != tt.wantErr {
				t.Errorf("CalcExpires() error = %v, wantErr %v", err, tt.wantErr)

				return
			}
			if gotExpires != 3600 {
				t.Errorf("CalcExpires() = %v, wantErr %v", 3600, gotExpires)

				return
			}
		})
	}
}

func CreateTestJwt(userID, email string) string {
	// https://firebase.google.com/docs/auth/admin/verify-id-tokens#verify_id_tokens_using_a_third-party_jwt_library
	token := jwt.New(jwt.SigningMethodRS256)

	claims, _ := token.Claims.(jwt.MapClaims)

	now := time.Now().UTC()

	claims["iss"] = "iss"
	claims["aud"] = "aud"
	claims["auth_time"] = now.Unix()
	claims["user_id"] = userID
	claims["sub"] = uuid.New()
	claims["iat"] = now.Unix()
	claims["exp"] = now.Add(time.Hour).Unix()
	claims["email"] = email
	claims["email_verified"] = false
	claims["firebase"] = ""

	jwt, err := token.SigningString()
	if err != nil {
		log.Fatalln(err.Error())
	}

	return jwt
}
