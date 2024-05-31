import { submitLoan } from "../LoanContractClient.mjs";

try {
  const res = submitLoan(
    "SCQN3XGRO65BHNSWLSHYIR4B65AHLDUQ7YLHGIWQ4677AZFRS77TCZRB",
    100,
    "CCQVSAP342JTUC6XRIP2T7VVRYO6JHQMZMUTX2W6NTE36OMOA4QKH5QS"
  );
} catch (err) {
  console.log(err);
}
