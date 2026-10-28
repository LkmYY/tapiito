<template>
  <div class="ac-magic-root relative min-h-screen overflow-hidden px-4 py-7 text-[#4b245f]">
    <div class="ac-magic-sky" aria-hidden="true"></div>
    <div class="ac-magic-stars" aria-hidden="true"></div>
    <span class="ac-magic-orb ac-magic-orb-a" aria-hidden="true"></span>
    <span class="ac-magic-orb ac-magic-orb-b" aria-hidden="true"></span>
    <span class="ac-magic-glint ac-magic-glint-a" aria-hidden="true">✦</span>
    <span class="ac-magic-glint ac-magic-glint-b" aria-hidden="true">✧</span>
    <span class="ac-magic-glint ac-magic-glint-c" aria-hidden="true">✦</span>
    <span class="ac-magic-glint ac-magic-glint-d" aria-hidden="true">✧</span>

    <div class="relative z-10 mx-auto w-full max-w-md">
      <header class="ac-magic-top">
        <span class="ac-magic-archive"><i aria-hidden="true">✦</i> STARLIGHT ARCHIVE</span>
        <span class="ac-magic-no">NO. {{ nfcId || 'PREVIEW' }}</span>
      </header>

      <div class="ac-magic-stage">
        <div class="ac-magic-crest" aria-hidden="true">
          <svg viewBox="0 0 120 120">
            <defs>
              <linearGradient id="acMagicStar" x1="0" y1="0" x2="1" y2="1">
                <stop offset="0" stop-color="#fff4c9" />
                <stop offset=".38" stop-color="#ff9fd4" />
                <stop offset=".75" stop-color="#d69bff" />
                <stop offset="1" stop-color="#8cb9ff" />
              </linearGradient>
            </defs>
            <path d="M60 13 71.5 44.5 105 45.5 78.5 66 88 98 60 78.5 32 98 41.5 66 15 45.5 48.5 44.5Z" fill="url(#acMagicStar)" stroke="#fff" stroke-width="2.6" stroke-linejoin="round" />
            <path d="M60 30 65 45 81 46 69 55 73 70 60 61 47 70 51 55 39 46 55 45Z" fill="rgba(255,255,255,.4)" stroke="none" />
            <path d="M95 17 97.5 25 106 27.5 97.5 30 95 38 92.5 30 84 27.5 92.5 25Z" fill="#fff" opacity=".92" />
            <path d="M25 88 27 94 33 96 27 98 25 104 23 98 17 96 23 94Z" fill="#fff" opacity=".7" />
          </svg>
        </div>

        <main class="ac-magic-card">
          <div class="ac-magic-shine" aria-hidden="true"></div>
          <span class="ac-magic-corner ac-magic-corner-a" aria-hidden="true">✦</span>
          <span class="ac-magic-corner ac-magic-corner-b" aria-hidden="true">✦</span>

          <p class="ac-magic-kicker">MAGICAL TRANSFORMATION RECORD</p>
          <h1>{{ p.name || '未命名角色' }}</h1>
          <p class="ac-magic-title">「{{ p.title }}」</p>

          <div class="ac-magic-tags">
            <span>{{ p.series || '未知作品' }}</span>
            <span class="ac-magic-element">{{ p.element }}</span>
          </div>

          <div class="ac-magic-ribbon" aria-hidden="true"><span>✦ STARLIGHT CHANGE ✦</span></div>

          <blockquote class="ac-magic-spell">
            <span class="ac-magic-spell-label">变身咒语 / INCANTATION</span>
            <p>{{ p.quote || '以星光为誓，把平凡的一天变成奇迹。' }}</p>
            <i class="ac-magic-moon" aria-hidden="true">☾</i>
          </blockquote>

          <footer class="ac-magic-foot">
            <div>
              <span>魔法联络 / SOCIAL</span>
              <b>@{{ p.social }}</b>
            </div>
            <div class="ac-magic-id">
              <span>星之编号 / STAR ID</span>
              <b>{{ nfcId || 'PREVIEW' }}</b>
            </div>
          </footer>
        </main>
      </div>

      <p class="ac-magic-caption">
        <span aria-hidden="true">✧</span>
        星光会记得每一次勇敢的变身
        <span aria-hidden="true">✧</span>
      </p>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps(['payload', 'nfcId'])
const p = computed(() => props.payload || {})
</script>

<style scoped>
.ac-magic-root { background: radial-gradient(circle at 16% 10%, rgba(255, 255, 255, 0.9), transparent 27%), radial-gradient(circle at 86% 18%, rgba(255, 183, 227, 0.72), transparent 34%), radial-gradient(circle at 22% 82%, rgba(190, 202, 255, 0.7), transparent 35%), linear-gradient(155deg, #ffe4f4 0%, #f5dcff 42%, #dae8ff 100%); font-family: 'PingFang SC', 'Hiragino Sans GB', 'Microsoft YaHei', sans-serif; }
.ac-magic-sky { position: absolute; inset: 0; background: radial-gradient(ellipse at 70% 8%, rgba(255, 255, 255, 0.6), transparent 30%), radial-gradient(ellipse at 40% 100%, rgba(255, 255, 255, 0.4), transparent 35%); pointer-events: none; }
.ac-magic-stars { position: absolute; inset: 0; background-image: radial-gradient(circle, rgba(255, 255, 255, 0.95) 0.8px, transparent 1.2px), radial-gradient(circle, rgba(255, 249, 190, 0.9) 0.7px, transparent 1.1px), radial-gradient(circle, rgba(158, 125, 255, 0.42) 0.7px, transparent 1.1px); background-size: 57px 57px, 83px 83px, 101px 101px; background-position: 11px 22px, 31px 7px, 52px 42px; opacity: 0.62; animation: ac-magic-twinkle 7s ease-in-out infinite; pointer-events: none; }
.ac-magic-orb { position: absolute; border-radius: 50%; filter: blur(14px); pointer-events: none; }
.ac-magic-orb-a { top: 20%; right: -50px; width: 150px; height: 150px; background: rgba(255, 174, 222, 0.4); }
.ac-magic-orb-b { bottom: 12%; left: -55px; width: 170px; height: 170px; background: rgba(173, 171, 255, 0.35); }
.ac-magic-glint { position: absolute; z-index: 1; font-family: Georgia, serif; color: #fff; text-shadow: 0 0 9px rgba(255, 255, 255, 0.95), 0 0 18px rgba(255, 137, 205, 0.65); animation: ac-magic-sparkle 4s ease-in-out infinite; pointer-events: none; }
.ac-magic-glint-a { top: 18%; left: 7%; font-size: 15px; }
.ac-magic-glint-b { top: 40%; right: 7%; font-size: 20px; animation-delay: 0.9s; }
.ac-magic-glint-c { top: 73%; left: 5%; font-size: 12px; animation-delay: 1.8s; }
.ac-magic-glint-d { top: 87%; right: 10%; font-size: 16px; animation-delay: 2.7s; }
.ac-magic-top { display: flex; align-items: center; justify-content: space-between; gap: 10px; font-family: Consolas, 'Courier New', monospace; font-size: 6.5px; letter-spacing: 0.14em; color: rgba(94, 44, 122, 0.55); animation: ac-magic-fade 0.9s both; }
.ac-magic-archive { display: inline-flex; align-items: center; gap: 5px; white-space: nowrap; }
.ac-magic-archive i { font-style: normal; color: #d746a3; }
.ac-magic-no { border: 1px solid rgba(154, 91, 183, 0.28); border-radius: 999px; background: rgba(255, 255, 255, 0.38); padding: 4px 7px; white-space: nowrap; }
.ac-magic-stage { position: relative; margin-top: 74px; }
.ac-magic-crest { position: absolute; top: -63px; left: 50%; z-index: 5; display: grid; place-items: center; width: 128px; height: 128px; border: 3px solid rgba(255, 255, 255, 0.95); border-radius: 50%; background: radial-gradient(circle at 33% 26%, #fffdfd 0%, #ffe6f5 30%, #e8c8ff 65%, #b88cff 100%); box-shadow: 0 16px 32px -14px rgba(126, 54, 165, 0.62), 0 0 0 7px rgba(255, 255, 255, 0.28), 0 0 40px rgba(255, 255, 255, 0.6), inset 0 0 22px rgba(255, 255, 255, 0.7); transform: translateX(-50%); animation: ac-magic-float 5.5s ease-in-out infinite; }
.ac-magic-crest::before { content: ''; position: absolute; inset: -28px; z-index: -1; border-radius: 50%; background: repeating-conic-gradient(from 0deg, rgba(255, 255, 255, 0.92) 0deg 4deg, transparent 4deg 18deg); mask-image: radial-gradient(circle, transparent 46%, #000 48%, transparent 70%); animation: ac-magic-spin 38s linear infinite; opacity: 0.85; }
.ac-magic-crest::after { content: ''; position: absolute; inset: 9px; border: 1px dashed rgba(255, 255, 255, 0.8); border-radius: 50%; }
.ac-magic-crest svg { position: relative; z-index: 1; width: 86px; height: 86px; filter: drop-shadow(0 3px 5px rgba(144, 57, 164, 0.28)); }
.ac-magic-card { position: relative; overflow: hidden; border: 1px solid transparent; border-radius: 30px; background: radial-gradient(circle at 18% 0%, rgba(255, 255, 255, 0.94), transparent 27%), linear-gradient(160deg, #fffdfd 0%, #fff8fd 52%, #f4ecff 100%) padding-box, linear-gradient(135deg, #ffc2e3, #e3c4ff 45%, #c4dcff 72%, #ffd9f0) border-box; padding: 81px 20px 19px; box-shadow: 0 30px 60px -30px rgba(108, 53, 151, 0.5), 0 0 0 1px rgba(255, 255, 255, 0.5), inset 0 0 45px rgba(255, 226, 247, 0.4); animation: ac-magic-reveal 1s cubic-bezier(0.16, 1, 0.3, 1) both; }
.ac-magic-shine { position: absolute; top: -40%; left: -65%; width: 45%; height: 180%; background: linear-gradient(100deg, transparent 10%, rgba(255, 255, 255, 0.68) 50%, transparent 90%); transform: rotate(13deg); animation: ac-magic-shine 8s ease-in-out infinite; pointer-events: none; }
.ac-magic-corner { position: absolute; color: rgba(206, 102, 179, 0.4); font-family: Georgia, serif; font-size: 13px; animation: ac-magic-sparkle 4.4s ease-in-out infinite; pointer-events: none; }
.ac-magic-corner-a { top: 17px; left: 17px; }
.ac-magic-corner-b { top: 17px; right: 17px; animation-delay: 2s; }
.ac-magic-kicker { position: relative; z-index: 2; text-align: center; font-family: Consolas, monospace; font-size: 5.8px; letter-spacing: 0.2em; color: rgba(134, 74, 158, 0.55); }
.ac-magic-card h1 { position: relative; z-index: 2; margin-top: 7px; text-align: center; font-family: 'Songti SC', 'Noto Serif SC', Georgia, serif; font-size: clamp(32px, 9vw, 41px); font-weight: 700; line-height: 1.12; letter-spacing: 0.05em; color: #a51d74; background: linear-gradient(100deg, #a51d74 0%, #d83f9f 46%, #7c4ad4 100%); -webkit-background-clip: text; background-clip: text; -webkit-text-fill-color: transparent; filter: drop-shadow(0 2px 0 rgba(255, 255, 255, 0.85)); overflow-wrap: anywhere; }
.ac-magic-title { position: relative; z-index: 2; margin-top: 6px; text-align: center; font-family: 'Songti SC', 'Noto Serif SC', serif; font-size: 11.5px; letter-spacing: 0.1em; color: #9b5f8f; overflow-wrap: anywhere; }
.ac-magic-tags { position: relative; z-index: 2; display: flex; flex-wrap: wrap; align-items: center; justify-content: center; gap: 6px; margin-top: 12px; }
.ac-magic-tags span { max-width: 100%; border: 1px solid rgba(199, 136, 205, 0.4); border-radius: 999px; background: rgba(255, 255, 255, 0.6); padding: 4px 10px; font-size: 8.5px; line-height: 1.4; letter-spacing: 0.05em; color: #76537e; overflow-wrap: anywhere; }
.ac-magic-tags .ac-magic-element { border-color: rgba(163, 117, 232, 0.55); background: linear-gradient(135deg, rgba(255, 226, 246, 0.9), rgba(230, 218, 255, 0.9)); color: #8a3e9e; font-weight: 600; }
.ac-magic-ribbon { position: relative; z-index: 2; display: flex; align-items: center; justify-content: center; height: 26px; margin: 14px -27px 0; background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.38) 50%, transparent), linear-gradient(90deg, #ff86c8, #c197ff 52%, #82b9ff); clip-path: polygon(0 0, 100% 0, 97.5% 50%, 100% 100%, 0 100%, 2.5% 50%); box-shadow: 0 5px 12px rgba(171, 91, 203, 0.18); }
.ac-magic-ribbon span { font-family: Consolas, monospace; font-size: 6.5px; font-weight: 700; letter-spacing: 0.18em; color: #fff; text-shadow: 0 1px 4px rgba(102, 42, 133, 0.35); white-space: nowrap; }
.ac-magic-spell { position: relative; z-index: 2; margin-top: 16px; border: 1px dashed rgba(188, 130, 213, 0.5); border-radius: 18px; background: radial-gradient(circle at 90% 8%, rgba(255, 255, 255, 0.9), transparent 24%), linear-gradient(145deg, #fff3fb, #f2e9ff); padding: 18px 16px 15px; text-align: center; }
.ac-magic-spell::before { content: ''; position: absolute; inset: 5px; border: 1px solid rgba(255, 255, 255, 0.75); border-radius: 14px; pointer-events: none; }
.ac-magic-spell-label { position: relative; z-index: 1; display: block; font-family: Consolas, monospace; font-size: 5.8px; letter-spacing: 0.16em; color: #b06aa8; }
.ac-magic-spell p { position: relative; z-index: 1; margin-top: 7px; font-family: 'Kaiti SC', 'STKaiti', 'Songti SC', serif; font-size: 15px; line-height: 1.8; color: #5c2d69; overflow-wrap: anywhere; }
.ac-magic-moon { position: absolute; right: 12px; bottom: 7px; font-family: Georgia, serif; font-size: 17px; font-style: normal; color: rgba(159, 104, 196, 0.4); }
.ac-magic-foot { position: relative; z-index: 2; display: grid; grid-template-columns: minmax(0, 1.25fr) minmax(0, 0.75fr); gap: 12px; margin-top: 15px; border-top: 1px solid rgba(195, 151, 208, 0.32); padding-top: 11px; }
.ac-magic-foot span { display: block; font-family: Consolas, monospace; font-size: 5.5px; letter-spacing: 0.08em; color: rgba(126, 77, 142, 0.52); }
.ac-magic-foot b { display: block; margin-top: 3px; font-family: Consolas, 'PingFang SC', monospace; font-size: 9.5px; font-weight: 500; line-height: 1.5; color: #76408a; overflow-wrap: anywhere; }
.ac-magic-id { text-align: right; }
.ac-magic-caption { display: flex; align-items: center; justify-content: center; gap: 8px; margin-top: 23px; text-align: center; font-size: 8px; letter-spacing: 0.16em; color: rgba(102, 55, 126, 0.5); animation: ac-magic-fade 1.2s 0.5s both; }
.ac-magic-caption span { color: rgba(218, 71, 163, 0.6); }
@keyframes ac-magic-reveal {
  from { opacity: 0; transform: translateY(34px) scale(0.96); filter: brightness(1.25); }
  to { opacity: 1; transform: none; filter: none; }
}
@keyframes ac-magic-fade {
  from { opacity: 0; }
  to { opacity: 1; }
}
@keyframes ac-magic-float {
  0%, 100% { transform: translateX(-50%) translateY(0); }
  50% { transform: translateX(-50%) translateY(-5px); }
}
@keyframes ac-magic-spin {
  to { transform: rotate(360deg); }
}
@keyframes ac-magic-sparkle {
  0%, 100% { opacity: 0.45; transform: scale(0.82) rotate(0deg); }
  50% { opacity: 1; transform: scale(1.12) rotate(16deg); }
}
@keyframes ac-magic-shine {
  0%, 62% { left: -65%; opacity: 0; }
  70% { opacity: 1; }
  92% { left: 130%; opacity: 0; }
  100% { left: 130%; opacity: 0; }
}
@keyframes ac-magic-twinkle {
  0%, 100% { opacity: 0.45; }
  50% { opacity: 0.75; }
}
@media (prefers-reduced-motion: reduce) {
  .ac-magic-root *, .ac-magic-root *::before, .ac-magic-root *::after { animation: none !important; transition: none !important; }
}
</style>
