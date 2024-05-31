import { getBalances } from "../LoanContractClient.mjs";

(async () => {
  try {
    const balances = await getBalances(
      "GCWY3M4VRW4NXJRI7IVAU3CC7XOPN6PRBG6I5M7TAOQNKZXLT3KAH362"
    );
    console.log(balances);
  } catch (err) {
    console.log(err);
  }
})();
