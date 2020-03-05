package chapter5

type IFCalcService interface {
	XXX(x int) int
	YYY(x, y int) int
}

type Calculator struct {
	service IFCalcService
}

func (c Calculator) Method(x, y, z int) int {
	result := c.service.XXX(x)
	result = c.service.XXX(result)
	result = c.service.YYY(result, z)
	return result + 10
}
