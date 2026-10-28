<template>
  <div class="ac-neo-root relative min-h-screen overflow-hidden px-4 pb-7 pt-5 text-[#e6fbff]">
    <div class="ac-neo-stars" aria-hidden="true"></div>
    <div class="ac-neo-grid" aria-hidden="true"></div>
    <span class="ac-neo-orb ac-neo-orb-a" aria-hidden="true"></span>
    <span class="ac-neo-orb ac-neo-orb-b" aria-hidden="true"></span>

    <svg class="ac-neo-skyline" viewBox="0 0 480 190" preserveAspectRatio="xMidYMax slice" aria-hidden="true">
      <defs>
        <linearGradient id="acNeoCity" x1="0" y1="0" x2="0" y2="1">
          <stop offset="0" stop-color="#071826" />
          <stop offset="1" stop-color="#02080f" />
        </linearGradient>
      </defs>
      <g fill="url(#acNeoCity)" stroke="#22d3ee" stroke-opacity=".18">
        <path d="M0 151h42V96h28v55h30V66h40v85h22v-38h34v38h40V86h34v65h30v-46h38v46h32V58h32v93h36v-38h32v38h30V98h32v53h30v-18h20v57H0z" />
      </g>
      <g fill="#67e8f9" opacity=".45">
        <rect x="53" y="77" width="5" height="4" /><rect x="66" y="77" width="5" height="4" />
        <rect x="53" y="91" width="5" height="4" /><rect x="80" y="91" width="5" height="4" />
        <rect x="148" y="80" width="5" height="4" /><rect x="161" y="80" width="5" height="4" />
        <rect x="148" y="98" width="5" height="4" /><rect x="175" y="98" width="5" height="4" />
        <rect x="244" y="101" width="5" height="4" /><rect x="257" y="101" width="5" height="4" />
        <rect x="310" y="84" width="5" height="4" /><rect x="323" y="84" width="5" height="4" />
        <rect x="365" y="73" width="4" height="4" /><rect x="379" y="73" width="4" height="4" />
        <rect x="415" y="110" width="5" height="4" /><rect x="428" y="110" width="5" height="4" />
      </g>
      <g fill="#f472b6" opacity=".5">
        <rect x="122" y="62" width="14" height="3" /><rect x="249" y="52" width="3" height="22" />
        <rect x="342" y="68" width="9" height="3" /><rect x="420" y="92" width="3" height="14" />
      </g>
      <path d="M78 66V34M96 66V43M406 58V27" stroke="#f472b6" stroke-opacity=".4" />
      <circle cx="78" cy="32" r="2" fill="#f472b6" /><circle cx="406" cy="25" r="2" fill="#f472b6" />
    </svg>

    <div class="relative z-10 mx-auto w-full max-w-md">
      <header class="ac-neo-top">
        <span class="ac-neo-live"><i></i> NEO-TOKYO / 24H SIGNAL</span>
        <span class="ac-neo-no">NO. {{ nfcId || 'PREVIEW' }}</span>
      </header>

      <main class="ac-neo-board">
        <span class="ac-neo-corner ac-neo-corner-a" aria-hidden="true"></span>
        <span class="ac-neo-corner ac-neo-corner-b" aria-hidden="true"></span>
        <span class="ac-neo-corner ac-neo-corner-c" aria-hidden="true"></span>
        <span class="ac-neo-corner ac-neo-corner-d" aria-hidden="true"></span>

        <aside class="ac-neo-rail" aria-hidden="true">
          <span class="ac-neo-rail-word">電脳</span>
          <span class="ac-neo-rail-sub">夜間営業</span>
          <span class="ac-neo-rail-line"></span>
          <span class="ac-neo-rail-tag">霓虹</span>
        </aside>

        <section class="ac-neo-panel">
          <div class="ac-neo-scan" aria-hidden="true"></div>
          <p class="ac-neo-file"><span>CHARACTER FILE</span> // 2099</p>
          <h1 class="ac-neo-name" :data-text="p.name || '未命名角色'">{{ p.name || '未命名角色' }}</h1>
          <p class="ac-neo-title">「{{ p.title }}」</p>

          <div class="ac-neo-rule"><i></i><span>SHINJUKU · 03:07</span><i></i></div>

          <blockquote class="ac-neo-quote">
            <span class="ac-neo-quote-mark" aria-hidden="true">“</span>
            <p>{{ p.quote || '……雨把霓虹冲成一片海，而我还站在灯下等你。' }}</p>
          </blockquote>

          <dl class="ac-neo-meta">
            <div><dt>作品 / SERIES</dt><dd>{{ p.series || '未知作品' }}</dd></div>
            <div><dt>属性 / ELEMENT</dt><dd class="ac-neo-elem">{{ p.element }}</dd></div>
            <div><dt>频道 / SOCIAL</dt><dd class="ac-neo-social">@{{ p.social }}</dd></div>
          </dl>
        </section>
      </main>

      <footer class="ac-neo-foot">
        <div class="ac-neo-barcode" aria-hidden="true"><i v-for="n in 26" :key="n"></i></div>
        <div class="ac-neo-foot-copy">
          <p>NFC / {{ nfcId || 'PREVIEW' }}</p>
          <p>NEON DISTRICT · 二〇九九</p>
        </div>
      </footer>

      <p class="ac-neo-ticker">NEON NEVER SLEEPS <span>◆</span> 東京 · 霓虹永夜 <span>◆</span> SIGNAL 98%</p>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps(['payload', 'nfcId'])
const p = computed(() => props.payload || {})
</script>

<style scoped>
.ac-neo-root { background: radial-gradient(ellipse at 12% 8%, rgba(244, 114, 182, 0.2), transparent 42%), radial-gradient(ellipse at 90% 24%, rgba(34, 211, 238, 0.18), transparent 45%), linear-gradient(180deg, #030b14 0%, #04131f 58%, #02080e 100%); font-family: 'PingFang SC', 'Hiragino Sans GB', 'Microsoft YaHei', sans-serif; }
.ac-neo-stars { position: absolute; inset: 0; background-image: radial-gradient(rgba(103, 232, 249, 0.55) 0.7px, transparent 0.9px), radial-gradient(rgba(244, 114, 182, 0.4) 0.6px, transparent 0.8px); background-size: 53px 53px, 89px 89px; background-position: 7px 12px, 31px 43px; opacity: 0.4; pointer-events: none; }
.ac-neo-grid { position: absolute; left: -28%; right: -28%; bottom: -190px; height: 470px; background-image: linear-gradient(rgba(34, 211, 238, 0.32) 1px, transparent 1px), linear-gradient(90deg, rgba(244, 114, 182, 0.25) 1px, transparent 1px); background-size: 34px 34px; transform: perspective(240px) rotateX(62deg); transform-origin: top center; mask-image: linear-gradient(to bottom, transparent, #000 28%, transparent 94%); opacity: 0.45; pointer-events: none; }
.ac-neo-orb { position: absolute; border-radius: 999px; filter: blur(54px); pointer-events: none; }
.ac-neo-orb-a { top: 18%; left: -80px; width: 210px; height: 210px; background: rgba(236, 72, 153, 0.2); }
.ac-neo-orb-b { top: 46%; right: -90px; width: 230px; height: 230px; background: rgba(6, 182, 212, 0.16); }
.ac-neo-skyline { position: absolute; bottom: 0; left: 0; width: 100%; height: 190px; opacity: 0.9; pointer-events: none; }
.ac-neo-top { display: flex; align-items: center; justify-content: space-between; gap: 10px; font-family: Consolas, 'Courier New', monospace; font-size: 7.5px; letter-spacing: 0.18em; color: rgba(103, 232, 249, 0.68); animation: ac-neo-fade 0.8s both; }
.ac-neo-live { display: inline-flex; align-items: center; gap: 6px; white-space: nowrap; }
.ac-neo-live i { width: 5px; height: 5px; border-radius: 50%; background: #f472b6; box-shadow: 0 0 9px #f472b6; animation: ac-neo-pulse 1.8s ease-in-out infinite; }
.ac-neo-no { color: rgba(255, 255, 255, 0.3); letter-spacing: 0.1em; white-space: nowrap; }
.ac-neo-board { position: relative; display: grid; grid-template-columns: 46px minmax(0, 1fr); margin-top: 18px; border: 1px solid rgba(34, 211, 238, 0.44); background: rgba(3, 12, 21, 0.9); box-shadow: 0 0 0 1px rgba(0, 0, 0, 0.7), 0 24px 70px -30px rgba(0, 0, 0, 0.95), 0 0 46px rgba(6, 182, 212, 0.1), inset 0 0 46px rgba(6, 182, 212, 0.05); animation: ac-neo-reveal 0.9s cubic-bezier(0.16, 1, 0.3, 1) both; }
.ac-neo-board::after { content: ''; position: absolute; top: 0; left: 0; right: 0; height: 1px; background: linear-gradient(90deg, transparent, #67e8f9 24%, #f472b6 68%, transparent); opacity: 0.85; }
.ac-neo-corner { position: absolute; z-index: 3; width: 17px; height: 17px; border-color: #67e8f9; border-style: solid; filter: drop-shadow(0 0 5px rgba(103, 232, 249, 0.55)); }
.ac-neo-corner-a { top: -1px; left: -1px; border-width: 2px 0 0 2px; }
.ac-neo-corner-b { top: -1px; right: -1px; border-width: 2px 2px 0 0; border-color: #f472b6; }
.ac-neo-corner-c { bottom: -1px; left: -1px; border-width: 0 0 2px 2px; border-color: #f472b6; }
.ac-neo-corner-d { bottom: -1px; right: -1px; border-width: 0 2px 2px 0; }
.ac-neo-rail { display: flex; flex-direction: column; align-items: center; gap: 13px; border-right: 1px solid rgba(34, 211, 238, 0.25); background: linear-gradient(180deg, rgba(244, 114, 182, 0.08), transparent 36%, rgba(34, 211, 238, 0.06)), #04101a; padding: 16px 0 14px; }
.ac-neo-rail-word { writing-mode: vertical-rl; text-orientation: upright; font-family: 'Hiragino Kaku Gothic ProN', 'Yu Gothic', sans-serif; font-size: 26px; font-weight: 800; letter-spacing: 0.12em; color: #8df3ff; text-shadow: 0 0 8px rgba(34, 211, 238, 0.8), 0 0 22px rgba(34, 211, 238, 0.45); }
.ac-neo-rail-sub { writing-mode: vertical-rl; font-family: 'Hiragino Kaku Gothic ProN', 'Yu Gothic', sans-serif; font-size: 8px; letter-spacing: 0.34em; color: rgba(244, 114, 182, 0.8); }
.ac-neo-rail-line { flex: 1; width: 1px; min-height: 52px; background: linear-gradient(180deg, rgba(103, 232, 249, 0.55), transparent); }
.ac-neo-rail-tag { border: 1px solid rgba(244, 114, 182, 0.45); padding: 3px 2px; writing-mode: vertical-rl; font-size: 7px; letter-spacing: 0.2em; color: #f9a8d4; }
.ac-neo-panel { position: relative; min-width: 0; overflow: hidden; padding: 20px 16px 18px; }
.ac-neo-scan { position: absolute; top: 0; left: 0; right: 0; height: 1px; background: rgba(103, 232, 249, 0.7); box-shadow: 0 0 12px rgba(103, 232, 249, 0.4); animation: ac-neo-scan 7s linear infinite; pointer-events: none; }
.ac-neo-file { font-family: Consolas, 'Courier New', monospace; font-size: 7px; letter-spacing: 0.2em; color: rgba(148, 163, 184, 0.65); }
.ac-neo-file span { color: #f472b6; }
.ac-neo-name { position: relative; z-index: 1; margin-top: 11px; font-family: 'Arial Black', 'PingFang SC', 'Microsoft YaHei', sans-serif; font-size: clamp(30px, 8.4vw, 39px); font-weight: 900; line-height: 1.05; letter-spacing: -0.045em; color: #f4feff; text-shadow: 0 0 18px rgba(34, 211, 238, 0.2); overflow-wrap: anywhere; }
.ac-neo-name::before, .ac-neo-name::after { content: attr(data-text); position: absolute; inset: 0; z-index: -1; font: inherit; letter-spacing: inherit; line-height: inherit; pointer-events: none; opacity: 0.42; }
.ac-neo-name::before { color: #f472b6; transform: translate(2px, -1px); clip-path: inset(16% 0 56% 0); animation: ac-neo-glitch-a 7s steps(1, end) infinite; }
.ac-neo-name::after { color: #22d3ee; transform: translate(-2px, 1px); clip-path: inset(58% 0 18% 0); animation: ac-neo-glitch-b 8.5s steps(1, end) infinite; }
.ac-neo-title { margin-top: 7px; font-family: 'Songti SC', 'Noto Serif SC', serif; font-size: 12px; letter-spacing: 0.16em; color: #f9a8d4; overflow-wrap: anywhere; }
.ac-neo-rule { display: flex; align-items: center; gap: 8px; margin: 15px 0 13px; font-family: Consolas, monospace; font-size: 6.5px; letter-spacing: 0.16em; color: rgba(103, 232, 249, 0.48); white-space: nowrap; }
.ac-neo-rule i { flex: 1; height: 1px; background: linear-gradient(90deg, transparent, rgba(34, 211, 238, 0.35)); }
.ac-neo-rule i:last-child { background: linear-gradient(90deg, rgba(34, 211, 238, 0.35), transparent); }
.ac-neo-quote { position: relative; border-left: 2px solid #f472b6; background: linear-gradient(90deg, rgba(244, 114, 182, 0.08), rgba(34, 211, 238, 0.03)); padding: 12px 12px 12px 13px; }
.ac-neo-quote-mark { position: absolute; top: -7px; right: 8px; font-family: Georgia, serif; font-size: 32px; line-height: 1; color: rgba(103, 232, 249, 0.22); }
.ac-neo-quote p { position: relative; font-size: 13px; line-height: 1.85; color: #d8f4f9; overflow-wrap: anywhere; }
.ac-neo-meta { margin-top: 16px; display: grid; gap: 0; }
.ac-neo-meta > div { display: flex; align-items: baseline; gap: 9px; border-bottom: 1px dotted rgba(103, 232, 249, 0.18); padding: 6px 1px; }
.ac-neo-meta > div:last-child { border-bottom: 0; }
.ac-neo-meta dt { flex: none; width: 88px; font-family: Consolas, 'Courier New', monospace; font-size: 6.3px; font-weight: 400; letter-spacing: 0.1em; color: rgba(103, 232, 249, 0.42); }
.ac-neo-meta dd { min-width: 0; flex: 1; text-align: right; font-size: 10.5px; line-height: 1.5; color: rgba(232, 249, 255, 0.85); overflow-wrap: anywhere; }
.ac-neo-meta dd.ac-neo-elem { color: #8df3ff; }
.ac-neo-meta dd.ac-neo-social { color: #f9a8d4; }
.ac-neo-foot { display: flex; align-items: center; gap: 14px; margin-top: 17px; animation: ac-neo-fade 1s 0.5s both; }
.ac-neo-barcode { display: flex; flex: 1; align-items: stretch; gap: 2px; height: 30px; opacity: 0.65; }
.ac-neo-barcode i { flex: 0 0 2px; background: #67e8f9; }
.ac-neo-barcode i:nth-child(3n) { flex-basis: 4px; background: #f472b6; }
.ac-neo-barcode i:nth-child(4n) { flex-basis: 1px; opacity: 0.55; }
.ac-neo-barcode i:nth-child(5n) { flex-basis: 3px; background: rgba(255, 255, 255, 0.7); }
.ac-neo-foot-copy { flex: none; text-align: right; font-family: Consolas, monospace; font-size: 6.3px; line-height: 1.8; letter-spacing: 0.13em; color: rgba(255, 255, 255, 0.29); }
.ac-neo-foot-copy p:first-child { color: rgba(103, 232, 249, 0.55); }
.ac-neo-ticker { margin-top: 15px; border-top: 1px solid rgba(34, 211, 238, 0.12); border-bottom: 1px solid rgba(34, 211, 238, 0.12); padding: 6px 0; text-align: center; font-family: Consolas, monospace; font-size: 6px; letter-spacing: 0.12em; color: rgba(103, 232, 249, 0.38); white-space: nowrap; animation: ac-neo-fade 1.2s 0.6s both; }
.ac-neo-ticker span { color: rgba(244, 114, 182, 0.55); margin: 0 5px; }
@keyframes ac-neo-reveal {
  from { opacity: 0; transform: translateY(28px) scale(0.97); filter: brightness(1.8); }
  to { opacity: 1; transform: none; filter: none; }
}
@keyframes ac-neo-fade {
  from { opacity: 0; }
  to { opacity: 1; }
}
@keyframes ac-neo-pulse {
  0%, 100% { opacity: 1; transform: scale(1); }
  50% { opacity: 0.35; transform: scale(0.8); }
}
@keyframes ac-neo-scan {
  0% { top: -2px; opacity: 0; }
  8% { opacity: 0.65; }
  72% { opacity: 0.25; }
  100% { top: 100%; opacity: 0; }
}
@keyframes ac-neo-glitch-a {
  0%, 91%, 100% { transform: translate(2px, -1px); }
  92% { transform: translate(-4px, 1px); }
  94% { transform: translate(3px, -2px); }
  96% { transform: translate(0, 0); }
}
@keyframes ac-neo-glitch-b {
  0%, 88%, 100% { transform: translate(-2px, 1px); }
  89% { transform: translate(4px, -1px); }
  92% { transform: translate(-3px, 2px); }
  94% { transform: translate(0, 0); }
}
@media (prefers-reduced-motion: reduce) {
  .ac-neo-root *, .ac-neo-root *::before, .ac-neo-root *::after { animation: none !important; transition: none !important; }
}
</style>
