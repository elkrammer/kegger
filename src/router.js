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

export default router;
