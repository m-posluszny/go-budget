package budgets

import "fmt"

type Category struct {
	Uid     string
	UserUid string
	Name    string
}

func GetCategories(userUid string) ([]Category, error) {
	fmt.Println(userUid)
	var cats []Category
	cats = append(cats, Category{"1", userUid, "Groceries"})
	cats = append(cats, Category{"2", userUid, "Home"})
	cats = append(cats, Category{"3", userUid, "Hobby"})
	return cats, nil
}
