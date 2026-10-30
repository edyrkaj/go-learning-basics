package main

import service_mod "golearn.com/interfaces/service_mod"

func main() {
	emailSvc := service_mod.EmailService{
		SMTPServer: "smtp.example.com",
	}

	nm := service_mod.NotificationManager{Sender: emailSvc}
	nm.NotifyUser("edyrkaj@gmail.com", "Hi Eledi!")
}
