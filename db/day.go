package db

import (
	"errors"
)

type ScheduleDetails []string

type ContactInfo struct {
	FirstName   string `json:"firstName" bson:"firstName"`
	LastName    string `json:"lastName" bson:"lastName"`
	Email       string `json:"email" bson:"email"`
	PhoneNumber string `json:"phoneNumber" bson:"phoneNumber"`
}

type Schedule struct {
	Id                string `json:"_id" bson:"_id"`
	Direction         string `json:"direction" bson:"direction"`
	Hour              string `json:"hour" bson:"hour"`
	importantDetails  ScheduleDetails
	clientContactInfo ContactInfo
	Done              bool `json:"done" bson:"done"`
}

type Day struct {
	Date               string `json:"date" bson:"date"`
	confirmedSchedules map[string]Schedule
	schedulesRequests  map[string]Schedule
	IsFreeDay          bool `json:"isFreeDay" bson:"isFreeDay"`
	MaxSchedules       int  `json:"maxSchedules" bson:"maxSchedules"`
}

func (d *Day) addScheduleReq(s *Schedule) error {
	if d.IsFreeDay {
		return errors.New("day " + d.Date + "is not available for schedule")
	}
	if d.isDayFull() {
		return errors.New("day " + d.Date + " is full")
	}
	d.schedulesRequests[s.Id] = *s
	return nil
}

func (d *Day) rejectScheduleReq(id string) error {
	_, _, err := removeScheduleFromMap(id, d.schedulesRequests)
	if err != nil {
		return err
	}
	return nil
}

func (d *Day) confirmScheduleReq(id string) error {
	if d.isDayFull() {
		return errors.New("day " + d.Date + " is full")
	}
	targetSchedule, remain, err := removeScheduleFromMap(id, d.schedulesRequests)
	d.schedulesRequests = remain
	if err != nil {
		return err
	}
	d.confirmedSchedules[targetSchedule.Id] = *targetSchedule
	if len(d.confirmedSchedules) == d.MaxSchedules {
		//redirect to select other day
		d.schedulesRequests = map[string]Schedule{}
	}
	return nil
}

func (d *Day) onScheduleDone(id string) error {

	if _, exist := d.confirmedSchedules[id]; exist {
		if d.confirmedSchedules[id].Done {
			return errors.New("The schedule is over")
		}
		entry := d.confirmedSchedules[id]
		entry.Done = true
		d.confirmedSchedules[id] = entry
	}
	return nil
}

func (d Day) isDayFull() bool {
	return d.MaxSchedules == len(d.confirmedSchedules)
}

func (s Schedule) getContactInfo() ContactInfo {
	return s.clientContactInfo
}

func (s *Schedule) updateContactInfo(newInfo ContactInfo) {
	s.clientContactInfo = newInfo
}

func removeScheduleFromMap(targetId string, src map[string]Schedule) (*Schedule, map[string]Schedule, error) {
	if targetSchedule, exist := src[targetId]; exist {
		delete(src, targetId)
		return &targetSchedule, src, nil
	} else {
		return nil, nil, errors.New("could not find schedule " + targetId)
	}
}
