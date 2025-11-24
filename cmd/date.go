package cmd

import (
    "fmt"
    "time"
    "strings"

    "dattool/internal/calendar"
    "github.com/spf13/cobra"
)

var onlyJalali bool
var onlyGregorian bool
var format string

var weekdaysFA = []string{
    "yekshanbe",
    "doshanbe",
    "seshanbe",
    "chaharshanbe",
    "panjshanbe",
    "jome",
    "shanbe",
}

var gregorianMonthsEN = []string{
    "January",
    "February",
    "March",
    "April",
    "May",
    "June",
    "July",
    "August",
    "September",
    "October",
    "November",
    "December",
}

var jalaliMonthsFA = []string{
    "Farvardin",
    "Ordibehesht",
    "Khordad",
    "Tir",
    "Mordad",
    "Shahrivar",
    "Mehr",
    "Aban",
    "Azar",
    "Dey",
    "Bahman",
    "Esfand",
}
var dateCmd = &cobra.Command{
    Use:   "date",
    Short: "Print current date",
    Run: func(cmd *cobra.Command, args []string) {
        now := time.Now()
        y, m, d := now.Date()
        h, min, sec := now.Clock()
        jy, jm, jd := calendar.GregorianToJalali(y, int(m), d)

        weekdayEN := now.Weekday().String()
        weekdayFA := weekdaysFA[int(now.Weekday())]
        if format != "" {
            out := format
            out = strings.ReplaceAll(out, "%Y", fmt.Sprintf("%04d",y))
            out = strings.ReplaceAll(out, "%m", fmt.Sprintf("%02d",m))
            out = strings.ReplaceAll(out, "%d", fmt.Sprintf("%02d",d))
            out = strings.ReplaceAll(out, "%gmonthname", gregorianMonthsEN[m-1])
            out = strings.ReplaceAll(out, "%JY", fmt.Sprintf("%04d",jy))
            out = strings.ReplaceAll(out, "%JM", fmt.Sprintf("%02d",jm))
            out = strings.ReplaceAll(out, "%JD", fmt.Sprintf("%02d",jd))
            out = strings.ReplaceAll(out, "%jmonthname", jalaliMonthsFA[jm-1])
            out = strings.ReplaceAll(out, "%W", weekdayEN)
            out = strings.ReplaceAll(out, "%JW", weekdayFA)
            out = strings.ReplaceAll(out, "%H", fmt.Sprintf("%02d", h))
            out = strings.ReplaceAll(out, "%M", fmt.Sprintf("%02d", min))
            out = strings.ReplaceAll(out, "%S", fmt.Sprintf("%02d", sec))

            fmt.Println(out)
            return
        }

        if onlyJalali && !onlyGregorian {
            fmt.Printf("%04d-%02d-%02d\n", jy, jm, jd)
            return
        }

        if onlyGregorian && !onlyJalali {
            fmt.Printf("%04d-%02d-%02d\n", y, m, d)
            return
        }
        fmt.Printf("Gregorian: %04d-%02d-%02d\n",y,m,d)
        fmt.Printf("Jalali: %04d-%02d-%02d\n",jy,jm,jd)
    },
}

func init() {
    dateCmd.Flags().BoolVarP(&onlyJalali, "jalali", "j", false, "Prinnt only jalali date")
    dateCmd.Flags().BoolVarP(&onlyGregorian, "gregorian", "g", false, "Prinnt only gregorian date")
    dateCmd.Flags().StringVarP(&format, "format", "f", "", "Custom date format (e.g. \"%Y%m%d\")")

    rootCmd.AddCommand(dateCmd)
}
