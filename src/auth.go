package src

import (
	"os"
	"strings"
)

//Auth user authentication
type IAuth interface {
	init()
	//pre-shared key encrypted handshake
	SharedKey() (string, bool)
	User(string) (string, bool)
}

//NewAuthWithRds from redis
func NewAuthWithRds() *IAuth {
	return nil
}

//NewAuthWithEnv from env
func NewAuthWithEnv() *IAuth {
	return nil
}

var _ IAuth = &env{}

type env struct {
}

func (e *env) init() {

}

func (e *env) SharedKey() (key string, ok bool) {
	key, ok = os.LookupEnv("TRAFFIC_SHARED")
	return
}
func (e *env) User(uname string) (pwd string, ok bool) {
	pwd, ok = os.LookupEnv("TRAFFIC_USER_" + strings.ToUpper(uname))
	return
}
