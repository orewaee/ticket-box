package services

import (
	"github.com/orewaee/ticket-box/internal/app/domain"
	"github.com/orewaee/ticket-box/internal/app/ports"
	"log"
)

type TopicService struct {
	topicRepo ports.TopicRepo
}

func NewTopicService(topicRepo ports.TopicRepo) *TopicService {
	return &TopicService{
		topicRepo: topicRepo,
	}
}

func (service *TopicService) AddTopic(topic *domain.Topic) error {
	if err := service.topicRepo.CreateTopic(topic); err != nil {
		return err
	}

	log.Println("topic added:", topic)

	return nil
}

func (service *TopicService) GetTopicById(id string) (*domain.Topic, error) {
	topic, err := service.topicRepo.ReadTopicById(id)
	if err != nil {
		return nil, err
	}

	return topic, nil
}

func (service *TopicService) GetTopicsByGuildId(guildId string) ([]*domain.Topic, error) {
	topics, err := service.topicRepo.ReadTopicsByGuildId(guildId)
	if err != nil {
		return nil, err
	}

	return topics, nil
}

func (service *TopicService) RemoveTopicById(id string) error {
	if err := service.topicRepo.DeleteTopicById(id); err != nil {
		return err
	}

	log.Println("topic removed:", id)

	return nil
}
