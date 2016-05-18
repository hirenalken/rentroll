package rlib

import (
	"fmt"
	"strings"
	"time"
)

// the CSV file format:
//    0         1             2     3       4             5             6          7                8
// Designation,RentableName, ASMTID,Amount, Start,        Stop,         Frequency, ProrationMethod, AcctRule
// REH,         "101",       1,     1000.00,"2014-07-01", "2015-11-08", 6,         4,               "d ${DFLTGENRCV} _, c ${DFLTGSRENT} ${UMR}, d ${DFLTLTL} ${UMR} _ -"
// REH,         "101",       1,     1200.00,"2015-11-21", "2016-11-21", 6,         4,               "d ${DFLTGENRCV} _, c ${DFLTGSRENT} ${UMR}, d ${DFLTLTL} ${UMR} ${aval(${DFLTGENRCV})} -"

// type Assessment struct {
// 	ASMID           int64     // unique id for this assessment
// 	BID             int64     // what business
// 	RID             int64     // the rentable
// 	ASMTID          int64     // what type of assessment
// 	RAID            int64     // associated Rental Agreement
// 	Amount          float64   // how much
// 	Start           time.Time // start time
// 	Stop            time.Time // stop time, may be the same as start time or later
// 	Frequency       int64     // 0 = one time only, 1 = secondly, 2 = minutely, 3 = hourly, 4 = daily, 5 = weekly, 6 = monthly, 7 = quarterly, 8 = yearly
// 	ProrationMethod int64     // 0 = one time only, 1 = secondly, 2 = minutely, 3 = hourly, 4 = daily, 5 = weekly, 6 = monthly, 7 = quarterly, 8 = yearly
// 	AcctRule        string    // expression showing how to account for the amount
// 	Comment         string
// 	LastModTime     time.Time
// 	LastModBy       int64
// }

// CreateAssessmentsFromCSV reads an assessment type string array and creates a database record for the assessment type
func CreateAssessmentsFromCSV(sa []string, lineno int, AsmtTypes *map[int64]AssessmentType) {
	var a Assessment
	var r Rentable
	var err error
	des := strings.ToLower(strings.TrimSpace(sa[0]))
	if des == "designation" {
		return // this is just the column heading
	}

	//-------------------------------------------------------------------
	// Make sure the business is in the database
	//-------------------------------------------------------------------
	if len(des) > 0 {
		b1, _ := GetBusinessByDesignation(des)
		if len(b1.Designation) == 0 {
			Ulog("CreateAssessmentsFromCSV: line %d - business with designation %s does net exist\n", lineno, sa[0])
			return
		}
		a.BID = b1.BID
	}

	//-------------------------------------------------------------------
	// Find and set the rentable
	//-------------------------------------------------------------------
	s := strings.TrimSpace(sa[1])
	if len(s) > 0 {
		r, err = GetRentableByName(s, a.BID)
		if err != nil {
			fmt.Printf("CreateAssessmentsFromCSV: line %d - Error loading rentable named: %s.  Error = %v\n", lineno, s, err)
			return
		}
		a.RID = r.RID
	}

	//-------------------------------------------------------------------
	// Get the dates
	//-------------------------------------------------------------------
	DtStart, err := time.Parse(RRDATEINPFMT, strings.TrimSpace(sa[4]))
	if err != nil {
		fmt.Printf("CreateAssessmentsFromCSV: line %d - invalid start date:  %s\n", lineno, sa[4])
		return
	}
	a.Start = DtStart

	DtStop, err := time.Parse(RRDATEINPFMT, strings.TrimSpace(sa[5]))
	if err != nil {
		fmt.Printf("CreateAssessmentsFromCSV: line %d - invalid stop date:  %s\n", lineno, sa[5])
		return
	}
	a.Stop = DtStop

	//-------------------------------------------------------------------
	// Find and set the rental agreement -- but we only need to worry about
	// this value if the rentable state is normal.  We can skip it otherwise
	// because the other state values mean that it is not covered by a rental
	// agreement
	//-------------------------------------------------------------------
	if r.State == RENTABLESTATEONLINE {
		ra, err := FindAgreementByRentable(a.RID, &DtStart, &DtStop)
		if err != nil {
			fmt.Printf("CreateAssessmentsFromCSV: line %d - Error finding rental agreement for rentable %s.  Error = %v\n", lineno, r.Name, err)
			return
		}
		a.RAID = ra.RAID
	}

	//-------------------------------------------------------------------
	// Determine the amount
	//-------------------------------------------------------------------
	a.Amount = FloatFromString(sa[3], "Amount is invalid")

	//-------------------------------------------------------------------
	// Assessment Type
	//-------------------------------------------------------------------
	a.ASMTID = IntFromString(sa[2], "Assessment type is invalid")
	_, ok := (*AsmtTypes)[a.ASMTID]
	if !ok {
		fmt.Printf("CreateAssessmentsFromCSV: line %d - Assessment type is invalid: %s\n", lineno, sa[2])
		return
	}

	//-------------------------------------------------------------------
	// Frequency
	//-------------------------------------------------------------------
	a.Frequency = IntFromString(sa[6], "Frequency value is invalid")
	if a.Frequency < 0 || a.Frequency > OCCTYPEYEARLY {
		fmt.Printf("CreateAssessmentsFromCSV: line %d - Frequency must be between %d and %d.  Found %d\n", lineno, OCCTYPESECONDLY, OCCTYPEYEARLY, a.Frequency)
		return
	}

	//-------------------------------------------------------------------
	// Proration
	//-------------------------------------------------------------------
	a.ProrationMethod = IntFromString(sa[7], "Proration value is invalid")
	if a.ProrationMethod < 0 || a.ProrationMethod > OCCTYPEYEARLY {
		fmt.Printf("CreateAssessmentsFromCSV: line %d - Proration must be between %d and %d.  Found %d\n", lineno, OCCTYPESECONDLY, OCCTYPEYEARLY, a.ProrationMethod)
		return
	}
	if a.ProrationMethod > a.Frequency {
		fmt.Printf("CreateAssessmentsFromCSV: line %d - Proration granularity (%d) must be more frequent than the Frequency (%d)\n", lineno, a.ProrationMethod, a.Frequency)
		return
	}

	//-------------------------------------------------------------------
	// Set the AcctRule.  No checking for now...
	//-------------------------------------------------------------------
	a.AcctRule = sa[8]

	//-------------------------------------------------------------------
	// Make sure everything that needs to be set actually got set...
	//-------------------------------------------------------------------
	if len(a.AcctRule) == 0 || (a.RAID == 0 && r.State == RENTABLESTATEONLINE) ||
		a.Amount == 0 || a.RID == 0 || a.ASMTID == 0 || a.BID == 0 {
		fmt.Printf("Skipping this record\n")
		return
	}

	err = InsertAssessment(&a)
	if err != nil {
		fmt.Printf("CreateAssessmentsFromCSV: line %d - error inserting assessment: %v\n", lineno, err)
	}

}

// LoadAssessmentsCSV loads a csv file with a chart of accounts and creates ledger markers for each
func LoadAssessmentsCSV(fname string, AsmtTypes *map[int64]AssessmentType) {
	t := LoadCSV(fname)
	for i := 0; i < len(t); i++ {
		CreateAssessmentsFromCSV(t[i], i+1, AsmtTypes)
	}
}