package assert

import "log"

func Assert(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func AssertMsg(err error, msg string) {
	if err != nil {
		log.Fatalf("%v: %v\n", msg, err)
	}
}
