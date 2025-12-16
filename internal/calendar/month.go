package calendar

import (
    "fmt"
    "strings"
    "time"
)

const (
    ansiReset = "\033[0m"
    ansiGreenBG = "\033[42;30m"
    ansiRed = "\033[31m"
)
// returns number of days in a Jalali month
func JalaliMonthDays(year, month int) int {
    if month <= 6 {
        return 31
    } else if month <= 11 {
        return 30
    }
    // Esfand
    if ((year+11)%33)%4 == 0 {
        return 30
    }
    return 29
}

// Build a printable Jalali calendar
func JalaliMonthCalendar(jy, jm, highlightDay int) string {
    var builder strings.Builder

    // Header
    monthNames := []string{
        "Farvardin", "Ordibehesht", "Khordad",
        "Tir", "Mordad", "Shahrivar",
        "Mehr", "Aban", "Azar",
        "Dey", "Bahman", "Esfand",
    }

    fmt.Fprintf(&builder, "  %s %d\n", monthNames[jm-1], jy)
    builder.WriteString("Sh Ye Do Se Ch Pa Jo\n")

    // Find first weekday
    gy, gm, gd := JalaliToGregorian(jy, jm, 1)
    first := time.Date(gy, time.Month(gm), gd, 0, 0, 0, 0, time.UTC)
    weekday := (int(first.Weekday()) + 1)%7

    // Leading spaces
    for i := 0; i < weekday; i++ {
        builder.WriteString("   ")
    }

    // Days
    days := JalaliMonthDays(jy, jm)
    for d := 1; d <= days; d++ {
        isFriday := weekday%7 == 6
        if d == highlightDay {
            fmt.Fprintf(&builder, "%s%2d%s ", ansiGreenBG, d, ansiReset)
        } else if isFriday {
            fmt.Fprintf(&builder, "%s%2d%s ", ansiRed, d, ansiReset)
        } else {
            fmt.Fprintf(&builder, "%2d ", d)
        }
        weekday++
        if weekday%7 == 0 {
            builder.WriteString("\n")
        }
    }

    builder.WriteString("\n")
    return builder.String()
}
