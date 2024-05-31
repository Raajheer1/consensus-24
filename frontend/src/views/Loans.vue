<template>
  <Page title="Loans">
    <div v-if="!loanStore.hasFetched" class="m-5">
      <Spinner />
    </div>
    <div v-else class="m-2">
      <div v-if="loanStore.loans.length > 0" class="grid grid-cols-1 gap-y-6">
        <div v-for="loan in loanStore.loans" :key="loan.id" class="rounded-t-xl border-b-2 border-main-gold pb-4">
          <div v-if="dateInPast(new Date(loan.funded_at), new Date())">
            <img
              v-if="!!loan.image_url"
              class="w-full max-h-48 object-cover rounded-t-xl -p-1"
              :src="loan.image_url"
              alt="loan"
            />
            <div class="m-2">
              <div class="flex justify-between text-xl font-bold">
                <p class="text-main-gold">{{ loan.title }}</p>
                <p class="text-main-blue">+{{ loan.interest_rate }}%</p>
              </div>
              <div class="my-2 flex font-semibold text-main-grey justify-between">
                <div class="flex">
                  <p>${{ loan.amount_raised }} <span class="font-normal">raised</span></p>
                  <p class="mx-1">({{ Math.round((loan.amount_raised / loan.goal_amount) * 100) }}%)</p>
                </div>
                <div class="flex my-auto justify-center text-main-grey">
                  <i class="mt-1 mx-2 fa-solid fa-user"></i>
                  <p>{{ loan.lenders?.length || 0 }}</p>
                </div>
              </div>
              <Secondary class="w-full" text="View Loan" color="blue" @click="redirectToLoan(loan.id)" />
            </div>
          </div>
        </div>
      </div>
      <div v-else>
        <p class="text-main-grey text-center">No loans available</p>
      </div>
    </div>
  </Page>
</template>

<script setup lang="ts">
import { onMounted } from "vue";
import useLoanStore from "@/stores/loan";
import { useRouter } from "vue-router";

// Components
import Page from "@/components/Page.vue";
import Spinner from "@/components/animations/Spinner.vue";
import Secondary from "@/components/buttons/Secondary.vue";

const loanStore = useLoanStore();
const router = useRouter();

onMounted(() => {
  if (!loanStore.hasFetched) {
    loanStore.fetchLoans();
  }

  // loanStore.createGeneric();
});

const redirectToLoan = (loanId: number): void => {
  router.push(`/loan/${loanId}?return=loans`);
};

const dateInPast = (firstDate: Date, secondDate: Date): boolean => {
  return firstDate.setHours(0, 0, 0, 0) <= secondDate.setHours(0, 0, 0, 0);
};
</script>

<style scoped></style>
