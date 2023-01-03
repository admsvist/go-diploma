package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/admsvist/go-diploma/country_codes"
	"github.com/admsvist/go-diploma/internal/pkg/repository"
	"github.com/admsvist/go-diploma/pkg/filereader"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

const countryCodesPath = "./codes.json"
const smsDataPath = "./../simulator/sms.data"
const voiceCallDataPath = "./../simulator/voice.data"
const emailDataPath = "./../simulator/email.data"
const billingDataPath = "./../simulator/billing.data"
const mmsUrl = "http://127.0.0.1:8383/mms"
const supportUrl = "http://127.0.0.1:8383/support"
const incidentUrl = "http://127.0.0.1:8383/accendent"

func main() {
	reader := filereader.New()

	country_codes.Init(reader, countryCodesPath)

	smsDataRepository := repository.NewSMSDataRepository()
	smsDataRepository.LoadData(reader, smsDataPath)

	mmsDataRepository := repository.NewMMSDataRepository()
	mmsDataRepository.LoadData(mmsUrl)

	voiceCallDataRepository := repository.NewVoiceCallDataRepository()
	voiceCallDataRepository.LoadData(reader, voiceCallDataPath)

	emailDataRepository := repository.NewEmailDataRepository()
	emailDataRepository.LoadData(reader, emailDataPath)

	billingDataRepository := repository.NewBillingDataRepository()
	billingDataRepository.LoadData(reader, billingDataPath)

	supportDataRepository := repository.NewSupportDataRepository()
	supportDataRepository.LoadData(supportUrl)

	incidentDataRepository := repository.NewIncidentDataRepository()
	incidentDataRepository.LoadData(incidentUrl)

	//fmt.Println(smsDataRepository.Data)
	//fmt.Println(mmsDataRepository.Data)
	//fmt.Println(voiceCallDataRepository.Data)
	//fmt.Println(emailDataRepository.Data)
	//fmt.Printf("%+v\n", billingDataRepository.Data[0])
	//fmt.Println(supportDataRepository.Data)
	fmt.Println(incidentDataRepository.Data)
}

func qmain() {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	r := mux.NewRouter()
	r.HandleFunc("/", handleConnection)

	srv := &http.Server{
		Addr: "127.0.0.1:8282",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r, // Pass our instance of gorilla/mux in.
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}

func handleConnection(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
