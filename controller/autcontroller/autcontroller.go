package autcontroller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/oriegenbi27/go-jwt-mux/models"
	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request) {

}

func Register(w http.ResponseWriter, r *http.Request) {

	// get data inpur json

	var userInput models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		log.Fatal("Gagagl Mendecode json")
	}

	defer r.Body.Close()

	// has password

	hasPassword, _ := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	userInput.Password = string(hasPassword)

	// insert ke db
	if err := models.DB.Create(&userInput).Error; err != nil {
		log.Fatal("gagal menyimpan data")
	}

	response, _ := json.Marshal(map[string]string{"message": "success"})
	w.Header().Add("Content-tyoe", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func Logout(w http.ResponseWriter, r *http.Request) {

}
