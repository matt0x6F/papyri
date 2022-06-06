package main

import (
	"flag"
	"os"
	"strconv"

	"github.com/adrg/xdg"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/mattouille/papyri/database"
)

func main() {
	log.Debug().Strs("args", os.Args).Msg("Arguments passed")

	// Close the default connections
	defer database.DefaultConnection.Close()
	err := database.DefaultConnection.Initialize()
	if err != nil {
		println("Error:", err.Error())

		os.Exit(1)
	}

	cfg, err := database.DefaultConnection.Config()
	if err != nil {
		println("Error:", err.Error())

		os.Exit(1)
	}

	log.Logger.Info().Interface("config", cfg)

	// Create application with options
	err = RunGUI()
	if err != nil {
		println("Error:", err)
	}
}

func init() {
	// sort out flags
	debug := flag.Bool("debug", false, "sets log level to debug")

	flag.Parse()

	// creates pretty print text logging
	zerolog.CallerMarshalFunc = func(file string, line int) string {
		short := file
		for i := len(file) - 1; i > 0; i-- {
			if file[i] == '/' {
				short = file[i+1:]
				break
			}
		}
		file = short
		return file + ":" + strconv.Itoa(line)
	}

	// Set two log modes
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		log.Debug().Str("method", "flag").Msg("Log level set to debug")
	}

	// print the line that triggered the log
	log.Logger = log.With().Caller().Logger()

	// write to stderr
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	// setup the database
	dbDir := xdg.DataHome
	db, _ := database.New(dbDir + "/" + database.DefaultName)
	database.DefaultConnection = db

	log.Info().Msg("Starting application")
}
