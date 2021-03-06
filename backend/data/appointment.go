package data

import (
	"time"
)

type Appointment struct {
	Id        int       `json:"id"`
	Patient   int       `json:"patient"`
	Provider  int       `json:"provider"`
	Location  string    `json:"location"`
	Day       int       `json:"day"`
	StartTime time.Time `json:"start_time"`
	Duration  int       `json:"duration"`
	Cancelled bool      `json:"cancelled"`
	Completed bool      `json:"completed"`
}

func Appointments() (appointments []Appointment, err error) {
	rows, err := Db.Query("select * from appointment")

	if err != nil {
		return
	}

	var a Appointment

	for rows.Next() {
		rows.Scan(&a.Id, &a.Patient, &a.Provider, &a.Location, &a.Day, &a.StartTime, &a.Duration, &a.Cancelled, &a.Completed)

		if err != nil {
			return
		}
		appointments = append(appointments, a)
	}
	return
}

func (a *Appointment) Create() (err error) {
	statement := "insert into appointment(patient_id, provider_id, location, appt_day, start_time, duration) values($1, $2, $3, $4, $5)"
	stmt, err := Db.Prepare(statement)

	if err != nil {
		return
	}

	defer stmt.Close()

	_, err = stmt.Exec(a.Patient, a.Provider, a.Location, a.Day, a.StartTime, a.Duration)

	return
}

func AppointmentById(id int) (a Appointment, err error) {
	statement := "select id, patient_id, provider_id, appt_day, start_time, duration, cancelled, completed from appointment where id = $1"

	err = Db.QueryRow(statement, id).Scan(&a.Id, &a.Patient, &a.Provider, &a.Location, &a.Day, &a.StartTime, &a.Duration, &a.Cancelled, &a.Completed)

	return

}

func (a *Appointment) Update() (err error) {
	statement := "update appointment set patient_id=$2, location=$3, appt_day=$4, start_time=$5, duration=$6  where id = $1"
	stmt, err := Db.Prepare(statement)

	if err != nil {
		return
	}

	defer stmt.Close()

	_, err = stmt.Exec(a.Id, a.Patient, a.Location, a.Day, a.StartTime, a.Duration)

	return
}

func (a *Appointment) Cancel() (err error) {
	_, err = Db.Exec("update appointment set cancelled = true where id = $1", a.Id)
	return
}

func (a *Appointment) Complete() (err error) {
	_, err = Db.Exec("update appointent set completed = true where id = $1", a.Id)
	return
}
