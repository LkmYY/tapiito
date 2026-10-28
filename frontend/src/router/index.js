import { createRouter, createWebHistory } from 'vue-router'
import Login from '../pages/Login.vue'
import AdminNfcList from '../pages/AdminNfcList.vue'
import AdminNfcForm from '../pages/AdminNfcForm.vue'
import AdminStats from '../pages/AdminStats.vue'
import PublicLanding from '../pages/PublicLanding.vue'
import PreviewPage from '../pages/PreviewPage.vue'
import ShowroomPage from '../pages/ShowroomPage.vue'
import OrderPage from '../pages/OrderPage.vue'
import ConsultPage from '../pages/ConsultPage.vue'
import AdminOrderList from '../pages/AdminOrderList.vue'
import HomeLanding from '../pages/HomeLanding.vue'
import AboutPage from '../pages/AboutPage.vue'
import LegalPage from '../pages/LegalPage.vue'
import FeedbackPage from '../pages/FeedbackPage.vue'
import AdminFeedbackList from '../pages/admin/AdminFeedbackList.vue'

const routes = [
  { path: '/', component: HomeLanding },
  { path: '/about', component: AboutPage },
  { path: '/privacy', component: LegalPage },
  { path: '/disclaimer', component: LegalPage },
  { path: '/feedback', component: FeedbackPage },
  { path: '/login', component: Login },
  { path: '/admin/nfc', component: AdminNfcList },
  { path: '/admin/nfc/new', component: AdminNfcForm },
  { path: '/admin/nfc/:id/edit', component: AdminNfcForm },
  { path: '/admin/stats', component: AdminStats },
  { path: '/admin/orders', component: AdminOrderList },
  { path: '/admin/feedback', component: AdminFeedbackList },
  { path: '/p/:nfcId', component: PublicLanding },
  { path: '/preview', component: PreviewPage },
  { path: '/showroom', component: ShowroomPage },
  // 下单页已下架：明码标价撤掉，改为预约咨询
  { path: '/order', redirect: '/consult' },
  { path: '/consult', component: ConsultPage }
]

const router = createRouter({
  history: createWebHistory(),
  routes,
  // 路由切换后把滚动条带回顶部。
  // 之前没有 scrollBehavior，Vue Router 会保留上一页的滚动位置，
  // 导致从长页面（首页/展厅）切到短页面时停在半截。
  // 这里不返回 savedPosition，前进后退统一回顶部，行为可预期。
  scrollBehavior(to) {
    if (to.hash) return { el: to.hash, behavior: 'smooth' }
    return { top: 0, left: 0 }
  },
})

const PUBLIC_PATHS = ['/', '/about', '/privacy', '/disclaimer', '/feedback', '/login']

router.beforeEach((to, from, next) => {
  const token = localStorage.getItem('token')

  if (to.path === '/login' && token) {
    next('/admin/nfc')
    return
  }

  const isPublic = PUBLIC_PATHS.includes(to.path) ||
    to.path.startsWith('/p/') ||
    to.path === '/preview' ||
    to.path === '/showroom' ||
    to.path === '/order' ||
    to.path === '/consult'

  if (isPublic) {
    next()
  } else if (token) {
    next()
  } else {
    next('/login')
  }
})

export default router
