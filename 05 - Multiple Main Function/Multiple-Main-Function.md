# Ringkasan Multiple Main Function

## Multiple Main Function
-   Pada Golang, ```functions``` bersifat unique, tidak boleh ada main function dengan nama yang sama
-   Jadi kalau kita buat ```functions``` yang sama pada 1 package, yang menjadi issue ketika kita menjalankan ```go build``` maka akan terjadi error ```duplicate declaration function```.

## Solusi ?
-   sekarang jalankan file satu persatu
-   Tapi nanti, di project asli, kita hnya akan membuat satu main function saja.

