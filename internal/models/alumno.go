package models

type Alumno struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Nombre    string `json:"nombre"`
	Apellido  string `json:"apellido"`
	TarjetaID string `json:"tarjeta_id"`
	HuellaID  string `json:"huella_id"`
	Estado    string `json:"estado"`
}
