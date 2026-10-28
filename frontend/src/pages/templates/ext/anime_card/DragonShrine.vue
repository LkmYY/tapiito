<template>
  <div class="ac-dragon-root relative min-h-screen overflow-hidden px-4 py-7 text-[#f2e7cd]">
    <div class="ac-dragon-stone" aria-hidden="true"></div>
    <div class="ac-dragon-halo" aria-hidden="true"></div>
    <div class="ac-dragon-dust" aria-hidden="true"></div>

    <svg class="ac-dragon-cloud ac-dragon-cloud-a" viewBox="0 0 220 96" aria-hidden="true">
      <g fill="none" stroke="currentColor" stroke-width="1" stroke-linecap="round">
        <path d="M8 71c20 0 19-17 39-17 16 0 18 13 34 13 14 0 15-11 28-11 14 0 14 14 31 14 13 0 21-6 27-14" />
        <path d="M36 54c10 0 12-10 22-10 9 0 11 7 20 7M118 59c8-9 20-10 28-4 7 5 17 3 22-4" />
      </g>
    </svg>
    <svg class="ac-dragon-cloud ac-dragon-cloud-b" viewBox="0 0 220 96" aria-hidden="true">
      <g fill="none" stroke="currentColor" stroke-width="1" stroke-linecap="round">
        <path d="M8 71c20 0 19-17 39-17 16 0 18 13 34 13 14 0 15-11 28-11 14 0 14 14 31 14 13 0 21-6 27-14" />
        <path d="M36 54c10 0 12-10 22-10 9 0 11 7 20 7" />
      </g>
    </svg>

    <div class="relative z-10 mx-auto w-full max-w-md">
      <header class="ac-dragon-top">
        <span class="ac-dragon-shrine"><i></i> 龍宮神殿 / RYUGU SHRINE</span>
        <span class="ac-dragon-no">SHRINE NO. {{ nfcId || 'PREVIEW' }}</span>
      </header>

      <main class="ac-dragon-tablet">
        <span class="ac-dragon-corner ac-dragon-corner-a" aria-hidden="true"></span>
        <span class="ac-dragon-corner ac-dragon-corner-b" aria-hidden="true"></span>
        <span class="ac-dragon-corner ac-dragon-corner-c" aria-hidden="true"></span>
        <span class="ac-dragon-corner ac-dragon-corner-d" aria-hidden="true"></span>

        <svg class="ac-dragon-art" viewBox="0 0 340 440" aria-hidden="true">
          <g fill="none" stroke="#d8b45e" stroke-linecap="round">
            <path d="M303 64c-41-28-97-12-109 28-13 45 39 57 60 88 23 33 3 81-40 85-41 4-76-28-65-63 9-29 48-39 66-16 14 19 2 44-18 47-16 2-28-11-25-24" stroke-width="1.2" stroke-opacity=".33" />
            <path d="M303 64c-9-13-8-28 2-39 4 11 3 22 0 30M295 61c11-10 25-14 39-11-8 9-18 14-29 17" stroke-width="1.2" stroke-opacity=".34" />
            <path d="M254 72c17-11 35-10 50-1" stroke-width="1" stroke-opacity=".22" />
            <path d="M128 139c-17-10-33 0-35 15M102 158c-17 1-28 14-23 28M79 191c-18 4-27 19-20 34M168 199c22 12 25 45 3 59M160 277c-29-1-47 25-40 50M112 329c-22 8-31 33-20 52" stroke-width="1" stroke-opacity=".18" />
            <circle cx="272" cy="81" r="2" fill="#d8b45e" stroke="none" />
          </g>
        </svg>

        <div class="ac-dragon-crest">
          <div class="ac-dragon-emblem">
            <svg viewBox="0 0 120 120" aria-hidden="true">
              <defs>
                <linearGradient id="acDragonGold" x1="0" y1="0" x2="1" y2="1">
                  <stop offset="0" stop-color="#f7e6b6" />
                  <stop offset=".52" stop-color="#d8b45e" />
                  <stop offset="1" stop-color="#9b7434" />
                </linearGradient>
              </defs>
              <g fill="none" stroke="url(#acDragonGold)" stroke-width="1.6" stroke-linecap="round" stroke-linejoin="round">
                <path d="M86 30c-15-12-37-9-46 7-9 17 7 34 24 33 16-1 27 13 21 27-6 15-26 18-38 7-10-9-8-22 3-27 8-4 18 2 17 11-1 8-9 11-15 7" />
                <path d="M86 30c-4-7-2-14 4-19 2 6 0 11-2 15M85 29c8-6 17-8 25-4-5 5-12 8-19 9" />
                <path d="M66 32c7-5 15-4 21 1M48 40c-8-1-15 4-18 11 8 0 14-2 19-7M43 75c-7 2-12 8-12 15 7-2 11-6 14-12" />
                <path d="M67 52c3 3 3 7 1 10M55 61c3 2 5 5 4 9M72 84c2 3 1 7-1 9M46 92c2 3 5 4 8 3" stroke-width="1.1" opacity=".7" />
              </g>
              <circle cx="73" cy="38" r="1.8" fill="#f6e6b8" />
            </svg>
          </div>
          <p class="ac-dragon-overline">御 祭 札</p>
          <span class="ac-dragon-sub">ENSHRINED · 龍神之契</span>
        </div>

        <section class="ac-dragon-profile">
          <p class="ac-dragon-role">「{{ p.title }}」</p>
          <h1>{{ p.name || '未命名角色' }}</h1>
          <p class="ac-dragon-series">出典 · {{ p.series || '未知作品' }}</p>
        </section>

        <div class="ac-dragon-divider" aria-hidden="true"><span></span><i>◇</i><span></span></div>

        <blockquote class="ac-dragon-oracle">
          <div class="ac-dragon-oracle-head"><span>神 託</span><small>ORACLE</small></div>
          <p>{{ p.quote || '……龙吟穿过深海，你的愿望已经听见。' }}</p>
          <span class="ac-dragon-seal" aria-hidden="true">龍</span>
        </blockquote>

        <div class="ac-dragon-details">
          <div><span>御力 / ELEMENT</span><b>{{ p.element }}</b></div>
          <div><span>参拜联络 / SOCIAL</span><b class="ac-dragon-social">@{{ p.social }}</b></div>
        </div>

        <footer class="ac-dragon-foot">
          <span>奉 納</span>
          <span class="ac-dragon-number">御札 {{ nfcId || 'PREVIEW' }}</span>
          <span>龍 宮</span>
        </footer>
      </main>

      <p class="ac-dragon-caption"><span>✦</span> 深海有龍 · 守願成真 <span>✦</span></p>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps(['payload', 'nfcId'])
const p = computed(() => props.payload || {})
</script>

<style scoped>
.ac-dragon-root { background: radial-gradient(circle at 18% 16%, rgba(37, 112, 100, 0.26), transparent 34%), radial-gradient(circle at 82% 72%, rgba(11, 69, 63, 0.5), transparent 42%), linear-gradient(180deg, #021b19 0%, #042b27 54%, #021713 100%); font-family: 'PingFang SC', 'Hiragino Sans GB', 'Microsoft YaHei', sans-serif; }
.ac-dragon-stone { position: absolute; inset: 0; background-image: linear-gradient(90deg, rgba(205, 178, 113, 0.045) 1px, transparent 1px), linear-gradient(rgba(205, 178, 113, 0.035) 1px, transparent 1px); background-size: 78px 54px; mask-image: radial-gradient(ellipse at center, #000, transparent 88%); pointer-events: none; }
.ac-dragon-halo { position: absolute; top: 4%; left: 50%; width: 360px; height: 360px; border-radius: 50%; background: radial-gradient(circle, rgba(213, 177, 94, 0.11), transparent 66%); filter: blur(20px); transform: translateX(-50%); pointer-events: none; }
.ac-dragon-dust { position: absolute; inset: 0; background-image: radial-gradient(rgba(227, 197, 123, 0.42) 0.6px, transparent 0.9px), radial-gradient(rgba(154, 205, 189, 0.35) 0.5px, transparent 0.8px); background-size: 61px 61px, 97px 97px; background-position: 9px 16px, 31px 45px; opacity: 0.42; animation: ac-dragon-drift 18s linear infinite; pointer-events: none; }
.ac-dragon-cloud { position: absolute; width: 230px; color: rgba(202, 173, 106, 0.28); pointer-events: none; }
.ac-dragon-cloud-a { top: 96px; left: -90px; }
.ac-dragon-cloud-b { right: -100px; bottom: 120px; transform: rotate(180deg); opacity: 0.65; }
.ac-dragon-top { display: flex; align-items: center; justify-content: space-between; gap: 10px; font-family: Consolas, 'Courier New', monospace; font-size: 6.5px; letter-spacing: 0.16em; color: rgba(226, 199, 133, 0.6); animation: ac-dragon-fade 0.9s both; }
.ac-dragon-shrine { display: inline-flex; align-items: center; gap: 6px; white-space: nowrap; }
.ac-dragon-shrine i { width: 5px; height: 5px; background: #d5b15e; transform: rotate(45deg); box-shadow: 0 0 8px rgba(213, 177, 94, 0.65); animation: ac-dragon-pulse 3s ease-in-out infinite; }
.ac-dragon-no { color: rgba(226, 199, 133, 0.36); white-space: nowrap; }
.ac-dragon-tablet { position: relative; overflow: hidden; margin-top: 20px; border: 1px solid rgba(205, 171, 95, 0.6); background: radial-gradient(circle at 50% 10%, rgba(87, 137, 121, 0.16), transparent 35%), linear-gradient(155deg, rgba(7, 51, 45, 0.97), rgba(3, 30, 27, 0.98) 62%, rgba(4, 37, 33, 0.99)); padding: 22px 20px 16px; box-shadow: 0 32px 65px -32px rgba(0, 0, 0, 0.9), 0 0 55px rgba(12, 84, 73, 0.14), inset 0 0 52px rgba(0, 0, 0, 0.3); animation: ac-dragon-reveal 1s cubic-bezier(0.16, 1, 0.3, 1) both; }
.ac-dragon-tablet::before { content: ''; position: absolute; inset: 7px; z-index: 1; border: 1px solid rgba(213, 177, 94, 0.16); pointer-events: none; }
.ac-dragon-tablet::after { content: ''; position: absolute; top: 0; left: 0; right: 0; height: 1px; z-index: 2; background: linear-gradient(90deg, transparent, #ecd7a3 32%, #a67e3d 68%, transparent); opacity: 0.75; }
.ac-dragon-corner { position: absolute; z-index: 3; width: 20px; height: 20px; border-color: rgba(220, 185, 109, 0.7); border-style: solid; pointer-events: none; }
.ac-dragon-corner-a { top: 13px; left: 13px; border-width: 1px 0 0 1px; }
.ac-dragon-corner-b { top: 13px; right: 13px; border-width: 1px 1px 0 0; }
.ac-dragon-corner-c { bottom: 13px; left: 13px; border-width: 0 0 1px 1px; }
.ac-dragon-corner-d { bottom: 13px; right: 13px; border-width: 0 1px 1px 0; }
.ac-dragon-art { position: absolute; top: 0; left: 50%; width: 340px; height: 440px; transform: translateX(-50%); opacity: 0.64; pointer-events: none; }
.ac-dragon-crest { position: relative; z-index: 3; display: flex; flex-direction: column; align-items: center; padding-top: 3px; }
.ac-dragon-emblem { position: relative; display: grid; place-items: center; width: 96px; height: 96px; border: 1px solid rgba(215, 180, 104, 0.36); border-radius: 50%; background: radial-gradient(circle, rgba(216, 180, 94, 0.08), transparent 68%); }
.ac-dragon-emblem::before { content: ''; position: absolute; inset: -15px; border-radius: 50%; background: repeating-conic-gradient(from 0deg, rgba(221, 189, 121, 0.32) 0deg 3deg, transparent 3deg 16deg); mask-image: radial-gradient(circle, transparent 45%, #000 47%, transparent 73%); animation: ac-dragon-spin 42s linear infinite; opacity: 0.52; }
.ac-dragon-emblem::after { content: ''; position: absolute; inset: 7px; border: 1px dashed rgba(215, 180, 104, 0.22); border-radius: 50%; }
.ac-dragon-emblem svg { position: relative; z-index: 1; width: 74px; height: 74px; filter: drop-shadow(0 0 9px rgba(215, 180, 104, 0.18)); }
.ac-dragon-overline { margin-top: 7px; font-family: 'Songti SC', 'Noto Serif SC', serif; font-size: 13px; letter-spacing: 0.4em; color: #e4cc96; }
.ac-dragon-sub { margin-top: 3px; font-family: Consolas, monospace; font-size: 5.5px; letter-spacing: 0.18em; color: rgba(213, 184, 117, 0.42); }
.ac-dragon-profile { position: relative; z-index: 3; padding-top: 16px; text-align: center; }
.ac-dragon-role { font-family: 'Kaiti SC', 'STKaiti', 'Songti SC', serif; font-size: 10px; letter-spacing: 0.14em; color: #c5a364; overflow-wrap: anywhere; }
.ac-dragon-profile h1 { margin-top: 6px; font-family: 'Songti SC', 'Noto Serif SC', Georgia, serif; font-size: clamp(32px, 9.2vw, 43px); font-weight: 700; line-height: 1.1; letter-spacing: 0.08em; color: #f3e5c2; text-shadow: 0 0 22px rgba(219, 182, 99, 0.16); overflow-wrap: anywhere; }
.ac-dragon-series { margin-top: 6px; font-size: 8.5px; letter-spacing: 0.1em; color: rgba(223, 204, 164, 0.5); overflow-wrap: anywhere; }
.ac-dragon-divider { display: flex; align-items: center; gap: 9px; position: relative; z-index: 3; margin: 14px 0 13px; color: rgba(216, 181, 101, 0.65); }
.ac-dragon-divider span { flex: 1; height: 1px; background: linear-gradient(90deg, transparent, rgba(211, 176, 98, 0.38)); }
.ac-dragon-divider span:last-child { background: linear-gradient(90deg, rgba(211, 176, 98, 0.38), transparent); }
.ac-dragon-divider i { font-style: normal; font-size: 8px; }
.ac-dragon-oracle { position: relative; z-index: 3; border: 1px solid rgba(179, 143, 72, 0.65); border-radius: 2px; background: radial-gradient(circle at 20% 10%, rgba(255, 251, 235, 0.7), transparent 28%), repeating-linear-gradient(2deg, transparent 0 6px, rgba(115, 88, 39, 0.025) 6px 7px), linear-gradient(150deg, #f1e6c9, #e7d5ab 62%, #efe0bf); padding: 15px 52px 16px 15px; box-shadow: 0 8px 20px -12px rgba(0, 0, 0, 0.7), inset 0 0 18px rgba(122, 89, 37, 0.12); }
.ac-dragon-oracle-head { display: flex; align-items: baseline; gap: 8px; border-bottom: 1px solid rgba(100, 70, 25, 0.2); padding-bottom: 5px; }
.ac-dragon-oracle-head span { font-family: 'Kaiti SC', 'STKaiti', serif; font-size: 10px; letter-spacing: 0.32em; color: #8a2d22; }
.ac-dragon-oracle-head small { font-family: Consolas, monospace; font-size: 5px; letter-spacing: 0.14em; color: rgba(95, 74, 35, 0.46); }
.ac-dragon-oracle p { margin-top: 8px; font-family: 'Kaiti SC', 'STKaiti', 'Songti SC', serif; font-size: 14.5px; line-height: 1.8; color: #203f37; overflow-wrap: anywhere; }
.ac-dragon-seal { position: absolute; right: 11px; bottom: 11px; display: flex; align-items: center; justify-content: center; width: 35px; height: 35px; border: 1.5px solid rgba(150, 42, 29, 0.75); border-radius: 2px; font-family: 'Kaiti SC', 'STKaiti', serif; font-size: 21px; line-height: 1; color: #9e2e22; transform: rotate(-6deg); opacity: 0.85; }
.ac-dragon-details { position: relative; z-index: 3; margin-top: 15px; border-top: 1px solid rgba(209, 175, 99, 0.17); }
.ac-dragon-details > div { display: flex; align-items: baseline; gap: 12px; border-bottom: 1px dotted rgba(209, 175, 99, 0.15); padding: 8px 1px; }
.ac-dragon-details span { flex: none; width: 118px; font-family: Consolas, monospace; font-size: 6px; letter-spacing: 0.06em; color: rgba(199, 169, 99, 0.55); }
.ac-dragon-details b { min-width: 0; flex: 1; text-align: right; font-size: 10.5px; font-weight: 500; line-height: 1.5; color: rgba(240, 229, 203, 0.84); overflow-wrap: anywhere; }
.ac-dragon-details b.ac-dragon-social { color: #dcc088; }
.ac-dragon-foot { position: relative; z-index: 3; display: flex; align-items: center; justify-content: space-between; gap: 10px; margin-top: 14px; border-top: 1px solid rgba(209, 175, 99, 0.18); padding-top: 11px; font-family: Consolas, monospace; font-size: 6px; letter-spacing: 0.16em; color: rgba(210, 179, 108, 0.46); }
.ac-dragon-number { color: rgba(227, 201, 141, 0.65); overflow-wrap: anywhere; text-align: center; }
.ac-dragon-caption { margin-top: 20px; text-align: center; font-family: 'Songti SC', 'Noto Serif SC', serif; font-size: 8.5px; letter-spacing: 0.22em; color: rgba(222, 198, 139, 0.44); animation: ac-dragon-fade 1.2s 0.5s both; }
.ac-dragon-caption span { color: rgba(213, 177, 94, 0.58); margin: 0 6px; }
@keyframes ac-dragon-reveal {
  from { opacity: 0; transform: translateY(32px) scale(0.98); filter: brightness(1.5); }
  to { opacity: 1; transform: none; filter: none; }
}
@keyframes ac-dragon-fade {
  from { opacity: 0; }
  to { opacity: 1; }
}
@keyframes ac-dragon-pulse {
  0%, 100% { opacity: 1; }
  50% { opacity: 0.32; }
}
@keyframes ac-dragon-spin {
  to { transform: rotate(360deg); }
}
@keyframes ac-dragon-drift {
  to { background-position: 70px 77px, -55px 142px; }
}
@media (prefers-reduced-motion: reduce) {
  .ac-dragon-root *, .ac-dragon-root *::before, .ac-dragon-root *::after { animation: none !important; transition: none !important; }
}
</style>
