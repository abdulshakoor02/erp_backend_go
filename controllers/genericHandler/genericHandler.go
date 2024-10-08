package genericHandler

import (
	"encoding/json"
	"fmt"

	"github.com/abdul/erp_backend/database/dbAdapter"
	"github.com/abdul/erp_backend/logger"
	"github.com/gofiber/fiber/v2"
)

var log = logger.Logger

func CreateHandler[T any](c *fiber.Ctx) error {
	// Get the authorization header
	// authHeader := c.Get("Authorization")

	// Get the JSON object from the body
	var genericData []T

	err := json.Unmarshal(c.Body(), &genericData)
	if err != nil {
		fmt.Println(err)
		log.Info().Msgf("error  %v", err)
		return c.Status(fiber.StatusBadRequest).SendString("Invalid JSON")
	}
	dbAdapter.DB.Debug()

	// Marshal the struct back to a JSON string
	newJSONData, err := json.Marshal(genericData)
	if err != nil {
		fmt.Println("Error:", err)
	}
	if err := dbAdapter.DB.Create(&genericData).Error; err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).SendString(fmt.Sprintf("%d", err))
	}

	// Print the JSON object and authorization header
	fmt.Println(string(newJSONData))

	// Return a 200 OK response
	return c.Status(fiber.StatusOK).SendString("Data created successfully")

}

func FindHandler[T any](c *fiber.Ctx) error {
	// Get the authorization header
	// authHeader := c.Get("Authorization")
	var response struct {
		Data  []T   `json:"data"`
		Count int64 `json:"count"`
	}
	var Count int64
	// Get the JSON object from the body
	var genericData struct {
		Where  T
		Limit  int32
		Offset int32
	}
	err := json.Unmarshal(c.Body(), &genericData)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).SendString("Invalid JSON")
	}

	db := dbAdapter.DB
	// Use the gorm.Statement to exclude certain fields from the Where clause
	var result []T

	if err := db.Model(&result).Where(&genericData.Where).Count(&Count).Error; err != nil {
		log.Info().Msgf("error  %v", err)
		return c.Status(fiber.StatusBadRequest).SendString("error could not process the query")
	}
	db = dbAdapter.DB

	if int(genericData.Limit) != 0 {
		db = db.Limit(int(genericData.Limit)).
			Offset(int(genericData.Offset)).
			Where(&genericData.Where)
	} else {
		// db = db.Where(&genericData.Where).Find(&result)
		db = db.Where(&genericData.Where)
	}

	if err := db.Find(&result).Error; err != nil {
		log.Err(err).Msgf("error  %v", err)
		return c.Status(fiber.StatusBadRequest).SendString("error could not process the query")
	}

	response.Count = Count
	response.Data = result

	// dbAdapter.DB.Find(&result)
	newJSONData2, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Error:", err)
	}
	// Print the JSON object and authorization header
	// fmt.Println(string(newJSONData))

	// Return a 200 OK response

	return c.Status(fiber.StatusOK).SendString(string(newJSONData2))
	// return c.Status(fiber.StatusOK).SendString("Data created successfully")
}

func UpdateHandler[W any, T any](c *fiber.Ctx) error {
	// Get the authorization header
	// authHeader := c.Get("Authorization")

	// Get the JSON object from the body
	var genericData struct {
		Where W
		Data  T
	}
	err := json.Unmarshal(c.Body(), &genericData)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).SendString("Invalid JSON")
	}

	// // Marshal the struct back to a JSON string
	// newJSONData, err := json.Marshal(genericData)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// }

	db := dbAdapter.DB
	// Use the gorm.Statement to exclude certain fields from the Where clause
	var result []T

	if err := db.Model(&result).Where(&genericData.Where).Updates(&genericData.Data).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("error could not process the query")
	}

	// dbAdapter.DB.Find(&result)
	// newJSONData2, err := json.Marshal(result)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// }
	// Print the JSON object and authorization header
	// fmt.Println(string(newJSONData))

	// Return a 200 OK response

	return c.Status(fiber.StatusOK).SendString("data updated")
	// return c.Status(fiber.StatusOK).SendString("Data created successfully")
}

func DeleteHandler[W any, T any](c *fiber.Ctx) error {
	// Get the authorization header
	// authHeader := c.Get("Authorization")

	// Get the JSON object from the body
	var Where W
	err := json.Unmarshal(c.Body(), &Where)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).SendString("Invalid JSON")
	}

	// // Marshal the struct back to a JSON string
	// newJSONData, err := json.Marshal(genericData)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// }

	db := dbAdapter.DB
	// Use the gorm.Statement to exclude certain fields from the Where clause
	var result []T

	if err := db.Where(Where).Delete(&result).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("error could not process the query")
	}

	// dbAdapter.DB.Find(&result)
	// newJSONData2, err := json.Marshal(result)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// }
	// Print the JSON object and authorization header
	// fmt.Println(string(newJSONData))

	// Return a 200 OK response

	return c.Status(fiber.StatusOK).SendString("data deleted")
	// return c.Status(fiber.StatusOK).SendString("Data created successfully")
}

func FindAssociatedHandler[T any](c *fiber.Ctx) error {

	type Associated struct {
		Column string                 `json:"column"`
		Where  map[string]interface{} `json:"where"`
	}

	type Query struct {
		Column   string `json:"column"`
		Operator string `json:"operator"`
		Value    string `json:"value"`
	}

	var response struct {
		Data  []T   `json:"data"`
		Count int64 `json:"count"`
	}
	var Count int64
	// Get the JSON object from the body
	var genericData struct {
		Where  []Query
		Joins  []Associated
		Limit  int32
		Offset int32
	}

	err := json.Unmarshal(c.Body(), &genericData)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).SendString("Invalid JSON")
	}

	newJSONData3, err := json.Marshal(genericData)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(string(newJSONData3))

	db := dbAdapter.DB
	// Use the gorm.Statement to exclude certain fields from the Where clause
	var result []T
	db = db.Model(&result)
	for _, v := range genericData.Joins {
		db = db.InnerJoins(v.Column)
		for i, v2 := range v.Where {
			cols := fmt.Sprintf("\"%v\".%v = ?", v.Column, i)
			db = db.Where(cols, v2)
		}
	}

	for i, v := range genericData.Where {
		if i == 0 {
			cols := fmt.Sprintf("%v %v ?", v.Column, v.Operator)
			fmt.Println(cols, v.Value)
			db = db.Where(cols, v.Value)
			continue
		}
		cols := fmt.Sprintf("%v %v ?", v.Column, v.Operator)
		fmt.Println(cols, v.Value)
		db = db.Or(cols, v.Value)
	}

	if err := db.Count(&Count).Error; err != nil {
		log.Info().Msgf("error  %v", err)
		return c.Status(fiber.StatusBadRequest).SendString("error could not process the query")
	}

	// if err := db.Model(&result).InnerJoins(joinWith, db.Where(&genericData.JoinWhere)).Where(&genericData.Where).Count(&Count).Error; err != nil {
	// 	log.Info().Msgf("error  %v", err)
	// 	return c.Status(fiber.StatusBadRequest).SendString("error could not process the query")
	// }
	db = dbAdapter.DB

	db = db.Model(&result)

	for _, v := range genericData.Joins {
		db = db.InnerJoins(v.Column)

		for i, v2 := range v.Where {
			cols := fmt.Sprintf("\"%v\".%v = ?", v.Column, i)
			fmt.Println(cols)
			db = db.Where(cols, v2)
		}
	}

	for i, v := range genericData.Where {
		if i == 0 {
			cols := fmt.Sprintf("%v %v ?", v.Column, v.Operator)
			db = db.Where(cols, v.Value)
			continue
		}
		cols := fmt.Sprintf("%v %v ?", v.Column, v.Operator)
		db = db.Or(cols, v.Value)
	}
	// db = db.Where(&genericData.Where)

	if genericData.Limit != 0 {
		db = db.Limit(int(genericData.Limit)).Offset(int(genericData.Offset))
	}

	// db = db.Model(&result).
	// 	InnerJoins(joinWith, db.Where(&genericData.JoinWhere)).
	// 	Where(&genericData.Where)
	// if genericData.Limit != 0 {
	// 	db = db.Limit(int(genericData.Limit)).Offset(int(genericData.Offset))
	// }
	if err := db.Find(&result).Error; err != nil {
		log.Err(err).Msgf("error  %v", err)
		return c.Status(fiber.StatusBadRequest).SendString("error could not process the query")
	}
	// if err := db.Raw("select tenants.id,tenants.name,phone,email,website,country_id,countries.name as country_name,status from tenants join countries on tenants.country_id = countries.id").Find(&result).Error; err != nil {
	// 	log.Err(err).Msgf("error  %v", err)
	// 	return c.Status(fiber.StatusBadRequest).SendString("error could not process the query")
	// }

	response.Count = Count
	response.Data = result

	// dbAdapter.DB.Find(&result)
	newJSONData2, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Error:", err)
	}
	// Print the JSON object and authorization header
	// fmt.Println(string(newJSONData))

	// Return a 200 OK response

	return c.Status(fiber.StatusOK).SendString(string(newJSONData2))
	// return c.Status(fiber.StatusOK).SendString("Data created successfully")
}
