package main

import (
	"fmt"
	"time"
)

type Bike struct {
	UserId  int
	BikeId  int
	Address string
	Time    time.Time
}

func (b *Bike) String() string {
	return fmt.Sprintf("BikeId: %d.UserId: %d.Address: %s.Time: %s\n", b.BikeId, b.UserId, b.Address, b.Time)
}
func (b *Bike) ShowAvailable() string {
	return fmt.Sprintf("BikeId: %d.Address: %s\n", b.BikeId, b.Address)
}
