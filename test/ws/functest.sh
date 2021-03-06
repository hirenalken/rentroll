#!/bin/bash

#---------------------------------------------------------------
# TOP is the directory where RentRoll begins. It is used
# in base.sh to set other useful directories such as ${BASHDIR}
#---------------------------------------------------------------
TOP=../..

TESTNAME="Web Services"
TESTSUMMARY="Test Web Services"
RRDATERANGE="-j 2016-07-01 -k 2016-08-01"

CREATENEWDB=0

#---------------------------------------------------------------
#  Use the testdb for these tests...
#---------------------------------------------------------------
echo "Create new database..."
mysql --no-defaults rentroll < restore.sql

source ../share/base.sh

echo "STARTING RENTROLL SERVER"
startRentRollServer

# Get Specificy PaymentType
# echo "request%3D%7B%22cmd%22%3A%22get%22%2C%22recid%22%3A0%2C%22name%22%3A%22paymentTypeGrid%22%7D" > request
# dojsonPOST "http://localhost:8270/v1/pmt/1/1" "request" "zz"  "WebService--PaymentTypes"

# get GLAccounts list for the business
dojsonGET "http://localhost:8270/v1/accountlist/2" "wa" "WebService--GetAccountsListForBusiness"

# get parent accounts list for the business
dojsonGET "http://localhost:8270/v1/parentaccounts/2" "wb" "WebService--GetParentAccountsListForBusiness"

# get post accounts list for the business
dojsonGET "http://localhost:8270/v1/postaccounts/2" "wc" "WebService--GetPostAccountsListForBusiness"

# Get Chart of Accounts
echo "request=%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%7D" > request
dojsonPOST "http://localhost:8270/v1/accounts/2" "request" "wd"  "WebService--ChartOfAccounts"

# Get Account details
echo "request=%7B%22cmd%22%3A%22get%22%2C%22recid%22%3A0%2C%22name%22%3A%22accountForm%22%7D" > request
dojsonPOST "http://localhost:8270/v1/account/2/97" "request" "we" "WebService--ChartOfAccounts-detail"

# dojsonPOST "http://localhost:8270/v1/account/2/97" "request" "we"  "WebService--GetAccountDetails"

# Create new Account
echo "request%3D%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22%22%2C%22record%22%3A%7B%22LID%22%3A0%2C%22BID%22%3A2%2C%22RAID%22%3A0%2C%22TCID%22%3A0%2C%22GLNumber%22%3A%22123456789%22%2C%22Name%22%3A%22SmokeTest%20GLAccount%22%2C%22AcctType%22%3A%22%22%2C%22Description%22%3A%22%22%2C%22LastModTime%22%3A%221%2F1%2F1900%22%2C%22LastModBy%22%3A0%2C%22BUD%22%3A%22%22%2C%22PLID%22%3A0%2C%22Status%22%3A0%2C%22Type%22%3A0%2C%22AllowPost%22%3A0%2C%22ManageToBudget%22%3A0%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/account/2/0" "request" "wf"  "WebService--CreateGLAccount"

# Update Account details
echo "request%3D%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22%22%2C%22record%22%3A%7B%22LID%22%3A108%2C%22BID%22%3A2%2C%22RAID%22%3A0%2C%22TCID%22%3A0%2C%22GLNumber%22%3A%229876543210%22%2C%22Name%22%3A%22SmokeTest%20GLAccount%22%2C%22AcctType%22%3A%22%22%2C%22Description%22%3A%22Update%20this%20Account%20(Smoke%20Test)%22%2C%22LastModTime%22%3A%221%2F1%2F1900%22%2C%22LastModBy%22%3A0%2C%22BUD%22%3A%22%22%2C%22PLID%22%3A0%2C%22Status%22%3A0%2C%22Type%22%3A0%2C%22AllowPost%22%3A1%2C%22ManageToBudget%22%3A0%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/account/2/108" "request" "wg"  "WebService--UpdateGLAccount"

# Delete Account
echo "request%3D%7B%22cmd%22%3A%22delete%22%2C%22LID%22%3A97%7D" > request
dojsonPOST "http://localhost:8270/v1/account/2/" "request" "wh"  "WebService--DeleteGLAccount"

# Get Transactants
echo "request=%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%2C%22sort%22%3A%5B%7B%22field%22%3A%22LastName%22%2C%22direction%22%3A%22asc%22%7D%5D%7D" > request
dojsonPOST "http://localhost:8270/v1/transactants/1" "request" "b"  "WebService--GetTransactants"

# Get Rentables
echo "request=%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%7D" > request
dojsonPOST "http://localhost:8270/v1/rentables/1" "request" "c"  "WebService--GetRentables"

# Get Receipts
# echo "request%3d%7b%22cmd%22%3a%22get%22%2c%22selected%22%3a%5b%5d%2c%22limit%22%3a100%2c%22offset%22%3a0%7d" > request
echo "request=%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%2C%22searchDtStart%22%3A%222016-08-01%22%2C%22searchDtStop%22%3A%222016-09-01%22%7D" > request
dojsonPOST "http://localhost:8270/v1/receipts/1" "request" "d"  "WebService--GetReceipts"

# Get Assessments
echo "request%3d%7b%22cmd%22%3a%22get%22%2c%22selected%22%3a%5b%5d%2c%22limit%22%3a100%2c%22offset%22%3a0%7d" > request
dojsonPOST "http://localhost:8270/v1/asms/1" "request" "e"  "WebService--GetAssessments"

# Get Assessment 1 from REX
echo "request=%7B%22cmd%22%3A%22get%22%2C%22recid%22%3A0%2C%22name%22%3A%22asmEpochForm%22%7D" > request
dojsonPOST "http://localhost:8270/v1/asm/REX/1" "request" "f"  "WebService--GetAnAssessment"

# Save the Assessment with an updated comment
echo "request=%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22asmInstForm%22%2C%22record%22%3A%7B%22recid%22%3A0%2C%22ASMID%22%3A43%2C%22BID%22%3A1%2C%22BUD%22%3A%22REX%22%2C%22PASMID%22%3A14%2C%22RID%22%3A1%2C%22Rentable%22%3A%22309%20Rexford%22%2C%22RAID%22%3A1%2C%22Amount%22%3A3750%2C%22Start%22%3A%2212%2F1%2F2016%22%2C%22Stop%22%3A%2212%2F2%2F2016%22%2C%22RentCycle%22%3A6%2C%22ProrationCycle%22%3A4%2C%22InvoiceNo%22%3A0%2C%22ARID%22%3A0%2C%22Comment%22%3A%22comment%20by%20sman%22%2C%22LastModTime%22%3A%226%2F6%2F2017%22%2C%22LastModBy%22%3A0%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/asm/REX/1" "request" "g"  "WebService--SaveAnAssessment"

# Get Receipt 5 from REX
echo "%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%7D" > request
dojsonPOST "http://localhost:8270/v1/receipt/REX/5" "request" "h"  "WebService--GetAReceipt"

# Save the Receipt 5 with an updated comment
# echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22receiptForm%22%2C%22record%22%3A%7B%22AcctRule%22%3A%22%22%2C%22Amount%22%3A3550%2C%22BID%22%3A1%2C%22BUD%22%3A%22REX%22%2C%22Comment%22%3A%22This%20comment%20was%20updated%20by%20a%20web%20service%20test%22%2C%22DocNo%22%3A%221631%22%2C%22Dt%22%3A%221%2F4%2F2016%22%2C%22LastModBy%22%3A0%2C%22LastModTime%22%3A%222%2F23%2F2017%22%2C%22OtherPayorName%22%3A%22%22%2C%22PMTID%22%3A1%2C%22PRCPTID%22%3A0%2C%22RAID%22%3A2%2C%22RCPTID%22%3A5%2C%22recid%22%3A0%7D%7D" > request
echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22receiptForm%22%2C%22record%22%3A%7B%22recid%22%3A0%2C%22RCPTID%22%3A5%2C%22PRCPTID%22%3A0%2C%22BID%22%3A1%2C%22BUD%22%3A%22REX%22%2C%22PMTID%22%3A1%2C%22Payor%22%3A%22Rita+Costea+(TCID%3A+3)%22%2C%22TCID%22%3A3%2C%22Dt%22%3A%221%2F4%2F2016%22%2C%22DocNo%22%3A%221631%22%2C%22Amount%22%3A3550%2C%22ARID%22%3A0%2C%22Comment%22%3A%22update+comment%22%2C%22OtherPayorName%22%3A%22%22%2C%22LastModTime%22%3A%227%2F18%2F2017%22%2C%22LastModBy%22%3A0%2C%22FLAGS%22%3A0%2C%22PmtTypeName%22%3A1%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/receipt/REX/5" "request" "i"  "WebService--SaveAReceipt"

# Create a NEW RECEIPT
# echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22receiptForm%22%2C%22record%22%3A%7B%22AcctRule%22%3A%22%22%2C%22Amount%22%3A1590.32%2C%22BID%22%3A1%2C%22BUD%22%3A%22REX%22%2C%22Comment%22%3A%22This%20is%20a%20NEW%20RECEIPT%20added%20by%20a%20web%20test%22%2C%22DocNo%22%3A%229876%22%2C%22Dt%22%3A%222%2F24%2F2017%22%2C%22LastModBy%22%3A0%2C%22LastModTime%22%3A%222%2F24%2F2017%22%2C%22OtherPayorName%22%3A%22%22%2C%22PMTID%22%3A1%2C%22PRCPTID%22%3A0%2C%22RAID%22%3A2%2C%22RCPTID%22%3A0%2C%22recid%22%3A0%7D%7D" > request
echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22receiptForm%22%2C%22record%22%3A%7B%22recid%22%3A0%2C%22RCPTID%22%3A0%2C%22PRCPTID%22%3A0%2C%22ARID%22%3A3%2C%22BID%22%3A1%2C%22BUD%22%3A%22REX%22%2C%22PMTID%22%3A2%2C%22PmtTypeName%22%3A2%2C%22Dt%22%3A%227%2F18%2F2017%22%2C%22DocNo%22%3A%222345%22%2C%22Payor%22%3A%22Aaron+Read+(TCID%3A+1)%22%2C%22TCID%22%3A1%2C%22Amount%22%3A3750%2C%22Comment%22%3A%22rent%22%2C%22OtherPayorName%22%3A%22%22%2C%22FLAGS%22%3A0%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/receipt/REX/0" "request" "j"  "WebService--InsertAReceipt"

# This receipt should FAIL -- TCID is 0
echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22receiptForm%22%2C%22record%22%3A%7B%22recid%22%3A0%2C%22RCPTID%22%3A0%2C%22PRCPTID%22%3A0%2C%22ARID%22%3A3%2C%22BID%22%3A1%2C%22BUD%22%3A%22REX%22%2C%22PMTID%22%3A2%2C%22PmtTypeName%22%3A2%2C%22Dt%22%3A%227%2F19%2F2017%22%2C%22DocNo%22%3A%222345%22%2C%22Payor%22%3A%22%22%2C%22TCID%22%3A0%2C%22Amount%22%3A40%2C%22Comment%22%3A%22%22%2C%22OtherPayorName%22%3A%22%22%2C%22FLAGS%22%3A0%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/receipt/REX/0" "request" "j1"  "WebService--InsertAReceipt-Fai"

# Create a NEW ASSESSMENT
echo "request=%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22asmEpochForm%22%2C%22record%22%3A%7B%22ARID%22%3A2%2C%22recid%22%3A0%2C%22RID%22%3A1%2C%22ASMID%22%3A0%2C%22PASMID%22%3A0%2C%22ATypeLID%22%3A0%2C%22InvoiceNo%22%3A0%2C%22RAID%22%3A1%2C%22BID%22%3A1%2C%22BUD%22%3A%22REX%22%2C%22Start%22%3A%226%2F6%2F2017%22%2C%22Stop%22%3A%226%2F6%2F2017%22%2C%22RentCycle%22%3A6%2C%22ProrationCycle%22%3A4%2C%22TCID%22%3A0%2C%22Amount%22%3A60%2C%22AcctRule%22%3A%22%22%2C%22Comment%22%3A%22%22%2C%22LastModTime%22%3A%226%2F6%2F2017%22%2C%22LastModBy%22%3A0%2C%22Rentable%22%3A%22309%2BRexford%22%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/asm/1/0" "request" "k"  "WebService--InsertAnAssessment"

# Test Transactant Typedown
dojsonGET "http://localhost:8270/v1/transactantstd/ISO?request=%7B%22search%22%3A%22s%22%2C%22max%22%3A250%7D" "l" "WebService--GetTransactantTypeDown"

# Create a NEW Rentable User
echo "%7B%22cmd%22%3A%22save%22%2C%22formname%22%3A%22tcidPicker%22%2C%22record%22%3A%7B%22recid%22%3A1%2C%22BID%22%3A5%2C%22TCID%22%3A373%2C%22RID%22%3A16%2C%22FirstName%22%3A%22Jason%22%2C%22MiddleName%22%3A%22%22%2C%22LastName%22%3A%22Thomas%22%2C%22DtStart%22%3A%223%2F5%2F2017%22%2C%22DtStop%22%3A%223%2F5%2F2018%22%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/ruser/ISO/16" "request" "m"  "WebService--InsertARentableUser"

# Create another NEW Rentable User -- same TCID
echo "%7B%22cmd%22%3A%22save%22%2C%22formname%22%3A%22tcidPicker%22%2C%22record%22%3A%7B%22recid%22%3A1%2C%22BID%22%3A5%2C%22TCID%22%3A373%2C%22RID%22%3A16%2C%22FirstName%22%3A%22Jason%22%2C%22MiddleName%22%3A%22%22%2C%22LastName%22%3A%22Thomas%22%2C%22DtStart%22%3A%223%2F5%2F2017%22%2C%22DtStop%22%3A%223%2F5%2F2018%22%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/ruser/ISO/16" "request" "n"  "WebService--InsertARentableUser"

# Delete a Rentable User
echo "request%3D%7B%22cmd%22%3A%22delete%22%2C%22selected%22%3A%5B1%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%2C%22TCID%22%3A373%7D" > request
dojsonPOST "http://localhost:8270/v1/ruser/ISO/16" "request" "o"  "WebService--DeleteARentableUser"

# Create another NEW RAID Payor -- same TCID
echo "%7B%22cmd%22%3A%22save%22%2C%22formname%22%3A%22tcidPicker%22%2C%22record%22%3A%7B%22recid%22%3A1%2C%22BID%22%3A5%2C%22TCID%22%3A367%2C%22RAID%22%3A16%2C%22FirstName%22%3A%22Eric%22%2C%22MiddleName%22%3A%22%22%2C%22LastName%22%3A%22Wilson%22%2C%22DtStart%22%3A%223%2F6%2F2017%22%2C%22DtStop%22%3A%223%2F6%2F2018%22%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/rapayor/ISO/16" "request" "p"  "WebService--InsertARAIDPayor"

# Read RAID Payors
echo "request=%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%7D" > request
dojsonPOST "http://localhost:8270/v1/rapayor/ISO/16" "request" "q"  "WebService--GetRAIDPayors"

# # Delete a RAID Payor that does not exist for the the specified RAID (it should go to RAID 20 but will go to RAID 16 instead)
echo "request=%7B%22cmd%22%3A%22delete%22%2C%22selected%22%3A%5B1049%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%2C%22TCID%22%3A367%2C%22DtStop%22%3A%223%2F11%2F2018%22%7D" > request
dojsonPOST "http://localhost:8270/v1/rapayor/ISO/16" "request" "r"  "WebService--DeleteARentablePayor-forceError"

# Delete a RAID Payor that does not exist for the the specified RAID
# echo "request%3D%7B%22cmd%22%3A%22delete%22%2C%22selected%22%3A%5B1%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%2C%22TCID%22%3A367%7D" > request
# dojsonPOST "http://localhost:8270/v1/rapayor/ISO/16" "request" "s"  "WebService--DeleteARentablePayor"

# # Read RAID Payors
# echo "request=%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%7D" > request
# dojsonPOST "http://localhost:8270/v1/rapayor/ISO/16" "request" "t"  "WebService--GetRAIDPayors"

# Read RAID Users
echo "request%3D%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%2C%22TCID%22%3A52%7D" > request
dojsonPOST "http://localhost:8270/v1/ruser/CCC/10" "request" "u"  "WebService--GetRAIDPayors"

# Test Transactant Typedown
dojsonGET "http://localhost:8270/v1/rentablestd/ISO?request%3D%7B%22search%22%3A%226%22%2C%22max%22%3A250%7D" "v" "WebService--GetRentableTypeDown"

# Search Payment Types
echo "request%3D%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%7D" > request
dojsonPOST "http://localhost:8270/v1/pmts/1" "request" "w"  "WebService--PaymentTypes-SearchAll"

# Get Specificy PaymentType - FORCE ERROR - no PMTID provided
echo "request%3D%7B%22cmd%22%3A%22get%22%2C%22recid%22%3A0%2C%22name%22%3A%22paymentTypeGrid%22%7D" > request
dojsonPOST "http://localhost:8270/v1/pmts/1" "request" "x"  "WebService--PaymentTypes-Get-ForceError"

# Get Specificy PaymentType
echo "request%3D%7B%22cmd%22%3A%22get%22%2C%22recid%22%3A0%2C%22name%22%3A%22paymentTypeGrid%22%7D" > request
dojsonPOST "http://localhost:8270/v1/pmts/1/1" "request" "y"  "WebService--PaymentTypes-Get"

# get Rentable types list for a business
dojsonGET "http://localhost:8270/v1/rtlist/2" "z" "WebService--GetRentableTypesForBusiness"

# get UI Values...
doPlainGET "http://localhost:8270/v1/uival/REX/app.AssessmentRules" "a1" "WebService--GetUIValue-app.AssessmentRules"
doPlainGET "http://localhost:8270/v1/uival/REX/app.ReceiptRules" "b1" "WebService--GetUIValue-app.ReceiptRules"

# rental Agreement typedown...
dojsonGET "http://localhost:8270/v1/rentalagrtd/CCC?request=%7B%22search%22%3A%22s%22%2C%22max%22%3A250%7D" "c1" "WebService--GetRentalAgreementTypeDown"


stopRentRollServer
echo "RENTROLL SERVER STOPPED"

echo "Restoring test database..."
mysql --no-defaults rentroll < restore.sql

logcheck
