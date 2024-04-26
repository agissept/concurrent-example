package process

import (
	"log"
	"time"
)

func WriteName(registrant int, table int) {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	time.Sleep(time.Millisecond * 1000)
	log.Println("Write name by registrant: ", registrant, "on table: ", table)
}

func GetNameTag(registrant int, table int) {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	time.Sleep(time.Millisecond * 1000)
	log.Println("Get name tag by registrant: ", registrant, "on table: ", table)
}

func GetSnack(registrant int, table int) {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	time.Sleep(time.Millisecond * 1000)
	log.Println("Get snack by registrant: ", registrant, "on table: ", table)
}
