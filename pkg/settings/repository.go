package settings

type Repository interface {
	Get() (Settings, error)
	Exists() (bool, error)
	Create(settings Settings) error
	Update(settings Settings) error
}
