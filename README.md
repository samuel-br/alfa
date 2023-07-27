This project consists of two services: the API service and the Manager service. It utilizes a relational database (PostgreSQL) to store transaction and advance payment data.

## Services

### API Service

The API service handles regular transactions and advance payments. It exposes endpoints to perform_advance and perform_transaction operations.

### Manager Service

The Manager service is responsible for updating data from the report file and performing debit transactions. It runs on a 12-hour interval to ensure periodic updates and debit transactions.

## Database

The relational database (PostgreSQL) stores two schemas: 'transaction' and 'advance'.

### Transaction Schema

The 'transaction' schema contains information about regular transactions:

- transaction_id (Primary Key): The transaction ID received from the call to perform_transaction.
- amount: The amount of money in the transaction.
- src_bank_account: Source bank account.
- dst_bank_account: Destination bank account.
- direction: Debit or credit.
- status: Transaction status (pending/success/fail).
- src_transaction: The transaction source (for advance payment transactions, it refers to the primary advance-pay transaction ID).

### Advance Schema

The 'advance' schema contains information about advance payment transactions:

- transaction_id (Primary Key): The transaction ID received from the call to perform_transaction for credit transactions during advance payment.
- dst_bank_account: Destination bank account for advance payment.
- amount: The amount of money in the advance payment.
- Debit_transaction: An array containing all the transaction IDs related to this advance payment.
- transaction: The number of successful debit transactions related to the advance payment.
- next_pay_date: The date for the next debit transaction.

## Classes (in the db package)

### TransactionClass

This class is responsible for CRUD operations on the 'transaction' table using ORM objects.

### AdvanceClass

This class is responsible for CRUD operations on the 'advance' table using ORM objects.

### BillingService Class

The BillingService class is injected with instances of the TransactionClass and AdvanceClass and contains the following methods:

- update: For updating data from the report file.
- advancePay: To perform an advance payment, including crediting the amount to the destination bank account and saving the transaction ID in the 'transaction' table and the 'advance' table. The next debit payment date is calculated and stored.
- debitPay: For managing debit payments, it queries all advance payments with the current date as the payment date (that are not completed yet) and performs debit transactions accordingly.
- getData: For getting report data from the Manager service.
- nextPayDateCal: For calculating the next payment date based on the Manager service's report data.
- performTransaction: For performing regular transactions.

## Flows

1. The system starts up, with the Manager service defining an interval of 12 hours to perform an update for the transaction data from the report file and then perform debit transactions.
2. The API service is up and listens for requests.
3. The API service triggers the BillingService to perform an advance payment (perform_advance), which involves crediting the amount to the destination bank account and saving the transaction ID in the 'transaction' table and the 'advance' table. The next debit payment date is calculated and stored.
4. Every 12 hours, the update is triggered in the Manager service, which downloads the report file, reads it, and updates the 'transaction' table. If a transaction has a source transaction of an advance payment transaction, the process updates the 'advance' table with new debit transaction data (adding the transaction ID to the table). If the status is successful, the transaction made for this advance payment is also updated.
5. After updating the data, the debitPay function is triggered to query all advance payments with the current date as the payment date (that are not completed yet) and perform debit transactions accordingly
