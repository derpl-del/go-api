package emailfunc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/smtp"
	"strings"

	"github.com/derpl-del/go-api.git/gocode/endecode"
	"github.com/derpl-del/go-api.git/gocode/strcode"
	"github.com/derpl-del/go-api.git/gocode/utilfunc"
)

//SenderEnv var
var SenderEnv strcode.SenderMail

const smtphost = "smtp.gmail.com"
const smtpprot = 587

//GenerateEmailSender func
func GenerateEmailSender(w http.ResponseWriter, r *http.Request) {
	SenderEnv = strcode.SenderMail{}
	ReqBody, _ := ioutil.ReadAll(r.Body)
	var Request strcode.SenderMail
	json.Unmarshal(ReqBody, &Request)
	SenderEnv.Sender = Request.Sender
	SenderEnv.Password = Request.Password
}

var subject string
var context string
var message string

//GenerateEmail Func
func GenerateEmail(tomail string, input string, choose string) {
	to := []string{tomail}
	//cc := []string{""}
	if choose == "user verify" {
		subject = "User Verify"
		context = endecode.GenerateEn(input)
		message = "Please Verify Your User" + "\n" + "http://localhost:9000/api/v1/verifyuser?req=" + context
	} else if choose == "invoice" {
		subject = "invoice"
		usermap := utilfunc.TokenizeWithValue("amount|price|total", input)
		message = "Your Invoice Is :" + "\n" + "Price : " + usermap["price"] + "\n" + "Amount : " + usermap["amount"] + "\n" + "=================\n" + "Total : " + usermap["total"]
	}
	if SenderEnv.Sender != "" && SenderEnv.Password != "" {
		err := sendMail(to, subject, message)
		if err != nil {
			log.Fatal(err.Error())
		}

		log.Println("Mail sent!")
	} else {
		log.Println("Set Sender First")
	}
}

func sendMail(to []string, subject, message string) error {
	body := "From: " + SenderEnv.Sender + "\n" +
		"To: " + strings.Join(to, ",") + "\n" +
		"Subject: " + subject + "\n\n" +
		message

	auth := smtp.PlainAuth("", SenderEnv.Sender, SenderEnv.Password, smtphost)
	smtpAddr := fmt.Sprintf("%s:%d", smtphost, smtpprot)

	err := smtp.SendMail(smtpAddr, auth, SenderEnv.Sender, to, []byte(body))
	if err != nil {
		return err
	}

	return nil
}
