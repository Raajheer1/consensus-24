<template>
  <button
    v-if="color === 'blue'"
    type="button"
    class="text-blue-400 border border-blue-400 hover:bg-blue-100 transition text-sm font-bold py-2 px-4 rounded"
  >
    <i v-if="!!props.icon" :class="`fas fa-${props.icon} mr-2`"></i>{{ props.text }}
  </button>
  <button
    v-else-if="color === 'red'"
    type="button"
    class="text-red-400 border border-red-400 hover:bg-red-100 transition text-sm font-bold py-2 px-4 rounded"
  >
    <i v-if="!!props.icon" :class="`fas fa-${props.icon} mr-2`"></i>{{ props.text }}
  </button>
  <button
    v-else-if="color === 'yellow'"
    type="button"
    class="text-yellow-400 border border-yellow-400 hover:bg-yellow-100 transition text-sm font-bold py-2 px-4 rounded"
  >
    <i v-if="!!props.icon" :class="`fas fa-${props.icon} mr-2`"></i>{{ props.text }}
  </button>
  <button
    v-else-if="color === 'green'"
    type="button"
    class="text-green-400 border border-green-400 hover:bg-green-100 transition text-sm font-bold py-2 px-4 rounded"
  >
    <i v-if="!!props.icon" :class="`fas fa-${props.icon} mr-2`"></i>{{ props.text }}
  </button>
  <button
    v-else
    type="button"
    class="text-gray-400 border border-gray-400 hover:bg-gray-100 transition text-sm font-bold py-2 px-4 rounded"
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
