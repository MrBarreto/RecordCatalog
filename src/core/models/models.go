package models

import "errors"

type RecordModel struct{
	ID          string
	Title       string
	Artist      string
	ReleaseYear uint16
	Status      string
}

func (r *RecordModel) Validate {
	if len(r.Title) == 0 || len(r.Artist) {
		return errors.new("The album must have a title and a artist")
	}
	return nil
}
