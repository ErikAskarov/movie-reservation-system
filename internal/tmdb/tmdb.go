package tmdb

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Movie struct {
	Adult        bool
	BackdropPath string `json:"backdrop_path"`
	// BelongsToCollection bool   `json:"belongs_to_collection"`
	BelongsToCollection CollectionShort `json:"belongs_to_collection"`
	Budget              uint32
	Genres              []struct {
		ID   int
		Name string
	}
	Homepage            string
	ID                  int
	ImdbID              string `json:"imdb_id"`
	OriginalLanguage    string `json:"original_language"`
	OriginalTitle       string `json:"original_title"`
	Overview            string
	Popularity          float32
	PosterPath          string `json:"poster_path"`
	ProductionCompanies []struct {
		ID        int
		Name      string
		LogoPath  string `json:"logo_path"`
		Iso3166_1 string `json:"origin_country"`
	} `json:"production_companies"`
	ProductionCountries []struct {
		Iso3166_1 string `json:"iso_3166_1"`
		Name      string
	} `json:"production_countries"`
	ReleaseDate     string `json:"release_date"`
	Revenue         uint32
	Runtime         uint32
	SpokenLanguages []struct {
		Iso639_1 string `json:"iso_639_1"`
		Name     string
	} `json:"spoken_languages"`
	Status  string
	Tagline string
	Title   string
	Video   bool
}

// MovieShort struct
type MovieShort struct {
	Adult         bool    `json:"adult"`
	BackdropPath  string  `json:"backdrop_path"`
	ID            int     `json:"id"`
	OriginalTitle string  `json:"original_title"`
	GenreIDs      []int32 `json:"genre_ids"`
	Popularity    float32 `json:"popularity"`
	PosterPath    string  `json:"poster_path"`
	ReleaseDate   string  `json:"release_date"`
	Title         string  `json:"title"`
	Overview      string  `json:"overview"`
	Video         bool    `json:"video"`
	VoteAverage   float32 `json:"vote_average"`
	VoteCount     uint32  `json:"vote_count"`
}

type CollectionShort struct {
	ID           int
	Name         string
	PosterPath   string `json:"poster_path"`
	BackdropPath string `json:"backdrop_path"`
}

type PopularMoviesResponse struct {
	Page    int          `json:"page"`
	Results []MovieShort `json:"results"` // Используем MovieShort для списка
	Total   int          `json:"total_results"`
}

func GetPopularMovies() ([]MovieShort, error) {
	url := "https://api.themoviedb.org/3/movie/popular?language=ru-RU&page=1"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiJjZmQwMWI1NTZmZTVhODRkZDdmZTNhZDdlOTE4M2VhOSIsIm5iZiI6MTc0MjYyMzEyOS43NjE5OTk4LCJzdWIiOiI2N2RlNTE5OTVmOTc3ODM2YWI3YTlkZTgiLCJzY29wZXMiOlsiYXBpX3JlYWQiXSwidmVyc2lvbiI6MX0.M9DxaSZVj8EmBncopBcHy4eip7WE3_tphRdSHjh46c0")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("Ошибка запроса:", err) // Обрабатываем ошибку сразу
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	//fmt.Println(string(body))

	var response PopularMoviesResponse
	if err := json.Unmarshal(body, &response); err != nil { // Десериализуем JSON в структуру
		return nil, err
	}

	return response.Results, nil
}

func GetMovieInfo(movie_id int) (Movie, error) {
	url := fmt.Sprintf("https://api.themoviedb.org/3/movie/%d?language=ru-RU", movie_id)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiJjZmQwMWI1NTZmZTVhODRkZDdmZTNhZDdlOTE4M2VhOSIsIm5iZiI6MTc0MjYyMzEyOS43NjE5OTk4LCJzdWIiOiI2N2RlNTE5OTVmOTc3ODM2YWI3YTlkZTgiLCJzY29wZXMiOlsiYXBpX3JlYWQiXSwidmVyc2lvbiI6MX0.M9DxaSZVj8EmBncopBcHy4eip7WE3_tphRdSHjh46c0")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("Ошибка запроса:", err) // Обрабатываем ошибку сразу
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	//fmt.Println(string(body))

	var movie Movie
	if err := json.Unmarshal(body, &movie); err != nil { // Десериализуем JSON в структуру
		return Movie{}, err
	}

	return movie, nil
}
