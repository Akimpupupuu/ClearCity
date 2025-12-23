package mailer

import (
	"fmt"

	"github.com/Akimpupupuu/ClearCity/internal/types"
	"gopkg.in/gomail.v2"
)

func Send(report *types.Report) error {
	reportString := fmt.Sprintf("id: %d\ntitle: %s\naddress: %s\ndescription: %s\ncreated at: %s",
		report.Id, report.Title, report.Address, report.Description, report.CreatedAt.Format("15:04:05 02-01-2006"))

	m := gomail.NewMessage()
	m.SetHeader("From", "damirmagdeev1054@gmail.com")
	m.SetHeader("To", "damirmagdeev1054@gmail.com")
	m.SetHeader("New report!")
	m.SetBody("text/html", reportString)

	d := gomail.NewDialer("smtp.gmail.com", 587, "damirmagdeev1054@gmail.com", "cbcezpfsoylkobvr")

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
