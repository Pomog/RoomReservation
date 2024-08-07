package main

import (
	"encoding/gob"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	"udemyCourse1/internal/config"
	"udemyCourse1/internal/driver"
	"udemyCourse1/internal/handlers"
	helpers "udemyCourse1/internal/helper"
	"udemyCourse1/internal/models"
	"udemyCourse1/internal/render"

	"github.com/alexedwards/scs/v2"
)

const port = ":8080"

var app config.AppConfig
var session *scs.SessionManager

var infolog *log.Logger
var errorlog *log.Logger

func main() {

	db, err := run()
	if err != nil {
		log.Fatal(err)
	}

	defer db.SQL.Close()
	defer close(app.MailChan)

	fmt.Println("Starting Mail listener")
	listenForMail()

	// msg := models.MailData{
	// 	To: "john@do.com",
	// 	From: "me@here.com",
	// 	Subject: "Test",
	// 	Content: "Hello",
	// }

	// app.MailChan <- msg

	fmt.Printf("Server starting on port %s\n", port)

	srv := &http.Server{
		Addr:    port,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}

func run() (*driver.DB, error) {
	fmt.Println("Starting app")

	// data in the session
	gob.Register(models.Reservation{})
	gob.Register(models.User{})
	gob.Register(models.Room{})
	gob.Register(models.Restriction{})
	gob.Register(map[string]int{})

	// read flags
	inProduction := flag.Bool("production", false, "Application is in production")
	useCashe := flag.Bool("cashe", false, "Use templates cache")
	dbHost := flag.String("dbhost", "localhost", "Database HOST")
	dbName := flag.String("dbname", "bookings", "Database Name")
	dbUser := flag.String("dbuser", "thoryur", "Database user name")
	dbPass := flag.String("dbpassword", "plot123123", "Database password")
	dbPort := flag.String("dbport", "5432", "Database PORT")
	dbSSL := flag.String("dbssl", "disable", "Database SSL (disable, prefer, require)")

	flag.Parse()
	if *dbName == "" || *dbUser==""{
		fmt.Println("Missing required flags")
		os.Exit(1)
	}

	mailChan := make(chan models.MailData)
	app.MailChan = mailChan

	// change this to true when in production
	app.InProduction = *inProduction
	app.UseCache = *useCashe

	// info log
	infolog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infolog

	// error log
	errorlog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorlog

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	// connect to database
	log.Println("Connecting to database...")
	connectionString := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s", *dbHost, *dbPort, *dbName, *dbUser, *dbPass, *dbSSL)
	db, err := driver.ConnectSQL(connectionString)
	if err != nil {
		log.Fatal("cannot connect to database! Dying...")
	}

	tc, err := render.CreateTemplateCashe()
	if err != nil {
		log.Fatal("cannot create template cache")
		return nil, err
	}
	app.TemplateCache = tc

	repo := handlers.NewRepo(&app, db)
	handlers.NewHandlers(repo)

	render.NewRenderer(&app)

	helpers.NewHelpers(&app)

	return db, nil
}

// test mail
/*
	from := "thoryur@gmail.com"
	password := ""
	to := "thoryur@gmail.com"
	subject := "Test Email"
	body := "Hello, this is a test email from Golang."

	msg := "To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" + body

	auth := smtp.PlainAuth("", from, password, "smtp.gmail.com")

	err = smtp.SendMail("smtp.gmail.com:587", auth, from, []string{to}, []byte(msg))
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Email sent successfully.")
	}
*/
