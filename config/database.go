package config

import (
	"os"
)

func init() {
	Configuration["database"] = map[string]interface{}{
		/*
			|--------------------------------------------------------------------------
			| Default Database Connection Name
			|--------------------------------------------------------------------------
			|
			| Here you may specify which of the database connections below you wish
			| to use as your default connection for all database work. Of course
			| you may use many connections at once using the Database library.
			|
		*/
		"default": os.Getenv("DB_CONNECTION"),

		/*
			|--------------------------------------------------------------------------
			| Database Connections
			|--------------------------------------------------------------------------
			|
			| Here are each of the database connections setup for your application.
			| Of course, examples of configuring each database platform that is
			| supported by Laravel is shown below to make development simple.
			|
			|
			| All database work in Laravel is done through the PHP PDO facilities
			| so make sure you have the driver for your particular database of
			| choice installed on your machine before you begin development.
			|
		*/
		"connections": map[string]interface{}{
			"pgsql": map[string]interface{}{
				"driver":   "pgsql",
				"host":     os.Getenv("DB_HOST"),
				"port":     os.Getenv("DB_PORT"),
				"database": os.Getenv("DB_DATABASE"),
				"username": os.Getenv("DB_USERNAME"),
				"password": os.Getenv("DB_PASSWORD"),
				"charset":  "utf8",
				"prefix":   "",
				"schema":   "public",
				"sslmode":  "prefer",
			},
		},
	}
}
