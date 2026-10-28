<template>
  <div class="ac-omikuji-root relative min-h-screen overflow-hidden px-5 py-8 text-[#3a2018]">
    <div class="ac-omikuji-waves"></div>
    <svg class="ac-omikuji-torii" viewBox="0 0 240 120" aria-hidden="true">
      <g fill="none" stroke="#7f1d1d" stroke-width="7" stroke-linecap="square">
        <path d="M18 18 H222" />
        <path d="M30 36 H210" />
        <path d="M56 20 V112" />
        <path d="M184 20 V112" />
        <path d="M56 52 H184" />
      </g>
      <rect x="48" y="10" width="16" height="10" fill="#7f1d1d" />
      <rect x="176" y="10" width="16" height="10" fill="#7f1d1d" />
    </svg>

    <div class="relative z-10 mx-auto w-full max-w-md pt-16">
      <div class="ac-omikuji-slip">
        <span class="ac-omikuji-knot"></span>
        <div class="ac-omikuji-head">
          <span>星海神社</span>
          <span>御神籤 · 第 {{ nfcId || '拾' }} 番</span>
        </div>

        <p class="ac-omikuji-fortune-label">本日の運勢</p>
        <p class="ac-omikuji-fortune">大吉</p>
        <div class="ac-omikuji-divider"><i>✦</i></div>

        <section class="ac-omikuji-profile">
          <p class="ac-omikuji-role">御祭神 · {{ p.title }}</p>
          <h1 class="ac-omikuji-name">{{ p.name || '未命名角色' }}</h1>
          <p class="ac-omikuji-series">御由緒 · {{ p.series || '未知作品' }}</p>
          <span class="ac-omikuji-element">{{ p.element }}</span>
        </section>

        <blockquote class="ac-omikuji-quote">
          <span>御神託</span>
          <p>{{ p.quote || '……願い事は、星に預けて。' }}</p>
        </blockquote>

        <div class="ac-omikuji-advice">
          <div><span>吉方位</span><b>東南 · 水辺</b></div>
          <div><span>縁起物</span><b>星屑のお守り</b></div>
          <div><span>授与所</span><b>@{{ p.social }}</b></div>
        </div>

        <footer class="ac-omikuji-foot">
          <span class="ac-omikuji-stamp">星海<br />神社</span>
          <p>お持ち帰りください<br /><small>NFC {{ nfcId || '0000' }} · 授与済</small></p>
        </footer>
      </div>

      <p class="ac-omikuji-caption">厄 除 け · 星 詣 で</p>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps(['payload', 'nfcId'])
const p = computed(() => props.payload || {})
</script>

<style scoped>
.ac-omikuji-root {
  background: #7f1d1d;
  font-family: 'Songti SC', 'Noto Serif SC', serif;
}
.ac-omikuji-waves {
  position: absolute;
  inset: 0;
  background-image: repeating-radial-gradient(circle at 50% 110%, transparent 0 18px, rgba(255, 237, 213, 0.16) 18px 19px);
  background-size: 80px 80px;
  opacity: 0.8;
}
.ac-omikuji-root::after {
  content: '';
  position: absolute;
  inset: 0;
  background: radial-gradient(ellipse at 50% 30%, rgba(255, 255, 255, 0.12), transparent 55%),
    linear-gradient(180deg, rgba(69, 10, 10, 0.25), rgba(69, 10, 10, 0.55));
  pointer-events: none;
}
.ac-omikuji-torii {
  position: absolute;
  top: 26px;
  left: 50%;
  width: 210px;
  transform: translateX(-50%);
  opacity: 0.55;
}
.ac-omikuji-slip {
  position: relative;
  background:
    linear-gradient(180deg, rgba(255, 251, 235, 0.85), transparent 26%),
    #fbf3dd;
  padding: 22px 22px 18px;
  box-shadow:
    0 26px 60px -24px rgba(40, 4, 4, 0.8),
    inset 0 0 40px rgba(180, 130, 70, 0.15);
  animation: ac-omikuji-drop 1.1s cubic-bezier(0.2, 1, 0.3, 1) both;
  transform-origin: top center;
  transition: transform 0.5s ease, box-shadow 0.5s ease;
}
.ac-omikuji-slip:hover {
  transform: rotate(-0.7deg) translateY(-3px);
  box-shadow: 0 34px 70px -24px rgba(40, 4, 4, 0.9), inset 0 0 40px rgba(180, 130, 70, 0.15);
}
.ac-omikuji-slip::before {
  content: '';
  position: absolute;
  inset: 8px;
  border: 1px solid rgba(153, 27, 27, 0.22);
  pointer-events: none;
}
.ac-omikuji-knot {
  position: absolute;
  top: -17px;
  left: 50%;
  width: 34px;
  height: 22px;
  transform: translateX(-50%);
  border: 2px solid #b91c1c;
  border-radius: 3px;
  background: #fff8e7;
}
.ac-omikuji-knot::before,
.ac-omikuji-knot::after {
  content: '';
  position: absolute;
  top: 3px;
  width: 12px;
  height: 12px;
  border: 1.5px solid #b91c1c;
  border-radius: 50%;
  background: #fbf3dd;
}
.ac-omikuji-knot::before { left: -9px; }
.ac-omikuji-knot::after { right: -9px; }
.ac-omikuji-head {
  display: flex;
  justify-content: space-between;
  border-bottom: 1px solid rgba(153, 27, 27, 0.25);
  padding-bottom: 8px;
  color: #9a3412;
  font-size: 8px;
  letter-spacing: 0.12em;
}
.ac-omikuji-fortune-label {
  margin-top: 18px;
  text-align: center;
  font-size: 9px;
  letter-spacing: 0.44em;
  color: #a16207;
}
.ac-omikuji-fortune {
  margin-top: 4px;
  text-align: center;
  font-family: 'Kaiti SC', 'STKaiti', serif;
  font-size: 54px;
  font-weight: 800;
  line-height: 1.1;
  letter-spacing: 0.16em;
  color: #991b1b;
  text-shadow: 2px 2px 0 rgba(252, 211, 77, 0.55);
}
.ac-omikuji-divider {
  display: flex;
  align-items: center;
  gap: 9px;
  margin: 12px 0 16px;
}
.ac-omikuji-divider::before,
.ac-omikuji-divider::after {
  content: '';
  flex: 1;
  height: 1px;
  background: linear-gradient(90deg, transparent, rgba(153, 27, 27, 0.32));
}
.ac-omikuji-divider::after { background: linear-gradient(90deg, rgba(153, 27, 27, 0.32), transparent); }
.ac-omikuji-divider i { font-style: normal; font-size: 8px; color: #b45309; }
.ac-omikuji-profile { text-align: center; }
.ac-omikuji-role { font-size: 10px; letter-spacing: 0.22em; color: #a16207; }
.ac-omikuji-name {
  margin-top: 6px;
  font-family: 'Kaiti SC', 'STKaiti', 'Songti SC', serif;
  font-size: 41px;
  font-weight: 700;
  line-height: 1.1;
  letter-spacing: 0.1em;
  color: #3f1d12;
  overflow-wrap: anywhere;
}
.ac-omikuji-series { margin-top: 4px; font-size: 9.5px; letter-spacing: 0.12em; color: #8a5a44; }
.ac-omikuji-element {
  display: inline-block;
  margin-top: 10px;
  border: 1px solid #991b1b;
  padding: 3px 12px;
  font-size: 10px;
  letter-spacing: 0.16em;
  color: #991b1b;
  background: rgba(185, 28, 28, 0.04);
}
.ac-omikuji-quote {
  position: relative;
  margin-top: 20px;
  border-top: 1px solid rgba(153, 27, 27, 0.2);
  border-bottom: 1px solid rgba(153, 27, 27, 0.2);
  background: rgba(185, 28, 28, 0.035);
  padding: 16px 12px 14px;
  text-align: center;
}
.ac-omikuji-quote > span {
  position: absolute;
  top: -7px;
  left: 50%;
  transform: translateX(-50%);
  background: #fbf3dd;
  padding: 0 9px;
  font-size: 8px;
  letter-spacing: 0.3em;
  color: #991b1b;
}
.ac-omikuji-quote p {
  font-family: 'Kaiti SC', 'STKaiti', serif;
  font-size: 17px;
  line-height: 1.85;
  color: #4a2a1c;
  overflow-wrap: anywhere;
}
.ac-omikuji-advice { margin-top: 16px; display: grid; gap: 4px; }
.ac-omikuji-advice > div {
  display: flex;
  align-items: baseline;
  gap: 10px;
  border-bottom: 1px dotted rgba(153, 27, 27, 0.22);
  padding: 5px 2px;
}
.ac-omikuji-advice span {
  flex: none;
  width: 48px;
  font-size: 8px;
  letter-spacing: 0.1em;
  color: #a16207;
}
.ac-omikuji-advice b {
  min-width: 0;
  flex: 1;
  text-align: right;
  font-family: 'Songti SC', serif;
  font-size: 11px;
  font-weight: 600;
  color: #5b3424;
  overflow-wrap: anywhere;
}
.ac-omikuji-foot {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 12px;
  margin-top: 18px;
}
.ac-omikuji-stamp {
  display: flex;
  width: 50px;
  height: 50px;
  align-items: center;
  justify-content: center;
  border: 2.5px solid #b91c1c;
  border-radius: 50%;
  color: #b91c1c;
  font-size: 9px;
  font-weight: 700;
  line-height: 1.3;
  letter-spacing: 0.1em;
  text-align: center;
  transform: rotate(-9deg);
  opacity: 0.82;
}
.ac-omikuji-foot p {
  text-align: right;
  font-size: 9px;
  letter-spacing: 0.1em;
  color: #8a5a44;
}
.ac-omikuji-foot small {
  font-family: Consolas, monospace;
  font-size: 6.5px;
  letter-spacing: 0.12em;
  color: #b08b73;
}
.ac-omikuji-caption {
  margin-top: 20px;
  text-align: center;
  font-size: 9px;
  letter-spacing: 0.5em;
  color: rgba(255, 237, 213, 0.55);
  animation: ac-omikuji-fade 1.3s 0.7s both;
}
@keyframes ac-omikuji-drop {
  from { opacity: 0; transform: translateY(-36px) rotate(2deg) scale(0.96); }
  60% { transform: translateY(5px) rotate(-0.6deg); }
  to { opacity: 1; transform: none; }
}
@keyframes ac-omikuji-fade { from { opacity: 0; } to { opacity: 1; } }
@media (prefers-reduced-motion: reduce) {
  .ac-omikuji-root *,
  .ac-omikuji-root *::before,
  .ac-omikuji-root *::after {
    animation: none !important;
    transition: none !important;
  }
}
</style>
