package ports

import "github.com/orewaee/ticket-box/internal/app/domain"

// UpdateTopicById(id string, topic *domain.Topic) error

type TopicRepo interface {
	CreateTopic(topic *domain.Topic) error
	ReadTopicById(id string) (*domain.Topic, error)
	ReadTopicsByGuildId(guildId string) ([]*domain.Topic, error)
	UpdateTopicById(id, emoji, name, description string) error
	DeleteTopicById(id string) error
}

// UpdateTopicById(id string, topic *domain.Topic) error

type TopicService interface {
	AddTopic(topic *domain.Topic) error
	GetTopicById(id string) (*domain.Topic, error)
	GetTopicsByGuildId(guildId string) ([]*domain.Topic, error)
	SetTopicEmoji(id, emoji string) error
	SetTopicName(id, name string) error
	SetTopicDescription(id, description string) error
	RemoveTopicById(id string) error
}
