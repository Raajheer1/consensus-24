import { createRouter, createWebHistory, NavigationGuard, Router } from "vue-router";
import useUserStore from "@/stores/user";
// import apiUrl from "@/utils/api";

declare module "vue-router" {
  interface RouteMeta {
    requiresAuth?: boolean;
    requiresRole?: string[] | string;
  }
}

const routes = [
  {
    path: "/",
    name: "Home",
    component: () => import("@/views/Home.vue"),
    meta: {
      requiresAuth: true,
    },
  },
  {
    path: "/ramp",
    name: "Ramp",
    component: () => import("@/views/Ramp.vue"),
    meta: {
      requiresAuth: true,
    },
  },
  {
    path: "/loans",
    name: "Loans",
    component: () => import("@/views/Loans.vue"),
    meta: {
      requiresAuth: true,
    },
  },
  {
    path: "/loan/:id",
    name: "Loan",
    component: () => import("@/views/LoanInfo.vue"),
    meta: {
      requiresAuth: true,
    },
  },
  {
    path: "/profile",
    name: "Profile",
    component: () => import("@/views/Profile.vue"),
    meta: {
      requiresAuth: true,
    },
  },
];

const buildRouter = (): Router => {
  const r = routes;

  return createRouter({
    scrollBehavior(to) {
      if (to.params.section) {
        return {
          el: `#${to.params.section}`,
          top: 75,
          behavior: "smooth",
        };
      }
      return {
        top: 0,
        behavior: "smooth",
      };
    },
    history: createWebHistory(),
    routes: r,
  });
};

const router = buildRouter();

const check: NavigationGuard = (to, from, next): void => {
  const userStore = useUserStore();

  if (!userStore.user) {
    // TODO - Redirect to login page
    // window.location.href = `${apiUrl}/v3/user/login?redirect=${window.location.href}`;
    // return;
  }

  next();
};

router.beforeEach((to, from, next) => {
  const userStore = useUserStore();
  if (!userStore.hasFetched) {
    if (userStore.loading === null || userStore.loading === undefined) {
      userStore.loading = userStore.fetchUser();
    }
    userStore.loading.then(() => {
      check(to, from, next);
    });
  } else {
    check(to, from, next);
  }
});

export default router;
