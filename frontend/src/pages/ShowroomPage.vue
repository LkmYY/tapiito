<template>
  <div class="min-h-screen bg-[#fbfaf7] font-sans text-[#14181d]">
    <!-- 顶部 -->
    <SiteHeader back-to="/" back-label="返回首页" cta-label="预约咨询" cta-to="/consult" />

    <!-- 标题区 -->
    <section class="mx-auto max-w-[1400px] px-4 pb-8 pt-14 sm:px-6">
      <p class="eyebrow">Showroom</p>
      <h1 class="mt-3 text-[34px] font-semibold leading-tight tracking-[-0.03em] md:text-[44px]">
        先随便看看，<br class="md:hidden" />再做决定
      </h1>
      <p class="mt-4 max-w-[520px] text-[15px] leading-relaxed text-[#14181d]/55">
        10 种类型 × 30 种风格，全部是真实渲染的页面。点开任意一张，手机上看到的就是这个样子 ——
        购买即完成首次写入；含 1 次免费内容调整，链接永远不变。
      </p>

      <div class="mt-8 flex flex-wrap items-center gap-2">
        <button class="chip" :class="{ 'chip-on': activeType === '' }" @click="activeType = ''">
          全部
          <span class="chip-num">{{ total }}</span>
        </button>
        <button
          v-for="key in visibleTypeOrder"
          :key="key"
          class="chip"
          :class="{ 'chip-on': activeType === key }"
          @click="activeType = key"
        >
          <component :is="templateTypes[key].icon" class="h-3.5 w-3.5" />
          {{ templateTypes[key].label }}
          <span class="chip-num">{{ templateTypes[key].styles.length }}</span>
        </button>
      </div>
    </section>

    <!-- 卡片墙 -->
    <section class="mx-auto max-w-[1400px] px-4 pb-24 sm:px-6">
      <div class="grid grid-cols-2 gap-3 sm:gap-6 md:grid-cols-3 xl:grid-cols-4">
        <a
          v-for="(card, i) in visibleCards"
          :key="activeType + '__' + card.type + '__' + card.style"
          :href="previewUrl(card)"
          target="_blank"
          rel="noopener"
          class="g-card reveal"
          :style="{ animationDelay: Math.min(i, 11) * 45 + 'ms' }"
        >
          <div class="g-window" :style="{ background: card.gradient || '#f1efea' }">
            <div class="g-fallback" aria-hidden="true">
              <span class="g-fallback-icon">
                <component :is="templateTypes[card.type].icon" class="h-5 w-5" />
              </span>
              <span class="g-fallback-label">{{ card.label }}</span>
              <span class="g-fallback-sub">{{ templateTypes[card.type].label }} · 点击预览</span>
            </div>
            <img
              v-if="!failed.has(card.key)"
              :src="imgUrl(card)"
              :alt="templateTypes[card.type].label + ' ' + card.label"
              loading="lazy"
              class="g-img"
              :class="{ 'g-img--wide': isWide(card) }"
              @error="markFailed(card)"
            />
            <div class="g-veil"></div>
            <span class="g-open">
              打开预览
              <ArrowUpRight class="h-3.5 w-3.5" />
            </span>
          </div>
          <div class="g-meta">
            <component :is="templateTypes[card.type].icon" class="g-meta-icon h-4 w-4 shrink-0 text-[#ff6b35]" />
            <span class="g-meta-type truncate font-medium">{{ templateTypes[card.type].label }}</span>
            <span class="g-meta-dot shrink-0 text-[#14181d]/30">·</span>
            <span class="g-meta-style truncate text-[#14181d]/55">{{ card.label }}</span>
          </div>
        </a>
      </div>
    </section>
  </div>
</template>

<script setup>
import { computed, ref } from 'vue'
import { useRoute } from 'vue-router'
import { ArrowLeft, ArrowUpRight, Info, MessageSquare, Nfc, ShoppingBag } from 'lucide-vue-next'
import { templateOrder, visibleTypeOrder, templateTypes } from '../templates/registry'
import SiteHeader from '../components/SiteHeader.vue'

const route = useRoute()
const activeType = ref(typeof route.query.type === 'string' && templateTypes[route.query.type] ? route.query.type : '')

const allCards = computed(() => {
  const out = []
  for (const key of visibleTypeOrder) {
    const t = templateTypes[key]
    for (const s of t.styles) {
      out.push({
        type: key,
        style: s.value || 'default',
        key: key + '__' + (s.value || 'default'),
        label: s.label || '默认',
        gradient: s.gradient || '',
      })
    }
  }
  return out
})

/** 缩略图缺失时回退到渐变海报，避免展厅出现破图 */
const failed = ref(new Set())
const markFailed = (card) => {
  failed.value = new Set(failed.value).add(card.key)
}

const total = computed(() => allCards.value.length)

const visibleCards = computed(() =>
  activeType.value ? allCards.value.filter((c) => c.type === activeType.value) : allCards.value
)

const imgUrl = (c) => `/showroom/${c.type}__${c.style}.webp`
/** 横屏类型需要特殊图片适配 */
const isWide = (c) => templateTypes[c.type]?.layout === 'landscape'

const previewUrl = (c) =>
  c.style === 'default' ? `/preview?type=${c.type}` : `/preview?type=${c.type}&style=${c.style}`
</script>

<style scoped>
.brand-chip {
  display: grid;
  place-items: center;
  width: 32px;
  height: 32px;
  border-radius: 10px;
  background: #14181d;
  transition: transform 0.3s cubic-bezier(0.22, 1, 0.36, 1);
}
.brand-chip:hover {
  transform: rotate(-6deg);
}

.nav-back {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  height: 34px;
  padding: 0 14px;
  border-radius: 999px;
  font-size: 13px;
  color: rgba(20, 24, 29, 0.65);
  border: 1px solid rgba(20, 24, 29, 0.1);
  transition: all 0.25s ease;
}
.nav-back:hover {
  color: #ff6b35;
  border-color: rgba(255, 107, 53, 0.4);
  transform: translateX(-2px);
}

.nav-about {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  height: 34px;
  padding: 0 12px;
  border-radius: 999px;
  font-size: 13px;
  color: rgba(20, 24, 29, 0.55);
  transition: all 0.25s ease;
}
.nav-about:hover {
  background: rgba(20, 24, 29, 0.05);
  color: #14181d;
}

@media (max-width: 380px) {
  .showroom-about {
    display: none;
  }
}

.nav-cta {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  height: 34px;
  padding: 0 14px;
  border-radius: 999px;
  background: #ff6b35;
  color: #fff;
  font-size: 13px;
  font-weight: 600;
  letter-spacing: 0.01em;
  box-shadow: 0 8px 20px -12px rgba(255, 107, 53, 0.9);
  transition: all 0.25s ease;
}
.nav-cta:hover {
  background: #f25a22;
  transform: translateY(-1px);
  box-shadow: 0 12px 26px -12px rgba(255, 107, 53, 0.95);
}

.eyebrow {
  font-size: 11px;
  font-weight: 600;
  letter-spacing: 0.2em;
  text-transform: uppercase;
  color: #ff6b35;
}

.chip {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  height: 36px;
  padding: 0 14px;
  border-radius: 999px;
  font-size: 13px;
  color: rgba(20, 24, 29, 0.62);
  background: #fff;
  border: 1px solid rgba(20, 24, 29, 0.08);
  transition: all 0.25s ease;
  cursor: pointer;
}
.chip:hover {
  border-color: rgba(255, 107, 53, 0.45);
  color: #14181d;
}
.chip-on {
  background: #14181d;
  border-color: #14181d;
  color: #fff;
}
.chip-on:hover {
  color: #fff;
}
.chip-num {
  font-size: 11px;
  opacity: 0.55;
  font-variant-numeric: tabular-nums;
}

.g-card {
  display: block;
  background: #fff;
  border: 1px solid rgba(20, 24, 29, 0.07);
  border-radius: 18px;
  overflow: hidden;
  transition: transform 0.35s cubic-bezier(0.22, 1, 0.36, 1), box-shadow 0.35s ease, border-color 0.35s ease;
}
.g-card:hover {
  transform: translateY(-5px);
  border-color: rgba(255, 107, 53, 0.35);
  box-shadow: 0 20px 44px -20px rgba(20, 24, 29, 0.3);
}

.g-window {
  position: relative;
  height: 300px;
  overflow: hidden;
  background: #f1efea;
}
.g-fallback {
  position: absolute;
  inset: 0;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  gap: 9px;
  padding: 22px;
  text-align: center;
}
.g-fallback-icon {
  display: grid;
  place-items: center;
  width: 46px;
  height: 46px;
  border-radius: 15px;
  color: #14181d;
  background: rgba(255, 255, 255, 0.92);
  box-shadow: 0 10px 24px -14px rgba(0, 0, 0, 0.45);
}
.g-fallback-label {
  padding: 6px 13px;
  border-radius: 999px;
  background: #fff;
  color: #14181d;
  font-size: 13px;
  font-weight: 600;
  letter-spacing: 0.01em;
  box-shadow: 0 10px 24px -16px rgba(0, 0, 0, 0.4);
}
.g-fallback-sub {
  color: rgba(255, 255, 255, 0.92);
  font-size: 11px;
  letter-spacing: 0.08em;
  text-shadow: 0 1px 8px rgba(0, 0, 0, 0.35);
}
.g-img {
  display: block;
  width: 100%;
  transition: transform 4.5s cubic-bezier(0.33, 0, 0.2, 1);
  will-change: transform;
}
/* 横屏类型（奇怪的想法）：整幅 contain，不裁切也不做长图上移 */
.g-img--wide {
  height: 300px;
  object-fit: contain;
  background: #05060a;
}
.g-card:hover .g-img--wide {
  transform: none;
}
.g-card:hover .g-img {
  transform: translateY(calc(-100% + 300px));
}

.g-veil {
  position: absolute;
  inset: 0;
  background: linear-gradient(to bottom, transparent 55%, rgba(20, 24, 29, 0.42));
  opacity: 0;
  transition: opacity 0.3s ease;
  pointer-events: none;
}
.g-card:hover .g-veil {
  opacity: 1;
}

.g-open {
  position: absolute;
  left: 14px;
  bottom: 12px;
  display: inline-flex;
  align-items: center;
  gap: 6px;
  height: 32px;
  padding: 0 13px;
  border-radius: 999px;
  background: #ff6b35;
  color: #fff;
  font-size: 12.5px;
  font-weight: 500;
  opacity: 0;
  transform: translateY(6px);
  transition: opacity 0.3s ease, transform 0.3s cubic-bezier(0.22, 1, 0.36, 1);
  pointer-events: none;
}
.g-card:hover .g-open {
  opacity: 1;
  transform: translateY(0);
}

.g-meta {
  display: flex;
  align-items: center;
  gap: 8px;
  min-width: 0;
  padding: 12px 14px 13px;
  font-size: 13px;
}
.g-meta-type {
  flex: 0 1 auto;
}
.g-meta-style {
  flex: 1 1 auto;
}

@media (max-width: 639px) {
  .g-window {
    height: 250px;
  }
  .g-meta {
    gap: 5px;
    padding: 10px 10px 11px;
    font-size: 12px;
  }
  .g-meta-icon {
    width: 14px;
    height: 14px;
  }
  .g-meta-dot {
    display: none;
  }
  .g-fallback {
    gap: 7px;
    padding: 14px;
  }
  .g-fallback-icon {
    width: 38px;
    height: 38px;
    border-radius: 12px;
  }
  .g-fallback-label {
    padding: 5px 10px;
    font-size: 12px;
  }
  .g-fallback-sub {
    font-size: 10px;
  }
  .g-open {
    left: 10px;
    bottom: 9px;
    height: 28px;
    padding: 0 10px;
    font-size: 11px;
  }
}

@media (hover: none) {
  .g-open {
    opacity: 1;
    transform: translateY(0);
  }
  .g-veil {
    opacity: 1;
  }
}

@keyframes rise {
  from {
    opacity: 0;
    transform: translateY(16px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}
.reveal {
  animation: rise 0.6s cubic-bezier(0.22, 1, 0.36, 1) both;
}

@media (prefers-reduced-motion: reduce) {
  .g-card,
  .g-img,
  .reveal,
  .g-open,
  .g-veil,
  .chip,
  .nav-back {
    transition: none !important;
    animation: none !important;
  }
}
</style>
