<template>
  <div
    class="flex justify-center items-center"
    :style="{ width: `${props.width}px`, height: `${props.height}px`, background: props.background }"
  >
    <img src="@/assets/logo-circle.png" class="w-20 animate-spin" />
  </div>
</template>

<script lang="ts" setup>
interface Props {
  width?: number;
  height?: number;
  background?: string;
}

const props = withDefaults(defineProps<Props>(), {
  width: 200,
  height: 200,
  background: "#6DAFFE",
});
</script>

<style scoped></style>
