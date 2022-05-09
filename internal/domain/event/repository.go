package event

import "fmt"

type Repository interface {
	FindAll() ([]Event, error)
	FindOne(id int64) (*Event, error)
}

const EventsCount int64 = 10

type repository struct {
	// Some internal data
}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) FindAll() ([]Event, error) {
	events := make([]Event, EventsCount)
	for i := int64(0); i < EventsCount; i++ {
		events[i] = Event{
			Id:   i + 1,
			Name: fmt.Sprintf("Event #%d", i+1),
		}
	}
	return events, nil
}

func (r *repository) FindOne(id int64) (*Event, error) {
	if id <= EventsCount {
		return &Event{
			Id:   id,
			Name: fmt.Sprintf("Event #%d", id),
		}, nil
	} else {
		return nil, nil
	}
}
