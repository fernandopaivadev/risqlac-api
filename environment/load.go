package environment

import (
	"risqlac-api/types"
)

var env types.Env

func Load() {
	env.JWT_SECRET = "0293rjdik023o3uihfdrca2q0o938hrfcd902382c9epjrfcd0239jrf"
	env.DATABASE_FILE = "risqlac.db"
}

func Get() types.Env {
	return env
}
