package models

import (
	"github.com/mlflow/mlflow-go/pkg/entities"
)

// Input mapped from table <inputs>.
type Input struct {
	ID              string     `db:"input_uuid"                          gorm:"column:input_uuid;not null"`
	SourceType      string     `db:"source_type"                         gorm:"column:source_type;primaryKey"`
	SourceID        string     `db:"source_id"                           gorm:"column:source_id;primaryKey"`
	DestinationType string     `db:"destination_type"                    gorm:"column:destination_type;primaryKey"`
	DestinationID   string     `db:"destination_id"                      gorm:"column:destination_id;primaryKey"`
	Tags            []InputTag `gorm:"foreignKey:InputID;references:ID"`
	Dataset         Dataset    `gorm:"foreignKey:ID;references:SourceID"`
}

func (i *Input) ToEntity() *entities.DatasetInput {
	tags := make([]*entities.InputTag, 0, len(i.Tags))
	for _, tag := range i.Tags {
		tags = append(tags, tag.ToEntity())
	}

	return &entities.DatasetInput{
		Tags:    tags,
		Dataset: i.Dataset.ToEntity(),
	}
}
