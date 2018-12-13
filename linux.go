// +build linux,cgo

package main

// #include <sys/time.h>
// static int setTime(long timedelta) {
//     struct timeval timeofday;
//     timeofday.tv_sec = timedelta;
//     return adjtime(&timeofday, NULL);
// }
import "C"

import (
	"fmt"
	"time"
)

func setTime(t time.Time) error {
	if err := C.setTime(C.long(t.Unix())); err != 0 {
		return fmt.Errorf("Error on setting time, error code: %d", err)
	}
	return nil
}
