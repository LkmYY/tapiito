<template>
  <div class="bc-ink-root min-h-screen px-6 py-9">
    <div class="bc-ink-mist bc-ink-mist-a"></div>
    <div class="bc-ink-mist bc-ink-mist-b"></div>

    <div class="bc-ink-scroll mx-auto w-full max-w-md">
      <header class="bc-ink-head">
        <span class="bc-ink-head-title">墨 帖</span>
        <span class="bc-ink-head-meta">{{ nfcId ? 'NO.' + nfcId : 'INK LEAF' }}</span>
      </header>

      <section class="bc-ink-hero">
        <svg class="bc-ink-wash" viewBox="0 0 200 200" aria-hidden="true">
          <defs>
            <radialGradient id="bcInkGrad" cx="42%" cy="38%" r="60%">
              <stop offset="0%" stop-color="#1c1917" stop-opacity="0.16" />
              <stop offset="58%" stop-color="#44403c" stop-opacity="0.1" />
              <stop offset="100%" stop-color="#78716c" stop-opacity="0" />
            </radialGradient>
          </defs>
          <path
            d="M99 12c27 4 48 20 57 42 8 21 2 44-12 60-15 17-37 26-56 22-19-3-36-17-43-35-8-19-4-41 8-56 12-16 28-28 46-33Z"
            fill="url(#bcInkGrad)"
          />
          <path d="M42 138c18 10 43 16 66 10 18-5 32-15 41-28" fill="none" stroke="#292524" stroke-opacity="0.14" stroke-width="6" stroke-linecap="round" />
        </svg>

        <p class="bc-ink-eyebrow">一纸相识 · 见字如面</p>
        <h1 class="bc-ink-name">{{ payload.name || '未命名' }}</h1>
        <p class="bc-ink-title">{{ payload.title }}</p>
        <p class="bc-ink-company">{{ payload.company }}</p>

        <div class="bc-ink-seal">
          <span>{{ initial }}</span>
        </div>
      </section>

      <div class="bc-ink-rule"><span></span></div>

      <section class="bc-ink-info">
        <div class="bc-ink-item">
          <span class="bc-ink-key">电话</span>
          <a v-if="payload.phone" :href="'tel:' + payload.phone" class="bc-ink-val">{{ payload.phone }}</a>
        </div>
        <div class="bc-ink-item">
          <span class="bc-ink-key">邮箱</span>
          <a v-if="payload.email" :href="'mailto:' + payload.email" class="bc-ink-val">{{ payload.email }}</a>
        </div>
        <div class="bc-ink-item">
          <span class="bc-ink-key">网站</span>
          <a
            v-if="payload.website"
            :href="payload.website"
            target="_blank"
            rel="noopener noreferrer"
            class="bc-ink-val"
          >
            {{ payload.website }}
          </a>
        </div>
        <div class="bc-ink-item">
          <span class="bc-ink-key">微信</span>
          <span class="bc-ink-val" :class="{ 'bc-ink-empty': !payload.wechat }">{{ payload.wechat }}</span>
        </div>
      </section>

      <footer class="bc-ink-foot">
        <span>落款于 · 云图</span>
        <span class="bc-ink-stamp">{{ initial }}</span>
        <span>壬寅年秋</span>
      </footer>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps(['payload', 'nfcId'])

const initial = computed(() => (props.payload?.name || '?').toString().trim()[0] || '?')
</script>

<style scoped>
.bc-ink-root {
  position: relative;
  overflow: hidden;
  background: #f4f0e5;
  font-family: 'Kaiti SC', 'STKaiti', 'KaiTi', 'Songti SC', serif;
}
.bc-ink-root::before {
  content: '';
  position: absolute;
  inset: 0;
  background-image: radial-gradient(rgba(68, 64, 60, 0.2) 0.5px, transparent 0.5px);
  background-size: 4px 4px;
  opacity: 0.22;
  pointer-events: none;
}
.bc-ink-mist {
  position: absolute;
  border-radius: 50%;
  filter: blur(40px);
  pointer-events: none;
}
.bc-ink-mist-a {
  top: -60px;
  right: -80px;
  width: 260px;
  height: 220px;
  background: rgba(120, 113, 108, 0.18);
}
.bc-ink-mist-b {
  bottom: -80px;
  left: -90px;
  width: 280px;
  height: 240px;
  background: rgba(168, 162, 158, 0.2);
}
.bc-ink-scroll {
  position: relative;
  z-index: 1;
  animation: bc-ink-unroll 1s cubic-bezier(0.22, 1, 0.36, 1) both;
}
.bc-ink-head {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding-bottom: 16px;
  border-bottom: 1px solid rgba(41, 37, 36, 0.24);
}
.bc-ink-head-title {
  font-size: 17px;
  letter-spacing: 0.4em;
  color: #292524;
}
.bc-ink-head-meta {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 7px;
  letter-spacing: 0.18em;
  color: rgba(68, 64, 60, 0.5);
}
.bc-ink-hero {
  position: relative;
  padding: 58px 0 44px;
  animation: bc-ink-fade 1s 0.15s both;
}
.bc-ink-wash {
  position: absolute;
  top: -12px;
  left: -34px;
  width: 290px;
  height: 290px;
  opacity: 0.9;
  animation: bc-ink-breathe 9s ease-in-out infinite;
}
.bc-ink-eyebrow {
  position: relative;
  font-size: 12px;
  letter-spacing: 0.3em;
  color: rgba(68, 64, 60, 0.55);
}
.bc-ink-name {
  position: relative;
  margin-top: 20px;
  font-family: 'Kaiti SC', 'STKaiti', 'KaiTi', 'Songti SC', serif;
  font-size: clamp(54px, 17vw, 76px);
  font-weight: 500;
  line-height: 1;
  letter-spacing: 0.14em;
  color: #1c1917;
  overflow-wrap: anywhere;
}
.bc-ink-title {
  position: relative;
  margin-top: 24px;
  font-size: 13px;
  letter-spacing: 0.24em;
  color: #44403c;
}
.bc-ink-company {
  position: relative;
  margin-top: 5px;
  font-size: 11px;
  letter-spacing: 0.18em;
  color: rgba(68, 64, 60, 0.55);
}
.bc-ink-seal {
  position: absolute;
  right: 4px;
  bottom: 36px;
  display: flex;
  width: 54px;
  height: 54px;
  align-items: center;
  justify-content: center;
  border: 2px solid #b91c1c;
  background: rgba(185, 28, 28, 0.06);
  transform: rotate(4deg);
  transition: transform 0.4s ease, background 0.4s ease;
}
.bc-ink-seal::before {
  content: '';
  position: absolute;
  inset: 3px;
  border: 1px solid rgba(185, 28, 28, 0.55);
}
.bc-ink-seal span {
  font-size: 25px;
  color: #b91c1c;
  line-height: 1;
}
.bc-ink-seal:hover {
  transform: rotate(-4deg) scale(1.06);
  background: rgba(185, 28, 28, 0.12);
}
.bc-ink-rule {
  position: relative;
  height: 1px;
  background: rgba(41, 37, 36, 0.25);
}
.bc-ink-rule span {
  position: absolute;
  top: -2px;
  left: 0;
  width: 82px;
  height: 5px;
  background: #292524;
  opacity: 0.85;
}
.bc-ink-info {
  padding: 12px 0 8px;
  animation: bc-ink-fade 1s 0.3s both;
}
.bc-ink-item {
  display: flex;
  align-items: baseline;
  gap: 20px;
  padding: 16px 2px;
  border-bottom: 1px solid rgba(41, 37, 36, 0.14);
}
.bc-ink-key {
  flex: none;
  width: 32px;
  font-size: 10px;
  letter-spacing: 0.2em;
  color: rgba(68, 64, 60, 0.5);
}
.bc-ink-val {
  min-width: 0;
  flex: 1;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  font-size: 14px;
  letter-spacing: 0.04em;
  color: #1c1917;
  text-decoration: none;
  transition: color 0.3s ease, letter-spacing 0.3s ease;
}
a.bc-ink-val:hover {
  color: #b91c1c;
  letter-spacing: 0.08em;
}
a.bc-ink-val:active {
  opacity: 0.6;
}
.bc-ink-empty {
  color: rgba(168, 162, 158, 0.75);
}
.bc-ink-foot {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 10px;
  padding-top: 38px;
  font-size: 9px;
  letter-spacing: 0.16em;
  color: rgba(68, 64, 60, 0.5);
}
.bc-ink-stamp {
  display: flex;
  width: 26px;
  height: 26px;
  align-items: center;
  justify-content: center;
  border: 1px solid #b91c1c;
  font-size: 12px;
  color: #b91c1c;
  transform: rotate(-5deg);
}
@keyframes bc-ink-unroll {
  from { opacity: 0; transform: translateY(20px); clip-path: inset(0 0 12% 0); }
  to { opacity: 1; transform: none; clip-path: inset(0 0 0 0); }
}
@keyframes bc-ink-fade {
  from { opacity: 0; }
  to { opacity: 1; }
}
@keyframes bc-ink-breathe {
  0%, 100% { transform: scale(1) rotate(0deg); opacity: 0.85; }
  50% { transform: scale(1.04) rotate(2deg); opacity: 1; }
}
@media (prefers-reduced-motion: reduce) {
  .bc-ink-root *,
  .bc-ink-root *::before,
  .bc-ink-root *::after {
    animation: none !important;
    transition: none !important;
  }
}
</style>
