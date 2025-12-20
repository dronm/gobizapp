package services

import (
	"errors"
	"fmt"
	"strings"

	"github.com/dronm/gobizapp/notif"

	"github.com/jackc/pgx/v5"
)

// Здесь собраны функции отправки различных сообщений
// специфичные для проекта
var stekloNotifier *notif.Notifier

func InitNotifier(host, appName, pwd string) {
	stekloNotifier = notif.NewNotifier(host, appName, pwd)
}

func NotifSend(batch []*notif.NotifMessage) error {
	if stekloNotifier == nil {
		return fmt.Errorf("stekloNotifier is not initialized")
	}
	resp, err := stekloNotifier.Send(batch)
	if err != nil {
		return err
	}

	var e strings.Builder
	for _, r := range resp {
		if r.Error != "" {
			if e.Len() > 0 {
				e.WriteString(", ")
			}
			e.WriteString(fmt.Sprintf("Error sending message ID: %d, text: %s", r.ID, r.Error))
		}
	}
	if e.Len() > 0 {
		return errors.New(e.String())
	}
	return nil
}

func RecoverPasswordEmail(conn *pgx.Conn, userId int, url string, newPassword string) error {
	msg, err := notif.NewEmailMessageFromSQL(conn, "email_recover_pwd", nil, nil, []any{userId, url, newPassword})
	if err != nil {
		return fmt.Errorf("notif.NewEmailMessageFromSQL() failed: %v", err)
	}

	notifMsg := notif.NewNotifMessage("confirm", notif.PROV_EMAIL)
	notifMsg.AddEmailMessage(msg)
	return NotifSend([]*notif.NotifMessage{notifMsg})
}

func CorrectTelForSMS(tel *string) {
	if len(*tel) == 10 {
		*tel = "7" + *tel
	}
}

