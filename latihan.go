package main
import "fmt"

func main(){
	var bulan, besaran, untung int
	var bulan1, besaran1, untung1 int
	var hasil, hasil1 int
	fmt.Scan(&bulan, &besaran, &untung)
	fmt.Scan(&bulan1, &besaran1, &untung1)
	hasil = bulan*besaran/untung
	hasil1 = bulan1*besaran1/untung1
	if hasil > hasil1{
		fmt.Print("Produk A")
	}else{
		fmt.Print("Produk B")
	}

}