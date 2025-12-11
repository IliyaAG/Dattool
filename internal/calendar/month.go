package calendar

import (
    "fmt"
    "strings"
    "time"
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
func JalaliMonthCalendar(jy, jm int) string {
    var builder strings.Builder

    // Header
    monthNames := []string{
        "Farvardin", "Ordibehesht", "Khordad",
        "Tir", "Mordad", "Shahrivar",
        "Mehr", "Aban", "Azar",
        "Dey", "Bahman", "Esfand",
    }

    fmt.Fprintf(&builder, "  %s %d\n", monthNames[jm-1], jy)
    builder.WriteString("Su Mo Tu We Th Fr Sa\n")

    // Find first weekday
    gy, gm, gd := JalaliToGregorian(jy, jm, 1)
    first := time.Date(gy, time.Month(gm), gd, 0, 0, 0, 0, time.UTC)
    weekday := int(first.Weekday()) // Sunday = 0

    // Leading spaces
    for i := 0; i < weekday; i++ {
        builder.WriteString("   ")
    }

    // Days
    days := JalaliMonthDays(jy, jm)
    for d := 1; d <= days; d++ {
        fmt.Fprintf(&builder, "%2d ", d)
        weekday++
        if weekday%7 == 0 {
            builder.WriteString("\n")
        }
    }

    builder.WriteString("\n")
    return builder.String()
}
