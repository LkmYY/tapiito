<template>
  <div class="ac-battle-root min-h-screen overflow-hidden px-4 py-8 text-[#f5e9d2]">
    <div class="ac-battle-bg"></div>
    <div class="ac-battle-floor"></div>
    <span class="ac-battle-aura"></span>

    <div class="relative z-10 mx-auto w-full max-w-md">
      <div class="ac-battle-stage">
        <div class="ac-battle-card">
          <div class="ac-battle-inner">
            <span class="ac-battle-holo" aria-hidden="true"></span>

            <header class="ac-battle-head">
              <span class="ac-battle-cost">
                <small>LV</small>
                <b>{{ cost }}</b>
              </span>
              <div class="ac-battle-heading">
                <p class="ac-battle-kicker">DUELIST ARCHIVE · 决斗档案</p>
                <h1 class="ac-battle-name">{{ p.name || '无名召唤兽' }}</h1>
              </div>
              <span class="ac-battle-rarity" aria-label="五星稀有度">
                <Star v-for="n in 5" :key="n" :size="6" fill="currentColor" aria-hidden="true" />
              </span>
            </header>

            <div class="ac-battle-art">
              <span class="ac-battle-art-glow"></span>
              <span class="ac-battle-art-grid"></span>

              <span class="ac-battle-ring ac-battle-ring-a"></span>
              <span class="ac-battle-ring ac-battle-ring-b"></span>
              <span class="ac-battle-ring ac-battle-ring-c"></span>

              <span class="ac-battle-sigil">{{ glyph }}</span>
              <span class="ac-battle-art-sub">SUMMONING SIGIL</span>

              <span class="ac-battle-set">SET {{ nfcId || 'PREVIEW' }}</span>
              <span class="ac-battle-series">{{ p.series || '未知卡包' }}</span>
              <span class="ac-battle-element">{{ p.element || '无属性' }}</span>
            </div>

            <div class="ac-battle-stats">
              <div class="ac-battle-stat">
                <span>ATK<small>攻击力</small></span>
                <b>{{ atk }}</b>
              </div>
              <div class="ac-battle-level">
                <span>RARITY</span>
                <b>UR</b>
                <i>ULTRA RARE</i>
              </div>
              <div class="ac-battle-stat ac-battle-def">
                <span>DEF<small>守备力</small></span>
                <b>{{ def }}</b>
              </div>
            </div>

            <div class="ac-battle-body">
              <p class="ac-battle-kind">【{{ p.element || '无属性' }}族 / 效果怪兽】</p>
              <p class="ac-battle-title">{{ p.title || '称号未定' }}</p>
              <p class="ac-battle-quote">「{{ p.quote || '当这张卡被召唤时，整个战场都会记住它的名字。' }}」</p>
              <div class="ac-battle-passive">
                <span>PASSIVE</span>
                <p>此卡在场时，己方场上所有卡片攻击力上升 500 点。</p>
              </div>
            </div>

            <footer class="ac-battle-foot">
              <span class="ac-battle-owner">
                <AtSign :size="10" aria-hidden="true" />
                所有者 {{ p.social || '未登记' }}
              </span>
              <span class="ac-battle-code">CARD NO. {{ nfcId || 'PREVIEW' }}</span>
            </footer>
          </div>

          <span class="ac-battle-corner ac-battle-corner-tl" aria-hidden="true"></span>
          <span class="ac-battle-corner ac-battle-corner-tr" aria-hidden="true"></span>
          <span class="ac-battle-corner ac-battle-corner-bl" aria-hidden="true"></span>
          <span class="ac-battle-corner ac-battle-corner-br" aria-hidden="true"></span>
        </div>
      </div>

      <p class="ac-battle-caption">TURN 01 · 召唤成功</p>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { AtSign, Star } from 'lucide-vue-next'

const props = defineProps(['payload', 'nfcId'])
const p = computed(() => props.payload || {})
const seed = computed(() => {
  const src = String(props.nfcId || p.value.name || 'PREVIEW')
  let total = 0
  for (let i = 0; i < src.length; i += 1) {
    total = (total * 33 + src.charCodeAt(i)) % 9973
  }
  return total
})
const glyph = computed(() => (p.value.element || '无').toString().trim().slice(0, 1) || '无')
const cost = computed(() => 5 + (seed.value % 5))
const atk = computed(() => (22 + (seed.value % 18)) * 100)
const def = computed(() => (16 + (seed.value % 14)) * 100)
</script>

<style scoped>
.ac-battle-root {
  position: relative;
  background: radial-gradient(ellipse at 50% 10%, rgba(154, 27, 26, 0.38), transparent 48%), radial-gradient(ellipse at 50% 98%, rgba(214, 157, 72, 0.12), transparent 42%), linear-gradient(180deg, #140707 0%, #0b0506 55%, #070304 100%);
  font-family: 'PingFang SC', 'Microsoft YaHei', sans-serif;
}
.ac-battle-bg {
  position: absolute; inset: 0;
  background: repeating-linear-gradient(90deg, transparent 0 39px, rgba(218, 166, 88, 0.035) 39px 40px), repeating-linear-gradient(0deg, transparent 0 39px, rgba(218, 166, 88, 0.028) 39px 40px);
  mask-image: radial-gradient(ellipse at 50% 46%, #000, transparent 72%); -webkit-mask-image: radial-gradient(ellipse at 50% 46%, #000, transparent 72%);
}
.ac-battle-floor {
  position: absolute; left: -20%; right: -20%; bottom: -15%; height: 50%;
  background: repeating-conic-gradient(from 90deg at 50% 100%, rgba(200, 145, 63, 0.07) 0deg 1deg, transparent 1deg 12deg); transform: perspective(400px) rotateX(64deg);
  transform-origin: bottom center; opacity: 0.65;
}
.ac-battle-aura {
  position: absolute; top: 13%; left: 50%; width: 430px; height: 430px; transform: translateX(-50%); border-radius: 50%;
  background: radial-gradient(circle, rgba(234, 67, 48, 0.14), rgba(146, 25, 24, 0.06) 44%, transparent 70%); filter: blur(10px); animation: ac-battle-aura 5s ease-in-out infinite;
}
.ac-battle-stage { perspective: 900px; }
.ac-battle-card {
  position: relative; padding: 5px; border-radius: 13px;
  background: linear-gradient(135deg, #fff0bd 0%, #c48a3a 13%, #f4d88f 27%, #8a5a20 49%, #f2cc7a 66%, #a06a27 84%, #ffe9a8 100%);
  box-shadow: 0 38px 70px -24px rgba(0, 0, 0, 0.9), 0 0 36px rgba(226, 160, 70, 0.18), 0 1px 0 rgba(255, 247, 211, 0.8) inset; animation: ac-battle-float 6s ease-in-out infinite;
  transition: transform 0.5s ease;
}
.ac-battle-card:hover { transform: translateY(-5px) rotateX(1.5deg); }
.ac-battle-inner {
  position: relative; overflow: hidden; border: 1px solid rgba(255, 226, 159, 0.34); border-radius: 8px;
  background: radial-gradient(circle at 50% 0%, rgba(129, 36, 28, 0.2), transparent 38%), linear-gradient(170deg, #160b0b 0%, #0d0607 56%, #080304 100%);
}
.ac-battle-inner::before { content: ''; position: absolute; inset: 5px; border: 1px solid rgba(218, 166, 88, 0.14); border-radius: 5px; pointer-events: none; z-index: 6; }
.ac-battle-holo {
  position: absolute; inset: 0; z-index: 5;
  background: linear-gradient( 116deg, transparent 18%, rgba(255, 222, 145, 0.22) 30%, rgba(255, 108, 74, 0.12) 39%, rgba(141, 209, 255, 0.1) 48%, rgba(255, 215, 141, 0.18) 56%, transparent 69% );
  background-size: 280% 120%; mix-blend-mode: screen; pointer-events: none; animation: ac-battle-sheen 5.5s ease-in-out infinite;
}
.ac-battle-head {
  position: relative; z-index: 2; display: flex; align-items: center; gap: 9px; padding: 17px 14px 12px 15px; border-bottom: 1px solid rgba(212, 159, 77, 0.25);
  background: linear-gradient(90deg, rgba(130, 28, 24, 0.22), transparent 70%);
}
.ac-battle-cost {
  display: flex; flex: none; flex-direction: column; align-items: center; justify-content: center; width: 38px; height: 42px;
  clip-path: polygon(50% 0, 100% 23%, 100% 77%, 50% 100%, 0 77%, 0 23%); background: linear-gradient(160deg, #92251f, #4d0d0c); color: #ffe1a1;
  box-shadow: 0 0 0 1px rgba(229, 181, 100, 0.5) inset;
}
.ac-battle-cost small { font-family: Consolas, monospace; font-size: 5px; letter-spacing: 0.14em; opacity: 0.75; }
.ac-battle-cost b { font-family: Georgia, 'Times New Roman', serif; font-size: 18px; line-height: 1.05; }
.ac-battle-heading { flex: 1; min-width: 0; }
.ac-battle-kicker {
  color: #a88a60; font-family: Consolas, 'Courier New', monospace; font-size: 6px; letter-spacing: 0.16em; white-space: nowrap; overflow: hidden; text-overflow: ellipsis;
}
.ac-battle-name {
  margin-top: 3px; color: #fff1d2; font-family: 'Songti SC', 'Noto Serif SC', 'STSong', serif; font-size: 29px; font-weight: 700; line-height: 1.1; letter-spacing: 0.04em;
  text-shadow: 0 2px 10px rgba(0, 0, 0, 0.55); overflow-wrap: anywhere;
}
.ac-battle-rarity {
  display: flex; flex: none; align-self: flex-start; gap: 1px; margin-top: 4px; color: #f4c86f; filter: drop-shadow(0 0 3px rgba(244, 200, 111, 0.65));
  animation: ac-battle-stars 2.8s ease-in-out infinite;
}
.ac-battle-art {
  position: relative; overflow: hidden; height: 188px; border-top: 1px solid rgba(255, 220, 144, 0.18); border-bottom: 1px solid rgba(212, 159, 77, 0.3);
  background: radial-gradient(circle at 50% 52%, rgba(191, 49, 36, 0.42), transparent 38%), radial-gradient(circle at 50% 50%, rgba(245, 200, 118, 0.1), transparent 62%), linear-gradient(180deg, #1e0c0b, #0a0405);
}
.ac-battle-art-glow {
  position: absolute; left: 50%; top: 50%; width: 200px; height: 200px; transform: translate(-50%, -50%); border-radius: 50%;
  background: radial-gradient(circle, rgba(255, 199, 103, 0.2), rgba(171, 46, 33, 0.12) 44%, transparent 68%); filter: blur(6px);
  animation: ac-battle-glow 3.6s ease-in-out infinite;
}
.ac-battle-art-grid {
  position: absolute; inset: 0;
  background: repeating-linear-gradient(30deg, transparent 0 26px, rgba(230, 185, 108, 0.05) 26px 27px), repeating-linear-gradient(-30deg, transparent 0 26px, rgba(230, 185, 108, 0.04) 26px 27px);
  mask-image: radial-gradient(circle at 50% 50%, #000, transparent 74%); -webkit-mask-image: radial-gradient(circle at 50% 50%, #000, transparent 74%);
}
.ac-battle-ring { position: absolute; left: 50%; top: 50%; border-radius: 50%; border: 1px solid rgba(240, 197, 118, 0.45); transform: translate(-50%, -50%); }
.ac-battle-ring-a { width: 144px; height: 144px; border-style: dashed; animation: ac-battle-spin 16s linear infinite; }
.ac-battle-ring-b { width: 114px; height: 114px; border-color: rgba(196, 82, 52, 0.55); border-width: 2px 1px; transform: translate(-50%, -50%) rotate(24deg); }
.ac-battle-ring-c { width: 166px; height: 166px; border-color: rgba(240, 197, 118, 0.16); }
.ac-battle-ring-b::before, .ac-battle-ring-b::after { content: '✦'; position: absolute; left: 50%; top: -8px; color: #e8bc70; font-size: 9px; transform: translateX(-50%); }
.ac-battle-ring-b::after { top: auto; bottom: -8px; color: #b94431; }
.ac-battle-sigil {
  position: absolute; left: 50%; top: 50%; transform: translate(-50%, -52%); color: #f9dda4; font-family: 'Kaiti SC', 'STKaiti', 'Songti SC', serif; font-size: 66px;
  font-weight: 700; line-height: 1; text-shadow: 0 0 18px rgba(255, 198, 104, 0.55), 0 0 48px rgba(199, 55, 40, 0.4); animation: ac-battle-sigil 4s ease-in-out infinite;
}
.ac-battle-art-sub {
  position: absolute; left: 50%; top: 67%; transform: translateX(-50%); color: rgba(231, 190, 119, 0.4); font-family: Consolas, monospace; font-size: 5px; letter-spacing: 0.3em;
  white-space: nowrap;
}
.ac-battle-set { position: absolute; top: 10px; left: 11px; color: rgba(226, 184, 108, 0.55); font-family: Consolas, monospace; font-size: 5.5px; letter-spacing: 0.08em; }
.ac-battle-series {
  position: absolute; top: 10px; right: 11px; max-width: 55%; border: 1px solid rgba(223, 177, 98, 0.3); border-radius: 2px; padding: 2px 5px 1px; color: #d8bc8b; font-size: 6.5px;
  letter-spacing: 0.08em; text-align: right; overflow: hidden; text-overflow: ellipsis; white-space: nowrap;
}
.ac-battle-element { position: absolute; left: 11px; bottom: 10px; border-left: 2px solid #dba94f; padding-left: 6px; color: #f0c879; font-size: 8.5px; letter-spacing: 0.12em; }
.ac-battle-stats {
  display: flex; align-items: stretch; border-bottom: 1px solid rgba(212, 159, 77, 0.28);
  background: repeating-linear-gradient(90deg, transparent 0 3px, rgba(255, 255, 255, 0.016) 3px 4px), #0b0506;
}
.ac-battle-stat { flex: 1; min-width: 0; padding: 8px 13px 7px; border-left: 2px solid #a52d22; }
.ac-battle-stat span {
  display: flex; align-items: center; gap: 5px; color: #c2634e; font-family: Consolas, 'Courier New', monospace; font-size: 8px; font-weight: 700; letter-spacing: 0.14em;
}
.ac-battle-stat span small { color: #816e5c; font-family: 'PingFang SC', sans-serif; font-size: 5.5px; font-weight: 400; letter-spacing: 0.1em; }
.ac-battle-stat b {
  display: block; margin-top: 1px; color: #ffdf9d; font-family: 'Arial Black', Impact, Consolas, sans-serif; font-size: 25px; line-height: 1; letter-spacing: -0.02em;
  font-variant-numeric: tabular-nums; text-shadow: 0 0 14px rgba(255, 193, 92, 0.16);
}
.ac-battle-def { border-left: 0; border-right: 2px solid #a52d22; text-align: right; }
.ac-battle-def span { justify-content: flex-end; color: #9f927f; }
.ac-battle-def b { color: #e8d6b3; }
.ac-battle-level { flex: none; display: flex; flex-direction: column; align-items: center; justify-content: center; padding: 6px 13px; color: #9c7b45; }
.ac-battle-level span { font-family: Consolas, monospace; font-size: 5px; letter-spacing: 0.2em; }
.ac-battle-level b { margin-top: 2px; color: #f0c466; font-family: Georgia, serif; font-size: 16px; line-height: 1; text-shadow: 0 0 12px rgba(245, 199, 106, 0.4); }
.ac-battle-level i { margin-top: 3px; color: #6e5c47; font-family: Consolas, monospace; font-size: 4px; font-style: normal; letter-spacing: 0.08em; white-space: nowrap; }
.ac-battle-body {
  position: relative; padding: 12px 15px 10px;
  background: linear-gradient(180deg, rgba(131, 32, 26, 0.11), transparent 44%), repeating-linear-gradient(0deg, transparent 0 19px, rgba(255, 255, 255, 0.018) 19px 20px);
}
.ac-battle-kind { color: #a98d61; font-size: 6.5px; letter-spacing: 0.1em; }
.ac-battle-title {
  margin-top: 4px; color: #f0cc8a; font-family: 'Songti SC', 'Noto Serif SC', serif; font-size: 13px; font-weight: 700; letter-spacing: 0.08em; overflow-wrap: anywhere;
}
.ac-battle-quote {
  margin-top: 7px; color: #d8c7ad; font-family: 'Songti SC', 'Noto Serif SC', serif; font-size: 12.5px; line-height: 1.72; letter-spacing: 0.03em; overflow-wrap: anywhere;
}
.ac-battle-passive { display: flex; align-items: flex-start; gap: 8px; margin-top: 10px; border-top: 1px dashed rgba(209, 158, 77, 0.27); padding-top: 8px; }
.ac-battle-passive span {
  flex: none; border: 1px solid rgba(195, 62, 43, 0.6); padding: 2px 4px 1px; color: #d0654d; font-family: Consolas, monospace; font-size: 5px; letter-spacing: 0.12em;
}
.ac-battle-passive p { color: #8f7f6b; font-size: 8px; line-height: 1.6; letter-spacing: 0.03em; }
.ac-battle-foot {
  position: relative; z-index: 2; display: flex; align-items: center; justify-content: space-between; gap: 10px; border-top: 1px solid rgba(212, 159, 77, 0.26);
  padding: 8px 15px 9px; background: #0a0405;
}
.ac-battle-owner {
  display: flex; align-items: center; gap: 4px; min-width: 0; color: #b9a17b; font-family: Consolas, 'Courier New', monospace; font-size: 8px; overflow-wrap: anywhere;
}
.ac-battle-owner svg { flex: none; color: #cf9a4b; }
.ac-battle-code { flex: none; color: #77654f; font-family: Consolas, monospace; font-size: 5px; letter-spacing: 0.06em; }
.ac-battle-corner { position: absolute; z-index: 7; width: 17px; height: 17px; border-color: rgba(255, 239, 192, 0.9); pointer-events: none; }
.ac-battle-corner-tl { top: 3px; left: 3px; border-top: 2px solid; border-left: 2px solid; }
.ac-battle-corner-tr { top: 3px; right: 3px; border-top: 2px solid; border-right: 2px solid; }
.ac-battle-corner-bl { bottom: 3px; left: 3px; border-bottom: 2px solid; border-left: 2px solid; }
.ac-battle-corner-br { bottom: 3px; right: 3px; border-bottom: 2px solid; border-right: 2px solid; }
.ac-battle-caption {
  margin-top: 24px; text-align: center; color: rgba(229, 185, 111, 0.5); font-family: Consolas, 'Courier New', monospace; font-size: 7.5px; letter-spacing: 0.34em;
  animation: ac-battle-fade 1.4s 0.7s both;
}
@keyframes ac-battle-float {
  0%, 100% { transform: translateY(0); }
  50% { transform: translateY(-5px); }
}
@keyframes ac-battle-sheen {
  0%, 22% { background-position: 180% 0; }
  62%, 100% { background-position: -80% 0; }
}
@keyframes ac-battle-spin {
  from { transform: translate(-50%, -50%) rotate(0deg); }
  to { transform: translate(-50%, -50%) rotate(360deg); }
}
@keyframes ac-battle-glow {
  0%, 100% { opacity: 0.7; transform: translate(-50%, -50%) scale(0.95); }
  50% { opacity: 1; transform: translate(-50%, -50%) scale(1.07); }
}
@keyframes ac-battle-sigil {
  0%, 100% { text-shadow: 0 0 18px rgba(255, 198, 104, 0.45), 0 0 44px rgba(199, 55, 40, 0.3); }
  50% { text-shadow: 0 0 26px rgba(255, 210, 125, 0.8), 0 0 64px rgba(199, 55, 40, 0.5); }
}
@keyframes ac-battle-stars {
  0%, 100% { opacity: 0.85; }
  50% { opacity: 1; }
}
@keyframes ac-battle-aura {
  0%, 100% { opacity: 0.7; transform: translateX(-50%) scale(0.96); }
  50% { opacity: 1; transform: translateX(-50%) scale(1.05); }
}
@keyframes ac-battle-fade {
  from { opacity: 0; }
  to { opacity: 1; }
}
@media (prefers-reduced-motion: reduce) {
  .ac-battle-root *, .ac-battle-root *::before, .ac-battle-root *::after { animation: none !important; transition: none !important; }
}
</style>
