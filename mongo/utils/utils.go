package utils

import (
	"fmt"
	"os"

	"github.com/JesusRJ/golearning/mongo/model"
	"github.com/jedib0t/go-pretty/table"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetCompanyID() primitive.ObjectID {
	id, _ := primitive.ObjectIDFromHex("6632bff5465065406a8f320a")
	return id
}

func getCommaSeparated[T interface{ String() string }](values []T) (result string) {
	for _, v := range values {
		result = fmt.Sprintf("%s, %s", result, v)
	}
	if result != "" {
		return result[1:]
	}
	return result
}

func PrintUsers(users []model.User) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"User", "Address", "Phones", "Company", "Pets"})

	for _, u := range users {
		phones := getCommaSeparated(u.Phone)
		pets := getCommaSeparated(u.Pets)

		t.AppendRows([]table.Row{
			{u.Name, u.Address, phones, u.Company, pets},
		})
	}

	t.AppendFooter(table.Row{"", "", "", "Count", len(users)})
	t.Render()
}
