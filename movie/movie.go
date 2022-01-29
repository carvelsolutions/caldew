package movie

type Movie struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func FindAll() ([]Movie, error) {
	var testData = []Movie{
		{ID: 1, Name: "Avengers"},
		{ID: 2, Name: "Ant-Man"},
		{ID: 3, Name: "Thor"},
		{ID: 4, Name: "Captain America"},
		{ID: 5, Name: "A Dogs Life"},
	}
	return testData, nil
}
