<template>
  <div v-if="open" class="fixed w-full h-full top-0 left-0 flex justify-center z-10">
    <div class="absolute w-full h-full bg-gray-900 opacity-50" @click="close"></div>

    <div class="absolute mt-32 max-h-full w-[90%]">
      <div class="container bg-white overflow-hidden rounded-xl">
        <div
          class="px-4 py-4 leading-none flex justify-between items-center font-medium text-sm bg-gray-100 border-b select-none"
        >
          <h3 class="text-xl">{{ title }}</h3>
          <div class="text-2xl hover:text-gray-600 cursor-pointer" @click="close">&#215;</div>
        </div>

        <div class="max-h-full px-4 py-4">
          <slot></slot>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    title: {
      type: String,
      default: "",
    },
    header: {
      type: String,
      required: false,
      default: "",
    },
  },
  data() {
    return {
      open: true,
    };
  },
  computed: {},
  mounted() {
    const onEscape = (e) => {
      if (e.key === "Esc" || e.key === "Escape") {
        this.close();
      }
    };

    document.addEventListener("keydown", onEscape);
  },
  beforeUnmount() {
    const onEscape = (e) => {
      if (e.key === "Esc" || e.key === "Escape") {
        this.close();
      }
    };

    document.removeEventListener("keydown", onEscape);
  },
  methods: {
    close() {
      this.open = false;
      this.$emit("close");
    },
  },
};
</script>
