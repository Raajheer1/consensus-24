<template>
  <Modal v-if="createPostModal" title="Create Post" width="sm" @close="toggleCreatePostModal()">
    <p class="font-bold text-gray-600 text-sm">Post Title</p>
    <input
      v-model="title"
      class="mb-2 py-1 border-b hover:border-b-main-blue outline-0 focus:border-transparent focus:ring-0 bg-transparent"
      placeholder="A weekly update"
    />

    <p class="mt-2 font-bold text-gray-600 text-sm">Description</p>
    <textarea
      v-model="description"
      class="w-full mb-2 py-1 border-b hover:border-b-main-blue outline-0 focus:border-transparent focus:ring-0 bg-transparent"
      placeholder="lorem ipsum dolor sit amet, consectetur adipiscing elit."
    />

    <div class="text-center mt-10 space-x-10">
      <button
        class="px-8 py-2 text-sm text-gray-600 focus:outline-none hover:underline"
        @click="toggleCreatePostModal()"
      >
        Cancel
      </button>
      <button
        class="px-8 py-2 text-sm rounded text-white bg-green-500 focus:outline-none hover:bg-green-400"
        @click="pushFeedMsg()"
      >
        Send Msg
      </button>
    </div>
  </Modal>

  <div class="w-full">
    <div class="flex w-full justify-between p-2">
      <h2 class="text-lg font-bold">{{ props.feedType.charAt(0).toUpperCase() + props.feedType.slice(1) }} Feed</h2>
      <button
        class="px-4 py-1 text-sm rounded text-white bg-green-500 focus:outline-none hover:bg-green-400"
        @click="toggleCreatePostModal()"
      >
        <i class="fas fa-plus"></i>
        Post
      </button>
    </div>
    <div class="flex w-full">
      <Spinner v-if="loading" />
      <div v-else class="w-full max-h-[150px] overflow-y-auto">
        <p v-if="feed.length === 0" class="text-center text-xs text-gray-500">No posts so far :(</p>
        <div
          v-for="(msg, idx) in feed"
          v-else
          :key="msg.id"
          class="p-1"
          :class="idx % 2 == 0 ? 'bg-gray-200' : 'bg-white'"
        >
          <div class="flex my-auto">
            <img
              :src="msg.created_by_user.profile_pic"
              class="max-h-8 rounded-full inline-block mt-0.5"
              alt="profile_pic"
            />
            <div class="grid grid-cols-1 mx-4">
              <p class="text-sm -mb-1">
                {{ msg.created_by_user.company_name || msg.created_by_user.first_name }}
                <span class="text-xs font-light">{{ new Date(msg.created_at).toLocaleString() }}</span>
              </p>
              <p class="text-sm font-semibold -my-1">{{ msg.title }}</p>
            </div>
          </div>
          <div class="w-full px-2 text-sm">
            {{ msg.description }}
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from "vue";
import { Feed } from "@/types";
import { API } from "@/utils/api";
import useUserStore from "@/stores/user";
import Spinner from "@/components/animations/Spinner.vue";
import Modal from "@/components/Modal.vue";

const userStore = useUserStore();

const props = defineProps<{
  feedType: string;
  id: number;
}>();

const feed = ref<Feed[]>([]);
const loading = ref<Boolean>(true);

const createPostModal = ref<Boolean>(false);
const title = ref<string>("");
const description = ref<string>("");
const toggleCreatePostModal = (): void => {
  title.value = "";
  description.value = "";
  createPostModal.value = !createPostModal.value;
};

const fetchFeed = async (): Promise<void> => {
  try {
    const { data } = await API.get(`/v1/feed/${props.feedType}/${props.id}`);
    feed.value = data;
  } catch (e) {
    console.log(e);
  } finally {
    loading.value = false;
  }
};

const pushFeedMsg = async (): Promise<void> => {
  if (userStore.user === null) {
    return;
  }
  try {
    await API.post(`/v1/feed/${props.feedType}/${props.id}`, {
      title: title.value,
      description: description.value,
      created_by: userStore.user.id,
    });
  } catch (e) {
    console.log(e);
  } finally {
    await fetchFeed();
    toggleCreatePostModal();
  }
};

onMounted(() => {
  fetchFeed();
});
</script>

<style scoped></style>
