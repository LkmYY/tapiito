<template>
  <div class="bc-bot-root min-h-screen px-5 py-9">
    <div class="bc-bot-wrap mx-auto flex min-h-[calc(100vh-4.5rem)] w-full max-w-md flex-col">
      <header class="bc-bot-head">
        <span>HERBARIUM / 07</span>
        <span>{{ nfcId ? 'SPECIMEN ' + nfcId : 'BOTANICAL PRESS' }}</span>
      </header>

      <div class="bc-bot-card">
        <svg class="bc-bot-leaf bc-bot-leaf-tl" viewBox="0 0 100 140" aria-hidden="true">
          <path d="M50 138C50 90 42 44 12 8c38 8 64 44 64 84 0 20-10 36-26 46Z" fill="none" stroke="currentColor" stroke-width="1.2" />
          <path d="M48 120c8-30 10-58 4-86" fill="none" stroke="currentColor" stroke-width="0.8" />
          <path d="M50 92c-12-4-20-12-24-24M52 74c10-4 17-11 21-22M46 56c-8-3-14-9-17-18" fill="none" stroke="currentColor" stroke-width="0.7" />
        </svg>
        <svg class="bc-bot-leaf bc-bot-leaf-br" viewBox="0 0 100 140" aria-hidden="true">
          <path d="M50 138C50 90 42 44 12 8c38 8 64 44 64 84 0 20-10 36-26 46Z" fill="none" stroke="currentColor" stroke-width="1.2" />
          <path d="M48 120c8-30 10-58 4-86" fill="none" stroke="currentColor" stroke-width="0.8" />
        </svg>

        <div class="bc-bot-oval">
          <p class="bc-bot-kicker">EST. IN GOOD COMPANY</p>
          <div class="bc-bot-initial">{{ initial }}</div>
          <h1 class="bc-bot-name">{{ payload.name || '未命名' }}</h1>
          <p class="bc-bot-role">{{ payload.title }}</p>
          <p class="bc-bot-company">{{ payload.company }}</p>
          <div class="bc-bot-sprig">❦</div>
        </div>

        <div class="bc-bot-divider">
          <span></span>
          <i>✿</i>
          <span></span>
        </div>

        <section class="bc-bot-info">
          <div class="bc-bot-col">
            <p class="bc-bot-label">TELEPHONE</p>
            <a v-if="payload.phone" :href="'tel:' + payload.phone">{{ payload.phone }}</a>
            <p class="bc-bot-label">EMAIL</p>
            <a v-if="payload.email" :href="'mailto:' + payload.email">{{ payload.email }}</a>
          </div>
          <div class="bc-bot-col">
            <p class="bc-bot-label">WEBSITE</p>
            <a v-if="payload.website" :href="payload.website" target="_blank" rel="noopener noreferrer">{{ payload.website }}</a>
            <p class="bc-bot-label">WECHAT</p>
            <span :class="{ 'bc-bot-empty': !payload.wechat }">{{ payload.wechat }}</span>
          </div>
        </section>

        <p class="bc-bot-quote">“May our paths cross, and the work grow.”</p>
      </div>

      <footer class="bc-bot-foot">
        <span>PRESSED &amp; PRESERVED</span>
        <span>FIELD NOTES · 2026</span>
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
.bc-bot-root {
  position: relative;
  overflow: hidden;
  background: #e8eddf;
  font-family: Georgia, 'Songti SC', 'PingFang SC', serif;
  color: #2f3e2e;
}
.bc-bot-root::before {
  content: '';
  position: absolute;
  inset: 0;
  background-image: radial-gradient(rgba(47, 62, 46, 0.18) 0.5px, transparent 0.5px);
  background-size: 5px 5px;
  opacity: 0.28;
}
.bc-bot-wrap {
  position: relative;
  z-index: 1;
}
.bc-bot-head {
  display: flex;
  justify-content: space-between;
  gap: 12px;
  padding-bottom: 14px;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6.5px;
  letter-spacing: 0.2em;
  color: rgba(47, 62, 46, 0.55);
  animation: bc-bot-fade 0.9s both;
}
.bc-bot-card {
  position: relative;
  flex: 1;
  border: 1px solid rgba(47, 62, 46, 0.28);
  background: #f7f4ea;
  padding: 26px 22px 22px;
  box-shadow: 0 26px 60px -38px rgba(47, 62, 46, 0.6);
  animation: bc-bot-in 1s cubic-bezier(0.22, 1, 0.36, 1) both;
}
.bc-bot-card::before {
  content: '';
  position: absolute;
  inset: 7px;
  border: 1px solid rgba(47, 62, 46, 0.12);
  pointer-events: none;
}
.bc-bot-leaf {
  position: absolute;
  color: rgba(47, 62, 46, 0.34);
  pointer-events: none;
}
.bc-bot-leaf-tl {
  top: 14px;
  left: 10px;
  width: 78px;
  height: 108px;
  transform: rotate(-18deg);
  animation: bc-bot-sway 8s ease-in-out infinite;
}
.bc-bot-leaf-br {
  right: 10px;
  bottom: 12px;
  width: 66px;
  height: 92px;
  transform: rotate(158deg);
  animation: bc-bot-sway 9s ease-in-out infinite reverse;
}
.bc-bot-oval {
  position: relative;
  margin: 14px auto 0;
  padding: 38px 22px 32px;
  border: 1px solid rgba(47, 62, 46, 0.36);
  border-radius: 50% / 30%;
  text-align: center;
}
.bc-bot-oval::before {
  content: '';
  position: absolute;
  inset: 5px;
  border: 1px dashed rgba(47, 62, 46, 0.2);
  border-radius: inherit;
}
.bc-bot-kicker {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6px;
  letter-spacing: 0.26em;
  color: rgba(47, 62, 46, 0.55);
}
.bc-bot-initial {
  display: flex;
  width: 54px;
  height: 54px;
  margin: 14px auto 0;
  align-items: center;
  justify-content: center;
  border: 1px solid rgba(47, 62, 46, 0.45);
  border-radius: 50%;
  font-size: 22px;
  color: #2f3e2e;
  transition: background 0.35s ease, color 0.35s ease, transform 0.35s ease;
}
.bc-bot-initial:hover {
  background: #2f3e2e;
  color: #f7f4ea;
  transform: rotate(-8deg);
}
.bc-bot-name {
  margin-top: 16px;
  font-family: Georgia, 'Songti SC', serif;
  font-size: clamp(34px, 11vw, 44px);
  font-weight: 400;
  line-height: 1.05;
  letter-spacing: -0.02em;
  overflow-wrap: anywhere;
}
.bc-bot-role {
  margin-top: 10px;
  font-size: 11px;
  font-style: italic;
  letter-spacing: 0.1em;
  color: #55664f;
}
.bc-bot-company {
  margin-top: 3px;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6.5px;
  letter-spacing: 0.18em;
  color: rgba(47, 62, 46, 0.5);
}
.bc-bot-sprig {
  margin-top: 12px;
  font-size: 12px;
  color: rgba(47, 62, 46, 0.55);
}
.bc-bot-divider {
  display: flex;
  align-items: center;
  gap: 10px;
  margin: 20px 0 4px;
  color: rgba(47, 62, 46, 0.6);
  font-size: 8px;
}
.bc-bot-divider span {
  flex: 1;
  height: 1px;
  background: rgba(47, 62, 46, 0.24);
}
.bc-bot-info {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 0;
  padding: 8px 0 6px;
}
.bc-bot-col {
  min-width: 0;
  padding: 0 12px;
}
.bc-bot-col:first-child {
  padding-left: 4px;
  border-right: 1px solid rgba(47, 62, 46, 0.18);
}
.bc-bot-col:last-child {
  padding-right: 4px;
}
.bc-bot-label {
  margin-top: 12px;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 5.5px;
  letter-spacing: 0.2em;
  color: rgba(47, 62, 46, 0.45);
}
.bc-bot-col .bc-bot-label:first-child {
  margin-top: 0;
}
.bc-bot-col a,
.bc-bot-col span {
  display: block;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  font-size: 10.5px;
  letter-spacing: 0.01em;
  color: #2f3e2e;
  text-decoration: none;
  transition: color 0.25s ease;
}
.bc-bot-col a:hover {
  color: #7c2d12;
  text-decoration: underline;
  text-underline-offset: 3px;
}
.bc-bot-col a:active {
  opacity: 0.6;
}
.bc-bot-empty {
  color: rgba(47, 62, 46, 0.28);
}
.bc-bot-quote {
  margin-top: 14px;
  padding-top: 12px;
  border-top: 1px solid rgba(47, 62, 46, 0.18);
  text-align: center;
  font-size: 9px;
  font-style: italic;
  letter-spacing: 0.03em;
  color: rgba(47, 62, 46, 0.6);
}
.bc-bot-foot {
  display: flex;
  justify-content: space-between;
  padding-top: 16px;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6px;
  letter-spacing: 0.18em;
  color: rgba(47, 62, 46, 0.42);
}
@keyframes bc-bot-in {
  from { opacity: 0; transform: translateY(22px) scale(0.99); }
  to { opacity: 1; transform: none; }
}
@keyframes bc-bot-fade {
  from { opacity: 0; }
  to { opacity: 1; }
}
@keyframes bc-bot-sway {
  0%, 100% { transform: rotate(-18deg) translateY(0); }
  50% { transform: rotate(-13deg) translateY(-4px); }
}
@media (prefers-reduced-motion: reduce) {
  .bc-bot-root *,
  .bc-bot-root *::before,
  .bc-bot-root *::after {
    animation: none !important;
    transition: none !important;
  }
}
</style>
