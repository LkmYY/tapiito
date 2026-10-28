<template>
  <div class="ac-manga-root relative min-h-screen overflow-hidden px-4 py-8 text-[#111]">
    <div class="ac-manga-dots"></div>
    <div class="relative z-10 mx-auto w-full max-w-md">
      <div class="ac-manga-page">
        <header class="ac-manga-meta">
          <span>少年エース · 第 {{ nfcId || '01' }} 话</span>
          <span>SCREENTONE / 网点纸</span>
        </header>

        <section class="ac-manga-hero">
          <div class="ac-manga-speed"></div>
          <div class="ac-manga-chapter">CH.01</div>
          <p class="ac-manga-kicker">CHARACTER FILE</p>
          <h1 class="ac-manga-name">{{ p.name || '未命名角色' }}</h1>
          <div class="ac-manga-title-wrap">
            <p class="ac-manga-title">{{ p.title }}</p>
          </div>
          <span class="ac-manga-series">{{ p.series || '未知作品' }}</span>
          <span class="ac-manga-element">{{ p.element }}</span>
        </section>

        <section class="ac-manga-talk">
          <div class="ac-manga-avatar">{{ initial }}</div>
          <div class="ac-manga-bubble">
            <p>{{ p.quote || '……还没有台词。' }}</p>
            <span class="ac-manga-tail"></span>
          </div>
        </section>

        <div class="ac-manga-bottom">
          <section class="ac-manga-cell">
            <p class="ac-manga-label">PROFILE</p>
            <dl>
              <div><dt>作品</dt><dd>{{ p.series || '未知作品' }}</dd></div>
              <div><dt>属性</dt><dd class="ac-manga-strong">{{ p.element }}</dd></div>
              <div><dt>社交</dt><dd class="ac-manga-handle">@{{ p.social }}</dd></div>
            </dl>
          </section>
          <section class="ac-manga-stamp-cell">
            <span class="ac-manga-stamp">承認<br />OK</span>
            <p class="ac-manga-barcode" aria-hidden="true"></p>
            <p class="ac-manga-id">NFC {{ nfcId || '0000' }}</p>
          </section>
        </div>

        <footer class="ac-manga-foot">
          <span>TO BE CONTINUED ▶</span>
          <span>© {{ p.series || '星海纪元' }}</span>
        </footer>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps(['payload', 'nfcId'])
const p = computed(() => props.payload || {})
const initial = computed(() => (p.value.name || '？').toString().trim().slice(0, 1) || '？')
</script>

<style scoped>
.ac-manga-root {
  background: #e9e2d4;
  font-family: 'PingFang SC', 'Microsoft YaHei', sans-serif;
}
.ac-manga-dots {
  position: absolute;
  inset: 0;
  background-image: radial-gradient(rgba(17, 17, 17, 0.16) 0.7px, transparent 0.8px);
  background-size: 5px 5px;
  opacity: 0.55;
}
.ac-manga-page {
  border: 3px solid #111;
  background: #f7f3e9;
  box-shadow: 9px 9px 0 #111;
  padding: 12px;
  animation: ac-manga-enter 0.8s cubic-bezier(0.2, 0.9, 0.2, 1) both;
  transition: transform 0.35s ease, box-shadow 0.35s ease;
}
.ac-manga-page:hover {
  transform: translate(-2px, -2px) rotate(-0.4deg);
  box-shadow: 13px 13px 0 #111;
}
.ac-manga-meta {
  display: flex;
  justify-content: space-between;
  border-bottom: 2px solid #111;
  padding: 2px 2px 7px;
  font-family: Consolas, monospace;
  font-size: 7.5px;
  letter-spacing: 0.14em;
}
.ac-manga-meta span:last-child { color: #777; }
.ac-manga-hero {
  position: relative;
  margin-top: 10px;
  height: 232px;
  overflow: hidden;
  border: 3px solid #111;
  background: #fff;
  padding: 16px 15px;
}
.ac-manga-speed {
  position: absolute;
  inset: -20%;
  background: repeating-conic-gradient(from 0deg at 80% 45%, #111 0deg 0.7deg, transparent 0.7deg 7deg);
  opacity: 0.1;
  animation: ac-manga-rays 26s linear infinite;
}
.ac-manga-hero::after {
  content: '';
  position: absolute;
  right: -55px;
  bottom: -70px;
  width: 230px;
  height: 230px;
  border: 3px solid #111;
  border-radius: 50%;
  background: radial-gradient(circle at 35% 30%, #ffd9e2, #e11d48 68%);
  opacity: 0.92;
}
.ac-manga-chapter {
  position: absolute;
  top: 10px;
  right: 12px;
  z-index: 2;
  font-family: Impact, 'Arial Black', sans-serif;
  font-size: 30px;
  line-height: 1;
  color: #e11d48;
  -webkit-text-stroke: 1px #111;
  transform: rotate(4deg);
}
.ac-manga-kicker {
  position: relative;
  z-index: 2;
  display: inline-block;
  background: #111;
  padding: 3px 7px;
  color: #fff;
  font-family: Consolas, monospace;
  font-size: 7.5px;
  letter-spacing: 0.2em;
}
.ac-manga-name {
  position: relative;
  z-index: 2;
  width: 74%;
  margin-top: 12px;
  font-family: 'Noto Sans SC', 'Microsoft YaHei', sans-serif;
  font-size: 40px;
  font-weight: 900;
  line-height: 1.04;
  letter-spacing: -0.03em;
  overflow-wrap: anywhere;
}
.ac-manga-title-wrap {
  position: absolute;
  z-index: 2;
  top: 48px;
  right: 12px;
  border-left: 2px solid #111;
  padding-left: 6px;
  background: rgba(247, 243, 233, 0.72);
}
.ac-manga-title {
  writing-mode: vertical-rl;
  font-family: 'Songti SC', 'Noto Serif SC', serif;
  font-size: 13px;
  letter-spacing: 0.3em;
  line-height: 1.6;
  max-height: 158px;
}
.ac-manga-series {
  position: absolute;
  z-index: 2;
  left: 15px;
  bottom: 45px;
  border: 2px solid #111;
  background: #fff;
  padding: 3px 8px;
  font-size: 10px;
  font-weight: 700;
  letter-spacing: 0.06em;
}
.ac-manga-element {
  position: absolute;
  z-index: 2;
  left: 15px;
  bottom: 15px;
  background: #e11d48;
  padding: 4px 10px;
  color: #fff;
  font-size: 10px;
  font-weight: 700;
  letter-spacing: 0.14em;
  transform: skewX(-8deg);
}
.ac-manga-talk {
  display: flex;
  align-items: flex-start;
  gap: 10px;
  margin-top: 10px;
  border: 3px solid #111;
  background: repeating-linear-gradient(45deg, #ece5d7 0 6px, #f7f3e9 6px 12px);
  padding: 13px 12px;
}
.ac-manga-avatar {
  display: flex;
  flex: none;
  width: 46px;
  height: 46px;
  align-items: center;
  justify-content: center;
  border: 2px solid #111;
  border-radius: 50%;
  background: #111;
  color: #fff;
  font-family: 'Songti SC', serif;
  font-size: 20px;
}
.ac-manga-bubble {
  position: relative;
  flex: 1;
  border: 2.5px solid #111;
  border-radius: 16px;
  background: #fff;
  padding: 12px 14px;
}
.ac-manga-bubble p {
  font-family: 'Songti SC', 'Noto Serif SC', serif;
  font-size: 14px;
  font-weight: 600;
  line-height: 1.8;
  overflow-wrap: anywhere;
}
.ac-manga-tail {
  position: absolute;
  left: -9px;
  top: 17px;
  width: 14px;
  height: 14px;
  border-left: 2.5px solid #111;
  border-bottom: 2.5px solid #111;
  background: #fff;
  transform: rotate(45deg);
}
.ac-manga-bottom {
  display: grid;
  grid-template-columns: 1.65fr 1fr;
  gap: 10px;
  margin-top: 10px;
}
.ac-manga-cell {
  border: 3px solid #111;
  background: #fff;
  padding: 10px 11px;
}
.ac-manga-label {
  font-family: Consolas, monospace;
  font-size: 7.5px;
  letter-spacing: 0.24em;
  color: #e11d48;
}
.ac-manga-cell dl { margin-top: 7px; display: grid; gap: 5px; }
.ac-manga-cell dl div { display: flex; gap: 7px; align-items: baseline; }
.ac-manga-cell dt {
  flex: none;
  width: 30px;
  color: #999;
  font-size: 9px;
}
.ac-manga-cell dd {
  min-width: 0;
  flex: 1;
  font-size: 11px;
  font-weight: 700;
  overflow-wrap: anywhere;
}
.ac-manga-strong { color: #e11d48; }
.ac-manga-handle { font-family: Consolas, monospace; font-size: 10px !important; }
.ac-manga-stamp-cell {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: space-between;
  border: 3px solid #111;
  background: #fff;
  padding: 9px 8px;
}
.ac-manga-stamp {
  display: flex;
  width: 56px;
  height: 56px;
  align-items: center;
  justify-content: center;
  border: 2.5px solid #e11d48;
  border-radius: 50%;
  color: #e11d48;
  font-family: 'Songti SC', serif;
  font-size: 10px;
  font-weight: 800;
  line-height: 1.25;
  letter-spacing: 0.1em;
  text-align: center;
  transform: rotate(-12deg);
  opacity: 0.9;
}
.ac-manga-barcode {
  width: 100%;
  height: 18px;
  background: repeating-linear-gradient(90deg, #111 0 1px, transparent 1px 3px, #111 3px 4px, transparent 4px 7px);
  opacity: 0.7;
}
.ac-manga-id {
  font-family: Consolas, monospace;
  font-size: 6.5px;
  letter-spacing: 0.14em;
  color: #777;
}
.ac-manga-foot {
  display: flex;
  justify-content: space-between;
  border-top: 2px solid #111;
  margin-top: 10px;
  padding: 6px 2px 1px;
  font-family: Consolas, monospace;
  font-size: 7.5px;
  letter-spacing: 0.12em;
}
@keyframes ac-manga-enter {
  from { opacity: 0; transform: translateY(26px) rotate(1.5deg); }
  to { opacity: 1; transform: none; }
}
@keyframes ac-manga-rays { to { transform: rotate(360deg); } }
@media (prefers-reduced-motion: reduce) {
  .ac-manga-root *,
  .ac-manga-root *::before,
  .ac-manga-root *::after {
    animation: none !important;
    transition: none !important;
  }
}
</style>
