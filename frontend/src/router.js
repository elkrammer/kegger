import Vue from "vue";
import Router from "vue-router";

Vue.use(Router);

import Login from "@/components/User/Login.vue";
import ListParties from "@/components/Parties/ListParties.vue";
import Invites from "@/components/Invites/Invites.vue";
import Invitation from "@/components/Invites/Invitation.vue";
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
            path: "/invite/:id",
            component: Invitation,
            name: "invitation",
            meta: { title: "View Invitation" },
            props: true
        },
        {
            path: "*",
            redirect: "/login"
        }
    ]
});

router.beforeEach((to, from, next) => {
    document.title = to.meta.title; //set title of each route
    let isLoggedIn = router.app.$options.store.getters["user/user"]
    let accessToken = localStorage.getItem("accessToken");
    let refreshToken = localStorage.getItem("refreshToken");

    if (accessToken) {
        router.app.$options.store.dispatch("user/setUserAndTokens", {
            accessToken: accessToken,
            refreshToken: refreshToken
        });
    }

    if (isLoggedIn && accessToken && refreshToken && to.meta.requiresAuth) {
        // user is logged in
        next();
    } else if (!isLoggedIn && !to.meta.requiresAuth) {
        // user is not logged in but the page doesn't require authentication
        next();
    } else {
        // user needs to authenticate
        next({ name: "login" });
    }
});

export default router;
