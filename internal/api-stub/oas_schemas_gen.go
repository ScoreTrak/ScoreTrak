// Code generated by ogen, DO NOT EDIT.

package api_stub

import (
	"time"
)

type ApiToken struct {
	Token string
}

// GetToken returns the value of Token.
func (s *ApiToken) GetToken() string {
	return s.Token
}

// SetToken sets the value of Token.
func (s *ApiToken) SetToken(val string) {
	s.Token = val
}

// Ref: #/components/schemas/Competition
type Competition struct {
	ID          OptString   `json:"id"`
	Name        OptString   `json:"name"`
	DisplayName OptString   `json:"display_name"`
	StartAt     OptDateTime `json:"start_at"`
	EndAt       OptDateTime `json:"end_at"`
	CreatedAt   OptDateTime `json:"created_at"`
	UpdatedAt   OptDateTime `json:"updated_at"`
}

// GetID returns the value of ID.
func (s *Competition) GetID() OptString {
	return s.ID
}

// GetName returns the value of Name.
func (s *Competition) GetName() OptString {
	return s.Name
}

// GetDisplayName returns the value of DisplayName.
func (s *Competition) GetDisplayName() OptString {
	return s.DisplayName
}

// GetStartAt returns the value of StartAt.
func (s *Competition) GetStartAt() OptDateTime {
	return s.StartAt
}

// GetEndAt returns the value of EndAt.
func (s *Competition) GetEndAt() OptDateTime {
	return s.EndAt
}

// GetCreatedAt returns the value of CreatedAt.
func (s *Competition) GetCreatedAt() OptDateTime {
	return s.CreatedAt
}

// GetUpdatedAt returns the value of UpdatedAt.
func (s *Competition) GetUpdatedAt() OptDateTime {
	return s.UpdatedAt
}

// SetID sets the value of ID.
func (s *Competition) SetID(val OptString) {
	s.ID = val
}

// SetName sets the value of Name.
func (s *Competition) SetName(val OptString) {
	s.Name = val
}

// SetDisplayName sets the value of DisplayName.
func (s *Competition) SetDisplayName(val OptString) {
	s.DisplayName = val
}

// SetStartAt sets the value of StartAt.
func (s *Competition) SetStartAt(val OptDateTime) {
	s.StartAt = val
}

// SetEndAt sets the value of EndAt.
func (s *Competition) SetEndAt(val OptDateTime) {
	s.EndAt = val
}

// SetCreatedAt sets the value of CreatedAt.
func (s *Competition) SetCreatedAt(val OptDateTime) {
	s.CreatedAt = val
}

// SetUpdatedAt sets the value of UpdatedAt.
func (s *Competition) SetUpdatedAt(val OptDateTime) {
	s.UpdatedAt = val
}

// Ref: #/components/schemas/CreateCompetitionRequest
type CreateCompetitionRequest struct {
	Name             string  `json:"name"`
	DisplayName      string  `json:"display_name"`
	ViewableToPublic OptBool `json:"viewable_to_public"`
}

// GetName returns the value of Name.
func (s *CreateCompetitionRequest) GetName() string {
	return s.Name
}

// GetDisplayName returns the value of DisplayName.
func (s *CreateCompetitionRequest) GetDisplayName() string {
	return s.DisplayName
}

// GetViewableToPublic returns the value of ViewableToPublic.
func (s *CreateCompetitionRequest) GetViewableToPublic() OptBool {
	return s.ViewableToPublic
}

// SetName sets the value of Name.
func (s *CreateCompetitionRequest) SetName(val string) {
	s.Name = val
}

// SetDisplayName sets the value of DisplayName.
func (s *CreateCompetitionRequest) SetDisplayName(val string) {
	s.DisplayName = val
}

// SetViewableToPublic sets the value of ViewableToPublic.
func (s *CreateCompetitionRequest) SetViewableToPublic(val OptBool) {
	s.ViewableToPublic = val
}

// NewOptBool returns new OptBool with value set to v.
func NewOptBool(v bool) OptBool {
	return OptBool{
		Value: v,
		Set:   true,
	}
}

// OptBool is optional bool.
type OptBool struct {
	Value bool
	Set   bool
}

// IsSet returns true if OptBool was set.
func (o OptBool) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptBool) Reset() {
	var v bool
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptBool) SetTo(v bool) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptBool) Get() (v bool, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptBool) Or(d bool) bool {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptDateTime returns new OptDateTime with value set to v.
func NewOptDateTime(v time.Time) OptDateTime {
	return OptDateTime{
		Value: v,
		Set:   true,
	}
}

// OptDateTime is optional time.Time.
type OptDateTime struct {
	Value time.Time
	Set   bool
}

// IsSet returns true if OptDateTime was set.
func (o OptDateTime) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptDateTime) Reset() {
	var v time.Time
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptDateTime) SetTo(v time.Time) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptDateTime) Get() (v time.Time, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptDateTime) Or(d time.Time) time.Time {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptString returns new OptString with value set to v.
func NewOptString(v string) OptString {
	return OptString{
		Value: v,
		Set:   true,
	}
}

// OptString is optional string.
type OptString struct {
	Value string
	Set   bool
}

// IsSet returns true if OptString was set.
func (o OptString) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptString) Reset() {
	var v string
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptString) SetTo(v string) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptString) Get() (v string, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptString) Or(d string) string {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// Ref: #/components/schemas/UpdateCompetitionRequest
type UpdateCompetitionRequest struct {
	Name        OptString `json:"name"`
	DisplayName OptString `json:"display_name"`
}

// GetName returns the value of Name.
func (s *UpdateCompetitionRequest) GetName() OptString {
	return s.Name
}

// GetDisplayName returns the value of DisplayName.
func (s *UpdateCompetitionRequest) GetDisplayName() OptString {
	return s.DisplayName
}

// SetName sets the value of Name.
func (s *UpdateCompetitionRequest) SetName(val OptString) {
	s.Name = val
}

// SetDisplayName sets the value of DisplayName.
func (s *UpdateCompetitionRequest) SetDisplayName(val OptString) {
	s.DisplayName = val
}
