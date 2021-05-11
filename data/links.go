package data

import (
	"encoding/json"
	"io"
	"time"
)

type Link struct {
	OriginalLink  string `json:"originalLink"`
	ShortenedCode string `json:"shortenedCode"`
	NumOfVisits   uint64 `json:"numOfVisits"`
	CreatedAt     string `json:"-"`
	UpdatedAt     string `json:"-"`
}

type Links []*Link

// Return all links from the database
func GetLinks() Links {
	return linksList
}

func AddLink(l *Link) {
	// Update date and time of the transaction
	l.CreatedAt = time.Now().UTC().String()
	l.UpdatedAt = time.Now().UTC().String()

	linksList = append(linksList, l)
}

func (l *Link) FromJson(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(l)
}

func (l *Links) ToJson(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(l)
}

func (l *Links) GetLinkByCode(code string) *Link {
	for i := 0; i < len(linksList); i++ {
		if linksList[i].ShortenedCode == code {
			return linksList[i]
		}
	}
	return &Link{}
}

// Provisional database
var linksList = []*Link{
	{
		OriginalLink:  "https://google.com",
		ShortenedCode: "1273fjak",
		NumOfVisits:   3,
		CreatedAt:     "2021-05-11T21:26:46.234Z",
		UpdatedAt:     "2021-05-11T21:26:46.234Z",
	},
	{
		OriginalLink:  "https://amazon.com",
		ShortenedCode: "3876uui",
		NumOfVisits:   12,
		CreatedAt:     "2021-05-11T23:25:11.234Z",
		UpdatedAt:     "2021-05-11T23:25:11.234Z",
	},
}
