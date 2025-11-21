package cmd

import (
    "fmt"
    "time"

    "dattool/internal/calendar"
    "github.com/spf13/cobra"
)

var onlyJalali bool
var onlyGregorian bool

var dateCmd = &cobra.Command{
    Use:   "date",
    Short: "Print current date",
    Run: func(cmd *cobra.Command, args []string) {
        now := time.Now()
        y, m, d := now.Date()
        jy, jm, jd := calendar.GregorianToJalali(y, int(m), d)

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
    rootCmd.AddCommand(dateCmd)
}
