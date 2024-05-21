package database

import (
	"database/sql"
	"errors"
	"log/slog"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func (db *Database) GetAllGroups(countUsersInSubgroups bool) ([]GroupWithUserCount, error) {
	slog.Debug("getting all groups", "countUsersInSubgroups", countUsersInSubgroups)

	var sql string
	if countUsersInSubgroups {
		sql = `
			SELECT
				"group".*,
				(
					WITH RECURSIVE "group_subset" AS (
						SELECT g."id" FROM "group" g WHERE g."id" = "group"."id" UNION ALL
						SELECT g."id" FROM "group" g JOIN "group_subset" gs ON g."parent_group" = gs."id"
					)
					SELECT count(u.*) FROM "user" u JOIN "group_subset" gs ON u."group" = gs."id"
				) "user_count"
			FROM "group"
		`
	} else {
		sql = `
			SELECT
				"group".*,
				(SELECT count(*) FROM "user" u WHERE u."group" = "group"."id") "user_count"
			FROM "group"
		`
	}

	groups := []GroupWithUserCount{}
	err := db.sqlx.Select(&groups, sql)
	if err != nil {
		return groups, err
	}

	return groups, nil
}

// ---

func (db *Database) GetGroup(groupId Id, includeSubgroups bool) (Group, []User, error) {
	slog.Debug("getting group", "groupId", groupId, "includeSubgroups", includeSubgroups)

	group := Group{}
	users := []User{}

	tx, err := db.sqlx.Beginx()
	if err != nil {
		return group, users, err
	}
	defer tx.Rollback()

	err = tx.Get(&group, `SELECT * FROM "group" WHERE "id" = $1`, groupId)
	if err != nil {
		if err == sql.ErrNoRows {
			return group, users, ErrGroupDoesntExist
		}

		return group, users, err
	}

	var sql string
	if includeSubgroups {
		sql = `
			WITH RECURSIVE "group_subset" AS (
				SELECT g."id" FROM "group" g WHERE g."id" = $1 UNION ALL
				SELECT g."id" FROM "group" g JOIN "group_subset" gs ON g."parent_group" = gs."id"
			)
			SELECT u.* FROM "user" u JOIN "group_subset" gs ON u."group" = gs."id"
		`
	} else {
		sql = `SELECT * FROM "user" WHERE "group" = $1`
	}

	err = tx.Select(&users, sql, groupId)
	if err != nil {
		return group, users, err
	}

	err = tx.Commit()
	if err != nil {
		return group, users, err
	}

	return group, users, nil
}

func (db *Database) CreateGroup(group NewGroup) (Id, error) {
	slog.Debug("creating group", "group", group)

	row := db.sqlx.QueryRow(`
		INSERT INTO "group" ("parent_group", "name")
		VALUES ($1, $2)
		RETURNING "id"
	`, group.ParentGroup, group.Name)

	var id Id
	err := row.Scan(&id)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == pgerrcode.ForeignKeyViolation {
				return 0, ErrParentGroupDoesntExist
			}
		}

		return 0, err
	}

	return id, nil
}

func (db *Database) UpdateGroup(groupId Id, group NewGroup) error {
	slog.Debug("updating group", "groupId", groupId, "group", group)

	res, err := db.sqlx.Exec(`
		UPDATE "group"
		SET "parent_group" = $1, "name" = $2
		WHERE "id" = $3
	`, group.ParentGroup, group.Name, groupId)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == pgerrcode.ForeignKeyViolation {
				return ErrParentGroupDoesntExist
			}
		}

		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return ErrGroupDoesntExist
	}

	return nil
}

func (db *Database) DeleteGroup(groupId Id) error {
	slog.Debug("deleting group", "groupId", groupId)

	res, err := db.sqlx.Exec(`DELETE FROM "group" WHERE id = $1`, groupId)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return ErrGroupDoesntExist
	}

	return nil
}

// ---

func (db *Database) GetUser(userId Id) (User, error) {
	slog.Debug("getting user", "userId", userId)

	user := User{}
	err := db.sqlx.Get(&user, `SELECT * FROM "user" WHERE "id" = $1`, userId)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, ErrUserDoesntExist
		}

		return user, err
	}

	return user, nil
}

func (db *Database) CreateUser(user NewUser) (Id, error) {
	slog.Debug("creating user", "user", user)

	row := db.sqlx.QueryRow(`
		INSERT INTO "user" ("group", "first_name", "last_name", "birth_year")
		VALUES ($1, $2, $3, $4)
		RETURNING "id"
	`, user.Group, user.FirstName, user.LastName, user.BirthYear)

	var id Id
	err := row.Scan(&id)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == pgerrcode.ForeignKeyViolation {
				return 0, ErrGroupDoesntExist
			}
		}

		return 0, err
	}

	return id, nil
}

func (db *Database) UpdateUser(userId Id, user NewUser) error {
	slog.Debug("updating user", "userId", userId, "user", user)

	res, err := db.sqlx.Exec(`
		UPDATE "user"
		SET "group" = $1, "first_name" = $2, "last_name" = $3, "birth_year" = $4
		WHERE "id" = $5
	`, user.Group, user.FirstName, user.LastName, user.BirthYear, userId)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == pgerrcode.ForeignKeyViolation {
				return ErrGroupDoesntExist
			}
		}

		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return ErrUserDoesntExist
	}

	return nil
}

func (db *Database) DeleteUser(userId Id) error {
	slog.Debug("deleting user", "userId", userId)

	res, err := db.sqlx.Exec(`DELETE FROM "user" WHERE id = $1`, userId)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return ErrUserDoesntExist
	}

	return nil
}
