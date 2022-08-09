package export

import (
	"fmt"
	"os"

	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
	"github.com/kachan28/liefer_club/internal/models"
	branchService "github.com/kachan28/liefer_club/internal/services/branch"
	timeService "github.com/kachan28/liefer_club/internal/services/time"
)

type CreatePdfProtocol struct{}

func (create CreatePdfProtocol) CreatePdfFile(result models.ResultModel, exportConfig models.ExportConfig, outputPath string, exportLang string) error {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.SetPageMargins(10, 10, 10)

	creationDate, err := timeService.TimeService{}.GetDateStringForExport(result.CreationDate)
	if err != nil {
		return err
	}
	result.CreationDate = creationDate

	create.buildHeading(m, result, exportConfig, exportLang)

	err = m.OutputFileAndClose(outputPath)
	if err != nil {
		fmt.Println("⚠️  Could not save PDF:", err)
		os.Exit(1)
	}

	fmt.Println("PDF saved successfully")
	return nil
}

func (create CreatePdfProtocol) buildHeading(m pdf.Maroto, result models.ResultModel, exportConfig models.ExportConfig, exportLang string) error {
	create.buildHeader1(m, result.Company.Name, result.CreationDate, exportConfig, exportLang)
	create.buildDivider(m)()
	create.buildHeader2(m, result.Company.Address.PrepareAddressForExport(), exportConfig, exportLang)
	create.buildHeader3(m, result.Company.TaxNumber, exportConfig, exportLang)
	create.buildHeader4(m, result.Company.TypeOfTaxationToString(), exportConfig, exportLang)
	create.buildDivider(m)()
	//check branch for not head
	isHead, err := branchService.CheckHeadBranch{}.BranchIsHead(result.Branch.Id)
	if err != nil {
		return err
	}
	if !isHead {
		create.buildHeader5(m, result.Branch.Name, exportConfig, exportLang)
		create.buildHeader7(m, result.Branch.Address.PrepareAddressForExport(), exportConfig, exportLang)
		create.buildDivider(m)()
	}
	for _, menu := range result.Menus {
		create.buildMenuBlock(m, *menu, exportConfig, exportLang)
		create.buildDivider(m)()
	}
	return nil
}

func (create CreatePdfProtocol) buildHeader1(m pdf.Maroto, companyName, creationDate string, exportConfig models.ExportConfig, exportLang string) {
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
			m.Col(2, func() {
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

func (create CreatePdfProtocol) buildHeader2(m pdf.Maroto, companyAddress string, exportConfig models.ExportConfig, exportLang string) {
	create.buildSingleTextRowWithLabel(m, exportConfig.Sections.Header.Header2.CompanyAddress.Properties.Lang[exportLang], companyAddress)()
}

func (create CreatePdfProtocol) buildHeader3(m pdf.Maroto, companyTaxNumber string, exportConfig models.ExportConfig, exportLang string) {
	create.buildSingleTextRowWithLabel(m, exportConfig.Sections.Header.Header3.CompanyTaxNumber.Properties.Lang[exportLang], companyTaxNumber)()
}

func (create CreatePdfProtocol) buildHeader4(m pdf.Maroto, typeOfTaxation string, exportConfig models.ExportConfig, exportLang string) {
	create.buildSingleTextRowWithLabel(m, exportConfig.Sections.Header.Header4.TypeOfTaxation.Properties.Lang[exportLang], typeOfTaxation)()
}

func (create CreatePdfProtocol) buildHeader5(m pdf.Maroto, branchName string, exportConfig models.ExportConfig, exportLang string) {
	create.buildSingleTextRowWithLabel(m, exportConfig.Sections.Header.Header5.BranchName.Properties.Lang[exportLang], branchName)()
}

func (create CreatePdfProtocol) buildHeader6(m pdf.Maroto, branchTaxNumber string, exportConfig models.ExportConfig, exportLang string) {
	create.buildSingleTextRowWithLabel(m, exportConfig.Sections.Header.Header6.BranchTaxNumber.Properties.Lang[exportLang], branchTaxNumber)()
}

func (create CreatePdfProtocol) buildHeader7(m pdf.Maroto, branchAddress string, exportConfig models.ExportConfig, exportLang string) {
	create.buildSingleTextRowWithLabel(m, exportConfig.Sections.Header.Header7.BranchAddress.Properties.Lang[exportLang], branchAddress)()
}

//
//
func (create CreatePdfProtocol) buildMenuBlock(m pdf.Maroto, menu models.Menu, exportConfig models.ExportConfig, exportLang string) {
	create.buildMenuBlock1(m, menu.Name, exportConfig, exportLang)
	create.buildMenuBlock2(m, exportConfig, exportLang)
	for _, dishGroup := range menu.DishGroups {
		create.buildDishGroupBlock(m, dishGroup, exportConfig, exportLang)
	}
	create.buildMenuBlock5(m, exportConfig, exportLang)
	for _, sideDishGroup := range menu.SideDishGroups {
		create.buildSideDishGroupBlock(m, sideDishGroup, exportConfig, exportLang)
	}
	create.buildMenuBlock8(m, exportConfig, exportLang)
	create.buildMenuBlock9(m, exportConfig, exportLang)
	for _, set := range menu.SpecialOffersAndSetMenus {
		for _, component := range set.Components {
			create.buildMenuBlock10(m, component.ToString())
		}
	}
	create.buildSpacer(m, 5)()
	create.buildMenuBlock11(m, exportConfig, exportLang)
	for _, set := range menu.SpecialOffersAndSetMenus {
		for _, offer := range set.Offers {
			create.buildMenuBlock12(m, offer.ToString())
		}
	}
}

func (create CreatePdfProtocol) buildMenuBlock1(m pdf.Maroto, menuName string, exportConfig models.ExportConfig, exportLang string) {
	create.buildSingleTextRowWithLabel(m, exportConfig.Sections.MenuSection.Block1.MenuName.Properties.Lang[exportLang], menuName)()
}

func (create CreatePdfProtocol) buildMenuBlock2(m pdf.Maroto, exportConfig models.ExportConfig, exportLang string) {
	create.buildSingleTextRowWithLabel(m, exportConfig.Sections.MenuSection.Block2.Properties.Lang[exportLang], "")()
}

func (create CreatePdfProtocol) buildDishGroupBlock(m pdf.Maroto, dishGroup models.DishGroup, exportConfig models.ExportConfig, exportLang string) {
	create.buildMenuBlock3(m, dishGroup.Name, exportConfig, exportLang)
	for _, dish := range dishGroup.Dishes {
		if len(dish.Name) == 0 {
			continue
		}
		create.buildMenuBlock4(m, dish.ToString())
	}
}

func (create CreatePdfProtocol) buildMenuBlock3(m pdf.Maroto, dishGroupName string, exportConfig models.ExportConfig, exportLang string) {
	create.buildSingleTextRowWithLabel(m, exportConfig.Sections.MenuSection.Block3.DishGroupName.Properties.Lang[exportLang], dishGroupName)()
}

func (create CreatePdfProtocol) buildMenuBlock4(m pdf.Maroto, dish string) {
	create.buildSingleTextRow(m, dish, consts.Normal)()
}

func (create CreatePdfProtocol) buildMenuBlock5(m pdf.Maroto, exportConfig models.ExportConfig, exportLang string) {
	create.buildSingleTextRow(m, exportConfig.Sections.MenuSection.Block5.Properties.Lang[exportLang], consts.Bold)()
}

func (create CreatePdfProtocol) buildSideDishGroupBlock(m pdf.Maroto, sideDishGroup models.SideDishGroup, exportConfig models.ExportConfig, exportLang string) {
	create.buildMenuBlock6(m, sideDishGroup.Name, exportConfig, exportLang)
	for _, sideDish := range sideDishGroup.SideDishes {
		if len(sideDish.Name) == 0 {
			continue
		}
		create.buildMenuBlock7(m, sideDish.ToString())
	}
}

func (create CreatePdfProtocol) buildMenuBlock6(m pdf.Maroto, sideDishGroupName string, exportConfig models.ExportConfig, exportLang string) {
	create.buildSingleTextRowWithLabel(m, exportConfig.Sections.MenuSection.Block6.SideDishGroupName.Properties.Lang[exportLang], sideDishGroupName)()
}

func (create CreatePdfProtocol) buildMenuBlock7(m pdf.Maroto, sideDish string) {
	create.buildSingleTextRow(m, sideDish, consts.Normal)()
}

func (create CreatePdfProtocol) buildMenuBlock8(m pdf.Maroto, exportConfig models.ExportConfig, exportLang string) {
	create.buildSingleTextRow(m, exportConfig.Sections.MenuSection.Block8.Properties.Lang[exportLang], consts.Bold)()
}

func (create CreatePdfProtocol) buildMenuBlock9(m pdf.Maroto, exportConfig models.ExportConfig, exportLang string) {
	create.buildSingleTextRow(m, exportConfig.Sections.MenuSection.Block9.Properties.Lang[exportLang], consts.Bold)()
}

func (create CreatePdfProtocol) buildMenuBlock10(m pdf.Maroto, component string) {
	create.buildSingleTextRow(m, component, consts.Normal)()
}

func (create CreatePdfProtocol) buildMenuBlock11(m pdf.Maroto, exportConfig models.ExportConfig, exportLang string) {
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
