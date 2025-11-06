package cmd

import (
    "fmt"
    "time"

    "github.com/spf13/cobra"
    "github.com/<username>/datool/internal/calendar"
)

var dateCmd = &cobra.Command{
    Use:   "date",
    Short: "show jalali date",
    Long:  "Yhis command displays todays jalali date or ...",
    Args:  cobra.MaximumNArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        if len(args) == 0 {
            now := time.Now()
            y, m, d := now.Date()
            jy, jm, jd := calendar.GregorianToJalali(y, int(m), d)
            fmt.Printf("%04d-%02d-%02d\n", jy, jm, jd)
            return
        }
    },
}

func init() {
    rootCmd.AddCommand(dateCmd)
}
