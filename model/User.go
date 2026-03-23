package model

import "time"

type User struct {
	Id    uint
	Pulsa float64
	Paket []Paket
}

type Paket struct {
	jumlah float64
	durasi time.Time
}
