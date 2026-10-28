<template>
  <div class="pt-pixel-root min-h-screen px-4 py-7 text-lime-100">
    <div class="pt-pixel-content mx-auto w-full max-w-md">
      <div class="pt-pixel-hud">
        <span>PET.ROM v3.1</span>
        <span class="pt-pixel-coin">◈ × {{ nfcId || '0086' }}</span>
      </div>

      <article class="pt-pixel-console">
        <div class="pt-pixel-scanlines" aria-hidden="true"></div>

        <header class="pt-pixel-titlebar">
          <span class="pt-pixel-dot pt-pixel-dot-r"></span>
          <span class="pt-pixel-dot pt-pixel-dot-y"></span>
          <span class="pt-pixel-dot pt-pixel-dot-g"></span>
          <span class="pt-pixel-file">PLAYER_01.SAV</span>
        </header>

        <div class="pt-pixel-screen">
          <div class="pt-pixel-sprite-wrap">
            <span class="pt-pixel-shadow" aria-hidden="true"></span>
            <div class="pt-pixel-sprite" aria-hidden="true">
              <svg viewBox="0 0 16 16" shape-rendering="crispEdges">
                <g fill="#c084fc">
                  <rect x="4" y="0" width="2" height="1" /><rect x="10" y="0" width="2" height="1" />
                  <rect x="3" y="1" width="3" height="2" /><rect x="10" y="1" width="3" height="2" />
                  <rect x="2" y="3" width="12" height="6" />
                  <rect x="1" y="9" width="14" height="4" />
                  <rect x="2" y="13" width="3" height="2" /><rect x="11" y="13" width="3" height="2" />
                </g>
                <g fill="#2e1065">
                  <rect x="5" y="5" width="2" height="2" /><rect x="9" y="5" width="2" height="2" />
                  <rect x="7" y="8" width="2" height="1" />
                  <rect x="4" y="10" width="2" height="1" /><rect x="10" y="10" width="2" height="1" />
                </g>
                <g fill="#fda4af">
                  <rect x="2" y="4" width="1" height="2" /><rect x="13" y="4" width="1" height="2" />
                  <rect x="7" y="9" width="2" height="1" />
                </g>
                <g fill="#facc15">
                  <rect x="1" y="8" width="1" height="1" /><rect x="14" y="8" width="1" height="1" />
                </g>
              </svg>
            </div>
            <span class="pt-pixel-lvl">LV.{{ payload.age ? payload.age.replace(/[^\d]/g, '') || '03' : '03' }}</span>
          </div>

          <div class="pt-pixel-nameplate">
            <span class="pt-pixel-mark">▶</span>
            <h1>{{ (payload.petName || '未命名').toUpperCase() }}</h1>
            <span class="pt-pixel-cursor">_</span>
          </div>
          <p class="pt-pixel-class">{{ payload.breed }} · COMPANION CLASS</p>

          <div class="pt-pixel-bars">
            <div class="pt-pixel-bar-row">
              <span>TRUST</span>
              <div class="pt-pixel-bar"><i style="width: 92%"></i></div>
              <span>92</span>
            </div>
            <div class="pt-pixel-bar-row">
              <span>ENERGY</span>
              <div class="pt-pixel-bar pt-pixel-bar-alt"><i style="width: 78%"></i></div>
              <span>78</span>
            </div>
          </div>

          <div class="pt-pixel-stats">
            <div class="pt-pixel-stat">
              <span>OWNER</span>
              <strong>{{ payload.ownerName }}</strong>
            </div>
            <div class="pt-pixel-stat">
              <span>SPAWN</span>
              <strong>{{ payload.address }}</strong>
            </div>
          </div>

          <div class="pt-pixel-alert">
            <span class="pt-pixel-alert-tag">! MEDICAL</span>
            <p>{{ payload.medical || 'NO DATA' }}</p>
          </div>

          <div class="pt-pixel-dialog">
            <span class="pt-pixel-dialog-name">{{ payload.petName || 'PET' }}：</span>
            <p>{{ payload.note || '如果我在外面迷路了，请帮我联系主人。' }}</p>
            <span class="pt-pixel-next" aria-hidden="true">▼</span>
          </div>
        </div>

        <a class="pt-pixel-call" :href="'tel:' + payload.phone">
          <span class="pt-pixel-call-key">[ CALL OWNER ]</span>
          <strong>{{ payload.phone }}</strong>
          <span class="pt-pixel-call-hint">按此拨号 · 走失必联</span>
        </a>

        <footer class="pt-pixel-footer">
          <span>HP ██████████ 10/10</span>
          <span>SAVE &amp; TAP</span>
        </footer>
      </article>

      <p class="pt-pixel-press">— PRESS START TO CALL —</p>
    </div>
  </div>
</template>

<script setup>
defineProps(['payload', 'nfcId'])
</script>

<style scoped>
.pt-pixel-root {
  background:
    linear-gradient(rgba(168, 85, 247, 0.055) 1px, transparent 1px),
    linear-gradient(90deg, rgba(168, 85, 247, 0.055) 1px, transparent 1px),
    radial-gradient(circle at 50% 20%, #2e1065 0%, #170b2e 48%, #0b0518 100%);
  background-size: 18px 18px, 18px 18px, 100% 100%;
  font-family: 'Courier New', 'Lucida Console', monospace;
}
.pt-pixel-content { animation: pt-pixel-boot 0.75s steps(7, end) both; }
.pt-pixel-hud { display: flex; justify-content: space-between; margin: 0 3px 8px; font-size: 9px; letter-spacing: 0.14em; color: #a78bfa; }
.pt-pixel-coin { color: #facc15; }
.pt-pixel-console {
  position: relative;
  padding: 0 0 14px;
  background: #1b1035;
  border: 3px solid #a78bfa;
  box-shadow: 0 0 0 3px #0b0518, 0 0 0 6px #6d28d9, 0 16px 30px rgba(0, 0, 0, 0.6);
  transition: transform 0.2s steps(2, end);
}
.pt-pixel-console:hover { transform: translate(-2px, -2px); }
.pt-pixel-scanlines { position: absolute; inset: 0; z-index: 4; pointer-events: none; background: repeating-linear-gradient(0deg, rgba(0, 0, 0, 0.2) 0 2px, transparent 2px 4px); opacity: 0.55; }
.pt-pixel-titlebar { display: flex; align-items: center; gap: 5px; padding: 7px 9px; background: #a78bfa; color: #1b1035; }
.pt-pixel-dot { width: 7px; height: 7px; background: #1b1035; }
.pt-pixel-dot-r { background: #f43f5e; }
.pt-pixel-dot-y { background: #facc15; }
.pt-pixel-dot-g { background: #22c55e; }
.pt-pixel-file { margin-left: auto; font-size: 8px; font-weight: 700; letter-spacing: 0.08em; }
.pt-pixel-screen { padding: 17px 17px 12px; background: #120a26; }
.pt-pixel-sprite-wrap { position: relative; display: flex; align-items: flex-end; justify-content: center; height: 128px; border-bottom: 2px dashed rgba(167, 139, 250, 0.4); margin-bottom: 12px; }
.pt-pixel-sprite { width: 112px; height: 112px; filter: drop-shadow(0 0 9px rgba(192, 132, 252, 0.65)); animation: pt-pixel-idle 2.4s ease-in-out infinite; }
.pt-pixel-sprite svg { display: block; width: 100%; height: 100%; }
.pt-pixel-shadow { position: absolute; bottom: 2px; left: 50%; width: 86px; height: 9px; margin-left: -43px; background: rgba(192, 132, 252, 0.18); border-radius: 50%; filter: blur(3px); }
.pt-pixel-lvl { position: absolute; top: 2px; right: 0; padding: 2px 5px; border: 2px solid #facc15; color: #facc15; font-size: 8px; font-weight: 700; }
.pt-pixel-nameplate { display: flex; align-items: baseline; gap: 5px; }
.pt-pixel-mark { color: #facc15; font-size: 16px; animation: pt-pixel-blink 1.2s steps(2, end) infinite; }
.pt-pixel-nameplate h1 { margin: 0; font-size: clamp(28px, 8.5vw, 38px); line-height: 1.1; font-weight: 700; letter-spacing: 0.02em; color: #f5f3ff; text-shadow: 3px 3px 0 #6d28d9; }
.pt-pixel-cursor { color: #a3e635; animation: pt-pixel-blink 1s steps(2, end) infinite; }
.pt-pixel-class { margin: 5px 0 12px; font-size: 9px; letter-spacing: 0.1em; color: #a78bfa; }
.pt-pixel-bars { display: grid; gap: 5px; padding: 8px 9px; border: 2px solid #4c1d95; background: rgba(109, 40, 217, 0.13); }
.pt-pixel-bar-row { display: grid; grid-template-columns: 48px 1fr 24px; gap: 6px; align-items: center; font-size: 8px; color: #c4b5fd; }
.pt-pixel-bar { height: 8px; background: #2e1065; border: 1px solid #7c3aed; padding: 1px; }
.pt-pixel-bar i { display: block; height: 100%; background: repeating-linear-gradient(90deg, #a3e635 0 4px, #65a30d 4px 6px); }
.pt-pixel-bar-alt i { background: repeating-linear-gradient(90deg, #f472b6 0 4px, #be185d 4px 6px); }
.pt-pixel-stats { display: grid; grid-template-columns: 0.8fr 1.2fr; gap: 7px; margin-top: 9px; }
.pt-pixel-stat { border: 1px solid #4c1d95; padding: 6px 7px; background: #170c30; }
.pt-pixel-stat span { display: block; font-size: 7px; letter-spacing: 0.14em; color: #7c6bb0; }
.pt-pixel-stat strong { display: block; margin-top: 2px; font-size: 10px; line-height: 1.45; font-weight: 700; color: #ddd6fe; }
.pt-pixel-alert { display: flex; gap: 7px; align-items: flex-start; margin-top: 9px; padding: 7px 8px; border: 2px solid #f43f5e; background: rgba(244, 63, 94, 0.1); }
.pt-pixel-alert-tag { flex: none; background: #f43f5e; color: #1b1035; padding: 2px 4px; font-size: 7px; font-weight: 700; }
.pt-pixel-alert p { margin: 0; font-size: 10px; line-height: 1.55; color: #fecdd3; }
.pt-pixel-dialog { position: relative; margin-top: 9px; padding: 8px 20px 8px 9px; border: 2px solid #e9d5ff; background: #f5f3ff; color: #2e1065; }
.pt-pixel-dialog-name { font-size: 9px; font-weight: 700; color: #6d28d9; }
.pt-pixel-dialog p { margin: 2px 0 0; font-family: 'PingFang SC', 'Microsoft YaHei', monospace; font-size: 10px; line-height: 1.55; }
.pt-pixel-next { position: absolute; right: 6px; bottom: 5px; color: #6d28d9; font-size: 8px; animation: pt-pixel-blink 1.3s steps(2, end) infinite; }
.pt-pixel-call { display: block; margin: 12px 14px 0; padding: 10px 12px; background: #a3e635; color: #1a2e05; text-decoration: none; border: 3px solid #65a30d; box-shadow: 4px 4px 0 #365314; transition: transform 0.12s steps(2, end), box-shadow 0.12s steps(2, end); }
.pt-pixel-call:hover { background: #bef264; }
.pt-pixel-call:active { transform: translate(3px, 3px); box-shadow: 1px 1px 0 #365314; }
.pt-pixel-call-key { display: block; font-size: 8px; font-weight: 700; letter-spacing: 0.12em; }
.pt-pixel-call strong { display: block; margin: 1px 0; font-size: clamp(26px, 8vw, 33px); font-weight: 700; letter-spacing: 0.01em; }
.pt-pixel-call-hint { font-size: 8px; letter-spacing: 0.08em; }
.pt-pixel-footer { display: flex; justify-content: space-between; margin: 12px 14px 0; padding-top: 8px; border-top: 1px dashed rgba(167, 139, 250, 0.45); font-size: 7px; letter-spacing: 0.08em; color: #8b7ab8; }
.pt-pixel-press { margin: 16px 0 0; text-align: center; font-size: 9px; letter-spacing: 0.18em; color: #a3e635; animation: pt-pixel-blink 1.5s steps(2, end) infinite; }
@keyframes pt-pixel-boot {
  from { opacity: 0; clip-path: inset(0 0 100% 0); }
  to { opacity: 1; clip-path: inset(0 0 0 0); }
}
@keyframes pt-pixel-idle {
  0%, 100% { transform: translateY(0); }
  50% { transform: translateY(-5px); }
}
@keyframes pt-pixel-blink {
  0%, 49% { opacity: 1; }
  50%, 100% { opacity: 0.25; }
}
@media (prefers-reduced-motion: reduce) {
  .pt-pixel-content, .pt-pixel-sprite, .pt-pixel-press, .pt-pixel-mark, .pt-pixel-cursor, .pt-pixel-next { animation: none; }
  .pt-pixel-console, .pt-pixel-call { transition: none; }
}
</style>
