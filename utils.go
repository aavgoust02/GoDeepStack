package GoDeepStack

import "log"

func check(err error) {
	if err != nil {
		log.Panic(err)
	}
}
