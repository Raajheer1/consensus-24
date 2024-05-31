import { RouteLocationRaw } from "vue-router";

export interface Link {
  title: string;
  icon?: string;
  to?: RouteLocationRaw;
  href?: string;
  subLinks?: Link[];
  showSubLinks?: boolean;
  roles?: string[];
  separator?: boolean;
  separatorTitle?: string;
}

export interface Account {
  id: number;
  first_name: string;
  last_name: string;
  email: string;
  company_name?: string;
  profile_pic?: string;
  secret_key: string;
  public_key: string;
}

export interface Loan {
  id: number;
  borrower_id: number;
  borrower: Account;
  goal_amount: number; // USDC
  amount_raised: number; // USDC
  number_of_payments: number;
  payment_schedule: string; // weekly, monthly
  interest_rate: number; // 0.05
  title: string;
  description: string;
  image_url: string;
  lenders?: LoadLender[];
  governance_votes: GovernanceVote[];
  loan_token_asset_code?: string; // 56 char alphanumeric
  created_at: string;
  funded_at?: string;
}

export interface LoadLender {
  id: number;
  loan: Loan;
  loan_id: number;
  lender_id: number;
  loan_amount: number;
  created_at: number;
}

export interface GovernanceVote {
  id: number;
  loan_id: number;
  amount_raised: number;
  voting_active: boolean;
  vote_passed: boolean;
  uuid: string;
  title: string;
  description: string;
  votes?: Vote[];
  created_at: string;
  end_at: string;
}

export interface Vote {
  id: number;
  governance_vote_id: number;
  lender_id: number;
  approves: boolean; // pass or fail
}

export interface Feed {
  id: number;
  title: string;
  description: string;
  created_by: uint;
  created_by_user: Account;
  created_at: string;
}
