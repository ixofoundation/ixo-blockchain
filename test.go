package main

type DID struct {
	Id string
}

type IIDImpl interface {
}

type IID struct {
	DID
	Other string
}

type Entity interface {
	DID
	// IsEntity() bool
}

type Project struct {
}

// func (p *Project) IsEntity() bool { return true }

type Investment struct {
	Entity
}

// func (p *Investment) IsEntity() bool { return true }

func X(did DID) {

}

func Y() {

	Project{}.
		X(IID{Id: "hi", Other: "hi"})
}
