<template>
  <div class="ac-det-root min-h-screen overflow-hidden px-4 py-9 text-[#2d2116]">
    <div class="ac-det-desk"></div>
    <div class="ac-det-lamp"></div>

    <div class="relative z-10 mx-auto w-full max-w-md">
      <div class="ac-det-folder">
        <span class="ac-det-tab">CASE FILE · 案件卷宗</span>
        <Paperclip class="ac-det-clip" :size="42" aria-hidden="true" />

        <div class="ac-det-sheet">
          <span class="ac-det-stamp">机密<br /><i>CONFIDENTIAL</i></span>

          <header class="ac-det-head">
            <div class="ac-det-topline">
              <span class="ac-det-badge">
                <Search :size="11" aria-hidden="true" />
                第七侦探事务所
              </span>
              <span class="ac-det-serial">FILE NO. {{ nfcId || 'PREVIEW' }}</span>
            </div>
            <p class="ac-det-kicker">PERSON OF INTEREST · 调查对象</p>
            <h1 class="ac-det-name">{{ p.name || '未署名对象' }}</h1>
            <p class="ac-det-title">「{{ p.title || '身份待查' }}」</p>
          </header>

          <div class="ac-det-photo-wrap">
            <div class="ac-det-photo">
              <svg class="ac-det-scene" viewBox="0 0 320 180" preserveAspectRatio="xMidYMid slice" aria-hidden="true">
                <defs>
                  <linearGradient id="detSky" x1="0" y1="0" x2="0" y2="1">
                    <stop offset="0" stop-color="#494137" />
                    <stop offset="0.55" stop-color="#252119" />
                    <stop offset="1" stop-color="#100d09" />
                  </linearGradient>
                  <radialGradient id="detGlow" cx="0.5" cy="0.5" r="0.5">
                    <stop offset="0" stop-color="#ffe6a8" stop-opacity="0.9" />
                    <stop offset="0.45" stop-color="#e3b867" stop-opacity="0.34" />
                    <stop offset="1" stop-color="#e3b867" stop-opacity="0" />
                  </radialGradient>
                </defs>
                <rect width="320" height="180" fill="url(#detSky)" />
                <circle cx="236" cy="48" r="58" fill="url(#detGlow)" class="ac-det-halo" />
                <path d="M0 148 L58 128 L104 140 L162 122 L224 138 L272 126 L320 141 V180 H0 Z" fill="#17130f" />
                <path d="M244 42 V148 M232 148 H256" stroke="#0a0806" stroke-width="3" />
                <path d="M231 42 H257 L251 27 H237 Z" fill="#0b0907" />
                <circle cx="244" cy="46" r="8" fill="#ffe1a0" opacity="0.85" />
                <g fill="#14100c">
                  <path d="M184 180 V145 c0-16 12-27 27-27 h10 c15 0 27 11 27 27 v35 Z" />
                  <circle cx="216" cy="103" r="12" />
                  <path d="M198 96 h36 l-5-10 h-26 Z" />
                  <path d="M204 85 h24 l4 11 h-32 Z" />
                  <path d="M158 180 L206 180 L194 164 L168 164 Z" opacity="0.42" />
                </g>
                <g class="ac-det-rain" stroke="#eadfc8" stroke-width="1" opacity="0.24">
                  <path d="M26 0 L12 44" /><path d="M68 8 L54 52" /><path d="M108 -4 L94 40" />
                  <path d="M148 12 L134 56" /><path d="M192 0 L178 44" /><path d="M226 18 L212 62" />
                  <path d="M272 4 L258 48" /><path d="M306 22 L292 66" /><path d="M44 70 L30 114" />
                  <path d="M286 96 L272 140" /><path d="M118 80 L104 124" /><path d="M166 100 L152 144" />
                </g>
              </svg>
              <span class="ac-det-photo-no">EVIDENCE A · 03/17</span>
              <span class="ac-det-photo-caption">案发现场 · 23:47</span>
            </div>
            <span class="ac-det-tape ac-det-tape-a"></span>
            <span class="ac-det-tape ac-det-tape-b"></span>
          </div>

          <div class="ac-det-clue" aria-hidden="true">
            <span class="ac-det-pin"></span>
            <span class="ac-det-clue-line"></span>
            <span class="ac-det-clue-label">线索 A · 身份核对</span>
            <span class="ac-det-clue-line"></span>
            <span class="ac-det-pin ac-det-pin-b"></span>
          </div>

          <div class="ac-det-fields">
            <div class="ac-det-field">
              <span>出自作品 / SERIES</span>
              <b>{{ p.series || '来源不明' }}</b>
            </div>
            <div class="ac-det-field">
              <span>属性 / ABILITY</span>
              <b class="ac-det-element">{{ p.element || '能力不详' }}</b>
            </div>
          </div>

          <blockquote class="ac-det-quote">
            <p class="ac-det-quote-label">证词 / WITNESS STATEMENT</p>
            <p class="ac-det-quote-text">“{{ p.quote || '真相只有一个，但它从不自己开口。' }}”</p>
          </blockquote>

          <footer class="ac-det-foot">
            <div class="ac-det-evidence">
              <span>证物袋 · 编号</span>
              <b>NO. {{ nfcId || 'PREVIEW' }}</b>
            </div>
            <div class="ac-det-contact">
              <span>联络人 / SNS</span>
              <b>@{{ p.social || '未登记' }}</b>
            </div>
            <p class="ac-det-sign">承办人签名 <i></i><em>已归档</em></p>
          </footer>
        </div>

        <div class="ac-det-folder-edge">
          <span>CONFIDENTIAL · DO NOT COPY</span>
          <span>{{ nfcId || 'PREVIEW' }}</span>
        </div>
      </div>

      <p class="ac-det-caption">CASE CLOSED? · 真相埋在咖啡渍里</p>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { Paperclip, Search } from 'lucide-vue-next'

const props = defineProps(['payload', 'nfcId'])
const p = computed(() => props.payload || {})
</script>

<style scoped>
.ac-det-root {
  position: relative;
  background: radial-gradient(ellipse at 78% 4%, rgba(255, 198, 108, 0.22), transparent 40%), radial-gradient(ellipse at 12% 96%, rgba(0, 0, 0, 0.55), transparent 52%), repeating-linear-gradient(88deg, rgba(255, 255, 255, 0.018) 0 2px, transparent 2px 24px), linear-gradient(155deg, #2c2119 0%, #1b130e 55%, #100b08 100%);
  font-family: 'Courier New', Courier, 'Noto Sans Mono CJK SC', monospace;
}
.ac-det-desk {
  position: absolute; inset: 0;
  background: repeating-linear-gradient(91deg, transparent 0 34px, rgba(0, 0, 0, 0.18) 34px 35px), radial-gradient(circle at 50% 30%, transparent, rgba(0, 0, 0, 0.45) 88%);
  opacity: 0.7;
}
.ac-det-lamp {
  position: absolute; top: -160px; right: -90px; width: 360px; height: 360px; border-radius: 50%; background: radial-gradient(circle, rgba(255, 206, 129, 0.24), transparent 66%);
  animation: ac-det-lamp 6s ease-in-out infinite;
}
.ac-det-folder {
  position: relative; border-radius: 5px; padding: 16px 12px 14px;
  background: repeating-linear-gradient(3deg, rgba(87, 59, 30, 0.035) 0 1px, transparent 1px 9px), linear-gradient(145deg, #c9a879, #b58e5d 52%, #a37d4e);
  box-shadow: 0 34px 70px -26px rgba(0, 0, 0, 0.85), 0 3px 0 rgba(255, 233, 195, 0.2) inset, 0 -2px 0 rgba(73, 47, 22, 0.3) inset; transform: rotate(-0.45deg);
  animation: ac-det-settle 0.9s cubic-bezier(0.2, 0.9, 0.25, 1) both;
}
.ac-det-tab {
  position: absolute; top: -15px; left: 22px; border-radius: 4px 4px 0 0; padding: 5px 13px 4px; background: #b28a58; color: #3e2c1a; font-size: 7px; font-weight: 700;
  letter-spacing: 0.16em; box-shadow: 0 -4px 12px rgba(0, 0, 0, 0.16);
}
.ac-det-clip {
  position: absolute; top: -24px; right: 27px; z-index: 4; color: #d8c08f; filter: drop-shadow(0 3px 2px rgba(0, 0, 0, 0.35)); transform: rotate(12deg);
  animation: ac-det-clip 5s ease-in-out infinite;
}
.ac-det-sheet {
  position: relative; overflow: hidden; padding: 24px 18px 16px;
  background: radial-gradient(circle at 86% 8%, rgba(173, 124, 67, 0.16), transparent 17%), radial-gradient(circle at 8% 92%, rgba(119, 82, 40, 0.1), transparent 22%), repeating-linear-gradient(0deg, rgba(116, 85, 48, 0.035) 0 1px, transparent 1px 6px), #efe2c7;
  box-shadow: 0 2px 7px rgba(55, 36, 18, 0.35), inset 0 0 46px rgba(146, 107, 60, 0.14); border: 1px solid rgba(93, 66, 34, 0.22);
}
.ac-det-sheet::before { content: ''; position: absolute; inset: 8px; border: 1px solid rgba(112, 80, 42, 0.16); pointer-events: none; }
.ac-det-sheet::after {
  content: ''; position: absolute; right: 20px; top: 26px; width: 58px; height: 58px; border: 5px solid rgba(122, 83, 42, 0.07); border-radius: 50%; pointer-events: none;
}
.ac-det-stamp {
  position: absolute; top: 78px; right: 12px; z-index: 2; display: grid; place-items: center; width: 76px; padding: 7px 4px 6px; border: 2px double #9d2c22; color: #9d2c22;
  text-align: center; font-size: 16px; font-weight: 700; line-height: 1; letter-spacing: 0.18em; transform: rotate(10deg); opacity: 0.66;
  animation: ac-det-stamp 0.6s 0.55s cubic-bezier(0.16, 1.4, 0.3, 1) both;
}
.ac-det-stamp i { margin-top: 4px; font-size: 4.5px; font-style: normal; letter-spacing: 0.1em; }
.ac-det-topline { display: flex; align-items: center; justify-content: space-between; gap: 8px; border-bottom: 1px solid rgba(91, 63, 32, 0.28); padding-bottom: 7px; }
.ac-det-badge { display: inline-flex; align-items: center; gap: 5px; color: #6b4c2a; font-size: 7px; font-weight: 700; letter-spacing: 0.14em; }
.ac-det-serial { color: #8c6a42; font-size: 6.5px; letter-spacing: 0.08em; white-space: nowrap; }
.ac-det-kicker { margin-top: 15px; color: #96744a; font-size: 6.5px; letter-spacing: 0.2em; }
.ac-det-name {
  max-width: 78%; margin-top: 6px; color: #24180d; font-family: 'Courier New', 'Noto Sans Mono CJK SC', monospace; font-size: 37px; font-weight: 700; line-height: 1.02;
  letter-spacing: -0.03em; overflow-wrap: anywhere;
}
.ac-det-title { margin-top: 7px; color: #7d2b20; font-family: 'Kaiti SC', 'STKaiti', 'Songti SC', serif; font-size: 15px; letter-spacing: 0.08em; overflow-wrap: anywhere; }
.ac-det-photo-wrap { position: relative; margin-top: 21px; transform: rotate(1.1deg); }
.ac-det-photo { position: relative; overflow: hidden; border: 6px solid #f8f2e4; border-bottom-width: 25px; background: #17130e; box-shadow: 0 10px 22px -10px rgba(54, 36, 18, 0.6); }
.ac-det-scene { display: block; width: 100%; height: 156px; filter: sepia(0.28) contrast(1.08) saturate(0.7); }
.ac-det-photo::after {
  content: ''; position: absolute; inset: 0; background-image: radial-gradient(rgba(255, 244, 220, 0.16) 0.5px, transparent 0.5px); background-size: 3px 3px;
  mix-blend-mode: overlay; opacity: 0.45; pointer-events: none;
}
.ac-det-halo { animation: ac-det-halo 4.5s ease-in-out infinite; }
.ac-det-rain { animation: ac-det-rain 2.8s linear infinite; }
.ac-det-photo-no { position: absolute; right: 6px; bottom: -18px; color: #584331; font-size: 6px; letter-spacing: 0.12em; }
.ac-det-photo-caption { position: absolute; left: 7px; bottom: -18px; color: #7a6148; font-size: 6px; letter-spacing: 0.08em; }
.ac-det-tape {
  position: absolute; z-index: 3; width: 45px; height: 13px; background: rgba(241, 226, 193, 0.72); box-shadow: 0 1px 2px rgba(64, 43, 22, 0.18); backdrop-filter: blur(1px);
}
.ac-det-tape-a { top: -8px; left: 16px; transform: rotate(-7deg); }
.ac-det-tape-b { right: 19px; bottom: 8px; transform: rotate(6deg); }
.ac-det-clue { display: flex; align-items: center; gap: 5px; margin: 20px 0 14px; transform: rotate(-1.2deg); }
.ac-det-pin {
  flex: none; width: 7px; height: 7px; border: 1px solid rgba(90, 20, 14, 0.5); border-radius: 50%; background: radial-gradient(circle at 35% 30%, #e8b39b, #a52a1e 60%, #5e120d);
  box-shadow: 0 1px 2px rgba(44, 15, 10, 0.5);
}
.ac-det-pin-b { width: 5px; height: 5px; }
.ac-det-clue-line { flex: 1; height: 1px; background: linear-gradient(90deg, rgba(155, 45, 33, 0.8), rgba(155, 45, 33, 0.35)); box-shadow: 0 1px 0 rgba(255, 247, 227, 0.35); }
.ac-det-clue-label { flex: none; color: #8b3024; font-size: 6px; letter-spacing: 0.16em; }
.ac-det-fields { display: grid; grid-template-columns: 1fr 1fr; gap: 12px; border-top: 1px dashed rgba(103, 73, 38, 0.35); padding-top: 12px; }
.ac-det-field { min-width: 0; }
.ac-det-field span { display: block; color: #967448; font-size: 6px; letter-spacing: 0.14em; }
.ac-det-field b {
  display: block; margin-top: 5px; border-bottom: 1px dotted rgba(93, 64, 32, 0.42); padding-bottom: 4px; color: #392817; font-size: 11px; font-weight: 700; letter-spacing: 0.03em;
  overflow-wrap: anywhere;
}
.ac-det-field .ac-det-element { color: #8d2b20; }
.ac-det-quote {
  position: relative; margin-top: 17px; border-left: 2px solid rgba(145, 38, 28, 0.65);
  background: repeating-linear-gradient(0deg, transparent 0 20px, rgba(122, 91, 53, 0.13) 20px 21px), rgba(233, 220, 193, 0.45); padding: 10px 10px 9px 12px;
}
.ac-det-quote-label { color: #9a6d3f; font-size: 6px; letter-spacing: 0.18em; }
.ac-det-quote-text {
  margin-top: 7px; color: #352518; font-family: 'Kaiti SC', 'STKaiti', 'Songti SC', serif; font-size: 15px; line-height: 1.65; letter-spacing: 0.02em; overflow-wrap: anywhere;
}
.ac-det-foot { display: grid; grid-template-columns: 1fr 1fr; gap: 10px; margin-top: 17px; border-top: 1px solid rgba(101, 71, 37, 0.3); padding-top: 11px; }
.ac-det-foot span { display: block; color: #9a764c; font-size: 5.5px; letter-spacing: 0.12em; }
.ac-det-foot b { display: block; margin-top: 3px; color: #564027; font-size: 8.5px; font-weight: 700; letter-spacing: 0.03em; overflow-wrap: anywhere; }
.ac-det-evidence b { display: inline-block; border: 1px solid rgba(92, 63, 31, 0.4); padding: 2px 5px; background: rgba(255, 250, 237, 0.42); }
.ac-det-sign { grid-column: 1 / -1; display: flex; align-items: center; gap: 7px; color: #8b6a43; font-family: 'Kaiti SC', 'STKaiti', serif; font-size: 8px; letter-spacing: 0.08em; }
.ac-det-sign i { flex: 1; height: 1px; background: rgba(105, 76, 41, 0.35); }
.ac-det-sign em { color: #9d2c22; font-size: 7px; font-style: normal; transform: rotate(-4deg); }
.ac-det-folder-edge { display: flex; justify-content: space-between; padding: 9px 3px 0; color: rgba(58, 38, 17, 0.68); font-size: 5.5px; letter-spacing: 0.13em; }
.ac-det-caption { margin-top: 22px; text-align: center; color: rgba(237, 217, 182, 0.44); font-size: 7px; letter-spacing: 0.28em; animation: ac-det-fade 1.4s 0.7s both; }
@keyframes ac-det-settle {
  from { opacity: 0; transform: translateY(24px) rotate(-2.4deg); }
  to { opacity: 1; transform: translateY(0) rotate(-0.45deg); }
}
@keyframes ac-det-stamp {
  from { opacity: 0; transform: rotate(10deg) scale(1.6); }
  to { opacity: 0.66; transform: rotate(10deg) scale(1); }
}
@keyframes ac-det-clip {
  0%, 100% { transform: rotate(12deg) translateY(0); }
  50% { transform: rotate(9deg) translateY(-2px); }
}
@keyframes ac-det-lamp {
  0%, 100% { opacity: 0.8; }
  50% { opacity: 1; }
}
@keyframes ac-det-halo {
  0%, 100% { opacity: 0.72; }
  50% { opacity: 1; }
}
@keyframes ac-det-rain {
  from { transform: translate(0, -14px); }
  to { transform: translate(-8px, 24px); }
}
@keyframes ac-det-fade {
  from { opacity: 0; }
  to { opacity: 1; }
}
@media (prefers-reduced-motion: reduce) {
  .ac-det-root *, .ac-det-root *::before, .ac-det-root *::after { animation: none !important; transition: none !important; }
}
</style>
