package builder

import (
	"fmt"
	"testing"
)

func TestBuilder(t *testing.T) {
	vegMeal := VegMeal()
	fmt.Println("veg meal item:")
	vegMeal.ShowItems()
	fmt.Println(fmt.Sprintf("veg meal，total cost：%f", vegMeal.getCost()))

	chickenMeal := ChickenMeal()
	fmt.Println("chicken meal item:")
	chickenMeal.ShowItems()
	fmt.Println(fmt.Sprintf("chicken meal，total cost：%f", chickenMeal.getCost()))
}
