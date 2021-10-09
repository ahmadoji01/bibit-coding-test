# bibit-coding-test
 
Berikut adalah daftar isi dari petunjuk pengecekan soal tes Bibit menggunakan Go

- [No. 1: Simple SQL Command]
- [No. 2: Go REST API dan gRPC Menggunakan API dari OMDB](#go-omdb)
- [No. 3: Refactor Code Go]
- [No. 4: Klasifikasi Kata Anagram]

***

### <a name="prep"></a>No. 2: Go REST API dan gRPC Menggunakan API dari OMDB

Untuk menjalankan program ini, diperlukan melakukan instalasi dengan menjalankan perintah berikut

    cd no2
    glide install -v

Lalu untuk menjalankan server dapat menggunakan perintah berikut

    go run app/main.go

Untuk pengecekan gRPC, dapat menjalankan program client dengan perintah berikut

    go run app/grpc_client/main.go

Untuk pengecekan unit test, dapat dilakukan dengan perintah berikut

    cd movie/usecase
    go test