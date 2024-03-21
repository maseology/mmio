package mmio

import "time"

func dateParse(s string) (time.Time, error) {
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		t, err = time.Parse("2006-01-02 15:04:05", s)
		if err != nil {
			t, err = time.Parse("2006-01-02T15:04:05", s)
			if err != nil {
				t, err = time.Parse("2006-01-02 15:04:05 +0000 UTC", s)
				if err != nil {
					t, err = time.Parse("2006-01-02 15:04:05+00:00", s)
					if err != nil {
						t, err = time.Parse("2006-01-02 15:04", s)
						if err != nil {
							return time.Time{}, err
						}
					}
				}
			}
		}
	}
	return t, nil
}

// // DayDate returns the input time as a date
// func DayDate(t time.Time) time.Time {
// 	year, month, day := t.Date()
// 	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
// }
