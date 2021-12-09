package main

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	time2 "time"
)

const layout = "Jan 2, 2006 at 3:04pm (MST)"

var config = mysql.Config{
	User:      "nastya",
	Passwd:    "1234",
	Net:       "tcp",
	Addr:      "localhost:3306",
	DBName:    "BikeService",
	Collation: "",
}

type repository interface {
	CreateBike(address string) (string, error)
	RentBike(userId, bikeId int) (string, error)
	ReturnBike(userId, bikeId int, newAddress string) (int, error)
	GetAll() ([]*Bike, error)
}
type Repository struct {
	db *sql.DB
}

func NewRepository() repository {
	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		return nil
	}
	return &Repository{db: db}
}
func (r *Repository) CreateBike(address string) (string, error) {
	_, err := r.db.Query("insert into bikes (userId,address,rentStartTime) values (?,?,?)", 0,
		address, time2.Now().Format(layout))
	if err != nil {
		return "", err
	}
	return "Bike was created", nil
}
func (r *Repository) GetAll() ([]*Bike, error) {
	rows, err := r.db.Query("select * from bikes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	bikes := make([]*Bike, 0)
	for rows.Next() {
		b := Bike{}
		var str string
		err := rows.Scan(&b.BikeId, &b.UserId, &b.Address, &str)
		if err != nil {
			return nil, err
		}
		if str != "" {
			time, err := time2.Parse(layout, str)
			if err != nil {
				return nil, err
			}
			b.Time = time
		} else {
			b.Time = time2.Now()
		}

		bikes = append(bikes, &b)
	}
	return bikes, nil
}
func (r *Repository) RentBike(userId, bikeId int) (string, error) {
	time := time2.Now().Format(layout)
	rAff, err := r.db.Exec("update bikes set userId=?,rentStartTime=? where bikeId=? && userId=?", userId, time, bikeId, 0)
	if err != nil {
		return "DB error", err
	}
	rows, err := rAff.RowsAffected()
	if err != nil {
		return "rows affected error", err
	}
	if rows == 0 {
		return "no such bike available", nil
	}
	return "bike rented successfully", nil
}
func (r *Repository) ReturnBike(userId, bikeId int, newAddress string) (int, error) {
	row := r.db.QueryRow("select * from bikes where userId=? && bikeId=?", userId, bikeId)
	var empty [3]string
	var t string
	err := row.Scan(&empty[0], &empty[1], &empty[2], &t)
	if err != nil {
		return 0, err
	}
	time1, err := time2.Parse(layout, t)
	if err != nil {
		return 0, err
	}
	time2 := time2.Now()
	res, err := r.db.Exec("update bikes set userId=?,address=? where bikeId=? && userId=?", 0, newAddress, bikeId, userId)
	if err != nil {
		return 0, err
	}
	rAff, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	if rAff == 0 {
		return 0, nil
	}
	return int(time2.Unix()-time1.Unix()) / 3600, nil
}
