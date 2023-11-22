package models

func GetAllCategories() []Category {
	var categories []Category
	DB.Find(&categories)
	return categories
}

func GetCategoryById(id int) Category {
	var category Category
	DB.Where("id = ?", id).First(&category)
	return category
}

func GetCategoryByName(name string) Category {
	var category Category
	DB.Where("name = ?", name).First(&category)
	return category
}

func CreateCategory(category Category) Category {
	DB.Create(&category)
	return category
}

func UpdateCategory(id int, category Category) Category {
	DB.Where("id= ? ", id).Updates(&category)
	return category
}

func DeleteCategory(id int) Category {
	var category Category
	DB.Where("id= ? ", id).First(&category)
	DB.Delete(category)
	return category
}
