package utils

import (
	"bizd/metion/model"
	"unsafe"
)

func EncapsulationPage(pageNumber int, pageSize int, tmp int64) model.ResponsePagination {
	var pagination model.ResponsePagination
	var total = *(*int)(unsafe.Pointer(&tmp))
	var totalPages int
	if total%pageSize == 0 {
		totalPages = total / pageSize
	} else {
		totalPages = total / (pageSize + 1)
	}
	pagination.Total = total
	if pageNumber-1 > 0 {
		pagination.Prev = pageNumber - 1
	} else {
		pagination.Prev = -1
	}
	if pageNumber+1 <= totalPages {
		pagination.Next = pageNumber + 1
	} else {
		pagination.Next = -1
	}
	pagination.Cur = pageNumber
	return pagination
}
