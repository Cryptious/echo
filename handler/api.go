package handler

import (
	"echo/server"
	"fmt"
	_ "mysql-master"
	"net/http"

	"github.com/labstack/echo"
)

type menu struct {
	Id_menu     string
	Nama_menu   string
	Deskripsi   string
	Jenis       string
	Harga       string
	Url_gambar  string
	Total_order string
}

var data []menu

func BacaData(c echo.Context) error {
	menu_makanan()

	return c.JSON(http.StatusOK, data)
}

func BacaPopuler(c echo.Context) error {
	menu_populer()

	return c.JSON(http.StatusOK, data)
}

func TambahData(c echo.Context) error {
	db, err := server.Koneksi()
	defer db.Close()

	var nama = c.FormValue("Nama_menu")
	var deskripsi = c.FormValue("Deskripsi")
	var harga = c.FormValue("Harga")
	var jenis = c.FormValue("Jenis")
	var url_gambar = c.FormValue("Url_gambar")

	_, err = db.Exec("Insert into tbl_menu values (?,?,?,?,?,?)", nil, nama, deskripsi, url_gambar, jenis, harga)

	if err != nil {
		fmt.Println("Menu Gagal Ditambahkan")
		return c.JSON(http.StatusOK, "Gagal Menambahkan Menu")
	} else {
		fmt.Println("Menu Berhasil Ditambahkan")
		return c.JSON(http.StatusOK, "Berhasil Menambahkan Menu")
	}
}

func InputOrder(c echo.Context) error {
	db, err := server.Koneksi()
	defer db.Close()

	var id = c.FormValue("id")
	var nama_pemesan = c.FormValue("nama_pemesan")
	var nomor_telepon = c.FormValue("nomor_telepon")
	var jumlah = c.FormValue("jumlah")
	var alamat = c.FormValue("alamat")

	_, err = db.Exec("Insert into tbl_order values (?,?,?,?,?,?)", nil, id, nama_pemesan, nomor_telepon, alamat, jumlah)

	if err != nil {
		fmt.Println("Menu Gagal Dipesan")
		return c.HTML(http.StatusOK, "<strong>Gagal PESAN</strong>")
	} else {
		fmt.Println("Pesanan Berhasil Dibuat")
		return c.HTML(http.StatusOK, "<script>alert('Berhasil Melakukan Pemesanan. Terima Kasih'); window.location='http://localhost:1323';</script>")
	}
	return c.Redirect(http.StatusSeeOther, "/")
}

func UbahData(c echo.Context) error {
	db, err := server.Koneksi()
	defer db.Close()

	var id = c.FormValue("Id_menu")
	var nama = c.FormValue("Nama_menu")
	var deskripsi = c.FormValue("Deskripsi")
	var harga = c.FormValue("Harga")
	var jenis = c.FormValue("Jenis")
	var url_gambar = c.FormValue("Url_gambar")

	_, err = db.Exec("Update tbl_menu set nama_menu = ?, deskripsi = ?, url_gambar = ?, jenis = ?, harga = ? where id_menu = ?", nama, deskripsi, url_gambar, jenis, harga, id)

	if err != nil {
		fmt.Println("Menu Gagal Diubah")
		return c.JSON(http.StatusOK, "Gagal Mengubah Menu")
	} else {
		fmt.Println("Menu Berhasil Diubah")
		return c.JSON(http.StatusOK, "Berhasil Mengubah Menu")
	}
}

func HapusData(c echo.Context) error {
	db, err := server.Koneksi()
	defer db.Close()

	var id = c.FormValue("Id_menu")

	_, err = db.Exec("Delete from tbl_menu where id_menu = ?", id)

	if err != nil {
		fmt.Println("Menu Gagal Dihapus")
		return c.JSON(http.StatusOK, "Gagal Menghapus Menu")
	} else {
		fmt.Println("Menu Berhasil DIhapu")
		return c.JSON(http.StatusOK, "Berhasil Menghapus Menu")
	}
}

func menu_makanan() {
	data = nil
	db, err := server.Koneksi()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()
	rows, err := db.Query("select * from tbl_menu")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer rows.Close()

	for rows.Next() {
		var each = menu{}
		var err = rows.Scan(&each.Id_menu, &each.Nama_menu, &each.Deskripsi, &each.Url_gambar, &each.Jenis, &each.Harga)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		data = append(data, each)
		fmt.Println(data)
	}
	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}
}

func menu_populer() {
	data = nil
	db, err := server.Koneksi()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()
	rows, err := db.Query("select * from vw_totalorder ORDER BY total_order DESC LIMIT 8")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer rows.Close()

	for rows.Next() {
		var each = menu{}
		var err = rows.Scan(&each.Id_menu, &each.Nama_menu, &each.Deskripsi, &each.Url_gambar, &each.Jenis, &each.Harga, &each.Total_order)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		data = append(data, each)
		fmt.Println(data)
	}
	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}
}
