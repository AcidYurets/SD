package repo

import (
	"calend/internal/modules/db"
	"calend/internal/modules/db/ent"
	"calend/internal/modules/db/schema"
	"calend/internal/modules/domain/tag/dto"
	"context"
)

type TagRepo struct {
	client *ent.Client
}

func NewTagRepo(client *ent.Client) *TagRepo {
	return &TagRepo{
		client: client,
	}
}

func (r *TagRepo) GetByUuid(ctx context.Context, uuid string) (*dto.Tag, error) {
	tag, err := r.client.Tag.Get(schema.SkipSoftDelete(ctx), uuid)
	if err != nil {
		return nil, db.WrapError(err)
	}

	return ToTagDTO(tag), nil
}

func (r *TagRepo) List(ctx context.Context) (dto.Tags, error) {
	tags, err := r.client.Tag.Query().All(ctx)
	if err != nil {
		return nil, db.WrapError(err)
	}

	return ToTagDTOs(tags), nil
}

func (r *TagRepo) Create(ctx context.Context, dtm *dto.CreateTag) (*dto.Tag, error) {
	tag, err := r.client.Tag.Create().
		SetName(dtm.Name).
		SetDescription(dtm.Description).
		Save(ctx)
	if err != nil {
		return nil, db.WrapError(err)
	}

	return ToTagDTO(tag), nil
}

func (r *TagRepo) Update(ctx context.Context, uuid string, dtm *dto.UpdateTag) (*dto.Tag, error) {
	tag, err := r.client.Tag.UpdateOneID(uuid).
		SetName(dtm.Name).
		SetDescription(dtm.Description).
		Save(ctx)
	if err != nil {
		return nil, db.WrapError(err)
	}

	return ToTagDTO(tag), nil
}

func (r *TagRepo) Delete(ctx context.Context, uuid string) error {
	err := r.client.Tag.DeleteOneID(uuid).Exec(ctx)
	if err != nil {
		return db.WrapError(err)
	}

	return nil
}

func (r *TagRepo) Restore(ctx context.Context, uuid string) (*dto.Tag, error) {
	tag, err := r.client.Tag.UpdateOneID(uuid).ClearDeletedAt().Save(ctx)
	if err != nil {
		return nil, db.WrapError(err)
	}

	return ToTagDTO(tag), nil
}

func ToTagDTO(model *ent.Tag) *dto.Tag {
	if model == nil {
		return nil
	}
	return &dto.Tag{
		Uuid:        model.ID,
		Name:        model.Name,
		Description: model.Description,
	}
}

func ToTagDTOs(models ent.Tags) dto.Tags {
	if models == nil {
		return nil
	}
	dtms := make(dto.Tags, len(models))
	for i := range models {
		dtms[i] = ToTagDTO(models[i])
	}
	return dtms
}
