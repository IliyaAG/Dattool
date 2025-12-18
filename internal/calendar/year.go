package calendar

import (
    "fmt"
    "strings"
)

func PrintJalaliYear(jy, cjy, cjm, cjd int) {

    fmt.Printf("          %d\n\n", jy)

    for row := 0; row < 4; row++ {

        months := [][]string{}

        for col := 0; col < 3; col++ {
            m := row*3 + col + 1

            highlight := 0
            if jy == cjy && m == cjm {
                highlight = cjd
            }

            cal := JalaliMonthCalendar(jy, m, highlight)
            lines := strings.Split(cal, "\n")
            months = append(months, lines)
        }

        maxLines := len(months[0])

        for i := 0; i < maxLines; i++ {
            for _, month := range months {
                if i < len(month) {
                    fmt.Printf("%-22s", month[i])
                }
            }
            fmt.Println()
        }

        fmt.Println()
    }
}
