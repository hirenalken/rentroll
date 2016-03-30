package main

import (
	"fmt"
	"rentroll/rlib"
)

func buildPreparedStatements() {
	var err error
	// Prepare("select deduction from deductions where uid=?")
	// Prepare("select type from compensation where uid=?")
	// Prepare("INSERT INTO compensation (uid,type) VALUES(?,?)")
	// Prepare("DELETE FROM compensation WHERE UID=?")
	// Prepare("update classes set Name=?,Designation=?,Description=?,lastmodby=? where ClassCode=?")
	// rlib.Errcheck(err)

	App.prepstmt.rentalAgreementByBusiness, err = App.dbrr.Prepare("SELECT RAID,RATID,BID,RID,UNITID,PID,LID,PrimaryTenant,RentalStart,RentalStop,Renewal,SpecialProvisions,LastModTime,LastModBy from rentalagreement where BID=?")
	rlib.Errcheck(err)
	App.prepstmt.getUnit, err = App.dbrr.Prepare("SELECT UNITID,BLDGID,UTID,RID,AVAILID,LastModTime,LastModBy FROM unit where UNITID=?")
	rlib.Errcheck(err)
	App.prepstmt.getTransactant, err = App.dbrr.Prepare("SELECT TCID,TID,PID,PRSPID,FirstName,MiddleName,LastName,PrimaryEmail,SecondaryEmail,WorkPhone,CellPhone,Address,Address2,City,State,PostalCode,Country,LastModTime,LastModBy FROM transactant WHERE TCID=?")
	rlib.Errcheck(err)
	App.prepstmt.getTenant, err = App.dbrr.Prepare("SELECT TID,TCID,Points,CarMake,CarModel,CarColor,CarYear,LicensePlateState,LicensePlateNumber,ParkingPermitNumber,AccountRep,DateofBirth,EmergencyContactName,EmergencyContactAddress,EmergencyContactTelephone,EmergencyAddressEmail,AlternateAddress,ElibigleForFutureOccupancy,Industry,Source,InvoicingCustomerNumber FROM tenant where TID=?")
	rlib.Errcheck(err)
	App.prepstmt.getRentable, err = App.dbrr.Prepare("SELECT RID,LID,RTID,BID,UNITID,Name,Assignment,Report,DefaultOccType,OccType,LastModTime,LastModBy FROM rentable where RID=?")
	rlib.Errcheck(err)
	App.prepstmt.getProspect, err = App.dbrr.Prepare("SELECT PRSPID,TCID,ApplicationFee FROM prospect where PRSPID=?")
	rlib.Errcheck(err)
	App.prepstmt.getPayor, err = App.dbrr.Prepare("SELECT PID,TCID,CreditLimit,EmployerName,EmployerStreetAddress,EmployerCity,EmployerState,EmployerZipcode,Occupation,LastModTime,LastModBy FROM payor where PID=?")
	rlib.Errcheck(err)
	App.prepstmt.getUnitSpecialties, err = App.dbrr.Prepare("SELECT USPID FROM unitspecialties where BID=? and UNITID=?")
	rlib.Errcheck(err)
	App.prepstmt.getUnitSpecialtyType, err = App.dbrr.Prepare("SELECT USPID,BID,Name,Fee,Description FROM unitspecialtytypes where USPID=?")
	rlib.Errcheck(err)
	App.prepstmt.getRentableType, err = App.dbrr.Prepare("SELECT RTID,BID,Name,Frequency,Proration,Report,ManageToBudget,LastModTime,LastModBy FROM rentabletypes where RTID=?")
	rlib.Errcheck(err)
	App.prepstmt.getUnitType, err = App.dbrr.Prepare("SELECT UTID,BID,Style,Name,SqFt,Frequency,Proration,LastModTime,LastModBy FROM unittypes where UTID=?")
	rlib.Errcheck(err)
	App.prepstmt.getUnitReceipts, err = App.dbrr.Prepare("SELECT RCPTID,BID,RAID,PMTID,Dt,Amount,AcctRule FROM receipt WHERE RAID=? and Dt>=? and Dt<?")
	rlib.Errcheck(err)
	App.prepstmt.getReceipt, err = App.dbrr.Prepare("SELECT RCPTID,BID,RAID,PMTID,Dt,Amount,AcctRule FROM receipt WHERE RCPTID=?")
	rlib.Errcheck(err)
	App.prepstmt.getUnitAssessments, err = App.dbrr.Prepare("SELECT ASMID,BID,RID,UNITID,ASMTID,RAID,Amount,Start,Stop,Frequency,ProrationMethod,AcctRule,LastModTime,LastModBy FROM assessments WHERE UNITID=? and Stop >= ? and Start < ?")
	rlib.Errcheck(err)
	App.prepstmt.getAllRentableAssessments, err = App.dbrr.Prepare("SELECT ASMID,BID,RID,UNITID,ASMTID,RAID,Amount,Start,Stop,Frequency,ProrationMethod,AcctRule,LastModTime,LastModBy FROM assessments WHERE RID=? and Stop >= ? and Start < ?")
	rlib.Errcheck(err)
	App.prepstmt.getAssessmentType, err = App.dbrr.Prepare("SELECT ASMTID,Name,Type,LastModTime,LastModBy FROM assessmenttypes WHERE ASMTID=?")
	rlib.Errcheck(err)
	s := fmt.Sprintf("SELECT ASMID,BID,RID,UNITID,ASMTID,RAID,Amount,Start,Stop,Frequency,ProrationMethod,AcctRule,LastModTime,LastModBy FROM assessments WHERE (ASMTID=%d or ASMTID=%d) and UNITID=?", SECURITYDEPOSIT, SECURITYDEPOSITASSESSMENT)
	App.prepstmt.getSecurityDepositAssessment, err = App.dbrr.Prepare(s)
	rlib.Errcheck(err)
	App.prepstmt.getUnitRentalAgreements, err = App.dbrr.Prepare("SELECT RAID,RATID,BID,RID,UNITID,PID,LID,PrimaryTenant,RentalStart,RentalStop,Renewal,SpecialProvisions,LastModTime,LastModBy from rentalagreement where unitid=? and RentalStop > ? and RentalStart < ?")
	rlib.Errcheck(err)
	App.prepstmt.getAllRentablesByBusiness, err = App.dbrr.Prepare("SELECT RID,LID,RTID,BID,UNITID,Name,Assignment,Report,LastModTime,LastModBy FROM rentable WHERE BID=?")
	rlib.Errcheck(err)
	App.prepstmt.getAllBusinessRentableTypes, err = App.dbrr.Prepare("SELECT RTID,BID,Name,Frequency,Proration,Report,ManageToBudget,LastModTime,LastModBy FROM rentabletypes WHERE BID=?")
	rlib.Errcheck(err)
	App.prepstmt.getAllBusinessUnitTypes, err = App.dbrr.Prepare("SELECT UTID,BID,Style,Name,SqFt,Frequency,Proration,LastModTime,LastModBy FROM unittypes WHERE BID=?")
	rlib.Errcheck(err)
	App.prepstmt.getBusiness, err = App.dbrr.Prepare("SELECT BID,Address,Address2,City,State,PostalCode,Country,Phone,Name,DefaultOccupancyType,ParkingPermitInUse,LastModTime,LastModBy from business where bid=?")
	rlib.Errcheck(err)
	App.prepstmt.getAllBusinessSpecialtyTypes, err = App.dbrr.Prepare("SELECT USPID,BID,Name,Fee,Description FROM unitspecialtytypes WHERE BID=?")
	rlib.Errcheck(err)
	App.prepstmt.getAllAssessmentsByBusiness, err = App.dbrr.Prepare("SELECT ASMID,BID,RID,UNITID,ASMTID,RAID,Amount,Start,Stop,Frequency,ProrationMethod,AcctRule,LastModTime,LastModBy FROM assessments WHERE BID=? and Start<? and Stop>?")
	rlib.Errcheck(err)
	App.prepstmt.getAllJournalsInRange, err = App.dbrr.Prepare("SELECT JID,BID,RAID,Dt,Amount,Type,ID from journal WHERE BID=? and ?<=Dt and Dt<?")
	rlib.Errcheck(err)
	App.prepstmt.getRentalAgreement, err = App.dbrr.Prepare("SELECT RAID,RATID,BID,RID,UNITID,PID,LID,PrimaryTenant,RentalStart,RentalStop,Renewal,SpecialProvisions,LastModTime,LastModBy from rentalagreement where RAID=?")
	rlib.Errcheck(err)
	App.prepstmt.getReceiptsInDateRange, err = App.dbrr.Prepare("SELECT RCPTID,BID,RAID,PMTID,Dt,Amount,AcctRule from receipt where BID=? and Dt >= ? and DT < ?")
	rlib.Errcheck(err)
	// App.prepstmt.getDefaultCashLedgerMarker, err = App.dbrr.Prepare("SELECT LMID,BID,GLNumber,State,Dt,Balance,Type,Name FROM ledgermarker where BID=? and DefaultAcct=1")
	// rlib.Errcheck(err)
	App.prepstmt.getReceiptAllocations, err = App.dbrr.Prepare("SELECT RCPTID,Amount,ASMID,AcctRule from receiptallocation where RCPTID=?")
	rlib.Errcheck(err)
	App.prepstmt.getJournalAllocation, err = App.dbrr.Prepare("SELECT JAID,JID,Amount,ASMID,AcctRule from journalallocation WHERE JAID=?")
	rlib.Errcheck(err)
	App.prepstmt.getJournalAllocations, err = App.dbrr.Prepare("SELECT JAID,JID,Amount,ASMID,AcctRule from journalallocation WHERE JID=?")
	rlib.Errcheck(err)
	App.prepstmt.getRentableMarketRates, err = App.dbrr.Prepare("SELECT RTID,MarketRate,DtStart,DtStop from rentablemarketrate WHERE RTID=?")
	rlib.Errcheck(err)
	App.prepstmt.getUnitMarketRates, err = App.dbrr.Prepare("SELECT UTID,MarketRate,DtStart,DtStop from unitmarketrate WHERE UTID=?")
	rlib.Errcheck(err)
	App.prepstmt.getAssessment, err = App.dbrr.Prepare("SELECT ASMID, BID, RID, UNITID, ASMTID, RAID, Amount, Start, Stop, Frequency, ProrationMethod, AcctRule, LastModTime, LastModBy from assessments where ASMID=?")
	rlib.Errcheck(err)
	App.prepstmt.getJournalMarker, err = App.dbrr.Prepare("SELECT JMID,BID,State,DtStart,DtStop from journalmarker where JMID=?")
	rlib.Errcheck(err)
	App.prepstmt.getJournalMarkers, err = App.dbrr.Prepare("SELECT JMID,BID,State,DtStart,DtStop from journalmarker ORDER BY JMID DESC LIMIT ?")
	rlib.Errcheck(err)
	App.prepstmt.getJournal, err = App.dbrr.Prepare("select JID,BID,RAID,Dt,Amount,Type,ID from journal where JID=?")
	rlib.Errcheck(err)

	App.prepstmt.insertJournal, err = App.dbrr.Prepare("INSERT INTO journal (BID,RAID,Dt,Amount,Type,ID) VALUES(?,?,?,?,?,?)")
	rlib.Errcheck(err)
	App.prepstmt.insertJournalAllocation, err = App.dbrr.Prepare("INSERT INTO journalallocation (JID,Amount,ASMID,AcctRule) VALUES(?,?,?,?)")
	rlib.Errcheck(err)
	App.prepstmt.insertJournalMarker, err = App.dbrr.Prepare("INSERT INTO journalmarker (BID,State,DtStart,DtStop) VALUES(?,?,?,?)")
	rlib.Errcheck(err)

	App.prepstmt.deleteJournalAllocations, err = App.dbrr.Prepare("DELETE FROM journalallocation WHERE JID=?")
	rlib.Errcheck(err)
	App.prepstmt.deleteJournalEntry, err = App.dbrr.Prepare("DELETE FROM journal WHERE JID=?")
	rlib.Errcheck(err)
	App.prepstmt.deleteJournalMarker, err = App.dbrr.Prepare("DELETE FROM journalmarker WHERE JMID=?")
	rlib.Errcheck(err)

	App.prepstmt.getLedgerMarkerByGLNo, err = App.dbrr.Prepare("SELECT LMID,BID,GLNumber,Status,State,DtStart,DtStop,Balance,Type,Name FROM ledgermarker WHERE BID=? and GLNumber=?")
	rlib.Errcheck(err)
	App.prepstmt.getLedgerMarkers, err = App.dbrr.Prepare("SELECT LMID,BID,GLNumber,Status,State,DtStart,DtStop,Balance,Type,Name from ledgermarker WHERE BID=? ORDER BY LMID DESC LIMIT ?")
	rlib.Errcheck(err)
	App.prepstmt.getAllLedgerMarkersInRange, err = App.dbrr.Prepare("SELECT LMID,BID,GLNumber,Status,State,DtStart,DtStop,Balance,Type,Name from ledgermarker WHERE BID=? and DtStart<=? and DtStop<?")
	rlib.Errcheck(err)

	App.prepstmt.getAllLedgersInRange, err = App.dbrr.Prepare("SELECT LID,BID,JID,JAID,GLNumber,Dt,Amount from ledger WHERE BID=? and ?<=Dt and Dt<?")
	rlib.Errcheck(err)
	App.prepstmt.getLedgerInRangeByGLNo, err = App.dbrr.Prepare("SELECT LID,BID,JID,JAID,GLNumber,Dt,Amount from ledger WHERE BID=? and GLNumber=? and ?<=Dt and Dt<?")
	rlib.Errcheck(err)
	App.prepstmt.getLedger, err = App.dbrr.Prepare("SELECT LID,BID,JID,JAID,GLNumber,Dt,Amount FROM ledger where LID=?")
	rlib.Errcheck(err)

	App.prepstmt.deleteLedgerEntry, err = App.dbrr.Prepare("DELETE FROM ledger WHERE LID=?")
	rlib.Errcheck(err)
	App.prepstmt.deleteLedgerMarker, err = App.dbrr.Prepare("DELETE FROM ledgermarker WHERE LMID=?")
	rlib.Errcheck(err)

	App.prepstmt.insertLedger, err = App.dbrr.Prepare("INSERT INTO ledger (BID,JID,JAID,GLNumber,Dt,Amount) VALUES(?,?,?,?,?,?)")
	rlib.Errcheck(err)
	App.prepstmt.insertLedgerMarker, err = App.dbrr.Prepare("INSERT INTO ledgermarker (LMID,BID,GLNumber,Status,State,DtStart,DtStop,Balance,Type,Name) VALUES(?,?,?,?,?,?,?,?,?,?)")
	rlib.Errcheck(err)
}
