package cmd

import (
    "fmt"
    "os"
    "strconv"
    "time"

    "dattool/internal/calendar"
    "github.com/spf13/cobra"
)

var jcalCmd = &cobra.Command{
    Use:   "jcal [year] [month]",
    Short: "Show Jalali calendar for a given month",
    Long:  "Displays a Jalali month calendar (similar to jcal). If no arguments are passed, current month is used.",

    Run: func(cmd *cobra.Command, args []string) {

        var jy, jm int

        // If user provided year/month
        if len(args) == 2 {
            y, err1 := strconv.Atoi(args[0])
            m, err2 := strconv.Atoi(args[1])
            if err1 != nil || err2 != nil {
                fmt.Println("Invalid year or month")
                os.Exit(1)
            }
            jy, jm = y, m
        } else {
            // Use current Gregorian date → convert to Jalali
            now := time.Now()
            jy, jm, _ = calendar.GregorianToJalali(now.Year(), int(now.Month()), now.Day())
        }
        now := time.Now()
        cjy, cjm, cjd := calendar.GregorianToJalali(now.Year(), int(now.Month()), now.Day())

        highlight := 0
        if jy == cjy && jm == cjm {
            highlight = cjd
        }

        // Print calendar
        fmt.Print(calendar.JalaliMonthCalendar(jy, jm, highlight))
    },
}

func init() {
    rootCmd.AddCommand(jcalCmd)
}
