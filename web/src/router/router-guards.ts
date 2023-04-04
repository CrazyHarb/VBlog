import { useHead } from "@unhead/vue";
import { loadingBar } from "@/utils/client";
import { useProjectSettingStore } from "@/store/modules/projectSetting";

const getCookie = (name) =>
  document.cookie.match(`[;\s+]?${name}=([^;]*)`)?.pop();

export function febore(router, pinia) {
  const settingStore = useProjectSettingStore(pinia);

  router.beforeEach(async (to, from, next) => {
    settingStore.Load.start();
    loadingBar.start();

    if (!to.name) {
      //判断有没有路由
      next({ name: "exception-404" });
    }
    if (to.meta.title != undefined) {
      useHead({ title: to.meta.title });
    }
    if (
      !import.meta.env.SSR &&
      to.matched[0].name == "admin" &&
      !getCookie("token")
    ) {
      //前往后台，判断是否登陆
      next({ name: "login" });
    } else if (to.name == "login") {
      //前往登陆，判断是否登陆
      if (!import.meta.env.SSR && getCookie("token")) {
        next({ name: "admin" });
      } else {
        next();
      }
    } else {
      next();
    }
  });

  router.afterEach((to, _, failure) => {
    loadingBar.finish();
    setInterval(() => {
      settingStore.Load.end();
    }, 2000);
  });

  router.onError((error) => {
    console.log(error, "路由错误");
  });
}
