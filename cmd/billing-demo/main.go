package main

import (
	"bufio"
	crand "crypto/rand"
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/combtis/billing"
)

func main() {
	//t := rand.NewPCG(uint64(time.Now().UnixNano()), 1024)

	bs := make([]byte, 32)
	crand.Read(bs)
	chacha8 := rand.NewChaCha8([32]byte(bs))
	t := rand.New(chacha8)

	billing := billing.NewBilling()
	fmt.Printf("balance: %.2f RUB\n", float64(billing.GetBalance())/100)

	wd, _ := os.Getwd()
	f, err := os.OpenFile(wd+"/demo.jsondb", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0o600)
	if err != nil {
		os.Exit(42)
	}
	defer f.Close()
	f.Seek(0, 0)

	count := int64(0)
	total := int64(0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		if text == "" {
			continue
		}
		v := &billing.Transaction{}
		if err := json.Unmarshal([]byte(text), v); err != nil {
			continue
		}
		tt, _ := strconv.ParseInt(v.Amount, 10, 64)
		total += tt
		v.Amount = v.Amount[:len(v.Amount)-2] + "." + v.Amount[len(v.Amount)-2:]
		count++
		fmt.Printf("%#+v\n", v)
	}

	var tr *billing.Transaction
	if len(os.Args) == 2 && os.Args[1] == "add" {
		tr = &billing.Transaction{
			Status:    "OK",
			Amount:    fmt.Sprintf("%d", t.Int64N(2999999)+1000000),
			Currency:  "RUB",
			Timestamp: time.Now(),
		}
		gs, _ := json.Marshal(tr)
		f.Write([]byte("\n" + string(gs[:])))
		tr.Amount = tr.Amount[:len(tr.Amount)-2] + "." + tr.Amount[len(tr.Amount)-2:]
		count++
	}
	if tr != nil {
		fmt.Printf("%#+v\n", tr)
	}
	fmt.Printf("Lines: %d; Middle: %.2f\nTotal: %.2f\n", count, float64(total)/2/100/float64(count), float64(total)/100)
}
