package builder

import "fmt"

// builder interface
type iBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() house
}

func getBuilder(builderType string) iBuilder {
	switch builderType {
	case "normal":
		return &normalBuilder{}
	case "high":
		return &highBuilder{}
	default:
		return nil
	}
}

// concrete builder - normal
type normalBuilder struct {
	windowsType string
	doorType    string
	floor       int
}

func newNormalBuilder() *normalBuilder {
	return &normalBuilder{}
}

func (n *normalBuilder) setWindowType() {
	n.windowsType = "Wooden Window"
}

func (n *normalBuilder) setDoorType() {
	n.doorType = "Wooden Door"
}

func (n *normalBuilder) setNumFloor() {
	n.floor = 1
}

func (n *normalBuilder) getHouse() house {
	return house{
		windowsType: n.windowsType,
		doorType:    n.doorType,
		floor:       n.floor,
	}
}

// concerate builder - high
type highBuilder struct {
	windowsType string
	doorType    string
	floor       int
}

func newHighBuilder() *highBuilder {
	return &highBuilder{}
}

func (h *highBuilder) setWindowType() {
	h.windowsType = "Glass Window"
}

func (h *highBuilder) setDoorType() {
	h.doorType = "Glass Door"
}

func (h *highBuilder) setNumFloor() {
	h.floor = 2
}

func (h *highBuilder) getHouse() house {
	return house{
		windowsType: h.windowsType,
		doorType:    h.doorType,
		floor:       h.floor,
	}
}

// product
type house struct {
	windowsType string
	doorType    string
	floor       int
}

func (h *house) show() string {
	return fmt.Sprintf("House: %s, %s, %d floor\n", h.windowsType, h.doorType, h.floor)
}

// director
type director struct {
	builder iBuilder
}

func newDirector(builder iBuilder) *director {
	return &director{
		builder: builder,
	}
}

func (d *director) construct() house {
	d.builder.setWindowType()
	d.builder.setDoorType()
	d.builder.setNumFloor()
	return d.builder.getHouse()
}
