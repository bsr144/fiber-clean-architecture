package patients

type patientUsecase struct {
	PatientRepository PatientRepository
	PatientFhirClient PatientFhirClient
}

func NewPatientUsecase(
	patientMongoRepository PatientRepository,
	patientFhirClient PatientFhirClient,
) PatientUsecase {
	return &patientUsecase{
		PatientRepository: patientMongoRepository,
		PatientFhirClient: patientFhirClient,
	}
}
