import { defineConfig } from "vite";
import mkcert from "vite-plugin-mkcert";
import vue from "@vitejs/plugin-vue";
import VueTypeImports from "vite-plugin-vue-type-imports";
import { fileURLToPath, URL } from "url";
import path from "path";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    mkcert({
      savePath: "./certs",
      force: true,
      hosts: ["localhost"],
    }),
    vue(),
    VueTypeImports(),
  ],
  resolve: {
    alias: {
      "@": fileURLToPath(new URL("./src", import.meta.url)),
    },
  },
  server: {
    watch: {
      usePolling: true,
    },
  },
});
