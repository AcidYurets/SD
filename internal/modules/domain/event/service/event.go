package service

import (
	"calend/internal/models/access"
	"calend/internal/models/err_const"
	"calend/internal/models/roles"
	"calend/internal/models/session"
	"calend/internal/modules/domain/event/dto"
	"calend/internal/modules/domain/event/elastic"
	tag_dto "calend/internal/modules/domain/tag/dto"
	"calend/internal/modules/logger"
	"context"
	"fmt"
	"go.uber.org/zap"
	"strings"
)

//go:generate mockgen -destination mock_event_test.go -package service . IEventRepo
//go:generate mockgen -destination mock_invitation_test.go -package service . IInvitationRepo

type IEventRepo interface {
	// GetByUuid получение по uuid события вместе со связанными сущностями
	GetByUuid(ctx context.Context, uuid string) (*dto.Event, error)
	// ListTagsByEventUuid получение тегов по uuid события
	ListTagsByEventUuid(ctx context.Context, uuid string) (tag_dto.Tags, error)
	// ListTagsByEventUuids получение тегов по uuids событий
	ListTagsByEventUuids(ctx context.Context, uuids []string) (tag_dto.Tags, error)
	// GetCheckingInfoByUuid получение по uuid события только необходимыми для проверки прав доступа полями
	GetCheckingInfoByUuid(ctx context.Context, uuid string) (*dto.Event, error)
	// ListAvailable ищет все доступные пользователю события вместе со связанными сущностями, т.е.
	//  1. события, которые он создал;
	//  2. события, к которым он приглашен.
	ListAvailable(ctx context.Context, userUuid string) (dto.Events, error)
	// Create создает событие без приглашений, возвращает событие без связанных сущностей
	Create(ctx context.Context, dtm *dto.CreateEvent) (*dto.Event, error)
	// Update обновляет событие не изменяя приглашения, возвращает событие без связанных сущностей
	Update(ctx context.Context, uuid string, dtm *dto.UpdateEvent) (*dto.Event, error)
	// Delete удаляет событие, не удаляя его приглашения
	Delete(ctx context.Context, uuid string) error
}

type IInvitationRepo interface {
	// CreateBulk разом создает несколько приглашений
	CreateBulk(ctx context.Context, dtms dto.CreateInvitations) (dto.Invitations, error)
	// DeleteByEventUuid удаляет все приглашения определенного события
	DeleteByEventUuid(ctx context.Context, eventUuid string) (int, error)
}

type EventService struct {
	eventRepo IEventRepo
	invRepo   IInvitationRepo

	elasticService *elastic.EventElasticService
}

func NewEventService(eRepo IEventRepo, iRepo IInvitationRepo, elasticService *elastic.EventElasticService) *EventService {
	return &EventService{
		eventRepo:      eRepo,
		invRepo:        iRepo,
		elasticService: elasticService,
	}
}

func (r *EventService) GetByUuid(ctx context.Context, uuid string) (*dto.Event, error) {
	if err := r.checkAvailable(ctx, uuid, access.ReadAccess); err != nil {
		return nil, err
	}

	return r.eventRepo.GetByUuid(ctx, uuid)
}

func (r *EventService) ListTagsByEventUuid(ctx context.Context, uuid string) (tag_dto.Tags, error) {
	return r.eventRepo.ListTagsByEventUuid(ctx, uuid)
}

func (r *EventService) ListTagsByEventUuids(ctx context.Context, uuids []string) (tag_dto.Tags, error) {
	return r.eventRepo.ListTagsByEventUuids(ctx, uuids)
}

// ListAvailable ищет все доступные пользователю события, т.е.
//  1. события, которые он создал;
//  2. события, к которым он приглашен.
func (r *EventService) ListAvailable(ctx context.Context) (dto.Events, error) {
	userUuid, err := session.GetUserUuidFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	return r.eventRepo.ListAvailable(ctx, userUuid)
}

func (r *EventService) Create(ctx context.Context, newEvent *dto.CreateEvent, newInvs dto.CreateInvitations) (*dto.Event, error) {
	userUuid, err := session.GetUserUuidFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	// Устанавливает создателя как текущего пользователя
	newEvent.CreatorUuid = userUuid

	// Создаем событие без приглашений
	createdEvent, err := r.eventRepo.Create(ctx, newEvent)
	if err != nil {
		return nil, fmt.Errorf("ошибка при создании события: %w", err)
	}

	for _, inv := range newInvs {
		inv.EventUuid = createdEvent.Uuid
	}

	// Создаем приглашения для события
	_, err = r.invRepo.CreateBulk(ctx, newInvs)
	if err != nil {
		return nil, fmt.Errorf("ошибка при создании приглашений: %w", err)
	}

	// Получаем событие с приглашениями
	eventWithInvs, err := r.eventRepo.GetByUuid(ctx, createdEvent.Uuid)
	if err != nil {
		return nil, fmt.Errorf("ошибка при получении события: %w", err)
	}

	// Переиндексируем эластик
	_, err = r.elasticService.ReindexByUuids(ctx, createdEvent.Uuid)
	if err != nil {
		return nil, err
	}

	return eventWithInvs, nil
}

func (r *EventService) AddInvitations(ctx context.Context, uuid string, newInvs dto.CreateInvitations) (*dto.Event, error) {
	if err := r.checkAvailable(ctx, uuid, access.InviteAccess); err != nil {
		return nil, err
	}

	for _, inv := range newInvs {
		inv.EventUuid = uuid
	}

	// Создаем приглашения для события
	if _, err := r.invRepo.CreateBulk(ctx, newInvs); err != nil {
		return nil, fmt.Errorf("ошибка при создании приглашений: %w", err)
	}

	// Получаем событие с приглашениями
	eventWithInvs, err := r.eventRepo.GetByUuid(ctx, uuid)
	if err != nil {
		return nil, fmt.Errorf("ошибка при получении события: %w", err)
	}

	// Переиндексируем эластик
	_, err = r.elasticService.ReindexByUuids(ctx, uuid)
	if err != nil {
		return nil, err
	}

	return eventWithInvs, nil
}

func (r *EventService) Update(ctx context.Context, uuid string, updEvent *dto.UpdateEvent, newInvs dto.CreateInvitations) (*dto.Event, error) {
	if err := r.checkAvailable(ctx, uuid, access.UpdateAccess); err != nil {
		return nil, err
	}

	// Обновляем событие без приглашений
	if _, err := r.eventRepo.Update(ctx, uuid, updEvent); err != nil {
		return nil, fmt.Errorf("ошибка при обновлении события: %w", err)
	}

	// Удаляем все приглашения события
	if _, err := r.invRepo.DeleteByEventUuid(ctx, uuid); err != nil {
		return nil, fmt.Errorf("ошибка при удалении приглашений: %w", err)
	}

	for _, inv := range newInvs {
		inv.EventUuid = uuid
	}

	// Создаем обновленные приглашения для события
	if _, err := r.invRepo.CreateBulk(ctx, newInvs); err != nil {
		return nil, fmt.Errorf("ошибка при создании приглашений: %w", err)
	}

	// Получаем событие с приглашениями
	eventWithInvs, err := r.eventRepo.GetByUuid(ctx, uuid)
	if err != nil {
		return nil, fmt.Errorf("ошибка при получении события: %w", err)
	}

	// Переиндексируем эластик
	_, err = r.elasticService.ReindexByUuids(ctx, uuid)
	if err != nil {
		return nil, err
	}

	return eventWithInvs, nil
}

func (r *EventService) Delete(ctx context.Context, uuid string) error {
	if err := r.checkAvailable(ctx, uuid, access.DeleteAccess); err != nil {
		return err
	}

	// Удаляем все приглашения события
	if _, err := r.invRepo.DeleteByEventUuid(ctx, uuid); err != nil {
		return fmt.Errorf("ошибка при удалении приглашений: %w", err)
	}

	// Удаляем событие
	if err := r.eventRepo.Delete(ctx, uuid); err != nil {
		return fmt.Errorf("ошибка при удалении события: %w", err)
	}

	// Переиндексируем эластик
	_, err := r.elasticService.ReindexByUuids(ctx, uuid)
	if err != nil {
		return err
	}

	return nil
}

func (r *EventService) checkAvailable(ctx context.Context, eventUuid string, opCode access.Type) error {
	ss, ok := session.GetSessionFromCtx(ctx)
	if !ok {
		return fmt.Errorf("ошибка при получении сессии")
	}

	// Если пользователь -- админ, то у него полный доступ
	if ss.Role == roles.Admin {
		return nil
	}

	userUuid := ss.UserUuid

	event, err := r.eventRepo.GetCheckingInfoByUuid(ctx, eventUuid)
	if err != nil {
		return fmt.Errorf("ошибка при получении события: %w", err)
	}

	// Проверяем, присутствуют ли в событии необходимые поля о создателе
	if event.CreatorUuid == "" {
		return err_const.ErrMissingRequiredFields
	}

	// Если текущий пользователь -- создатель, то у него полный доступ
	if event.CreatorUuid == userUuid {
		return nil
	}

	for _, inv := range event.Invitations {
		// Проверяем, присутствуют ли в приглашении необходимые поля о пользователе и праве доступа
		if inv == nil || inv.UserUuid == "" || inv.AccessRightCode == "" {
			return err_const.ErrMissingRequiredFields
		}

		// Если пользователь приглашен к событию, то проверяем его права доступа
		if inv.UserUuid == userUuid {
			// Если есть необходимое право
			if strings.Contains(inv.AccessRightCode.String(), string(opCode)) {
				return nil
			}
		}
	}

	logger.GetFromCtx(ctx).Warn("недостаточно прав",
		zap.String("event_uuid", eventUuid),
		zap.String("operation_code", string(opCode)),
	)
	return fmt.Errorf("%w: код операции = <%s>", err_const.ErrAccessDenied, opCode)
}
