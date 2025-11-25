package calendar

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
    if jy > 979 {
        gy = 1600
        jy -= 979
    } else {
        gy = 621
    }

    days := 365*jy + jy/33*8 + (jy%33+3)/4

    if jm < 7 {
        days += (jm - 1) * 31
    } else {
        days += (jm-7)*30 + 186
    }

    days += jd - 1

    gy += 400 * (days / 146097)
    days %= 146097

    leap := true
    if days >= 36525 {
        days--
        gy += 100 * (days / 36524)
        days %= 36524

        if days >= 365 {
            days++
        } else {
            leap = false
        }
    }

    gy += 4 * (days / 1461)
    days %= 1461

    if days >= 366 {
        leap = false
        days--
        gy += days / 365
        days %= 365
    }

    // month and day
    g_dm := [...]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
    if leap {
        g_dm[1] = 29
    }

    var i int
    for i = 0; i < 12 && days >= g_dm[i]; i++ {
        days -= g_dm[i]
    }

    gm = i + 1
    gd = days + 1

    return
}
