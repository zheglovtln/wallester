package helpers

// Need for serverside rendering, now i use client-side rendering.
type DataTables struct {
	Draw int `json:"draw"`
	RecordsTotal int64 `json:"recordsTotal"`
	RecordsFiltered int64 `json:"recordsFiltered"`
}

func ProcessDataTables(count int64) DataTables {
	var DT DataTables
	DT.Draw = 1
	DT.RecordsFiltered = count
	DT.RecordsTotal = count
	return DT
}




