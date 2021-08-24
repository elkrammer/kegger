import Vue from "vue";
import Router from "vue-router";

import store from "@/store/";

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
            redirect: "/parties",
            meta: { requiresAuth: true },
        }
    ]
});

// verify if user is authenticated
router.beforeEach(async (to, from, next) => {
    let accessToken = localStorage.getItem("accessToken");

    if (!to.meta.requiresAuth) {
      next();
    }

    else if (to.meta.requiresAuth && (accessToken || store.state.user.isLoggedIn)) {
        // user is already authenticated - validate existing access token
        await store.dispatch("user/validateToken", accessToken);
        next();
    }

    else if (to.meta.requiresAuth && (!accessToken || !store.state.user.isLoggedIn)) {
      next({ name: "login" });
    }

    else {
      next();
    }
});

export default router;
