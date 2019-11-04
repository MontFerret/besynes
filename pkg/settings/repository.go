package settings

type Repository interface {
	Get() (SettingsDetails, error)
	Exists() (bool, error)
	Create(settings SettingsDetails) error
	Update(settings SettingsDetails) error
}
