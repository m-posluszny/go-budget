package accounts

type BudgetType string

const BudgetTypeChecking BudgetType = "checking"
const BudgetTypeSavings BudgetType = "savings"
const BudgetTypeInvestments BudgetType = "investments"
const BudgetTypeOffBudget BudgetType = "off-budget"

var BudgetTypes = []BudgetType{
	BudgetTypeChecking,
	BudgetTypeSavings,
	BudgetTypeInvestments,
	BudgetTypeOffBudget,
}
