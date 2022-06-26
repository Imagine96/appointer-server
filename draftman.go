package main

import (
	"fmt"
)

type Agenda map[string]Day

type Draftman struct {
	Id          string         `json:"_id" bson:"_id"`
	ContactInfo ContactInfo    `json:"contactInfo" bson:"contactInfo"`
	Agenda      Agenda         `json:"agenda,omitempty" bson:"agenda,omitempty"`
	History     map[string]Day `json:"history,omitempty" bson:"history,omitempty"`
}

func (d Draftman) getAgenda() Agenda {
	return d.Agenda
}

func (d Draftman) getContactInfo() ContactInfo {
	return d.ContactInfo
}

func (d *Draftman) updateContactInfo(newInfo ContactInfo) {
	d.ContactInfo = newInfo
}

func (d *Draftman) addScheduleReqToDate(date string, s Schedule) error {
	if err := d.Agenda.checkDate(date); err != nil {
		return err
	}

	if v, exist := d.Agenda[date]; exist {
		v.addScheduleReq(&s)
		d.Agenda[date] = v
	} else {
		d.Agenda[date] = Day{date, map[string]Schedule{}, map[string]Schedule{s.Id: s}, false, 6}
	}
	return nil
}

func (d *Draftman) addConfirmedScheduleToDate(date string, s Schedule) error {
	if err := d.Agenda.checkDate(date); err != nil {
		return err
	}

	if v, exist := d.Agenda[date]; exist {
		v.confirmedSchedules[s.Id] = s
	} else {
		d.Agenda[date] = Day{date, map[string]Schedule{s.Id: s}, map[string]Schedule{}, false, 6}
	}
	return nil
}

func (d *Draftman) confirmScheduleReq(date string, id string) error {
	if err := d.Agenda.checkDate(date); err != nil {
		return err
	}

	if v, exist := d.Agenda[date]; exist {
		if _, exist := v.schedulesRequests[id]; exist {
			v.confirmScheduleReq(id)
			d.Agenda[date] = v
			return nil
		}
		return fmt.Errorf("could find not schedule %v on day %v", id, date)
	}
	return fmt.Errorf("could not find any schedule on day %v", date)
}

func (d *Draftman) rejectScheduleReq(date string, id string) error {
	if v, exist := d.Agenda[date]; exist {
		if err := v.rejectScheduleReq(id); err != nil {
			return err
		}
		return nil
	}
	return fmt.Errorf("could not find schedule %v on day %v", id, date)
}

func (d *Draftman) toggleFreeDay(date string) {
	if v, exist := d.Agenda[date]; exist {
		v.IsFreeDay = !v.IsFreeDay
		d.Agenda[date] = v
		return
	}
	d.Agenda[date] = Day{date, map[string]Schedule{}, map[string]Schedule{}, true, 0}
}

func (d *Draftman) onDayEnd(today string) {
	if v, exist := d.Agenda[today]; exist {
		d.History[today] = v
		delete(d.Agenda, today)
	}
}

func (d *Draftman) setScheduleOver(date string, id string) error {
	if v, exist := d.Agenda[date]; exist {
		if s, exist := v.confirmedSchedules[id]; exist {
			s.Done = true
			return nil
		}
		return fmt.Errorf("could not schedule %v on day %v", id, date)
	}
	return fmt.Errorf("could not find any schedule on day %v", date)
}

func (d *Draftman) changeScheduleDay(sDate string, sId string, targetDate string) error {

	if v, exist := d.Agenda[sDate]; exist {
		if _, exist := v.schedulesRequests[sId]; exist {
			s, src, err := removeScheduleFromMap(sId, v.schedulesRequests)
			if err != nil {
				return err
			}
			v.schedulesRequests = src
			d.addScheduleReqToDate(targetDate, *s)
			return nil
		} else if _, exist := v.confirmedSchedules[sId]; exist {
			s, src, err := removeScheduleFromMap(sId, v.confirmedSchedules)
			if err != nil {
				return err
			}
			v.confirmedSchedules = src
			d.addConfirmedScheduleToDate(targetDate, *s)
		}
		d.Agenda[sDate] = v
		return nil
	}
	return fmt.Errorf("could not find any schedule on day %v ", sDate)
}

func (a Agenda) checkDate(date string) error {
	if v, exist := a[date]; exist {
		if v.isDayFull() {
			return fmt.Errorf("day %v is full", date)
		}
	}
	return nil
}
