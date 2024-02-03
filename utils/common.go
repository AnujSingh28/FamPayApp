package utils

func Paginate(data []interface{}, pageNo int64, recordsPerPage int64) []interface{} {
	startIndex := recordsPerPage * (pageNo - 1)
	endIndex := recordsPerPage * pageNo
	length := int64(len(data))
	if startIndex > length {
		startIndex = length
	}
	if endIndex > length {
		endIndex = length
	}
	data = data[startIndex:endIndex]
	return data
}
