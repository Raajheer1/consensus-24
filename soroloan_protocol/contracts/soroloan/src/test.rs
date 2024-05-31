#![cfg(test)]
extern crate std;

use crate::{token, LoanContractClient, LoanContract};
use soroban_sdk::{testutils::{Address as _}, Address, Env, Vec, String, vec, symbol_short};

#[test]
fn test() {
    let e = Env::default();
    e.mock_all_auths();

    let contract_id = e.register_contract(None, LoanContract);
    let loan_contract = LoanContractClient::new(&e, &contract_id);

    let token_admin = Address::generate(&e);

    let (principal_token_client, principal_token_admin_client) = create_token_contract(&e, &token_admin);
    let (bond_token_client, bond_token_admin_client) = create_token_contract(&e, &token_admin);
    let lender = Address::generate(&e);

    principal_token_admin_client.mint(&lender, &1000);
    bond_token_admin_client.mint(&loan_contract.address, &1000);

    loan_contract.loan(
        &lender,
        &principal_token_client.address,
        &bond_token_client.address,
        &1000,
    );

    //Check that lender has lone tokens
    assert_eq!(bond_token_client.balance(&lender), 1000);
    //Check that lender has no principal tokens
    assert_eq!(principal_token_client.balance(&lender), 0);
}

#[test]
fn test_tabulate() {
    let e = Env::default();
    e.mock_all_auths();

    let contract_id = e.register_contract(None, LoanContract);
    let governance_contract = LoanContractClient::new(&e, &contract_id);
    let borrower = Address::generate(&e);
    let token_admin = Address::generate(&e);
    let (principal_token_client, principal_token_admin_client) = create_token_contract(&e, &token_admin);
    principal_token_admin_client.mint(&governance_contract.address, &1000000);
    let (bond_token_client, bond_token_admin_client) = create_token_contract(&e, &token_admin);

    let mut votes_list = Vec::new(&e);
    let mut yes_votes = 0;
    let mut no_votes = 0;
    for i in 0..100 {
        let fund_ammount = &(i * 100);
        let voter = Address::generate(&e);
        bond_token_admin_client.mint(&voter, &fund_ammount);
        let vote = if generate_random(&e) {
            yes_votes += fund_ammount;
            String::from_str(&e, "yes")
        } else {
            no_votes += fund_ammount;
            String::from_str(&e, "no")
        };
        let message = String::from_str(&e, "msg");
        let signature = String::from_str(&e, "sig");
        votes_list.push_front(vec![&e, voter.to_string(), vote, message, signature]);
    }

    let expected_result = if is_more_than_75_percent(yes_votes as i128, no_votes as i128) {
        symbol_short!("yes")
    } else {
        symbol_short!("no")
    };

    let result = governance_contract.tabulate(
        &borrower,
        &bond_token_client.address,
        &principal_token_client.address,
        &votes_list,
    );

    assert_eq!(result, expected_result)
}

fn create_token_contract<'a>(
    e: &Env,
    admin: &Address,
) -> (token::Client<'a>, token::StellarAssetClient<'a>) {
    let addr = e.register_stellar_asset_contract(admin.clone());
    (
        token::Client::new(e, &addr),
        token::StellarAssetClient::new(e, &addr),
    )
}

fn generate_random(e: &Env) -> bool {
    let addr1 = Address::generate(e);
    let addr2 = Address::generate(e);
    addr1 < addr2
}

fn is_more_than_75_percent(a: i128, b: i128) -> bool {
    let a_float = a as f32;
    let b_float = b as f32;
    a_float > (a_float + b_float) * 0.75
}