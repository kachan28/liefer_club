package export

import (
	"fmt"
	"os"

	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
	"github.com/kachan28/liefer_club/internal/models"
	timeService "github.com/kachan28/liefer_club/internal/services/time"
)

const exportLang = "de"

type CreatePdfProtocol struct{}

func (create CreatePdfProtocol) CreatePdfFile(result models.ResultModel, exportConfig models.ExportConfig, outputPath string) error {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.SetPageMargins(10, 10, 10)

	creationDate, err := timeService.TimeService{}.GetDateStringForExport(result.CreationDate)
	if err != nil {
		return err
	}
	result.CreationDate = creationDate

	create.buildHeading(m, result, exportConfig)
	// buildFooter(m)
	// buildFruitList(m)
	// buildSignature(m)

	err = m.OutputFileAndClose(outputPath)
	if err != nil {
		fmt.Println("⚠️  Could not save PDF:", err)
		os.Exit(1)
	}

	fmt.Println("PDF saved successfully")
	return nil
}

func (create CreatePdfProtocol) buildHeading(m pdf.Maroto, result models.ResultModel, exportConfig models.ExportConfig) {
	create.buildHeader1(m, result.Company.Name, result.CreationDate, exportConfig)
	create.buildDivider(m)()
	create.buildHeader2(m, result.Company.Address.PrepareAddressForExport(), exportConfig)
	create.buildHeader3(m, result.Company.TaxNumber, exportConfig)
	create.buildHeader4(m, result.Company.TypeOfTaxation, exportConfig)
	create.buildDivider(m)()
	create.buildHeader5(m, result.Branch.Name, exportConfig)
	create.buildHeader7(m, result.Branch.Address.PrepareAddressForExport(), exportConfig)
	create.buildDivider(m)()
	for _, menu := range result.Menus {
		create.buildMenuBlock(m, *menu, exportConfig)
		create.buildDivider(m)()
	}
	// m.Row(10, func() {
	// 	m.Col(12, func() {
	// 		_ = m.Barcode("https://divrhino.com", props.Barcode{
	// 			Percent:    75,
	// 			Proportion: props.Proportion{Width: 50, Height: 10},
	// 			Center:     true,
	// 		})
	// 	})
	// })
}

func (create CreatePdfProtocol) buildHeader1(m pdf.Maroto, companyName, creationDate string, exportConfig models.ExportConfig) {
	m.Row(15, func() {
		m.Col(2, func() {
			m.Col(2, func() {
				m.Text(exportConfig.Sections.Header.Header1.Properties.CompanyName.Properties.Lang[exportLang], props.Text{
					Style: consts.Bold,
					Align: consts.Left,
				})
			})
			m.Col(3, func() {
				m.Text(companyName, props.Text{
					Style: consts.Bold,
					Align: consts.Left,
				})
			})
		})
		m.Col(4, func() {
			m.Col(3, func() {
				m.Text(exportConfig.Sections.Header.Header1.Properties.CreationDate.Properties.Lang[exportLang], props.Text{
					Style: consts.Bold,
					Align: consts.Left,
				})
			})
			m.Col(3, func() {
				m.Text(creationDate, props.Text{
					Style: consts.Bold,
					Align: consts.Left,
				})
			})
		})
	})
}

func (create CreatePdfProtocol) buildHeader2(m pdf.Maroto, companyAddress string, exportConfig models.ExportConfig) {
	create.buildSingleTextRowWithLabel(m, exportConfig.Sections.Header.Header2.CompanyAddress.Properties.Lang[exportLang], companyAddress)()
}

func (create CreatePdfProtocol) buildHeader3(m pdf.Maroto, companyTaxNumber string, exportConfig models.ExportConfig) {
	create.buildSingleTextRowWithLabel(m, exportConfig.Sections.Header.Header3.CompanyTaxNumber.Properties.Lang[exportLang], companyTaxNumber)()
}

func (create CreatePdfProtocol) buildHeader4(m pdf.Maroto, typeOfTaxation string, exportConfig models.ExportConfig) {
	create.buildSingleTextRowWithLabel(m, exportConfig.Sections.Header.Header4.TypeOfTaxation.Properties.Lang[exportLang], typeOfTaxation)()
}

func (create CreatePdfProtocol) buildHeader5(m pdf.Maroto, branchName string, exportConfig models.ExportConfig) {
	create.buildSingleTextRowWithLabel(m, exportConfig.Sections.Header.Header5.BranchName.Properties.Lang[exportLang], branchName)()
}

func (create CreatePdfProtocol) buildHeader6(m pdf.Maroto, branchTaxNumber string, exportConfig models.ExportConfig) {
	create.buildSingleTextRowWithLabel(m, exportConfig.Sections.Header.Header6.BranchTaxNumber.Properties.Lang[exportLang], branchTaxNumber)()
}

func (create CreatePdfProtocol) buildHeader7(m pdf.Maroto, branchAddress string, exportConfig models.ExportConfig) {
	create.buildSingleTextRowWithLabel(m, exportConfig.Sections.Header.Header7.BranchAddress.Properties.Lang[exportLang], branchAddress)()
}

//
//
func (create CreatePdfProtocol) buildMenuBlock(m pdf.Maroto, menu models.Menu, exportConfig models.ExportConfig) {
	create.buildMenuBlock1(m, menu.Name, exportConfig)
	create.buildMenuBlock2(m, exportConfig)
	for _, dishGroup := range menu.DishGroups {
		create.buildDishGroupBlock(m, dishGroup, exportConfig)
	}
	create.buildMenuBlock5(m, exportConfig)
	for _, sideDishGroup := range menu.SideDishGroups {
		create.buildSideDishGroupBlock(m, sideDishGroup, exportConfig)
	}
	create.buildMenuBlock8(m, exportConfig)
	create.buildMenuBlock9(m, exportConfig)
	for _, set := range menu.SpecialOffersAndSetMenus {
		for _, component := range set.Components {
			create.buildMenuBlock10(m, component.ToString())
		}
	}
	create.buildSpacer(m, 5)()
	create.buildMenuBlock11(m, exportConfig)
	for _, set := range menu.SpecialOffersAndSetMenus {
		for _, offer := range set.Offers {
			create.buildMenuBlock12(m, offer.ToString())
		}
	}
}

func (create CreatePdfProtocol) buildMenuBlock1(m pdf.Maroto, menuName string, exportConfig models.ExportConfig) {
	create.buildSingleTextRowWithLabel(m, exportConfig.Sections.MenuSection.Block1.MenuName.Properties.Lang[exportLang], menuName)()
}

func (create CreatePdfProtocol) buildMenuBlock2(m pdf.Maroto, exportConfig models.ExportConfig) {
	create.buildSingleTextRowWithLabel(m, exportConfig.Sections.MenuSection.Block2.Properties.Lang[exportLang], "")()
}

func (create CreatePdfProtocol) buildDishGroupBlock(m pdf.Maroto, dishGroup models.DishGroup, exportConfig models.ExportConfig) {
	create.buildMenuBlock3(m, dishGroup.Name, exportConfig)
	for _, dish := range dishGroup.Dishes {
		create.buildMenuBlock4(m, dish.ToString())
	}
}

func (create CreatePdfProtocol) buildMenuBlock3(m pdf.Maroto, dishGroupName string, exportConfig models.ExportConfig) {
	create.buildSingleTextRowWithLabel(m, exportConfig.Sections.MenuSection.Block3.DishGroupName.Properties.Lang[exportLang], dishGroupName)()
}

func (create CreatePdfProtocol) buildMenuBlock4(m pdf.Maroto, dish string) {
	create.buildSingleTextRow(m, dish, consts.Normal)()
}

func (create CreatePdfProtocol) buildMenuBlock5(m pdf.Maroto, exportConfig models.ExportConfig) {
	create.buildSingleTextRow(m, exportConfig.Sections.MenuSection.Block5.Properties.Lang[exportLang], consts.Bold)()
}

func (create CreatePdfProtocol) buildSideDishGroupBlock(m pdf.Maroto, sideDishGroup models.SideDishGroup, exportConfig models.ExportConfig) {
	create.buildMenuBlock6(m, sideDishGroup.Name, exportConfig)
	for _, sideDish := range sideDishGroup.SideDishes {
		create.buildMenuBlock7(m, sideDish.ToString())
	}
}

func (create CreatePdfProtocol) buildMenuBlock6(m pdf.Maroto, sideDishGroupName string, exportConfig models.ExportConfig) {
	create.buildSingleTextRowWithLabel(m, exportConfig.Sections.MenuSection.Block6.SideDishGroupName.Properties.Lang[exportLang], sideDishGroupName)()
}

func (create CreatePdfProtocol) buildMenuBlock7(m pdf.Maroto, sideDish string) {
	create.buildSingleTextRow(m, sideDish, consts.Normal)()
}

func (create CreatePdfProtocol) buildMenuBlock8(m pdf.Maroto, exportConfig models.ExportConfig) {
	create.buildSingleTextRow(m, exportConfig.Sections.MenuSection.Block8.Properties.Lang[exportLang], consts.Bold)()
}

func (create CreatePdfProtocol) buildMenuBlock9(m pdf.Maroto, exportConfig models.ExportConfig) {
	create.buildSingleTextRow(m, exportConfig.Sections.MenuSection.Block9.Properties.Lang[exportLang], consts.Bold)()
}

func (create CreatePdfProtocol) buildMenuBlock10(m pdf.Maroto, component string) {
	create.buildSingleTextRow(m, component, consts.Normal)()
}

func (create CreatePdfProtocol) buildMenuBlock11(m pdf.Maroto, exportConfig models.ExportConfig) {
	create.buildSingleTextRow(m, exportConfig.Sections.MenuSection.Block11.Properties.Lang[exportLang], consts.Bold)()
}

func (create CreatePdfProtocol) buildMenuBlock12(m pdf.Maroto, offer string) {
	create.buildSingleTextRow(m, offer, consts.Normal)()
}

//
//common pdf build funcs
func (create CreatePdfProtocol) buildSingleTextRow(m pdf.Maroto, value string, textBold consts.Style) func() {
	return func() {
		m.Row(10, func() {
			m.Col(12, func() {
				m.Text(value, props.Text{
					Style:           textBold,
					Align:           consts.Left,
					VerticalPadding: 1.2,
				})
			})
		})
	}
}

func (create CreatePdfProtocol) buildSingleTextRowWithLabel(m pdf.Maroto, label, value string) func() {
	return func() {
		m.Row(10, func() {
			m.Col(3, func() {
				m.Text(label, props.Text{
					Style: consts.Bold,
					Align: consts.Left,
				})
			})
			m.Col(9, func() {
				m.Text(value, props.Text{
					Style: consts.Normal,
					Align: consts.Left,
				})
			})
		})
	}
}

func (create CreatePdfProtocol) buildDivider(m pdf.Maroto) func() {
	return func() {
		m.Row(3, func() {
			m.Line(3)
		})
	}
}

func (create CreatePdfProtocol) buildSpacer(m pdf.Maroto, height float64) func() {
	return func() {
		m.Row(height, func() {})
	}
}

// func buildFruitList(m pdf.Maroto) {
// 	headings := getHeadings()
// 	// contents := data.FruitList(20)
// 	contents := [][]string{{"Apple", "Red and juicy", "2.00"}, {"Orange", "Orange and juicy", "3.00"}}
// 	purpleColor := getPurpleColor()

// 	m.SetBackgroundColor(getTealColor())
// 	m.Row(10, func() {
// 		m.Col(12, func() {
// 			m.Text("Grocery List", props.Text{
// 				Top:    2,
// 				Size:   13,
// 				Color:  color.NewWhite(),
// 				Family: consts.Courier,
// 				Style:  consts.Bold,
// 				Align:  consts.Center,
// 			})
// 		})
// 	})

// 	m.SetBackgroundColor(color.NewWhite())

// 	m.TableList(headings, contents, props.TableList{
// 		HeaderProp: props.TableListContent{
// 			Size:      9,
// 			GridSizes: []uint{3, 7, 2},
// 		},
// 		ContentProp: props.TableListContent{
// 			Size:      8,
// 			GridSizes: []uint{3, 7, 2},
// 		},
// 		Align:                consts.Left,
// 		AlternatedBackground: &purpleColor,
// 		HeaderContentSpace:   1,
// 		Line:                 false,
// 	})

// 	m.Row(20, func() {
// 		m.ColSpace(7)
// 		m.Col(2, func() {
// 			m.Text("Total:", props.Text{
// 				Top:   5,
// 				Style: consts.Bold,
// 				Size:  8,
// 				Align: consts.Right,
// 			})
// 		})
// 		m.Col(3, func() {
// 			m.Text("$ XXXX.00", props.Text{
// 				Top:   5,
// 				Style: consts.Bold,
// 				Size:  8,
// 				Align: consts.Center,
// 			})
// 		})
// 	})
// }

// func buildSignature(m pdf.Maroto) {
// 	m.Row(15, func() {
// 		m.Col(5, func() {
// 			m.QrCode("https://divrhino.com", props.Rect{
// 				Left:    0,
// 				Top:     5,
// 				Center:  false,
// 				Percent: 100,
// 			})
// 		})

// 		m.ColSpace(2)

// 		m.Col(5, func() {
// 			m.Signature("Signed by", props.Font{
// 				Size:   8,
// 				Style:  consts.Italic,
// 				Family: consts.Courier,
// 			})
// 		})
// 	})
// }

// func buildFooter(m pdf.Maroto) {
// 	begin := time.Now()
// 	m.SetAliasNbPages("{nb}")
// 	m.SetFirstPageNb(1)

// 	m.RegisterFooter(func() {
// 		m.Row(20, func() {
// 			m.Col(6, func() {
// 				m.Text(begin.Format("02/01/2006"), props.Text{
// 					Top:   10,
// 					Size:  8,
// 					Color: getGreyColor(),
// 					Align: consts.Left,
// 				})
// 			})

// 			m.Col(6, func() {
// 				m.Text("Page "+strconv.Itoa(m.GetCurrentPage())+" of {nb}", props.Text{
// 					Top:   10,
// 					Size:  8,
// 					Style: consts.Italic,
// 					Color: getGreyColor(),
// 					Align: consts.Right,
// 				})
// 			})

// 		})
// 	})
// }

// func getHeadings() []string {
// 	return []string{"Fruit", "Description", "Price"}
// }

// // Colours

// func getPurpleColor() color.Color {
// 	return color.Color{
// 		Red:   210,
// 		Green: 200,
// 		Blue:  230,
// 	}
// }

// func getTealColor() color.Color {
// 	return color.Color{
// 		Red:   3,
// 		Green: 166,
// 		Blue:  166,
// 	}
// }

// func getGreyColor() color.Color {
// 	return color.Color{
// 		Red:   206,
// 		Green: 206,
// 		Blue:  206,
// 	}
// }
