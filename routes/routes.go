package routes

import (
	"medcare/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.Static("/signin", "C:\\Users\\HP\\OneDrive\\Desktop\\HMS\\MedCare\\frontend\\signin")
r.Static("/home", "C:\\Users\\HP\\OneDrive\\Desktop\\HMS\\MedCare\\frontend\\home")
r.Static("/signup", "C:\\Users\\HP\\OneDrive\\Desktop\\HMS\\MedCare\\frontend\\signup")
r.Static("/doctorsignup", "C:\\Users\\HP\\OneDrive\\Desktop\\HMS\\MedCare\\frontend\\doctorsignup")
r.Static("/doctorsignin", "C:\\Users\\HP\\OneDrive\\Desktop\\HMS\\MedCare\\frontend\\doctorsignin")
r.Static("/doctorhome", "C:\\Users\\HP\\OneDrive\\Desktop\\HMS\\MedCare\\frontend\\doctorhome")
r.Static("/patienthome", "C:\\Users\\HP\\OneDrive\\Desktop\\HMS\\MedCare\\frontend\\patienthome")
r.Static("/patientcreateappointment", "C:\\Users\\HP\\OneDrive\\Desktop\\HMS\\MedCare\\frontend\\patientcreateappointment")
r.Static("/doctorenroll", "C:\\Users\\HP\\OneDrive\\Desktop\\HMS\\MedCare\\frontend\\doctorenroll")
r.Static("/listappointmentfordoctor", "C:\\Users\\HP\\OneDrive\\Desktop\\HMS\\MedCare\\frontend\\doctorlistappointment")
r.Static("/createprescriptionfordoctor", "C:\\Users\\HP\\OneDrive\\Desktop\\HMS\\MedCare\\frontend\\doctorprescription")
r.Static("/patientreport", "C:\\Users\\HP\\OneDrive\\Desktop\\HMS\\MedCare\\frontend\\patientreport")
r.Static("/doctorreport", "C:\\Users\\HP\\OneDrive\\Desktop\\HMS\\MedCare\\frontend\\doctorreport")
r.Static("/patientprescriptionlist", "C:\\Users\\HP\\OneDrive\\Desktop\\HMS\\MedCare\\frontend\\patientprescription")

	r.POST("/signupcustomer", handlers.CustomerSignup)
	r.POST("/signincustomer", handlers.CustomerSignIn)
	r.POST("/signupdoctor", handlers.DoctorSignup)
	r.POST("/signindoctor", handlers.DoctorSignin)
	r.POST("/bookappointment", handlers.BookAppointment)
	r.POST("/listappointments", handlers.ListAppointment)
	r.POST("/createprescription", handlers.CreatePrescription)
	r.POST("/listpatientreport", handlers.ListPatientReport)
	r.POST("/listpatientprescription", handlers.ListPrescription)
}
