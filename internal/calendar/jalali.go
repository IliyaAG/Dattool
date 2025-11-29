package calendar

import "time"

var gdm = [...]int{0, 31, 59, 90, 120, 151, 181, 212, 243, 273, 304, 334}

func GregorianToJalali(gy, gm, gd int) (jy, jm, jd int) {
    var jy2, days int

    if gy > 1600 {
        jy2 = 979
        gy -= 1600
    } else {
        jy2 = 0
        gy -= 621
    }

    if gm > 2 {
        days = 365*gy + (gy+3)/4 - (gy+99)/100 + (gy+399)/400 + gdm[gm-1] + gd - 80
    } else {
        days = 365*gy + (gy+3)/4 - (gy+99)/100 + (gy+399)/400 + gdm[gm-1] + gd - 79
    }

    jy2 += 33 * (days / 12053)
    days %= 12053
    jy2 += 4 * (days / 1461)
    days %= 1461

    if days > 365 {
        jy2 += (days - 1) / 365
        days = (days - 1) % 365
    }

    if days < 186 {
        jm = 1 + days/31
        jd = 1 + days%31
    } else {
        jm = 7 + (days-186)/30
        jd = 1 + (days-186)%30
    }
    jy = jy2
    return
}

func JalaliToGregorian(jy, jm, jd int) (gy, gm, gd int) {
    est := jy + 621

    start := time.Date(est-1, 1, 1, 0, 0, 0, 0, time.UTC)
    end := time.Date(est+2, 1, 1, 0, 0, 0, 0, time.UTC)

    for t := start; t.Before(end); t = t.AddDate(0, 0, 1) {
        y, m, d := t.Date()
        jy2, jm2, jd2 := GregorianToJalali(y, int(m), d)
        if jy2 == jy && jm2 == jm && jd2 == jd {
            return y, int(m), d
        }
    }

    return 0, 0, 0
}
