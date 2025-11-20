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
