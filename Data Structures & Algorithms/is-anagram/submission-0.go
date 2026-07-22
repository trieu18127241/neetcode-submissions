func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    annagramS := make(map[string]int)
    annagramT := make(map[string]int)
    for _, val := range s {
        annagramS[string(val)] += 1
    }
    for _, val := range t {
        annagramT[string(val)] += 1
    }
    for val, numS := range annagramS {
        numT, ok := annagramT[val]; 
        fmt.Println("checking s: ", string(val), numS)
        fmt.Println("checking t: ", string(val), numT)
        if !ok || numS != numT {
            return false
        } 
    }
    return true
}
