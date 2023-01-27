ANSWER

QUESTION NO 1 

    Hasil dari pertanyaan nomor 1 akan tampil di log saat menjalankan program, inputnya digunakan sebagai parameter pada func MyAnswer() di answer/my_answer.go

QUESTION NO 2

    Hasil dari pertanyaan nomor 2 akan tampil di log saat menjalankan program, inputnya digunakan sebagai parameter pada func MyAnswer() di answer/my_answer.go

QUESTION NO 3

EXISTING

	1.FROM golang
	2.ADD . /go/src/github.com/telkomdev/indihome/backend
	3.WORKDIR /go/src/github.com/telkomdev/indihome
	4.RUN go get github.com/tools/godep
	5.RUN godep restore
	6.RUN go install github.com/telkomdev/indihome
	7.ENTRYPOINT /go/bin/indihome
	8.LISTEN 80

NEW

	1.FROM golang
	2.ADD . /go/src/github.com/telkomdev/indihome
	3.WORKDIR /go/src/github.com/telkomdev/indihome
	4.RUN go install github.com/tools/godep@latest
	5.RUN godep save
	6.RUN godep restore
	7.RUN go build -o indihome.exe
	8.ENTRYPOINT ./indihome.exe
	9.EXPOSE 80

EXPLANATION

	1.Hapus /backend pada line 2 agar sesuai dengan WORKDIR pada line 3
    2.Ubah command go get menjadi go install pada line 4 dan tambahkan @latest sebagai versinya agar command godep pada line selanjutnya dapat berjalan
    3.Tambahkan command RUN godep save sebelum RUN godep restore untuk mendaftarkan depedency yang ada ke Godeps.json
    4.Ubah command go install pada line 6 menjadi go build agar dapat menghasilkan binary file yang bisa di run
    5.Karena hasil binary file yang sudah dibuild sebelumnya terdapat di directory luar, ganti pathnya menjadi ./outputnamefile.exe
    6.Command LISTEN salah, seharusnya menggunakan EXPOSE, namun saat di expose port tersebut belum bisa digunakan selengkapnya harus menggunakan command -p saat docker run

QUESTION NO 4

    Menurut saya tujuan penggunaan microservices adalah untuk mempermudah proses improvement atau bug fixing dari sebuah fitur yang ada didalam aplikasi, 
    contoh sederhananya adalah jika terdapat bug atau error pada suatu fitur hingga mengakibatkan exit program jika aplikasi tersebut menggunakan monolith maka akan berdampak pada fitur lainnya atau keseluruhan aplikasi namun jika menggunakan microservices maka hanya service atau fitur itu saja yang akan kena dampaknya

QUESTION NO 5

    Indexing merupakan proses untuk mempercepat pencarian query, menurut saya indexing memiliki proses dengan memfokuskan proses pencarian pada kolom yang diindexing. Tanpa indexing proses pencarian akan diproses secara merata pada semua kolom dan menghasilkan query yang lambat jika jumlah row dan kolom banyak

QUESTION NO 6
    
    1.Ubah value di file .env sesuai dengan konfigurasi local (databasenya menggunakan PostgreSQL)
    2.Tambahkan database sesuai dengan PG_DBNAME di .env
    3.Jalankan perintah 'go run main.go'
    CONTOH REQUEST:

    METHOD GET
    localhost:8080/cart?nama_produk=produk A&kuantitas=18

    METHOD POST
    localhost:8080/cart
        BODY
        {
            "nama_produk": "produk C",
            "kuantitas": 9
        }
    
    METHOD DELETE
    localhost:8080/cart/f29f3139-f0b1-429c-b2c8-5e9ad0ce2314