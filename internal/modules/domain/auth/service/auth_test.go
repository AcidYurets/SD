package service

import (
	"calend/internal/models/err_const"
	"calend/internal/models/session"
	"calend/internal/modules/domain/auth/dto"
	user_dto "calend/internal/modules/domain/user/dto"
	"context"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

const secret = "123"

// hash - хэш для пароля Pass123!
const hash = "$2a$14$qXllSjS8CzmsBObgqZNTT.BmwX4dp8xHs.1IAAITXdUixsDtyJzsK"

const validToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODMwMjkwMDAsImlhdCI6MTY4MDAxNDMxOCwianRpIjoiZjMyMjM4MmMtOWVhNy00NGJiLWE4OTctZWEyMWZjODk0N2MyIiwic2Vzc2lvbiI6eyJTSUQiOiI1MDk4MjlkOC1hYjAwLTQ2MDgtYThkOS00MzJiZjdmOWIxMTgiLCJVc2VyVXVpZCI6IjEyM2U0NTY3LWU4OWItMTJkMy1hNDU2LTQyNjY1NTQ0MDAwMCJ9fQ.KZXQNLByEg2lz6gxmr2Pbk5TIpPzXeAIxCaqJncwJ08"
const expiredToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODAwMDEwMDAsImlhdCI6MTY4MDAxNDMxOCwianRpIjoiZjMyMjM4MmMtOWVhNy00NGJiLWE4OTctZWEyMWZjODk0N2MyIiwic2Vzc2lvbiI6eyJTSUQiOiI1MDk4MjlkOC1hYjAwLTQ2MDgtYThkOS00MzJiZjdmOWIxMTgiLCJVc2VyVXVpZCI6IjEyM2U0NTY3LWU4OWItMTJkMy1hNDU2LTQyNjY1NTQ0MDAwMCJ9fQ.Yw_xFpqeDbwhVTZ1NtdMUq2d04wglTm_2KFE7atiXb8"
const invalidToken = "invalid_eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODAwMDEwMDAsImlhdCI6MTY4MDAxNDMxOCwianRpIjoiZjMyMjM4MmMtOWVhNy00NGJiLWE4OTctZWEyMWZjODk0N2MyIiwic2Vzc2lvbiI6eyJTSUQiOiI1MDk4MjlkOC1hYjAwLTQ2MDgtYThkOS00MzJiZjdmOWIxMTgiLCJVc2VyVXVpZCI6IjEyM2U0NTY3LWU4OWItMTJkMy1hNDU2LTQyNjY1NTQ0MDAwMCJ9fQ.Yw_xFpqeDbwhVTZ1NtdMUq2d04wglTm_2KFE7atiXb8"

// Идентификаторы из токена
const SID = "509829d8-ab00-4608-a8d9-432bf7f9b118"
const UserUuid = "123e4567-e89b-12d3-a456-426655440000"

func TestAuthService_SignUp(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expectedUser := &user_dto.User{
		Uuid:         "123e4567-e89b-12d3-a456-426655440000",
		Phone:        "89197628803",
		Login:        "user1",
		PasswordHash: "",
	}

	type fields struct {
		repo   *MockIUserRepo
		secret string
		expect func(repo *MockIUserRepo)
	}
	type args struct {
		ctx     context.Context
		newUser *dto.NewUser
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *user_dto.User
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "invalid",
			fields: fields{
				repo:   NewMockIUserRepo(ctrl),
				secret: secret,
				expect: func(repo *MockIUserRepo) {
					repo.EXPECT().Create(gomock.Any(), gomock.Any()).Times(0)
				},
			},
			args: args{
				ctx: context.Background(),
				newUser: &dto.NewUser{
					Login:    "user1",
					Password: "pass",
					Phone:    "89197628803",
				},
			},
			want:    nil,
			wantErr: assert.Error,
		},
		{
			name: "valid",
			fields: fields{
				repo:   NewMockIUserRepo(ctrl),
				secret: secret,
				expect: func(repo *MockIUserRepo) {
					repo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(expectedUser, nil)
				},
			},
			args: args{
				ctx: context.Background(),
				newUser: &dto.NewUser{
					Login:    "user1",
					Password: "Pass123!",
					Phone:    "89197628803",
				},
			},
			want:    expectedUser,
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &AuthService{
				repo:   tt.fields.repo,
				secret: tt.fields.secret,
			}

			tt.fields.expect(tt.fields.repo)

			got, err := r.SignUp(tt.args.ctx, tt.args.newUser)
			if !tt.wantErr(t, err, fmt.Sprintf("SignUp(%v, %v)", tt.args.ctx, tt.args.newUser)) {
				return
			}

			assert.Equalf(t, tt.want, got, "SignUp(%v, %v)", tt.args.ctx, tt.args.newUser)
		})
	}
}

func TestAuthService_Login(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	foundUser := &user_dto.User{
		Uuid:         "123e4567-e89b-12d3-a456-426655440000",
		Phone:        "89197628803",
		Login:        "user1",
		PasswordHash: hash,
	}
	expectedJWT := &dto.JWT{
		Session: &session.Session{
			UserUuid: "123e4567-e89b-12d3-a456-426655440000",
		},
	}

	type fields struct {
		repo   *MockIUserRepo
		secret string
		expect func(repo *MockIUserRepo)
	}
	type args struct {
		ctx             context.Context
		userCredentials *dto.UserCredentials
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *dto.JWT
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "invalid login",
			fields: fields{
				repo:   NewMockIUserRepo(ctrl),
				secret: secret,
				expect: func(repo *MockIUserRepo) {
					repo.EXPECT().GetByLogin(gomock.Any(), "user2").Return(nil, err_const.ErrDatabaseRecordNotFound)
				},
			},
			args: args{
				ctx: context.Background(),
				userCredentials: &dto.UserCredentials{
					Login:    "user2",
					Password: "pass",
				},
			},
			want:    nil,
			wantErr: assert.Error,
		},
		{
			name: "invalid password",
			fields: fields{
				repo:   NewMockIUserRepo(ctrl),
				secret: secret,
				expect: func(repo *MockIUserRepo) {
					repo.EXPECT().GetByLogin(gomock.Any(), "user1").Return(foundUser, nil)
				},
			},
			args: args{
				ctx: context.Background(),
				userCredentials: &dto.UserCredentials{
					Login:    "user1",
					Password: "invalid",
				},
			},
			want:    nil,
			wantErr: assert.Error,
		},
		{
			name: "valid",
			fields: fields{
				repo:   NewMockIUserRepo(ctrl),
				secret: secret,
				expect: func(repo *MockIUserRepo) {
					repo.EXPECT().GetByLogin(gomock.Any(), "user1").Return(foundUser, nil)
				},
			},
			args: args{
				ctx: context.Background(),
				userCredentials: &dto.UserCredentials{
					Login:    "user1",
					Password: "Pass123!",
				},
			},
			want:    expectedJWT,
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &AuthService{
				repo:   tt.fields.repo,
				secret: tt.fields.secret,
			}
			tt.fields.expect(tt.fields.repo)

			got, err := r.Login(tt.args.ctx, tt.args.userCredentials)
			if !tt.wantErr(t, err, fmt.Sprintf("Login(%v, %v)", tt.args.ctx, tt.args.userCredentials)) {
				return
			}

			if tt.want != nil && got != nil {
				assert.Equalf(t, tt.want.Session.UserUuid, got.Session.UserUuid, "Login(%v, %v)", tt.args.ctx, tt.args.userCredentials)
			}
		})
	}
}

func TestAuthService_AuthAccessToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expectedSession := &session.Session{
		SID:      SID,
		UserUuid: UserUuid,
	}

	type fields struct {
		repo   *MockIUserRepo
		secret string
		expect func(repo *MockIUserRepo)
	}
	type args struct {
		in0         context.Context
		accessToken string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *session.Session
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "invalid token",
			fields: fields{
				repo:   NewMockIUserRepo(ctrl),
				secret: secret,
			},
			args: args{
				in0:         context.Background(),
				accessToken: invalidToken,
			},
			want:    nil,
			wantErr: assert.Error,
		},
		{
			name: "expired token",
			fields: fields{
				repo:   NewMockIUserRepo(ctrl),
				secret: secret,
			},
			args: args{
				in0:         context.Background(),
				accessToken: expiredToken,
			},
			want:    nil,
			wantErr: assert.Error,
		},
		{
			name: "valid",
			fields: fields{
				repo:   NewMockIUserRepo(ctrl),
				secret: secret,
			},
			args: args{
				in0:         context.Background(),
				accessToken: validToken,
			},
			want:    expectedSession,
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &AuthService{
				repo:   tt.fields.repo,
				secret: tt.fields.secret,
			}
			got, err := r.AuthAccessToken(tt.args.in0, tt.args.accessToken)
			if !tt.wantErr(t, err, fmt.Sprintf("AuthAccessToken(%v, %v)", tt.args.in0, tt.args.accessToken)) {
				return
			}
			assert.Equalf(t, tt.want, got, "AuthAccessToken(%v, %v)", tt.args.in0, tt.args.accessToken)
		})
	}
}

func TestAuthService_RefreshAccessToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		repo   IUserRepo
		secret string
	}
	type args struct {
		in0          context.Context
		refreshToken string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *dto.Tokens
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "invalid refresh token",
			fields: fields{
				repo:   NewMockIUserRepo(ctrl),
				secret: secret,
			},
			args: args{
				in0:          context.Background(),
				refreshToken: invalidToken,
			},
			want:    nil,
			wantErr: assert.Error,
		},
		{
			name: "expired refresh token",
			fields: fields{
				repo:   NewMockIUserRepo(ctrl),
				secret: secret,
			},
			args: args{
				in0:          context.Background(),
				refreshToken: expiredToken,
			},
			want:    nil,
			wantErr: assert.Error,
		},
		{
			name: "valid",
			fields: fields{
				repo:   NewMockIUserRepo(ctrl),
				secret: secret,
			},
			args: args{
				in0:          context.Background(),
				refreshToken: validToken,
			},
			want:    nil,
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &AuthService{
				repo:   tt.fields.repo,
				secret: tt.fields.secret,
			}
			got, err := r.RefreshAccessToken(tt.args.in0, tt.args.refreshToken)
			if !tt.wantErr(t, err, fmt.Sprintf("RefreshAccessToken(%v, %v)", tt.args.in0, tt.args.refreshToken)) {
				return
			}

			_ = got
		})
	}
}
