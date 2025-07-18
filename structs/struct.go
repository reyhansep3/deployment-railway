package structs

type Bioskop struct {
	ID     int     `json:"id"`
	Nama   string  `json:"nama"`
	Lokasi string  `json:"Lokasi"`
	Rating float64 `json:"Rating"`
}
