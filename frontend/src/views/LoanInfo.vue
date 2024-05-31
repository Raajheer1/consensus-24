<template>
  <Page v-if="loan === null" title="Err not found">
    <div class="flex items-center justify-center h-screen bg-main-blue">
      <h1 class="text-white text-4xl font-bold">404</h1>
    </div>
  </Page>
  <Page
    v-else
    :title="loan.title"
    :subtitle="`Posted on: ${new Date(loan.created_at).toLocaleDateString()} at
    ${new Date(loan.created_at).toLocaleTimeString()}`"
    button
    :button-target="`/${route.query.return || ''}`"
  >
    <Modal v-if="showPaymentModal" title="Payment" width="sm" @close="showPaymentModal = false">
      <div class="text-center">
        <img v-if="loan.image_url" class="mx-auto rounded-xl" :src="loan?.image_url" alt="Loan Image" />
        <h2 class="flex font-bold text-main-gold text-6xl mt-10">
          <span class="my-auto">$</span>
          <!--            TODO - Max wallet or amount remaining  -->
          <input
            v-model="paymentAmount"
            class="min-w-24 max-w-full text-center border-b border-main-blue outline-0 focus:border-transparent focus:ring-0 bg-transparent"
            type="number"
            min="1"
            :max="loan.goal_amount - loan.amount_raised"
          />
        </h2>
        <h5 class="text-main-grey font-semibold">$175 Balance</h5>
      </div>

      <div class="text-center mt-10 space-x-10">
        <button
          class="px-8 py-2 text-sm text-gray-600 focus:outline-none hover:underline"
          @click="showPaymentModal = false"
        >
          Cancel
        </button>
        <button
          class="px-8 py-2 text-sm rounded text-white bg-green-500 focus:outline-none hover:bg-green-400"
          @click="createLoanLender()"
        >
          Send Payment
        </button>
      </div>
    </Modal>

    <Modal v-if="showCreateVoteModal" title="Start a new vote." width="sm" @close="closeCreateVote()">
      <p class="font-bold text-gray-600 text-sm">Title</p>
      <input
        v-model="title"
        class="mb-2 py-1 border-b hover:border-b-main-blue outline-0 focus:border-transparent focus:ring-0 bg-transparent"
        placeholder="New City"
      />

      <p class="mt-2 font-bold text-gray-600 text-sm">Description</p>
      <textarea
        v-model="description"
        class="w-full mb-2 py-1 border-b hover:border-b-main-blue outline-0 focus:border-transparent focus:ring-0 bg-transparent"
        placeholder="With our ever growing business its time that we have a new city to expand to."
      />

      <div class="text-center mt-4 space-x-10">
        <button class="px-8 py-2 text-sm text-gray-600 focus:outline-none hover:underline" @click="closeCreateVote()">
          Cancel
        </button>
        <button
          class="px-8 py-2 text-sm rounded text-white bg-green-500 focus:outline-none hover:bg-red-400"
          @click="createGovVote()"
        >
          Create
        </button>
      </div>
    </Modal>

    <Modal v-if="showVotingModal" title="Governance Vote" width="sm" @close="toggleVoteModal(null)">
      <p class="font-bold text-gray-600 text-sm">Title</p>
      <p v-if="governanceVote.title !== ''" class="pb-2">{{ governanceVote.title || "Err Rendering" }}</p>
      <p v-else class="pb-2">Release Funding</p>
      <p class="font-bold text-gray-600 text-sm mt-1">Description</p>
      <p v-if="governanceVote.title" class="pb-2">{{ governanceVote.description || "No Description Provided." }}</p>
      <p v-else class="pb-2">
        By voting 'Yes' you are agreeing to release funds to
        <span class="underline decoration-main-blue font-semibold">
          {{ borrower!.company_name || borrower!.first_name }}
        </span>
        for the loan titled
        <span class="underline decoration-main-blue font-semibold">{{ loan!.title }}</span
        >.
      </p>
      <p class="text-xs italic mt-4">
        Note: Funds will be released only if a supermajority (75%) of voting tokens approve by the end of the Governance
        Vote or when the last voter has cast their vote.
      </p>
      <div class="text-center mt-4 space-x-10">
        <button
          class="px-8 py-2 text-sm rounded text-white bg-red-500 focus:outline-none hover:bg-red-400"
          @click="castVote(false)"
        >
          No
        </button>
        <button
          class="px-8 py-2 text-sm rounded text-white bg-green-400 focus:outline-none hover:bg-red-400"
          @click="castVote(true)"
        >
          Yes
        </button>
      </div>
    </Modal>

    <Modal v-if="showVoteInfoModal" title="Governance Vote" width="sm" @close="toggleVoteInfoModal(null)">
      <p class="font-bold text-gray-600 text-sm">Title</p>
      <p v-if="voteInfo!.title !== ''" class="pb-2">{{ voteInfo!.title || "Err Rendering" }}</p>
      <p v-else class="pb-2">Release Funding</p>
      <p class="font-bold text-gray-600 text-sm mt-1">Description</p>
      <p v-if="voteInfo!.title" class="w-full pb-2">{{ voteInfo!.description || "No Description Provided." }}</p>
      <div class="grid grid-cols-2 items-center gap-y-2 mt-2">
        <div>
          <p class="font-bold text-gray-600 text-sm">Voting</p>
          <p>{{ voteInfo!.voting_active ? "In Progress" : "Completed" }}</p>
        </div>
        <div>
          <p class="font-bold text-gray-600 text-sm">Your Vote</p>
          <p>
            {{ voteInfo!.votes.filter((vote) => vote.lender_id === userStore.user.id)[0].approves ? "For" : "Against" }}
          </p>
        </div>

        <div>
          <p class="font-bold text-gray-600 text-sm">Votes(Received/Total)</p>
          <p>{{ voteInfo!.votes.length }}/{{ loanStore.fetchLoan(voteInfo!.loan_id)!.lenders!.length }}</p>
        </div>
        <div>
          <p class="font-bold text-gray-600 text-sm">Time Left</p>
          <p>
            {{ Math.ceil((new Date(voteInfo!.end_at).getTime() - new Date().getTime()) / (1000 * 60 * 60 * 24)) }} Days
          </p>
        </div>
      </div>
      <div class="mt-4">
        <Feed :id="voteInfo.id" feed-type="governance" />
      </div>
      <div class="text-center mt-4 space-x-10">
        <button
          class="px-8 py-2 text-sm text-gray-600 focus:outline-none hover:underline"
          @click="toggleVoteInfoModal()"
        >
          Cancel
        </button>
      </div>
    </Modal>

    <Modal v-if="showUserModal" title="Borrower Information" width="sm" @close="toggleUserModal(null)">
      <Profile v-if="userModalUser != null" :user="userModalUser" :show-secret="false" />
      <div v-else>
        <p class="text-center">No user found.</p>
      </div>
      <Feed v-if="userModalUser != null" :id="userModalUser!.id" feed-type="account" />
      <div class="text-center mt-4 space-x-10">
        <button
          class="px-8 py-2 text-sm text-gray-600 focus:outline-none hover:underline"
          @click="toggleUserModal(null)"
        >
          Close
        </button>
      </div>
    </Modal>

    <img v-if="!!loan.image_url" class="w-full h-48 object-cover rounded-lg" :src="loan.image_url" alt="loan" />

    <div class="relative w-full bg-gray-300 rounded-full h-4">
      <div
        class="bg-main-gold h-4 rounded-full"
        :style="{ width: `${(loan.amount_raised / loan.goal_amount) * 100}%` }"
      ></div>
      <span class="absolute top-1/2 -translate-y-1/2 left-1 text-xs text-white">$0</span>
      <span class="absolute top-1/2 -translate-y-1/2 right-1 text-xs text-white">${{ loan.goal_amount }}</span>
    </div>

    <div>
      <div>
        <p class="font-bold text-gray-600 text-sm">Description</p>
        <p class="pb-2">{{ loan.description }}</p>
      </div>

      <div class="grid grid-cols-2 items-center">
        <div>
          <p class="font-bold text-gray-600 text-sm">Interest Rate (APY)</p>
          <p>{{ loan.interest_rate }}%</p>
        </div>
        <div>
          <p class="font-bold text-gray-600 text-sm">Term</p>
          <p>
            {{ loan.number_of_payments
            }}{{ loan.payment_schedule == "monthly" ? "mo" : loan.payment_schedule == "weekly" ? "wk" : "yr" }}
          </p>
        </div>
      </div>
      <!--      amLender: {{ loanLending }}-->
      <div v-if="loanLending != null" class="grid grid-cols-1 gap-y-2 mb-2">
        <div class="grid grid-cols-2 items-center">
          <div>
            <p class="font-bold text-gray-600 text-sm">Lent Amount</p>
            <p>${{ loanLending.loan_amount }}</p>
          </div>
          <div>
            <p class="font-bold text-gray-600 text-sm">Lent Date</p>
            <p>{{ new Date(loanLending.created_at).toLocaleDateString() }}</p>
          </div>
        </div>
        <div
          v-if="loan.funded_at != '0001-01-01T00:00:00Z' && loan.funded_at != null"
          class="grid grid-cols-2 items-center"
        >
          <div>
            <p class="font-bold text-gray-600 text-sm">Fund Date</p>
            <p>{{ new Date(loan.funded_at).toLocaleDateString() }}</p>
          </div>
          <div>
            <p class="font-bold text-gray-600 text-sm">Completion Date</p>
            <p>
              {{ new Date().setDate(new Date(loan.funded_at).getDate() + payoutDaysRemaining) }}
              ({{ payoutDaysRemaining }}d)
            </p>
          </div>
        </div>
        <div v-if="loan!.governance_votes.length != 0">
          <h5 class="text-lg font-bold">Governance Votes</h5>
          <table class="w-full text-center table-auto">
            <thead>
              <tr class="font-bold text-gray-600 text-sm">
                <th>Date</th>
                <th>Type</th>
                <th>Vote</th>
                <th>Status</th>
                <th>Action</th>
              </tr>
            </thead>
            <tbody class="text-gray-600 text-sm">
              <tr v-for="govVote in loan!.governance_votes" :key="govVote.id">
                <td>
                  {{ new Date(govVote.created_at).getMonth() + 1 }}/{{ new Date(govVote.created_at).getDate() + 1 }}
                </td>
                <td>{{ govVote.title === "" ? "Funding" : "Social" }}</td>
                <td>
                  {{
                    govVote.voting_active
                      ? hasVoted(govVote)
                        ? "Sent"
                        : "Pending"
                      : govVote.vote_passed
                        ? "Passed"
                        : "Failed"
                  }}
                </td>
                <td>{{ govVote.voting_active ? "Active" : "Inactive" }}</td>
                <td>
                  <button
                    v-if="govVote.voting_active && !hasVoted(govVote)"
                    class="my-auto w-12 h-8 text-white bg-main-blue rounded-2xl"
                    @click="toggleVoteModal(govVote)"
                  >
                    <i class="fa-solid fa-check-to-slot"></i>
                  </button>
                  <button
                    v-else
                    class="my-auto w-12 h-8 text-white bg-main-grey rounded-2xl"
                    @click="toggleVoteInfoModal(govVote)"
                  >
                    <i class="fa-solid fa-info"></i>
                  </button>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>

    <Feed :id="loan!.id" feed-type="loan" />

    <div>
      <h2 class="text-lg font-bold">Borrower Information</h2>
      <div class="flex w-full">
        <img
          class="w-14 h-14 rounded-full border-2 border-main-gold"
          :src="borrower?.profile_pic"
          alt="Profile Picture"
        />
        <div class="mx-5 my-auto">
          <h6 v-if="!borrower?.company_name" class="text-xl font-semibold">
            {{ borrower?.first_name }} {{ borrower?.last_name }}
          </h6>
          <h6 v-else class="text-xl font-semibold">{{ borrower?.company_name }}</h6>
          <p>{{ borrower?.email }}</p>
        </div>
        <div
          class="cursor-pointer flex w-14 h-14 bg-main-blue rounded-xl ml-auto my-auto text-white"
          @click="toggleUserModal(borrower)"
        >
          <i class="w-6 h-6 mx-auto my-auto fa-solid fa-info"></i>
        </div>
      </div>
    </div>
    <div class="flex w-full justify-between">
      <Primary
        v-if="loanLending != null"
        class="w-[49%]"
        text="Call Governance Vote"
        color="blue"
        @click="showCreateVoteModal = true"
      />
      <Primary
        v-if="loanLending != null"
        class="w-[49%]"
        text="Call Funding Vote"
        color="green"
        @click="createGovVote()"
      />
      <Primary v-else class="w-4/5 mx-auto" text="Lend" color="green" @click="showPaymentModal = true" />
    </div>
  </Page>
</template>

<script setup lang="ts">
import { API } from "@/utils/api";
import { ref, onMounted, computed } from "vue";
import useLoanStore from "@/stores/loan";
import useUserStore from "@/stores/user";
import { useRoute } from "vue-router";
import type { Loan, Account, GovernanceVote } from "@/types";
import { notify } from "notiwind";

// Components
import Page from "@/components/Page.vue";
import Primary from "@/components/buttons/Primary.vue";
import Modal from "@/components/Modal.vue";
import Feed from "@/components/Feed.vue";
import Profile from "@/components/profile/Profile.vue";

const route = useRoute();
const loanStore = useLoanStore();
const userStore = useUserStore();

const loan = ref<Loan | null>(null);
const borrower = ref<Account | null>(null);

const showVotingModal = ref(false);
const showCreateVoteModal = ref(false);
const showPaymentModal = ref(false);
const showVoteInfoModal = ref(false);

const loanLending = computed(() => {
  return loanStore.my_loan_lendings.filter((lend) => lend.loan_id === loan.value?.id)[0];
});

const governanceVote = ref<GovernanceVote | null>(null);

const payoutDaysRemaining = computed(() => {
  if (loan.value === null) return 0;
  if (loan.value.funded_at !== "0001-01-01T00:00:00Z") {
    let offset = loan.value.number_of_payments;
    if (loan.value.payment_schedule === "weekly") {
      offset *= 7;
    }
    if (loan.value.payment_schedule === "monthly") {
      offset *= 30;
    }
    if (loan.value.payment_schedule === "yearly") {
      offset *= 365;
    }

    return offset;
  }

  return 0;
});

onMounted(async () => {
  if (!loanStore.hasFetched) {
    await loanStore.fetchLoans();
  }

  const loanId = Number(route.params.id);
  loan.value = loanStore.fetchLoan(loanId);

  if (loan.value !== null) {
    borrower.value = await userStore.fetchBorrower(loan.value.borrower_id);
  }
});

const voteInfo = ref<GovernanceVote | null>(null);

const toggleVoteInfoModal = (govVote: GovernanceVote | null): void => {
  showVoteInfoModal.value = !showVoteInfoModal.value;
  voteInfo.value = govVote;
};

const toggleVoteModal = (govVote: GovernanceVote | null): void => {
  showVotingModal.value = !showVotingModal.value;
  governanceVote.value = govVote;
};

const castVote = async (vote: boolean): void => {
  if (loan.value === null) return;

  try {
    await API.post(`/v1/loan/${loan.value.id}/gv/${governanceVote.value?.id}/vote`, {
      lender_id: userStore.user!.id,
      approves: vote,
    });
  } catch (e) {
    console.log(e);
    return;
  } finally {
    await loanStore.fetchLoans();
    loan.value = loanStore.fetchLoan(loan.value.id);
    toggleVoteModal(null);
  }
};

const title = ref<string>("");
const description = ref<string>("");

const closeCreateVote = (): void => {
  showCreateVoteModal.value = false;
  title.value = "";
  description.value = "";
};

const createGovVote = async (): Promise<void> => {
  // Check if funding vote is already going if so return
  if (title.value === "") {
    // check if funding requirement reached
    if (loan.value!.goal_amount * 0.75 > loan.value!.amount_raised) {
      notify(
        {
          group: "br",
          title: "Funding Requirement Not Met",
          text: "The funding requirement has not been met to call a funding vote(75% of goal required).",
        },
        4000
      );
      return;
    }

    let pass = true;
    loan.value!.governance_votes.forEach((govVote) => {
      if (govVote.title === "" && govVote.voting_active) {
        notify(
          {
            group: "br",
            title: "Funding Vote Already Active",
            text: "A funding vote is already active for this loan.",
          },
          4000
        );
        pass = false;
      }
    });

    description.value = "";

    if (!pass) return;
  }

  try {
    await API.post(`/v1/loan/${loan.value?.id}/gv`, {
      title: title.value,
      description: description.value,
    });
  } catch (e) {
    console.log(e);
  } finally {
    await loanStore.fetchLoans();
    loan.value = loanStore.fetchLoan(loan.value!.id);
    showCreateVoteModal.value = false;
    title.value = "";
    description.value = "";
  }

  notify(
    {
      group: "br",
      title: "Governance Vote Called",
      text: "A governance vote has been called for this loan.",
    },
    4000
  );
};

const hasVoted = (govVote: GovernanceVote): boolean => {
  return govVote.votes.filter((vote) => vote.lender_id === userStore.user.id).length > 0;
};

const paymentAmount = ref<number>(25);

// TODO - Raaj
const sendPayment = async (): Promise<void> => {};

const createLoanLender = async (): Promise<void> => {
  if (loan.value === null) return;

  try {
    await API.post(`/v1/loan/${loan.value.id}/lender`, {
      lender_id: userStore.user!.id,
      loan_amount: 25,
    });

    await sendPayment();
  } catch (e) {
    console.log(e);
    return;
  } finally {
    await loanStore.fetchLoans();
    loan.value = loanStore.fetchLoan(loan.value.id);
    showPaymentModal.value = false;
  }
};

const showUserModal = ref<Boolean>(false);
const userModalUser = ref<Account | null>(null);
const toggleUserModal = (userLocal: Account | null): void => {
  showUserModal.value = !showUserModal.value;
  userModalUser.value = userLocal;
};
</script>

<style scoped></style>
