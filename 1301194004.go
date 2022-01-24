/*  NIM     : 1301194004
    NAMA    : Yantrisnandra Akbar Maulino
    KELAS   : IF-43-02
    TOPIK   : D [ Transaksi Tarik Tunasi]
*/
package main

import "fmt"
import "time"
import "strings"
import "strconv"

//TYPE RECORD NASABAH
type nasabah struct{
    jenisrekening string
    nama string
    nik string
    norek string
    bayarterakhir int
    saldo int
    daftar string
    sekarang string
    waktu string
    now, jamsekarang string
}

//VARIABEL
const N int = 1000
type TypeArr [N]nasabah
var masukan TypeArr

//JUMLAH NASABAH
var total int
var validrekening int
var username string
var password string




// FUNGSI AREA
func ceknik(key string)bool{
    var inik int
    var cek int
    inik = 0
    for inik < total{
        if key == masukan[inik].nik {
            cek = cek + 1
        }else if key != masukan[inik].nik{
            cek =cek
        }
        inik = inik + 1
    }
    return cek == 0
}

// WAKTU
func cekday()string {
        waktusekarang := time.Now()
        Year := waktusekarang.Year()
        Month:= int(waktusekarang.Month())
        Day  := waktusekarang.Day()



        Tahun := strconv.Itoa(Year)
        Bulan := strconv.Itoa(Month)
        Tanggal := strconv.Itoa(Day)
       

        hari := []string{Tanggal ,Bulan, Tahun}
        sekarang := strings.Join(hari, "/")

        haridaftar := sekarang

        return haridaftar
}



func cektime()string{
    waktusekarang := time.Now()
    Hour := waktusekarang.Hour()
    Minute := waktusekarang.Minute()
    Second := waktusekarang.Second()

    Jam := strconv.Itoa(Hour)
    Menit := strconv.Itoa(Minute)
    Detik  := strconv.Itoa(Second)

    waktu := []string{Jam, Menit, Detik}
    today := strings.Join(waktu, ":")

    waktudaftar := today

    return waktudaftar
}



// NAMA
func nama()string{

    var namalengkap, namadepan,namabelakang string
    
    fmt.Print("Nama Depan: ")
    fmt.Scan(&namadepan)
    fmt.Print("Nama Belakang: ")
    fmt.Scan(&namabelakang)

    lengkap := []string{namadepan, namabelakang}
    namalengkap = strings.Join(lengkap, " ")

    return namalengkap
}

//PROSEDUR AREA

func Customer(silver, gold, platinum *int){
    var saldo, setoran int
    var transaksiterakhir int
    var key string
    var norek string
    var n1, n2, n3 int
    var jenis string
    var namalengkap string
    var keluar string
    var i int
    var valid int


    total = total + 1

    fmt.Println("|==========================================================|")
    fmt.Println("|  LOGIN CUSTOMER SERVICE, MASUKKAN USERNAME DAN PASSWORD  |")
    fmt.Println("|==========================================================|")
    fmt.Print("Username Anda : ")
    fmt.Scan(&username)
    fmt.Print("Password Anda : ")
    fmt.Scan(&password)

    for username != "Akbar" || password != "Maulino" {
        fmt.Println("Password yang Anda Masukan salah")
        fmt.Println("masukan username dan  Password Lagi")
        fmt.Print("Username Anda : ")
        fmt.Scan(&username)
        fmt.Print("Password Anda : ")
        fmt.Scan(&password)

    }


    fmt.Println("|==========================================================|")
    fmt.Println("|============== | Pendaftaran Nasabah Baru | ==============|")
    fmt.Println("|==========================================================|\n")



   //NIK & Validasi
   fmt.Println("|==========================================================|")
   fmt.Println("|================ | Syarat Memasukan NIK | ================|")
   fmt.Println("|==========================================================|")
   fmt.Println("| 1. NIK Valid Jika Berupa 16 Digit Angka                  |")
   fmt.Println("| 2. NIK Tidak Boleh Duplikat                              |")
   fmt.Println("|==========================================================|")
   fmt.Println("|==========================================================|") 
   fmt.Print("Masukan NIK Baru : ")
   fmt.Scan(&key)
   digit := strings.Split(key,"")
   valid = 0
   i = 0

   for i < len(digit) && len(digit) == 16 && ceknik(key) {
        if key[i] >= 48 && key[i] <= 57 {
            valid = valid + 1
        }
    i = i + 1
    }
    for valid != 16 {
        fmt.Println("|==========================================================|")
        fmt.Println("|        Tidak Valid, Masukkan Lagi NIK yang Benar         |")
        fmt.Println("|==========================================================|")
            
        fmt.Print("Masukkan NIK Lain : ")
        fmt.Scan(&key)
        digit := strings.Split(key, "")
        valid = 0
        i = 0
        for i < len(digit) && len(digit) == 16 && ceknik(key) {
            if key[i] >= 48 && key[i] <= 57 {
                valid = valid + 1
        }
    i++
    }
}




    //Nama
    namalengkap = nama() 



    //Jenis Rekening
    fmt.Println("|==========================================================|")
    fmt.Println("|=================== | Jenis Rekening | ===================|")
    fmt.Println("|==========================================================|")
    fmt.Println("| 1. Silver                                                |")
    fmt.Println("| 2. Gold                                                  |")
    fmt.Println("| 3. Platinum                                              |")
    fmt.Println("|==========================================================|")
    fmt.Println("|==========================================================|")
    fmt.Print("Jenis Rekening: ")
    fmt.Scan(&jenis)
    jenis = strings.ToLower(jenis)

    for (jenis != "silver") && (jenis != "platinum") && (jenis != "gold") {
        fmt.Println("|==========================================================|")
        fmt.Println("|======== | Jenis yang Anda Masukkan Tidak Valid | ========|")
        fmt.Println("|==========================================================|\n")
        fmt.Print("Masukkan Jenis Rekening Yang Valid: ")
        fmt.Scan(&jenis)
    }

    

    //No Rekening
    if jenis == "silver" {
        *silver++
        n1 = *silver / 100
        n2 = *silver / 10
        n3 = *silver % 10
    }else if jenis == "platinum" {
        *platinum++
        n1 = *platinum / 100
        n2 = *platinum / 10
        n3 = *platinum % 10
    }else if jenis == "gold" {
        *gold++
        n1 = *gold / 100
        n2 = *gold / 10
        n3 = *gold % 10
    }

    a := strconv.Itoa(n1)
    b := strconv.Itoa(n2)
    c := strconv.Itoa(n3)
    
    if jenis == "silver"{
        norek = "XYZ-S" + a + b + c
    } else if jenis == "platinum" {
        norek = "XYZ-P"+ a + b + c
    }else if jenis == "gold" {
        norek = "XYZ-G"+ a + b + c
    }
    


    // Setoran
    fmt.Println("|==========================================================|")
    fmt.Println("|==================== | Setoran Awal | ====================|")
    fmt.Println("| 1. Setoran Awal Minimal 7500                             |")
    fmt.Println("| 2. Setoran Awal Harus Bernilai Positif                   |")
    fmt.Println("|==========================================================|")
    fmt.Println("|==========================================================|")
    fmt.Print("Masukan Setoran Awal: Rp. ")
    fmt.Scan(&setoran) 
    
    for setoran < 7500 {
        fmt.Print("Masukkan Setoran Awal: Rp.")
        fmt.Scan(&setoran)
    }
    transaksiterakhir = setoran



    // Saldo
    saldo = 0
    saldo = saldo + setoran


   // waktu
    sekarang := cekday()
    haridaftar := sekarang
    now := cektime()
    waktudaftar := now

    masukan[total].nama = namalengkap
    masukan[total].nik = key
    masukan[total].bayarterakhir = transaksiterakhir
    masukan[total].jenisrekening = jenis 
    masukan[total].norek = norek
    masukan[total].saldo = saldo
    masukan[total].daftar = haridaftar
    masukan[total].sekarang=sekarang
    masukan[total].now = waktudaftar

        fmt.Println("|==============================================================================================|")
        fmt.Println("|======================= | Selamat Anda Sudah Menjadi Nasabah Bank XYZ | ======================|")
        fmt.Println("|==============================================================================================|\n")
        fmt.Println("|====================================== | DATA NASABAH | ======================================|")
        fmt.Println("Nomor Rekening             :",masukan[total].norek)
        fmt.Println("NIK                        :",masukan[total].nik)
        fmt.Println("Nama                       :",masukan[total].nama)                
        fmt.Println("Jenis Rekening             :",masukan[total].jenisrekening)
        fmt.Println("Saldo                      :",masukan[total].saldo)
        fmt.Println("Nominal Transaksi Terakhir :",masukan[total].bayarterakhir)
        fmt.Println("Tanggal Pendaftaran        :",masukan[total].daftar)
        fmt.Println("Jam Pendaftaran            :",masukan[total].now)
        fmt.Println("Tanggal Transaksi Terakhir :",masukan[total].sekarang)
        fmt.Println("|==============================================================================================|")
        fmt.Println("|========================= | Ketik menu Untuk Kembali ke menu awal | ==========================|")
        fmt.Println("|==============================================================================================|")
        fmt.Println("Masukan Pilihan : ")
        fmt.Scan(&keluar)

        for keluar != "menu" {
            fmt.Println("Masukan Pilihan : ")
            fmt.Scan(&keluar)
        }
}



func teller(){
    var noreki string
    var penarikan int
    var hitungrek, cekrekening int
    var keluar string

// PENCARIAN NO REKENING
    fmt.Println("|==========================================================|")
    fmt.Println("|       LOGIN TELLER, MASUKKAN USERNAME DAN PASSWORD       |")
    fmt.Println("|==========================================================|")
    fmt.Print("Username Anda : ")
    fmt.Scan(&username)
    fmt.Print("Password Anda : ")
    fmt.Scan(&password)

    for username != "Akbar" || password != "Maulino" {
        fmt.Println("Passwor Yang Anda Masukan Salah")
        fmt.Println("Masukan Username Dan Password Lagi")
        fmt.Print("Username Anda : ")
        fmt.Scan(&username)
        fmt.Print("Password Anda : ")
        fmt.Scan(&password) 
    }

        hitungrek = 0
        fmt.Print("Masukkan Nomor Rekening Yang Sudah Terdaftar : ")
        fmt.Scan(&noreki)
        for (hitungrek < N-1) && (masukan[hitungrek].norek != noreki) {
            hitungrek = hitungrek + 1
        }
        if masukan[hitungrek].norek == noreki {
            cekrekening = hitungrek
        }   


        if masukan[cekrekening].jenisrekening == "silver" {
            fmt.Println("|=========================================================================|")
            fmt.Println("|============== | Maximal Penarikan Uang Sebesar Rp.40000 | ==============|")
            fmt.Println("|=========================================================================|")
            fmt.Print("Masukkan jumlah penarikan : Rp. ")
            fmt.Scan(&penarikan)
            for (penarikan % 10 != 0) || (penarikan >= 40000 && penarikan > 0 ) {
                fmt.Print("Masukkan jumlah penarikan lagi : Rp. ")
                fmt.Scan(&penarikan)
            }    
        } else if masukan[cekrekening].jenisrekening == "gold" {
            fmt.Println("|=========================================================================|")
            fmt.Println("|============= | Maximal Penarikan Uang Sebesar Rp.180000 | ==============|")
            fmt.Println("|=========================================================================|")
            fmt.Print("Masukkan jumlah penarikan : Rp.  ")
            fmt.Scan(&penarikan)
            for (penarikan % 10 != 0) || (penarikan >= 180000 && penarikan > 0 ) {
                fmt.Print("Masukkan jumlah penarikan lagi : Rp. ")
                fmt.Scan(&penarikan)
            }
        } else if masukan[cekrekening].jenisrekening == "platinum" {
            fmt.Println("|=========================================================================|")
            fmt.Println("|============= | Maximal Penarikan Uang Sebesar Rp.450000 | ==============|")
            fmt.Println("|=========================================================================|")
            fmt.Print("Masukkan jumlah penarikan : Rp. ")
            fmt.Scan(&penarikan)
            for (penarikan % 10 != 0) || (penarikan >= 450000 && penarikan > 0 ) {
                fmt.Print("Masukkan jumlah penarikan lagi: ")
                fmt.Scan(&penarikan)
            }
        }
        for penarikan < 0{
            fmt.Println("Masukan jumlah penarikan lagi : ")
            fmt.Scan(&penarikan)
        }
            masukan[cekrekening].saldo = masukan[cekrekening].saldo - penarikan
            masukan[cekrekening].bayarterakhir = penarikan


        //WAKTU
        masukan[cekrekening].sekarang = cekday()
        masukan[cekrekening].jamsekarang = cektime()
        fmt.Println("|==============================================================================================|")
        fmt.Println("|==================================== | DATA NASABAH | ========================================|")
        fmt.Println("Nomor Rekening             :", masukan[cekrekening].norek )
        fmt.Println("NIK                        :", masukan[cekrekening].nik)
        fmt.Println("Nama                       :", masukan[cekrekening].nama)
        fmt.Println("Jenis Rekening             :", masukan[cekrekening].jenisrekening)
        fmt.Println("Saldo                      :", masukan[cekrekening].saldo)
        fmt.Println("Nominal Transaksi Terakhir :", masukan[cekrekening].bayarterakhir )
        fmt.Println("Tanggal Pendaftaran        :", masukan[cekrekening].daftar)
        fmt.Println("Waktu Pendaftaran          :", masukan[cekrekening].now)
        fmt.Println("Tanggal Transaksi Terakhir :", masukan[cekrekening].sekarang)
        fmt.Println("Jam Transaksi Terakhir     :", masukan[cekrekening].jamsekarang)
        fmt.Println("|==============================================================================================|")
        fmt.Println("|============================ | Ketik Menu Untuk Kembali ke Menu awal | =======================|")
        fmt.Println("|==============================================================================================|")
        fmt.Println("Masukan Pilihan : ")
        fmt.Scan(&keluar)
        for keluar != "menu"{
            fmt.Println("Masukan Pilihan : ")
            fmt.Scan(&keluar)
        }
}




func manager(){
    var noreki string
    var keluar string
    var counterrek int
    var setahun int
    var pilih int
    var counterwaktu, thnnasabah, blnnasabah int


    fmt.Println("|==========================================================|")
    fmt.Println("|      LOGIN MANAGER, MASUKKAN USERNAME DAN PASSWORD       |")
    fmt.Println("|==========================================================|")
    fmt.Print("Username Anda : ")
    fmt.Scan(&username)
    fmt.Print("Password Anda : ")
    fmt.Scan(&password)

    for username != "Akbar" || password != "Maulino" {
        fmt.Println("Password yang Anda Masukan salah")
        fmt.Println("Masukan Username Dan Password Anda Lagi ")
        fmt.Print("Username Anda : ")
        fmt.Scan(&username)
        fmt.Print("Password Anda : ")
        fmt.Scan(&password)
    }
    fmt.Println("|=================================================================================================|")
    fmt.Println("|========================================== | MENU | =============================================|")
    fmt.Println("|=================================================================================================|")
    fmt.Println("| 1. Menampilkan data Nasabah                                                                     |")
    fmt.Println("| 2. Menampilkan jumlah Data Nasabah terurut berdasarkan nama                                     |")
    fmt.Println("| 3. Menampilkan jumlah dan data semua nasabah yang tidak melakukan transaksi lebih dari 12 bulan |")
    fmt.Println("| 4. Menampilkan jumlah dan data nasabah per jenis rekening                                       |")
    fmt.Println("|=================================================================================================|")
    fmt.Println("|=================================================================================================|")
    fmt.Print("Masukkan pilihan : ")
    fmt.Scan(&pilih)


        if pilih == 1 {


            fmt.Print("Masukkan Nomor Rekening: ")
            fmt.Scan(&noreki)

            counterrek = 0
            for (counterrek < N-1) && (masukan[counterrek].norek != noreki) {
                counterrek = counterrek + 1
            }
            if masukan[counterrek].norek == noreki {
                validrekening = counterrek
                fmt.Println("|==================================== | DATA NASABAH | ========================================|")
                fmt.Println("No Rekening                :",masukan[counterrek].norek)
                fmt.Println("NIK                        :",masukan[counterrek].nik )
                fmt.Println("Nama                       :",masukan[counterrek].nama )
                fmt.Println("Jenis Rekening             :",masukan[counterrek].jenisrekening )
                fmt.Println("Saldo                      :",masukan[counterrek].saldo )
                fmt.Println("Nominal Transaksi Terakhir :",masukan[counterrek].bayarterakhir )
                fmt.Println("Tanggal Pendaftaran        :",masukan[counterrek].daftar)
                fmt.Println("Tanggal Transaksi Terakhir :",masukan[counterrek].sekarang)
                fmt.Println("|=============================================================================================|")
                fmt.Println("|=========================== | Pilih menu Untuk Kembali ke menu awal | =======================|")
                fmt.Println("|=============================================================================================|")
                fmt.Println("Masukan Pilihan : ")
                fmt.Scan(&keluar)
                for keluar != "menu"{
                    fmt.Println("Masukan Pilihan : ")
                    fmt.Scan(&keluar)
                }

    

            }else  {
                fmt.Println("Nomor Rekening Tidak Ditemukan ")
                fmt.Println("Masukan : ")
                fmt.Scan(&keluar)
                for keluar == "menu" {
                    fmt.Println("Masukan Pilihan :")
                    fmt.Scan(&keluar)
                }
            }


        }else if pilih == 2 {


                tampilSemua()
                fmt.Println("|======================================================================================================|")
                fmt.Println("|================================ | Pilih menu Untuk Kembali ke menu awal | ===========================|")
                fmt.Println("|======================================================================================================|")
                fmt.Print("Masukkan : ")
                fmt.Scan(&keluar)
                for keluar != "menu" {
                    fmt.Println("Masukkan :")
                    fmt.Scan(&keluar)
                }



        }else if pilih == 3 {

            waktu := time.Now()
            counterwaktu = 1

            for counterwaktu <= total {
                waktuterakhir := masukan[counterwaktu].sekarang
                pecah := strings.Split(waktuterakhir, "/")
                thnnasabah,_  = strconv.Atoi(pecah[2])
                blnnasabah,_  = strconv.Atoi(pecah[1])
                tahunsekarang := waktu.Year()
                bulansekarang := int(waktu.Month())
                
                if (tahunsekarang - thnnasabah == 1  && bulansekarang - blnnasabah == 0) || (tahunsekarang - thnnasabah > 2 ) {
                    fmt.Println("|==================================== | DATA NASABAH | ========================================|")
                    fmt.Println("Nama                       :", masukan[counterwaktu].nama)
                    fmt.Println("Nomor Rekening             :", masukan[counterwaktu].norek )
                    fmt.Println("NIK                        :", masukan[counterwaktu].nik)
                    fmt.Println("Jenis Rekening             :", masukan[counterwaktu].jenisrekening)
                    fmt.Println("Saldo                      :", masukan[counterwaktu].saldo)
                    fmt.Println("Nominal Transaksi Terakhir :", masukan[counterwaktu].bayarterakhir )
                    fmt.Println("Tanggal Pendaftaran        :", masukan[counterwaktu].daftar)
                    fmt.Println("Tanggal Transaksi Terakhir :", masukan[counterwaktu].sekarang)
                    fmt.Println("|==============================================================================================|\n")
                    setahun = setahun + 1
                }  
                counterwaktu++
           }
           if setahun == 0 {
                fmt.Println("|=========================================================================================================================================================|")
                fmt.Println("|======================================= | TIDAK ADA NASABAH YANG TIDAK MELAKUKAN TRANSAKSI LEBIH DARI 12 BULAN | ========================================|")
                fmt.Println("|=========================================================================================================================================================|")
                fmt.Println()
           }
           fmt.Println("| JUMLAH NASABAH : ",setahun," | ")
           
            fmt.Println("|=============================================================================================|")
            fmt.Println("|=========================== | Pilih menu Untuk Kembali ke menu awal | =======================|")
            fmt.Println("|=============================================================================================|")
            fmt.Print("Masukkan : ")
            fmt.Scan(&keluar)
            for keluar != "menu" {
                fmt.Println("Masukkan :")
                fmt.Scan(&keluar)
            }

        }else if pilih == 4 {

            tampilRekening()

            fmt.Println("|=========================================================================================================================|")
            fmt.Println("|========================================= | Pilih menu Untuk Kembali ke menu awal | =====================================|")
            fmt.Println("|=========================================================================================================================|")
            fmt.Print("Masukkan : ")
            fmt.Scan(&keluar)
            for keluar != "menu" {
                fmt.Println("Masukkan :")
                fmt.Scan(&keluar)
            }
        }
    }






    func tampilSemua(){
    var  j, max int
    var t nasabah
    for i:=0;i < total ;i++{
        max = i
        j = i + 1
        for j < total {
            if masukan[max].norek > masukan[j].norek {
                max = j
            }
            j++
        }
        t = masukan[i]
        masukan[i] = masukan[max]
        masukan[max] = t
    }
    fmt.Println("|======================================================================================================|")
    fmt.Println("|==================================== | DATA SELURUH NASABAH | ========================================|")
    fmt.Println("|======================================================================================================|\n")
    j = 1
    for i:=1; i <= total;i++{
        fmt.Println("|================= |",j,"| ==================|")
        j++
        fmt.Print("NIK anda adalah     : ")
        fmt.Println(masukan[i].nik)
        fmt.Print("Nama                : ")
        fmt.Println(masukan[i].nama)
        fmt.Print("Jenis rekening      : ")
        fmt.Println(masukan[i].jenisrekening)
        fmt.Print("Saldo               : ")
        fmt.Println(masukan[i].saldo)
        fmt.Print("Nomor rekening      : ")
        fmt.Println(masukan[i].norek)
        fmt.Print("Tanggal Pendaftaran :")
        fmt.Println(masukan[i].daftar)
        fmt.Print("Jam Pendaftaran     :")
        fmt.Println(masukan[i].now)
        
    }
    if j == 0{
        fmt.Println("|==================================================================================================================|")
        fmt.Println("|==================================== | TIDAK ADA NASABAH YANG TERDAFTAR | ========================================|")
        fmt.Println("|==================================================================================================================|\n")
    }else{
    	fmt.Println()
        fmt.Println("Terdapat ", (j-1), " nasabah di bank XYZ\n")
    }
}



func tampilRekening(){
    var valid bool
    var t nasabah
    var j  int
    var rekening string
    

    fmt.Println("|===================================================================================================================|")
    fmt.Println("|==================================== | DATA NASABAH BERDASARKAN REKENING | ========================================|")
    fmt.Println("|===================================================================================================================|\n")
    fmt.Print("Masukkan jenis rekening untuk menampilkan data : ")
    fmt.Scan(&rekening)
    valid = false
    for !valid{
        if rekening != "silver" && rekening != "gold" && rekening != "platinum"{
            fmt.Println("Jenis rekening tersebut tidak ada !\n")
            valid = false
            fmt.Print("Masukkan jenis rekening untuk menampilkan data : ")
            fmt.Scan(&rekening)
        }else{
            valid = true
        }
    }
    i := 1
    for i < total{
        t = masukan[i]
        j = i - 1
        for j >= 0 &&  masukan[j].norek > t.norek{
            masukan[j+1] = masukan[j]
            j = j - 1
        }
        masukan[j+1] = t
        i++
    }
    
    k:=0
    for l:=1; l <= total; l++{
        if masukan[l].jenisrekening == rekening{
            k++
            fmt.Println("================= ",k," ===================")
            fmt.Print("NIK anda adalah     : ")
            fmt.Println(masukan[l].nik)
            fmt.Print("Nama                : ")
            fmt.Println(masukan[l].nama)
            fmt.Print("Jenis rekening      : ")
            fmt.Println(masukan[l].jenisrekening)
            fmt.Print("Saldo               : ")
            fmt.Println(masukan[l].saldo)
            fmt.Print("Nomor rekening      : ")
            fmt.Println(masukan[l].norek)
        }
    }
    if k == 0{
        fmt.Println("|=========================================================================================================================|")
        fmt.Println("|==================================== | TIDAK ADA JENIS REKENING YANG TERDAFTAR | ========================================|")
        fmt.Println("|=========================================================================================================================|\n")
    }else{
        fmt.Println("Jumlah data yang terdapat pada jenis rekening ", rekening," adalah ", k)
    }
}



func main(){
    var pilihan, silver, gold, platinum int



    silver = 0
    gold = 0
    platinum = 0

    pilihan = 0
    for pilihan != 4 {
        fmt.Println("|==========================================================|")
        fmt.Println("|====== | Selamat Datang di Transaksi Tarik Tunai | =======|")
        fmt.Println("|==========================================================|")
        fmt.Println("Silakan Pilih Menu Dibawah Ini")
        fmt.Println("1.Login Customer Service")
        fmt.Println("2.Login Teller")
        fmt.Println("3.Login Manager")
        fmt.Println("4.Exit")
        fmt.Print("Masukan Pilihan  : ")
        fmt.Scan(&pilihan)

        if pilihan == 1 {

            Customer(&silver, &gold, &platinum)
        }else if pilihan == 2 {

                 teller()
            
            
             // TRANSAKSI 2
        } else if pilihan == 3 {
            
            manager()
        
        }
    }
        if pilihan == 4 {
        fmt.Println("|==========================================================|")
        fmt.Println("|==================== | Terima Kasih | ====================|")
        fmt.Println("|==========================================================|")
    }   
   


}