package computation

import "strings"

func Strminus(a, b string)string{
   return strings.ReplaceAll(a, b, "")
}

func Umnostr(a string, numstr int)string{
    return strings.Repeat(a, numstr)
}

func Delstr(a string, numstr int)string{
    leght := len(a)
    if numstr <= 0 || leght < numstr{
        return ""
    }
    return a[:leght/numstr]
}
