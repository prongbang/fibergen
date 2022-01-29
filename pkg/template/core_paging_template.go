package template

type corePagingTemplate struct {
}

func (c *corePagingTemplate) Text() []byte {
	return []byte("package core\n\nimport (\n\t\"math\"\n)\n\nconst PagingLimitDefault = 20\n\ntype Paging struct {\n\tList  interface{} `json:\"list\"`\n\tPage  int64       `json:\"page\"`\n\tLimit int64       `json:\"limit\"`\n\tCount int64       `json:\"count\"`\n\tTotal int64       `json:\"total\"`\n\tStart int64       `json:\"start\"`\n\tEnd   int64       `json:\"end\"`\n}\n\nfunc Offset(pageNo int64, limitNo int64) int64 {\n\treturn (limitNo * pageNo) - limitNo\n}\n\nfunc Pagination(pageNo int64, limitNo int64, getCount func() int64, getData func(limit int64, offset int64) interface{}) Paging {\n\tvar pageCount = math.Ceil(float64(getCount()) / float64(limitNo))\n\tpageCountInt := int64(pageCount)\n\tif pageNo <= 0 {\n\t\tpageNo = 1\n\t}\n\toffset := (limitNo * pageNo) - limitNo\n\tvar startRow = (pageNo - 1) * limitNo\n\tvar endRow = startRow + limitNo - 1\n\n\tdata := getData(limitNo, offset)\n\n\treturn Paging{\n\t\tList:  data,\n\t\tPage:  pageNo,\n\t\tLimit: limitNo,\n\t\tCount: pageCountInt,\n\t\tTotal: getCount(),\n\t\tStart: startRow,\n\t\tEnd:   endRow,\n\t}\n}\n")
}

func CorePagingTemplate() Template {
	return &corePagingTemplate{}
}
