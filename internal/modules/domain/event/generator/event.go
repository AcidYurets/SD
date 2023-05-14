package generator

import (
	"calend/internal/models/session"
	"calend/internal/modules/domain/event/dto"
	"calend/internal/modules/domain/event/elastic"
	"context"
	"fmt"
)

type IEventRepo interface {
	// Create создает событие без приглашений, возвращает событие без связанных сущностей
	Create(ctx context.Context, dtm *dto.CreateEvent) (*dto.Event, error)
}

type IInvitationRepo interface {
	// CreateBulk разом создает несколько приглашений
	CreateBulk(ctx context.Context, dtms dto.CreateInvitations) (dto.Invitations, error)
}

type EventGenerator struct {
	eventRepo IEventRepo
	invRepo   IInvitationRepo

	elasticService *elastic.EventElasticService
}

func NewEventGenerator(eRepo IEventRepo, iRepo IInvitationRepo, elasticService *elastic.EventElasticService) *EventGenerator {
	return &EventGenerator{
		eventRepo:      eRepo,
		invRepo:        iRepo,
		elasticService: elasticService,
	}
}

func (r *EventGenerator) Generate(ctx context.Context) error {
	newEvent, newInvs, err := r.generateEventWithInvitations()
	if err != nil {
		return err
	}

	userUuid, err := session.GetUserUuidFromCtx(ctx)
	if err != nil {
		return err
	}

	// Устанавливает создателя как текущего пользователя
	newEvent.CreatorUuid = userUuid

	// Создаем событие без приглашений
	createdEvent, err := r.eventRepo.Create(ctx, newEvent)
	if err != nil {
		return fmt.Errorf("ошибка при создании события: %w", err)
	}

	for _, inv := range newInvs {
		inv.EventUuid = createdEvent.Uuid
	}

	// Создаем приглашения для события
	_, err = r.invRepo.CreateBulk(ctx, newInvs)
	if err != nil {
		return fmt.Errorf("ошибка при создании приглашений: %w", err)
	}

	// Переиндексируем эластик
	_, err = r.elasticService.ReindexByUuids(ctx, createdEvent.Uuid)
	if err != nil {
		return err
	}

	return nil
}

func (r *EventGenerator) generateEventWithInvitations() (*dto.CreateEvent, dto.CreateInvitations, error) {
	return nil, nil, nil
}
