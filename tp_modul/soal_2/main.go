package main

import "fmt"

const NMAX int = 5

type tabInt [NMAX]int

func bacaData(A *tabInt, n *int) {
	var x int
	fmt.Scan(&x)
	if *n < NMAX {
		A[*n] = x
		*n++
	} else {
		fmt.Println("Array penuh")
	}
}

func cetakData(A tabInt, n int) {
	if n == 0 {
		fmt.Println("Array kosong")
	} else {
		for i := 0; i < n; i++ {
			fmt.Print(A[i], " ")
		}
		fmt.Println()
	}
}

func hapusData(A *tabInt, n *int) {
	if *n > 0 {
		*n--
	} else {
		fmt.Println("Array tidak bisa dihapus")
	}
}

func main() {
	var data tabInt
	var nData int

	bacaData(&data, &nData)
	bacaData(&data, &nData)
	cetakData(data, nData)
	hapusData(&data, &nData)
	cetakData(data, nData)
	hapusData(&data, &nData)
	cetakData(data, nData)
	nData = 0
	bacaData(&data, &nData)
	bacaData(&data, &nData)
	cetakData(data, nData)
	hapusData(&data, &nData)
	cetakData(data, nData)
	hapusData(&data, &nData)
	cetakData(data, nData)
	hapusData(&data, &nData)
	cetakData(data, nData)
}