package transaction

import (
	"sort"
	"time"
)

const (
	Second = iota
	Minute
	Hour
	Day
	Week
	Month
	Year
)

var truncs = [7]time.Duration{time.Second, time.Minute, time.Hour}

type Transaction struct {
	Value     int
	Timestamp time.Time
}

func roundPeriod(period time.Time, interval int) time.Time {
	switch interval {
	case Day:
		return period.Truncate(time.Hour * 24)
	case Month:
		startOfMonth := time.Date(period.Year(), period.Month(), 1, 0, 0, 0, 0, period.Location())
		return startOfMonth
	case Year:
		startOfYear := time.Date(period.Year(), 1, 1, 0, 0, 0, 0, period.Location())
		return startOfYear
	default:
		return period.Truncate(truncs[interval])
	}
}

func Format(transactions []*Transaction, interval int) []*Transaction {
	// In the example the values are sorted by timestamp
	// but we can't be sure
	if len(transactions) == 0 {
		return nil
	}
	sort.SliceStable(transactions, func(i int, j int) bool {
		return transactions[i].Timestamp.Unix() > transactions[j].Timestamp.Unix()
	})
	res := make([]*Transaction, 0, len(transactions))
	res = append(res, &Transaction{
		transactions[0].Value,
		roundPeriod(transactions[0].Timestamp, interval),
	})
	for _, trans := range transactions {
		roundedTimestamp := roundPeriod(trans.Timestamp, interval)
		if roundedTimestamp == res[len(res)-1].Timestamp {
			continue
		}
		res = append(res, &Transaction{
			trans.Value,
			roundedTimestamp,
		})
	}
	return res
}
