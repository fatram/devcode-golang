package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/fatram/devcode-golang/config"
	"github.com/fatram/devcode-golang/controller/http"
	"github.com/fatram/devcode-golang/internal/connector"
)

func beforeTerminate() {
	fmt.Println("Good bye!")
}

// SetupCloseHandler creates a 'listener' on a new goroutine which will notify the
// program if it receives an interrupt from the OS. We then handle this by calling
// our clean up procedure and exiting the program.
func setupCloseHandler() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\r- Ctrl+C pressed in Terminal")
		beforeTerminate()
		os.Exit(0)
	}()
}

// @title Golang Test Documentation
// @version 1.0
// @description This is a API server for Simple Login Online

// @contact.name Fatur Rahman
// @contact.email frfatram@gmail.com

// @host localhost
// @BasePath /
func main() {
	setupCloseHandler()
	config.ReadConfig(".env")
	initializeDatabase()
	http.NewHttpController().Start("", config.Configuration().Port)
}

func initializeDatabase() {
	db := connector.LoadMysqlDatabase()
	sql := `
			DROP TABLE IF EXISTS activities;
		/*!40101 SET @saved_cs_client     = @@character_set_client */;
		/*!40101 SET character_set_client = utf8 */;
		CREATE TABLE activities (
		activity_id INT NOT NULL AUTO_INCREMENT,
		title varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
		email varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
		created_at int(11) DEFAULT NULL,
		updated_at int(11) DEFAULT NULL,
		PRIMARY KEY (activity_id)
		) AUTO_INCREMENT=1 ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

		DROP TABLE IF EXISTS todos;
		/*!40101 SET @saved_cs_client     = @@character_set_client */;
		/*!40101 SET character_set_client = utf8 */;
		CREATE TABLE todos (
		todo_id INT NOT NULL AUTO_INCREMENT,
		activity_group_id INT NOT NULL,
		title varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
		is_active tinyint(1) NOT NULL,
		priority varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
		created_at int(11) DEFAULT NULL,
		updated_at int(11) DEFAULT NULL,
		PRIMARY KEY (todo_id)
		) AUTO_INCREMENT=1 ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

	`
	_, err := db.Exec(sql)
	if err != nil {
		log.Panic(err)
	}
}
