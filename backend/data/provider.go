package data

// view everything about the patient
// create appointments
// summary of suggestions
// create documentation
// send messages
// upload things?

type Provider struct {
	Id        int
	IsAdmin   bool
	VidyoRoom string
}

//type Credential struct {
//	RD
//	MD
//	DO
//	DC
//	PhD
//}
//
//func (credential *Credential) ToString(c Credential) (s string) {
//	switch c {
//	case RD:
//		s = "RD"
//	case MD:
//		s = "MD"
//	case DO:
//		s = "DO"
//	case DC:
//		s = "DC"
//	case PhD:
//		s = "PhD"
//	}
//
//	return
//}
