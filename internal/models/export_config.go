package models

type ExportConfig struct {
	Currency string         `json:"&curren"`
	Sections SectionsConfig `json:"sections"`
}

type SectionsConfig struct {
	Header      HeaderSection `json:"header section"`
	MenuSection MenuSection   `json:"menu section"`
}

type HeaderSection struct {
	Header1 Header1Block `json:"header-block-1"`
	Header2 Header2Block `json:"header-block-2"`
	Header3 Header3Block `json:"header-block-3"`
	Header4 Header4Block `json:"header-block-4"`
	Header5 Header5Block `json:"header-block-5"`
	Header6 Header6Block `json:"header-block-6"`
	Header7 Header7Block `json:"header-block-7"`
}

//companyName and creation date block
type Header1Block struct {
	Properties Header1BlockProperties `json:"properties"`
}

type Header1BlockProperties struct {
	CompanyName  DisplayProperties `json:"company name"`
	CreationDate DisplayProperties `json:"report creation date"`
}

type Header2Block struct {
	CompanyAddress DisplayProperties `json:"company address"`
}

type Header3Block struct {
	CompanyTaxNumber DisplayProperties `json:"company tax number"`
}

type Header4Block struct {
	TypeOfTaxation DisplayProperties `json:"type of taxation"`
}

type Header5Block struct {
	BranchName DisplayProperties `json:"branch name"`
}

type Header6Block struct {
	BranchTaxNumber DisplayProperties `json:"branch tax number"`
}

type Header7Block struct {
	BranchAddress DisplayProperties `json:"branch address"`
}

//menu blocks
type MenuSection struct {
	Block1  MenuBlock1        `json:"menu block 1"`
	Block2  DisplayProperties `json:"menu block 2"`
	Block3  MenuBlock3        `json:"menu block 3"`
	Block5  DisplayProperties `json:"menu block 5"`
	Block6  MenuBlock6        `json:"menu block 6"`
	Block8  DisplayProperties `json:"menu block 8"`
	Block9  DisplayProperties `json:"menu block 9"`
	Block11 DisplayProperties `json:"menu block 11"`
}

type MenuBlock1 struct {
	MenuName DisplayProperties `json:"menu name"`
}

type MenuBlock3 struct {
	DishGroupName DisplayProperties `json:"dish group"`
}

type MenuBlock6 struct {
	SideDishGroupName DisplayProperties `json:"sideDish group"`
}

//
//
//
//
//
//Common properties for blocks

type DisplayProperties struct {
	Properties struct {
		Lang      map[string]string `json:"label"`
		TextAlign string            `json:"text-align"`
	} `json:"properties"`
}

func (e *ExportConfig) PrepareLabels() {
	e.Sections.Header.Header1.Properties.CompanyName.prepareLabels()
	e.Sections.Header.Header1.Properties.CreationDate.prepareLabels()
	e.Sections.Header.Header2.CompanyAddress.prepareLabels()
	e.Sections.Header.Header3.CompanyTaxNumber.prepareLabels()
	e.Sections.Header.Header4.TypeOfTaxation.prepareLabels()
	e.Sections.Header.Header5.BranchName.prepareLabels()
	e.Sections.Header.Header6.BranchTaxNumber.prepareLabels()
	e.Sections.Header.Header7.BranchAddress.prepareLabels()
	//
	e.Sections.MenuSection.Block1.MenuName.prepareLabels()
	e.Sections.MenuSection.Block3.DishGroupName.prepareLabels()
	e.Sections.MenuSection.Block6.SideDishGroupName.prepareLabels()
}

func (d *DisplayProperties) prepareLabels() {
	for labelIndex := range d.Properties.Lang {
		d.Properties.Lang[labelIndex] += ": "
	}
}
