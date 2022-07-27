package app

type Conf struct {
	Dbs        databaseList
	Connection connection
}

func (c *Conf) GetConf() {
	c.Dbs.getConf()
	c.Connection.getConf()
}
