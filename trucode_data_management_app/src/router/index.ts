import { createRouter, createWebHistory } from 'vue-router'
import RegisterView from '@/views/RegisterView.vue'
import LoginView from '@/views/LoginView.vue'
import PeopleDataView from '@/views/PeopleDataView.vue'
import DashboardView from '@/views/DashboardView.vue'
import { useAuthStore } from '@/stores/auth'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/register',
      name: 'register',
      component: RegisterView
    },
    {
      path: '/login',
      name: 'login',
      component: LoginView
    },
    {
      path: '/peopleData',
      name: 'peopleData',
      component: PeopleDataView,
      meta: {
        requiresAuth: true
      }
    },
    {
      path: '/dashboard',
      name: 'dashboard',
      component: DashboardView,
      meta: {
        requiresAuth: true
      }
    },
    {
      path: '/',
      redirect: '/login'
    }
  ]
})

router.beforeEach((to, from, next) => {
  const authStore = useAuthStore()
  if (to.meta.requiresAuth && !authStore.isLoggedIn) {
    return next({ name: 'login' });
  }
  if (authStore.isLoggedIn && (to.name === 'login' || to.name === 'register')) {
    return next({ name: 'peopleData' });
  }
  if (authStore.isLoggedIn && to.name === 'dashboard' && from.name === 'peopleData') {
    return next();
  }
  next();
})

export default router
