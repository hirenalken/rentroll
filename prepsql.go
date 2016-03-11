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

	App.prepstmt.rentalAgreementByBusiness, err = App.dbrr.Prepare("SELECT RAID,RATID,BID,RID,UNITID,PID,PrimaryTenant,RentalStart,RentalStop,Renewal,SpecialProvisions,LastModTime,LastModBy from rentalagreement where BID=?")
	rlib.Errcheck(err)
	App.prepstmt.getUnit, err = App.dbrr.Prepare("SELECT UNITID,BLDGID,UTID,RID,AVAILID,LastModTime,LastModBy FROM unit where UNITID=?")
	rlib.Errcheck(err)
	App.prepstmt.getLedger, err = App.dbrr.Prepare("SELECT LID,GLNumber,Dt,Balance FROM ledger where LID=?")
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
	App.prepstmt.getRentableType, err = App.dbrr.Prepare("SELECT RTID,BID,Name,Amount,Frequency,Proration,LastModTime,LastModBy FROM rentabletypes where RTID=?")
	rlib.Errcheck(err)
	App.prepstmt.getUnitType, err = App.dbrr.Prepare("SELECT UTID,BID,Style,Name,SqFt,MarketRate,Frequency,Proration,LastModTime,LastModBy FROM unittypes where UTID=?")
	rlib.Errcheck(err)
	App.prepstmt.getUnitReceipts, err = App.dbrr.Prepare("SELECT RCPTID,PID,PMTID,Amount,Dt FROM receipt WHERE RAID=? and Dt>=? and Dt<?")
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
	App.prepstmt.getUnitRentalAgreements, err = App.dbrr.Prepare("SELECT RAID,RATID,BID,RID,UNITID,PID,PrimaryTenant,RentalStart,RentalStop,Renewal,SpecialProvisions,LastModTime,LastModBy from rentalagreement where unitid=? and RentalStop > ? and RentalStart < ?")
	rlib.Errcheck(err)
	App.prepstmt.getAllRentablesByBusiness, err = App.dbrr.Prepare("SELECT RID,LID,RTID,BID,UNITID,Name,Assignment,Report,LastModTime,LastModBy FROM rentable WHERE BID=?")
	rlib.Errcheck(err)
	App.prepstmt.getAllBusinessRentableTypes, err = App.dbrr.Prepare("SELECT RTID,BID,Name,Amount,Frequency,Proration,LastModTime,LastModBy FROM rentabletypes WHERE BID=?")
	rlib.Errcheck(err)
	App.prepstmt.getAllBusinessUnitTypes, err = App.dbrr.Prepare("SELECT UTID,BID,Style,Name,SqFt,MarketRate,Frequency,Proration,LastModTime,LastModBy FROM unittypes WHERE BID=?")
	rlib.Errcheck(err)
	App.prepstmt.getBusiness, err = App.dbrr.Prepare("SELECT BID,Address,Address2,City,State,PostalCode,Country,Phone,Name,DefaultOccupancyType,ParkingPermitInUse,LastModTime,LastModBy from business where bid=?")
	rlib.Errcheck(err)
	App.prepstmt.getAllBusinessSpecialtyTypes, err = App.dbrr.Prepare("SELECT USPID,BID,Name,Fee,Description FROM unitspecialtytypes WHERE BID=?")
	rlib.Errcheck(err)
	App.prepstmt.getAllAssessmentsByBusiness, err = App.dbrr.Prepare("SELECT ASMID,BID,RID,UNITID,ASMTID,RAID,Amount,Start,Stop,Frequency,ProrationMethod,AcctRule,LastModTime,LastModBy FROM assessments WHERE BID=? and Start<? and Stop>?")
	rlib.Errcheck(err)
	App.prepstmt.getLedgerByGLNo, err = App.dbrr.Prepare("SELECT LID,GLNumber,Dt,Balance,Name FROM ledger WHERE GLNumber=?")
	rlib.Errcheck(err)
	App.prepstmt.getRentalAgreement, err = App.dbrr.Prepare("SELECT RAID,RATID,BID,RID,UNITID,PID,PrimaryTenant,RentalStart,RentalStop,Renewal,SpecialProvisions,LastModTime,LastModBy from rentalagreement where RAID=?")
	rlib.Errcheck(err)
	App.prepstmt.getReceiptsInDateRange, err = App.dbrr.Prepare("SELECT RCPTID,BID,PID,RAID,PMTID,Dt,Amount from receipt where BID=? and Dt >= ? and DT < ?")
	rlib.Errcheck(err)
	App.prepstmt.getReceiptAllocations, err = App.dbrr.Prepare("SELECT RCPTID,Amount,ASMID from receiptallocation where RCPTID=?")
	rlib.Errcheck(err)
	App.prepstmt.getAssessment, err = App.dbrr.Prepare("SELECT ASMID, BID, RID, UNITID, ASMTID, RAID, Amount, Start, Stop, Frequency, ProrationMethod, AcctRule, LastModTime, LastModBy from assessments where ASMID=?")
	rlib.Errcheck(err)
}
