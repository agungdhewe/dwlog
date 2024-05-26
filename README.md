# DwLogger
Suatu modul lightweight logger untuk development applikasi golang, serta idependent dari 3rd party modul lain.
Modul ini hanya menambahkan beberapa fungsi ke logger standard golang, dengan fungsi `Info()`, `Log()`, `Warning()` dan `Error()`.
Fungsi-fungsi tersebut digunakan lebih untuk keperluan logging di screen, dengan menambahkan beberapa visual sehingga diharapkan lebih memudahkan proses debugging. 

### Contoh Penggunaan
    dlog  :=  dwlog.New()
    dlog.Info("ini info")
    dlog.Log("coba screen logging script")
    dlog.Warning("ini adalah warning")
    dlog.Error("ini log untuk error")
    
### Contoh Tampilan saat running
![contoh tampilan](https://transfashionindonesia.com/wp-content/uploads/2023/12/AIGNER-Gift-Ideas-2023-1-selected-400x400-1.jpg)
