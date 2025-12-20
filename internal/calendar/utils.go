package calendar

func PrevMonth(y, m int) (int, int) {
    if m == 1 {
        return y - 1, 12
    }
    return y, m - 1
}

func NextMonth(y, m int) (int, int) {
    if m == 12 {
        return y + 1, 1
    }
    return y, m + 1
}
