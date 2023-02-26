// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type File interface {
	IsFile()
}

type BasePaginationFilter struct {
	Count  *int `json:"count"`
	Offset *int `json:"offset"`
}

type DeploymentPlanInput struct {
	JSONSerialize *string `json:"jsonSerialize"`
	ExternalLink  *string `json:"externalLink"`
}

type DocumentReadable struct {
	ID          string  `json:"id"`
	Name        *string `json:"name"`
	Extension   *string `json:"extension"`
	ContentType *string `json:"contentType"`
	Blob        string  `json:"blob"`
	Size        int     `json:"size"`
	Preview     *string `json:"preview"`
}

func (DocumentReadable) IsFile() {}

type FileLink struct {
	ID          string  `json:"id"`
	Name        *string `json:"name"`
	Extension   *string `json:"extension"`
	ContentType *string `json:"contentType"`
	Link        string  `json:"link"`
}

func (FileLink) IsFile() {}

type Infrastructure struct {
	Project  *Project       `json:"project"`
	Platform *string        `json:"platform"`
	Status   *ServiceStatus `json:"status"`
	UpStatus *bool          `json:"upStatus"`
	UpTime   *int           `json:"upTime"`
}

type Project struct {
	ID           string        `json:"id"`
	Key          string        `json:"key"`
	Name         string        `json:"name"`
	Description  *string       `json:"description"`
	Technologies []*Technology `json:"technologies"`
}

type ProjectFilter struct {
	Name         *string   `json:"name"`
	Technologies []*string `json:"technologies"`
	Count        *int      `json:"count"`
	Offset       *int      `json:"offset"`
}

type ProjectInput struct {
	ID           *string            `json:"id"`
	Key          *string            `json:"key"`
	Name         *string            `json:"name"`
	Description  *string            `json:"description"`
	Technologies []*TechnologyInput `json:"technologies"`
}

type RequestServiceInput struct {
	Name *string `json:"name"`
}

type Technology struct {
	ID          string  `json:"id"`
	Key         string  `json:"key"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

type TechnologyInput struct {
	ID          *string `json:"id"`
	Key         *string `json:"key"`
	Name        *string `json:"name"`
	Description *string `json:"description"`
}

type UploadDeploymentPlanInput struct {
	Project *ProjectFilter       `json:"project"`
	Plan    *DeploymentPlanInput `json:"plan"`
}

type ServiceStatus string

const (
	ServiceStatusCreated ServiceStatus = "CREATED"
	ServiceStatusExhaust ServiceStatus = "EXHAUST"
	ServiceStatusDown    ServiceStatus = "DOWN"
)

var AllServiceStatus = []ServiceStatus{
	ServiceStatusCreated,
	ServiceStatusExhaust,
	ServiceStatusDown,
}

func (e ServiceStatus) IsValid() bool {
	switch e {
	case ServiceStatusCreated, ServiceStatusExhaust, ServiceStatusDown:
		return true
	}
	return false
}

func (e ServiceStatus) String() string {
	return string(e)
}

func (e *ServiceStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ServiceStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ServiceStatus", str)
	}
	return nil
}

func (e ServiceStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
