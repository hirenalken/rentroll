package rlib

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

//  CSV file format:
//      0     1         2        3        4           5        6            7            8
// RATRefNum,BID,PrimaryTenant,Payor,RentalStart,RentalStop,Renewal,SpecialProvisions,RentableName, ...
// 		"RAT001","REH","866-123-4567","866-123-4567","2004-01-01","2015-11-08",1,"",101
// 		"RAT001","REH","866-123-4567","866-123-4567","2004-01-01","2017-07-04",1,"",107
// 		"RAT001","REH","homerj@springfield.com","866-123-4567","2015-11-21","2016-11-21",1,"",101,102

// CreateRentalAgreement creates database records for the rental agreement defined in sa[]
func CreateRentalAgreement(sa []string) {
	var ra RentalAgreement
	var payor AgreementPayor
	var m []AgreementRentable

	des := strings.ToLower(strings.TrimSpace(sa[0]))
	if des == "ratrefnum" {
		return // this is just the column heading
	}

	//-------------------------------------------------------------------
	// Make sure the business is in the database
	//-------------------------------------------------------------------
	if len(des) > 0 {
		b1, _ := GetRentalAgreementTemplateByRefNum(des)
		if len(b1.ReferenceNumber) == 0 {
			Ulog("CreateRentalAgreement: business with designation %s does net exist\n", des)
			return
		}
		ra.RATID = b1.RATID
	}

	//-------------------------------------------------------------------
	// See if the biz exists, if so, set the BID
	//-------------------------------------------------------------------
	cmpdes := strings.TrimSpace(sa[1])
	if len(cmpdes) > 0 {
		b2, _ := GetBusinessByDesignation(cmpdes)
		if len(b2.Name) == 0 {
			fmt.Printf("CreateRentalAgreement: could not find business named %s\n", cmpdes)
			return
		}
		ra.BID = b2.BID
	}

	//-------------------------------------------------------------------
	//  Determine the primary tenant
	//-------------------------------------------------------------------
	s := strings.TrimSpace(sa[2]) // either the email address or the phone number
	t := GetTransactantByPhoneOrEmail(s)
	if t.TID == 0 {
		fmt.Printf("CreateRentalAgreement: could not find tenant with contact information %s\n", s)
		return
	}
	ra.PrimaryTenant = t.TID

	//-------------------------------------------------------------------
	//  Determine the payor
	//-------------------------------------------------------------------
	s = strings.TrimSpace(sa[3]) // either the email address or the phone number
	t = GetTransactantByPhoneOrEmail(s)
	if t.PID == 0 {
		fmt.Printf("CreateRentalAgreement: could not find tenant with contact information %s\n", s)
		return
	}
	payor.PID = t.PID

	//-------------------------------------------------------------------
	// Get the dates
	//-------------------------------------------------------------------
	DtStart, err := time.Parse(RRDATEINPFMT, strings.TrimSpace(sa[4]))
	if err != nil {
		fmt.Printf("CreateRentalAgreement: invalid start date:  %s\n", sa[4])
		return
	}
	ra.RentalStart = DtStart

	DtStop, err := time.Parse(RRDATEINPFMT, strings.TrimSpace(sa[5]))
	if err != nil {
		fmt.Printf("CreateRentalAgreement: invalid stop date:  %s\n", sa[5])
		return
	}
	ra.RentalStop = DtStop

	s = strings.TrimSpace(sa[6])
	if len(s) > 0 {
		i, err := strconv.Atoi(s)
		if err != nil {
			fmt.Printf("CreatePeopleFromCSV: Renewal value is invalid: %s\n", s)
			return
		}
		ra.Renewal = int64(i)
	}

	ra.SpecialProvisions = sa[7]

	// the rest of the arguments are rentables that are associated with
	// this rental agreement
	for i := 8; i < len(sa); i++ {
		s = strings.TrimSpace(sa[i])
		r, _ := GetRentableByName(s, ra.BID)

		if len(r.Name) > 0 {
			var ar AgreementRentable
			ar.RID = r.RID
			ar.DtStart = DtStart
			ar.DtStop = DtStop
			m = append(m, ar)
		}
	}

	// First write the rental agreement record, then write the agreementrentables and agreement payors
	RAID, err := InsertRentalAgreement(&ra)
	if nil != err {
		fmt.Printf("CreateRentalAgreement: error inserting RentalAgreement = %v\n", err)
	}
	for i := 0; i < len(m); i++ {
		m[i].RAID = RAID
		InsertAgreementRentable(&m[i])
	}
	payor.RAID = RAID
	payor.DtStart = DtStart
	payor.DtStop = DtStop
	InsertAgreementPayor(&payor)
}

// LoadRentalAgreementCSV loads a csv file with rental specialty types and processes each one
func LoadRentalAgreementCSV(fname string) {
	t := LoadCSV(fname)
	for i := 0; i < len(t); i++ {
		CreateRentalAgreement(t[i])
	}
}