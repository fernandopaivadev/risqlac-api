package environment

import "risqlac-api/types"

func Get() types.EnvironmentVariables {
	return env
}
