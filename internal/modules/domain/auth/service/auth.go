package service

import (
	"calend/internal/models/err_const"
	"calend/internal/models/roles"
	"calend/internal/models/session"
	"calend/internal/modules/config"
	"calend/internal/modules/domain/auth/dto"
	user_dto "calend/internal/modules/domain/user/dto"
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"time"
	"unicode"
)

const AccessTokenLiveTime = 30000 // TODO: Уменьшить
const RefreshTokenLiveTime = 360

//go:generate mockgen -destination mock_test.go -package service . IUserRepo

type IUserRepo interface {
	GetByLogin(ctx context.Context, login string) (*user_dto.User, error)
	Create(ctx context.Context, dtm *user_dto.CreateUser) (*user_dto.User, error)
}

type AuthService struct {
	repo   IUserRepo
	secret string
}

func NewAuthService(repo IUserRepo, config config.Config) *AuthService {
	return &AuthService{
		repo:   repo,
		secret: config.Secret,
	}
}

// AuthAccessToken проверяет доступ к системе по токену доступа
func (r *AuthService) AuthAccessToken(_ context.Context, accessToken string) (*session.Session, error) {

	tokenClaims, err := r.parseToken(accessToken)
	if err != nil {
		return nil, err
	}

	if tokenClaims.Session == nil {
		return nil, fmt.Errorf("%w: сессия пользователя не найдена", err_const.ErrUnauthorized)
	}

	return tokenClaims.Session, nil
}

// RefreshAccessToken обновляет токен доступа по токену обновления
func (r *AuthService) RefreshAccessToken(_ context.Context, refreshToken string) (*dto.Tokens, error) {
	tokenClaims, err := r.parseToken(refreshToken)
	if err != nil {
		return nil, fmt.Errorf("%w: отсутствует токен или передан в некорректном формате", err_const.ErrUnauthorized)
	}

	tokenSession := tokenClaims.Session
	if tokenSession == nil {
		return nil, fmt.Errorf("%w: сессия пользователя не найдена", err_const.ErrUnauthorized)
	}

	if err != nil {
		return nil, err
	}

	expTime := time.Minute * time.Duration(AccessTokenLiveTime) // Срок жизни токена доступа
	// При использовании токена обновления срок жизни заново не продлеваем
	nbfTime := time.Until(tokenClaims.ExpiresAt.Time) // Срок жизни токена обновления

	// Генерируем пару токенов
	token, err := r.createTokenPair(tokenSession, expTime, nbfTime)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", err_const.ErrUnauthorized, err.Error()) //nolint:errorlint
	}

	return token, nil
}

// Login аутентификация и получение токенов и сессии
func (r *AuthService) Login(ctx context.Context, userCredentials *dto.UserCredentials) (*dto.JWT, error) {
	if userCredentials.Login == "" {
		return nil, fmt.Errorf("%w: некорректный логин", err_const.ErrUnauthorized)
	}

	currentUser, err := r.repo.GetByLogin(ctx, userCredentials.Login)
	if err != nil {
		return nil, fmt.Errorf("%w: пользователя с таким логином не существует", err_const.ErrUnauthorized)
	}

	// Валидация пароля
	if !r.validateUserPassword(userCredentials, currentUser) {
		return nil, fmt.Errorf("%w: неверный пароль", err_const.ErrUnauthorized)
	}

	ss := &session.Session{
		SID:      uuid.NewString(),
		UserUuid: currentUser.Uuid,
	}

	// Сроки жизни токенов
	expTime := time.Minute * time.Duration(AccessTokenLiveTime)  // Срок жизни токена доступа
	nbfTime := time.Minute * time.Duration(RefreshTokenLiveTime) // Срок жизни токена обновления

	token, err := r.createTokenPair(ss, expTime, nbfTime)
	if err != nil {
		return nil, err
	}

	return &dto.JWT{
		Tokens:  *token,
		Session: ss,
	}, nil
}

// SignUp регистрация в системе
func (r *AuthService) SignUp(ctx context.Context, newUser *dto.NewUser) (*user_dto.User, error) {
	err := r.passwordValidation(newUser.Password)
	if err != nil {
		return nil, err
	}

	// Формируем хэш пароля
	passwordHash, err := hashPassword(newUser.Password)
	if err != nil {
		return nil, err
	}

	// Новые пользователи получают роль простого пользователя
	role := roles.SimpleUser
	createUser := &user_dto.CreateUser{
		Phone:        newUser.Phone,
		Login:        newUser.Login,
		PasswordHash: passwordHash,
		Role:         role,
	}

	// Создаем пользователя
	createdUser, err := r.repo.Create(ctx, createUser)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

func (r *AuthService) parseToken(jwtToken string) (*dto.TokenClaims, error) {
	var claims dto.TokenClaims

	token, err := jwt.ParseWithClaims(jwtToken, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("%w: неожиданный метод подписи: %v", err_const.ErrInvalidToken, token.Header["alg"])
		}
		return []byte(r.secret), nil
	})
	if err != nil {
		return nil, fmt.Errorf("%w: %s", err_const.ErrInvalidToken, err)
	}

	if !token.Valid {
		return nil, err_const.ErrInvalidToken
	}

	if claims.Valid() != nil {
		return nil, err_const.ErrInvalidToken
	}

	return &claims, nil
}

// passwordValidation валидирует пароль
func (r *AuthService) passwordValidation(pass string) error {
	runes := []rune(pass)

	// минимум 8 символов
	if len(runes) < 8 {
		return fmt.Errorf("пароль не соответствует требованиям")
	}

	hasNumbers := false
	hasLower := false
	hasUpper := false
	hasSymbols := false

	symbols := map[rune]struct{}{
		'~': {}, '!': {}, '@': {}, '#': {}, '$': {}, '%': {},
		'^': {}, '&': {}, '*': {}, '(': {}, ')': {}, '_': {},
		'-': {}, '+': {}, '=': {}, '{': {}, '}': {}, '[': {},
		']': {}, '\\': {}, '|': {}, ':': {}, ';': {}, '"': {},
		'\'': {}, '<': {}, '>': {}, ',': {}, '.': {}, '?': {}, '/': {},
	}

	for _, r := range runes {
		// минимум одна цифра
		if !hasNumbers && unicode.IsNumber(r) {
			hasNumbers = true
		}
		// минимум одна прописная буква
		if !hasLower && unicode.IsLower(r) {
			hasLower = true
		}
		// минимум одна заглавная буква
		if !hasUpper && unicode.IsUpper(r) {
			hasUpper = true
		}
		// минимум один спец символ
		if !hasSymbols {
			if _, ok := symbols[r]; ok {
				hasSymbols = true
			}
		}
	}

	if !hasNumbers || !hasLower || !hasUpper || !hasSymbols {
		return fmt.Errorf("пароль не соответствует требованиям")
	}
	return nil
}

func (r *AuthService) validateUserPassword(userCredentials *dto.UserCredentials, user *user_dto.User) bool {
	return checkPasswordHash(userCredentials.Password, user.PasswordHash)
}

// createTokenPair создает пару токенов
// exp - срок действия токена доступа
// nbf - срок действия токена обновления
func (r *AuthService) createTokenPair(session *session.Session, exp time.Duration, nbf time.Duration) (*dto.Tokens, error) {
	// Определяем общий uuid для токенов
	tokenUuid := uuid.NewString()
	// Дата создания токена
	issuedAt := time.Now()

	accessTokenClaims := dto.TokenClaims{}
	accessTokenClaims.Session = session
	accessTokenClaims.IssuedAt = &jwt.NumericDate{Time: issuedAt}
	accessTokenClaims.ExpiresAt = &jwt.NumericDate{Time: issuedAt.Add(exp)}
	accessTokenClaims.ID = tokenUuid

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	accessTokenSign, err := accessToken.SignedString([]byte(r.secret))
	if err != nil {
		return nil, fmt.Errorf("%w: ошибка генерации токена доступа", err)
	}

	refreshTokenClaims := dto.TokenClaims{}
	refreshTokenClaims.Session = session
	refreshTokenClaims.IssuedAt = &jwt.NumericDate{Time: issuedAt}
	refreshTokenClaims.ExpiresAt = &jwt.NumericDate{Time: issuedAt.Add(nbf)}
	refreshTokenClaims.ID = tokenUuid

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	refreshTokenSign, err := refreshToken.SignedString([]byte(r.secret))
	if err != nil {
		return nil, fmt.Errorf("%w: ошибка генерации токена обновления", err)
	}

	return &dto.Tokens{
		AccessToken:  accessTokenSign,
		RefreshToken: refreshTokenSign,
	}, nil
}
