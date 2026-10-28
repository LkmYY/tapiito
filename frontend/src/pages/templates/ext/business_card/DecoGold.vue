<template>
  <div class="bc-deco-root min-h-screen px-5 py-9">
    <div class="bc-deco-fan bc-deco-fan-top"></div>
    <div class="bc-deco-fan bc-deco-fan-bottom"></div>

    <div class="bc-deco-wrap mx-auto flex min-h-[calc(100vh-4.5rem)] w-full max-w-md flex-col">
      <header class="bc-deco-head">
        <span>MCMXXVI</span>
        <span class="bc-deco-head-mark">✦</span>
        <span>PRIVATE SALON</span>
      </header>

      <div class="bc-deco-frame">
        <span class="bc-deco-corner bc-deco-corner-tl"></span>
        <span class="bc-deco-corner bc-deco-corner-tr"></span>
        <span class="bc-deco-corner bc-deco-corner-bl"></span>
        <span class="bc-deco-corner bc-deco-corner-br"></span>

        <div class="bc-deco-ornament"><i></i><span>◆</span><i></i></div>

        <section class="bc-deco-hero">
          <p class="bc-deco-kicker">THE GOLDEN REGISTER</p>
          <div class="bc-deco-monogram">{{ initial }}</div>
          <h1 class="bc-deco-name">{{ payload.name || '未命名' }}</h1>
          <p class="bc-deco-role">{{ payload.title }}</p>
          <p class="bc-deco-company">{{ payload.company }}</p>
        </section>

        <div class="bc-deco-divider">
          <span></span><i>◇</i><span></span>
        </div>

        <section class="bc-deco-contacts">
          <div class="bc-deco-row">
            <span class="bc-deco-label">TEL</span>
            <a v-if="payload.phone" :href="'tel:' + payload.phone">{{ payload.phone }}</a>
          </div>
          <div class="bc-deco-row">
            <span class="bc-deco-label">MAIL</span>
            <a v-if="payload.email" :href="'mailto:' + payload.email">{{ payload.email }}</a>
          </div>
          <div class="bc-deco-row">
            <span class="bc-deco-label">WEB</span>
            <a v-if="payload.website" :href="payload.website" target="_blank" rel="noopener noreferrer">{{ payload.website }}</a>
          </div>
          <div class="bc-deco-row">
            <span class="bc-deco-label">WX</span>
            <span :class="{ 'bc-deco-empty': !payload.wechat }">{{ payload.wechat }}</span>
          </div>
        </section>

        <div class="bc-deco-ornament bc-deco-ornament-low"><i></i><span>◆</span><i></i></div>
        <p class="bc-deco-serial">Nº {{ nfcId || '000' }}</p>
      </div>

      <footer class="bc-deco-foot">
        <span>BY APPOINTMENT ONLY</span>
        <span>EST. YUNTU</span>
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
.bc-deco-root {
  position: relative;
  overflow: hidden;
  background: #0a0908;
  font-family: 'Times New Roman', Georgia, 'Songti SC', serif;
  color: #d4af37;
}
.bc-deco-root::before {
  content: '';
  position: absolute;
  inset: 0;
  background: radial-gradient(ellipse at 50% 20%, rgba(212, 175, 55, 0.12), transparent 58%),
    radial-gradient(ellipse at 50% 100%, rgba(212, 175, 55, 0.07), transparent 50%);
  pointer-events: none;
}
.bc-deco-fan {
  position: absolute;
  left: 50%;
  width: 560px;
  height: 280px;
  margin-left: -280px;
  background: repeating-conic-gradient(from 262deg at 50% 100%, rgba(212, 175, 55, 0.16) 0deg 0.5deg, transparent 0.5deg 7deg);
  -webkit-mask-image: radial-gradient(ellipse at 50% 100%, #000 0%, transparent 68%);
  mask-image: radial-gradient(ellipse at 50% 100%, #000 0%, transparent 68%);
  pointer-events: none;
}
.bc-deco-fan-top {
  top: -150px;
  transform: rotate(180deg);
}
.bc-deco-fan-bottom {
  bottom: -150px;
}
.bc-deco-wrap {
  position: relative;
  z-index: 1;
}
.bc-deco-head {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding-bottom: 16px;
  font-size: 7px;
  letter-spacing: 0.3em;
  color: rgba(212, 175, 55, 0.6);
  animation: bc-deco-fade 0.9s both;
}
.bc-deco-head-mark {
  font-size: 9px;
  color: #d4af37;
}
.bc-deco-frame {
  position: relative;
  border: 1px solid rgba(212, 175, 55, 0.55);
  padding: 30px 20px 24px;
  background: linear-gradient(180deg, rgba(212, 175, 55, 0.045), transparent 40%, rgba(212, 175, 55, 0.03));
  box-shadow: inset 0 0 0 4px #0a0908, inset 0 0 0 5px rgba(212, 175, 55, 0.25);
  animation: bc-deco-in 1s cubic-bezier(0.22, 1, 0.36, 1) both;
}
.bc-deco-corner {
  position: absolute;
  width: 22px;
  height: 22px;
  border-color: #d4af37;
  border-style: solid;
}
.bc-deco-corner-tl { top: 7px; left: 7px; border-width: 1px 0 0 1px; }
.bc-deco-corner-tr { top: 7px; right: 7px; border-width: 1px 1px 0 0; }
.bc-deco-corner-bl { bottom: 7px; left: 7px; border-width: 0 0 1px 1px; }
.bc-deco-corner-br { bottom: 7px; right: 7px; border-width: 0 1px 1px 0; }
.bc-deco-ornament {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 9px;
  color: rgba(212, 175, 55, 0.8);
  font-size: 7px;
}
.bc-deco-ornament i {
  width: 54px;
  height: 1px;
  background: linear-gradient(90deg, transparent, #d4af37);
}
.bc-deco-ornament i:last-child {
  background: linear-gradient(90deg, #d4af37, transparent);
}
.bc-deco-hero {
  padding: 26px 0 6px;
  text-align: center;
  animation: bc-deco-fade 1s 0.2s both;
}
.bc-deco-kicker {
  font-size: 7px;
  letter-spacing: 0.38em;
  color: rgba(212, 175, 55, 0.55);
}
.bc-deco-monogram {
  display: flex;
  width: 66px;
  height: 66px;
  margin: 20px auto 0;
  align-items: center;
  justify-content: center;
  border: 1px solid rgba(212, 175, 55, 0.7);
  transform: rotate(45deg);
  transition: transform 0.5s ease, background 0.5s ease;
}
.bc-deco-monogram::after {
  content: '';
  position: absolute;
  width: 54px;
  height: 54px;
  border: 1px solid rgba(212, 175, 55, 0.25);
}
.bc-deco-monogram:hover {
  transform: rotate(135deg);
  background: rgba(212, 175, 55, 0.1);
}
.bc-deco-hero {
  position: relative;
}
.bc-deco-hero .bc-deco-monogram {
  font-size: 22px;
}
.bc-deco-hero .bc-deco-monogram::before {
  content: attr(data-x);
}
.bc-deco-monogram span {
  transform: rotate(-45deg);
}
.bc-deco-name {
  margin-top: 38px;
  font-family: 'Didot', 'Bodoni MT', 'Times New Roman', Georgia, serif;
  font-size: clamp(40px, 13vw, 56px);
  font-weight: 400;
  line-height: 1;
  letter-spacing: 0.02em;
  text-transform: uppercase;
  color: #f5e6b8;
  overflow-wrap: anywhere;
  text-shadow: 0 0 30px rgba(212, 175, 55, 0.25);
}
.bc-deco-role {
  margin-top: 16px;
  font-size: 10px;
  letter-spacing: 0.32em;
  color: #d4af37;
}
.bc-deco-company {
  margin-top: 6px;
  font-size: 9px;
  font-style: italic;
  letter-spacing: 0.12em;
  color: rgba(245, 230, 184, 0.45);
}
.bc-deco-divider {
  display: flex;
  align-items: center;
  gap: 10px;
  margin: 28px 0 4px;
}
.bc-deco-divider span {
  flex: 1;
  height: 1px;
  background: rgba(212, 175, 55, 0.35);
}
.bc-deco-divider i {
  font-style: normal;
  font-size: 7px;
  color: rgba(212, 175, 55, 0.7);
}
.bc-deco-contacts {
  padding: 4px 0;
}
.bc-deco-row {
  display: flex;
  align-items: baseline;
  gap: 14px;
  padding: 11px 4px;
  border-bottom: 1px solid rgba(212, 175, 55, 0.14);
}
.bc-deco-label {
  flex: none;
  width: 32px;
  font-size: 6.5px;
  letter-spacing: 0.18em;
  color: rgba(212, 175, 55, 0.6);
}
.bc-deco-row a,
.bc-deco-row span,
.bc-deco-row em {
  min-width: 0;
  flex: 1;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  font-size: 12px;
  font-style: normal;
  letter-spacing: 0.05em;
  text-align: right;
  color: #efe0b8;
  text-decoration: none;
  transition: color 0.25s ease, letter-spacing 0.25s ease;
}
.bc-deco-row a:hover {
  color: #ffd76e;
  letter-spacing: 0.09em;
}
.bc-deco-row a:active {
  opacity: 0.6;
}
.bc-deco-empty,
.bc-deco-row em {
  color: rgba(239, 224, 184, 0.26);
}
.bc-deco-ornament-low {
  margin-top: 20px;
}
.bc-deco-serial {
  margin-top: 12px;
  text-align: center;
  font-size: 6.5px;
  letter-spacing: 0.24em;
  color: rgba(212, 175, 55, 0.45);
}
.bc-deco-foot {
  display: flex;
  justify-content: space-between;
  margin-top: auto;
  padding-top: 20px;
  font-size: 6.5px;
  letter-spacing: 0.24em;
  color: rgba(212, 175, 55, 0.4);
}
@keyframes bc-deco-in {
  from { opacity: 0; transform: translateY(22px) scale(0.985); }
  to { opacity: 1; transform: none; }
}
@keyframes bc-deco-fade {
  from { opacity: 0; }
  to { opacity: 1; }
}
@media (prefers-reduced-motion: reduce) {
  .bc-deco-root *,
  .bc-deco-root *::before,
  .bc-deco-root *::after {
    animation: none !important;
    transition: none !important;
  }
}
</style>
