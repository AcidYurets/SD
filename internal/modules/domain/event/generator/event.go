package generator

import (
	"calend/internal/models/access"
	dto2 "calend/internal/modules/domain/access_right/dto"
	"calend/internal/modules/domain/event/dto"
	"calend/internal/modules/domain/event/elastic"
	tag_dto "calend/internal/modules/domain/tag/dto"
	user_dto "calend/internal/modules/domain/user/dto"
	"calend/internal/utils/ptr"
	"calend/internal/utils/random"
	"context"
	"fmt"
	"time"
)

type IEventRepo interface {
	// Create создает событие без приглашений, возвращает событие без связанных сущностей
	Create(ctx context.Context, dtm *dto.CreateEvent) (*dto.Event, error)
}

type IInvitationRepo interface {
	// CreateBulk разом создает несколько приглашений
	CreateBulk(ctx context.Context, dtms dto.CreateInvitations) (dto.Invitations, error)
}

type IUserRepo interface {
	// List возвращает список пользователей
	List(ctx context.Context) (user_dto.Users, error)
}

type IAccessRightRepo interface {
	// List возвращает список прав доступа
	List(ctx context.Context) (dto2.AccessRights, error)
}

type ITagsRepo interface {
	// List возвращает список тегов
	List(ctx context.Context) (tag_dto.Tags, error)
}

type EventGenerator struct {
	eventRepo IEventRepo
	invRepo   IInvitationRepo
	userRepo  IUserRepo
	arRepo    IAccessRightRepo
	tagRepo   ITagsRepo

	elasticService *elastic.EventElasticService
}

func NewEventGenerator(
	eRepo IEventRepo,
	iRepo IInvitationRepo,
	uRepo IUserRepo,
	arRepo IAccessRightRepo,
	tRepo ITagsRepo,
	elasticService *elastic.EventElasticService,
) *EventGenerator {

	return &EventGenerator{
		eventRepo: eRepo,
		invRepo:   iRepo,
		userRepo:  uRepo,
		arRepo:    arRepo,
		tagRepo:   tRepo,

		elasticService: elasticService,
	}
}

func (r *EventGenerator) Generate(ctx context.Context, count uint) error {
	createdEventUuids := make([]string, 0)

	for i := uint(0); i < count; i++ {
		newEvent, newInvs, err := r.generateEventWithInvitations(ctx, i)
		if err != nil {
			return err
		}

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

		createdEventUuids = append(createdEventUuids, createdEvent.Uuid)
	}

	// Переиндексируем эластик
	_, err := r.elasticService.ReindexByUuids(ctx, createdEventUuids...)
	if err != nil {
		return err
	}

	return nil
}

func (r *EventGenerator) generateEventWithInvitations(ctx context.Context, num uint) (*dto.CreateEvent, dto.CreateInvitations, error) {
	userUuids, err := r.usersUuids(ctx)
	if err != nil {
		return nil, nil, err
	}
	arCodes, err := r.accessRightCodes(ctx)
	if err != nil {
		return nil, nil, err
	}
	tagsUuids, err := r.tagsUuids(ctx)
	if err != nil {
		return nil, nil, err
	}
	eventTypes := []string{"Встреча", "День рождения", "Праздник", "Другое"}

	creatorUuid, userUuids := random.FromSliceWithRemove(userUuids)

	invsCount := random.IntRange(0, 7)
	newInvs := make(dto.CreateInvitations, invsCount)

	for i := 0; i < invsCount; i++ {
		var userUuid string
		userUuid, userUuids = random.FromSliceWithRemove(userUuids)

		newInvs[i] = &dto.CreateInvitation{
			UserUuid:        userUuid,
			AccessRightCode: random.FromSlice(arCodes),
		}
	}

	newEvent := &dto.CreateEvent{
		Timestamp:   random.TimestampRange(time.Now(), time.Now().AddDate(0, 0, 7)),
		Name:        fmt.Sprintf("Событие%d", num),
		Description: ptr.String(fmt.Sprintf("Описание события %d", num)),
		Type:        random.FromSlice(eventTypes),
		IsWholeDay:  random.Bool(),
		TagUuids:    random.NFromSlice(tagsUuids, random.IntRange(1, 5)),
		CreatorUuid: creatorUuid,
	}

	return newEvent, newInvs, nil
}

func (r *EventGenerator) usersUuids(ctx context.Context) ([]string, error) {
	users, err := r.userRepo.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("ошибка при получении пользователей: %w", err)
	}

	usersUuids := make([]string, len(users))
	for i, user := range users {
		usersUuids[i] = user.Uuid
	}

	return usersUuids, nil
}

func (r *EventGenerator) accessRightCodes(ctx context.Context) ([]access.Type, error) {
	ars, err := r.arRepo.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("ошибка при получении прав доступа: %w", err)
	}

	arsCodes := make([]access.Type, len(ars))
	for i, ar := range ars {
		arsCodes[i] = ar.Code
	}

	return arsCodes, nil
}

func (r *EventGenerator) tagsUuids(ctx context.Context) ([]string, error) {
	tags, err := r.tagRepo.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("ошибка при получении тегов: %w", err)
	}

	tagsUuids := make([]string, len(tags))
	for i, tag := range tags {
		tagsUuids[i] = tag.Uuid
	}

	return tagsUuids, nil
}
