package cmd

import (
    "fmt"
    "time"

    "dattool/internal/calendar"
    "github.com/spf13/cobra"
)

var dateCmd = &cobra.Command{
    Use:   "date",
    Short: "Print today's Jalali date",
    Run: func(cmd *cobra.Command, args []string) {
        now := time.Now()
        y, m, d := now.Date()
        jy, jm, jd := calendar.GregorianToJalali(y, int(m), d)
        fmt.Printf("%04d-%02d-%02d\n", jy, jm, jd)
    },
}

func init() {
    rootCmd.AddCommand(dateCmd)
}
