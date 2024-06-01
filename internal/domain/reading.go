package domain

// Entidades de negocio principales de la app

type Reading struct {
	DeviceID  string
	Timestamp int64
	Value     float64
}
