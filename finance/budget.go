package finance

import (
	"context"
	"pontistudios/gogogo/graph/model"
)

type FinanceResolver struct{}

/**
 * CalculateGoal calculates the annual and monthly pre-tax income required to reach the user's goals.
 *
 * This function considers the user's current expenses, tax rate, and goals to determine the total amount of money needed to cover expenses and reach goals.
 *
 * It calculates the total amount of money required to cover the user's current expenses while they work toward their goals, and the total amount of money required to cover the user's goals.
 *
 * The function then determines the pre-tax income the user must earn to afford their goals, and finally, the total annual and monthly pre-tax income required to reach the user's goals.
 *
 * Parameters:
 * - budget: The user's current budget, including expenses and tax rate.
 * - years: The number of years the user has to reach their goals.
 * - goals: A slice of GoalInput representing the user's goals.
 *
 * Returns:
 * - annualPreTaxIncome: The total annual pre-tax income required to reach the user's goals.
 * - monthlyPreTaxIncome: The total monthly pre-tax income required to reach the user's goals.
 * - error: An error if any occurs during the calculation.
 */
func (s *FinanceResolver) CalculateGoal(ctx context.Context, budget model.BudgetInput, years int, goals []*model.GoalInput) (float64, float64, error) {
	months := years * 12
	// Create a map to store the total amount of money needed per month
	expensesPerMonth := 0.0

	// Calculate the total amount of money needed per month
	for _, expense := range budget.Expenses {
		if expense.Cadence == model.ExpenseCadenceMonthly {
			expensesPerMonth += expense.Amount
		}
	}

	/**
	 * This is the total amount of money required to cover the
	 * user's current expenses while they work toward their goals.
	 */
	totalExpenses := expensesPerMonth * float64(months)

	/**
	 * This is the total amount of money required to cover the
	 * user's current expenses while they work toward their goals.
	 */
	totalRequiredForGoals := 0.0
	for _, goal := range goals {
		totalRequiredForGoals += goal.Amount
	}

	/**
	 * This is the amount of pre-tax income the user must earn to
	 * afford their goals.
	 */
	preTaxIncome := (totalRequiredForGoals + totalExpenses) / (1 - budget.TaxRate)

	/**
	 * This is the total annual pre-tax income required to reach the user's goals.
	 */
	annualPreTaxIncome := preTaxIncome / 3

	/**
	 * This is the total monthly pre-tax income required to reach the user's goals.
	 */
	monthlyPreTaxIncome := preTaxIncome / 12

	return annualPreTaxIncome, monthlyPreTaxIncome, nil
}
