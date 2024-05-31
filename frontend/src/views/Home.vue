<template>
  <div class="min-h-[100vh-61px] from-main-blue to-main-grey bg-gradient-to-b">
    <Page>
      <div class="w-full text-center mb-20 -mt-4">
        <p class="text-main-light-grey text-xl mt-0">Wallet Balance</p>
        <h2 class="text-white font-extrabold text-8xl">$175</h2>
      </div>
      <div class="-mx-6 flex flex-col">
        <div class="text-center text-white my-4 mx-2 bg-main-gold rounded-2xl -mt-12 py-2">
          <div class="grid grid-cols-2 divide-x divide-main-blue">
            <div class="w-full text-center">
              <h2 class="font-bold text-2xl underline decoration-main-blue">{{ loanStore.avgReturn }}%</h2>
              <p class="text-main-grey text-lg mt-0">Est. Return</p>
            </div>
            <div class="w-full text-center">
              <h2 class="font-bold text-2xl underline decoration-main-gold">{{ loanStore.my_loans.length }}</h2>
              <p class="text-main-grey text-lg mt-0">Active Loans</p>
            </div>
          </div>
        </div>

        <!------------------------------------------------- LOAN LIST ------------------------------------------------->
        <div v-if="loanStore.my_loan_lendings.length > 0" class="grid gap-y-1 grid-cols-1 mx-2">
          <div
            v-for="lending in loanStore.my_loan_lendings"
            :key="lending.id"
            class="cursor-pointer rounded-2xl bg-main-grey"
            @click="redirectToLoan(lending.loan.id)"
          >
            <div v-if="lending.loan.borrower.id !== 0" class="flex w-full px-2 py-2">
              <img
                class="w-14 h-14 rounded-full border-2 border-main-gold"
                :src="lending.loan.borrower.profile_pic"
                alt="Profile Picture"
              />
              <div class="mx-2.5 my-auto text-white w-1/2">
                <h6 class="text-xl font-semibold text-left truncate">{{ lending.loan.title }}</h6>
                <p class="text-main-light-grey">
                  {{ lending.loan_amount }} {{ lending.loan.borrower.company_name }} Loan Tkns
                </p>
              </div>
              <div class="flex ml-auto text-3xl my-auto text-green-400">+{{ lending.loan.interest_rate }}%</div>
            </div>
          </div>
        </div>
      </div>
    </Page>
  </div>
</template>
<script setup lang="ts">
import Page from "@/components/Page.vue";

import useUserStore from "@/stores/user";
import useLoanStore from "@/stores/loan";
import { onMounted } from "vue";
import { useRouter } from "vue-router";

const router = useRouter();
const userStore = useUserStore();
const loanStore = useLoanStore();

onMounted(() => {
  if (!userStore.user) {
    return;
  }
  loanStore.fetchMyLoans(`${userStore.user.id}`);
});

const redirectToLoan = (loanId: number): void => {
  router.push(`/loan/${loanId}`);
};
</script>
