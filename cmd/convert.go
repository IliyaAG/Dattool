package cmd

import (
    "errors"
    "fmt"
    "strings"

    "dattool/internal/calendar"
    "github.com/spf13/cobra"
)

var convertTo string

var convertCmd = &cobra.Command{
    Use:   "convert <date>",
    Short: "Convert Gregorian <-> Jalali date",
    Args:  cobra.ExactArgs(1),
    RunE: func(cmd *cobra.Command, args []string) error {

        input := args[0]

        year, month, day, err := parseDate(input)
        if err != nil {
            return err
        }

        // ---- Forced conversion ----
        if convertTo == "jalali" || convertTo == "j" {
            jy, jm, jd := calendar.GregorianToJalali(year, month, day)
            fmt.Printf("%04d-%02d-%02d\n", jy, jm, jd)
            return nil
        }

        if convertTo == "gregorian" || convertTo == "g" {
            gy, gm, gd := calendar.JalaliToGregorian(year, month, day)
            fmt.Printf("%04d-%02d-%02d\n", gy, gm, gd)
            return nil
        }

        // ---- Auto detect mode ----
        if year > 1700 {
            // Gregorian → Jalali
            jy, jm, jd := calendar.GregorianToJalali(year, month, day)
            fmt.Printf("%04d-%02d-%02d\n", jy, jm, jd)
            return nil
        }

        // Jalali → Gregorian
        gy, gm, gd := calendar.JalaliToGregorian(year, month, day)
        fmt.Printf("%04d-%02d-%02d\n", gy, gm, gd)
        return nil
    },
}

func init() {
    convertCmd.Flags().StringVarP(&convertTo, "to", "t", "", "Force conversion (jalali|gregorian)")
    rootCmd.AddCommand(convertCmd)
}

func parseDate(s string) (int, int, int, error) {

    // Allow both 2025-11-21 and 2025/11/21
    separator := "-"
    if strings.Contains(s, "/") {
        separator = "/"
    }

    parts := strings.Split(s, separator)
    if len(parts) != 3 {
        return 0, 0, 0, errors.New("invalid date format, expected YYYY-MM-DD or YYYY/MM/DD")
    }

    var y, m, d int
    _, err := fmt.Sscanf(s, fmt.Sprintf("%%d%s%%d%s%%d", separator, separator), &y, &m, &d)
    if err != nil {
        return 0, 0, 0, errors.New("cannot parse date numbers")
    }

    return y, m, d, nil
}
