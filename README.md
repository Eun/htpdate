# htpdate

A simple htpdate implementation in golang

```
usage: htpdate [<flags>] [<command>...]

Flags:
      --help     Show context-sensitive help (also try --help-long and --help-man).
  -s, --server=pool.ntp.org... ...
                 server to use for synchronization, multiple values possible
  -d, --debug    turn debug on
      --version  Show application version.

Args:
  [<command>]  command to set the date. You can use following directives to format the date:

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
                     htpdate --server pool.ntp.org --server ntp.org date --set "%a, %d %b %Y %H:%M:%S %Z"

                 If not specified htpdate tries to set the date by itself
```
