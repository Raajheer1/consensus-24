<template>
  <div class="p-6 md:p-12">
    <div v-if="props.button" class="flex justify-between max-w-screen">
      <div class="w-3/4">
        <h1 v-if="props.title.length > 0" class="text-main-blue text-3xl font-bold">{{ props.title }}</h1>
        <p v-if="props.subtitle.length > 0" class="text-xs text-main-grey">{{ props.subtitle }}</p>
      </div>
      <Primary
        v-if="props.buttonType == 'primary'"
        :text="props.buttonText"
        :to="props.buttonTarget"
        color="red"
        class="my-auto"
        @click="buttonRedirect"
      >
        {{ props.buttonText }}
      </Primary>
      <Secondary v-else :text="props.buttonText" :to="props.buttonTarget" class="my-auto" @click="buttonRedirect">
        {{ props.buttonText }}
      </Secondary>
    </div>
    <h1 v-else class="text-main-blue text-3xl font-bold">{{ props.title }}</h1>
    <div class="flex divide-x divide-gray-300 mt-6">
      <div class="w-full space-y-6">
        <slot>
          <div class="m-5">
            <p>Failed to render page. If you are viewing this page please open a ticket.</p>
          </div>
        </slot>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import Primary from "@/components/buttons/Primary.vue";
import Secondary from "@/components/buttons/Secondary.vue";
import { useRouter } from "vue-router";

const router = useRouter();

const props = defineProps({
  title: { type: String, required: false, default: "" },
  subtitle: { type: String, required: false, default: "" },
  button: { type: Boolean, default: false },
  buttonType: { type: String, default: "primary" },
  buttonText: { type: String, default: "Back" },
  buttonTarget: { type: String, default: "/" },
});

const buttonRedirect = (): void => {
  router.push(props.buttonTarget);
};
</script>

<style scoped></style>
