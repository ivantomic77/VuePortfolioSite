import { createApp } from "vue";
import SimpleAnalytics from "simple-analytics-vue";
import App from "./App.vue";

import "./assets/main.css";

createApp(App).use(SimpleAnalytics, {skip: false}).mount("#app");
