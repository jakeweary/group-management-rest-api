package database

func (db *Database) SeedWithFakeData() {
	users, _ := db.CreateGroup(NewGroup{Name: "users"})
	mods, _ := db.CreateGroup(NewGroup{Name: "mods", ParentGroup: &users})
	admins, _ := db.CreateGroup(NewGroup{Name: "admins", ParentGroup: &mods})

	_, _ = db.CreateUser(NewUser{Group: users, FirstName: "User", LastName: "One", BirthYear: 2000})
	_, _ = db.CreateUser(NewUser{Group: users, FirstName: "User", LastName: "Two", BirthYear: 2000})
	_, _ = db.CreateUser(NewUser{Group: users, FirstName: "User", LastName: "Three", BirthYear: 2000})

	_, _ = db.CreateUser(NewUser{Group: mods, FirstName: "Mod", LastName: "One", BirthYear: 2000})
	_, _ = db.CreateUser(NewUser{Group: mods, FirstName: "Mod", LastName: "Two", BirthYear: 2000})
	_, _ = db.CreateUser(NewUser{Group: mods, FirstName: "Mod", LastName: "Three", BirthYear: 2000})

	_, _ = db.CreateUser(NewUser{Group: admins, FirstName: "Admin", LastName: "One", BirthYear: 2000})
	_, _ = db.CreateUser(NewUser{Group: admins, FirstName: "Admin", LastName: "Two", BirthYear: 2000})
	_, _ = db.CreateUser(NewUser{Group: admins, FirstName: "Admin", LastName: "Three", BirthYear: 2000})
}
