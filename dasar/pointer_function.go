package main

import "fmt"

type Address struct {
	city, province, country string
}

func changeCountryToIndonesia(address *Address) {
	address.country = "Indonesia"
}

func main() {
	var address1 Address = Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	var address2 *Address = &address1 // & tanda ini di maksudkan menjadikan variable address2 itu sebagai pointer ke variable address1
	var address3 *Address = &address1
	var address4 *Address = new(Address)

	address2.city = "Bhedeng"

	*address3 = Address{"Magelang", "Jawa Tengah", "Indonesia"} // penambahan operator * akan memaksa semua value berubah pada memory yang sama
	address4.city = "Kenthu"
	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)
	fmt.Println(address4)

	var alamat = Address{"Jakarta", "DKI Jakarta", ""}
	changeCountryToIndonesia(&alamat)
	fmt.Println(alamat)

}
