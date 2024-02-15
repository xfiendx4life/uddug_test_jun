package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"path"
	"time"

	"github.com/xfiendx4life/uddug_test_jun/internal/pkg/transaction"
)

func fromMapToTransaction(m []map[string]int) []*transaction.Transaction {
	inputData := make([]*transaction.Transaction, 0, len(m))
	for _, item := range m {
		inputData = append(inputData, &transaction.Transaction{
			Value:     item["value"],
			Timestamp: time.Unix(int64(item["timestamp"]), 0),
		})
	}
	return inputData
}

func fromTransactionToMap(m []*transaction.Transaction) []map[string]int {
	inputData := make([]map[string]int, 0)
	for _, item := range m {
		inputData = append(inputData, map[string]int{
			"value":     item.Value,
			"timestamp": int(item.Timestamp.Unix()),
		})
	}
	return inputData
}

func main() {
	var filename, interval string

	flag.StringVar(&filename, "filename", "input.json", "file containing data to test function")
	flag.StringVar(&interval, "interval", "day", "desireable interval to filter data [second, minute, hour, day, month, year]")
	flag.Parse()
	p, err := os.Getwd()
	if err != nil {
		log.Fatalf("can't get filepath %s", err)
	}
	log.Println(path.Join(p, filename))
	file, err := os.ReadFile(path.Join(p, filename))
	if err != nil {
		log.Fatalf("can't read file %s", err)
	}
	test := []map[string]int{}

	if err := json.Unmarshal(file, &test); err != nil {
		log.Fatalf("can't unmarshal file %s", err)
	}
	res := transaction.Format(fromMapToTransaction(test), interval)

	toWrite, err := json.Marshal(fromTransactionToMap(res))
	if err != nil {
		log.Fatalf("can't marshal result %s", err)
	}
	if err = os.WriteFile(path.Join(p, "output.json"), toWrite, os.ModeAppend); err != nil {
		log.Fatalf("can't write to file result %s", err)
	}

}
