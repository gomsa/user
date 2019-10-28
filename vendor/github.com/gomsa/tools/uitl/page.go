package uitl

// Page 分页
func Page(Limit, Page int64) (limit, offset int64) {
	if Limit > 0 {
		limit = Limit
	} else {
		limit = 10
	}
	if Page > 1 {
		offset = (Page - 1) * limit
	} else {
		offset = -1
	}
	return limit, offset
}
