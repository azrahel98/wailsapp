import {
    NavigationGuardNext,
    RouteLocationNormalized,
    RouteRecordRaw,
    createRouter,
    createWebHashHistory
  } from 'vue-router'
  
import dashboard from '@views/dashboard.vue'
import login from '@views/login.vue'
import Main from '@views/dash/main.vue'
import Buscar from '@views/dash/buscar.vue'
  
  const routes: RouteRecordRaw[] = [
    {
      path: '/login',
      name: 'login',
      component: login
    },
    {
      path: '/',
      name: 'home',
      component: dashboard,
      beforeEnter: (to, from, next) => middleware(to, from, next),
      children: [
        {
          path: '/',
          name: 'dashboard',
          component: Main
        },
        {
          path: '/buscar',
          name: 'buscar',
          component: Buscar
        },
        {
          path: 'perfil/:dni',
          name: 'perfil',
      component: () => import('@views/dash/perfil.vue')
        },
        {
          path: 'boleta',
          name: 'boleta',
      component: () => import('@views/dash/boleta.vue')
        },{
          path: 'add',
          name: 'add',
      component: () => import('@views/dash/add.vue')
        }
      ]
    },
  ]
  
  const middleware = async (
    _to: RouteLocationNormalized,
    _from: RouteLocationNormalized,
    next: NavigationGuardNext
  ) => {
    const token = localStorage.getItem('jwt')
    console.log(token)
    if (token == null) {
      return next({
        name: 'login'
      })
    }
  
    return next()
  }
  
  export const router = createRouter({
    history: createWebHashHistory(),
    routes,
    linkExactActiveClass: 'active',
    linkActiveClass: 'active'
  })