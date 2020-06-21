package models

import (
	b64 "encoding/base64"
	"encoding/json"
	"github.com/gobuffalo/buffalo/binding"
	"github.com/gobuffalo/pop/v5"
	"github.com/gobuffalo/validate/v3"
	"github.com/gobuffalo/validate/v3/validators"
	"github.com/gofrs/uuid"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// Image is used by pop to map your images database table to your go code.
type Image struct {
	ID           uuid.UUID     `json:"id" db:"id"`
	CreatedAt    time.Time     `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time     `json:"updated_at" db:"updated_at"`
	Image        []byte        `json:"image" db:"image"`
	ImageType    string        `json:"image_type" db:"image_type"`
	EncodedImage string        `db:"-"`
	Avatar       *binding.File `db:"-"`
}

// String is not required by pop and may be deleted
func (i Image) String() string {
	ji, _ := json.Marshal(i)
	return string(ji)
}

// Images is not required by pop and may be deleted
type Images []Image

// String is not required by pop and may be deleted
func (i Images) String() string {
	ji, _ := json.Marshal(i)
	return string(ji)
}

func (i *Image) BeforeValidate(tx *pop.Connection) error {
	var b []byte
	b, err := ioutil.ReadAll(i.Avatar.File)
	if err != nil {
		return err
	}
	ttype := http.DetectContentType(b)
	i.ImageType = ttype
	i.Image = b
	return nil
}

func (i *Image) AfterFind(tx *pop.Connection) error {
	i.EncodedImage = b64.StdEncoding.EncodeToString(i.Image)
	return nil
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (i *Image) Validate(tx *pop.Connection) (*validate.Errors, error) {
	var err error
	return validate.Validate(
		&validators.StringIsPresent{Field: i.ImageType, Name: "ImageType"},
		&validators.FuncValidator{
			Field:   i.ImageType,
			Name:    "ImageType",
			Message: "%s must be an Image",
			Fn: func() bool {
				if i.ImageType == "text/plain; charset=utf-8" && strings.HasSuffix(i.Avatar.FileHeader.Filename, ".svg") {
					i.ImageType = "image/svg+xml"
				}
				_, found := find([]string{"image/svg+xml", "image/png", "image/jpeg", "image/gif"}, i.ImageType)
				return found
			},
		},
	), err
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (i *Image) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (i *Image) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
