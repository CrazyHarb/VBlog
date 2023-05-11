import { createSSRApp } from "vue";
import App from "./App.vue";
import VueQrcode from "@chenfengyuan/vue-qrcode";
import { createHead } from "@unhead/vue";
import { createPinia } from "pinia";
import { createRouter } from "./router";
import { febore } from "./router/router-guards";
import { i18n } from "./lang";
import naive from "naive-ui";
import { setup } from "@css-render/vue3-ssr";
import Particles from "vue3-particles";

export function createApp() {
  const app = createSSRApp(App);
  app.component(VueQrcode.name, VueQrcode);

  // 挂载head插件
  const head = createHead();
  app.use(head);

  // 挂载状态管理
  const pinia = createPinia();
  app.use(pinia);

  // 挂载路由
  const router = createRouter();
  febore(router, pinia);
  app.use(router);

  //国际化语言
  app.use(i18n);

  const { collect } = setup(app);
  app.use(naive);
  app.use(Particles);
  return { app, router, pinia, head, collect };
}
