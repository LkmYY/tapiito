import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import { createPinia } from 'pinia'
import ElementPlus from 'element-plus'
import { MotionPlugin } from '@vueuse/motion'
import 'element-plus/dist/index.css'
import './assets/main.css'
import './styles/theme.css'

import {
  Plus, Search, Edit, Trash2, Save,
  BarChart2, X, User, Lock, Loader2,
  Phone, Mail
} from 'lucide-vue-next'

const app = createApp(App)
app.use(createPinia())
app.use(router)
app.use(ElementPlus)
app.use(MotionPlugin)

const icons = { Plus, Search, Edit, Trash2, Save, BarChart2, X, User, Lock, Loader2, Phone, Mail }
Object.entries(icons).forEach(([name, comp]) => app.component(name, comp))

// 预渲染会在 #app 里留下静态 HTML（给爬虫看）。
// 浏览器端渲染前必须清空，否则两份内容并存会出现重复文案。
const rootEl = document.querySelector('#app')
if (rootEl) rootEl.innerHTML = ''

app.mount('#app')
