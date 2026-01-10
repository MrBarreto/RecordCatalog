package models

import "errors"

type RecordModel struct {
	ID          string
	Title       string
	Artist      string
	ReleaseYear uint16
	Status      string
}

func (r *RecordModel) Validate() error {
	if r.Title == "" || r.Artist == "" {
		return errors.New("The album must have a title and a artist")
	}
	return nil
}
