package service

import (
	"calend/internal/models/err_const"
	"calend/internal/modules/domain/event/dto"
	"context"
	"fmt"
)

//go:generate mockgen -destination mock_test.go -package service . IEventRepo

type IEventRepo interface {
	GetByUuid(ctx context.Context, uuid string) (*dto.Event, error)
	// Create создает событие без приглашений
	Create(ctx context.Context, dtm *dto.CreateEvent) (*dto.Event, error)
	Update(ctx context.Context, uuid string, dtm *dto.UpdateEvent) (*dto.Event, error)
	Delete(ctx context.Context, uuid string) error

	// ListAvailable ищет все доступные пользователю события, т.е.
	//  1. события, которые он создал;
	//  2. события, к которым он приглашен.
	ListAvailable(ctx context.Context, userUuid string) (dto.Events, error)
	// IsAvailable проверяет, доступно ли событие пользователю, т.е.
	//  1. создал ли он это событие?
	//  2. приглашен ли он к этому событию?
	IsAvailable(ctx context.Context, eventUuid string, userUuid string) (bool, error)
}

type IInvitationRepo interface {
	CreateBulk(ctx context.Context, dtms dto.CreateInvitations) (dto.Invitations, error)
	DeleteByEventUuid(ctx context.Context, eventUuid string) error
}

type EventService struct {
	eventRepo IEventRepo
	invRepo   IInvitationRepo
}

func NewEventService(eRepo IEventRepo, iRepo IInvitationRepo) *EventService {
	return &EventService{
		eventRepo: eRepo,
		invRepo:   iRepo,
	}
}

func (r *EventService) GetByUuid(ctx context.Context, uuid string) (*dto.Event, error) {
	if err := r.checkAvailable(ctx, uuid); err != nil {
		return nil, err
	}

	return r.eventRepo.GetByUuid(ctx, uuid)
}

func (r *EventService) ListAvailable(ctx context.Context) (dto.Events, error) {
	// TODO: сессия из контекста
	userUuid := ""

	return r.eventRepo.ListAvailable(ctx, userUuid)
}

func (r *EventService) CreateWithInvitations(ctx context.Context, newEvent *dto.CreateEvent, newInvs dto.CreateEventInvitations) (*dto.Event, error) {
	// Создаем событие без приглашений
	createdEvent, err := r.eventRepo.Create(ctx, newEvent)
	if err != nil {
		return nil, fmt.Errorf("ошибка при создании события: %w", err)
	}

	var createInvs dto.CreateInvitations
	for _, inv := range newInvs {
		createInv := &dto.CreateInvitation{
			EventUuid:       createdEvent.Uuid,
			UserUuid:        inv.UserUuid,
			AccessRightCode: inv.AccessRightCode,
		}
		createInvs = append(createInvs, createInv)
	}

	// Создаем приглашения для события
	_, err = r.invRepo.CreateBulk(ctx, createInvs)
	if err != nil {
		return nil, fmt.Errorf("ошибка при создании приглашений: %w", err)
	}

	// Получаем событие с приглашениями
	eventWithInvs, err := r.eventRepo.GetByUuid(ctx, createdEvent.Uuid)
	if err != nil {
		return nil, fmt.Errorf("ошибка при получении события: %w", err)
	}

	return eventWithInvs, nil
}

func (r *EventService) Update(ctx context.Context, uuid string, updEvent *dto.UpdateEvent, newInvs dto.CreateEventInvitations) (*dto.Event, error) {
	if err := r.checkAvailable(ctx, uuid); err != nil {
		return nil, err
	}

	// Обновляем событие без приглашений
	if _, err := r.eventRepo.Update(ctx, uuid, updEvent); err != nil {
		return nil, fmt.Errorf("ошибка при обновлении события: %w", err)
	}

	// Удаляем все приглашения события
	if err := r.invRepo.DeleteByEventUuid(ctx, uuid); err != nil {
		return nil, fmt.Errorf("ошибка при удалении приглашений: %w", err)
	}

	var createInvs dto.CreateInvitations
	for _, inv := range newInvs {
		createInv := &dto.CreateInvitation{
			EventUuid:       uuid,
			UserUuid:        inv.UserUuid,
			AccessRightCode: inv.AccessRightCode,
		}
		createInvs = append(createInvs, createInv)
	}

	// Создаем обновленные приглашения для события
	if _, err := r.invRepo.CreateBulk(ctx, createInvs); err != nil {
		return nil, fmt.Errorf("ошибка при создании приглашений: %w", err)
	}

	// Получаем событие с приглашениями
	eventWithInvs, err := r.eventRepo.GetByUuid(ctx, uuid)
	if err != nil {
		return nil, fmt.Errorf("ошибка при получении события: %w", err)
	}

	return eventWithInvs, nil
}

func (r *EventService) Delete(ctx context.Context, uuid string) error {
	if err := r.checkAvailable(ctx, uuid); err != nil {
		return err
	}

	// Удаляем все приглашения события
	if err := r.invRepo.DeleteByEventUuid(ctx, uuid); err != nil {
		return fmt.Errorf("ошибка при удалении приглашений: %w", err)
	}

	// Удаляем событие
	if err := r.eventRepo.Delete(ctx, uuid); err != nil {
		return fmt.Errorf("ошибка при удалении события: %w", err)
	}

	return nil
}

func (r *EventService) checkAvailable(ctx context.Context, eventUuid string, opCode string) error {
	// TODO: сессия из контекста
	userUuid := ""

	event, err := r.eventRepo.GetByUuid(ctx, eventUuid)
	if err != nil {
		return err
	}

	for _, inv := range event.Invitations {
		// Если пользователь приглашен к событию -- проверяем его права доступа
		if inv.User.Uuid == userUuid {
			inv.AccessRight.Code
		}
	}

	available, err := r.eventRepo.IsAvailable(ctx, eventUuid, userUuid)
	if err != nil {
		return fmt.Errorf("ошибка при проверки доступности события: %w", err)
	}
	if !available {
		return fmt.Errorf("%w: код операции = <%s>", err_const.ErrAccessDenied, opCode)
	}

	return nil
}
