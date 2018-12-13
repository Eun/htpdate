package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/leekchan/timeutil"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	serversArg       = kingpin.Flag("server", "server to use for synchronization, multiple values possible").Short('s').Default("pool.ntp.org", "ntp.org").URLList()
	debugFlag        = kingpin.Flag("debug", "turn debug on").Short('d').Default("false").Bool()
	commandToSetDate = kingpin.Arg("command", `command to set the date. You can use following directives to format the date:
    %a    Weekday as locale’s abbreviated name.                             (Sun, Mon, ..., Sat)
    %A    Weekday as locale’s full name.                                    (Sunday, Monday, ..., Saturday)
    %w    Weekday as a decimal number, where 0 is Sunday and 6 is Saturday  (0, 1, ..., 6)
    %d    Day of the month as a zero-padded decimal number.                 (01, 02, ..., 31)
    %b    Month as locale’s abbreviated name.                               (Jan, Feb, ..., Dec)
    %B    Month as locale’s full name.                                      (January, February, ..., December)
    %m    Month as a zero-padded decimal number.                            (01, 02, ..., 12)
    %y    Year without century as a zero-padded decimal number.             (00, 01, ..., 99)
    %Y    Year with century as a decimal number.                            (1970, 1988, 2001, 2013)
    %H    Hour (24-hour clock) as a zero-padded decimal number.             (00, 01, ..., 23)
    %I    Hour (12-hour clock) as a zero-padded decimal number.             (01, 02, ..., 12)
    %p    Meridian indicator.                                               (AM, PM)
    %M    Minute as a zero-padded decimal number.                           (00, 01, ..., 59)
    %S    Second as a zero-padded decimal number.                           (00, 01, ..., 59)
    %f    Microsecond as a decimal number, zero-padded on the left.         (000000, 000001, ..., 999999)
    %z    UTC offset in the form +HHMM or -HHMM                             (+0000)
    %Z    Time zone name                                                    (UTC)
    %j    Day of the year as a zero-padded decimal number                   (001, 002, ..., 366)
    %U    Week number of the year (Sunday as the first day of the week) as a zero padded decimal number. All days in a new year preceding the first Sunday are considered to be in week 0.
                                                                            (00, 01, ..., 53)
    %W    Week number of the year (Monday as the first day of the week) as a decimal number. All days in a new year preceding the first Monday are considered to be in week 0.
                                                                            (00, 01, ..., 53)
    %c    Date and time representation.                                     (Tue Aug 16 21:30:00 1988)
    %x    Date representation.                                              (08/16/88)
    %X    Time representation.                                              (21:30:00)
    %%    A literal '%' character.                                          (%)

    Example:
        htpdate --server pool.ntp.org --server ntp.org -- date --set="%a, %d %b %Y %H:%M:%S %Z"
    
    If not specified htpdate tries to set the date by itself
`).Strings()
)

func main() {
	kingpin.Version("1.0")

	kingpin.Parse()

	logDebug := func(s string, v ...interface{}) {
		log.Printf(s, v...)
	}

	if !*debugFlag {
		logDebug = func(s string, v ...interface{}) {}
	}

	if len(*serversArg) <= 0 {
		fmt.Fprintf(os.Stderr, "No servers to synchronization\n")
		os.Exit(1)
		return
	}

	var dateToSet *time.Time

	for _, s := range *serversArg {
		if s.Scheme == "" {
			s.Scheme = "http"
		}
		logDebug("Querying `%s'", s.String())
		req, err := http.NewRequest(http.MethodHead, s.String(), nil)
		if err != nil {
			logDebug("Unable to create request for `%s': %v", s.String(), err)
			continue
		}
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			logDebug("Unable to do request to `%s': %v", s.String(), err)
			continue
		}
		date := res.Header.Get("Date")
		if date == "" {
			logDebug("`%s' has no date", s.String())
			continue
		}
		t, err := time.Parse(time.RFC1123, date)
		if err != nil {
			logDebug("`%s' has invalid date (`%s') was not RFC1123", s.String(), date)
			continue
		}
		logDebug("`%s' reports `%s'", s.String(), t.String())
		dateToSet = &t
		break
	}
	if dateToSet == nil {
		logDebug("No time to set")
		fmt.Fprintf(os.Stderr, "No server responded a valid time\n")
		os.Exit(1)
		return
	}
	logDebug("Setting time to %s", (*dateToSet).String())

	if len(*commandToSetDate) > 0 {
		for i := range *commandToSetDate {
			(*commandToSetDate)[i] = timeutil.Strftime(dateToSet, (*commandToSetDate)[i])
		}
		logDebug("Running %v %d", *commandToSetDate, len(*commandToSetDate))
		var args []string
		if len(*commandToSetDate) > 1 {
			args = (*commandToSetDate)[1:]
		}
		cmd := exec.Command((*commandToSetDate)[0], args...)
		cmd.Stdin = os.Stdin
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		if err := cmd.Run(); err != nil {
			fmt.Fprintf(os.Stderr, "Error on running `%s': %v", strings.Join(*commandToSetDate, " "), err)
			os.Exit(1)
			return
		}
		os.Exit(1)
		return
	}

	if err := setTime(*dateToSet); err != nil {
		panic(err)
	}
}
