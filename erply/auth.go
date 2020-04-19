package erply

import (
	"github.com/breathbath/erplyapi/cache"
	"github.com/breathbath/go_utils/utils/env"
	log "github.com/sirupsen/logrus"
	"time"
)

//Session holds Erply Session data
type Session struct {
	sessionID string
}

//IsValid checks if session is valid for another API call
func (s Session) IsValid() bool {
	return s.sessionID != ""
}

//AuthProvider retrieves session data from auth env vars
type AuthProvider struct {
	curSession  Session
	CacheClient cache.Client
	AuthFunc func (username string, password string, clientCode string) (string, error)
}

//GetSession reads session either from memory or cache or calls API to establish a new one
func (ap *AuthProvider) GetSession() (s Session, err error) {
	log.Debug("Will read session data")
	if ap.curSession.IsValid() {
		log.Debug("Read session data from memory cache")
		return ap.curSession, nil
	}

	log.Debug("Will read session data from cache")
	s, err = ap.getSessionFromCache()
	if err != nil {
		return
	}

	if s.IsValid() {
		log.Debug("Session exists in cache will return it")
		ap.curSession = s
		return
	}

	sessionKey, err := ap.getAuthUserFromAPI()
	if err != nil {
		return
	}

	err = ap.storeSessionInCache(sessionKey)
	if err != nil {
		return
	}

	ap.curSession = Session{sessionID: sessionKey}

	return ap.curSession, nil
}

func (ap *AuthProvider) getSessionFromCache() (sess Session, err error) {
	var sessionID string
	found, err := ap.CacheClient.Read(sessionCacheKey, &sessionID)
	if err != nil {
		return
	}

	if found {
		sess = Session{sessionID: sessionID}
	}

	return
}

func (ap *AuthProvider) getAuthUserFromAPI() (sessionKey string, err error) {
	clientCode, err := env.ReadEnvOrError("ERPLY_CLIENT_CODE")
	if err != nil {
		return
	}

	username, err := env.ReadEnvOrError("ERPLY_USERNAME")
	if err != nil {
		return
	}

	pass, err := env.ReadEnvOrError("ERPLY_PASS")
	if err != nil {
		return
	}
	log.Debug("Will read session from API")
	sessionKey, err = ap.AuthFunc(username, pass, clientCode)
	if err == nil {
		log.Debug("Successfully read session from API")
	}
	return
}

func (ap *AuthProvider) storeSessionInCache(sessionKey string) (err error) {
	timeout := time.Second * time.Duration(env.ReadEnvInt("ERPLY_DATA_CACHE_TTL_SECONDS", 0))

	err = ap.CacheClient.Store(sessionCacheKey, sessionKey, timeout)
	return
}
