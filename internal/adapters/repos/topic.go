package repos

import (
	"errors"
	"github.com/orewaee/ticket-box/internal/app/domain"
)

type TopicRepo struct {
	topics []*domain.Topic
}

func NewTopicRepo() *TopicRepo {
	return &TopicRepo{
		topics: make([]*domain.Topic, 0),
	}
}

func (repo *TopicRepo) CreateTopic(topic *domain.Topic) error {
	repo.topics = append(repo.topics, topic)
	return nil
}

func (repo *TopicRepo) ReadTopicById(id string) (*domain.Topic, error) {
	for _, topic := range repo.topics {
		if topic.Id == id {
			return topic, nil
		}
	}

	return nil, errors.New("topic not found")
}

func (repo *TopicRepo) ReadTopicsByGuildId(guildId string) ([]*domain.Topic, error) {
	topics := make([]*domain.Topic, 0)
	for _, topic := range repo.topics {
		if topic.GuildId == guildId {
			topics = append(topics, topic)
		}
	}

	return topics, nil
}

func (repo *TopicRepo) DeleteTopicById(id string) error {
	for i, topic := range repo.topics {
		if topic.Id == id {
			repo.topics = append(repo.topics[:i], repo.topics[i+1:]...)
			return nil
		}
	}

	return errors.New("topic not found")
}
