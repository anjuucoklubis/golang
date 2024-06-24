# Ringkasan Tipe Data Number

## 2 tipe jenis data number, yaitu:
-   ```Integer``` (bilangan bulat)
-   ```Floating Point``` (bilangan desimal)

### Tipe Integer (1)
    *negatif & positif atau int ( int)*
-   ```int8  ```  (min: -128, max: 127)
-   ```int16``` (min: -32768, max: 32767)
-   ```int32``` (min: -2147483648, max: 2147483647)
-   ```int64``` (min: -9223372036854775808, max: 9223372036854775807)

### Tipe Integer (2)  
    *tidak negatif atau uint (unsigned int)
-   ```uint8  ```  (min: 0, max: 255)
-   ```uint16``` (min: 0, max: 65535)
-   ```uint32``` (min: 0, max: 4294967295)
-   ```uint64``` (min: 0, max: 18446744073709551615)

### Tipe Floating Point (1)  
-   ```float32```  (min: 1.18 x 10^-38, max: 3.4 x 10^38)
-   ```float64``` (min: 2.23 x 10^-308, max: 1.80 x 10^308)
-   ```complex64``` (complex numbers with float32 real and imaginary parts)
-   ```complex128``` (complex numbers with float64 real and imaginary parts)

## ALIAS (nama lain):
-   ```byte``` (Alias untuk ```uint8```)
-   ```rune``` (Alias untuk ```int32```)
-   ```int ``` (Alias untuk ```Minimal int32```)
-   ```uint``` (Alias untuk ```Minimal uint32```)


## Notes:
-   Kenapa tidak menggunakan ```int64``` saja ketika angkanya kecil ataupun besar ?  **karena ketika menggunakan int64, nantinya ukuran yang harus disimpan didalam memory penyimpanan komputer kita semakin besar juga, jadi baiknnya kita sesuaikan dengan angka yang akan kita gunakan**

