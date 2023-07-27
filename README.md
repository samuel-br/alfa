# alfa

services

we have two services running independently 
api service - the api for the call perform_advance
manager service - responsible for the update of the data from the report file and perform the debit transaction


Database

a relational database (postgres) - 

Schemas

transaction schema
transaction_id -  the primary key is the transaction id we receive from the call to perform_transaction
amount - the amount money in transaction
src_bank_account - source bank account
dst_bank account - destination bank account
direction - debit/credit
status - pending/success/fail
src_transaction - the transaction source: the same as transaction id in regular transaction, if transaction is part of the debits transaction of advance payment the source transaction will be the transaction id of the primary advance-pay transaction

advance schema

transaction_id -   the primary key is the transaction id we receive from the call to perform_transaction for the credit transaction when preform advance
dst_bank_account -  destination bank account
amount - the amount money in transaction
Debit_transaction - array contain all the transactions id related to this advance-pay
transaction - the number of success debit transaction 
next_pay_date - a date for the next debit transaction 

relational database is a good match here because the data structure and the support of SQL for some of the query we need to perform


classes
We have two classes in the db package, each one is for a different table in the database.
The class performs crud operations on the database using ORM objects.

Another main class is the BillingService class 
this class injected with instances of the classes of the db package 
and have the methods
update - for updating the data from the report file
advancePay - to perform advance payment
debitPay - for managing the debit payments 
getData - for getting the data of the report
nextPayDateCal - for calculate the next pay date
performTransaction -for perform regular transaction


Flows
system is up - the main function for the manager service defines an interval to perform an update for the transaction from the report file every 12 hours and only after to perform the debit transaction after the data is updated.
In addition api service is up and listen for requests

api triggered
api service using Billing-service performs Advance-pay, first make a request for credit the amount to dst bank account then save the transaction_id in the transaction table
and the advance-pay in the advance table and calculate the next debit pay

once every 12 hours(I used a time interval on the service itself for convenience, in production can use cron-job or similar) update is triggered \
the function downloads the file, reads it and updates  the transaction table.
If  a transaction have source transaction of advance pay transaction
the process update the advance table too with the new debit transaction data(add the transaction id for the table) and if the status if success update the transaction made for this advance-pay

after update the data debitPay function is triggered 
query all the advance pay with the pay date of the current date that not completed yet
and perform debit transaction on behalf 
Billing service

ARCHITECRUTE
we will have two services running independly 
