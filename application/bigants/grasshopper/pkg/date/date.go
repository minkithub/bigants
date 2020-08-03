package date

import (
	"encoding/json"
	"errors"
	"time"
)

func Parse(ds string) (d Date, err error) {
	t, err := time.Parse("2006-01-02", ds)
	if err != nil {
		return d, err
	}

	return Date(t), nil
}

type Date time.Time

// Year returns the year in which d occurs.
func (d Date) Year() int { return time.Time(d).Year() }

// Month returns the month of the year specified by d.
func (d Date) Month() time.Month { return time.Time(d).Month() }

// Day returns the day of the month specified by d.
func (d Date) Day() int { return time.Time(d).Day() }

func (d Date) String() string { return time.Time(d).Format("2006-01-02") }

func (d Date) AddDays(days int) Date {
	return Date(time.Time(d).Add(time.Hour * 24 * time.Duration(days)))
}

func (d *Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

func (d *Date) UnmarshalJSON(b []byte) (err error) {
	if string(b) == "null" {
		return nil
	}
	var src string
	if err = json.Unmarshal(b, &src); err != nil {
		return err
	}
	*d, err = Parse(src)
	return err
}

func (d *Date) Scan(src interface{}) (err error) {
	switch src.(type) {
	case string:
		*d, err = Parse(src.(string))
	case []byte:
		*d, err = Parse(string(src.([]byte)))
	case time.Time:
		*d = Date(src.(time.Time))
	default:
		err = errors.New("Incompatible type for Date")
	}
	return err
}
