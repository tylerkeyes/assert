package assert

import "log"

func assert(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func assert_msg(err error, msg string) {
	if err != nil {
		log.Fatalf("%v: %v\n", msg, err)
	}
}
