CREATE TABLE IF NOT EXISTS loans (
    id VARCHAR(255) PRIMARY KEY,
    borrower_id VARCHAR(255) NOT NULL,
    principal_amount FLOAT NOT NULL,
    rate FLOAT NOT NULL,
    roi FLOAT NOT NULL,
    agreement_link VARCHAR(255) NOT NULL,
    state VARCHAR(50) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS approvals (
    loan_id VARCHAR(255) PRIMARY KEY,
    validator_id VARCHAR(255) NOT NULL,
    approval_date TIMESTAMP NOT NULL,
    proof_picture VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS investments (
    loan_id VARCHAR(255) NOT NULL,
    investor_id VARCHAR(255) NOT NULL,
    amount FLOAT NOT NULL,
    invested_at TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS disbursements (
    loan_id VARCHAR(255) PRIMARY KEY,
    field_officer_id VARCHAR(255) NOT NULL,
    disbursement_date TIMESTAMP NOT NULL,
    agreement_letter VARCHAR(255) NOT NULL
);