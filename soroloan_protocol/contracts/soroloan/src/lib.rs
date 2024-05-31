#![no_std]
/*
## Soroloan
This is the implementation of the Soroloan smart lending protocol. It includes routines for issuing debt/bond tokens
and a governance system for distributing the loan to the borrower once the lenders have reached consensus.

# Loan (loan) issues
1. Any principal could be exchanged for any loan token. There needs to be something stateful that
enforces supported conversions. I like that its generalized and not walled into USDC, but further
development is needed to secure the exchange.
2. No repayment mechanism

# Governance (tabulate) issues:
1. The contract only knows the loan ammount to be the total stake of the
voters. If lenders do not vote, their stake won't be included in the loan.
We have identified a number of ways around this and will implement them in
the coming weeks
2. Social votes aren't yet supported in the contract.
3. I'm certain there is a better way to pass the votes to the contract
than in a vector of vectors. Or at least there is certainly a better way
to marshal them into objects and validate them.
4. There is nothing stopping anyone from calling this contract and getting
our USDC reserves (LOL). This is just a proof of concept to lay the groundwork
for moving money around.
5. Checking that the vote is more than 75% using i128 is very expensive.
6. There is no on chain record of the vote passing other than the funds being
transferred to the borrower
*/
use soroban_sdk::{contract, contractimpl, contracterror, symbol_short, vec, Env, Symbol, Vec, Address, token, String};

#[contracterror]
#[derive(Copy, Clone, Debug, Eq, PartialEq, PartialOrd, Ord)]
#[repr(u32)]
pub enum Error {
    NotEnoughUsdc = 1,
    NotEnoughLoanToken = 2,
    VoteMissingValidAddress = 3,
    VoteMissingVote = 4,
    VoteMissingMessage = 5,
    VoteMissingSignature = 6,
    VoterBalanceInsufficient = 7,
}

pub struct Vote {
    pub voter: Address,
    pub vote: String,
    pub message: String,
    pub signature: String,
}

#[contract]
pub struct LoanContract;

#[contractimpl]
impl LoanContract {
    pub fn loan(e: Env,
                lender: Address,
                principal_token_code: Address,
                bond_token_code: Address,
                lone_amount: i128) -> Result<u32, Error> {
        lender.require_auth();
        let contract_wallet = e.current_contract_address();
        let bond_token_client = token::Client::new(&e, &bond_token_code);
        let principal_token_client = token::Client::new(&e, &principal_token_code);

        let bond_token_balance = bond_token_client.balance(&contract_wallet);
        let principal_token_balance = principal_token_client.balance(&lender);

        if principal_token_balance < lone_amount {
            return Err(Error::NotEnoughUsdc);
        }

        if bond_token_balance < lone_amount {
            return Err(Error::NotEnoughLoanToken);
        }

        principal_token_client.transfer(&lender, &contract_wallet, &lone_amount);
        bond_token_client.transfer(&contract_wallet, &lender, &lone_amount);

        return Ok(1);
    }
    pub fn tabulate(e: Env,
                    borrower: Address,
                    bond_token: Address,
                    principal_token: Address,
                    votes: Vec<Vec<String>>) -> Result<Symbol, Error> {
        let mut no_votes = 0;
        let mut yes_votes = 0;
        for vote_vec in votes {
            //Validate vote
            let voter_address_string = match vote_vec.get(0) {
                Some(v) => v,
                None => String::from_str(&e, ""),
            };
            if voter_address_string == String::from_str(&e, "") {
                return Err(Error::VoteMissingValidAddress);
            }
            let voter_vote_string = match vote_vec.get(1) {
                Some(v) => v,
                None => String::from_str(&e, ""),
            };
            if voter_vote_string == String::from_str(&e, "") {
                return Err(Error::VoteMissingVote);
            }
            let voter_message_string = match vote_vec.get(2) {
                Some(v) => v,
                None => String::from_str(&e, ""),
            };
            if voter_message_string == String::from_str(&e, "") {
                return Err(Error::VoteMissingMessage);
            }
            let voter_signature_string = match vote_vec.get(3) {
                Some(v) => v,
                None => String::from_str(&e, ""),
            };

            let vote = Vote {
                voter: Address::from_string(&voter_address_string),
                vote: voter_vote_string,
                message: voter_message_string,
                signature: voter_signature_string,
            };

            /*TODO: Validate signiture using Crypto::ed25519_verify()
            //Message and signature will need to come through as bytes
            Crypo::ed25519_verify(&vote.message, &vote.signature, &vote.voter)
            */

            //Check account balances
            let bond_token_client = token::Client::new(&e, &bond_token);
            let bond_token_balance = bond_token_client.balance(&vote.voter);

            if vote.vote == String::from_str(&e, "yes") {
                yes_votes += bond_token_balance;
            } else if vote.vote == String::from_str(&e, "no") {
                no_votes += bond_token_balance;
            }
        }
        //calculate vote result
        let vote_result = if Self::is_more_than_75_percent(yes_votes, no_votes) {
            let principal_token_client = token::Client::new(&e, &principal_token);
            let contract_wallet = e.current_contract_address();
            let loan_ammount = yes_votes + no_votes;
            principal_token_client.transfer(&contract_wallet, &borrower, &loan_ammount);

            symbol_short!("yes")
        } else {
            symbol_short!("no")
        };

        return Ok(vote_result);
    }
    fn is_more_than_75_percent(a: i128, b: i128) -> bool {
        let a_float = a as f32;
        let b_float = b as f32;
        a_float > (a_float + b_float) * 0.75
    }

    pub fn hello(e: Env, to: Symbol) -> Vec<Symbol> {
        vec![&e, symbol_short!("Hello"), to]
    }
}

mod test;
