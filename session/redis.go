// Package session is a redis session implementation.
package session

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"github.com/dronm/session"
	_ "github.com/dronm/session/redis"

	"github.com/dronm/gobizapp/logger"
)

const SessCookieKey = "_s"

var (
	SessManager    *session.Manager
	//SessQueryParam = "token"
)

type SessConfiger interface {
	GetMaxLifeTime() int64
	GetMaxIdleTime() int64
	GetDestroyAllTime() string
}

type RedisConfiger interface {
	GetConnect() string
	GetNamespace() string
}

func Initialize(sessConf SessConfiger, redisConf RedisConfiger, lv string) error {
	var err error
	SessManager, err = session.NewManager("redis",
		sessConf.GetMaxLifeTime(),
		sessConf.GetMaxIdleTime(),
		sessConf.GetDestroyAllTime(),
		redisConf.GetConnect(),
		redisConf.GetNamespace(),
	)
	if err != nil {
		panic(fmt.Sprintf("session.NewManager: %v", err))
	}

	var sessLv session.LogLevel
	switch lv {
	case "debug":
		sessLv = session.LOG_LEVEL_DEBUG
	case "warn":
		sessLv = session.LOG_LEVEL_WARN
	default:
		sessLv = session.LOG_LEVEL_ERROR
	}

	lw := logger.NewLogWriter()
	SessManager.StartGC(lw, sessLv)

	return nil
}

func GenSessionID() (string, error) {
	sessionID := make([]byte, SessManager.GetSessionIDLen())
	_, err := rand.Read(sessionID)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(sessionID), nil
}
