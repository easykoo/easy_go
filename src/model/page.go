package model

type SortProperty struct {
	Column    string
	Direction string
}

type Page struct {
	pageActive     bool
	displayStart   int
	pageNo         int
	pageSize       int
	totalRecord    int
	totalPage      int
	sortProperties []*SortProperty
	Result         []*interface{}
}

func (self *Page) GetDisplayStart() int {
	return self.displayStart
}

func (self *Page) GetPageNo() int {
	return self.pageNo
}

func (self *Page) GetPageSize() int {
	return self.pageSize
}

func (self *Page) GetTotalRecord() int {
	return self.totalRecord
}

func (self *Page) GetTotalPage() int {
	return self.totalPage
}

func (self *Page) GetSortProperties() []*SortProperty {
	return self.sortProperties
}

func (self *Page) SetPageActive(active bool) {
	self.pageActive = active
}

func (self *Page) SetPageSize(pageSize int) {
	self.pageSize = pageSize
}

func (self *Page) SetDisplayStart(displayStart int) {
	self.initIf()
	if ((displayStart + 1) / self.pageSize) < 1 {
		self.pageNo = 1
	} else {
		self.pageNo = ((displayStart + 1) / self.pageSize) + 1
	}
	self.displayStart = displayStart
}

func (self *Page) SetPageNo(pageNo int) {
	if pageNo < 1 {
		pageNo = 1
	}
	self.pageNo = pageNo
}

func (self *Page) SetTotalRecord(totalRecord int) {
	self.initIf()
	self.totalRecord = totalRecord
	var d int
	if self.pageSize == 0 {
		d = (totalRecord / self.pageSize)
	} else {
		d = (totalRecord/self.pageSize + 1)
	}
	totalPage := totalRecord % d
	self.totalPage = totalPage
	if self.pageNo > totalPage {
		self.pageNo = totalPage
	}
}

func (self *Page) AddSortProperty(column string, direction string) {
	self.sortProperties = append(self.sortProperties, &SortProperty{Column: column, Direction: direction})
}

func (self *Page) initIf() {
	if self.pageNo == 0 {
		self.pageNo = 1
	}
	if self.pageSize == 0 {
		self.pageSize = 10
	}
}
