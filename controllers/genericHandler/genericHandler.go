package genericHandler

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/abdul/erp_backend/database/dbAdapter"
	"github.com/abdul/erp_backend/logger"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	// dbAdapter.DB.Debug()

	// Marshal the struct back to a JSON string
	if err := dbAdapter.DB.Create(&genericData).Error; err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).SendString(fmt.Sprintf("%d", err))
	}

	newJSONData2, err := json.Marshal(genericData)
	if err != nil {
		fmt.Println("Error:", err)
	}
	// Print the JSON object and authorization header

	// Return a 200 OK response
	return c.Status(fiber.StatusOK).SendString(string(newJSONData2))

}

func FindHandler[T any](c *fiber.Ctx) error {
	// Get the authorization header
	// authHeader := c.Get("Authorization")

	var wg sync.WaitGroup

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
		Where   T
		Find    []Query `json:"find"`
		Limit   int32
		Offset  int32
		OrderBy string
	}
	err := json.Unmarshal(c.Body(), &genericData)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).SendString("Invalid JSON")
	}

	// Use the gorm.Statement to exclude certain fields from the Where clause
	var result []T

	var countErr, resultErr error

	wg.Add(2)

	go func() {
		defer wg.Done()

		db := dbAdapter.DB
		for i, v := range genericData.Find {
			if i == 0 {
				cols := fmt.Sprintf("%v %v ?", v.Column, v.Operator)
				db = db.Where(cols, v.Value)
				continue
			}
			cols := fmt.Sprintf("%v %v ?", v.Column, v.Operator)
			db = db.Or(cols, v.Value)
		}
		countErr = db.Model(&result).Where(&genericData.Where).Count(&Count).Error
	}()

	go func() {
		defer wg.Done()

		db := dbAdapter.DB

		for i, v := range genericData.Find {
			if i == 0 {
				cols := fmt.Sprintf("%v %v ?", v.Column, v.Operator)
				db = db.Where(cols, v.Value)
				continue
			}
			cols := fmt.Sprintf("%v %v ?", v.Column, v.Operator)
			db = db.Or(cols, v.Value)
		}
		if int(genericData.Limit) != 0 {
			db = db.Limit(int(genericData.Limit)).
				Offset(int(genericData.Offset)).
				Where(&genericData.Where)
		} else {
			// db = db.Where(&genericData.Where).Find(&result)
			db = db.Where(&genericData.Where)
		}

		if genericData.OrderBy != "" {
			db = db.Order(genericData.OrderBy)
		}

		resultErr = db.Find(&result).Error
	}()

	wg.Wait()

	if countErr != nil || resultErr != nil {
		log.Err(countErr).Err(resultErr).Msg("Query error")
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

	if err := db.Where(Where).Unscoped().Delete(&result).Error; err != nil {
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
		Column  string
		OrderBy string
		Find    map[string]any
		Where   []Query
		Joins   []Associated
		Limit   int32
		Offset  int32
	}

	err := json.Unmarshal(c.Body(), &genericData)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).SendString("Invalid JSON")
	}

	tenantId := c.Locals("tenant_id")

	// newJSONData3, err := json.Marshal(genericData)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// }
	// fmt.Println(string(newJSONData3))

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
	var orCondCount *gorm.DB

	for i, v := range genericData.Where {
		if i == 0 {
			cols := fmt.Sprintf("%v %v ?", v.Column, v.Operator)
			fmt.Println(cols, v.Value)
			orCondCount = db.Where(cols, v.Value)
			continue
		}
		cols := fmt.Sprintf("%v %v ?", v.Column, v.Operator)
		fmt.Println(cols, v.Value)
		orCondCount = db.Or(cols, v.Value)
	}

	if orCondCount != nil {
		db = db.Where(orCondCount)
	}

	if tenantId != "" && genericData.Column != "" {
		clause := fmt.Sprintf("\"%v\".\"tenant_id\" = ? ", genericData.Column)
		db = db.Where(clause, tenantId)
	}

	for key, value := range genericData.Find {
		clause := fmt.Sprintf("\"%v\".\"%v\" = ? ", genericData.Column, key)
		db = db.Where(clause, value)
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

	var orCondRes *gorm.DB

	for i, v := range genericData.Where {
		if i == 0 {
			cols := fmt.Sprintf("%v %v ?", v.Column, v.Operator)
			orCondRes = db.Where(cols, v.Value)
			continue
		}
		cols := fmt.Sprintf("%v %v ?", v.Column, v.Operator)
		orCondRes = db.Or(cols, v.Value)
	}

	if orCondRes != nil {
		db = db.Where(orCondRes)
	}

	if tenantId != "" && genericData.Column != "" {
		clause := fmt.Sprintf("\"%v\".\"tenant_id\" = ? ", genericData.Column)
		db = db.Where(clause, tenantId)
	}

	if genericData.Limit != 0 {
		db = db.Limit(int(genericData.Limit)).Offset(int(genericData.Offset))
	}

	if genericData.OrderBy != "" {
		db = db.Order(genericData.OrderBy)
	}

	for key, value := range genericData.Find {
		clause := fmt.Sprintf("\"%v\".\"%v\" = ? ", genericData.Column, key)
		db = db.Where(clause, value)
	}

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

func UpsertHandler[T any](c *fiber.Ctx) error {
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
	// dbAdapter.DB.Debug()

	// Marshal the struct back to a JSON string
	if err := dbAdapter.DB.Clauses(clause.OnConflict{DoNothing: true}).Create(&genericData).Error; err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).SendString(fmt.Sprintf("%d", err))
	}

	newJSONData2, err := json.Marshal(genericData)
	if err != nil {
		fmt.Println("Error:", err)
	}
	// Print the JSON object and authorization header

	// Return a 200 OK response
	return c.Status(fiber.StatusOK).SendString(string(newJSONData2))

}
