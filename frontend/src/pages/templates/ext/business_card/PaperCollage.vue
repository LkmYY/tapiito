<template>
  <div class="bc-collage-root min-h-screen px-5 py-8">
    <div class="bc-collage-board mx-auto w-full max-w-md">
      <header class="bc-collage-head">
        <span class="bc-collage-stamp">HELLO</span>
        <span class="bc-collage-head-note">my little card<br />{{ nfcId ? '#' + nfcId : '—' }}</span>
      </header>

      <div class="bc-collage-stack">
        <div class="bc-collage-tape bc-collage-tape-a"></div>
        <article class="bc-collage-photo">
          <div class="bc-collage-photo-frame">
            <span>{{ initial }}</span>
            <i class="bc-collage-spark bc-collage-spark-a"></i>
            <i class="bc-collage-spark bc-collage-spark-b"></i>
          </div>
          <p class="bc-collage-caption">that's me, probably.</p>
        </article>

        <div class="bc-collage-tape bc-collage-tape-b"></div>
        <article class="bc-collage-namecard">
          <p class="bc-collage-hi">nice to meet you!</p>
          <h1>{{ payload.name || '未命名' }}</h1>
          <p class="bc-collage-role">{{ payload.title }}</p>
          <p class="bc-collage-company">@ {{ payload.company }}</p>
        </article>

        <div class="bc-collage-sticky">
          <span>call me maybe</span>
          <a v-if="payload.phone" :href="'tel:' + payload.phone">{{ payload.phone }}</a>
        </div>

        <article class="bc-collage-note">
          <div class="bc-collage-note-row">
            <span>MAIL</span>
            <a v-if="payload.email" :href="'mailto:' + payload.email">{{ payload.email }}</a>
          </div>
          <div class="bc-collage-note-row">
            <span>WEB</span>
            <a v-if="payload.website" :href="payload.website" target="_blank" rel="noopener noreferrer">{{ payload.website }}</a>
          </div>
          <div class="bc-collage-note-row">
            <span>WECHAT</span>
            <p>{{ payload.wechat }}</p>
          </div>
          <div class="bc-collage-arrow">→</div>
        </article>

        <div class="bc-collage-ticket">
          <span>ADMIT ONE FRIEND</span>
          <strong>{{ initial }}</strong>
          <span>NO. {{ nfcId || '000' }}</span>
        </div>

        <div class="bc-collage-doodle">✿</div>
      </div>

      <footer class="bc-collage-foot">
        <span>PASTED WITH CARE</span>
        <span>{{ payload.company }}</span>
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
.bc-collage-root {
  position: relative;
  overflow: hidden;
  background: #c9a97c;
  font-family: 'Segoe Print', 'Bradley Hand', 'PingFang SC', 'Microsoft YaHei', cursive;
}
.bc-collage-root::before {
  content: '';
  position: absolute;
  inset: 0;
  background:
    radial-gradient(circle at 18% 12%, rgba(255, 255, 255, 0.3), transparent 32%),
    repeating-linear-gradient(0deg, rgba(92, 64, 38, 0.06) 0 1px, transparent 1px 4px),
    repeating-linear-gradient(90deg, rgba(92, 64, 38, 0.04) 0 1px, transparent 1px 5px);
  pointer-events: none;
}
.bc-collage-board {
  position: relative;
  z-index: 1;
}
.bc-collage-head {
  display: flex;
  align-items: flex-start;
  justify-content: space-between;
  padding: 2px 4px 16px;
}
.bc-collage-stamp {
  display: inline-block;
  border: 2px dashed #7c2d12;
  padding: 5px 10px;
  font-family: 'Courier New', monospace;
  font-size: 12px;
  font-weight: 700;
  letter-spacing: 0.16em;
  color: #7c2d12;
  transform: rotate(-4deg);
  background: rgba(255, 247, 237, 0.45);
}
.bc-collage-head-note {
  text-align: right;
  font-size: 11px;
  line-height: 1.5;
  color: rgba(69, 26, 3, 0.7);
  transform: rotate(2deg);
}
.bc-collage-stack {
  position: relative;
  padding-top: 8px;
}
.bc-collage-tape {
  position: absolute;
  z-index: 5;
  width: 92px;
  height: 26px;
  background: rgba(254, 240, 138, 0.78);
  box-shadow: 0 1px 2px rgba(69, 26, 3, 0.15);
  clip-path: polygon(2% 8%, 98% 0, 100% 92%, 0 100%);
}
.bc-collage-tape-a {
  top: 0;
  left: 50%;
  transform: translateX(-50%) rotate(-5deg);
}
.bc-collage-tape-b {
  top: 268px;
  right: 12px;
  transform: rotate(8deg);
  background: rgba(244, 114, 182, 0.45);
}
.bc-collage-photo {
  position: relative;
  z-index: 3;
  width: 62%;
  margin: 0 auto;
  padding: 12px 12px 30px;
  background: #fffdf7;
  box-shadow: 6px 10px 24px -10px rgba(69, 26, 3, 0.45);
  transform: rotate(-3.5deg);
  transition: transform 0.4s cubic-bezier(0.22, 1, 0.36, 1);
  animation: bc-collage-slap 0.8s cubic-bezier(0.22, 1, 0.36, 1) both;
}
.bc-collage-photo:hover {
  transform: rotate(0deg) translateY(-4px);
}
.bc-collage-photo-frame {
  position: relative;
  display: flex;
  aspect-ratio: 1 / 1;
  align-items: center;
  justify-content: center;
  overflow: hidden;
  background:
    radial-gradient(circle at 30% 25%, #fbcfe8, transparent 55%),
    linear-gradient(150deg, #fde68a, #fca5a5 55%, #c4b5fd);
}
.bc-collage-photo-frame span {
  font-family: Georgia, 'Songti SC', serif;
  font-size: 84px;
  color: #fffdf7;
  text-shadow: 3px 4px 0 rgba(69, 26, 3, 0.18);
}
.bc-collage-spark {
  position: absolute;
  width: 12px;
  height: 12px;
  border: 2px solid #fffdf7;
  transform: rotate(45deg);
}
.bc-collage-spark-a { top: 14px; right: 16px; }
.bc-collage-spark-b { bottom: 20px; left: 18px; width: 7px; height: 7px; }
.bc-collage-caption {
  position: absolute;
  left: 0;
  right: 0;
  bottom: 9px;
  text-align: center;
  font-size: 10px;
  color: #92400e;
}
.bc-collage-namecard {
  position: relative;
  z-index: 2;
  margin: -12px 4px 0;
  padding: 28px 20px 20px;
  border: 1px solid rgba(120, 53, 15, 0.22);
  background: #fff7ed;
  box-shadow: 4px 8px 22px -12px rgba(69, 26, 3, 0.5);
  transform: rotate(1.4deg);
  animation: bc-collage-rise 0.8s 0.15s cubic-bezier(0.22, 1, 0.36, 1) both;
}
.bc-collage-hi {
  font-size: 11px;
  color: #b45309;
}
.bc-collage-namecard h1 {
  margin-top: 8px;
  font-family: Georgia, 'Songti SC', serif;
  font-size: 38px;
  line-height: 1;
  letter-spacing: -0.03em;
  color: #431407;
  overflow-wrap: anywhere;
}
.bc-collage-role {
  margin-top: 8px;
  font-family: 'Courier New', monospace;
  font-size: 11px;
  font-weight: 700;
  letter-spacing: 0.1em;
  color: #9a3412;
}
.bc-collage-company {
  margin-top: 3px;
  font-size: 11px;
  color: #a16207;
}
.bc-collage-sticky {
  position: relative;
  z-index: 4;
  width: 72%;
  margin: 16px 0 0 auto;
  padding: 16px 14px 14px;
  background: #fef08a;
  box-shadow: 3px 6px 14px -8px rgba(69, 26, 3, 0.55);
  transform: rotate(2.6deg);
  transition: transform 0.3s ease;
}
.bc-collage-sticky:hover {
  transform: rotate(-1deg) translateY(-2px);
}
.bc-collage-sticky span {
  display: block;
  font-size: 9px;
  color: #a16207;
}
.bc-collage-sticky a,
.bc-collage-sticky em {
  display: block;
  margin-top: 5px;
  font-family: 'Courier New', monospace;
  font-size: 14px;
  font-style: normal;
  font-weight: 700;
  letter-spacing: 0.02em;
  color: #422006;
  text-decoration: none;
  overflow-wrap: anywhere;
}
.bc-collage-sticky a:hover {
  text-decoration: underline wavy #b45309;
}
.bc-collage-note {
  position: relative;
  z-index: 2;
  margin: -20px 34px 0 0;
  padding: 20px 16px 26px;
  border-left: 3px solid #fb7185;
  background: #fffdf7;
  box-shadow: 4px 8px 20px -14px rgba(69, 26, 3, 0.6);
  transform: rotate(-1.8deg);
  animation: bc-collage-rise 0.9s 0.28s cubic-bezier(0.22, 1, 0.36, 1) both;
}
.bc-collage-note-row {
  display: flex;
  align-items: baseline;
  gap: 10px;
  padding: 7px 0;
  border-bottom: 1px dashed rgba(180, 83, 9, 0.28);
}
.bc-collage-note-row:last-of-type {
  border-bottom: 0;
}
.bc-collage-note-row span {
  flex: none;
  width: 42px;
  font-family: 'Courier New', monospace;
  font-size: 8px;
  font-weight: 700;
  letter-spacing: 0.08em;
  color: #b45309;
}
.bc-collage-note-row a,
.bc-collage-note-row em,
.bc-collage-note-row p {
  min-width: 0;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  font-size: 11px;
  font-style: normal;
  color: #431407;
  text-decoration: none;
}
.bc-collage-note-row a:hover {
  color: #be123c;
  text-decoration: underline;
}
.bc-collage-note-row em {
  color: #d6b28a;
}
.bc-collage-arrow {
  position: absolute;
  right: 12px;
  bottom: 8px;
  font-size: 18px;
  color: #fb7185;
  transform: rotate(-8deg);
}
.bc-collage-ticket {
  position: relative;
  z-index: 5;
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin: 18px 8px 0;
  padding: 9px 12px;
  border: 1px dashed #78350f;
  background: rgba(255, 247, 237, 0.8);
  font-family: 'Courier New', monospace;
  font-size: 7px;
  letter-spacing: 0.1em;
  color: #78350f;
  transform: rotate(1deg);
}
.bc-collage-ticket strong {
  font-family: Georgia, serif;
  font-size: 18px;
}
.bc-collage-doodle {
  position: absolute;
  top: 238px;
  right: 0;
  font-size: 28px;
  color: #9d174d;
  transform: rotate(16deg);
  animation: bc-collage-wiggle 5s ease-in-out infinite;
}
.bc-collage-foot {
  display: flex;
  justify-content: space-between;
  margin-top: 22px;
  padding: 0 4px;
  font-family: 'Courier New', monospace;
  font-size: 7px;
  letter-spacing: 0.14em;
  color: rgba(69, 26, 3, 0.5);
}
@keyframes bc-collage-slap {
  from { opacity: 0; transform: translateY(-24px) rotate(-9deg) scale(1.04); }
  to { opacity: 1; transform: rotate(-3.5deg); }
}
@keyframes bc-collage-rise {
  from { opacity: 0; transform: translateY(18px) rotate(0deg); }
  to { opacity: 1; }
}
@keyframes bc-collage-wiggle {
  0%, 100% { transform: rotate(16deg); }
  50% { transform: rotate(24deg) translateY(-3px); }
}
@media (prefers-reduced-motion: reduce) {
  .bc-collage-root *,
  .bc-collage-root *::before,
  .bc-collage-root *::after {
    animation: none !important;
    transition: none !important;
  }
}
</style>
