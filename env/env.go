package env

import "os"

type Env string

const (
	local Env = "local"
	cloud Env = "cloud"
	empty Env = ""
)

var env Env

func Init() {
	e := Env(os.Getenv("ENV"))

	env = e
}

func Get() Env {
	return env
}

func (e Env) String() string {
	return string(e)
}

func (e Env) IsLocal() bool {
	return e == local
}
