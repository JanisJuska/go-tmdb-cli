package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"text/tabwriter"
)

type MovieResponse struct {
	Dates        *Dates    `json:"dates,omitempty"` // pointer = optional
	Page         int       `json:"page"`
	Results      []Results `json:"results"`
	TotalPages   int       `json:"total_pages"`
	TotalResults int       `json:"total_results"`
}

type Results struct {
	Adult            bool    `json:"adult"`
	BackdropPath     string  `json:"backdrop_path"`
	GenreIds         []int   `json:"genre_ids"`
	ID               int     `json:"id"`
	OriginalLanguage string  `json:"original_language"`
	OriginalTitle    string  `json:"original_title"`
	Overview         string  `json:"overview"`
	Popularity       float64 `json:"popularity"`
	PosterPath       string  `json:"poster_path"`
	ReleaseDate      string  `json:"release_date"`
	Title            string  `json:"title"`
	Video            bool    `json:"video"`
	VoteAverage      float64 `json:"vote_average"`
	VoteCount        int     `json:"vote_count"`
}

type Dates struct {
	Maximum string `json:"maximum"`
	Minimum string `json:"minimum"`
}

func main() {
	typeFlag := flag.String("type", "", "movie type: playing, popular, top, upcoming")
	fullFlag := flag.Bool("full", false, "Show full (not truncated) movie titles")

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

	flag.Parse()

	var reqType string
	var reqTString string

	switch *typeFlag {
	case "playing":
		reqType = "now_playing"
		reqTString = "Now Playing"
	case "popular":
		reqType = "popular"
		reqTString = "Popular"
	case "top":
		reqType = "top_rated"
		reqTString = "Top Rated"
	case "upcoming":
		reqType = "upcoming"
		reqTString = "Upcoming"
	default:
		fmt.Println("Missing required --type flag")
		flag.Usage()
		return
	}

	fmt.Printf("Showing %v movies from TMDB:\n\n", reqTString)
	fmt.Fprintln(w, "#\tTitle\tRelease Date\tLang\tPopularity Score\tVote Avg")
	fmt.Fprintln(w, "\t\t\t\t\t")

	url := fmt.Sprintf("https://api.themoviedb.org/3/movie/%v?language=en-US&page=1", reqType)

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	request.Header.Add("cept", "application/json")
	request.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiI2MzliMjI5ZDE3NjRmMTM1YWVmNWNkNzZmOTEyMTk3YiIsIm5iZiI6MTc3NjcwMDUzMi43MDQsInN1YiI6IjY5ZTY0Yzc0ZTQyZGJhMzFmNzgwMmQ2NyIsInNjb3BlcyI6WyJhcGlfcmVhZCJdLCJ2ZXJzaW9uIjoxfQ.DWfykdh2ydM9jS8d5LIGYuef6NXW3xIOvt0di6NFBEw")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var unmarhsaledList MovieResponse

	err = json.Unmarshal(body, &unmarhsaledList)
	if err != nil {
		log.Fatal(err)
	}

	for k, v := range unmarhsaledList.Results {
		title := v.Title

		if !*fullFlag {
			title = formatTitle(title, 30)
		}

		fmt.Fprintf(w, "%d\t%s\t%s\t%s\t%.2f\t%.1f\n",
			k+1,
			title,
			v.ReleaseDate,
			v.OriginalLanguage,
			v.Popularity,
			v.VoteAverage,
		)
	}

	w.Flush()
}

func formatTitle(title string, max int) string {
	if len(title) > max {
		return title[:max-1] + "…"
	}
	return title
}
