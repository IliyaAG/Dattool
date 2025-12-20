package calendar

import (
    "fmt"
    "strings"
)

func PrintThreeMonths(jy, jm, cjy, cjm, cjd int) {

    py, pm := PrevMonth(jy, jm)
    ny, nm := NextMonth(jy, jm)

    months := []struct {
        y, m int
    }{
        {py, pm},
        {jy, jm},
        {ny, nm},
    }

    rendered := [][]string{}

    for _, mo := range months {

        highlight := 0
        if mo.y == cjy && mo.m == cjm {
            highlight = cjd
        }

        cal := JalaliMonthCalendar(mo.y, mo.m, highlight)
        lines := strings.Split(cal, "\n")
        rendered = append(rendered, lines)
    }

    maxLines := len(rendered[0])

    for i := 0; i < maxLines; i++ {
        for _, m := range rendered {
            if i < len(m) {
                fmt.Printf("%-22s", m[i])
            }
        }
        fmt.Println()
    }
}
