package error_code

// import "fmt"

// func ErrorNoError() error {
// 	return nil
// }
// func ErrorGeneralError() error {
// 	return fmt.Errorf("system: General app error.")
// }

// // Failed to query databse
// func ErrorQueryDBError() error {
// 	return fmt.Errorf("system: Database query error.")
// }

// // Web APP (GRPC) Failed to set web session
// func ErrorSetWebSession() error {
// 	return fmt.Errorf("webApp: Failed to set session.")
// }

// // Web App (GRPC) Failed to login for bad account/password
// func ErrorWebLoginBadAcc() error {
// 	return fmt.Errorf("webApp: Failed to login for bad account/password.")
// }

// // Web App (GRPC) Failed to login for email and or phone number is not verified yet.
// func ErrorWebLoginNotVerifiedAcc() error {
// 	return fmt.Errorf("webApp: Failed to login for email and or phone number is not verified.")
// }

// // Web App (GRPC) return invalid session
// func ErrorWebInvalidSession() error {
// 	return fmt.Errorf("webApp: Invalid session. Do logout ASAP.")
// }

// func ErrorWebRegisterAlreadyRegistered() error {
// 	return fmt.Errorf("webApp: Email already registered. Please login.")
// }

const (
	ErrorSuccess = "000"
	ErrorUnverifiedAccount = "101"
	ErrorAccountAlreadyExists = "102"
	ErrorAuthStateNotExist = "103"
	ErrorBadAccount = "106"
	ErrorBadPassword = "107"
	ErrorWebInvalidSession = "704"
	ErrorServer = "900"
	ErrorInvalidParameter = "902"
	ErrorSendEmail = "903"
	ErrorNoOTP = "904"
	ErrorBadOTP = "905"
)