package conditions

const (
	RequestInvalid = "INVALID"
	RequestRejected = "REJECTED"
	RequestVipAccess = "VIP_ACCESS"
	RequestStandardAccess = "STANDARD_ACCESS"
	RequestLimitedAccess = "LIMITED_ACCESS"

	GradeF = "F"
	GradeD = "D"
	GradeC = "C"
	GradeB = "B"
	GradeA = "A"
	ErrGradeInvalid = "INVALID"
)

func ClassifyRequest(age int, hasID bool, balance float64, isVIP bool) string {
	if age <= 0 || balance < 0 {
		return RequestInvalid
	} else if age < 18 || hasID == false {
		return RequestRejected
	} else if isVIP && balance >= 10000 {
		return RequestVipAccess
	} else if balance >= 1000 {
		return RequestStandardAccess
	}
	
	return RequestLimitedAccess
}

func EvaluateGrade(score int) string {
	switch {
	case score >= 0 && score <= 59:
		return GradeF
	case score >= 60 && score <= 69:
		return GradeD
	case score >= 70 && score <= 79:
		return GradeC
	case score >= 80 && score <= 89:
		return GradeB
	case score >= 90 && score <= 100:
		return GradeA
	}
	return ErrGradeInvalid
}
