<template>
  <div class="border-b">
    <h3 class="font-bold">Personal Details</h3>
  </div>
  <div class="m-5">
    <div v-if="!!editableUser" class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-x-5 gap-y-4">
      <div class="col-span-1 flex flex-col">
        <p class="font-bold text-gray-600 text-sm">First name</p>
        <input
          v-model="editableUser.first_name"
          class="my-1 py-1 border-b hover:border-b-main-blue outline-0 focus:border-transparent focus:ring-0 bg-transparent"
          placeholder="John"
          disabled
        />
      </div>
      <div class="col-span-1 flex flex-col">
        <p class="font-bold text-gray-600 text-sm">Last name</p>
        <input
          v-model="editableUser.last_name"
          class="my-1 py-1 border-b hover:border-b-main-blue outline-0 focus:border-transparent focus:ring-0 bg-transparent"
          placeholder="Doe"
          disabled
        />
      </div>
      <div class="col-span-1 md:col-span-2 flex flex-col">
        <p class="font-bold text-gray-600 text-sm">Email</p>
        <input
          v-model="editableUser.email"
          class="my-1 py-1 border-b hover:border-b-main-blue outline-0 focus:border-transparent focus:ring-0 bg-transparent"
          placeholder="bob@bobbybbaby.com"
          disabled
        />
      </div>
      <div class="col-span-1 md:col-span-2 flex flex-col">
        <p class="font-bold text-gray-600 text-sm">Public Key</p>
        <input
          v-model="editableUser.public_key"
          class="my-1 py-1 border-b hover:border-b-main-blue outline-0 focus:border-transparent focus:ring-0 bg-transparent"
          placeholder="abcddd"
          disabled
        />
      </div>
      <div v-if="props.showSecret" class="col-span-1 md:col-span-2 flex flex-col">
        <p class="font-bold text-gray-600 text-sm">Private Key</p>
        <input
          v-model="editableUser.secret_key"
          class="my-1 py-1 border-b hover:border-b-main-blue outline-0 focus:border-transparent focus:ring-0 bg-transparent"
          placeholder="ABCD****"
        />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from "vue";

import { Account } from "@/types";
import { b } from "vite/dist/node/types.d-aGj9QkWt";

const props = defineProps<{
  user: Account | null;
  showSecret?: boolean;
}>();

const editableUser = ref<Account | null>(props.user);

onMounted(() => {
  if (editableUser.value) {
    if (editableUser.value?.secret_key) {
      editableUser.value.secret_key = `${editableUser.value?.secret_key.substring(0, 4)}*****************`;
    }
  }
});
</script>

<style scoped></style>
