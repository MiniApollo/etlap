// RESTful API-hoz szükséges funkciók megvalósítása.
//
// A csomag az alábbi komponensekből épül fel:
//   - auth.go: Közvetítő szoftver(middleware) az adminisztrációs felületnek.
//   - customer:.go: Vásárlói adatokhoz kapcsolódó API végpontok.
//   - food.go: Ételadatokhoz kapcsolódó API végpontok.
//   - orders.go: Rendelések kezelésére szolgáló API végpontok.
//   - restapi.go: Deklarálja az adatszerkezeteket a csomag számára.
//
// Használathoz változók beállítása szükséges.
package restapi

import (
	"database/sql"
)

// Globális változó az adatbázis kapcsolat kezelésére.
//
// Nagy kezdőbetűk jelentése hogy nyilvános, globális.
//
// https://stackoverflow.com/questions/35038864/how-to-access-global-variables
var Db *sql.DB

// Vásárlók SQL tábla adatszerkezete.
type customer struct {
	VasarloID    int64           `json:"VasarloID"` // Not needed, because of auto increment
	Nev          string          `json:"Nev" binding:"required"`
	Email        string          `json:"Email" binding:"required"`
	Telefonszam  string          `json:"Telefonszam" binding:"required"`
	LeadasiIdo   string          `json:"LeadasiIdo"`
	ElkeszultIdo sql.NullString  `json:"ElkeszultIdo"`
	Osszeg       sql.NullFloat64 `json:"Osszeg"`
}

// Rendelések SQL tábla adatszerkezete.
type order struct {
	VasarloID int     `json:"VasarloID"`
	EtelID    int     `json:"EtelID"`
	Ar        float64 `json:"Ar"`
	Darab     int     `json:"Darab"`
}

// Összes rendelés Vásárlókra bontva.
type fullOrder struct {
	Customer customer    `json:"Customer" binding:"required"`
	Foods    []foodOrder `json:"Foods" binding:"required"`
}

// Adatszerkezet egy étel rendelésnek a tárolására.
type foodOrder struct {
	EtelID int `json:"EtelID" binding:"required"`
	Volume int `json:"Volume" binding:"required"`
}

// Adatszerkezet az étel minden adatának tárolására és rendelt mennyiségével.
type foodWithVolume struct {
	Food   food `json:"Food" binding:"required"`
	Volume int  `json:"Volume" binding:"required"`
}

// Ételek SQL tábla adatszerkezete.
type food struct {
	EtelID  int     `json:"EtelID"`
	Nev     string  `json:"Nev" binding:"required"`
	Leiras  string  `json:"Leiras" `
	KepPath string  `json:"KepPath" `
	Ar      float64 `json:"Ar" binding:"required"`
}
