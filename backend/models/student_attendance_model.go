package models

import "time"

type AttendanceMark int

const (
	Absent AttendanceMark = iota
	Present
	Leave
)

func (a AttendanceMark) String() string {
	return [...]string{"Absent", "Present", "Leave"}[a]
}

type StudentAttendance struct {
	StudentID string                          `json:"_id" bson:"_id"`
	SubjectID string                          `json:"subject_id" bson:"subject_id"`
	Attendance map[time.Time]AttendanceMark   `json:"attendance" bson:"attendance"`
}