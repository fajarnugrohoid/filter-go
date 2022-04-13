package main

import (
	"fmt"
	"sort"
	"time"
)

type Movie struct {
	Title string  // movie title
	Year  int     // movie release year
	Rate  float32 // movie rating
}

// byYear sorts all movies by release year
type byYear []*Movie

func (m byYear) Len() int           { return len(m) }
func (m byYear) Less(i, j int) bool { return m[i].Year < m[j].Year }
func (m byYear) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }

// byTitle sorts all movies by title
type byTitle []*Movie

func (m byTitle) Len() int           { return len(m) }
func (m byTitle) Less(i, j int) bool { return m[i].Title < m[j].Title }
func (m byTitle) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }

// byRate sorts all movies by rate
type byRate []*Movie

func (m byRate) Len() int           { return len(m) }
func (m byRate) Less(i, j int) bool { return m[i].Rate < m[j].Rate }
func (m byRate) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }

func main() {
	start := time.Now()

	movies := []*Movie{
		&Movie{"The 400 Blows", 1959, 8.1},
		&Movie{"La Haine", 1995, 8.1},
		&Movie{"The Godfather", 1972, 9.2},
		&Movie{"The Godfather: Part II", 1974, 9},
		&Movie{"Mafioso", 1962, 7.7}}

	displayMovies("Movies (unsorted)", movies)

	sort.Sort(byYear(movies))
	displayMovies("Movies sorted by year", movies)

	sort.Sort(byTitle(movies))
	displayMovies("Movies sorted by title", movies)

	sort.Sort(sort.Reverse(byRate(movies)))
	displayMovies("Movies sorted by rate", movies)

	timeElapsed := time.Since(start)
	fmt.Printf("The `for` loop took %s", timeElapsed)
}

func displayMovies(header string, movies []*Movie) {
	fmt.Println(header)
	for _, m := range movies {
		fmt.Printf("\t- %s (%d) R:%.1f\n", m.Title, m.Year, m.Rate)
	}
	fmt.Println()
}
