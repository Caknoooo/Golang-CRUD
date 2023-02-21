package pasiencontroller

import (
	"html/template"
	"net/http"

	"github.com/jeypc/go-crud/entities"
	"github.com/jeypc/go-crud/models"
	"github.com/jeypc/go-crud/libraries"
)

var pasienModel = models.NewPasienModel()
var validation = libraries.NewValidation()

func Index(response http.ResponseWriter, request *http.Request) {
	pasien, _ := pasienModel.FindAll()

	data := map[string]interface{}{
		"pasien": pasien,
	}

	temp, err := template.ParseFiles("views/pasien/index.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(response, data)
}

func Add(response http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		temp, err := template.ParseFiles("views/pasien/add.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(response, nil)
	}	else if request.Method == http.MethodPost{
		request.ParseForm()
		
		var pasien entities.Pasien

		pasien.NamaLengkap = request.Form.Get("nama_lengkap")
		pasien.NIK = request.Form.Get("nik")
		pasien.JenisKelamin = request.Form.Get("jenis_kelamin")
		pasien.TempatLahir = request.Form.Get("tempat_lahir")
		pasien.TanggalLahir = request.Form.Get("tanggal_lahir")
		pasien.Alamat = request.Form.Get("alamat")
		pasien.NoHp = request.Form.Get("no_hp")

		var data = make(map[string]interface{})

		vErrors := validation.Struct(pasien)

		if vErrors != nil{
			data["validation"] = vErrors
			data["pasien"] = pasien
		} else{
			data["pesan"] = "Data pasien berhasil disimpan"
			pasienModel.Create(pasien)
		}
		
		temp, _ := template.ParseFiles("views/pasien/add.html")
		temp.Execute(response, data)
	}
}

func Edit(response http.ResponseWriter, request *http.Request) {

}

func Delete(response http.ResponseWriter, request *http.Request) {

}
