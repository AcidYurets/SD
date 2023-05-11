package roles

import (
	"calend/internal/utils/slice"
	"context"
	"fmt"
)

type Type string

// Допустимые поли
const (
	SimpleUser  Type = "simple_user"  // Простой пользователь
	PremiumUser Type = "premium_user" // Премиум пользователь
	Admin       Type = "admin"        // Администратор
)

func (s Type) String() string {
	return string(s)
}

func (s Type) IsValid() bool {
	return s.isValid()
}

func (s Type) Validate() error {
	if !s.isValid() {
		return fmt.Errorf("некорректное значиение типа roles.Type")
	}

	return nil
}

func (s Type) isValid() bool {
	if slice.Contains([]Type{SimpleUser, PremiumUser, Admin}, s) {
		return true
	}

	return false
}

type needChangeKey struct{}

func CheckNeedChange(ctx context.Context) bool {
	changeNeeded, ok := ctx.Value(needChangeKey{}).(bool)
	return ok && changeNeeded
}

func SetNeedChange(ctx context.Context) context.Context {
	return context.WithValue(ctx, needChangeKey{}, true)
}

type useSuperUserKey struct{}

func CheckUseSuperUser(ctx context.Context) bool {
	useSuperuser, ok := ctx.Value(useSuperUserKey{}).(bool)
	return ok && useSuperuser
}

func SetUseSuperUser(ctx context.Context) context.Context {
	return context.WithValue(ctx, useSuperUserKey{}, true)
}
