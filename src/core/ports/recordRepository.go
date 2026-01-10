package ports

import {
	"github.com/MrBarreto/RecordCatalog/src/core/models/models.go"
}

type RecordRepository interface {
	CreateRecord(record models.RecordModel) error
	GetAlbumsByArtist(artist string) ([]models.RecordModel, error)
	GetAvailableArtists() ([]string, error)
	GetAlbumByID(ID string) (models.RecordModel, error)
	GetAlbums() ([]models.RecordModel, error)
}
