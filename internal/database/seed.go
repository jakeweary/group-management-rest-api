package database

import "log/slog"

func (db *Database) SeedWithFakeData() error {
	slog.Info("seeding db with fake data")

	// groups

	users, err := db.CreateGroup(NewGroup{Name: "users"})
	if err != nil {
		return err
	}
	mods, err := db.CreateGroup(NewGroup{Name: "mods", ParentGroup: &users})
	if err != nil {
		return err
	}
	admins, err := db.CreateGroup(NewGroup{Name: "admins", ParentGroup: &mods})
	if err != nil {
		return err
	}

	// group: users

	_, err = db.CreateUser(NewUser{Group: users, FirstName: "User", LastName: "One", BirthYear: 2000})
	if err != nil {
		return err
	}
	_, err = db.CreateUser(NewUser{Group: users, FirstName: "User", LastName: "Two", BirthYear: 2000})
	if err != nil {
		return err
	}
	_, err = db.CreateUser(NewUser{Group: users, FirstName: "User", LastName: "Three", BirthYear: 2000})
	if err != nil {
		return err
	}

	// group: mods

	_, err = db.CreateUser(NewUser{Group: mods, FirstName: "Mod", LastName: "One", BirthYear: 2000})
	if err != nil {
		return err
	}
	_, err = db.CreateUser(NewUser{Group: mods, FirstName: "Mod", LastName: "Two", BirthYear: 2000})
	if err != nil {
		return err
	}
	_, err = db.CreateUser(NewUser{Group: mods, FirstName: "Mod", LastName: "Three", BirthYear: 2000})
	if err != nil {
		return err
	}

	// group: admins

	_, err = db.CreateUser(NewUser{Group: admins, FirstName: "Admin", LastName: "One", BirthYear: 2000})
	if err != nil {
		return err
	}
	_, err = db.CreateUser(NewUser{Group: admins, FirstName: "Admin", LastName: "Two", BirthYear: 2000})
	if err != nil {
		return err
	}
	_, err = db.CreateUser(NewUser{Group: admins, FirstName: "Admin", LastName: "Three", BirthYear: 2000})
	if err != nil {
		return err
	}

	return nil
}
