fun insertSort(arr: Array<Int>) {
    for (i in 1 until arr.size) {
        //println("loop i = $i")
        for (j in (i - 1) downTo 0) {
            //println("loop j = $j")
            if (arr[j] > arr[j + 1]) {
                // swap element j and j + 1
                arr[j] = arr[j + 1].also {arr[j + 1] = arr[j]}
            } else {
                //println("break, i = $i, j = $j")
                break
            }
        }
    }
}

fun printArray(arr: Array<Int>) {
    print("(")
    for (i in arr) {
        print("$i, ")
    }
    println(")")
}

fun main() {
    println("Hello, World!")
    val a = arrayOf(45, 34, 78, 12, 34, 32, 29, 64)
    val b = arrayOf(12, 29, 32, 34, 34, 45, 64, 78)
    insertSort(a)

    printArray(a)
    for (i in 0 until a.size) {
        if (a[i] != b[i]) {
            println("ERROR: NOT Equals!!! a[$i](${a[i]}) != b[$i](${b[i]})")
            return
        }
    }

    println("Equals!!!")
}
