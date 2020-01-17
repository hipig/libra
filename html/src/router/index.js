import Vue from 'vue'
import VueRouter from 'vue-router'
import Login from '@/views/login/Login'
import Main from '@/views/main/Main'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'login',
    component: Login
  },
  {
    path: '/admin',
    name: 'admin',
    component: Main
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
