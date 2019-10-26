import Vue from "vue";
import Router from "vue-router";

Vue.use(Router);

import HelloWorld from "@/components/HelloWorld.vue";
import Login from "@/components/User/Login.vue";

const router = new Router({
  mode: "history",
  routes: [
    {
      path: "/login",
      component: Login,
      name: "login",
      meta: { title: "Login" }
    },
    {
      path: "/home",
      component: HelloWorld,
      name: "home",
      meta: { title: "Home", requiresAuth: true }
    },
    {
      path: "*",
      redirect: "/login"
    }
  ]
});

router.beforeEach((to, from, next) => {
  // set the title of each route
  document.title = to.meta.title;
  let accessToken = localStorage.getItem("accessToken") ? localStorage.getItem("accessToken"): null;
  let refreshToken = localStorage.getItem("refreshToken") ? localStorage.getItem("refreshToken"): null;
  if (accessToken) {
    router.app.$options.store.dispatch("user/setUserAndTokens", {
      accessToken: accessToken,
      refreshToken: refreshToken
    });
  }

  // if the user is not logged in do not allow acess into protected pages.
  if (to.meta.requiresAuth && !router.app.$options.store.getters["user/user"]) {
    next({ name: "login" });
  }

  next();

});

export default router;
