package rlib

func buildPreparedStatements() {
	var err error
	// Prepare("select deduction from deductions where uid=?")
	// Prepare("select type from compensation where uid=?")
	// Prepare("INSERT INTO compensation (uid,type) VALUES(?,?)")
	// Prepare("DELETE FROM compensation WHERE UID=?")
	// Prepare("update classes set Name=?,Designation=?,Description=?,lastmodby=? where ClassCode=?")
	// Errcheck(err)

	//===============================
	//  Agreement Payor
	//===============================
	RRdb.Prepstmt.InsertAgreementPayor, err = RRdb.dbrr.Prepare("INSERT INTO agreementpayors (RAID,PID,DtStart,DtStop) VALUES(?,?,?,?)")
	Errcheck(err)

	//===============================
	//  Agreement Rentable
	//===============================
	RRdb.Prepstmt.InsertAgreementRentable, err = RRdb.dbrr.Prepare("INSERT INTO agreementrentables (RAID,RID,DtStart,DtStop) VALUES(?,?,?,?)")
	Errcheck(err)
	RRdb.Prepstmt.FindAgreementByRentable, err = RRdb.dbrr.Prepare("SELECT RAID,RID,DtStart,DtStop from agreementrentables where RID=? and DtStop>=? and DtStart<=?")
	Errcheck(err)

	//===============================
	//  Agreement Tenant
	//===============================
	RRdb.Prepstmt.InsertAgreementTenant, err = RRdb.dbrr.Prepare("INSERT INTO agreementtenants (RAID,TID,DtStart,DtStop) VALUES(?,?,?,?)")
	Errcheck(err)

	//===============================
	//  Assessments
	//===============================
	RRdb.Prepstmt.GetAssessment, err = RRdb.dbrr.Prepare("SELECT ASMID, BID, RID, ASMTID, RAID, Amount, Start, Stop, Accrual, ProrationMethod, AcctRule,Comment, LastModTime, LastModBy from assessments WHERE ASMID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetAllAssessmentsByBusiness, err = RRdb.dbrr.Prepare("SELECT ASMID,BID,RID,ASMTID,RAID,Amount,Start,Stop,Accrual,ProrationMethod,AcctRule,Comment,LastModTime,LastModBy FROM assessments WHERE BID=? and Start<? and Stop>=?")
	Errcheck(err)
	RRdb.Prepstmt.GetAllRentableAssessments, err = RRdb.dbrr.Prepare("SELECT ASMID,BID,RID,ASMTID,RAID,Amount,Start,Stop,Accrual,ProrationMethod,AcctRule,Comment,LastModTime,LastModBy FROM assessments WHERE RID=? and Stop >= ? and Start < ?")
	Errcheck(err)
	RRdb.Prepstmt.InsertAssessment, err = RRdb.dbrr.Prepare("INSERT INTO assessments (ASMID,BID,RID,ASMTID,RAID,Amount,Start,Stop,Accrual,ProrationMethod,AcctRule,Comment,LastModBy) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?)")
	Errcheck(err)

	//===============================
	//  AssessmentType
	//===============================
	RRdb.Prepstmt.GetAssessmentType, err = RRdb.dbrr.Prepare("SELECT ASMTID,OccupancyRqd,Name,Description,LastModTime,LastModBy FROM assessmenttypes WHERE ASMTID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetAssessmentTypeByName, err = RRdb.dbrr.Prepare("SELECT ASMTID,OccupancyRqd,Name,Description,LastModTime,LastModBy FROM assessmenttypes WHERE Name=?")
	Errcheck(err)
	RRdb.Prepstmt.InsertAssessmentType, err = RRdb.dbrr.Prepare("INSERT INTO assessmenttypes (OccupancyRqd,Name,Description,LastModBy) VALUES(?,?,?,?)")
	Errcheck(err)

	//===============================
	//  Building
	//===============================
	RRdb.Prepstmt.InsertBuilding, err = RRdb.dbrr.Prepare("INSERT INTO building (BID,Address,Address2,City,State,PostalCode,Country,LastModBy) VALUES(?,?,?,?,?,?,?,?)")
	Errcheck(err)
	RRdb.Prepstmt.InsertBuildingWithID, err = RRdb.dbrr.Prepare("INSERT INTO building (BLDGID,BID,Address,Address2,City,State,PostalCode,Country,LastModBy) VALUES(?,?,?,?,?,?,?,?,?)")
	Errcheck(err)
	RRdb.Prepstmt.GetBuilding, err = RRdb.dbrr.Prepare("SELECT BLDGID,BID,Address,Address2,City,State,PostalCode,Country,LastModTime,LastModBy FROM building WHERE BLDGID=?")
	Errcheck(err)

	//==========================================
	// Business
	//==========================================
	RRdb.Prepstmt.GetAllBusinesses, err = RRdb.dbrr.Prepare("SELECT BID,DES,Name,DefaultAccrual,ParkingPermitInUse,LastModTime,LastModBy FROM business")
	Errcheck(err)
	RRdb.Prepstmt.GetBusiness, err = RRdb.dbrr.Prepare("SELECT BID,DES,Name,DefaultAccrual,ParkingPermitInUse,LastModTime,LastModBy FROM business WHERE BID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetBusinessByDesignation, err = RRdb.dbrr.Prepare("SELECT BID,DES,Name,DefaultAccrual,ParkingPermitInUse,LastModTime,LastModBy FROM business WHERE DES=?")
	Errcheck(err)
	RRdb.Prepstmt.GetAllBusinessSpecialtyTypes, err = RRdb.dbrr.Prepare("SELECT RSPID,BID,Name,Fee,Description FROM rentablespecialtytypes WHERE BID=?")
	Errcheck(err)
	RRdb.Prepstmt.InsertBusiness, err = RRdb.dbrr.Prepare("INSERT INTO business (DES,Name,DefaultAccrual,ParkingPermitInUse,LastModBy) VALUES(?,?,?,?,?)")
	Errcheck(err)

	//==========================================
	// Custom Attribute
	//==========================================
	RRdb.Prepstmt.InsertCustomAttribute, err = RRdb.dbrr.Prepare("INSERT INTO customattr (Type,Name,Value,LastModBy) VALUES(?,?,?,?)")
	Errcheck(err)
	RRdb.Prepstmt.GetCustomAttribute, err = RRdb.dbrr.Prepare("SELECT CID,Type,Name,Value,LastModTime,LastModBy FROM customattr where CID=?")
	Errcheck(err)
	RRdb.Prepstmt.DeleteCustomAttribute, err = RRdb.dbrr.Prepare("DELETE FROM customattr where CID=?")
	Errcheck(err)

	//==========================================
	// Custom Attribute Ref
	//==========================================
	RRdb.Prepstmt.InsertCustomAttributeRef, err = RRdb.dbrr.Prepare("INSERT INTO customattrref (ElementType,ID,CID) VALUES(?,?,?)")
	Errcheck(err)
	RRdb.Prepstmt.GetCustomAttributeRefs, err = RRdb.dbrr.Prepare("SELECT CID FROM customattrref where ElementType=? and ID=?")
	Errcheck(err)
	RRdb.Prepstmt.DeleteCustomAttributeRef, err = RRdb.dbrr.Prepare("DELETE FROM customattrref where CID=? and ElementType=? and ID=?")
	Errcheck(err)

	//==========================================
	// JOURNAL
	//==========================================
	RRdb.Prepstmt.GetJournal, err = RRdb.dbrr.Prepare("select JID,BID,RAID,Dt,Amount,Type,ID,Comment,LastModTime,LastModBy from journal WHERE JID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetJournalMarker, err = RRdb.dbrr.Prepare("SELECT JMID,BID,State,DtStart,DtStop from journalmarker WHERE JMID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetJournalMarkers, err = RRdb.dbrr.Prepare("SELECT JMID,BID,State,DtStart,DtStop from journalmarker ORDER BY JMID DESC LIMIT ?")
	Errcheck(err)
	RRdb.Prepstmt.InsertJournal, err = RRdb.dbrr.Prepare("INSERT INTO journal (BID,RAID,Dt,Amount,Type,ID,Comment,LastModBy) VALUES(?,?,?,?,?,?,?,?)")
	Errcheck(err)
	RRdb.Prepstmt.GetAllJournalsInRange, err = RRdb.dbrr.Prepare("SELECT JID,BID,RAID,Dt,Amount,Type,ID,Comment,LastModTime,LastModBy from journal WHERE BID=? and ?<=Dt and Dt<?")
	Errcheck(err)
	RRdb.Prepstmt.GetJournalAllocation, err = RRdb.dbrr.Prepare("SELECT JAID,JID,RID,Amount,ASMID,AcctRule from journalallocation WHERE JAID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetJournalAllocations, err = RRdb.dbrr.Prepare("SELECT JAID,JID,RID,Amount,ASMID,AcctRule from journalallocation WHERE JID=?")
	Errcheck(err)

	RRdb.Prepstmt.InsertJournalAllocation, err = RRdb.dbrr.Prepare("INSERT INTO journalallocation (JID,RID,Amount,ASMID,AcctRule) VALUES(?,?,?,?,?)")
	Errcheck(err)
	RRdb.Prepstmt.InsertJournalMarker, err = RRdb.dbrr.Prepare("INSERT INTO journalmarker (BID,State,DtStart,DtStop) VALUES(?,?,?,?)")
	Errcheck(err)

	RRdb.Prepstmt.DeleteJournalAllocations, err = RRdb.dbrr.Prepare("DELETE FROM journalallocation WHERE JID=?")
	Errcheck(err)
	RRdb.Prepstmt.DeleteJournalEntry, err = RRdb.dbrr.Prepare("DELETE FROM journal WHERE JID=?")
	Errcheck(err)
	RRdb.Prepstmt.DeleteJournalMarker, err = RRdb.dbrr.Prepare("DELETE FROM journalmarker WHERE JMID=?")
	Errcheck(err)

	//==========================================
	// LEDGER
	//==========================================
	LDGRfields := "LID,BID,RAID,GLNumber,Status,Type,Name,AcctType,RAAssociated,LastModTime,LastModBy"
	RRdb.Prepstmt.GetLedgerByGLNo, err = RRdb.dbrr.Prepare("SELECT " + LDGRfields + " FROM ledger WHERE BID=? and GLNumber=?")
	Errcheck(err)
	RRdb.Prepstmt.GetLedgerByType, err = RRdb.dbrr.Prepare("SELECT " + LDGRfields + " FROM ledger WHERE BID=? and Type=?")
	Errcheck(err)
	RRdb.Prepstmt.GetLedger, err = RRdb.dbrr.Prepare("SELECT " + LDGRfields + " FROM ledger WHERE LID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetLedgerList, err = RRdb.dbrr.Prepare("SELECT " + LDGRfields + " FROM ledger WHERE BID=? ORDER BY GLNumber ASC")
	Errcheck(err)
	RRdb.Prepstmt.GetDefaultLedgers, err = RRdb.dbrr.Prepare("SELECT " + LDGRfields + " FROM ledger WHERE BID=? and Type>=10 ORDER BY GLNumber ASC")
	Errcheck(err)
	RRdb.Prepstmt.InsertLedger, err = RRdb.dbrr.Prepare("INSERT INTO ledger (BID,RAID,GLNumber,Status,Type,Name,AcctType,RAAssociated,LastModBy) VALUES(?,?,?,?,?,?,?,?,?)")
	Errcheck(err)
	RRdb.Prepstmt.DeleteLedger, err = RRdb.dbrr.Prepare("DELETE FROM ledger WHERE LID=?")
	Errcheck(err)
	RRdb.Prepstmt.UpdateLedger, err = RRdb.dbrr.Prepare("UPDATE ledger SET BID=?,RAID=?,GLNumber=?,Status=?,Type=?,Name=?,AcctType=?,RAAssociated=?,LastModBy=? WHERE LID=?")
	Errcheck(err)

	//==========================================
	// LEDGER ENTRY
	//==========================================
	LEfields := "LEID,BID,JID,JAID,GLNumber,Dt,Amount,Comment,LastModTime,LastModBy"
	RRdb.Prepstmt.GetAllLedgerEntriesInRange, err = RRdb.dbrr.Prepare("SELECT " + LEfields + " from ledgerentry WHERE BID=? and ?<=Dt and Dt<?")
	Errcheck(err)
	RRdb.Prepstmt.GetLedgerEntriesInRangeByGLNo, err = RRdb.dbrr.Prepare("SELECT " + LEfields + " from ledgerentry WHERE BID=? and GLNumber=? and ?<=Dt and Dt<? ORDER BY JAID ASC")
	Errcheck(err)
	RRdb.Prepstmt.GetLedgerEntry, err = RRdb.dbrr.Prepare("SELECT " + LEfields + " FROM ledgerentry where LEID=?")
	Errcheck(err)
	RRdb.Prepstmt.InsertLedgerEntry, err = RRdb.dbrr.Prepare("INSERT INTO ledgerentry (BID,JID,JAID,GLNumber,Dt,Amount,Comment,LastModBy) VALUES(?,?,?,?,?,?,?,?)")
	Errcheck(err)
	RRdb.Prepstmt.DeleteLedgerEntry, err = RRdb.dbrr.Prepare("DELETE FROM ledgerentry WHERE LEID=?")
	Errcheck(err)

	//==========================================
	// LEDGER MARKER
	//==========================================
	LMfields := "LMID,LID,BID,DtStart,DtStop,Balance,State,LastModTime,LastModBy"
	RRdb.Prepstmt.GetLatestLedgerMarkerByLID, err = RRdb.dbrr.Prepare("SELECT " + LMfields + " FROM ledgermarker WHERE BID=? and LID=? ORDER BY DtStop DESC")
	Errcheck(err)
	RRdb.Prepstmt.GetLedgerMarkerByDateRange, err = RRdb.dbrr.Prepare("SELECT " + LMfields + " FROM ledgermarker WHERE BID=? and LID=? and DtStop>? and DtStart<? ORDER BY LID ASC")
	Errcheck(err)
	RRdb.Prepstmt.GetLedgerMarkers, err = RRdb.dbrr.Prepare("SELECT " + LMfields + " FROM ledgermarker WHERE BID=? ORDER BY LMID DESC LIMIT ?")
	Errcheck(err)
	RRdb.Prepstmt.GetAllLedgerMarkersInRange, err = RRdb.dbrr.Prepare("SELECT " + LMfields + " FROM ledgermarker WHERE BID=? and DtStop>? and DtStart<=?")
	Errcheck(err)
	RRdb.Prepstmt.DeleteLedgerMarker, err = RRdb.dbrr.Prepare("DELETE FROM ledgermarker WHERE LMID=?")
	Errcheck(err)
	RRdb.Prepstmt.InsertLedgerMarker, err = RRdb.dbrr.Prepare("INSERT INTO ledgermarker (LID,BID,DtStart,DtStop,Balance,State,LastModBy) VALUES(?,?,?,?,?,?,?)")
	Errcheck(err)
	RRdb.Prepstmt.UpdateLedgerMarker, err = RRdb.dbrr.Prepare("UPDATE ledgermarker SET LMID=?,LID=?,BID=?,DtStart=?,DtStop=?,Balance=?,State=?,LastModBy=? WHERE LMID=?")
	Errcheck(err)

	//==========================================
	// PAYMENT TYPES
	//==========================================
	RRdb.Prepstmt.InsertPaymentType, err = RRdb.dbrr.Prepare("INSERT INTO paymenttypes (BID,Name,Description,LastModBy) VALUES(?,?,?,?)")
	Errcheck(err)
	RRdb.Prepstmt.GetPaymentTypesByBusiness, err = RRdb.dbrr.Prepare("SELECT PMTID,BID,Name,Description,LastModTime,LastModBy FROM paymenttypes WHERE BID=?")
	Errcheck(err)

	//==========================================
	// PAYOR
	//==========================================
	RRdb.Prepstmt.InsertPayor, err = RRdb.dbrr.Prepare("INSERT INTO payor (TCID,CreditLimit,EmployerName,EmployerStreetAddress,EmployerCity,EmployerState,EmployerPostalCode,EmployerEmail,EmployerPhone,Occupation,LastModBy) VALUES(?,?,?,?,?,?,?,?,?,?,?)")
	Errcheck(err)
	RRdb.Prepstmt.GetPayor, err = RRdb.dbrr.Prepare("SELECT PID,TCID,CreditLimit,EmployerName,EmployerStreetAddress,EmployerCity,EmployerState,EmployerPostalCode,EmployerEmail,EmployerPhone,Occupation,LastModTime,LastModBy FROM payor where PID=?")
	Errcheck(err)

	//==========================================
	// PROSPECT
	//==========================================
	RRdb.Prepstmt.InsertProspect, err = RRdb.dbrr.Prepare("INSERT INTO prospect (TCID,ApplicationFee,LastModBy) VALUES(?,?,?)")
	Errcheck(err)
	RRdb.Prepstmt.GetProspect, err = RRdb.dbrr.Prepare("SELECT PRSPID,TCID,ApplicationFee,LastModTime,LastModBy FROM prospect where PRSPID=?")
	Errcheck(err)

	//==========================================
	// RECEIPT
	//==========================================
	RRdb.Prepstmt.GetReceipt, err = RRdb.dbrr.Prepare("SELECT RCPTID,BID,RAID,PMTID,Dt,Amount,AcctRule,Comment,LastModTime,LastModBy FROM receipt WHERE RCPTID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetReceiptsInDateRange, err = RRdb.dbrr.Prepare("SELECT RCPTID,BID,RAID,PMTID,Dt,Amount,AcctRule,Comment,LastModTime,LastModBy from receipt WHERE BID=? and Dt >= ? and DT < ?")
	Errcheck(err)
	RRdb.Prepstmt.InsertReceipt, err = RRdb.dbrr.Prepare("INSERT INTO receipt (RCPTID,BID,RAID,PMTID,Dt,Amount,AcctRule,Comment,LastModBy) VALUES(?,?,?,?,?,?,?,?,?)")
	Errcheck(err)
	RRdb.Prepstmt.DeleteReceipt, err = RRdb.dbrr.Prepare("DELETE FROM receipt WHERE RCPTID=?")
	Errcheck(err)

	//==========================================
	// RECEIPT ALLOCATION
	//==========================================
	RRdb.Prepstmt.GetReceiptAllocations, err = RRdb.dbrr.Prepare("SELECT RCPTID,Amount,ASMID,AcctRule from receiptallocation WHERE RCPTID=?")
	Errcheck(err)
	RRdb.Prepstmt.InsertReceiptAllocation, err = RRdb.dbrr.Prepare("INSERT INTO receiptallocation (RCPTID,Amount,ASMID,AcctRule) VALUES(?,?,?,?)")
	Errcheck(err)
	RRdb.Prepstmt.DeleteReceiptAllocations, err = RRdb.dbrr.Prepare("DELETE FROM receiptallocation WHERE RCPTID=?")
	Errcheck(err)

	//===============================
	//  Rentable
	//===============================
	RRdb.Prepstmt.InsertRentable, err = RRdb.dbrr.Prepare("INSERT INTO rentable (RTID,BID,Name,Assignment,Report,DefaultOccType,OccType,State,LastModBy) VALUES(?,?,?,?,?,?,?,?,?)")
	Errcheck(err)
	RRdb.Prepstmt.GetRentable, err = RRdb.dbrr.Prepare("SELECT RID,RTID,BID,Name,Assignment,Report,DefaultOccType,OccType,State,LastModTime,LastModBy FROM rentable WHERE RID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetRentableByName, err = RRdb.dbrr.Prepare("SELECT RID,RTID,BID,Name,Assignment,Report,DefaultOccType,OccType,State,LastModTime,LastModBy FROM rentable WHERE Name=? AND BID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetAllRentablesByBusiness, err = RRdb.dbrr.Prepare("SELECT RID,RTID,BID,Name,Assignment,Report,DefaultOccType,OccType,State,LastModTime,LastModBy FROM rentable WHERE BID=?")
	Errcheck(err)

	//===============================
	//  Rental Agreement
	//===============================
	RRdb.Prepstmt.GetRentalAgreementByBusiness, err = RRdb.dbrr.Prepare("SELECT RAID,RATID,BID,PrimaryTenant,RentalStart,RentalStop,OccStart,OccStop,Renewal,SpecialProvisions,LastModTime,LastModBy from rentalagreement where BID=?")
	Errcheck(err)
	RRdb.Prepstmt.InsertRentalAgreement, err = RRdb.dbrr.Prepare("INSERT INTO rentalagreement (RATID,BID,PrimaryTenant,RentalStart,RentalStop,OccStart,OccStop,Renewal,SpecialProvisions,LastModBy) VALUES(?,?,?,?,?,?,?,?,?,?)")
	Errcheck(err)
	RRdb.Prepstmt.GetRentalAgreement, err = RRdb.dbrr.Prepare("SELECT RAID,RATID,BID,PrimaryTenant,RentalStart,RentalStop,OccStart,OccStop,Renewal,SpecialProvisions,LastModTime,LastModBy from rentalagreement WHERE RAID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetAllRentalAgreements, err = RRdb.dbrr.Prepare("SELECT RAID from rentalagreement WHERE BID=?")
	Errcheck(err)

	//===============================
	//  Rental Agreement Template
	//===============================
	RRdb.Prepstmt.GetAllRentalAgreementTemplates, err = RRdb.dbrr.Prepare("SELECT RATID,ReferenceNumber,RentalAgreementType,LastModTime,LastModBy FROM rentalagreementtemplate")
	Errcheck(err)
	RRdb.Prepstmt.GetRentalAgreementTemplate, err = RRdb.dbrr.Prepare("SELECT RATID,ReferenceNumber,RentalAgreementType,LastModTime,LastModBy FROM rentalagreementtemplate WHERE RATID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetRentalAgreementTemplateByRefNum, err = RRdb.dbrr.Prepare("SELECT RATID,ReferenceNumber,RentalAgreementType,LastModTime,LastModBy FROM rentalagreementtemplate WHERE ReferenceNumber=?")
	Errcheck(err)
	RRdb.Prepstmt.InsertRentalAgreementTemplate, err = RRdb.dbrr.Prepare("INSERT INTO rentalagreementtemplate (ReferenceNumber,RentalAgreementType,LastModBy) VALUES(?,?,?)")
	Errcheck(err)

	//===============================
	//  RentableSpecialty
	//===============================
	RRdb.Prepstmt.InsertRentableSpecialtyType, err = RRdb.dbrr.Prepare("INSERT INTO rentablespecialtytypes (RSPID,BID,Name,Fee,Description) VALUES(?,?,?,?,?)")
	Errcheck(err)
	RRdb.Prepstmt.GetSpecialtyByName, err = RRdb.dbrr.Prepare("SELECT RSPID,BID,Name,Fee,Description FROM rentablespecialtytypes WHERE BID=? and Name=?")
	Errcheck(err)
	RRdb.Prepstmt.GetRentableSpecialties, err = RRdb.dbrr.Prepare("SELECT RSPID FROM rentablespecialties WHERE BID=? and RID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetRentableSpecialty, err = RRdb.dbrr.Prepare("SELECT RSPID,BID,Name,Fee,Description FROM rentablespecialtytypes WHERE RSPID=?")
	Errcheck(err)

	//===============================
	//  Rentable Type
	//===============================
	RRdb.Prepstmt.GetRentableType, err = RRdb.dbrr.Prepare("SELECT RTID,BID,Style,Name,Accrual,Proration,Report,ManageToBudget,LastModTime,LastModBy FROM rentabletypes WHERE RTID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetRentableTypeByStyle, err = RRdb.dbrr.Prepare("SELECT RTID,BID,Style,Name,Accrual,Proration,Report,ManageToBudget,LastModTime,LastModBy FROM rentabletypes WHERE Style=? and BID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetAllBusinessRentableTypes, err = RRdb.dbrr.Prepare("SELECT RTID,BID,Style,Name,Accrual,Proration,Report,ManageToBudget,LastModTime,LastModBy FROM rentabletypes WHERE BID=?")
	Errcheck(err)
	RRdb.Prepstmt.InsertRentableType, err = RRdb.dbrr.Prepare("INSERT INTO rentabletypes (RTID,BID,Style,Name,Accrual,Proration,Report,ManageToBudget,LastModBy) VALUES(?,?,?,?,?,?,?,?,?)")
	Errcheck(err)

	//===============================
	//  RentableMarketRates
	//===============================
	RRdb.Prepstmt.GetRentableMarketRates, err = RRdb.dbrr.Prepare("SELECT RTID,MarketRate,DtStart,DtStop from rentablemarketrate WHERE RTID=?")
	Errcheck(err)
	RRdb.Prepstmt.InsertRentableMarketRates, err = RRdb.dbrr.Prepare("INSERT INTO rentablemarketrate (RTID,MarketRate,DtStart,DtStop) VALUES(?,?,?,?)")
	Errcheck(err)

	RRdb.Prepstmt.GetAgreementRentables, err = RRdb.dbrr.Prepare("SELECT RAID,RID,DtStart,DtStop from agreementrentables WHERE RID=? and ?<DtStop and ?>=DtStart")
	Errcheck(err)
	RRdb.Prepstmt.GetAgreementPayors, err = RRdb.dbrr.Prepare("SELECT RAID,PID,DtStart,DtStop from agreementpayors WHERE RAID=? and ?<DtStop and ?>=DtStart")
	Errcheck(err)
	RRdb.Prepstmt.GetAgreementTenants, err = RRdb.dbrr.Prepare("SELECT RAID,TID,DtStart,DtStop from agreementtenants WHERE RAID=? and ?<DtStop and ?>=DtStart")
	Errcheck(err)
	RRdb.Prepstmt.GetAgreementsForRentable, err = RRdb.dbrr.Prepare("SELECT RAID,RID,DtStart,DtStop from agreementrentables WHERE RID=? and ?<DtStop and ?>=DtStart")
	Errcheck(err)

	//==========================================
	// TENANT
	//==========================================
	RRdb.Prepstmt.InsertTenant, err = RRdb.dbrr.Prepare("INSERT INTO tenant (TCID,Points,CarMake,CarModel,CarColor,CarYear,LicensePlateState,LicensePlateNumber,ParkingPermitNumber,AccountRep,DateofBirth,EmergencyContactName,EmergencyContactAddress,EmergencyContactTelephone,EmergencyEmail,AlternateAddress,ElibigleForFutureOccupancy,Industry,Source,InvoicingCustomerNumber,LastModBy) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	Errcheck(err)
	RRdb.Prepstmt.GetTenant, err = RRdb.dbrr.Prepare("SELECT TID,TCID,Points,CarMake,CarModel,CarColor,CarYear,LicensePlateState,LicensePlateNumber,ParkingPermitNumber,AccountRep,DateofBirth,EmergencyContactName,EmergencyContactAddress,EmergencyContactTelephone,EmergencyEmail,AlternateAddress,ElibigleForFutureOccupancy,Industry,Source,InvoicingCustomerNumber,LastModTime,LastModBy FROM tenant where TID=?")
	Errcheck(err)

	//==========================================
	// TRANSACTANT
	//==========================================
	RRdb.Prepstmt.InsertTransactant, err = RRdb.dbrr.Prepare("INSERT INTO transactant (TID,PID,PRSPID,FirstName,MiddleName,LastName,CompanyName,IsCompany,PrimaryEmail,SecondaryEmail,WorkPhone,CellPhone,Address,Address2,City,State,PostalCode,Country,LastModBy) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	Errcheck(err)
	RRdb.Prepstmt.UpdateTransactant, err = RRdb.dbrr.Prepare("UPDATE transactant SET TID=?,PID=?,PRSPID=?,FirstName=?,MiddleName=?,LastName=?,CompanyName=?,IsCompany=?,PrimaryEmail=?,SecondaryEmail=?,WorkPhone=?,CellPhone=?,Address=?,Address2=?,City=?,State=?,PostalCode=?,Country=?,LastModBy=? WHERE TCID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetTransactant, err = RRdb.dbrr.Prepare("SELECT TCID,TID,PID,PRSPID,FirstName,MiddleName,LastName,CompanyName,IsCompany,PrimaryEmail,SecondaryEmail,WorkPhone,CellPhone,Address,Address2,City,State,PostalCode,Country,LastModTime,LastModBy FROM transactant WHERE TCID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetAllTransactants, err = RRdb.dbrr.Prepare("SELECT TCID,TID,PID,PRSPID,FirstName,MiddleName,LastName,CompanyName,IsCompany,PrimaryEmail,SecondaryEmail,WorkPhone,CellPhone,Address,Address2,City,State,PostalCode,Country,LastModTime,LastModBy FROM transactant")
	Errcheck(err)
	RRdb.Prepstmt.FindTransactantByPhoneOrEmail, err = RRdb.dbrr.Prepare("SELECT TCID,TID,PID,PRSPID,FirstName,MiddleName,LastName,CompanyName,IsCompany,PrimaryEmail,SecondaryEmail,WorkPhone,CellPhone,Address,Address2,City,State,PostalCode,Country,LastModTime,LastModBy FROM transactant where WorkPhone=? OR CellPhone=? or PrimaryEmail=? or SecondaryEmail=?")
	Errcheck(err)

}
