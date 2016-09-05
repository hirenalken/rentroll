#!/bin/bash

TESTNAME="CSV Loader Test"
TESTSUMMARY="Load all csv files through loader and validate the database after loading"

source ../share/base.sh
RRCTX="-G ${BUD} -g 12/1/15,1/1/16"

#./newbiz -b nb.csv -f rprefs.csv -n rprtrate.csv -t rpsprate.csv -l strlists.csv -R rt.csv -u custom.csv -d depository.csv -s specialties.csv -D bldg.csv -p people.csv -r rentable.csv -T rat.csv -C ra.csv -E pets.csv -a rp.csv -c coa.csv -A asmt.csv -P pmt.csv -e rcpt.csv -U assigncustom.csv -O nt.csv -m depmeth.csv -y deposit.csv -S sources.csv >${LOGFILE} 2>&1

mysqlverify "a"  "-b nb.csv"           "NewBusinesses"	            "select BID,BUD,Name,DefaultRentCycle,DefaultProrationCycle,DefaultGSRPC,LastModBy from Business;"
mysqlverify "b"  "-l strlists.csv"     "StringLists"	            "select SLID,BID,Name,LastModBy from StringList;"
mysqlverify "c"  " "	               "SLString"	            "select SLSID,SLID,Value,LastModBy from SLString;"
mysqlverify "d"  "-R rt.csv"           "RentableTypes"	            "select RTID,BID,Style,Name,RentCycle,Proration,GSRPC,ManageToBudget,LastModBy from RentableTypes;"
mysqlverify "e"  " "                   "RentableMarketRates"	    "select * from RentableMarketRate;"
mysqlverify "f"  "-m depmeth.csv"      "Deposit Methods"            "select * from DepositMethod;"
mysqlverify "g"  "-S sources.csv"      "Sources"	            "select SourceSLSID,BID,Name,Industry from DemandSource;"
mysqlverify "h"  "-s specialties.csv"  "RentableSpecialtyTypes"     "select * from RentableSpecialty;"
mysqlverify "i"  "-D bldg.csv"         "Buildings"	            "select BLDGID,BID,Address,Address2,City,State,PostalCode,Country,LastModBy from Building;"
mysqlverify "j"  "-d depository.csv"   "Depositories"	            "select DEPID,BID,Name,AccountNo,LastModBy from Depository;"
mysqlverify "n"  "-p people.csv"       "Transactants"	            "select TCID,BID,FirstName,MiddleName,LastName,CompanyName,IsCompany,PrimaryEmail,SecondaryEmail,WorkPhone,CellPhone,Address,Address2,City,State,PostalCode,Country,LastModBy from Transactant;"
mysqlverify "o"  " "                   "Users"	                    "select TCID,Points,CarMake,CarModel,CarColor,CarYear,LicensePlateState,LicensePlateNumber,ParkingPermitNumber,DateofBirth,EmergencyContactName,EmergencyContactAddress,EmergencyContactTelephone,EmergencyEmail,AlternateAddress,EligibleFutureUser,Industry,SourceSLSID from User;"
mysqlverify "p"  " "                   "Payors"	                    "select TCID,CreditLimit,TaxpayorID,AccountRep,LastModBy from Payor;"
mysqlverify "q"  " "                   "Prospects"	            "select TCID,EmployerName,EmployerStreetAddress,EmployerCity,EmployerState,EmployerPostalCode,EmployerEmail,EmployerPhone,Occupation,ApplicationFee,LastModBy from Prospect;"
mysqlverify "k"  "-r rentable.csv"     "Rentables"	            "select RID,BID,Name,AssignmentTime,LastModBy from Rentable;"
mysqlverify "l"  " "                   "RentableTypeRef"	    "select RID,RTID,RentCycle,ProrationCycle,DtStart,DtStop,LastModBy from RentableTypeRef;"
mysqlverify "m"  " "                   "RentableStatus"	            "select RID,Status,DtStart,DtStop,LastModBy from RentableStatus;"
mysqlverify "r"  "-T rat.csv"          "RentalAgreementTemplates"   "select RATID,BID,RATemplateName,LastModBy from RentalAgreementTemplate;"
mysqlverify "s"  "-C ra.csv"           "RentalAgreements"	    "select RAID,RATID,BID,AgreementStart,AgreementStop,Renewal,SpecialProvisions,LastModBy from RentalAgreement;"
mysqlverify "t"  "-E pets.csv"         "Pets"	                    "select PETID,RAID,Type,Breed,Color,Weight,Name,DtStart,DtStop,LastModBy from RentalAgreementPets;"
mysqlverify "u"  ""           "Notes"	                    "select NID,PNID,Comment,LastModBy from Notes;"
mysqlverify "v"  " "                   "AgreementRentables"	    "select * from RentalAgreementRentables;"
mysqlverify "w"  " "                   "AgreementPayors"	    "select * from RentalAgreementPayors;"
mysqlverify "x"  "-c coa.csv"          "ChartOfAccounts"	    "select LID,PLID,BID,RAID,GLNumber,Status,Type,Name,AcctType,RAAssociated,AllowPost,LastModBy from GLAccount;"
mysqlverify "y"  " "                   "LedgerMarkers"	            "select LMID,LID,BID,Dt,Balance,State,LastModBy from LedgerMarker;"
mysqlverify "z"  "-a rp.csv"           "RatePlan"	            "select RPID,BID,Name,LastModBy from RatePlan;"
mysqlverify "a1" "-f rprefs.csv"       "RatePlanRef"	            "select RPRID,RPID,DtStart,DtStop,FeeAppliesAge,MaxNoFeeUsers,AdditionalUserFee,PromoCode,CancellationFee,FLAGS,LastModBy from RatePlanRef;"
mysqlverify "b1" "-n rprtrate.csv"     "RatePlanRefRTRate"	    "select * from RatePlanRefRTRate;"
mysqlverify "c1" "-t rpsprate.csv"     "RatePlanRefSPRate"	    "select * from RatePlanRefSPRate;"
mysqlverify "d1" "-A asmt.csv ${RRCTX}"         "Assessments"	            "select ASMID,BID,RID,ATypeLID,RAID,Amount,Start,Stop,RentCycle,ProrationCycle,AcctRule,Comment,LastModBy from Assessments;"
mysqlverify "e1" "-P pmt.csv"          "PaymentTypes"	            "select PMTID,BID,Name,Description,LastModBy from PaymentTypes;"
mysqlverify "f1" "-e rcpt.csv ${RRCTX}"         "PaymentAllocations"	    "select * from ReceiptAllocation order by Amount ASC;"
mysqlverify "g1" " "                   "Receipts"	            "select RCPTID,BID,RAID,PMTID,Dt,Amount,AcctRule,Comment,LastModBy from Receipt;"
mysqlverify "h1" "-u custom.csv"       "CustomAttributes"	    "select CID,Type,Name,Value,LastModBy from CustomAttr;"
mysqlverify "i1" "-U assigncustom.csv" "CustomAttributesAssignment" "select * from CustomAttrRef;"
mysqlverify "j1" "-O nt.csv"           "NoteTypes"	            "select NTID,BID,Name,LastModBy from NoteType;"
mysqlverify "k1" "-y deposit.csv"      "Deposits"	            "select DID,BID,Dt,DEPID,Amount,LastModBy from Deposit;"

logcheck
