<template>
  <div
    class="pt-[45px] w-full flex justify-between lg:justify-end border-b border-gray-300 px-4 lg:px-10 py-2.5 m-0 my-auto text-main-white lg:text-gray-500 text-xl bg-main-blue lg:bg-transparent"
  >
    <button
      class="my-auto lg:hidden text-xl leading-10 rounded-md hover:text-main-gold hover:scale-110 transition-all"
      @click="toggleNav()"
    >
      <i class="fa-solid fa-bars"></i>
    </button>
    <img src="@/assets/logo.png" class="my-auto block lg:hidden w-auto max-h-10" alt="logo" />
    <div class="my-auto space-x-8">
      <div class="relative inline-block text-left">
        <button type="button" @click="toggleNotification">
          <i class="fa-regular fa-bell hover:text-main-gold lg:hover:text-main-blue"></i>
        </button>

        <div
          v-if="isNotificationOpen"
          class="absolute z-10 bg-white divide-y divide-gray-100 rounded-lg shadow w-64 right-0 top-full mt-2.5"
        >
          <ul class="py-2 text-sm text-gray-700 flex flex-col">
            <li class="py-2 px-4">
              <div class="flex my-auto justify-between">
                <div>
                  <h2 class="text-main-blue text-xl font-bold">Notifications</h2>
                </div>
                <div class="my-auto">
                  <router-link to="Profile#notifications" class="text-main-gold text-xs font-bold">
                    View All
                  </router-link>
                </div>
              </div>
            </li>
          </ul>
        </div>
      </div>

      <div class="relative inline-block text-left">
        <button type="button" @click="toggleProfile">
          <i class="fa-regular fa-user hover:text-main-gold lg:hover:text-main-blue"></i>
        </button>

        <div
          v-if="isProfileOpen"
          class="absolute z-10 bg-white divide-y divide-gray-100 rounded-lg shadow w-64 right-0 top-full mt-2.5"
        >
          <ul class="py-2 text-sm text-gray-700 flex flex-col">
            <li class="py-2 px-4">
              <div class="flex my-auto justify-between">
                <div>
                  <h2 class="text-main-blue text-xl font-bold -mb-1">{{ userStore.fullName }}</h2>
                  <p>{{ userStore.user?.email }}</p>
                </div>
              </div>
            </li>
            <router-link
              to="Profile"
              class="font-bold text-md py-2.5 px-4 hover:bg-gray-100 text-gray-600"
              active-class="text-main-gold"
            >
              Profile
            </router-link>
            <li class="border-t border-gray-200 mt-2 pt-2"></li>
            <a class="cursor-pointer" @click="userStore.logout()">
              <p class="font-bold text-md py-2.5 px-4 hover:bg-gray-100 text-gray-600">Sign Out</p>
            </a>
          </ul>
        </div>
      </div>
    </div>

    <MobileSidebar :class="!isNavOpen ? 'opacity-0 invisible' : 'opacity-100 visible'" @toggle="toggleNav()" />
  </div>
</template>
<script setup lang="ts">
import { ref } from "vue";
import MobileSidebar from "@/views/partials/MobileSidebar.vue";
import useUserStore from "@/stores/user";

const isProfileOpen = ref<boolean>(false);
const isNotificationOpen = ref<boolean>(false);
const isNavOpen = ref<boolean>(false);
const userStore = useUserStore();

const toggleNav = (): void => {
  isProfileOpen.value = false;
  isNavOpen.value = !isNavOpen.value;
};

const toggleProfile = (): void => {
  isNavOpen.value = false;
  isNotificationOpen.value = false;
  // sleep for 50ms cause im too lazy to actually fix the animation issue
  setTimeout(() => {
    isProfileOpen.value = !isProfileOpen.value;
  }, 50);
};

const toggleNotification = (): void => {
  isNavOpen.value = false;
  isProfileOpen.value = false;
  // sleep for 50ms cause im too lazy to actually fix the animation issue
  setTimeout(() => {
    isNotificationOpen.value = !isNotificationOpen.value;
  }, 50);
};
</script>
