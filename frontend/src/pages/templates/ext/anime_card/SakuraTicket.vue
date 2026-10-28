<template>
  <div class="ac-ticket-root relative min-h-screen overflow-hidden px-5 py-9 text-[#4b2432]">
    <div class="ac-ticket-sky"></div>
    <span v-for="n in 9" :key="n" class="ac-ticket-petal" :class="'ac-ticket-p' + n"></span>

    <div class="relative z-10 mx-auto w-full max-w-md">
      <div class="ac-ticket">
        <div class="ac-ticket-main">
          <header class="ac-ticket-head">
            <span><Flower2 :size="11" />SAKURA THEATER</span>
            <span>ADMIT ONE</span>
          </header>

          <p class="ac-ticket-serial">SERIES {{ nfcId || '0001' }} · GATE 03</p>
          <h1 class="ac-ticket-name">{{ p.name || '未命名角色' }}</h1>
          <p class="ac-ticket-title">{{ p.title }}</p>

          <div class="ac-ticket-meta">
            <div><span>作品</span><b>{{ p.series || '未知作品' }}</b></div>
            <div><span>属性</span><b>{{ p.element }}</b></div>
            <div><span>席位</span><b>A-{{ nfcId || '12' }}</b></div>
          </div>

          <blockquote class="ac-ticket-quote">
            <p>{{ p.quote || '……开场前，先看看星星吧。' }}</p>
          </blockquote>

          <div class="ac-ticket-seat">
            <span>ROW A</span><span>SEAT {{ nfcId || '12' }}</span><span>19:30</span>
          </div>
        </div>

        <div class="ac-ticket-stub">
          <p class="ac-ticket-stub-top">NFC {{ nfcId || '0000' }}</p>
          <div class="ac-ticket-admit">ADMIT ONE</div>
          <div class="ac-ticket-barcode"></div>
          <p class="ac-ticket-issued">ISSUED TO</p>
          <b class="ac-ticket-social">@{{ p.social }}</b>
          <p class="ac-ticket-cut">✂ - - - - - - -</p>
        </div>
      </div>

      <p class="ac-ticket-caption">桜 の 夜 に 、 君 と 一 緒 に</p>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { Flower2 } from 'lucide-vue-next'

const props = defineProps(['payload', 'nfcId'])
const p = computed(() => props.payload || {})
</script>

<style scoped>
.ac-ticket-root {
  background: #f7dfe7;
  font-family: 'PingFang SC', 'Microsoft YaHei', sans-serif;
}
.ac-ticket-sky {
  position: absolute;
  inset: 0;
  background:
    radial-gradient(circle at 82% 12%, rgba(255, 255, 255, 0.8), transparent 26%),
    radial-gradient(circle at 12% 86%, rgba(251, 113, 133, 0.25), transparent 32%),
    linear-gradient(160deg, #fce7f3, #fbcfe8 45%, #fecdd3);
}
.ac-ticket-petal {
  position: absolute;
  width: 11px;
  height: 8px;
  border-radius: 80% 0 80% 0;
  background: linear-gradient(135deg, #fda4af, #fb7185);
  opacity: 0.55;
  animation: ac-ticket-fall 14s linear infinite;
}
.ac-ticket-p1 { left: 8%; top: -6%; animation-delay: 0s; }
.ac-ticket-p2 { left: 22%; top: -8%; animation-delay: 2s; width: 8px; height: 6px; }
.ac-ticket-p3 { left: 36%; top: -4%; animation-delay: 4.5s; }
.ac-ticket-p4 { left: 52%; top: -7%; animation-delay: 1.2s; width: 7px; height: 5px; }
.ac-ticket-p5 { left: 66%; top: -5%; animation-delay: 6s; }
.ac-ticket-p6 { left: 78%; top: -9%; animation-delay: 3.2s; width: 9px; height: 7px; }
.ac-ticket-p7 { left: 90%; top: -4%; animation-delay: 7.5s; }
.ac-ticket-p8 { left: 45%; top: -8%; animation-delay: 9s; width: 6px; height: 5px; }
.ac-ticket-p9 { left: 84%; top: -3%; animation-delay: 11s; width: 8px; height: 6px; }
.ac-ticket {
  position: relative;
  display: flex;
  border-radius: 16px;
  background: #fffaf8;
  box-shadow: 0 26px 60px -24px rgba(190, 24, 93, 0.45), 0 2px 0 rgba(255, 255, 255, 0.9) inset;
  animation: ac-ticket-enter 0.95s cubic-bezier(0.2, 1, 0.3, 1) both;
  transition: transform 0.5s ease, box-shadow 0.5s ease;
}
.ac-ticket:hover {
  transform: translateY(-5px) rotate(-0.6deg);
  box-shadow: 0 36px 74px -24px rgba(190, 24, 93, 0.55), 0 2px 0 rgba(255, 255, 255, 0.9) inset;
}
.ac-ticket::before,
.ac-ticket::after {
  content: '';
  position: absolute;
  left: 70.5%;
  width: 20px;
  height: 20px;
  border-radius: 50%;
  background: #f7dfe7;
  z-index: 2;
  box-shadow: inset 0 0 6px rgba(190, 24, 93, 0.12);
}
.ac-ticket::before { top: -10px; }
.ac-ticket::after { bottom: -10px; }
.ac-ticket-main { flex: 1; min-width: 0; padding: 20px 16px 16px 20px; }
.ac-ticket-head {
  display: flex;
  justify-content: space-between;
  border-bottom: 1px solid #fbcfe8;
  padding-bottom: 8px;
  font-family: Consolas, monospace;
  font-size: 7px;
  letter-spacing: 0.14em;
  color: #be185d;
}
.ac-ticket-head span:first-child { display: inline-flex; align-items: center; gap: 4px; }
.ac-ticket-head span:last-child { color: #f472b6; }
.ac-ticket-serial {
  margin-top: 13px;
  font-family: Consolas, monospace;
  font-size: 6.5px;
  letter-spacing: 0.16em;
  color: #c99aa9;
}
.ac-ticket-name {
  margin-top: 4px;
  font-size: 35px;
  font-weight: 800;
  line-height: 1.08;
  letter-spacing: 0.02em;
  color: #831843;
  overflow-wrap: anywhere;
}
.ac-ticket-title {
  margin-top: 5px;
  font-size: 11.5px;
  letter-spacing: 0.2em;
  color: #be185d;
}
.ac-ticket-meta {
  display: grid;
  gap: 4px;
  margin-top: 15px;
  border-top: 1px dashed #f9a8d4;
  border-bottom: 1px dashed #f9a8d4;
  padding: 9px 0;
}
.ac-ticket-meta div { display: flex; align-items: baseline; gap: 9px; }
.ac-ticket-meta span {
  flex: none;
  width: 32px;
  font-size: 8px;
  letter-spacing: 0.1em;
  color: #c99aa9;
}
.ac-ticket-meta b {
  min-width: 0;
  flex: 1;
  text-align: right;
  font-size: 10.5px;
  font-weight: 600;
  color: #6b2b40;
  overflow-wrap: anywhere;
}
.ac-ticket-quote {
  margin-top: 13px;
  border-left: 2px solid #fb7185;
  padding-left: 10px;
}
.ac-ticket-quote p {
  font-family: 'Songti SC', 'Noto Serif SC', serif;
  font-size: 12.5px;
  font-style: italic;
  line-height: 1.8;
  color: #8c4058;
  overflow-wrap: anywhere;
}
.ac-ticket-seat {
  display: flex;
  justify-content: space-between;
  margin-top: 15px;
  font-family: Consolas, monospace;
  font-size: 7px;
  letter-spacing: 0.1em;
  color: #b87c8e;
}
.ac-ticket-stub {
  position: relative;
  flex: none;
  width: 96px;
  border-left: 2px dashed #f9a8d4;
  padding: 20px 10px 14px;
  display: flex;
  flex-direction: column;
  align-items: center;
}
.ac-ticket-stub-top {
  font-family: Consolas, monospace;
  font-size: 6px;
  letter-spacing: 0.12em;
  color: #c99aa9;
}
.ac-ticket-admit {
  margin-top: 16px;
  writing-mode: vertical-rl;
  font-family: 'Arial Black', Impact, sans-serif;
  font-size: 17px;
  letter-spacing: 0.2em;
  color: #e11d48;
}
.ac-ticket-barcode {
  width: 100%;
  height: 27px;
  margin-top: 14px;
  background: repeating-linear-gradient(90deg, #831843 0 1px, transparent 1px 3px, #831843 3px 4px, transparent 4px 6px, #831843 6px 8px, transparent 8px 11px);
  opacity: 0.7;
}
.ac-ticket-issued {
  margin-top: 13px;
  font-family: Consolas, monospace;
  font-size: 5.5px;
  letter-spacing: 0.14em;
  color: #c99aa9;
}
.ac-ticket-social {
  margin-top: 2px;
  max-width: 100%;
  text-align: center;
  font-family: Consolas, monospace;
  font-size: 8.5px;
  font-weight: 500;
  color: #be185d;
  overflow-wrap: anywhere;
}
.ac-ticket-cut {
  margin-top: auto;
  padding-top: 13px;
  font-family: Consolas, monospace;
  font-size: 7px;
  color: #e5b8c5;
  white-space: nowrap;
}
.ac-ticket-caption {
  margin-top: 24px;
  text-align: center;
  font-family: 'Songti SC', serif;
  font-size: 10px;
  letter-spacing: 0.32em;
  color: #b05c76;
  animation: ac-ticket-fade 1.4s 0.6s both;
}
@keyframes ac-ticket-enter {
  from { opacity: 0; transform: translateX(-30px) rotate(-2deg); }
  to { opacity: 1; transform: none; }
}
@keyframes ac-ticket-fade { from { opacity: 0; } to { opacity: 1; } }
@keyframes ac-ticket-fall {
  0% { transform: translateY(-10vh) rotate(0deg) translateX(0); opacity: 0; }
  10% { opacity: 0.6; }
  100% { transform: translateY(112vh) rotate(520deg) translateX(40px); opacity: 0; }
}
@media (prefers-reduced-motion: reduce) {
  .ac-ticket-root *,
  .ac-ticket-root *::before,
  .ac-ticket-root *::after {
    animation: none !important;
    transition: none !important;
  }
}
</style>
