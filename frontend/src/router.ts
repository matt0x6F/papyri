import { createRouter, createWebHashHistory, RouteRecordRaw } from "vue-router"

import Bookmarks from "./views/Bookmarks.vue"
import Settings from "./views/Settings.vue"
import Sources from "./views/Sources.vue"

const routes:Array<RouteRecordRaw> = [
  {
    path: "/",
    name: "Bookmarks",
    component: Bookmarks,
  },
  {
    path: "/settings",
    name: "Settings",
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: Settings,
  },
  {
    path: "/sources",
    name: "Sources",
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: Sources,
  },
];

const router = createRouter({
  history: createWebHashHistory(),
  routes,
});

export default router;