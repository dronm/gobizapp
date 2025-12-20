package services

import (
	"context"
	"errors"

	"github.com/dronm/ds/pgds"
	"github.com/dronm/session"
)

var ErrEvHandlerNotDefined = errors.New("EventService: EvHandler not set")

type EventHandler interface {
	SubscribeToEvent(sessionID string, eventID string) error
	UnsubscribeFromEvent(sessionID string, eventID string) error
	PublishEvent(publisherID, eventID string, payload any) error
}
var EvHandler EventHandler

type EventService struct {
	DB      *pgds.PgProvider
	Session session.Session
	QueryID string
}

func (s *EventService) SetDB(db *pgds.PgProvider) {
	s.DB = db
}

func (s *EventService) SetSession(sess session.Session) {
	s.Session = sess
}

func (s *EventService) SetQueryID(queryID string) {
	s.QueryID = queryID
}

func NewEventService(db *pgds.PgProvider, sess session.Session) *EventService {
	return &EventService{DB: db, Session: sess}
}

func (s *EventService) Subscribe(ctx context.Context, events []string) error {
	if EvHandler == nil {
		return ErrEvHandlerNotDefined
	}
	sessID := s.Session.SessionID()
	for _, ev := range events {
		if err := EvHandler.SubscribeToEvent(sessID, ev); err != nil {
			return err
		}
	}
	return nil
}

func (s *EventService) Unsubscribe(ctx context.Context, events []string) error {
	if EvHandler == nil {
		return ErrEvHandlerNotDefined
	}
	
	sessID := s.Session.SessionID()
	for _, ev := range events {
		if err := EvHandler.UnsubscribeFromEvent(sessID, ev); err != nil {
			return err
		}
	}
	return nil
}
