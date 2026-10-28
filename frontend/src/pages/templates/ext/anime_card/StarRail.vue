<template>
  <div class="ac-rail-root relative min-h-screen overflow-hidden px-4 py-7 text-[#eef4ff]">
    <div class="ac-rail-stars" aria-hidden="true"></div>
    <div class="ac-rail-nebula" aria-hidden="true"></div>

    <svg class="ac-rail-chart" viewBox="0 0 420 420" aria-hidden="true">
      <defs>
        <radialGradient id="acRailNode" cx="50%" cy="50%" r="50%">
          <stop offset="0" stop-color="#fff8dd" />
          <stop offset="1" stop-color="#e9d8a6" stop-opacity="0" />
        </radialGradient>
      </defs>
      <g fill="none" stroke="#c4b184" stroke-opacity=".25">
        <ellipse cx="210" cy="210" rx="184" ry="78" transform="rotate(-26 210 210)" />
        <ellipse cx="210" cy="210" rx="148" ry="56" transform="rotate(31 210 210)" stroke-dasharray="2 7" />
        <path d="M24 289 81 247 128 260 171 176 222 203 269 132 306 167 390 112" />
        <path d="M56 98 108 123 141 83 191 112 252 71 318 95 369 57" stroke-opacity=".15" />
      </g>
      <g fill="#e9d8a6">
        <circle cx="24" cy="289" r="2.5" /><circle cx="81" cy="247" r="1.8" /><circle cx="128" cy="260" r="2.5" />
        <circle cx="171" cy="176" r="2" /><circle cx="222" cy="203" r="2.6" /><circle cx="269" cy="132" r="1.7" />
        <circle cx="306" cy="167" r="2.4" /><circle cx="390" cy="112" r="2" /><circle cx="252" cy="71" r="1.7" />
      </g>
      <circle cx="222" cy="203" r="20" fill="url(#acRailNode)" opacity=".7" />
      <circle cx="222" cy="203" r="2.8" fill="#fff8dd" />
    </svg>

    <span class="ac-rail-planet" aria-hidden="true"></span>

    <div class="relative z-10 mx-auto w-full max-w-md">
      <header class="ac-rail-top">
        <span class="ac-rail-brand">
          <svg viewBox="0 0 24 24" aria-hidden="true">
            <rect x="5.5" y="3.5" width="13" height="14" rx="4.5" fill="none" stroke="currentColor" stroke-width="1.4" />
            <path d="M9 8h6M9 11.5h6M7.5 21l2.5-4M16.5 21 14 17" fill="none" stroke="currentColor" stroke-width="1.4" stroke-linecap="round" />
            <circle cx="9" cy="14.2" r="1" fill="currentColor" /><circle cx="15" cy="14.2" r="1" fill="currentColor" />
          </svg>
          ASTRA RAIL
        </span>
        <span class="ac-rail-route-code">ORBITAL LINE · 07</span>
      </header>

      <main class="ac-rail-ticket">
        <div class="ac-rail-shimmer" aria-hidden="true"></div>
        <div class="ac-rail-head">
          <div>
            <p class="ac-rail-overline">DEEP SPACE BOARDING PASS</p>
            <h2>星际列车 <span>STAR RAIL</span></h2>
          </div>
          <div class="ac-rail-stamp" aria-label="已检票">
            已检票
            <b>CHECKED</b>
          </div>
        </div>

        <div class="ac-rail-journey">
          <div class="ac-rail-station">
            <span>DEPARTURE</span>
            <b>地球站</b>
            <small>TERRA · 00:00</small>
          </div>
          <div class="ac-rail-track" aria-hidden="true">
            <i class="ac-rail-track-line"></i>
            <span class="ac-rail-track-dot ac-rail-dot-a"></span>
            <span class="ac-rail-track-dot ac-rail-dot-b"></span>
            <span class="ac-rail-track-dot ac-rail-dot-c"></span>
            <span class="ac-rail-train">
              <svg viewBox="0 0 24 24">
                <path d="M5 17 19 7M13 7h6v6" fill="none" stroke="currentColor" stroke-width="1.6" stroke-linecap="round" stroke-linejoin="round" />
              </svg>
            </span>
          </div>
          <div class="ac-rail-station ac-rail-station-end">
            <span>ARRIVAL</span>
            <b>星海站</b>
            <small>ASTRA · ∞</small>
          </div>
        </div>

        <section class="ac-rail-profile">
          <p class="ac-rail-label">PASSENGER / 乘客</p>
          <h1>{{ p.name || '未命名角色' }}</h1>
          <p class="ac-rail-title">「{{ p.title }}」</p>
        </section>

        <dl class="ac-rail-facts">
          <div class="ac-rail-fact-wide">
            <dt>出处作品 / ORIGIN</dt>
            <dd>{{ p.series || '未知作品' }}</dd>
          </div>
          <div>
            <dt>属性 / CLASS</dt>
            <dd class="ac-rail-element">{{ p.element }}</dd>
          </div>
          <div>
            <dt>座席 / SEAT</dt>
            <dd>07A · 星窗</dd>
          </div>
        </dl>

        <blockquote class="ac-rail-quote">
          <span aria-hidden="true">“</span>
          <p>{{ p.quote || '……下一站，是尚未命名的星海。' }}</p>
        </blockquote>

        <div class="ac-rail-stub">
          <div class="ac-rail-stub-copy">
            <p>CONTACT / 星际通讯</p>
            <b>@{{ p.social }}</b>
            <small>TICKET NO. {{ nfcId || 'PREVIEW' }} · FIRST LIGHT</small>
          </div>
          <div class="ac-rail-barcode" aria-hidden="true"><i v-for="n in 22" :key="n"></i></div>
          <p class="ac-rail-serial">{{ nfcId || 'PREVIEW' }}</p>
        </div>
      </main>

      <p class="ac-rail-caption">
        <span aria-hidden="true">✦</span>
        列车不停靠昨日 · 请携带好回忆
        <span aria-hidden="true">✦</span>
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
.ac-rail-root { background: radial-gradient(ellipse at 24% 6%, rgba(57, 84, 138, 0.28), transparent 38%), radial-gradient(ellipse at 82% 34%, rgba(175, 137, 67, 0.1), transparent 42%), linear-gradient(180deg, #030817 0%, #061127 56%, #030916 100%); font-family: 'PingFang SC', 'Microsoft YaHei', sans-serif; }
.ac-rail-stars { position: absolute; inset: 0; background-image: radial-gradient(rgba(255, 255, 255, 0.75) 0.6px, transparent 0.9px), radial-gradient(rgba(233, 216, 166, 0.55) 0.7px, transparent 1px), radial-gradient(rgba(147, 197, 253, 0.4) 0.5px, transparent 0.8px); background-size: 63px 63px, 101px 101px, 41px 41px; background-position: 10px 18px, 38px 7px, 18px 42px; opacity: 0.55; animation: ac-rail-twinkle 8s ease-in-out infinite; pointer-events: none; }
.ac-rail-nebula { position: absolute; top: 12%; right: -20%; width: 330px; height: 330px; border-radius: 50%; background: radial-gradient(circle, rgba(90, 119, 180, 0.18), transparent 65%); filter: blur(28px); pointer-events: none; }
.ac-rail-chart { position: absolute; top: 40px; left: 50%; width: 520px; transform: translateX(-50%); opacity: 0.58; animation: ac-rail-orbit 90s linear infinite; pointer-events: none; }
.ac-rail-planet { position: absolute; top: 125px; right: -30px; width: 92px; height: 92px; border-radius: 50%; border: 1px solid rgba(233, 216, 166, 0.16); background: radial-gradient(circle at 32% 28%, rgba(255, 250, 231, 0.36), transparent 22%), radial-gradient(circle at 68% 74%, rgba(15, 31, 61, 0.9), transparent 65%), linear-gradient(135deg, #6b7fa5, #1c2f50 55%, #0a1730); box-shadow: 0 0 40px rgba(147, 197, 253, 0.12); opacity: 0.55; pointer-events: none; }
.ac-rail-top { display: flex; align-items: center; justify-content: space-between; gap: 10px; animation: ac-rail-fade 0.9s both; }
.ac-rail-brand { display: inline-flex; align-items: center; gap: 7px; font-family: Georgia, 'Songti SC', serif; font-size: 11px; letter-spacing: 0.16em; color: #e9d8a6; }
.ac-rail-brand svg { width: 18px; height: 18px; filter: drop-shadow(0 0 5px rgba(233, 216, 166, 0.28)); }
.ac-rail-route-code { border: 1px solid rgba(233, 216, 166, 0.25); padding: 4px 7px; font-family: Consolas, monospace; font-size: 6.5px; letter-spacing: 0.15em; color: rgba(233, 216, 166, 0.6); white-space: nowrap; }
.ac-rail-ticket { position: relative; margin-top: 20px; padding: 20px 18px 0; border: 1px solid transparent; background: linear-gradient(155deg, rgba(8, 25, 49, 0.98), rgba(4, 14, 32, 0.99)) padding-box, linear-gradient(140deg, #f4e6bb, rgba(233, 216, 166, 0.16) 32%, #8f7138 66%, #ead79d) border-box; clip-path: polygon(0 16px, 16px 0, calc(100% - 16px) 0, 100% 16px, 100% calc(100% - 16px), calc(100% - 16px) 100%, 16px 100%, 0 calc(100% - 16px)); filter: drop-shadow(0 30px 45px rgba(0, 0, 0, 0.48)) drop-shadow(0 0 22px rgba(233, 216, 166, 0.07)); animation: ac-rail-reveal 1s cubic-bezier(0.16, 1, 0.3, 1) both; }
.ac-rail-ticket::before { content: ''; position: absolute; inset: 7px; border: 1px solid rgba(233, 216, 166, 0.12); pointer-events: none; }
.ac-rail-ticket::after { content: ''; position: absolute; inset: 0; background: radial-gradient(rgba(233, 216, 166, 0.16) 0.5px, transparent 0.7px); background-size: 17px 17px; opacity: 0.24; pointer-events: none; }
.ac-rail-shimmer { position: absolute; top: 0; left: -60%; z-index: 1; width: 45%; height: 2px; background: linear-gradient(90deg, transparent, #fff6d8, transparent); box-shadow: 0 0 12px rgba(255, 246, 216, 0.75); animation: ac-rail-shimmer 7.5s ease-in-out infinite; pointer-events: none; }
.ac-rail-head { position: relative; z-index: 2; display: flex; align-items: flex-start; justify-content: space-between; gap: 12px; border-bottom: 1px solid rgba(233, 216, 166, 0.18); padding-bottom: 14px; }
.ac-rail-overline { font-family: Consolas, monospace; font-size: 6.5px; letter-spacing: 0.22em; color: rgba(233, 216, 166, 0.5); }
.ac-rail-head h2 { margin-top: 6px; font-family: 'Songti SC', 'Noto Serif SC', serif; font-size: 19px; font-weight: 600; letter-spacing: 0.1em; color: #f6edd4; }
.ac-rail-head h2 span { display: block; margin-top: 2px; font-family: Georgia, serif; font-size: 8px; letter-spacing: 0.34em; color: #c6ae73; }
.ac-rail-stamp { flex: none; display: flex; flex-direction: column; align-items: center; justify-content: center; width: 57px; height: 57px; border: 1px solid rgba(233, 216, 166, 0.5); border-radius: 50%; font-size: 8px; letter-spacing: 0.12em; line-height: 1.4; color: #e9d8a6; transform: rotate(9deg); opacity: 0.82; }
.ac-rail-stamp b { font-family: Consolas, monospace; font-size: 5.5px; font-weight: 400; letter-spacing: 0.12em; }
.ac-rail-journey { position: relative; z-index: 2; display: flex; align-items: center; gap: 9px; padding: 16px 1px 13px; }
.ac-rail-station { flex: none; min-width: 56px; }
.ac-rail-station span { display: block; font-family: Consolas, monospace; font-size: 5.5px; letter-spacing: 0.1em; color: rgba(233, 216, 166, 0.42); }
.ac-rail-station b { display: block; margin-top: 2px; font-family: 'Songti SC', serif; font-size: 14px; font-weight: 600; letter-spacing: 0.06em; color: #f5edd9; }
.ac-rail-station small { display: block; margin-top: 2px; font-family: Consolas, monospace; font-size: 5.5px; letter-spacing: 0.06em; color: rgba(233, 216, 166, 0.4); }
.ac-rail-station-end { text-align: right; }
.ac-rail-track { position: relative; flex: 1; height: 20px; }
.ac-rail-track-line { position: absolute; top: 50%; left: 0; right: 0; height: 1px; background: repeating-linear-gradient(90deg, rgba(233, 216, 166, 0.55) 0 5px, transparent 5px 10px); }
.ac-rail-track-dot { position: absolute; top: 50%; width: 5px; height: 5px; border: 1px solid #e9d8a6; border-radius: 50%; background: #08182e; transform: translateY(-50%); }
.ac-rail-dot-a { left: 3%; opacity: 0.5; }
.ac-rail-dot-b { left: 48%; opacity: 0.8; }
.ac-rail-dot-c { right: 2%; opacity: 0.5; }
.ac-rail-train { position: absolute; top: 50%; left: 12%; display: flex; align-items: center; justify-content: center; width: 22px; height: 22px; border-radius: 50%; background: #102b4d; color: #f5e6bd; box-shadow: 0 0 0 3px rgba(233, 216, 166, 0.09), 0 0 14px rgba(233, 216, 166, 0.3); transform: translateY(-50%); animation: ac-rail-ride 8s ease-in-out infinite alternate; }
.ac-rail-train svg { width: 14px; height: 14px; }
.ac-rail-profile { position: relative; z-index: 2; border-top: 1px solid rgba(233, 216, 166, 0.16); padding-top: 14px; }
.ac-rail-label { font-family: Consolas, monospace; font-size: 6px; letter-spacing: 0.24em; color: rgba(233, 216, 166, 0.48); }
.ac-rail-profile h1 { margin-top: 5px; font-family: 'Songti SC', 'Noto Serif SC', serif; font-size: clamp(31px, 8.6vw, 41px); font-weight: 700; line-height: 1.1; letter-spacing: 0.03em; color: #fbf4e1; text-shadow: 0 0 24px rgba(233, 216, 166, 0.12); overflow-wrap: anywhere; }
.ac-rail-title { margin-top: 5px; font-size: 11.5px; letter-spacing: 0.14em; color: #cdb175; overflow-wrap: anywhere; }
.ac-rail-facts { position: relative; z-index: 2; display: grid; grid-template-columns: minmax(0, 1fr) minmax(0, 1fr); gap: 0; margin-top: 15px; border-top: 1px solid rgba(233, 216, 166, 0.14); border-bottom: 1px solid rgba(233, 216, 166, 0.14); }
.ac-rail-facts > div { min-width: 0; border-right: 1px solid rgba(233, 216, 166, 0.1); padding: 9px 8px 9px 0; }
.ac-rail-facts > div:nth-child(2) { border-right: 0; padding-left: 10px; }
.ac-rail-facts > div:nth-child(3) { border-right: 0; padding-left: 0; }
.ac-rail-fact-wide { grid-column: 1 / -1; border-right: 0 !important; border-bottom: 1px solid rgba(233, 216, 166, 0.1); padding-left: 0 !important; }
.ac-rail-facts dt { font-family: Consolas, monospace; font-size: 5.5px; letter-spacing: 0.14em; color: rgba(233, 216, 166, 0.38); }
.ac-rail-facts dd { margin-top: 3px; font-size: 10.5px; line-height: 1.5; color: rgba(238, 244, 255, 0.82); overflow-wrap: anywhere; }
.ac-rail-facts dd.ac-rail-element { color: #a7c7f5; }
.ac-rail-quote { position: relative; z-index: 2; margin-top: 13px; padding: 2px 0 13px 22px; }
.ac-rail-quote > span { position: absolute; top: -2px; left: 1px; font-family: Georgia, serif; font-size: 30px; line-height: 1; color: rgba(233, 216, 166, 0.32); }
.ac-rail-quote p { font-family: 'Songti SC', 'Noto Serif SC', serif; font-size: 13.5px; line-height: 1.85; color: rgba(244, 239, 223, 0.9); overflow-wrap: anywhere; }
.ac-rail-stub { position: relative; z-index: 2; display: flex; align-items: center; gap: 12px; margin: 0 -18px; border-top: 1px dashed rgba(233, 216, 166, 0.28); padding: 12px 18px 14px; background: rgba(1, 8, 20, 0.42); }
.ac-rail-stub-copy { min-width: 0; flex: 1; }
.ac-rail-stub-copy p { font-family: Consolas, monospace; font-size: 5.5px; letter-spacing: 0.18em; color: rgba(233, 216, 166, 0.4); }
.ac-rail-stub-copy b { display: block; margin-top: 3px; font-family: Consolas, monospace; font-size: 11px; font-weight: 400; color: #f0dcaa; overflow-wrap: anywhere; }
.ac-rail-stub-copy small { display: block; margin-top: 3px; font-family: Consolas, monospace; font-size: 5.5px; letter-spacing: 0.08em; color: rgba(238, 244, 255, 0.28); overflow-wrap: anywhere; }
.ac-rail-barcode { display: flex; flex: none; align-items: stretch; gap: 2px; width: 62px; height: 26px; opacity: 0.68; }
.ac-rail-barcode i { flex: 1; background: #e9d8a6; }
.ac-rail-barcode i:nth-child(2n) { flex: 0 0 1px; opacity: 0.45; }
.ac-rail-barcode i:nth-child(3n) { flex: 0 0 3px; background: #b7c9e9; }
.ac-rail-serial { position: absolute; right: 18px; bottom: 4px; font-family: Consolas, monospace; font-size: 4.5px; letter-spacing: 0.08em; color: rgba(233, 216, 166, 0.22); }
.ac-rail-caption { display: flex; align-items: center; justify-content: center; gap: 8px; margin-top: 20px; text-align: center; font-size: 8px; letter-spacing: 0.14em; color: rgba(233, 216, 166, 0.38); animation: ac-rail-fade 1.2s 0.5s both; }
.ac-rail-caption span { font-size: 7px; }
@keyframes ac-rail-reveal {
  from { opacity: 0; transform: translateY(32px) scale(0.97); }
  to { opacity: 1; transform: none; }
}
@keyframes ac-rail-fade {
  from { opacity: 0; }
  to { opacity: 1; }
}
@keyframes ac-rail-twinkle {
  0%, 100% { opacity: 0.45; }
  50% { opacity: 0.68; }
}
@keyframes ac-rail-orbit {
  to { transform: translateX(-50%) rotate(360deg); }
}
@keyframes ac-rail-shimmer {
  0%, 65% { left: -60%; opacity: 0; }
  72% { opacity: 1; }
  90% { left: 120%; opacity: 0; }
  100% { left: 120%; opacity: 0; }
}
@keyframes ac-rail-ride {
  from { left: 8%; }
  to { left: 68%; }
}
@media (prefers-reduced-motion: reduce) {
  .ac-rail-root *, .ac-rail-root *::before, .ac-rail-root *::after { animation: none !important; transition: none !important; }
}
</style>
