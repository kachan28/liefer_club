package branch

import (
	"github.com/kachan28/liefer_club/app"
	"github.com/kachan28/liefer_club/internal/db"
)

const (
	dbName = "firma_pizzanova_db"
	head   = 1
)

type CheckHeadBranch struct{}

func (c CheckHeadBranch) BranchIsHead(branchID int64) (bool, error) {
	conf := new(app.Conf)
	conf.GetConf()
	conn, err := db.MakeConnection(conf, dbName)
	if err != nil {
		return false, err
	}
	isHead, err := conn.GetBranchHeadParam(branchID)
	if isHead == head {
		return true, nil
	}
	return false, nil
}
