package ports

type Record struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Artist      string `json:"artist"`
	ReleaseYear uint16 `json:"releaseyear"`
	Status      string `json:"status"`
}

type RecordService interface {
	GetAlbumsByArtist(artist string) ([]Record, error)
	GetAvailableArtists() ([]string, error)
	GetAlbumByID(ID string) (Record, error)
	GetAlbums() ([]Record, error)
}
