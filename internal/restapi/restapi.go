package restapi

import (
	"database/sql"
)

// https://stackoverflow.com/questions/35038864/how-to-access-global-variables
var Db *sql.DB

type customer struct {
	VasarloID    int64           `json:"VasarloID"` // Not needed, because of auto increment
	Nev          string          `json:"Nev" binding:"required"`
	Email        string          `json:"Email" binding:"required"`
	Telefonszam  string          `json:"Telefonszam" binding:"required"`
	LeadasiIdo   string          `json:"LeadasiIdo"`
	ElkeszultIdo sql.NullString  `json:"ElkeszultIdo"`
	Osszeg       sql.NullFloat64 `json:"Osszeg"`
}

type order struct {
	VasarloID int     `json:"VasarloID"`
	EtelID    int     `json:"EtelID"`
	Ar        float64 `json:"Ar"`
	Darab     int     `json:"Darab"`
}

type fullOrder struct {
	Customer customer    `json:"Customer" binding:"required"`
	Foods    []foodOrder `json:"Foods" binding:"required"`
}

type foodOrder struct {
	EtelID int `json:"EtelID" binding:"required"`
	Volume int `json:"Volume" binding:"required"`
}
type foodWithVolume struct {
	Food   food `json:"Food" binding:"required"`
	Volume int  `json:"Volume" binding:"required"`
}

type food struct {
	EtelID  int     `json:"EtelID"`
	Nev     string  `json:"Nev" binding:"required"`
	Leiras  string  `json:"Leiras" `
	KepPath string  `json:"KepPath" `
	Ar      float64 `json:"Ar" binding:"required"`
}
