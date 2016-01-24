package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
	"github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"
)

/*LearningResource LearningResource learning resource

swagger:model LearningResource
*/
type LearningResource struct {

	/* ID id

	Required: true
	Read Only: true
	*/
	ID int64 `json:"id,omitempty"`

	/* Languages languages
	 */
	Languages []string `json:"languages,omitempty"`

	/* Name name

	Required: true
	*/
	Name string `json:"name,omitempty"`

	/* Screenshots screenshots
	 */
	Screenshots []string `json:"screenshots,omitempty"`

	/* Type type

	Required: true
	*/
	Type string `json:"type,omitempty"`

	/* URL url

	Required: true
	*/
	URL string `json:"url,omitempty"`
}

// Validate validates this learning resource
func (m *LearningResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLanguages(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateScreenshots(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LearningResource) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *LearningResource) validateLanguages(formats strfmt.Registry) error {

	if swag.IsZero(m.Languages) { // not required
		return nil
	}

	for i := 0; i < len(m.Languages); i++ {

		if err := validate.RequiredString("languages"+"."+strconv.Itoa(i), "body", string(m.Languages[i])); err != nil {
			return err
		}

	}

	return nil
}

func (m *LearningResource) validateName(formats strfmt.Registry) error {

	if err := validate.RequiredString("name", "body", string(m.Name)); err != nil {
		return err
	}

	return nil
}

func (m *LearningResource) validateScreenshots(formats strfmt.Registry) error {

	if swag.IsZero(m.Screenshots) { // not required
		return nil
	}

	for i := 0; i < len(m.Screenshots); i++ {

		if err := validate.RequiredString("screenshots"+"."+strconv.Itoa(i), "body", string(m.Screenshots[i])); err != nil {
			return err
		}

	}

	return nil
}

func (m *LearningResource) validateType(formats strfmt.Registry) error {

	if err := validate.RequiredString("type", "body", string(m.Type)); err != nil {
		return err
	}

	return nil
}

func (m *LearningResource) validateURL(formats strfmt.Registry) error {

	if err := validate.RequiredString("url", "body", string(m.URL)); err != nil {
		return err
	}

	return nil
}
