package bill

import "server/models"

func YearTotal(monthlyBill []models.Bill) (float64, error){
	var total float64
	for _, bill := range monthlyBill{
		total += bill.Price
	}
	return total, nil
}

func CreateNewBill(bills *[]models.Bill, yearTotal float64) *models.MonthlyYearlyBill{
	return &models.MonthlyYearlyBill{
		Bills: *bills,
		YearTotal: yearTotal,
	}
} 