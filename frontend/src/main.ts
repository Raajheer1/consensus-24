import "./style.css";
import { createApp } from "vue";
import { createPinia } from "pinia";
import Notifications from "notiwind";

import App from "./App.vue";
import router from "./router";

const app = createApp(App);
app.use(router).use(createPinia()).use(Notifications).mount("#app");
