# Ringkasan Variable

-   ```Variable``` adalah untuk menyimpan data
-   Digunakan untuk mengakses data yang sama dimanapun
-   Pada Golang, ```Variable``` hanya bisa menyimpan ```tipe data``` yang sama, jadi harus membuat beberapa variable untuk menyimpan data yang berbeda tipe
-   Menggunakan keyword ```var```, ikuti ```nama variable``` dan ```type data``` nya. 

## Tipe Data Variable (tidak wajib menggunakan typedata)

-   Ketika kita membuat ```variable```, wajib menyertakan ```tipe data``` variable tersebut
-   Namun, kalau kita langsung memasukan ``` data``` atau ```value``` dari variable tersebut, tidak wajib untuk menyertakan atau menyebutkan ```tipe data``` dari variable tersebut.


## Kata Kunci Var

-   di ```Golang```, kata kunci ```var``` saat membuat variable tidalah wajib.
-   Asalkan saat membuat variable langsung diberikan atau disebutkan ```value``` atau ```datanya```
-   Agar tidak perlu menggunakan ```var```, kita harus menggunakan kata kunci ``` :=``` saat menginisialisasikan data pada variable tersebut.

## Deklrasi Multi Variable

-   di ```Golang```, kita bisa membuat ```variable``` secara sekaligus banyak
-   Code yang dbuat akan lebih bagus dan mudah di baca

### Notes :
-   Jika variable dengan menggunakan ```" := "``` sudah digunakan, dan ketika ingin mengganti ```value``` dari variable sebelumnya ?, kita harus menggunakan ```" = "``` , tidak boleh lagi menggunakan ```":="```, karena itu dianggap ```membuat ulang``` variable yang sudah ada.

-   ```variable``` yang telah dibuat dan di deklarasi, harus digunakan akan di ```call / use ```, kalau tidak nanti errro, jadi semua ```variable``` di golang yang dibuat dan di deklarasi, harus digunakan.