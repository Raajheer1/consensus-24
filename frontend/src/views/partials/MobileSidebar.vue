<template>
  <div class="fixed top-[60px] left-0 w-screen h-screen z-50 transition-all lg:invisible lg:opacity-0">
    <div class="top-0 left-0 w-full h-full absolute bg-black bg-opacity-80 content-['']" @click="$emit('toggle')"></div>
    <div class="relative bg-main-blue w-[275px] max-w-[80%] p-4 text-sm rounded-br">
      <div v-for="(link, idx) in SidebarLinks" :key="idx">
        <div v-if="link.separator" class="mt-2 flex py-1 my-auto">
          <h2 v-if="link.separatorTitle !== undefined" class="mr-4 text-main-white text-md font-bold">
            {{ link.separatorTitle }}
          </h2>
          <hr class="bg-extra-dark-purple dark:bg-extra-dark-purple my-2.5 flex-grow" />
        </div>
        <router-link
          v-if="link.to"
          class="flex rounded p-3 hover:bg-white text-main-white hover:bg-opacity-20 hover:text-main-gold items-center justify-between my-0.5"
          active-class="bg-white bg-opacity-10 font-bold text-main-gold"
          :to="link.to"
          @click="$emit('toggle')"
        >
          <div class="flex items-center gap-x-4">
            <span class="w-4 mx-auto text-center">
              <i :class="link.icon"></i>
            </span>
            <h2>{{ link.title }}</h2>
          </div>
          <button v-if="link.subLinks" @click="toggleSubLink(link.title)">
            <i :class="showSubLinks.includes(link.title) ? 'fas fa-chevron-down' : 'fas fa-chevron-up'"></i>
          </button>
        </router-link>
        <div
          v-else
          class="flex rounded p-3 hover:bg-white text-main-white hover:bg-opacity-20 hover:text-main-gold items-center justify-between my-0.5"
          active-class="bg-white bg-opacity-10 font-bold text-main-gold"
          @click="$emit('toggle')"
        >
          <div class="flex items-center gap-x-4">
            <span class="w-4 mx-auto text-center">
              <i :class="link.icon"></i>
            </span>
            <h2>{{ link.title }}</h2>
          </div>
          <button v-if="link.subLinks" @click="toggleSubLink(link.title)">
            <i :class="showSubLinks.includes(link.title) ? 'fas fa-chevron-down' : 'fas fa-chevron-up'"></i>
          </button>
        </div>
        <div v-if="showSubLinks.includes(link.title)">
          <div v-for="(subLink, index) in link.subLinks" :key="index" class="pl-4">
            <router-link
              v-if="subLink.to"
              class="flex rounded p-3 hover:bg-white text-main-white hover:bg-opacity-20 hover:text-main-gold items-center justify-between my-0.5"
              active-class="bg-white bg-opacity-10 font-bold text-main-gold"
              :to="subLink.to"
              @click="$emit('toggle')"
            >
              <div class="flex items-center gap-x-4">
                <span class="w-4 mx-auto text-center">
                  <i :class="subLink.icon"></i>
                </span>
                <h2>{{ subLink.title }}</h2>
              </div>
            </router-link>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { onMounted, ref } from "vue";
import SidebarLinks from "@/data/sidebar";

defineEmits(["toggle"]);

const showSubLinks = ref<string[]>([]);

const toggleSubLink = (title: string): void => {
  if (showSubLinks.value?.includes(title)) {
    showSubLinks.value = showSubLinks.value?.filter((link) => link !== title);
  } else {
    showSubLinks.value?.push(title);
  }
};

onMounted(() => {
  SidebarLinks.forEach((link) => {
    if (link.showSubLinks) {
      showSubLinks.value?.push(link.title);
    }
  });
});
</script>

<style scoped></style>
