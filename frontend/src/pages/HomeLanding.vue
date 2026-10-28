<template>
  <div class="min-h-screen bg-[#fbfaf7] font-sans text-[#14181d] selection:bg-[#ff6b35]/20">
    <!-- ============ Hero ============ -->
    <header class="relative overflow-hidden">
      <div class="pointer-events-none absolute inset-0">
        <div class="absolute -left-24 -top-32 h-[420px] w-[420px] rounded-full bg-[#ff6b35]/15 blur-[100px]" />
        <div class="absolute right-[-10%] top-[-10%] h-[380px] w-[380px] rounded-full bg-[#ffd166]/30 blur-[100px]" />
        <div class="absolute bottom-[-180px] left-1/3 h-[420px] w-[420px] rounded-full bg-[#7bdff2]/20 blur-[110px]" />
        <div class="hero-grid absolute inset-0" />
      </div>

      <SiteHeader transparent class="relative z-20" />

      <div class="relative z-10 mx-auto max-w-6xl px-6 pb-16 pt-10 md:pb-24 md:pt-16">
        <div class="grid items-center gap-12 md:grid-cols-[minmax(0,1fr)_320px]">
          <div>
            <div
              v-motion
              :initial="{ opacity: 0, y: 14 }"
              :enter="{ opacity: 1, y: 0, transition: { duration: 600 } }"
              class="mb-6 inline-flex items-center gap-2 rounded-full border border-[#14181d]/10 bg-white/70 px-3 py-1.5 text-[12px] font-medium text-[#14181d]/70 backdrop-blur"
            >
              <span class="relative flex h-2 w-2">
                <span class="absolute inline-flex h-full w-full animate-ping rounded-full bg-[#ff6b35] opacity-60" />
                <span class="relative inline-flex h-2 w-2 rounded-full bg-[#ff6b35]" />
              </span>
              Tap it alive.
            </div>

            <h1
              v-motion
              :initial="{ opacity: 0, y: 18 }"
              :enter="{ opacity: 1, y: 0, transition: { duration: 700, delay: 80 } }"
              class="text-[44px] font-semibold leading-[1.08] tracking-[-0.03em] md:text-[68px]"
            >
              碰一下，<br />
              它就<span class="alive-word">活了</span>
            </h1>

            <p
              v-motion
              :initial="{ opacity: 0, y: 18 }"
              :enter="{ opacity: 1, y: 0, transition: { duration: 700, delay: 160 } }"
              class="mt-6 max-w-[30rem] text-[16px] leading-[1.75] text-[#14181d]/60 md:text-[17px]"
            >
              一张薄薄的标签，贴上你的名片、菜单、门店或随身物品。
              手机一碰，内容自己跳出来——购买即完成首次写入，含 1 次免费内容调整。
            </p>

            <div
              v-motion
              :initial="{ opacity: 0, y: 18 }"
              :enter="{ opacity: 1, y: 0, transition: { duration: 700, delay: 240 } }"
              class="mt-9 flex flex-wrap items-center gap-3"
            >
              <router-link to="/showroom" class="cta-main">
                逛逛展厅
                <ArrowUpRight class="h-4 w-4" />
              </router-link>
            </div>

            <div class="mt-9 flex flex-wrap items-center gap-x-6 gap-y-2 text-[12px] text-[#14181d]/45">
              <span class="inline-flex items-center gap-1.5">
                <Check class="h-3.5 w-3.5 text-[#ff6b35]" /> 无需装 App
              </span>
              <span class="inline-flex items-center gap-1.5">
                <Check class="h-3.5 w-3.5 text-[#ff6b35]" /> 购买即完成首次写入，含 1 次免费调整
              </span>
              <span class="inline-flex items-center gap-1.5">
                <Check class="h-3.5 w-3.5 text-[#ff6b35]" /> 10 种类型 · 300 种风格
              </span>
            </div>
          </div>

          <!-- 右：会呼吸的标签卡 -->
          <div class="relative hidden md:block">
            <div class="tap-stage">
              <div class="tap-ring tap-ring-1" />
              <div class="tap-ring tap-ring-2" />
              <div class="tap-ring tap-ring-3" />
              <div class="tap-card">
                <div class="flex items-center justify-between">
                  <span class="font-mono text-[9px] uppercase tracking-[0.22em] text-[#14181d]/35">NFC · 9999</span>
                  <span class="h-1.5 w-1.5 rounded-full bg-[#ff6b35]" />
                </div>
                <div class="mt-6">
                  <p class="text-[26px] font-semibold leading-tight tracking-tight">拾壹造型</p>
                  <p class="mt-1 text-[11px] tracking-[0.06em] text-[#14181d]/45">预约名片 · 美业沙龙</p>
                </div>
                <div class="mt-6 space-y-2 border-t border-[#14181d]/8 pt-4">
                  <div class="flex items-center justify-between text-[11px]">
                    <span class="text-[#14181d]/40">设计师精剪</span>
                    <span class="font-mono tabular-nums">¥168</span>
                  </div>
                  <div class="flex items-center justify-between text-[11px]">
                    <span class="text-[#14181d]/40">法式烫</span>
                    <span class="font-mono tabular-nums">¥588</span>
                  </div>
                </div>
              </div>
              <div class="tap-hint">
                <MousePointerClick class="h-3.5 w-3.5" />
                碰一碰
              </div>
            </div>
          </div>
        </div>
      </div>
    </header>

    <!-- ============ 类型跑马灯 ============ -->
    <div class="border-y border-[#14181d]/8 bg-[#14181d] py-3">
      <div class="ticker-mask overflow-hidden">
        <div class="ticker-track">
          <span v-for="(item, i) in tickerItems" :key="i" class="ticker-item">
            {{ item }}
            <span class="ticker-dot" />
          </span>
        </div>
      </div>
    </div>

    <!-- ============ 类型布局（bento） ============ -->
    <section class="mx-auto max-w-6xl px-6 pb-20 pt-16 md:pt-24">
      <div class="mb-10 flex flex-wrap items-end justify-between gap-4">
        <div>
          <p class="mb-3 text-[11px] font-semibold uppercase tracking-[0.2em] text-[#ff6b35]">What's inside</p>
          <h2 class="text-[32px] font-semibold tracking-[-0.02em] md:text-[40px]">九种标签，各有各的活法</h2>
          <p class="mt-3 max-w-lg text-[14px] leading-relaxed text-[#14181d]/55">
            九种常用类型，每种都有 30 种设计风格。选一个最像你的，购买即完成首次写入；含 1 次免费内容调整。
          </p>
        </div>
      </div>

      <div class="grid grid-cols-1 gap-4 sm:grid-cols-2 lg:grid-cols-12">
        <router-link
          v-for="(t, i) in typeCards"
          :key="t.key"
          :to="`/showroom?type=${t.key}`"
          class="tcard group"
          :class="t.span"
        >
          <div class="flex items-start justify-between gap-3">
            <span class="ticon" :style="{ background: t.iconBg, color: t.iconColor }">
              <component :is="t.icon" class="h-[18px] w-[18px]" />
            </span>
            <span class="flex items-center gap-1.5">
              <span class="tcount">{{ t.styles.length }} 风格</span>
              <ArrowUpRight class="tarrow h-4 w-4" />
            </span>
          </div>

          <div class="mt-5">
            <p class="text-[15.5px] font-semibold tracking-tight">{{ t.label }}</p>
            <p class="mt-1.5 text-[12px] leading-relaxed text-[#14181d]/50">{{ t.desc }}</p>
          </div>

          <div class="tflip mt-4" aria-hidden="true">
            <div v-for="(column, ci) in t.columns" :key="ci" class="tcol">
              <div
                class="tcol-track"
                :class="{ 'tcol-snap': flipSnap }"
                :style="{ '--flip-index': flipStep, transitionDelay: `${i * CARD_STAGGER}ms` }"
              >
                <span
                  v-for="(s, si) in column"
                  :key="`${t.key}-${ci}-${si}`"
                  class="tthumb"
                  :style="{ background: s.gradient }"
                  :title="s.label"
                >
                  <img
                    :src="`/showroom/thumbs/${t.key}__${s.value}.jpg`"
                    alt=""
                    width="116"
                    height="164"
                    :loading="si === 0 ? 'eager' : 'lazy'"
                    :fetchpriority="si === 0 ? 'low' : 'auto'"
                    decoding="async"
                    draggable="false"
                    :class="{ 'timg--wide': templateTypes[t.key].layout === 'landscape' }"
                    @error="onThumbError"
                  />
                </span>
              </div>
            </div>
          </div>
        </router-link>
      </div>
    </section>

    <!-- ============ 三步 ============ -->
    <section class="border-y border-[#14181d]/8 bg-white/60">
      <div class="mx-auto max-w-6xl px-6 py-20">
        <div class="mb-12 max-w-lg">
          <p class="mb-3 text-[11px] font-semibold uppercase tracking-[0.2em] text-[#ff6b35]">How it works</p>
          <h2 class="text-[32px] font-semibold tracking-[-0.02em] md:text-[40px]">三步，让它活过来</h2>
        </div>

        <div class="grid gap-5 md:grid-cols-3">
          <div v-for="(step, i) in steps" :key="i" class="step-card group">
            <span class="step-num">{{ String(i + 1).padStart(2, '0') }}</span>
            <component :is="step.icon" class="mt-5 h-5 w-5 text-[#ff6b35]" />
            <h3 class="mt-4 text-[16px] font-semibold tracking-tight">{{ step.title }}</h3>
            <p class="mt-2 text-[12.5px] leading-relaxed text-[#14181d]/50">{{ step.desc }}</p>
          </div>
        </div>
      </div>
    </section>

    <!-- ============ 结尾 CTA ============ -->
    <section class="mx-auto max-w-6xl px-6 py-20">
      <div class="relative overflow-hidden rounded-[26px] bg-[#14181d] px-8 py-14 text-center text-white md:px-16">
        <div class="pointer-events-none absolute -left-16 -top-20 h-64 w-64 rounded-full bg-[#ff6b35]/30 blur-[90px]" />
        <div class="pointer-events-none absolute -bottom-24 -right-10 h-72 w-72 rounded-full bg-[#ffd166]/20 blur-[90px]" />
        <div class="relative">
          <p class="text-[12px] uppercase tracking-[0.24em] text-white/40">Tap it alive.</p>
          <h2 class="mx-auto mt-4 max-w-3xl text-balance text-[30px] font-semibold leading-[1.18] tracking-[-0.02em] md:text-[42px]">
            给你的东西，<br class="hidden md:block" />一个会说话的身份
          </h2>
          <p class="mx-auto mt-4 max-w-md text-[14px] leading-relaxed text-white/55">
            从一张标签开始，把该说的话都放进去。购买即完成首次写入，含 1 次免费内容调整。
          </p>
          <router-link to="/showroom" class="cta-invert mt-8">
            去展厅看看
            <ArrowUpRight class="h-4 w-4" />
          </router-link>
        </div>
      </div>
    </section>

    <!-- ============ 页脚 ============ -->
    <footer class="border-t border-[#14181d]/8">
      <div class="mx-auto flex max-w-6xl flex-col items-center justify-between gap-4 px-6 py-8 text-[12px] text-[#14181d]/40 sm:flex-row">
        <div class="flex items-center gap-2">
          <span class="brand-chip-sm">
            <img src="/tapiito-icon-white.png" alt="Tapiito" class="h-[15px] w-[15px]" draggable="false" />
          </span>
          <span class="font-medium text-[#14181d]/70">Tapiito</span>
          <span class="hidden sm:inline">· Tap it alive. 碰一下，它就活了</span>
        </div>
        <div class="flex flex-wrap items-center justify-center gap-x-5 gap-y-2 sm:justify-end">
          <router-link to="/showroom" class="hover:text-[#14181d]/70 transition-colors">展厅</router-link>
          <router-link to="/about" class="hover:text-[#14181d]/70 transition-colors">关于我们</router-link>
          <router-link to="/feedback" class="hover:text-[#14181d]/70 transition-colors">我有意见</router-link>
          <router-link to="/privacy" class="hover:text-[#14181d]/70 transition-colors">隐私政策</router-link>
          <router-link to="/disclaimer" class="hover:text-[#14181d]/70 transition-colors">免责声明</router-link>
          <span>&copy; {{ new Date().getFullYear() }} Tapiito</span>
        </div>
      </div>
    </footer>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, onBeforeUnmount } from 'vue'
import {
  Nfc, ArrowUpRight, ArrowRight, Check, MousePointerClick,
  ScanLine, PencilLine, Radio
} from 'lucide-vue-next'
import { templateTypes, visibleTypeOrder } from '../templates/registry.js'
import SiteHeader from '../components/SiteHeader.vue'

/** 每种类型的定位说明与 bento 跨列规则 */
const TYPE_META = {
  business_card: { desc: '碰一碰，联系方式直接存进通讯录', span: 'lg:col-span-5' },
  child_info: { desc: '需要帮助时，好心人一眼看到关键信息', span: 'lg:col-span-4' },
  emergency_contact: { desc: '紧急时刻，第一时间找到家人', span: 'lg:col-span-3' },
  company_website: { desc: '一张标签，就是一页企业官网', span: 'lg:col-span-3' },
  product_display: { desc: '产品参数、卖点、购买入口，一次讲清', span: 'lg:col-span-4' },
  recipe: { desc: '菜单、价格、招牌菜，一碰就看全', span: 'lg:col-span-5' },
  custom: { desc: '花店、面馆、健身房、书院，都能装下', span: 'sm:col-span-2 lg:col-span-12' },
  nightclub: { desc: '夜场社交，一张卡留下今晚的相遇', span: 'lg:col-span-4' },
  anime_card: { desc: '二次元角色闪卡，收藏与交换都合适', span: 'lg:col-span-4' },
  pet_tag: { desc: '宠物走失时，让路人第一时间联系到你', span: 'lg:col-span-4' }
}

const PAGE_SIZE = 5
const FLIP_PAGES = 6
const FLIP_MS = 520
const CARD_STAGGER = 20
const HOME_TYPE_KEYS = visibleTypeOrder.filter((key) => key !== 'custom')
const FLIP_SNAP_MS = FLIP_MS + (HOME_TYPE_KEYS.length - 1) * CARD_STAGGER + 40
const flipStep = ref(0)
const flipSnap = ref(false)
let flipTimer = null
let snapTimer = null
const prefersReducedMotion = typeof window !== 'undefined'
  && window.matchMedia?.('(prefers-reduced-motion: reduce)').matches

/** 5 列竖向翻页：每列 6 张，末尾补第一张做无缝回环 */
const buildColumns = (styles) =>
  Array.from({ length: PAGE_SIZE }, (_, col) => {
    const column = []
    for (let page = 0; page < FLIP_PAGES; page += 1) {
      const style = styles[page * PAGE_SIZE + col]
      if (style) column.push(style)
    }
    if (column.length) column.push(column[0])
    return column
  })

const typeCards = computed(() =>
  HOME_TYPE_KEYS.map((key) => {
    const t = templateTypes[key]
    const meta = TYPE_META[key] || {}
    return {
      ...t,
      desc: meta.desc || '',
      span: meta.span || 'lg:col-span-4',
      columns: buildColumns(t.styles)
    }
  })
)

const onThumbError = (event) => {
  event.currentTarget.style.display = 'none'
}

const tickerItems = computed(() => {
  const base = HOME_TYPE_KEYS.map((k) => templateTypes[k].label)
  return [...base, ...base, ...base]
})

const steps = [
  { icon: ScanLine, title: '写一张标签', desc: '把链接写进任意 NFC 标签，贴在名片、餐桌或书包上。' },
  { icon: MousePointerClick, title: '碰一下就打开', desc: '手机靠近，页面立刻出现，不用下载任何 App。' },
  { icon: PencilLine, title: '首次写入包含在购买内', desc: '购买即完成首次写入；含 1 次免费内容调整，后续调整另收制作费。' }
]

const resetFlip = () => {
  flipSnap.value = true
  flipStep.value = 0
  requestAnimationFrame(() => {
    requestAnimationFrame(() => {
      flipSnap.value = false
    })
  })
}

onMounted(() => {
  if (prefersReducedMotion) return
  flipTimer = setInterval(() => {
    if (document.hidden) return
    if (flipStep.value >= FLIP_PAGES) {
      resetFlip()
      return
    }
    flipStep.value += 1
    if (flipStep.value === FLIP_PAGES) {
      clearTimeout(snapTimer)
      snapTimer = setTimeout(resetFlip, FLIP_SNAP_MS)
    }
  }, 1600)
})

onBeforeUnmount(() => {
  clearInterval(flipTimer)
  clearTimeout(snapTimer)
})

</script>

<style scoped>
/* ---------- 背景网格 ---------- */
.hero-grid {
  background-image:
    linear-gradient(to right, rgba(20, 24, 29, 0.04) 1px, transparent 1px),
    linear-gradient(to bottom, rgba(20, 24, 29, 0.04) 1px, transparent 1px);
  background-size: 64px 64px;
  mask-image: radial-gradient(circle at 30% 20%, black, transparent 68%);
  -webkit-mask-image: radial-gradient(circle at 30% 20%, black, transparent 68%);
}

/* ---------- 品牌 ---------- */
.brand-chip {
  display: flex;
  height: 30px;
  width: 30px;
  align-items: center;
  justify-content: center;
  border-radius: 10px;
  transition: transform 0.3s cubic-bezier(0.22, 1, 0.36, 1);

  background: #15151B;}
.group:hover .brand-chip {
  transform: translateY(-1px) rotate(-4deg);
}
.brand-chip-sm {
  display: flex;
  height: 20px;
  width: 20px;
  align-items: center;
  justify-content: center;
  border-radius: 7px;

  background: #15151B;}

/* ---------- 导航按钮 ---------- */
.nav-ghost {
  border-radius: 10px;
  padding: 8px 14px;
  font-size: 13px;
  font-weight: 500;
  color: rgba(20, 24, 29, 0.65);
  transition: color 0.2s ease, background 0.2s ease;
}
.nav-ghost:hover {
  color: #14181d;
  background: rgba(20, 24, 29, 0.05);
}
.nav-solid {
  display: inline-flex;
  align-items: center;
  gap: 5px;
  border-radius: 10px;
  background: #14181d;
  padding: 8px 14px;
  font-size: 13px;
  font-weight: 600;
  color: #fff;
  transition: transform 0.2s cubic-bezier(0.22, 1, 0.36, 1), background 0.2s ease;
}
.nav-solid:hover {
  background: #ff6b35;
  transform: translateY(-1px);
}

/* ---------- Hero 主按钮 ---------- */
.cta-main {
  display: inline-flex;
  align-items: center;
  gap: 7px;
  border-radius: 12px;
  background: #14181d;
  padding: 13px 20px;
  font-size: 14px;
  font-weight: 600;
  color: #fff;
  transition: transform 0.22s cubic-bezier(0.22, 1, 0.36, 1), background 0.22s ease, box-shadow 0.22s ease;
}
.cta-main:hover {
  background: #ff6b35;
  transform: translateY(-2px);
  box-shadow: 0 14px 26px -14px rgba(255, 107, 53, 0.75);
}
.cta-ghost {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  border-radius: 12px;
  border: 1px solid rgba(20, 24, 29, 0.14);
  background: rgba(255, 255, 255, 0.6);
  padding: 12px 18px;
  font-size: 13.5px;
  font-weight: 500;
  color: rgba(20, 24, 29, 0.72);
  transition: border-color 0.22s ease, color 0.22s ease, transform 0.22s ease;
}
.cta-ghost:hover {
  border-color: rgba(20, 24, 29, 0.3);
  color: #14181d;
  transform: translateY(-1px);
}
.cta-invert {
  display: inline-flex;
  align-items: center;
  gap: 7px;
  border-radius: 12px;
  background: #fff;
  padding: 13px 22px;
  font-size: 14px;
  font-weight: 600;
  color: #14181d;
  transition: transform 0.22s cubic-bezier(0.22, 1, 0.36, 1), background 0.22s ease;
}
.cta-invert:hover {
  background: #ffd166;
  transform: translateY(-2px);
}

/* ---------- 强调字 ---------- */
.alive-word {
  position: relative;
  color: #ff6b35;
  white-space: nowrap;
}
.alive-word::after {
  content: '';
  position: absolute;
  left: -2%;
  right: -2%;
  bottom: 0.08em;
  height: 0.24em;
  border-radius: 999px;
  background: #ffd166;
  opacity: 0.55;
  z-index: -1;
}

/* ---------- 呼吸标签卡 ---------- */
.tap-stage {
  position: relative;
  display: flex;
  height: 340px;
  align-items: center;
  justify-content: center;
}
.tap-ring {
  position: absolute;
  border-radius: 999px;
  border: 1px solid rgba(255, 107, 53, 0.35);
  animation: tap-pulse 3.4s ease-out infinite;
}
.tap-ring-1 { height: 190px; width: 190px; animation-delay: 0s; }
.tap-ring-2 { height: 250px; width: 250px; animation-delay: 1.1s; }
.tap-ring-3 { height: 310px; width: 310px; animation-delay: 2.2s; }

@keyframes tap-pulse {
  0% { transform: scale(0.72); opacity: 0; }
  35% { opacity: 1; }
  100% { transform: scale(1.06); opacity: 0; }
}

.tap-card {
  position: relative;
  z-index: 2;
  width: 220px;
  border-radius: 18px;
  border: 1px solid rgba(20, 24, 29, 0.09);
  background: #fff;
  padding: 18px;
  box-shadow:
    0 2px 4px rgba(20, 24, 29, 0.05),
    0 26px 50px -26px rgba(20, 24, 29, 0.42);
  animation: card-float 6s ease-in-out infinite;
}

@keyframes card-float {
  0%, 100% { transform: translateY(0) rotate(-1.5deg); }
  50% { transform: translateY(-8px) rotate(0.5deg); }
}

.tap-hint {
  position: absolute;
  bottom: 6px;
  left: 50%;
  z-index: 3;
  display: inline-flex;
  transform: translateX(-50%);
  align-items: center;
  gap: 5px;
  border-radius: 999px;
  background: #14181d;
  padding: 6px 12px;
  font-size: 11px;
  font-weight: 500;
  color: #fff;
}

/* ---------- 跑马灯 ---------- */
.ticker-mask {
  mask-image: linear-gradient(to right, transparent, black 8%, black 92%, transparent);
  -webkit-mask-image: linear-gradient(to right, transparent, black 8%, black 92%, transparent);
}
.ticker-track {
  display: flex;
  width: max-content;
  animation: ticker-scroll 38s linear infinite;
}
.ticker-item {
  display: inline-flex;
  align-items: center;
  gap: 14px;
  padding-right: 14px;
  font-size: 12px;
  letter-spacing: 0.1em;
  color: rgba(255, 255, 255, 0.55);
  white-space: nowrap;
}
.ticker-dot {
  height: 3px;
  width: 3px;
  border-radius: 999px;
  background: #ff6b35;
}

@keyframes ticker-scroll {
  from { transform: translateX(0); }
  to { transform: translateX(-33.333%); }
}

/* ---------- 类型卡（贴纸感） ---------- */
.tcard {
  position: relative;
  border-radius: 18px;
  border: 1px solid rgba(20, 24, 29, 0.09);
  background: #fff;
  padding: 18px;
  transition: transform 0.28s cubic-bezier(0.22, 1, 0.36, 1), box-shadow 0.28s ease, border-color 0.28s ease;
}
.tcard:hover {
  transform: translateY(-3px);
  border-color: rgba(20, 24, 29, 0.18);
  box-shadow: 0 18px 34px -20px rgba(20, 24, 29, 0.35);
}
.ticon {
  display: flex;
  height: 36px;
  width: 36px;
  align-items: center;
  justify-content: center;
  border-radius: 11px;
}
.tcount {
  font-size: 10.5px;
  font-weight: 500;
  letter-spacing: 0.04em;
  color: rgba(20, 24, 29, 0.35);
}
.tflip {
  display: grid;
  grid-template-columns: repeat(5, minmax(0, 1fr));
  gap: 6px;
  height: 78px;
}
.tcol {
  overflow: hidden;
  border-radius: 8px;
  background: #ece9e3;
  box-shadow: 0 1px 2px rgba(20, 24, 29, 0.05);
}
.tcol-track {
  display: flex;
  flex-direction: column;
  transform: translate3d(0, calc(var(--flip-index, 0) * -78px), 0);
  transition: transform 520ms cubic-bezier(0.22, 1, 0.36, 1);
  will-change: transform;
}
.tcol-snap {
  transition: none;
}
/* 横屏缩略图：居中裁切（覆盖 .tthumb img 的 top center，否则只看到画面顶部黑边） */
.tthumb img.timg--wide {
  object-fit: cover;
  object-position: center;
}
.tthumb {
  display: block;
  height: 78px;
  flex: 0 0 78px;
  overflow: hidden;
  background: #ece9e3;
}
.tthumb img {
  display: block;
  width: 100%;
  height: 100%;
  object-fit: cover;
  object-position: top center;
  user-select: none;
}
.tarrow {
  color: rgba(20, 24, 29, 0.18);
  transition: color 0.24s ease, transform 0.24s ease;
}
.tcard:hover .tarrow {
  color: #ff6b35;
  transform: translate(2px, -2px);
}

/* ---------- 步骤 ---------- */
.step-card {
  position: relative;
  border-radius: 18px;
  border: 1px solid rgba(20, 24, 29, 0.09);
  background: #fbfaf7;
  padding: 22px;
  transition: transform 0.28s cubic-bezier(0.22, 1, 0.36, 1), border-color 0.28s ease;
}
.step-card:hover {
  transform: translateY(-3px);
  border-color: rgba(255, 107, 53, 0.4);
}
.step-num {
  font-family: ui-monospace, SFMono-Regular, Menlo, monospace;
  font-size: 11px;
  letter-spacing: 0.12em;
  color: rgba(20, 24, 29, 0.3);
}

@media (prefers-reduced-motion: reduce) {
  .ticker-track,
  .tap-ring,
  .tap-card {
    animation: none;
  }
  .tcol-track {
    transform: none !important;
    transition: none;
  }
}
</style>
