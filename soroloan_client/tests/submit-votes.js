/*
This exercises the tabulate function of the Soroloan smart contract.
Generates 100 random votes and submits for a tally.
*/
const {
  Keypair,
  Contract,
  SorobanRpc,
  TransactionBuilder,
  Networks,
  BASE_FEE,
  xdr,
  XdrLargeInt,
  Address,
} = require("@stellar/stellar-sdk");
const waitForTransactionConfirmation = require("./waitForTxConf");

(async () => {
  const server = new SorobanRpc.Server(
    "https://soroban-testnet.stellar.org:443"
  );
  //Bond asset contract
  const bondAssetContractAddress =
    "CCQVSAP342JTUC6XRIP2T7VVRYO6JHQMZMUTX2W6NTE36OMOA4QKH5QS";
  //Principal asset contract
  const principalAssetContractAddress =
    "CBIELTK6YBZJU5UP2WWQEUCYKLPU6AUNZ2BQ4WWFEIE3USCIHMXQDAMA";
  //Soroloan smart contract
  const soroloanContractAddress =
    "CDEHQZUOGIPALIFYWTL63CMJFZ5PWID5T5CJNRTXGZZ72H63DXGVFM2T";
  const borrowerKeypair = Keypair.fromSecret(
    "SCQN3XGRO65BHNSWLSHYIR4B65AHLDUQ7YLHGIWQ4677AZFRS77TCZRB"
  );
  const sourceKeypair = Keypair.fromSecret(
    "SCQN3XGRO65BHNSWLSHYIR4B65AHLDUQ7YLHGIWQ4677AZFRS77TCZRB"
  );
  const sourceAccount = await server.getAccount(
    sourceKeypair.publicKey()
  );
  const soroloanContract = new Contract(loanContractAddress);
  //generate 100 random votes
  const votes = [];
  for (let i = 0; i < 100; i++) {
    const voterKeypair = Keypair.random().publicKey();
    const vote = Math.random() > 0.5 ? "YES" : "NO";
    const message = `Vote number ${i + 1}: ${vote}`;
    const signature = voterKeypair
      .sign(Buffer.from(message))
      .toString("base64");

    votes.push({
      voter: new Address(
        voterKeypair.publicKey()
      ).toScAddress(),
      vote: vote,
      message: message,
      signature: signature,
    });
  }

  console.log(
    "Borrower public key:",
    borrowerKeypair.publicKey()
  );
  const lenderAccount = await server.getAccount(
    lenderKeypair.publicKey()
  );

  let transferTransaction = new TransactionBuilder(
    sourceAccount,
    {
      fee: BASE_FEE,
      networkPassphrase: Networks.TESTNET,
    }
  )
    .addOperation(
      // pub fn tabulate(e: Env,
      //     borrower: Address,
      //     bond_token: Address,
      //     principal_token: Address,
      //     votes: Vec<Vec<String>>)
      soroloanContract.call(
        "tabulate",
        //     borrower: Address,
        xdr.ScVal.scvAddress(
          new Address(borrowerKeypair.publicKey()).toScAddress()
        ),
        //     bond_token: Address,
        xdr.ScVal.scvAddress(
          new Address(bondAssetContractAddress).toScAddress()
        ),
        //     principal_token: Address,
        xdr.ScVal.scvAddress(
          new Address(bondAssetContractAddress).toScAddress()
        ),
        //     votes: Vec<Vec<String>>)
        new XdrLargeInt("i128", 100).toI128()
      )
    )
    .setTimeout(30)
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

  console.log(
    `Signed prepared transaction XDR: ${preparedTransaction
      .toEnvelope()
      .toXDR("base64")}`
  );

  let sendResponse = await server.sendTransaction(
    preparedTransaction
  );
  await waitForTransactionConfirmation(server, sendResponse);
})();
