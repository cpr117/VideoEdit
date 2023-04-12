// @User CPR
package dao

import (
	"errors"
	"gorm.io/gorm"
)

// 数据库对象
var DB *gorm.DB

// 通用CRUD

func Create[T any](data *T) {
	err := DB.Create(&data).Error
	if err != nil {
		panic(err)
	}
}

func GetOne[T any](query string, args ...any) (data T) {
	err := DB.Where(query, args...).First(&data).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		panic(err)
	}
	return
}

func Update[T any](data *T, slt ...string) {
	var err error
	db := DB.Model(&data)
	if len(slt) > 0 {
		err = db.Select(slt).Updates(&data).Error
	} else {
		err = db.Updates(&data).Error
	}
	if err != nil {
		panic(err)
	}
}

// UpdatesMap
// [批量]更新: map 的字段就是要更新的字段 (map 可以更新零值), 通过条件可以实现[单行]更新
func UpdatesMap[T any](maps map[string]any, query string, args ...any) {
	err := DB.Model(new(T)).Where(query, args...).Updates(maps).Error
	if err != nil {
		panic(err)
	}
}

func Updates[T any](data *T, query string, args ...any) {
	err := DB.Model(new(T)).Where(query, args...).Updates(&data).Error
	if err != nil {
		panic(err)
	}
}

func List[T any](slt, order, query string, args ...any) (data T) {
	db := DB.Model(new(T)).Select(slt)
	if order != "" {
		db = db.Order(order)
	}
	if query != "" {
		db = db.Where(query, args...)
	}
	err := db.Find(&data).Error
	if err != nil {
		panic(err)
	}
	return
}

// Delete
// 删除
func Delete[T any](query string, args ...any) {
	db := DB
	if query != "" {
		db = db.Where(query, args...)
	}
	if err := db.Delete(new(T)).Error; err != nil {
		panic(err)
	}

}

// Count
func Count[T any](query string, args ...any) (cnt int64) {
	db := DB.Model(new(T))
	if query != "" {
		db = db.Where(query, args...)
	}
	if err := db.Count(&cnt).Error; err != nil {
		panic(err)
	}
	return
}

// First
func First[T any](query string, args ...any) (data T) {
	db := DB.Where(query, args...)
	err := db.First(&data).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		panic(err)
	}
	return
}
