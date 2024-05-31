export async function waitForTransactionConfirmation(
  server,
  transactionResponse
) {
  console.log("Transaction status", transactionResponse.status);
  if (transactionResponse.status === "PENDING") {
    let getResponse = await server.getTransaction(
      transactionResponse.hash
    );
    // Poll `getTransaction` until the status is not "NOT_FOUND"
    while (getResponse.status === "NOT_FOUND") {
      console.log("Waiting for transaction confirmation...");
      // See if the transaction is complete
      getResponse = await server.getTransaction(
        transactionResponse.hash
      );
      // Wait one second
      await new Promise((resolve) => setTimeout(resolve, 1000));
    }

    console.log(
      `getTransaction response: ${JSON.stringify(getResponse)}`
    );

    if (getResponse.status === "SUCCESS") {
      // Make sure the transaction's resultMetaXDR is not empty
      if (!getResponse.resultMetaXdr) {
        throw "Empty resultMetaXDR in getTransaction response";
      }
      // Find the return value from the contract and return it
      let transactionMeta = getResponse.resultMetaXdr;
      let returnValue = transactionMeta
        .v3()
        .sorobanMeta()
        .returnValue();
      console.log(`Transaction result: ${returnValue.value()}`);
    } else {
      throw `Transaction failed: ${getResponse.resultXdr}`;
    }
  } else {
    throw transactionResponse.errorResult;
  }
}
