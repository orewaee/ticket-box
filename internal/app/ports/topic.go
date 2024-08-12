package ports

import "github.com/orewaee/ticket-box/internal/app/domain"

// UpdateTopicById(id string, topic *domain.Topic) error

type TopicRepo interface {
	CreateTopic(topic *domain.Topic) error
	ReadTopicById(id string) (*domain.Topic, error)
	ReadTopicsByGuildId(guildId string) ([]*domain.Topic, error)
	DeleteTopicById(id string) error
}

// UpdateTopicById(id string, topic *domain.Topic) error

type TopicService interface {
	AddTopic(topic *domain.Topic) error
	GetTopicById(id string) (*domain.Topic, error)
	GetTopicsByGuildId(guildId string) ([]*domain.Topic, error)
	RemoveTopicById(id string) error
}
