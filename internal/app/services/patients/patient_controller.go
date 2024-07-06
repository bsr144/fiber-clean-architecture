package patients

type PatientController struct {
	PatientUsecase PatientUsecase
}

func NewPatientController(patientUsecase PatientUsecase) *PatientController {
	return &PatientController{
		PatientUsecase: patientUsecase,
	}
}
