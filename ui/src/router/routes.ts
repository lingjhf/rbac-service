import {
  NavigationGuardNext,
  RouteLocationNormalized,
  RouteRecordRaw,
} from 'vue-router'
import Cookies from 'js-cookie'
export abstract class RouteName {
  static get home() {
    return 'home'
  }
  static get tenant() {
    return 'tenant'
  }
  static get tenantChild() {
    return 'tenantChild'
  }
  static get user() {
    return 'user'
  }
  static get role() {
    return 'role'
  }
  static get permission() {
    return 'permission'
  }
  static get auth() {
    return 'auth'
  }
  static get signup() {
    return 'signup'
  }
  static get login() {
    return 'login'
  }
  static get unauthorized() {
    return 'unauthorized'
  }
  static get notFound() {
    return 'notFound'
  }
}

export const routes: RouteRecordRaw[] = [
  {
    path: '/test',
    name: 'test',
    component: () => import('@/pages/test.vue'),
  },
  {
    path: '/',
    name: RouteName.home,
    component: () => import('@/pages/home'),
    beforeEnter: [requiredSignup],
  },
  {
    path: '/tenant/:id',
    name: RouteName.tenant,
    component: () => import('@/pages/tenant'),
    beforeEnter: [requiredSignup],
  },
  {
    path: '/auth',
    name: RouteName.auth,
    redirect: { name: RouteName.login },
    component: () => import('@/pages/auth'),
    children: [
      {
        path: 'signup',
        name: RouteName.signup,
        component: () => import('@/pages/signup'),
      },
      {
        path: 'login',
        name: RouteName.login,
        component: () => import('@/pages/login'),
      },
    ],
  },
  {
    path: '/unauthorized',
    name: RouteName.unauthorized,
    component: () => import('@/pages/errors').then((v) => v.Unauthorized),
  },
  {
    path: '/:pathMatch(.*)*',
    name: RouteName.notFound,
    component: () => import('@/pages/errors').then((v) => v.NotFound),
  },
]

function requiredSignup(
  to: RouteLocationNormalized,
  form: RouteLocationNormalized,
  next: NavigationGuardNext
) {
  const token = Cookies.get('token')
  if (!token) {
    return next({ name: 'auth' })
  }
  return next()
}
