package model

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/subramanyam/go-rest-api/database"
)

//Retrieving all Employees data
func GetAllEmp(emp *[]Employee) (err error) {
	if err = database.DB.Find(emp).Error; err != nil {
		return err
	}
	return nil
}

//Inserting data
func CreateEmp(emp *Employee) (err error) {
	if err = database.DB.Create(emp).Error; err != nil {
		return err
	}
	return nil
}

//getting employee record with ID
func GetEmpByID(emp *Employee, id string) (err error) {
	if err = database.DB.Where("id = ?", id).First(emp).Error; err != nil {
		return err
	}
	return nil
}

//Updating Employee
func EmpUpdate(emp *Employee, id string) (err error) {
	fmt.Println(emp)
	database.DB.Save(emp)
	return nil
}

//Deleting Employee
func EmpDelete(emp *Employee, id string) (err error) {
	database.DB.Where("id = ?", id).Delete(emp)
	return nil
}
