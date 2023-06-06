import {createRouter, createWebHashHistory, RouteRecordRaw} from "vue-router";


const routes: RouteRecordRaw[] = [
    {
        path: "/",
        name:"home",
        component: ()=> import("../pages/Home.vue"),
    },
    {
        path: "/about",
        name:"about",
        component: ()=> import("../pages/About.vue"),
    }
]

export  default  createRouter({
    history: createWebHashHistory(),
    routes
})