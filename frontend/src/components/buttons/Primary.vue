<template>
  <button
    v-if="color === 'blue'"
    type="button"
    class="bg-blue-400 hover:bg-blue-600 border border-blue-400 transition text-white font-bold py-2 px-4 rounded"
  >
    <i v-if="!!props.icon" :class="`fas fa-${props.icon} mr-2`"></i>{{ props.text }}
  </button>
  <button
    v-else-if="color === 'red'"
    type="button"
    class="bg-red-400 hover:bg-red-600 border border-red-400 transition text-white font-bold py-2 px-4 rounded"
  >
    <i v-if="!!props.icon" :class="`fas fa-${props.icon} mr-2`"></i>{{ props.text }}
  </button>
  <button
    v-else-if="color === 'yellow'"
    type="button"
    class="bg-yellow-400 hover:bg-yellow-600 border border-yellow-400 transition text-white font-bold py-2 px-4 rounded"
  >
    <i v-if="!!props.icon" :class="`fas fa-${props.icon} mr-2`"></i>{{ props.text }}
  </button>
  <button
    v-else-if="color === 'green'"
    type="button"
    class="bg-green-400 hover:bg-green-600 border border-green-400 transition text-white font-bold py-2 px-4 rounded"
  >
    <i v-if="!!props.icon" :class="`fas fa-${props.icon} mr-2`"></i>{{ props.text }}
  </button>
  <button
    v-else
    type="button"
    class="bg-gray-400 hover:bg-gray-600 border border-gray-400 transition text-white font-bold py-2 px-4 rounded"
  >
    <i v-if="!!props.icon" :class="`fas fa-${props.icon} mr-2`"></i>{{ props.text }}
  </button>
</template>
<script setup lang="ts">
const props = defineProps<{
  text: string;
  icon?: string;
  color?: string;
}>();
</script>
<style scoped></style>
