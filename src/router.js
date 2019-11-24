import Vue from "vue";
import Router from "vue-router";

Vue.use(Router);

import Login from "@/components/User/Login.vue";
import ListParties from "@/components/Parties/ListParties.vue";
import Invites from "@/components/Invites/Invites.vue";
import Settings from "@/components/Settings/Settings.vue";

const router = new Router({
  mode: "history",
  linkExactActiveClass: 'is-active',
  routes: [
    {
      path: "/login",
      component: Login,
      name: "login",
      meta: { title: "Login" }
    },
    {
      path: "/parties",
      component: ListParties,
      name: "list_parties",
      meta: { title: "List Parties", requiresAuth: true }
    },
    {
      path: "/invites",
      component: Invites,
      name: "invites",
      meta: { title: "Invites", requiresAuth: true }
    },
    {
      path: "/settings",
      component: Settings,
      name: "settings",
      meta: { title: "Settings", requiresAuth: true }
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
