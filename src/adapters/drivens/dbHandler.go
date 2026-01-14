package drivens

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/MrBarreto/RecordCatalog/src/core/models"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func CreatingConnection() (*pgxpool.Pool, error) {
	err := godotenv.Load("../go.env")
	if err != nil {
		return nil, errors.New("Error loading .env file")
	}
	dbURL := os.Getenv("DATABASE_URL")

	dbpool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		errorMessage := fmt.Sprintf("Unable to create connection pool: %v\n", err)
		return nil, errors.New(errorMessage)
	}

	return dbpool, nil
}

type dbHandler struct {
	connection *pgxpool.Pool
}

func (d *dbHandler) CreateRecord(record models.RecordModel) (int, error) {
	var novoID int
	sql := "INSERT INTO records (title, artist, releaseyear, status) VALUES ($1, $2, $3, $4)"
	err := d.connection.QueryRow(context.Background(), sql, record.Title, record.Artist, record.ReleaseYear, record.Status).Scan(&novoID)
	if err != nil {
		errString := fmt.Sprintf("addAlbum: %v", err)
		return 0, errors.New(errString)
	}
	return novoID, nil
}

func (d *dbHandler) GetAlbumsByArtist(artist string) ([]models.RecordModel, error) {
	sql := "SELECT title FROM records WHERE artist = $1"
	rows, err := d.connection.Query(context.Background(), sql, artist)

	if err != nil {
		return nil, err
	}

	var Recordlist []models.RecordModel

	defer rows.Close()

	for rows.Next() {
		var Record models.RecordModel
		if err := rows.Scan(&Record.ID, &Record.Title, &Record.Artist, &Record.ReleaseYear, &Record.Status); err != nil {
			return nil, fmt.Errorf("albumsByArtist %q: %v", artist, err)
		}

		Recordlist = append(Recordlist, Record)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", artist, err)
	}

	return Recordlist, nil
}

func (d *dbHandler) GetAvailableArtists() ([]string, error) {
	sql := "SELECT artist FROM records"
	rows, err := d.connection.Query(context.Background(), sql)

	if err != nil {
		return nil, err
	}

	var Artistlist []string

	defer rows.Close()

	for rows.Next() {
		var Artist string
		if err := rows.Scan(&Artist); err != nil {
			return nil, fmt.Errorf("GetAvailableArtists: %v", err)
		}

		Artistlist = append(Artistlist, Artist)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("GetAvailableArtists %q: %v", err)
	}

	return Artistlist, nil
}
