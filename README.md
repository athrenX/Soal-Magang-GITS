# Soal-Magang-GITS

Pada soal 1:
Saya diminta untuk membuat program yang menghasilkan deret angka berdasarkan rumus A000124 dari OEIS. Deret ini dimulai dari angka 1, lalu setiap elemen berikutnya ditambahkan dengan nilai indeks + 1, sehingga membentuk pola seperti 1, 2, 4, 7, 11, 16, .... Untuk menghitungnya, saya pakai perulangan for dan menambahkan nilai ke variabel current di setiap iterasi. Output akhirnya ditampilkan dalam format string yang dipisahkan tanda strip (-). Soal ini membantu saya memahami bagaimana membuat pola angka menggunakan logika sederhana.

Pada soal 2:
Soal ini berkaitan dengan sistem Dense Ranking, yaitu sistem peringkat di mana skor yang sama akan mendapat peringkat yang sama, dan skor selanjutnya akan menyesuaikan jumlah skor unik di atasnya. Saya menggunakan map untuk menyimpan skor unik, lalu saya urutkan dari yang tertinggi ke terendah. Setelah itu, saya cek skor milik GITS dan tentukan peringkatnya dengan cara membandingkan terhadap skor-skor tersebut. Soal ini menurut saya lumayan menantang karena harus mengatur logika pengurutan dan pencocokan nilai secara efisien.

Pada soal 3 :
Saya membuat program yang bisa mengecek apakah string berisi bracket (kurung buka dan tutup) tersusun secara seimbang. Yang dicek hanya tanda (), {}, dan []. Untuk menyelesaikannya, saya pakai stack (dalam bentuk slice) dan mapping pasangan kurung. Setiap kali ketemu tanda buka, saya masukkan ke stack, dan saat ketemu tanda tutup, saya cek apakah pasangannya sesuai. Kalau tidak sesuai atau stack masih sisa di akhir, berarti tidak seimbang. Soal ini bagus banget buat memahami konsep stack dan cocok buat latihan logika pemrograman dasar.

Kompleksitas waktu dari algoritma ini adalah O(n), di mana n adalah panjang string yang diberikan. Hal ini karena setiap karakter hanya diproses satu kali, baik itu untuk dimasukkan ke stack atau untuk dicocokkan dan dikeluarkan dari stack. Proses pengecekan, penambahan, dan penghapusan dari stack semuanya dilakukan dalam waktu konstan O(1). Jadi, secara keseluruhan, algoritma ini efisien dan optimal untuk menyelesaikan masalah balanced bracket.
