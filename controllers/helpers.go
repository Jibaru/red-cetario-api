package controllers

import "strconv"

func parseUint(s string) uint {
    u, _ := strconv.ParseUint(s, 10, 64)
    return uint(u)
}
