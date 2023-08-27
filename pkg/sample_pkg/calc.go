package sample


func Calculate(i ...int) int {
    var num int
    for _, v := range i {
        num += v
    }

    return num
}