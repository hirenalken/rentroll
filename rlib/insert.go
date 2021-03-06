package rlib

func insertError(err error, n string, a interface{}) error {
	if nil != err {
		Ulog("Insert%s: error inserting %s:  %v\n", n, n, err)
		Ulog("%s = %#v\n", n, a)
	}
	return err
}

// InsertAR writes a new AR record to the database. If the record is successfully written,
// the ARID field is set to its new value.
func InsertAR(a *AR) (int64, error) {
	var rid = int64(0)
	res, err := RRdb.Prepstmt.InsertAR.Exec(a.BID, a.Name, a.ARType, a.DebitLID, a.CreditLID, a.Description, a.RARequired, a.DtStart, a.DtStop, a.FLAGS, a.DefaultAmount, a.CreateBy, a.LastModBy)
	if nil == err {
		id, err := res.LastInsertId()
		if err == nil {
			rid = int64(id)
			a.ARID = rid
		}
	} else {
		err = insertError(err, "AR", *a)
	}
	return rid, err
}

// InsertAssessment writes a new assessmenttype record to the database. If the record is successfully written,
// the ASMID field is set to its new value.
func InsertAssessment(a *Assessment) (int64, error) {
	var rid = int64(0)

	//
	// DEBUG...
	//
	// if a.FLAGS&0x4 == 0 {
	// 	fmt.Printf(">>> INSERTING ASSESSMENT WITH FLAGS bit 2 not set.  FLAGS = %x\n", a.FLAGS)
	// 	debug.PrintStack()
	// 	// os.Exit(1)
	// }

	res, err := RRdb.Prepstmt.InsertAssessment.Exec(a.PASMID, a.RPASMID, a.BID, a.RID, a.ATypeLID, a.RAID, a.Amount, a.Start, a.Stop, a.RentCycle, a.ProrationCycle, a.InvoiceNo, a.AcctRule, a.ARID, a.FLAGS, a.Comment, a.CreateBy, a.LastModBy)
	if nil == err {
		id, err := res.LastInsertId()
		if err == nil {
			rid = int64(id)
			a.ASMID = rid
		}
	} else {
		err = insertError(err, "Insert", *a)
	}
	return rid, err
}

// InsertBuilding writes a new Building record to the database
func InsertBuilding(a *Building) (int64, error) {
	var rid = int64(0)
	res, err := RRdb.Prepstmt.InsertBuilding.Exec(a.BID, a.Address, a.Address2, a.City, a.State, a.PostalCode, a.Country, a.CreateBy, a.LastModBy)
	if nil == err {
		id, err := res.LastInsertId()
		if err == nil {
			rid = int64(id)
		}
	} else {
		err = insertError(err, "Building", *a)
	}
	return rid, err
}

// InsertBuildingWithID writes a new Building record to the database with the supplied bldgid
// the Building ID must be set in the supplied Building struct ptr (a.BLDGID).
func InsertBuildingWithID(a *Building) (int64, error) {
	var rid = int64(0)
	res, err := RRdb.Prepstmt.InsertBuildingWithID.Exec(a.BLDGID, a.BID, a.Address, a.Address2, a.City, a.State, a.PostalCode, a.Country, a.CreateBy, a.LastModBy)
	if nil == err {
		id, err := res.LastInsertId()
		if err == nil {
			rid = int64(id)
		}
	} else {
		err = insertError(err, "Building", *a)
	}
	return rid, err
}

// InsertBusiness writes a new Business record.
// returns the new Business ID and any associated error
func InsertBusiness(b *Business) (int64, error) {
	var bid = int64(0)
	res, err := RRdb.Prepstmt.InsertBusiness.Exec(b.Designation, b.Name, b.DefaultRentCycle, b.DefaultProrationCycle, b.DefaultGSRPC, b.CreateBy, b.LastModBy)
	if nil == err {
		id, err := res.LastInsertId()
		if err == nil {
			bid = int64(id)
			b.BID = bid
		}
		RRdb.BUDlist[b.Designation] = bid
	}
	return bid, err
}

// InsertCustomAttribute writes a new User record to the database
func InsertCustomAttribute(a *CustomAttribute) (int64, error) {
	var tid = int64(0)
	res, err := RRdb.Prepstmt.InsertCustomAttribute.Exec(a.BID, a.Type, a.Name, a.Value, a.Units, a.CreateBy, a.LastModBy)
	if nil == err {
		id, err := res.LastInsertId()
		if err == nil {
			tid = int64(id)
		}
	} else {
		err = insertError(err, "CustomAttribute", *a)
	}
	return tid, err
}

// InsertCustomAttributeRef writes a new assessmenttype record to the database
func InsertCustomAttributeRef(a *CustomAttributeRef) error {
	_, err := RRdb.Prepstmt.InsertCustomAttributeRef.Exec(a.ElementType, a.BID, a.ID, a.CID, a.CreateBy)
	return err
}

// InsertDemandSource writes a new DemandSource record to the database
func InsertDemandSource(a *DemandSource) (int64, error) {
	var tid = int64(0)
	res, err := RRdb.Prepstmt.InsertDemandSource.Exec(a.BID, a.Name, a.Industry, a.CreateBy, a.LastModBy)
	if nil == err {
		id, err := res.LastInsertId()
		if err == nil {
			tid = int64(id)
		}
	} else {
		err = insertError(err, "DemandSource", *a)
	}
	return tid, err
}

// InsertDeposit writes a new Deposit record to the database
func InsertDeposit(a *Deposit) (int64, error) {
	var rid = int64(0)
	res, err := RRdb.Prepstmt.InsertDeposit.Exec(a.BID, a.DEPID, a.DPMID, a.Dt, a.Amount, a.ClearedAmount, a.FLAGS, a.CreateBy, a.LastModBy)
	if nil == err {
		id, err := res.LastInsertId()
		if err == nil {
			rid = int64(id)
			a.DID = rid
		}
	} else {
		err = insertError(err, "Deposit", *a)
	}
	return rid, err
}

// InsertDepositMethod writes a new DepositMethod record to the database
func InsertDepositMethod(a *DepositMethod) error {
	_, err := RRdb.Prepstmt.InsertDepositMethod.Exec(a.BID, a.Method, a.CreateBy, a.LastModBy)
	if nil != err {
		return insertError(err, "DepositMethod", *a)
	}
	return err
}

// InsertDepositPart writes a new DepositPart record to the database
func InsertDepositPart(a *DepositPart) error {
	_, err := RRdb.Prepstmt.InsertDepositPart.Exec(a.DID, a.BID, a.RCPTID, a.CreateBy, a.LastModBy)
	if nil != err {
		return insertError(err, "DepositPart", *a)
	}
	return err
}

// InsertDepository writes a new Depository record to the database
func InsertDepository(a *Depository) (int64, error) {
	var id = int64(0)
	res, err := RRdb.Prepstmt.InsertDepository.Exec(a.BID, a.LID, a.Name, a.AccountNo, a.CreateBy, a.LastModBy)
	if nil == err {
		x, err := res.LastInsertId()
		if err == nil {
			id = int64(x)
			a.DEPID = id
		}
	} else {
		err = insertError(err, "Depository", *a)
	}
	return id, err
}

//======================================
//  EXPENSE
//======================================

// InsertExpense writes a new Expense record to the database
func InsertExpense(a *Expense) error {
	var rid = int64(0)
	res, err := RRdb.Prepstmt.InsertExpense.Exec(a.RPEXPID, a.BID, a.RID, a.RAID, a.Amount, a.Dt, a.AcctRule, a.ARID, a.FLAGS, a.Comment, a.CreateBy, a.LastModBy)
	if nil == err {
		id, err := res.LastInsertId()
		if err == nil {
			rid = int64(id)
			a.EXPID = rid
		}
	} else {
		return insertError(err, "Expense", *a)
	}
	return nil
}

//======================================
//  INVOICE
//======================================

// InsertInvoice writes a new Invoice record to the database
func InsertInvoice(a *Invoice) (int64, error) {
	var rid = int64(0)
	res, err := RRdb.Prepstmt.InsertInvoice.Exec(a.BID, a.Dt, a.DtDue, a.Amount, a.DeliveredBy, a.CreateBy, a.LastModBy)
	if nil == err {
		id, err := res.LastInsertId()
		if err == nil {
			rid = int64(id)
		}
	} else {
		err = insertError(err, "Invoice", *a)
	}
	return rid, err
}

// InsertInvoiceAssessment writes a new InvoiceAssessment record to the database
func InsertInvoiceAssessment(a *InvoiceAssessment) error {
	_, err := RRdb.Prepstmt.InsertInvoiceAssessment.Exec(a.InvoiceNo, a.BID, a.ASMID, a.CreateBy)
	if nil != err {
		return insertError(err, "DepositPart", *a)
	}
	return err
}

// InsertInvoicePayor writes a new InvoicePayor record to the database
func InsertInvoicePayor(a *InvoicePayor) error {
	_, err := RRdb.Prepstmt.InsertInvoicePayor.Exec(a.InvoiceNo, a.BID, a.PID, a.CreateBy)
	if nil != err {
		return insertError(err, "DepositPayor", *a)
	}
	return err
}

// InsertJournal writes a new Journal entry to the database
func InsertJournal(j *Journal) (int64, error) {
	var id = int64(0)
	res, err := RRdb.Prepstmt.InsertJournal.Exec(j.BID, j.Dt, j.Amount, j.Type, j.ID, j.Comment, j.CreateBy, j.LastModBy)
	if nil == err {
		nid, err := res.LastInsertId()
		if err == nil {
			id = int64(nid)
			j.JID = id
		}
	}
	return id, err
}

// InsertJournalAllocationEntry writes a new JournalAllocation record to the database. Also sets JAID with its
// newly assigned id.
func InsertJournalAllocationEntry(ja *JournalAllocation) error {
	// debug.PrintStack()
	res, err := RRdb.Prepstmt.InsertJournalAllocation.Exec(ja.BID, ja.JID, ja.RID, ja.RAID, ja.TCID, ja.RCPTID, ja.Amount, ja.ASMID, ja.EXPID, ja.AcctRule, ja.CreateBy)
	if nil == err {
		id, err := res.LastInsertId()
		if err == nil {
			ja.JAID = int64(id)
		}
	}
	return err
}

// InsertJournalMarker writes a new JournalMarker record to the database
func InsertJournalMarker(jm *JournalMarker) error {
	_, err := RRdb.Prepstmt.InsertJournalMarker.Exec(jm.BID, jm.State, jm.DtStart, jm.DtStop, jm.CreateBy, jm.LastModBy)
	return err
}

//======================================
//  LEDGER MARKER
//======================================

// InsertLedgerMarker writes a new LedgerMarker record to the database
func InsertLedgerMarker(l *LedgerMarker) error {
	res, err := RRdb.Prepstmt.InsertLedgerMarker.Exec(l.LID, l.BID, l.RAID, l.RID, l.TCID, l.Dt, l.Balance, l.State, l.CreateBy, l.LastModBy)
	if nil == err {
		id, err := res.LastInsertId()
		if err == nil {
			l.LMID = int64(id)
		}
	} else {
		Ulog("InsertLedgerMarker: err = %#v\n", err)
	}
	return err
}

// InsertLedgerEntry writes a new LedgerEntry to the database
func InsertLedgerEntry(l *LedgerEntry) (int64, error) {
	var rid = int64(0)
	res, err := RRdb.Prepstmt.InsertLedgerEntry.Exec(l.BID, l.JID, l.JAID, l.LID, l.RAID, l.RID, l.TCID, l.Dt, l.Amount, l.Comment, l.CreateBy, l.LastModBy)
	if nil == err {
		id, err := res.LastInsertId()
		if err == nil {
			rid = int64(id)
			l.LEID = rid
		}
	} else {
		Ulog("Error inserting LedgerEntry:  %v\n", err)
	}
	return rid, err
}

// InsertLedger writes a new GLAccount to the database
func InsertLedger(l *GLAccount) (int64, error) {
	var rid = int64(0)
	//                                            PLID, BID,     RAID,  TCID,   GLNumber,   Status,   Name,   AcctType,   AllowPost,  FLAGS,   Description, CreateBy, LastModBy
	res, err := RRdb.Prepstmt.InsertLedger.Exec(l.PLID, l.BID, l.RAID, l.TCID, l.GLNumber, l.Status, l.Name, l.AcctType, l.AllowPost, l.FLAGS, l.Description, l.CreateBy, l.LastModBy)
	if nil == err {
		id, err := res.LastInsertId()
		if err == nil {
			rid = int64(id)
			l.LID = rid
		}
	} else {
		Ulog("Error inserting GLAccount:  %v\n", err)
	}
	return rid, err
}

//======================================
// NOTE
//======================================

// InsertNote writes a new Note to the database
func InsertNote(a *Note) (int64, error) {
	var rid = int64(0)
	res, err := RRdb.Prepstmt.InsertNote.Exec(a.BID, a.NLID, a.PNID, a.NTID, a.RID, a.RAID, a.TCID, a.Comment, a.CreateBy, a.LastModBy)
	if nil == err {
		id, err := res.LastInsertId()
		if err == nil {
			rid = int64(id)
		}
	} else {
		Ulog("Error inserting Note:  %v\n", err)
	}
	return rid, err
}

//======================================
// NOTE LIST
//======================================

// InsertNoteList inserts a new wrapper for a notelist into the database
func InsertNoteList(a *NoteList) (int64, error) {
	var rid = int64(0)
	res, err := RRdb.Prepstmt.InsertNoteList.Exec(a.BID, a.CreateBy, a.LastModBy)
	if nil == err {
		id, err := res.LastInsertId()
		if err == nil {
			rid = int64(id)
		}
	} else {
		Ulog("Error inserting NoteList:  %v\n", err)
	}
	return rid, err
}

//======================================
// NOTE TYPE
//======================================

// InsertNoteType writes a new NoteType to the database
func InsertNoteType(a *NoteType) (int64, error) {
	var rid = int64(0)
	res, err := RRdb.Prepstmt.InsertNoteType.Exec(a.BID, a.Name, a.CreateBy, a.LastModBy)
	if nil == err {
		id, err := res.LastInsertId()
		if err == nil {
			rid = int64(id)
		}
	} else {
		Ulog("Error inserting NoteType:  %v\n", err)
	}
	return rid, err
}

//=======================================================
//  RATE PLAN
//=======================================================

// InsertRatePlan writes a new RatePlan record to the database
func InsertRatePlan(a *RatePlan) (int64, error) {
	var tid = int64(0)
	res, err := RRdb.Prepstmt.InsertRatePlan.Exec(a.BID, a.Name, a.CreateBy, a.LastModBy)
	if nil == err {
		id, err := res.LastInsertId()
		if err == nil {
			tid = int64(id)
		}
	} else {
		err = insertError(err, "RatePlan", *a)
	}
	a.RPID = tid
	return tid, err
}

// InsertRatePlanRef writes a new RatePlanRef record to the database
func InsertRatePlanRef(a *RatePlanRef) (int64, error) {
	var tid = int64(0)
	res, err := RRdb.Prepstmt.InsertRatePlanRef.Exec(a.BID, a.RPID, a.DtStart, a.DtStop, a.FeeAppliesAge, a.MaxNoFeeUsers, a.AdditionalUserFee, a.PromoCode, a.CancellationFee, a.FLAGS, a.CreateBy, a.LastModBy)
	if nil == err {
		id, err := res.LastInsertId()
		if err == nil {
			tid = int64(id)
		}
	} else {
		err = insertError(err, "RatePlanRef", *a)
	}
	a.RPRID = tid
	return tid, err
}

// InsertRatePlanRefRTRate writes a new RatePlanRefRTRate record to the database
func InsertRatePlanRefRTRate(a *RatePlanRefRTRate) error {
	_, err := RRdb.Prepstmt.InsertRatePlanRefRTRate.Exec(a.RPRID, a.BID, a.RTID, a.FLAGS, a.Val, a.CreateBy)
	if nil != err {
		return insertError(err, "RatePlanRefRTRate", *a)
	}
	return err
}

// InsertRatePlanRefSPRate writes a new RatePlanRefSPRate record to the database
func InsertRatePlanRefSPRate(a *RatePlanRefSPRate) error {
	_, err := RRdb.Prepstmt.InsertRatePlanRefSPRate.Exec(a.RPRID, a.BID, a.RTID, a.RSPID, a.FLAGS, a.Val, a.CreateBy)
	if nil != err {
		return insertError(err, "RatePlanRefSPRate", *a)
	}
	return err
}

//=======================================================
//  PAYMENT
//=======================================================

// InsertPaymentType writes a new assessmenttype record to the database
func InsertPaymentType(a *PaymentType) error {
	res, err := RRdb.Prepstmt.InsertPaymentType.Exec(a.BID, a.Name, a.Description, a.CreateBy, a.LastModBy)
	if nil == err {
		id, err := res.LastInsertId()
		if err == nil {
			a.PMTID = int64(id)
		}
	} else {
		return insertError(err, "Payor", *a)
	}
	return err
}

// InsertPayor writes a new User record to the database
func InsertPayor(a *Payor) (int64, error) {
	var tid = int64(0)
	res, err := RRdb.Prepstmt.InsertPayor.Exec(a.TCID, a.BID, a.CreditLimit, a.TaxpayorID, a.AccountRep, a.EligibleFuturePayor, a.CreateBy, a.LastModBy)
	if nil == err {
		id, err := res.LastInsertId()
		if err == nil {
			tid = int64(id)
		}
	} else {
		err = insertError(err, "Payor", *a)
	}
	return tid, err
}

// InsertProspect writes a new User record to the database
func InsertProspect(a *Prospect) (int64, error) {
	var tid = int64(0)
	res, err := RRdb.Prepstmt.InsertProspect.Exec(a.TCID, a.BID, a.EmployerName, a.EmployerStreetAddress, a.EmployerCity,
		a.EmployerState, a.EmployerPostalCode, a.EmployerEmail, a.EmployerPhone, a.Occupation, a.ApplicationFee,
		a.DesiredUsageStartDate, a.RentableTypePreference, a.FLAGS, a.Approver, a.DeclineReasonSLSID, a.OtherPreferences,
		a.FollowUpDate, a.CSAgent, a.OutcomeSLSID, a.FloatingDeposit, a.RAID, a.CreateBy, a.LastModBy)
	if nil == err {
		id, err := res.LastInsertId()
		if err == nil {
			tid = int64(id)
		}
	} else {
		err = insertError(err, "Prospect", *a)
	}
	return tid, err
}

// InsertRentable writes a new Rentable record to the database
func InsertRentable(a *Rentable) (int64, error) {
	var rid = int64(0)
	res, err := RRdb.Prepstmt.InsertRentable.Exec(a.BID, a.RentableName, a.AssignmentTime, a.CreateBy, a.LastModBy)
	if nil == err {
		id, err := res.LastInsertId()
		if err == nil {
			rid = int64(id)
		}
	} else {
		err = insertError(err, "Rentable", *a)
	}
	return rid, err
}

//=======================================================
//  R E C E I P T
//=======================================================

// InsertReceipt writes a new Receipt record to the database. If the record is successfully written,
// the RCPTID field is set to its new value.
func InsertReceipt(r *Receipt) (int64, error) {
	var tid = int64(0)
	res, err := RRdb.Prepstmt.InsertReceipt.Exec(r.PRCPTID, r.BID, r.TCID, r.PMTID, r.DEPID, r.DID, r.Dt, r.DocNo, r.Amount, r.AcctRuleReceive, r.ARID, r.AcctRuleApply, r.FLAGS, r.Comment, r.OtherPayorName, r.CreateBy, r.LastModBy)
	if nil == err {
		id, err := res.LastInsertId()
		if err == nil {
			tid = int64(id)
			r.RCPTID = tid
		}
	} else {
		err = insertError(err, "Receipt", *r)
	}
	return tid, err
}

// InsertReceiptAllocation writes a new ReceiptAllocation record to the database
func InsertReceiptAllocation(a *ReceiptAllocation) (int64, error) {
	var tid = int64(0)
	res, err := RRdb.Prepstmt.InsertReceiptAllocation.Exec(a.RCPTID, a.BID, a.RAID, a.Dt, a.Amount, a.ASMID, a.FLAGS, a.AcctRule, a.CreateBy, a.LastModBy)
	if nil == err {
		id, err := res.LastInsertId()
		if err == nil {
			tid = int64(id)
			a.RCPAID = tid
		}
	} else {
		err = insertError(err, "ReceiptAllocation", *a)
	}
	return tid, err
}

// InsertRentalAgreement writes a new RentalAgreement record to the database
func InsertRentalAgreement(a *RentalAgreement) (int64, error) {
	var tid = int64(0)
	res, err := RRdb.Prepstmt.InsertRentalAgreement.Exec(a.RATID, a.BID, a.NLID, a.AgreementStart, a.AgreementStop, a.PossessionStart, a.PossessionStop, a.RentStart, a.RentStop, a.RentCycleEpoch, a.UnspecifiedAdults, a.UnspecifiedChildren, a.Renewal, a.SpecialProvisions, a.LeaseType, a.ExpenseAdjustmentType, a.ExpensesStop, a.ExpenseStopCalculation, a.BaseYearEnd, a.ExpenseAdjustment, a.EstimatedCharges, a.RateChange, a.NextRateChange, a.PermittedUses, a.ExclusiveUses, a.ExtensionOption, a.ExtensionOptionNotice, a.ExpansionOption, a.ExpansionOptionNotice, a.RightOfFirstRefusal, a.FLAGS, a.CreateBy, a.LastModBy)
	if nil == err {
		id, err := res.LastInsertId()
		if err == nil {
			tid = int64(id)
			a.RAID = tid
		}
	} else {
		err = insertError(err, "RentalAgreement", *a)
	}
	return tid, err
}

// InsertRentalAgreementPayor writes a new User record to the database
func InsertRentalAgreementPayor(a *RentalAgreementPayor) (int64, error) {
	var tid = int64(0)
	res, err := RRdb.Prepstmt.InsertRentalAgreementPayor.Exec(a.RAID, a.BID, a.TCID, a.DtStart, a.DtStop, a.FLAGS, a.CreateBy)
	if nil == err {
		id, err := res.LastInsertId()
		if err == nil {
			tid = int64(id)
			a.RAPID = tid
		}
	} else {
		err = insertError(err, "RentalAgreementPayor", *a)
	}
	return tid, err
}

// InsertRentalAgreementPet writes a new User record to the database
func InsertRentalAgreementPet(a *RentalAgreementPet) (int64, error) {
	var tid = int64(0)
	res, err := RRdb.Prepstmt.InsertRentalAgreementPet.Exec(a.BID, a.RAID, a.Type, a.Breed, a.Color, a.Weight, a.Name, a.DtStart, a.DtStop, a.CreateBy, a.LastModBy)
	if nil == err {
		id, err := res.LastInsertId()
		if err == nil {
			tid = int64(id)
		}
	} else {
		err = insertError(err, "RentalAgreementPet", *a)
	}
	return tid, err
}

// InsertRentalAgreementRentable writes a new User record to the database
func InsertRentalAgreementRentable(a *RentalAgreementRentable) (int64, error) {
	var tid = int64(0)
	res, err := RRdb.Prepstmt.InsertRentalAgreementRentable.Exec(a.RAID, a.BID, a.RID, a.CLID, a.ContractRent, a.RARDtStart, a.RARDtStop, a.CreateBy)
	if nil == err {
		id, err := res.LastInsertId()
		if err == nil {
			tid = int64(id)
			a.RARID = tid
		}
	} else {
		err = insertError(err, "RentalAgreementRentable", *a)
	}
	return tid, err
}

//=======================================================
//  RENTAL AGREEMENT TEMPLATE
//=======================================================

// InsertRentalAgreementTemplate writes a new User record to the database
func InsertRentalAgreementTemplate(a *RentalAgreementTemplate) (int64, error) {
	var tid = int64(0)
	res, err := RRdb.Prepstmt.InsertRentalAgreementTemplate.Exec(a.BID, a.RATemplateName, a.CreateBy, a.LastModBy)
	if nil == err {
		id, err := res.LastInsertId()
		if err == nil {
			tid = int64(id)
		}
	} else {
		err = insertError(err, "RentalAgreementTemplate", *a)
	}
	return tid, err
}

// InsertRentableSpecialty writes a new RentableSpecialty record to the database
func InsertRentableSpecialty(a *RentableSpecialty) error {
	_, err := RRdb.Prepstmt.InsertRentableSpecialtyType.Exec(a.BID, a.Name, a.Fee, a.Description, a.CreateBy)
	return err
}

// InsertRentableMarketRates writes a new marketrate record to the database
func InsertRentableMarketRates(r *RentableMarketRate) error {
	_, err := RRdb.Prepstmt.InsertRentableMarketRates.Exec(r.RTID, r.BID, r.MarketRate, r.DtStart, r.DtStop, r.CreateBy)
	return err
}

// InsertRentableType writes a new RentableType record to the database
func InsertRentableType(a *RentableType) (int64, error) {
	var rid = int64(0)
	res, err := RRdb.Prepstmt.InsertRentableType.Exec(a.BID, a.Style, a.Name, a.RentCycle, a.Proration, a.GSRPC, a.ManageToBudget, a.CreateBy, a.LastModBy)
	if nil == err {
		id, err := res.LastInsertId()
		if err == nil {
			rid = int64(id)
		}
	} else {
		Ulog("Error inserting RentableType:  %v\n", err)
	}
	return rid, err
}

// InsertRentableSpecialtyRef writes a new RentableSpecialty record to the database
func InsertRentableSpecialtyRef(a *RentableSpecialtyRef) error {
	_, err := RRdb.Prepstmt.InsertRentableSpecialtyRef.Exec(a.BID, a.RID, a.RSPID, a.DtStart, a.DtStop, a.CreateBy, a.LastModBy)
	return err
}

// InsertRentableStatus writes a new RentableStatus record to the database
func InsertRentableStatus(a *RentableStatus) error {
	res, err := RRdb.Prepstmt.InsertRentableStatus.Exec(a.RID, a.BID, a.DtStart, a.DtStop, a.DtNoticeToVacate, a.Status, a.CreateBy, a.LastModBy)
	if nil != err {
		return insertError(err, "RentableStatus", *a)
	}
	id, err := res.LastInsertId()
	if err == nil {
		a.RSID = int64(id)
	}
	return err

}

// InsertRentableTypeRef writes a new RentableTypeRef record to the database
func InsertRentableTypeRef(a *RentableTypeRef) error {
	res, err := RRdb.Prepstmt.InsertRentableTypeRef.Exec(a.RID, a.BID, a.RTID, a.OverrideRentCycle, a.OverrideProrationCycle, a.DtStart, a.DtStop, a.CreateBy, a.LastModBy)
	if nil != err {
		return insertError(err, "RentableTypeRef", *a)
	}
	id, err := res.LastInsertId()
	if err == nil {
		a.RTRID = int64(id)
	}
	return err
}

// InsertRentableUser writes a new User record to the database
func InsertRentableUser(a *RentableUser) error {
	res, err := RRdb.Prepstmt.InsertRentableUser.Exec(a.RID, a.BID, a.TCID, a.DtStart, a.DtStop, a.CreateBy)
	if nil != err {
		return insertError(err, "RentableUser", *a)
	}
	id, err := res.LastInsertId()
	if err == nil {
		a.RUID = int64(id)
	}
	return err
}

// InsertStringList writes a new StringList record to the database
func InsertStringList(a *StringList) (int64, error) {
	var tid = int64(0)
	res, err := RRdb.Prepstmt.InsertStringList.Exec(a.BID, a.Name, a.CreateBy, a.LastModBy)
	if nil == err {
		id, err := res.LastInsertId()
		if err == nil {
			tid = int64(id)
		}
	} else {
		err = insertError(err, "StringList", *a)
	}
	a.SLID = tid
	InsertSLStrings(a)
	return tid, err
}

// InsertSLStrings writes a the list of strings in a StringList to the database
// THIS SHOULD BE PUT IN A TRANSACTION
func InsertSLStrings(a *StringList) {
	// DeleteSLStrings(a.SLID)
	for i := 0; i < len(a.S); i++ {
		a.S[i].SLID = a.SLID
		_, err := RRdb.Prepstmt.InsertSLString.Exec(a.BID, a.SLID, a.S[i].Value, a.CreateBy, a.S[i].LastModBy)
		if nil != err {
			Ulog("InsertSLString: error:  %v\n", err)
		}
	}
}

// InsertTransactant writes a new Transactant record to the database
func InsertTransactant(a *Transactant) (int64, error) {
	var tid = int64(0)
	res, err := RRdb.Prepstmt.InsertTransactant.Exec(a.BID, a.NLID, a.FirstName, a.MiddleName, a.LastName, a.PreferredName, a.CompanyName, a.IsCompany, a.PrimaryEmail, a.SecondaryEmail, a.WorkPhone, a.CellPhone, a.Address, a.Address2, a.City, a.State, a.PostalCode, a.Country, a.Website, a.CreateBy, a.LastModBy)
	if nil == err {
		id, err := res.LastInsertId()
		if err == nil {
			tid = int64(id)
			a.TCID = tid
		}
	} else {
		err = insertError(err, "Transactant", *a)
	}
	return tid, err
}

// InsertUser writes a new User record to the database
func InsertUser(a *User) (int64, error) {
	var tid = int64(0)
	res, err := RRdb.Prepstmt.InsertUser.Exec(a.TCID, a.BID, a.Points, a.DateofBirth, a.EmergencyContactName, a.EmergencyContactAddress, a.EmergencyContactTelephone, a.EmergencyEmail, a.AlternateAddress, a.EligibleFutureUser, a.Industry, a.SourceSLSID, a.CreateBy, a.LastModBy)
	if nil == err {
		id, err := res.LastInsertId()
		if err == nil {
			tid = int64(id)
		}
	} else {
		err = insertError(err, "User", *a)
	}
	return tid, err
}

// InsertVehicle writes a new Vehicle record to the database
func InsertVehicle(a *Vehicle) (int64, error) {
	var tid = int64(0)
	res, err := RRdb.Prepstmt.InsertVehicle.Exec(a.TCID, a.BID, a.VehicleType, a.VehicleMake, a.VehicleModel, a.VehicleColor, a.VehicleYear, a.LicensePlateState, a.LicensePlateNumber, a.ParkingPermitNumber, a.DtStart, a.DtStop, a.CreateBy, a.LastModBy)
	if nil == err {
		id, err := res.LastInsertId()
		if err == nil {
			tid = int64(id)
		}
	} else {
		err = insertError(err, "Vehicle", *a)
	}
	return tid, err
}
