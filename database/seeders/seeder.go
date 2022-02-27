package seeders

import "goapihub/pkg/seed"

func Initialize()  {
	seed.SetRunOrder([]string{
		"SeedUsersTable",
	})
}