package method

import (
	"fmt"
	"strings"
)

type Method string

const (
	Delete Method = "DELETE"
	Insert Method = "INSERT"
)

const (
	DeleteLower string = "delete "
	DeleteUpper string = "DELETE "
	InsertLower string = "insert "
	InsertUpper string = "INSERT "
)

func FindMethod(sql string) (Method, error) {
	if strings.Contains(sql, DeleteLower) || strings.Contains(sql, DeleteUpper) {
		return Delete, nil
	}

	if strings.Contains(sql, InsertLower) || strings.Contains(sql, InsertUpper) {
		return Insert, nil
	}

	return "", fmt.Errorf("not found method")
}

func FindMethodIndex(sql string) (int, error) {
	if index := strings.Index(sql, DeleteLower); index >= 0 {
		return index, nil
	}

	if index := strings.Index(sql, DeleteUpper); index >= 0 {
		return index, nil
	}

	if index := strings.Index(sql, InsertLower); index >= 0 {
		return index, nil
	}

	if index := strings.Index(sql, InsertUpper); index >= 0 {
		return index, nil
	}

	return 0, fmt.Errorf("not found method")
}

func RemoveMethodAfterText(sql string) (string, error) {
	index, err := FindMethodIndex(sql)
	if err != nil {
		return sql, err
	}

	return sql[index:], nil
}
