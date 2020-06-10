package http

type Course struct {
	Course        float64 `json:"course"`
	CurrencyShort string  `json:"currencyShort"`
}

type CourseResponse struct {
	Symbol       string `json:"symbol"`
	SourceCourse Course `json:"sourceCourse"`
	TargetCourse Course `json:"targetCourse"`
}
