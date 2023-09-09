package handler

import (
	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/mongo"
	"mongoPractice/modles"
	"mongoPractice/mongoDB/mflix"
	"mongoPractice/mongoDB/restaurant"
	"mongoPractice/utils"
	"net/http"
	"strconv"
)

func Health(w http.ResponseWriter, r *http.Request) {
	utils.RespondJSON(w, http.StatusOK, struct {
		Status string `json:"status"`
	}{Status: "server is running"})
}

// Restaurant GetRestaurant We are able to get the restaurant for a given name
func Restaurant(w http.ResponseWriter,r *http.Request){
	// get the name from the path
	restName := chi.URLParam(r, "name")
	// now we just need to give the name of the collection and vala..
	restaurantInfo,err := restaurant.DB{Name: "sample_restaurants"}.GetRestaurantDetail(restName)
	if err != nil{
		if err == mongo.ErrNoDocuments{
			utils.RespondError(w,http.StatusNotFound,err,"no restaurant with this name")
			return
		}
		utils.RespondError(w,http.StatusBadRequest,err,"bad request")
	}

	utils.RespondJSON(w, http.StatusOK, struct {
		Restaurant modles.Restaurant `json:"restaurant"`
	}{*restaurantInfo})

}

func AllRestaurant(w http.ResponseWriter,r *http.Request){
	cuisine := r.URL.Query().Get("cuisine")
	restaurantInfo,err := restaurant.DB{Name: "sample_restaurants"}.GetAllRestaurantDetail(cuisine)
	if err != nil{
		if err == mongo.ErrNoDocuments{
			utils.RespondError(w,http.StatusNotFound,err,"no restaurant with this name")
			return
		}
		utils.RespondError(w,http.StatusBadRequest,err,"bad request")
	}

	utils.RespondJSON(w, http.StatusOK, struct {
		Restaurant []modles.Restaurant `json:"restaurant"`
	}{*restaurantInfo})

}

func MovieWithIMDB(w http.ResponseWriter,r *http.Request){
	imdb := r.URL.Query().Get("imdb")
	votes := r.URL.Query().Get("votes")

	imdbValue, err := strconv.ParseFloat(imdb, 32)
	if err != nil{
		utils.RespondError(w,http.StatusUnprocessableEntity,err,"invalid imdb value")
		return
	}

	votesValue,err := strconv.Atoi(votes)
	if err != nil{
		utils.RespondError(w,http.StatusUnprocessableEntity,err,"invalid votes value")
		return
	}

	movies,err := mflix.DB{Name: "sample_mflix"}.GetMovieWithIMDB(float32(imdbValue),votesValue)
	if err != nil{
		if err == mongo.ErrNoDocuments{
			utils.RespondError(w,http.StatusNotFound,err,"no restaurant with this name")
			return
		}
		utils.RespondError(w,http.StatusBadRequest,err,"bad request")
	}

	utils.RespondJSON(w, http.StatusOK, struct {
		Movies []modles.Movie `json:"movies"`
	}{*movies})

}

func MovieWithCast(w http.ResponseWriter,r *http.Request)  {
	movies,err := mflix.DB{Name: "sample_mflix"}.GetMovieWithCast([]string{})
	if err != nil{
		if err == mongo.ErrNoDocuments{
			utils.RespondError(w,http.StatusNotFound,err,"no restaurant with this name")
			return
		}
		utils.RespondError(w,http.StatusBadRequest,err,"bad request")
	}

	utils.RespondJSON(w, http.StatusOK, struct {
		Movies []modles.Movie `json:"movies"`
	}{*movies})
}