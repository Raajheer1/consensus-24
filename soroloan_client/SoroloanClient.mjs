import {
  Keypair,
  Contract,
  SorobanRpc,
  TransactionBuilder,
  Networks,
  BASE_FEE,
  xdr,
  Horizon,
  Address,
  XdrLargeInt,
} from "@stellar/stellar-sdk";
import { waitForTransactionConfirmation } from "./waitForTxConf.mjs";

const server = new SorobanRpc.Server(
  "https://soroban-testnet.stellar.org:443"
);
const horizonServer = new Horizon.Server(
  "https://horizon-testnet.stellar.org"
);
const principalAssetContractAddress =
  "CBIELTK6YBZJU5UP2WWQEUCYKLPU6AUNZ2BQ4WWFEIE3USCIHMXQDAMA";
const soroloanContractAddress =
  "CDEHQZUOGIPALIFYWTL63CMJFZ5PWID5T5CJNRTXGZZ72H63DXGVFM2T";
const soroloanContract = new Contract(soroloanContractAddress);

// Sends stellar transaction that calls the loan function of the soroloan contract
// The principal asset (USDC) is exchanged for a bond asset
export async function submitLoan(
  lenderSecretKey,
  loanAmmount,
  bondAssetContractAddress
) {
  const lenderKeypair = Keypair.fromSecret(lenderSecretKey);
  console.log(
    "Submitting loan of ammount ",
    loanAmmount,
    "for lender ",
    lenderKeypair.publicKey()
  );
  const lenderAccount = await server.getAccount(
    lenderKeypair.publicKey()
  );
  let transferTransaction = new TransactionBuilder(
    lenderAccount,
    {
      fee: BASE_FEE,
      networkPassphrase: Networks.TESTNET,
    }
  )
    .addOperation(
      // pub fn loan(e: Env,
      //     lender: Address,
      //     principal_token_code: Address,
      //     loan_token_code: Address,
      //     lone_amount: i128) -> Result<u32, Error>
      soroloanContract.call(
        "loan",
        // lender: Address
        xdr.ScVal.scvAddress(
          new Address(lenderKeypair.publicKey()).toScAddress()
        ),
        // principal_token_code: Address
        xdr.ScVal.scvAddress(
          new Address(
            principalAssetContractAddress
          ).toScAddress()
        ),
        // loan_token_code: Address
        xdr.ScVal.scvAddress(
          new Address(bondAssetContractAddress).toScAddress()
        ),
        // loan_amount: i128
        new XdrLargeInt("i128", loanAmmount).toI128()
      )
    )
    .setTimeout(100)
    .build();
  let preparedTransaction;
  try {
    preparedTransaction = await server.prepareTransaction(
      transferTransaction
    );
  } catch (error) {
    console.error("Error preparing transaction:", error);
    return;
  }
  preparedTransaction.sign(lenderKeypair);
  const transactionResult = await server.sendTransaction(
    preparedTransaction
  );
  const transactionResponse = await server.sendTransaction(
    preparedTransaction
  );
  await waitForTransactionConfirmation(
    server,
    transactionResponse
  );
  console.log("Loan request successful");
  return true;
}

export async function getBalances(publicKey) {
  const account = await horizonServer.loadAccount(publicKey);
  return account.balances;
}

export async function issueBondToken(
  issuerSecretKey,
  bondTokenCode,
  bondTokenAmount
) {
  const issuerKeypair = Keypair.fromSecret(issuerSecretKey);
  const issuerAccount = await horizonServer.loadAccount(
    issuerKeypair.publicKey()
  );
  const bondAsset = new Asset(
    bondTokenCode,
    issuerKeypair.publicKey()
  );
  const trustlineTransaction = new TransactionBuilder(
    issuerAccount,
    {
      fee: BASE_FEE,
      networkPassphrase: Networks.TESTNET,
    }
  )
    .addOperation(
      Operation.changeTrust({
        asset: bondAsset,
        limit: bondTokenAmount,
      })
    )
    .setTimeout(100)
    .build();
  trustlineTransaction.sign(issuerKeypair);
  const trustlineResponse = await server.submitTransaction(
    trustlineTransaction
  );
  await waitForTransactionConfirmation(
    server,
    trustlineResponse.hash
  );
  console.log("Trustline transaction successful");
}

//WIP: See tests/submit-votes.js
export async function submitVotes() {}
