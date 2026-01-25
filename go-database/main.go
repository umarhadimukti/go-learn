package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

func initializeDatabase(ctx context.Context) (*pgxpool.Pool, error) {
	// DSN postgresql
	dsn := "postgres://postgres:12344321@localhost:5432/db-contact"

	// Create connection config
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}

	config.MaxConns = 10
	config.MinConns = 2
	config.MaxConnIdleTime = 5 * time.Minute
	config.MaxConnLifetime = 30 * time.Minute

	// Connection with pool
	db, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, err
	}

	// Tes PING
	if err := db.Ping(ctx); err != nil {
		return nil, err
	}
	fmt.Println("Connected using PostgreSQL (pgx)")

	return db, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func insertData(ctx context.Context, db *pgxpool.Pool) {
	hashedPassword, _ := hashPassword("test1234")
	_, err := db.Exec(
		ctx,
		"INSERT INTO users VALUES ($1, $2, $3)",
		"John Denver",
		"jdever@example.test",
		hashedPassword,
	)
	if err != nil {
		log.Fatal("Failed to execute insert statement:", err)
	}
	fmt.Println("Saved!")
}

type User struct {
	Username string
	Name string
	Password string
}

func getAllData(ctx context.Context, db *pgxpool.Pool) ([]User, error) {
	rows, err := db.Query(ctx, "SELECT username, name, password FROM users ORDER BY name ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Username, &user.Name, &user.Password); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func getDataByName(ctx context.Context, db *pgxpool.Pool, name string) (*User, error) {
	var user User
	err := db.QueryRow(
		ctx,
		`SELECT username, name, password FROM users WHERE name = $1`,
		name,
	).Scan(&user.Username, &user.Name, &user.Password)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func updateData(ctx context.Context, db *pgxpool.Pool, username string, newName string) (error) {
	cmd, err := db.Exec(
		ctx,
		`UPDATE users SET name = $1 WHERE username = $2`,
		newName,
		username,
	)
	if err != nil {
		return err
	}
	if cmd.RowsAffected() == 0 {
		return fmt.Errorf("No user updated")
	}
	return nil
}

func transactionData(ctx context.Context, db *pgxpool.Pool, users []User) error {
	tx, err := db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)
	for _, user := range users {
		_, err := db.Exec(
			ctx,
			"INSERT INTO users VALUES ($1, $2, $3)",
			user.Username,
			user.Name,
			user.Password,
		)
		if err != nil {
			return err
		}
	}
	return tx.Commit(ctx)
}

func main() {
	// Context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	// Initialize database
	db, err := initializeDatabase(ctx)
	if err != nil {
		log.Fatal("DB Unreachable:", err)
	}
	defer db.Close()

	// Insert data
	// insertData(ctx, db)

	// Get all data
	users, err := getAllData(ctx, db)
	if err != nil {
		log.Fatal("Failed to get all data:", err)
	}
	for i, u := range users {
		fmt.Println(fmt.Sprint(i+1) + ".", u.Name)
	}
	fmt.Println("==== GET BY NAME ====")

	// Get data by name
	user, err := getDataByName(ctx, db, "Sugio Takeda")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Name: " + user.Name, "\nUsername: " + user.Username)
	fmt.Println("==== UPDATE ====")

	// Update by username
	err = updateData(ctx, db, "ekok", "Eko Kurniawanx")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Data updated!")
	fmt.Println("==== INSERT MANY USING TX ====")
	
	// Insert many users using transaction
	usersList := []User{
		{Username: "ilhamgots", Name: "Ilham Gots", Password: "test1234"},
		{Username: "windabarusadar", Name: "Winda Barusadar", Password: "test1234"},
		{Username: "kurrrniawan", Name: "Kurrrniawan", Password: "test1234"},
	}
	err = transactionData(ctx, db, usersList)
	if err != nil {
		log.Fatal(err)
	}
}
