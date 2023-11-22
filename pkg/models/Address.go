package models

// // user can see products list
func GetAllAddresses() []Address {
	var addresses []Address
	DB.Find(&addresses)
	return addresses
}

// // user can see product page
func GetAddressById(id int) Address {
	var address Address
	DB.Where("id = ?", id).First(&address)
	return address
}

func GetAddressByName(name string) Address {
	var address Address
	DB.Where("name = ?", name).First(&address)
	return address
}

func CreateAddress(address Address) (Address, error) {
	resp := DB.Create(&address)
	if resp.Error != nil {
		return address, resp.Error
	} else {
		return address, nil
	}
}
func UpdateAddress(id int, address Address) Address {
	DB.Where("id=?", id).Updates(&address)
	return address
}
func DeleteAddress(id int) Address {
	var address Address
	DB.Where("id=?", id).First(&address)
	DB.Delete(address)
	return address
}
