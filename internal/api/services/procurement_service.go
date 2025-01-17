package services

import (
	"fmt"
	"web-api/internal/pkg/database"
	"web-api/internal/pkg/models/request"
	"web-api/internal/pkg/models/types"
)

type ProcurementService struct {
	*BaseService
}

var Procurement = &ProcurementService{}

func (s *ProcurementService) GetDataService(requestParams *request.ProcurementRequest) ([]types.YWCP, error) {
	var YWCPs []types.YWCP

	// Kết nối cơ sở dữ liệu
	db, err := database.RPConnection()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return nil, err
	}
	dbInstance, _ := db.DB()
	defer dbInstance.Close() // Đóng kết nối cơ sở dữ liệu khi xong

	// Truy vấn với tham số
	query := `
		SELECT *, DATEDIFF(DAY, Startdate, EndDate) AS DiffDay
		FROM (
			SELECT DDBH,
				   (SELECT TOP 1 USERDATE FROM KCRKS WHERE CGBH = ddzl.DDBH ORDER BY USERDATE) AS StartDate,
				   (SELECT TOP 1 INDATE FROM YWCP WHERE DDBH = ddzl.DDBH ORDER BY INDATE DESC) AS EndDate
			FROM (
				SELECT *,
					   (SELECT SUM(QTY) FROM YWCP WHERE DDBH = ddzl.DDBH) AS OKQTY
				FROM (
					SELECT DDBH, Pairs
					FROM ddzl
					WHERE LEFT(BUYNO, 6) = ? AND GSBH = ? AND pairs > 1
				) ddzl
			) ddzl
			WHERE pairs = OKQTY
		) YWCP1
		ORDER BY DiffDay
	`

	// Truyền tham số
	err = db.Raw(query, requestParams.BUYNO, requestParams.GSBH).Scan(&YWCPs).Error
	if err != nil {
		fmt.Println("Query error:", err)
		return nil, err
	}

	return YWCPs, nil
}

func (s *ProcurementService) GetAverageService(requestParams *request.ProcurementRequest) (float32, error) {

	// Kết nối cơ sở dữ liệu
	db, err := database.RPConnection()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return 0, err
	}
	dbInstance, _ := db.DB()
	defer dbInstance.Close() // Đóng kết nối cơ sở dữ liệu khi xong

	// Truy vấn với tham số
	query := `
		SELECT CAST(AVG(CAST(s.DiffDay AS FLOAT)) AS DECIMAL(10, 4)) AS Average
		FROM (
			SELECT *, DATEDIFF(DAY, Startdate, EndDate) AS DiffDay
			FROM (
				SELECT DDBH,
					   (SELECT TOP 1 USERDATE FROM KCRKS WHERE CGBH = ddzl.DDBH ORDER BY USERDATE) AS StartDate,
					   (SELECT TOP 1 INDATE FROM YWCP WHERE DDBH = ddzl.DDBH ORDER BY INDATE DESC) AS EndDate
				FROM (
					SELECT *, (SELECT SUM(QTY) FROM YWCP WHERE DDBH = ddzl.DDBH) AS OKQTY
					FROM (
						SELECT DDBH, Pairs
						FROM ddzl
						WHERE LEFT(BUYNO, 6) = ? AND GSBH = ? AND Pairs > 1
					) ddzl
				) ddzl
				WHERE Pairs = OKQTY
			) YWCP1
		) AS s;
	`
	var average float32
	// Truyền tham số
	err = db.Raw(query, requestParams.BUYNO, requestParams.GSBH).Scan(&average).Error
	if err != nil {
		fmt.Println("Query error:", err)
		return 0, err
	}

	return average, nil
}
