package handler

import (
	"bufio"
	"encoding/base64"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jung-kurt/gofpdf"
	"github.com/sahilgarg96/DBTNT/logging"
	"github.com/sahilgarg96/DBTNT/redis"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Response struct {
	Success bool   `json:"success"`
	Output  string `json:"error"`
}

var Logger = logging.NewLogger()

// generate pdf and email
func GeneratePdf(w http.ResponseWriter, r *http.Request) {

	if r.Header.Get("Content-Type") != "application/json" {
		msg := "Content-Type header is not application/json"
		http.Error(w, msg, http.StatusUnsupportedMediaType)
		return
	}

	vars := mux.Vars(r)
	userId := vars["user_id"]

	expiryTime := time.Duration(301 * 1000 * 1000 * 1000) // 5 mins 1 second
	currentTs := time.Now().Unix()

	if !redis.SetValue("pdftsuser_"+userId, strconv.FormatInt(currentTs, 10), expiryTime) {
		Logger.Errorf("error setting timestamp for userId " + userId)
	}

	var pdfOutput map[string]interface{}
	var resp Response
	//decode json
	err := json.NewDecoder(r.Body).Decode(&pdfOutput)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.Image("static/db.jpeg", -5, 5, 10, 0, false, "", 0, "")
	pdf.SetFont("Arial", "B", 12)
	count := 0
	for _, value := range pdfOutput {
		word := value.(string)
		wo := ""
		var st []string
		for i := range word {
			if i == 0 {
				wo = strconv.Itoa(count+1) + "."
			}
			wo = wo + string(word[i])
			if len((wo)) > 60 {
				st = append(st, wo)
				wo = ""
			}
		}
		st = append(st, wo)
		for i := range st {
			pdf.Cell(float64(1), 10, st[i])
			pdf.Ln(6)

		}
		pdf.Ln(12)
		count++
	}

	err = pdf.OutputFileAndClose("static/pdfs/" + userId + "_" + strconv.FormatInt(currentTs, 10) + ".pdf")

	if err != nil {
		resp.Output = err.Error()
		Logger.Errorf("error generating pdf for userId " + userId)
	}

	if resp.Output != "" {
		resp.Success = false
	} else {
		resp.Success = true
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(&resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

//uses sendgrid and send email
func SendEmail(fromMail string, toMail string, fileName string) error {

	from := mail.NewEmail("DoubtNut", fromMail)

	subject := "Similar Questions related to Search"

	to := mail.NewEmail("User Name", toMail)

	plainTextContent := "Please find below the relevant content"
	htmlContent := "<strong>List</strong>"

	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)

	pdf := mail.NewAttachment()
	fileHandle, err := os.Open("static/pdfs/" + fileName)
	defer fileHandle.Close()
	if err != nil {
		return err
	}
	reader := bufio.NewReader(fileHandle)
	content, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}
	// Encode as base64.
	encoded := base64.StdEncoding.EncodeToString(content)
	pdf.SetContent(encoded)
	pdf.SetType("application/pdf")
	pdf.SetFilename("questions.pdf")
	pdf.SetDisposition("attachment")

	//message.AddAttachment(fileHandle)
	message.AddAttachment(pdf)
	client := sendgrid.NewSendClient("Sendgrid_api_key")
	_, err = client.Send(message)
	if err != nil {
		return err
	}

	return nil

}
