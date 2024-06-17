package main

func trocar(p1, p2 int) (segundo int, primeiro int) {
	segundo = p2
    primeiro = p1
    return
}

func main() {
    a, b := trocar(1, 2)
    println(a, b)
}