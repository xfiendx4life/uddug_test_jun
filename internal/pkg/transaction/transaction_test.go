package transaction_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/xfiendx4life/uddug_test_jun/internal/pkg/transaction"
)

var cases = map[string]map[string][]*transaction.Transaction{
	"default": {
		"input": {
			{
				Value:     4456,
				Timestamp: time.Unix(1616026248, 0),
			},
			{
				Value:     4231,
				Timestamp: time.Unix(1616022648, 0),
			},
			{
				Value:     5212,
				Timestamp: time.Unix(1616019048, 0),
			},
			{
				Value:     4321,
				Timestamp: time.Unix(1615889448, 0),
			},
			{
				Value:     4567,
				Timestamp: time.Unix(1615871448, 0),
			},
		},
		"output": {
			{
				Value:     4456,
				Timestamp: time.Unix(1616025600, 0),
			},
			{
				Value:     4231,
				Timestamp: time.Unix(1615939200, 0),
			},
			{
				Value:     4321,
				Timestamp: time.Unix(1615852800, 0),
			},
		},
	},
	"unsorted": {
		"input": {
			{
				Value:     4231,
				Timestamp: time.Unix(1616022648, 0),
			},
			{
				Value:     4321,
				Timestamp: time.Unix(1615889448, 0),
			},
			{
				Value:     5212,
				Timestamp: time.Unix(1616019048, 0),
			},
			{
				Value:     4567,
				Timestamp: time.Unix(1615871448, 0),
			},
			{
				Value:     4456,
				Timestamp: time.Unix(1616026248, 0),
			},
		},
		"output": {
			{
				Value:     4456,
				Timestamp: time.Unix(1616025600, 0),
			},
			{
				Value:     4231,
				Timestamp: time.Unix(1615939200, 0),
			},
			{
				Value:     4321,
				Timestamp: time.Unix(1615852800, 0),
			},
		},
	},
	"empty": {
		"input":  {},
		"output": nil,
	},
	"byHour": {
		"input": {
			{
				Value:     4456,
				Timestamp: time.Unix(1616026248, 0),
			},
			{
				Value:     4231,
				Timestamp: time.Unix(1616022648, 0),
			},
			{
				Value:     5212,
				Timestamp: time.Unix(1616019048, 0),
			},
			{
				Value:     4321,
				Timestamp: time.Unix(1615889448, 0),
			},
			{
				Value:     4567,
				Timestamp: time.Unix(1615871448, 0),
			},
		},
		"output": {
			{
				Value:     4456,
				Timestamp: time.Unix(1616025600, 0),
			},
			{
				Value:     4231,
				Timestamp: time.Unix(1616022000, 0),
			},
			{
				Value:     5212,
				Timestamp: time.Unix(1616018400, 0),
			},
			{
				Value:     4321,
				Timestamp: time.Unix(1615888800, 0),
			},
			{
				Value:     4567,
				Timestamp: time.Unix(1615870800, 0),
			},
		},
	},
	"byMonth": {
		"input": {
			{
				Value:     4456,
				Timestamp: time.Unix(1616026248, 0),
			},
			{
				Value:     4231,
				Timestamp: time.Unix(1616022648, 0),
			},
			{
				Value:     5212,
				Timestamp: time.Unix(1616019048, 0),
			},
			{
				Value:     4321,
				Timestamp: time.Unix(1615889448, 0),
			},
			{
				Value:     4567,
				Timestamp: time.Unix(1615871448, 0),
			},
		},
		"output": {
			{
				Value:     4456,
				Timestamp: time.Date(2021, 3, 1, 0, 0, 0, 0, time.Local),
			},
		},
	},
	"byYear": {
		"input": {
			{
				Value:     4456,
				Timestamp: time.Unix(1616026248, 0),
			},
			{
				Value:     4231,
				Timestamp: time.Unix(1616022648, 0),
			},
			{
				Value:     5212,
				Timestamp: time.Unix(1616019048, 0),
			},
			{
				Value:     4321,
				Timestamp: time.Unix(1615889448, 0),
			},
			{
				Value:     4567,
				Timestamp: time.Unix(1615871448, 0),
			},
		},
		"output": {
			{
				Value:     4456,
				Timestamp: time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local),
			},
		},
	},
}

func TestTransactionFormat(t *testing.T) {
	res := transaction.Format(cases["default"]["input"], transaction.Day)
	for i, trans := range res {
		assert.Equal(t, cases["default"]["output"][i], trans)
		// fmt.Println(trans.Timestamp)
	}
}
func TestTransactionUnsorted(t *testing.T) {
	res := transaction.Format(cases["unsorted"]["input"], transaction.Day)
	for i, trans := range res {
		assert.Equal(t, cases["default"]["output"][i], trans)
	}
}

func TestTransactionEmpty(t *testing.T) {
	res := transaction.Format(cases["empty"]["input"], transaction.Day)
	assert.Nil(t, res)
}

func TestTransactionByHour(t *testing.T) {
	res := transaction.Format(cases["byHour"]["input"], transaction.Hour)
	for i, trans := range res {
		assert.Equal(t, cases["byHour"]["output"][i], trans)
	}
}

func TestTransactionByMonth(t *testing.T) {
	res := transaction.Format(cases["byMonth"]["input"], transaction.Month)
	for i, trans := range res {
		assert.Equal(t, cases["byMonth"]["output"][i], trans)
	}
}

func TestTransactionByYear(t *testing.T) {
	res := transaction.Format(cases["byYear"]["input"], transaction.Year)
	for i, trans := range res {
		assert.Equal(t, cases["byYear"]["output"][i], trans)
	}
}
