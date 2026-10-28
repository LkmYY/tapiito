<template>
  <div class="ac-gacha-root relative min-h-screen overflow-hidden px-5 py-8 text-white">
    <div class="ac-gacha-rays"></div>
    <div class="ac-gacha-vignette"></div>
    <div class="ac-gacha-dust"></div>

    <div class="relative z-10 mx-auto w-full max-w-md">
      <header class="ac-gacha-top">
        <span class="ac-gacha-tag">LIMITED PICK-UP</span>
        <span class="ac-gacha-top-id">ID {{ nfcId || '0000' }}</span>
      </header>

      <section class="ac-gacha-hero">
        <div class="ac-gacha-watermark">UR</div>
        <article class="ac-gacha-card">
          <i class="ac-gacha-corner ac-gacha-c1"></i>
          <i class="ac-gacha-corner ac-gacha-c2"></i>
          <i class="ac-gacha-corner ac-gacha-c3"></i>
          <i class="ac-gacha-corner ac-gacha-c4"></i>

          <p class="ac-gacha-overline"><Sparkles :size="12" />SUMMON SUCCESS</p>
          <p class="ac-gacha-rate">SSR · 0.600%</p>

          <div class="ac-gacha-stars" aria-label="五星角色">
            <Star v-for="n in 5" :key="n" :size="17" fill="currentColor" />
          </div>

          <h1 class="ac-gacha-name">{{ p.name || '未命名角色' }}</h1>
          <p class="ac-gacha-title">「{{ p.title }}」</p>

          <div class="ac-gacha-ribbon">
            <span>{{ p.series || '未知作品' }}</span>
            <i></i>
            <span class="ac-gacha-element">{{ p.element }}</span>
          </div>

          <blockquote class="ac-gacha-quote">{{ p.quote || '尚未收录台词。' }}</blockquote>

          <div class="ac-gacha-sign">
            <span>PLAYER</span>
            <b><AtSign :size="12" />{{ p.social }}</b>
          </div>
        </article>
      </section>

      <footer class="ac-gacha-foot">
        <div class="ac-gacha-progress"><span></span></div>
        <p>召唤记录已保存 · 再次召唤需 300 星石</p>
        <p class="ac-gacha-serial">CARD SERIES / A-{{ nfcId || '0000' }}</p>
      </footer>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { AtSign, Sparkles, Star } from 'lucide-vue-next'

const props = defineProps(['payload', 'nfcId'])
const p = computed(() => props.payload || {})
</script>

<style scoped>
.ac-gacha-root {
  background: radial-gradient(circle at 50% 34%, #3b1c05 0%, #150b03 44%, #060402 100%);
  font-family: 'Optima', 'PingFang SC', 'Microsoft YaHei', sans-serif;
}
.ac-gacha-rays {
  position: absolute;
  top: -30%;
  left: 50%;
  width: 900px;
  height: 900px;
  transform: translateX(-50%);
  background: repeating-conic-gradient(from 0deg, rgba(251, 191, 36, 0.22) 0deg 7deg, transparent 7deg 18deg);
  mask-image: radial-gradient(circle, #000 8%, transparent 62%);
  animation: ac-gacha-spin 60s linear infinite;
  pointer-events: none;
}
.ac-gacha-vignette {
  position: absolute;
  inset: 0;
  background: radial-gradient(ellipse at 50% 38%, transparent 30%, rgba(0, 0, 0, 0.72) 100%);
  pointer-events: none;
}
.ac-gacha-dust {
  position: absolute;
  inset: 0;
  background-image:
    radial-gradient(rgba(253, 230, 138, 0.55) 0.7px, transparent 0.8px),
    radial-gradient(rgba(255, 255, 255, 0.3) 0.6px, transparent 0.7px);
  background-size: 47px 47px, 83px 83px;
  background-position: 0 0, 24px 38px;
  opacity: 0.5;
  animation: ac-gacha-dust 16s linear infinite;
  pointer-events: none;
}
.ac-gacha-top {
  display: flex;
  align-items: center;
  justify-content: space-between;
  font-family: Consolas, monospace;
  font-size: 8px;
  letter-spacing: 0.24em;
  color: rgba(253, 230, 138, 0.55);
  animation: ac-gacha-fade 0.8s both;
}
.ac-gacha-tag {
  border: 1px solid rgba(253, 230, 138, 0.45);
  padding: 4px 8px;
  background: rgba(120, 53, 15, 0.3);
}
.ac-gacha-top-id { color: rgba(255, 255, 255, 0.3); }
.ac-gacha-hero {
  position: relative;
  margin-top: 60px;
  animation: ac-gacha-reveal 1.1s cubic-bezier(0.16, 1, 0.3, 1) both;
}
.ac-gacha-watermark {
  position: absolute;
  top: -78px;
  left: 50%;
  transform: translateX(-50%);
  font-family: Georgia, 'Times New Roman', serif;
  font-size: 170px;
  font-weight: 700;
  line-height: 1;
  letter-spacing: -0.06em;
  color: transparent;
  -webkit-text-stroke: 1px rgba(253, 230, 138, 0.28);
  filter: drop-shadow(0 0 26px rgba(251, 191, 36, 0.25));
  user-select: none;
}
.ac-gacha-card {
  position: relative;
  border: 1px solid rgba(253, 230, 138, 0.55);
  background:
    linear-gradient(165deg, rgba(58, 36, 8, 0.96), rgba(15, 9, 3, 0.98) 62%),
    #0d0803;
  padding: 30px 22px 24px;
  box-shadow:
    0 0 0 1px rgba(0, 0, 0, 0.6),
    0 0 60px rgba(245, 158, 11, 0.16),
    inset 0 0 50px rgba(120, 53, 15, 0.28);
  transition: transform 0.5s ease, box-shadow 0.5s ease;
}
.ac-gacha-card:hover {
  transform: translateY(-5px) rotateX(3deg) rotateY(-4deg);
  box-shadow:
    0 0 0 1px rgba(0, 0, 0, 0.6),
    0 0 90px rgba(245, 158, 11, 0.3),
    inset 0 0 60px rgba(146, 64, 14, 0.35);
}
.ac-gacha-card::before {
  content: '';
  position: absolute;
  inset: 7px;
  border: 1px solid rgba(253, 230, 138, 0.18);
  pointer-events: none;
}
.ac-gacha-corner {
  position: absolute;
  width: 22px;
  height: 22px;
  border-color: #fcd34d;
  border-style: solid;
  filter: drop-shadow(0 0 6px rgba(252, 211, 77, 0.6));
}
.ac-gacha-c1 { top: -1px; left: -1px; border-width: 2px 0 0 2px; }
.ac-gacha-c2 { top: -1px; right: -1px; border-width: 2px 2px 0 0; }
.ac-gacha-c3 { bottom: -1px; left: -1px; border-width: 0 0 2px 2px; }
.ac-gacha-c4 { bottom: -1px; right: -1px; border-width: 0 2px 2px 0; }
.ac-gacha-overline {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 6px;
  font-family: Consolas, monospace;
  font-size: 8.5px;
  letter-spacing: 0.34em;
  color: #fcd34d;
}
.ac-gacha-rate {
  margin-top: 6px;
  text-align: center;
  font-family: Consolas, monospace;
  font-size: 8px;
  letter-spacing: 0.2em;
  color: rgba(255, 255, 255, 0.28);
}
.ac-gacha-stars {
  display: flex;
  justify-content: center;
  gap: 7px;
  margin-top: 22px;
  color: #fbbf24;
  filter: drop-shadow(0 0 9px rgba(251, 191, 36, 0.65));
}
.ac-gacha-stars svg { animation: ac-gacha-pop 0.6s cubic-bezier(0.2, 1.6, 0.4, 1) both; }
.ac-gacha-stars svg:nth-child(2) { animation-delay: 0.08s; }
.ac-gacha-stars svg:nth-child(3) { animation-delay: 0.16s; }
.ac-gacha-stars svg:nth-child(4) { animation-delay: 0.24s; }
.ac-gacha-stars svg:nth-child(5) { animation-delay: 0.32s; }
.ac-gacha-name {
  margin-top: 18px;
  text-align: center;
  font-family: 'Songti SC', 'Noto Serif SC', serif;
  font-size: 43px;
  font-weight: 700;
  line-height: 1.08;
  letter-spacing: 0.04em;
  background: linear-gradient(180deg, #fffbe8 10%, #fcd34d 52%, #b45309 100%);
  -webkit-background-clip: text;
  background-clip: text;
  color: transparent;
  filter: drop-shadow(0 0 18px rgba(251, 191, 36, 0.35));
  overflow-wrap: anywhere;
}
.ac-gacha-title {
  margin-top: 9px;
  text-align: center;
  font-size: 12.5px;
  letter-spacing: 0.22em;
  color: rgba(253, 230, 138, 0.8);
}
.ac-gacha-ribbon {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 10px;
  margin-top: 24px;
  border-top: 1px solid rgba(253, 230, 138, 0.16);
  border-bottom: 1px solid rgba(253, 230, 138, 0.16);
  padding: 10px 6px;
  font-size: 11px;
  letter-spacing: 0.12em;
  color: rgba(255, 255, 255, 0.62);
}
.ac-gacha-ribbon i { width: 3px; height: 3px; border-radius: 50%; background: #fbbf24; }
.ac-gacha-element { color: #fcd34d; }
.ac-gacha-quote {
  margin-top: 22px;
  text-align: center;
  font-family: 'Songti SC', 'Noto Serif SC', serif;
  font-size: 14.5px;
  line-height: 1.9;
  color: rgba(255, 255, 255, 0.86);
  overflow-wrap: anywhere;
}
.ac-gacha-sign {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-top: 26px;
  border-top: 1px dashed rgba(253, 230, 138, 0.22);
  padding-top: 13px;
}
.ac-gacha-sign span {
  font-family: Consolas, monospace;
  font-size: 8px;
  letter-spacing: 0.24em;
  color: rgba(255, 255, 255, 0.3);
}
.ac-gacha-sign b {
  display: inline-flex;
  align-items: center;
  gap: 5px;
  font-family: Consolas, monospace;
  font-size: 11px;
  font-weight: 400;
  color: #fcd34d;
  overflow-wrap: anywhere;
}
.ac-gacha-foot {
  margin-top: 26px;
  text-align: center;
  animation: ac-gacha-fade 1s 0.6s both;
}
.ac-gacha-progress {
  height: 3px;
  width: 100%;
  background: rgba(253, 230, 138, 0.14);
  overflow: hidden;
}
.ac-gacha-progress span {
  display: block;
  width: 82%;
  height: 100%;
  background: linear-gradient(90deg, #92400e, #fcd34d);
  box-shadow: 0 0 12px rgba(252, 211, 77, 0.7);
  animation: ac-gacha-grow 1.4s 0.5s cubic-bezier(0.2, 0.8, 0.2, 1) both;
}
.ac-gacha-foot p {
  margin-top: 10px;
  font-size: 10px;
  letter-spacing: 0.08em;
  color: rgba(255, 255, 255, 0.32);
}
.ac-gacha-serial {
  font-family: Consolas, monospace;
  font-size: 7.5px !important;
  letter-spacing: 0.24em !important;
  color: rgba(253, 230, 138, 0.3) !important;
}
@keyframes ac-gacha-reveal {
  from { opacity: 0; transform: translateY(46px) scale(0.9); }
  to { opacity: 1; transform: none; }
}
@keyframes ac-gacha-fade { from { opacity: 0; } to { opacity: 1; } }
@keyframes ac-gacha-spin { to { transform: translateX(-50%) rotate(360deg); } }
@keyframes ac-gacha-dust { to { background-position: 47px 94px, -59px 121px; } }
@keyframes ac-gacha-pop {
  from { opacity: 0; transform: scale(0.2) rotate(-40deg); }
  to { opacity: 1; transform: none; }
}
@keyframes ac-gacha-grow { from { width: 0; } to { width: 82%; } }
@media (prefers-reduced-motion: reduce) {
  .ac-gacha-root *,
  .ac-gacha-root *::before,
  .ac-gacha-root *::after {
    animation: none !important;
    transition: none !important;
  }
}
</style>
