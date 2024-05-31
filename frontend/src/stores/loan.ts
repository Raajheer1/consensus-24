import { defineStore } from "pinia";
import { API } from "@/utils/api";
import { LoadLender, Loan } from "@/types";

interface LoanState {
  loans: Loan[];
  my_loan_lendings: LoadLender[];
  my_loans: Loan[];
  error: string | null;
  fetching: boolean;
  hasFetched: boolean;
}

const useLoanStore = defineStore({
  id: "loan",
  state: () =>
    ({
      loans: [],
      my_loan_lendings: [],
      my_loans: [],
      error: null,
      fetching: false,
      hasFetched: false,
    }) as LoanState,
  getters: {
    avgReturn: (state) => {
      if (state.my_loans.length === 0) return 0;
      const sum = state.my_loans.reduce((acc, loan) => acc + loan.interest_rate, 0);
      return sum / state.my_loans.length;
    },
  },
  actions: {
    async fetchLoans(): Promise<void> {
      this.fetching = true;
      try {
        const { data } = await API.get("/v1/loan");
        this.loans = data;
      } catch (e) {
        console.log(e);
        this.loans = [];
      } finally {
        this.fetching = false;
        this.hasFetched = true;
      }
    },
    async fetchMyLoans(account_id: string): Promise<void> {
      if (!account_id || account_id === "") {
        return;
      }
      this.fetching = true;
      try {
        const { data } = await API.get(`v1/account/${account_id}/loan_lendings`);
        this.my_loan_lendings = data;

        // Use Sets for Efficient Lookups
        const existingLoanIds = new Set(this.my_loans.map((loan) => loan.id));

        // Filter & Extend
        this.my_loans.push(
          ...this.my_loan_lendings
            .filter((loanLending) => !existingLoanIds.has(loanLending.loan.id))
            .map((loanLending) => loanLending.loan)
        );
      } catch (e) {
        this.loans = [];
      } finally {
        this.fetching = false;
      }
    },
    fetchLoan(id: number): Loan | null {
      return this.loans.find((loan) => loan.id === id) || null;
    },
  },
});

export default useLoanStore;
