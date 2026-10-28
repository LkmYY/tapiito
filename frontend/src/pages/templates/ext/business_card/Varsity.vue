<template>
  <div class="bc-varsity-root min-h-screen px-5 py-8">
    <div class="bc-varsity-wrap mx-auto flex min-h-[calc(100vh-4rem)] w-full max-w-md flex-col">
      <header class="bc-varsity-head">
        <span class="bc-varsity-est">EST. MMXXVI</span>
        <span class="bc-varsity-club">THE CONNECTION CLUB</span>
      </header>

      <div class="bc-varsity-stripe"></div>

      <section class="bc-varsity-hero">
        <div class="bc-varsity-badge">
          <div class="bc-varsity-badge-inner">
            <span class="bc-varsity-star">★</span>
            <span class="bc-varsity-badge-initial">{{ initial }}</span>
            <span class="bc-varsity-badge-txt">YUNTU<br />SOCIETY</span>
            <span class="bc-varsity-star">★</span>
          </div>
        </div>
        <p class="bc-varsity-kicker">MEMBER OF THE YEAR</p>
        <h1 class="bc-varsity-name">{{ payload.name || '未命名' }}</h1>
        <p class="bc-varsity-role">{{ payload.title }}</p>
        <p class="bc-varsity-company">{{ payload.company }}</p>
        <span class="bc-varsity-number">07</span>
      </section>

      <div class="bc-varsity-stripe bc-varsity-stripe-low"></div>

      <section class="bc-varsity-roster">
        <div class="bc-varsity-roster-head">
          <span>TEAM ROSTER / CONTACT</span>
          <span>{{ nfcId || 'NO. 07' }}</span>
        </div>
        <div class="bc-varsity-row">
          <span class="bc-varsity-pos">TEL</span>
          <a v-if="payload.phone" :href="'tel:' + payload.phone">{{ payload.phone }}</a>
          <span class="bc-varsity-num">01</span>
        </div>
        <div class="bc-varsity-row">
          <span class="bc-varsity-pos">MAIL</span>
          <a v-if="payload.email" :href="'mailto:' + payload.email">{{ payload.email }}</a>
          <span class="bc-varsity-num">02</span>
        </div>
        <div class="bc-varsity-row">
          <span class="bc-varsity-pos">WEB</span>
          <a v-if="payload.website" :href="payload.website" target="_blank" rel="noopener noreferrer">{{ payload.website }}</a>
          <span class="bc-varsity-num">03</span>
        </div>
        <div class="bc-varsity-row">
          <span class="bc-varsity-pos">WX</span>
          <span :class="{ 'bc-varsity-empty': !payload.wechat }">{{ payload.wechat }}</span>
          <span class="bc-varsity-num">04</span>
        </div>
      </section>

      <footer class="bc-varsity-foot">
        <span class="bc-varsity-pennant">GO<br />TEAM</span>
        <p>TRADITION · CONNECTION<br />GOOD SPORTSMANSHIP</p>
        <span class="bc-varsity-year">'26</span>
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
.bc-varsity-root {
  position: relative;
  overflow: hidden;
  background: #f4eddb;
  font-family: Georgia, 'Songti SC', 'PingFang SC', serif;
  color: #431407;
}
.bc-varsity-root::before {
  content: '';
  position: absolute;
  inset: 0;
  background-image: radial-gradient(rgba(67, 20, 7, 0.16) 0.5px, transparent 0.5px);
  background-size: 5px 5px;
  opacity: 0.3;
}
.bc-varsity-wrap {
  position: relative;
  z-index: 1;
}
.bc-varsity-head {
  display: flex;
  justify-content: space-between;
  align-items: baseline;
  padding-bottom: 10px;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6.5px;
  letter-spacing: 0.18em;
  color: #7f1d1d;
  animation: bc-varsity-fade 0.8s both;
}
.bc-varsity-club {
  font-weight: 700;
}
.bc-varsity-stripe {
  height: 12px;
  background: repeating-linear-gradient(90deg, #7f1d1d 0 16px, #f4eddb 16px 24px, #1e3a8a 24px 40px, #f4eddb 40px 48px);
  border-top: 2px solid #7f1d1d;
  border-bottom: 2px solid #7f1d1d;
  transform-origin: left;
  animation: bc-varsity-stripe 1s 0.15s cubic-bezier(0.22, 1, 0.36, 1) both;
}
.bc-varsity-hero {
  position: relative;
  padding: 34px 0 26px;
  text-align: center;
  animation: bc-varsity-in 0.9s 0.1s cubic-bezier(0.22, 1, 0.36, 1) both;
}
.bc-varsity-badge {
  display: flex;
  width: 142px;
  height: 142px;
  margin: 0 auto;
  align-items: center;
  justify-content: center;
  border-radius: 50%;
  background: #7f1d1d;
  box-shadow: 0 0 0 4px #f4eddb, 0 0 0 6px #7f1d1d, 6px 10px 22px -10px rgba(67, 20, 7, 0.55);
  transition: transform 0.5s cubic-bezier(0.22, 1, 0.36, 1);
}
.bc-varsity-badge:hover {
  transform: rotate(-5deg) scale(1.03);
}
.bc-varsity-badge-inner {
  display: flex;
  width: 124px;
  height: 124px;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  gap: 2px;
  border: 2px dashed rgba(244, 237, 219, 0.55);
  border-radius: 50%;
  color: #f4eddb;
}
.bc-varsity-badge-initial {
  font-size: 42px;
  font-weight: 700;
  line-height: 1;
}
.bc-varsity-badge-txt {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 5.5px;
  line-height: 1.5;
  letter-spacing: 0.14em;
}
.bc-varsity-star {
  font-size: 7px;
  color: #fbbf24;
}
.bc-varsity-kicker {
  margin-top: 20px;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6px;
  letter-spacing: 0.3em;
  color: #1e3a8a;
}
.bc-varsity-name {
  margin-top: 10px;
  font-family: Georgia, 'Songti SC', serif;
  font-size: clamp(38px, 12.5vw, 52px);
  font-weight: 700;
  line-height: 1;
  letter-spacing: -0.03em;
  text-transform: uppercase;
  overflow-wrap: anywhere;
  color: #7f1d1d;
  text-shadow: 2px 2px 0 rgba(251, 191, 36, 0.65);
}
.bc-varsity-role {
  margin-top: 12px;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 8.5px;
  letter-spacing: 0.16em;
  color: #1e3a8a;
}
.bc-varsity-company {
  margin-top: 3px;
  font-size: 10px;
  font-style: italic;
  color: rgba(67, 20, 7, 0.55);
}
.bc-varsity-number {
  position: absolute;
  top: 30px;
  right: 2px;
  font-size: 46px;
  font-weight: 700;
  line-height: 1;
  color: rgba(127, 29, 29, 0.14);
}
.bc-varsity-stripe-low {
  background: repeating-linear-gradient(90deg, #1e3a8a 0 16px, #f4eddb 16px 24px, #7f1d1d 24px 40px, #f4eddb 40px 48px);
}
.bc-varsity-roster {
  padding-top: 16px;
  animation: bc-varsity-fade 1s 0.35s both;
}
.bc-varsity-roster-head {
  display: flex;
  justify-content: space-between;
  padding-bottom: 6px;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 5.5px;
  letter-spacing: 0.18em;
  color: rgba(67, 20, 7, 0.55);
}
.bc-varsity-row {
  display: grid;
  grid-template-columns: 46px 1fr 24px;
  align-items: baseline;
  gap: 6px;
  padding: 11px 2px;
  border-top: 1px solid rgba(127, 29, 29, 0.28);
  transition: background 0.2s ease, color 0.2s ease, padding-left 0.2s ease;
}
.bc-varsity-row:hover {
  background: #7f1d1d;
  color: #f4eddb;
  padding-left: 8px;
}
.bc-varsity-pos {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6.5px;
  font-weight: 700;
  letter-spacing: 0.1em;
  color: #1e3a8a;
  transition: color 0.2s ease;
}
.bc-varsity-row:hover .bc-varsity-pos {
  color: #fbbf24;
}
.bc-varsity-row a,
.bc-varsity-row span:nth-child(2),
.bc-varsity-row em {
  min-width: 0;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 10px;
  font-style: normal;
  color: inherit;
  text-decoration: none;
  transition: color 0.2s ease;
}
.bc-varsity-row a:hover {
  text-decoration: underline;
}
.bc-varsity-row a:active {
  opacity: 0.6;
}
.bc-varsity-row em,
.bc-varsity-empty {
  color: rgba(67, 20, 7, 0.32);
}
.bc-varsity-row:hover em,
.bc-varsity-row:hover .bc-varsity-empty {
  color: rgba(244, 237, 219, 0.5);
}
.bc-varsity-num {
  text-align: right;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 5.5px;
  color: rgba(67, 20, 7, 0.4);
}
.bc-varsity-row:hover .bc-varsity-num {
  color: rgba(244, 237, 219, 0.6);
}
.bc-varsity-foot {
  display: flex;
  align-items: center;
  gap: 14px;
  margin-top: auto;
  padding-top: 24px;
}
.bc-varsity-pennant {
  border: 2px solid #7f1d1d;
  background: #fbbf24;
  padding: 5px 12px;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 8px;
  font-weight: 700;
  line-height: 1.3;
  letter-spacing: 0.1em;
  color: #7f1d1d;
  clip-path: polygon(0 0, 100% 0, 100% 78%, 50% 100%, 0 78%);
  transition: transform 0.3s ease;
}
.bc-varsity-pennant:hover {
  transform: rotate(-4deg) translateY(-2px);
}
.bc-varsity-foot p {
  flex: 1;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 5.5px;
  line-height: 1.7;
  letter-spacing: 0.1em;
  color: rgba(67, 20, 7, 0.5);
}
.bc-varsity-year {
  font-size: 20px;
  font-weight: 700;
  color: #7f1d1d;
}
@keyframes bc-varsity-in {
  from { opacity: 0; transform: translateY(20px); }
  to { opacity: 1; transform: none; }
}
@keyframes bc-varsity-fade {
  from { opacity: 0; }
  to { opacity: 1; }
}
@keyframes bc-varsity-stripe {
  from { transform: scaleX(0); }
  to { transform: scaleX(1); }
}
@media (prefers-reduced-motion: reduce) {
  .bc-varsity-root *,
  .bc-varsity-root *::before,
  .bc-varsity-root *::after {
    animation: none !important;
    transition: none !important;
  }
}
</style>
