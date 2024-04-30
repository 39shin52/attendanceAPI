package entity

import "time"

type User struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Affiliation string `json:"affiliation"`
}

type Attendance struct {
	ID   string    `json:"id"`
	Name string    `json:"name"`
	Time time.Time `json:"time"`
}

func ConvertUser(id string, name string, affiliation string) *User {
	return &User{
		ID:          id,
		Name:        name,
		Affiliation: affiliation,
	}
}

func ConvertAttendance(id string, name string, time time.Time) *Attendance {
	return &Attendance{
		ID:   id,
		Name: name,
		Time: time,
	}
}
