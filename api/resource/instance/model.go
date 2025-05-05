package instance

import (
	"strconv"

	"github.com/google/uuid"
)

type DTO struct {
	ID      string `json:"id"`
	Type    string `json:"type"`
	CPU     string `json:"cpu"`
	RAM     string `json:"memory"`
	Storage string `json:"storage"`
	User    string `json:"user"`
}

type Form struct {
	Type    string `json:"type" form:"required,max=255"`
	CPU     string `json:"cpu" form:"required,numeric,max=2"`
	RAM     string `json:"memory" form:"required,numeric,max=10"`
	Storage string `json:"storage" form:"required,numeric,max=2048"`
	User    string `json:"user" form:"required,max=255"`
}

type Instance struct {
	ID      uuid.UUID
	Type    string
	CPU     int
	RAM     int
	Storage int
	User    string
}

type Instances []*Instance

func (i *Instance) ToDto() *DTO {
	return &DTO{
		ID:      i.ID.String(),
		Type:    i.Type,
		CPU:     strconv.Itoa(i.CPU),
		RAM:     strconv.Itoa(i.RAM),
		Storage: strconv.Itoa(i.Storage),
		User:    i.User,
	}
}

func (is Instances) ToDto() []*DTO {
	dtos := make([]*DTO, len(is))
	for i, v := range is {
		dtos[i] = v.ToDto()
	}

	return dtos
}

func (f *Form) ToModel() *Instance {
	instanceCPU, err := strconv.Atoi(f.CPU)
	if err != nil {
		instanceCPU = 0
	}

	instanceRAM, err := strconv.Atoi(f.RAM)
	if err != nil {
		instanceRAM = 0
	}

	instanceStorage, err := strconv.Atoi(f.Storage)
	if err != nil {
		instanceStorage = 256
	}

	return &Instance{
		Type:    f.Type,
		CPU:     instanceCPU,
		RAM:     instanceRAM,
		Storage: instanceStorage,
		User:    f.User,
	}

}
