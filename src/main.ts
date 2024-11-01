import { createApp } from "vue";
import SimpleAnalytics from "simple-analytics-vue";
import App from "./App.vue";
import router from "./router"

import "./assets/main.css";

createApp(App).use(router).use(SimpleAnalytics, {skip: import.meta.env.NODE_ENV !== "production"}).mount("#app");