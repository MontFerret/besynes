package execution

import "runtime"

type Settings struct {
	PoolSize int
}

func NewDefaultSettings() Settings {
	return Settings{
		PoolSize: runtime.NumCPU() * 20,
	}
}
