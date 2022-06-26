package client

import (
	"net/url"
	"time"
)

type Hostitals []struct {
	Edrpou string `json:"edrpou"`
	Region string `json:"region"`
	Name   string `json:"name"`
	Stage  string `json:"stage"`
}

type BasicCovid struct {
	Edrpou     string                 `json:"EDRPOU"`
	ReportDate string                 `json:"report_date"`
	Updated    string                 `json:"updated"`
	Data       map[string]interface{} `json:"data"`
}

type Protectives struct {
	basicProtectives BasicCovid
}

type Equipment struct {
	basicEquipment BasicCovid
}

type Medicines struct {
	basicProtectives BasicCovid
}

type Beds struct {
	basicMedicines BasicCovid
}

type Patients struct {
	basicPatients BasicCovid
}

type Staff struct {
	basicStaff BasicCovid
}

type CovidArgs struct {
	Edrpou     string
	ReportData time.Time
}

func (args CovidArgs) QueryParams() url.Values {
	q := make(url.Values)

	if args.Edrpou != "" {
		q.Add("edrpou", args.Edrpou)
	}

	if !args.ReportData.IsZero() {
		q.Add("report_date", args.ReportData.Format("2006-01-02"))
	}

	return q
}
