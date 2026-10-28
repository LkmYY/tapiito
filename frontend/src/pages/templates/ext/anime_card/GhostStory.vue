<template>
  <div class="ac-ghost-root min-h-screen overflow-hidden px-4 py-9 text-[#e9e1d4]">
    <div class="ac-ghost-sky"></div>
    <div class="ac-ghost-grain"></div>
    <span class="ac-ghost-fog ac-ghost-fog-a"></span>
    <span class="ac-ghost-fog ac-ghost-fog-b"></span>

    <div class="relative z-10 mx-auto w-full max-w-md">
      <div class="ac-ghost-candle" aria-hidden="true">
        <span class="ac-ghost-candle-glow"></span>
        <span class="ac-ghost-flame"></span>
        <span class="ac-ghost-wick"></span>
        <span class="ac-ghost-candle-body"></span>
        <span class="ac-ghost-wax ac-ghost-wax-a"></span>
        <span class="ac-ghost-wax ac-ghost-wax-b"></span>
        <span class="ac-ghost-candle-plate"></span>
      </div>

      <div class="ac-ghost-sheet">
        <span class="ac-ghost-ofuda" aria-hidden="true">
          <i>禁</i><i>忌</i>
          <em>{{ nfcId || 'PREVIEW' }}</em>
        </span>

        <header class="ac-ghost-head">
          <div>
            <p class="ac-ghost-kicker">KWAIDAN NOTES · 怪谈夜话</p>
            <p class="ac-ghost-night">第 {{ nfcId || '七' }} 夜 · 子时三刻</p>
          </div>
          <Moon class="ac-ghost-moon" :size="20" aria-hidden="true" />
        </header>

        <div class="ac-ghost-rule">
          <span></span>
          <b>今夜の怪異</b>
          <span></span>
        </div>

        <h1 class="ac-ghost-name">{{ p.name || '无名夜行者' }}</h1>
        <p class="ac-ghost-title">{{ p.title || '来历不明的存在' }}</p>

        <div class="ac-ghost-meta">
          <div>
            <span>出自作品</span>
            <b>{{ p.series || '无人知晓的旧谈' }}</b>
          </div>
          <div>
            <span>灵属性</span>
            <b class="ac-ghost-element">{{ p.element || '阴 · 未鉴定' }}</b>
          </div>
        </div>

        <blockquote class="ac-ghost-quote">
          <span class="ac-ghost-quote-mark" aria-hidden="true">“</span>
          <p class="ac-ghost-quote-label">独白 / WHISPER</p>
          <p class="ac-ghost-quote-text">{{ p.quote || '别吹灭那盏灯，我还没讲完我的故事。' }}</p>
          <span class="ac-ghost-quote-watermark" aria-hidden="true">怪</span>
        </blockquote>

        <div class="ac-ghost-warning">
          <span class="ac-ghost-talisman">禁</span>
          <p>子时之后，请勿翻至第七页。</p>
          <span class="ac-ghost-talisman ac-ghost-talisman-b">忌</span>
        </div>

        <footer class="ac-ghost-foot">
          <div class="ac-ghost-social">
            <span>聞き手 · 记录者</span>
            <b>@{{ p.social || '未留下姓名' }}</b>
          </div>
          <span class="ac-ghost-serial">ARCHIVE NO. {{ nfcId || 'PREVIEW' }}</span>
        </footer>
      </div>

      <p class="ac-ghost-caption">烛火灭时，请勿回头。</p>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { Moon } from 'lucide-vue-next'

const props = defineProps(['payload', 'nfcId'])
const p = computed(() => props.payload || {})
</script>

<style scoped>
.ac-ghost-root {
  position: relative;
  background: radial-gradient(ellipse at 50% 0%, rgba(122, 74, 128, 0.22), transparent 44%), radial-gradient(circle at 18% 72%, rgba(135, 37, 47, 0.1), transparent 32%), linear-gradient(175deg, #120d18 0%, #0b0710 48%, #070409 100%);
  font-family: 'Songti SC', 'Noto Serif SC', 'STSong', serif;
}
.ac-ghost-sky {
  position: absolute; inset: 0;
  background: radial-gradient(ellipse at 76% 12%, rgba(58, 44, 76, 0.55), transparent 55%), repeating-radial-gradient(circle at 30% 24%, transparent 0 42px, rgba(130, 100, 160, 0.026) 43px, transparent 45px);
}
.ac-ghost-grain {
  position: absolute; inset: -20px;
  background-image: radial-gradient(rgba(255, 244, 226, 0.22) 0.5px, transparent 0.5px), radial-gradient(rgba(0, 0, 0, 0.5) 0.7px, transparent 0.7px);
  background-position: 0 0, 2px 1px; background-size: 4px 4px, 5px 5px; opacity: 0.22; animation: ac-ghost-grain 0.9s steps(2) infinite; pointer-events: none;
}
.ac-ghost-fog {
  position: absolute; border-radius: 50%; background: radial-gradient(ellipse, rgba(159, 133, 177, 0.13), transparent 68%); filter: blur(26px);
  animation: ac-ghost-drift 15s ease-in-out infinite alternate;
}
.ac-ghost-fog-a { top: 30%; left: -100px; width: 280px; height: 180px; }
.ac-ghost-fog-b { right: -110px; bottom: 12%; width: 300px; height: 220px; animation-delay: -6s; opacity: 0.7; }
.ac-ghost-candle { position: relative; height: 92px; margin: 0 auto -6px; width: 120px; z-index: 2; }
.ac-ghost-candle-glow {
  position: absolute; left: 50%; top: -34px; width: 168px; height: 168px; transform: translateX(-50%); border-radius: 50%;
  background: radial-gradient(circle, rgba(255, 167, 79, 0.3), rgba(255, 133, 61, 0.08) 42%, transparent 70%); filter: blur(8px);
  animation: ac-ghost-breathe 3.4s ease-in-out infinite;
}
.ac-ghost-flame {
  position: absolute; left: 50%; top: 4px; width: 15px; height: 29px; transform: translateX(-50%) rotate(-2deg); border-radius: 52% 48% 46% 54% / 70% 68% 32% 30%;
  background: radial-gradient(ellipse at 50% 76%, #fff6d8 0 15%, #ffca76 40%, #f07733 72%, rgba(180, 52, 22, 0.2) 100%);
  box-shadow: 0 0 16px rgba(255, 171, 78, 0.75), 0 0 46px rgba(255, 128, 45, 0.3); animation: ac-ghost-flicker 2.1s ease-in-out infinite;
}
.ac-ghost-wick { position: absolute; left: 50%; top: 32px; width: 2px; height: 9px; transform: translateX(-50%); background: #2a1c12; }
.ac-ghost-candle-body {
  position: absolute; left: 50%; bottom: 9px; width: 37px; height: 49px; transform: translateX(-50%); border-radius: 3px 3px 7px 7px;
  background: linear-gradient(90deg, #b8a281 0%, #efe0c3 28%, #f7ecd6 48%, #cbb590 76%, #9e8666 100%);
  box-shadow: 0 10px 22px -8px rgba(0, 0, 0, 0.85), inset 0 2px 3px rgba(255, 249, 232, 0.8);
}
.ac-ghost-candle-body::after {
  content: ''; position: absolute; left: 8px; top: 2px; width: 5px; height: 38px; border-radius: 4px; background: linear-gradient(180deg, rgba(255, 251, 238, 0.8), transparent);
  opacity: 0.55;
}
.ac-ghost-wax { position: absolute; z-index: 2; width: 6px; border-radius: 3px; background: linear-gradient(180deg, #f4e7cf, #d5c3a4); }
.ac-ghost-wax-a { left: 47px; bottom: 44px; height: 19px; border-radius: 3px 3px 5px 5px; }
.ac-ghost-wax-b { right: 48px; bottom: 36px; height: 28px; width: 5px; opacity: 0.8; }
.ac-ghost-candle-plate {
  position: absolute; left: 50%; bottom: 0; width: 88px; height: 12px; transform: translateX(-50%); border-radius: 50%;
  background: radial-gradient(ellipse at 50% 30%, #6b5a52, #2b2224 62%, #120e12); box-shadow: 0 8px 18px rgba(0, 0, 0, 0.6);
}
.ac-ghost-sheet {
  position: relative; padding: 24px 20px 17px;
  background: radial-gradient(circle at 82% 12%, rgba(112, 72, 120, 0.2), transparent 30%), radial-gradient(circle at 10% 90%, rgba(66, 26, 52, 0.32), transparent 42%), repeating-linear-gradient(92deg, rgba(255, 255, 255, 0.013) 0 1px, transparent 1px 8px), linear-gradient(160deg, #1b1522 0%, #120d18 58%, #0d0912 100%);
  border: 1px solid rgba(180, 143, 192, 0.2); box-shadow: 0 34px 70px -24px rgba(0, 0, 0, 0.9), inset 0 0 60px rgba(0, 0, 0, 0.4), inset 0 1px 0 rgba(255, 224, 195, 0.09);
  clip-path: polygon(0% 0.5%, 99.2% 0%, 100% 99.4%, 0.7% 100%); animation: ac-ghost-rise 1s cubic-bezier(0.22, 1, 0.36, 1) both;
}
.ac-ghost-sheet::before { content: ''; position: absolute; inset: 9px; border: 1px solid rgba(176, 145, 188, 0.1); pointer-events: none; }
.ac-ghost-sheet::after {
  content: ''; position: absolute; left: 20px; right: 20px; bottom: 84px; height: 1px; background: linear-gradient(90deg, transparent, rgba(170, 130, 182, 0.25), transparent);
}
.ac-ghost-ofuda {
  position: absolute; top: -12px; right: 11px; z-index: 3; display: flex; flex-direction: column; align-items: center; width: 40px; min-height: 132px;
  border: 1.5px solid rgba(192, 45, 38, 0.85);
  background: repeating-linear-gradient(0deg, transparent 0 16px, rgba(150, 44, 37, 0.08) 16px 17px), linear-gradient(180deg, #efe4ce, #dccdb4); padding: 9px 0 7px; color: #b72620;
  box-shadow: 0 10px 20px -12px rgba(0, 0, 0, 0.9); transform-origin: top center; animation: ac-ghost-sway 5.5s ease-in-out infinite;
}
.ac-ghost-ofuda::before { content: ''; position: absolute; inset: 4px; border: 1px solid rgba(174, 38, 31, 0.4); }
.ac-ghost-ofuda i { font-family: 'Kaiti SC', 'STKaiti', serif; font-size: 19px; font-style: normal; font-weight: 700; line-height: 1.35; }
.ac-ghost-ofuda em {
  margin-top: 7px; font-family: Consolas, monospace; font-size: 5px; font-style: normal; letter-spacing: 0.08em; writing-mode: vertical-rl; color: rgba(126, 28, 23, 0.65);
  overflow-wrap: anywhere;
}
.ac-ghost-head { display: flex; align-items: flex-start; justify-content: space-between; gap: 10px; padding-right: 42px; }
.ac-ghost-kicker { color: #a48bab; font-family: Consolas, 'Courier New', monospace; font-size: 6.5px; letter-spacing: 0.22em; }
.ac-ghost-night { margin-top: 4px; color: #cbb9cf; font-size: 8.5px; letter-spacing: 0.14em; }
.ac-ghost-moon { flex: none; color: #cab18b; opacity: 0.8; filter: drop-shadow(0 0 6px rgba(224, 191, 134, 0.4)); }
.ac-ghost-rule { display: flex; align-items: center; gap: 8px; margin-top: 19px; color: #9e818e; }
.ac-ghost-rule span { flex: 1; height: 1px; background: linear-gradient(90deg, transparent, rgba(174, 135, 152, 0.36)); }
.ac-ghost-rule span:last-child { background: linear-gradient(90deg, rgba(174, 135, 152, 0.36), transparent); }
.ac-ghost-rule b { font-size: 8px; font-weight: 400; letter-spacing: 0.34em; }
.ac-ghost-name {
  margin-top: 10px; color: #f1e6d6; font-family: 'Kaiti SC', 'STKaiti', 'Songti SC', serif; font-size: 44px; font-weight: 700; line-height: 1.06; letter-spacing: 0.05em;
  text-shadow: 0 0 24px rgba(179, 126, 197, 0.22), 2px 2px 0 rgba(0, 0, 0, 0.5); overflow-wrap: anywhere;
}
.ac-ghost-title {
  display: inline-block; margin-top: 8px; border-bottom: 1px solid rgba(184, 48, 41, 0.55); padding-bottom: 3px; color: #d3a7a0; font-size: 12px; letter-spacing: 0.16em;
  overflow-wrap: anywhere;
}
.ac-ghost-meta {
  display: grid; grid-template-columns: 1fr 1fr; gap: 12px; margin-top: 20px; border-top: 1px solid rgba(158, 128, 168, 0.18); border-bottom: 1px solid rgba(158, 128, 168, 0.18);
  padding: 11px 0;
}
.ac-ghost-meta div { min-width: 0; }
.ac-ghost-meta span { display: block; color: #806d86; font-size: 6.5px; letter-spacing: 0.2em; }
.ac-ghost-meta b { display: block; margin-top: 5px; color: #ddd1c3; font-size: 11px; font-weight: 400; letter-spacing: 0.06em; overflow-wrap: anywhere; }
.ac-ghost-meta .ac-ghost-element { color: #d2645b; }
.ac-ghost-quote {
  position: relative; overflow: hidden; margin-top: 17px; border-left: 2px solid rgba(184, 48, 41, 0.7);
  background: linear-gradient(90deg, rgba(184, 48, 41, 0.08), transparent 55%), rgba(255, 255, 255, 0.025); padding: 12px 13px 14px 15px;
}
.ac-ghost-quote-mark { position: absolute; right: 4px; top: -18px; color: rgba(196, 161, 179, 0.14); font-family: Georgia, serif; font-size: 82px; line-height: 1; }
.ac-ghost-quote-label { position: relative; color: #927c99; font-family: Consolas, monospace; font-size: 6.5px; letter-spacing: 0.22em; }
.ac-ghost-quote-text {
  position: relative; margin-top: 8px; color: #e5d8c7; font-family: 'Kaiti SC', 'STKaiti', 'Songti SC', serif; font-size: 16px; line-height: 1.85; letter-spacing: 0.04em;
  overflow-wrap: anywhere;
}
.ac-ghost-quote-watermark {
  position: absolute; right: 10px; bottom: -21px; color: rgba(188, 143, 160, 0.055); font-family: 'Kaiti SC', 'STKaiti', serif; font-size: 78px; line-height: 1;
  pointer-events: none;
}
.ac-ghost-warning { display: flex; align-items: center; gap: 8px; margin-top: 16px; }
.ac-ghost-warning p { flex: 1; text-align: center; color: #9b7f88; font-size: 7.5px; letter-spacing: 0.16em; overflow-wrap: anywhere; }
.ac-ghost-talisman {
  display: grid; flex: none; place-items: center; width: 19px; height: 28px; border: 1px solid rgba(186, 49, 42, 0.8); background: rgba(239, 227, 207, 0.08); color: #c74a3f;
  font-family: 'Kaiti SC', 'STKaiti', serif; font-size: 11px; transform: rotate(-4deg);
}
.ac-ghost-talisman-b { transform: rotate(4deg); opacity: 0.82; }
.ac-ghost-foot { display: flex; align-items: flex-end; justify-content: space-between; gap: 12px; margin-top: 22px; padding-top: 12px; border-top: 1px solid rgba(158, 128, 168, 0.22); }
.ac-ghost-social { min-width: 0; }
.ac-ghost-social span { display: block; color: #806d86; font-size: 6.5px; letter-spacing: 0.18em; }
.ac-ghost-social b {
  display: block; margin-top: 4px; color: #d8c8be; font-family: Consolas, monospace; font-size: 10px; font-weight: 400; letter-spacing: 0.04em; overflow-wrap: anywhere;
}
.ac-ghost-serial { flex: none; color: #6f5d76; font-family: Consolas, monospace; font-size: 5.5px; letter-spacing: 0.12em; writing-mode: vertical-rl; text-orientation: mixed; }
.ac-ghost-caption { margin-top: 23px; text-align: center; color: rgba(173, 147, 183, 0.44); font-size: 8px; letter-spacing: 0.4em; animation: ac-ghost-fade 1.7s 0.8s both; }
@keyframes ac-ghost-rise {
  from { opacity: 0; transform: translateY(28px); filter: blur(3px); }
  to { opacity: 1; transform: none; filter: blur(0); }
}
@keyframes ac-ghost-flicker {
  0%, 100% { transform: translateX(-50%) rotate(-2deg) scale(1); opacity: 1; }
  18% { transform: translateX(-52%) rotate(2deg) scale(0.96, 1.03); opacity: 0.88; }
  34% { transform: translateX(-49%) rotate(-4deg) scale(1.03, 0.97); opacity: 1; }
  58% { transform: translateX(-51%) rotate(1deg) scale(0.98, 1.02); opacity: 0.92; }
  76% { transform: translateX(-50%) rotate(-1deg) scale(1.02, 0.99); opacity: 1; }
}
@keyframes ac-ghost-breathe {
  0%, 100% { opacity: 0.75; transform: translateX(-50%) scale(0.96); }
  50% { opacity: 1; transform: translateX(-50%) scale(1.06); }
}
@keyframes ac-ghost-drift {
  from { transform: translate3d(0, 0, 0) scale(1); }
  to { transform: translate3d(46px, -28px, 0) scale(1.14); }
}
@keyframes ac-ghost-grain {
  0% { transform: translate(0, 0); }
  25% { transform: translate(-1px, 1px); }
  50% { transform: translate(1px, 0); }
  75% { transform: translate(0, -1px); }
  100% { transform: translate(-1px, 0); }
}
@keyframes ac-ghost-sway {
  0%, 100% { transform: rotate(-1.4deg); }
  50% { transform: rotate(1.8deg); }
}
@keyframes ac-ghost-fade {
  from { opacity: 0; }
  to { opacity: 1; }
}
@media (prefers-reduced-motion: reduce) {
  .ac-ghost-root *, .ac-ghost-root *::before, .ac-ghost-root *::after { animation: none !important; transition: none !important; }
}
</style>
