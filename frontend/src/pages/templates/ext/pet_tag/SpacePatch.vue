<template>
  <div class="pt-space-root min-h-screen px-5 py-8 text-indigo-50">
    <div class="pt-space-content mx-auto w-full max-w-md">
      <header class="pt-space-head">
        <span>MISSION CONTROL · HOME</span>
        <span>SIGNAL {{ nfcId || '0086' }}</span>
      </header>

      <article class="pt-space-card">
        <div class="pt-space-stars" aria-hidden="true">
          <span v-for="n in 30" :key="n" :style="{ left: ((n * 37) % 96) + '%', top: ((n * 53) % 88) + '%', opacity: 0.25 + ((n % 5) * 0.14) }"></span>
        </div>
        <div class="pt-space-orbit" aria-hidden="true"></div>
        <div class="pt-space-planet" aria-hidden="true"></div>

        <div class="pt-space-patch">
          <div class="pt-space-patch-inner">
            <span class="pt-space-patch-top">MISSION</span>
            <h1>{{ payload.petName || '未命名' }}</h1>
            <span class="pt-space-patch-sub">{{ payload.breed }}</span>
            <div class="pt-space-rocket" aria-hidden="true">
              <span class="pt-space-rocket-body"></span>
              <span class="pt-space-rocket-window"></span>
              <span class="pt-space-rocket-fin pt-space-fin-l"></span>
              <span class="pt-space-rocket-fin pt-space-rocket-fin-r"></span>
              <span class="pt-space-rocket-flame"></span>
            </div>
            <span class="pt-space-patch-bottom">RETURN TO EARTH · {{ payload.age || '—' }}</span>
          </div>
          <span class="pt-space-stitch" aria-hidden="true"></span>
        </div>

        <div class="pt-space-log">
          <div class="pt-space-log-head">
            <span>CREW MANIFEST</span>
            <span>LOG 04</span>
          </div>
          <div class="pt-space-row">
            <span>COMMANDER</span>
            <p>{{ payload.ownerName }}</p>
          </div>
          <div class="pt-space-row">
            <span>BASECAMP</span>
            <p>{{ payload.address }}</p>
          </div>
          <div class="pt-space-row pt-space-row-alert">
            <span>FLIGHT NOTE</span>
            <p>{{ payload.medical || '无特殊记录' }}</p>
          </div>
        </div>

        <div class="pt-space-note">
          <span class="pt-space-note-icon" aria-hidden="true">✦</span>
          <p>{{ payload.note || '怕生，请勿强行抱抱。请先联系地面控制中心。' }}</p>
        </div>

        <a class="pt-space-phone" :href="'tel:' + payload.phone">
          <span class="pt-space-phone-top">
            <span class="pt-space-signal" aria-hidden="true"></span>
            GROUND CONTROL · 紧急联络
          </span>
          <strong>{{ payload.phone }}</strong>
          <span class="pt-space-phone-sub">OPEN CHANNEL · 24H · 请协助返航</span>
        </a>

        <footer class="pt-space-foot">
          <span>MISSION: SAFE RETURN</span>
          <span>NFC {{ nfcId || '086' }}</span>
        </footer>
      </article>

      <p class="pt-space-below">— even the smallest explorer deserves a way home —</p>
    </div>
  </div>
</template>

<script setup>
defineProps(['payload', 'nfcId'])
</script>

<style scoped>
.pt-space-root {
  background:
    radial-gradient(circle at 50% 0, rgba(99, 102, 241, 0.2), transparent 40%),
    #080b1d;
  font-family: 'Avenir Next', 'PingFang SC', sans-serif;
}
.pt-space-content { animation: pt-space-arrive 0.95s cubic-bezier(0.2, 0.8, 0.25, 1) both; }
.pt-space-head { display: flex; justify-content: space-between; margin: 0 5px 12px; font-family: 'Courier New', monospace; font-size: 7px; letter-spacing: 0.13em; color: #818cf8; }
.pt-space-card {
  position: relative;
  padding: 19px 18px 15px;
  background:
    radial-gradient(circle at 70% 4%, rgba(99, 102, 241, 0.16), transparent 33%),
    linear-gradient(165deg, #111735 0%, #0d1230 58%, #080d22 100%);
  border: 1px solid rgba(129, 140, 248, 0.35);
  border-radius: 18px;
  box-shadow: 0 24px 44px rgba(0, 0, 0, 0.5), inset 0 0 40px rgba(79, 70, 229, 0.07);
  overflow: hidden;
  transition: transform 0.35s ease, border-color 0.35s ease;
}
.pt-space-card:hover { transform: translateY(-4px); border-color: rgba(129, 140, 248, 0.6); }
.pt-space-stars { position: absolute; inset: 0; pointer-events: none; }
.pt-space-stars span { position: absolute; width: 2px; height: 2px; border-radius: 50%; background: #e0e7ff; }
.pt-space-orbit { position: absolute; top: -98px; left: 50%; width: 325px; height: 325px; margin-left: -162px; border: 1px solid rgba(129, 140, 248, 0.16); border-radius: 50%; pointer-events: none; }
.pt-space-planet { position: absolute; top: 8px; right: -16px; width: 44px; height: 44px; border-radius: 50%; background: radial-gradient(circle at 34% 28%, #c7d2fe, #6366f1 55%, #312e81); opacity: 0.55; box-shadow: 0 0 22px rgba(99, 102, 241, 0.35); }
.pt-space-patch { position: relative; z-index: 1; width: 214px; height: 214px; margin: 9px auto 14px; padding: 7px; border-radius: 50%; background: conic-gradient(from 20deg, #fbbf24, #b45309, #fde68a, #92400e, #fbbf24); box-shadow: 0 13px 24px rgba(0, 0, 0, 0.4), 0 0 0 3px rgba(251, 191, 36, 0.12); transition: transform 0.45s ease; animation: pt-space-float 5s ease-in-out infinite; }
.pt-space-card:hover .pt-space-patch { transform: rotate(-3deg) scale(1.02); }
.pt-space-patch-inner { position: relative; height: 100%; display: flex; flex-direction: column; align-items: center; justify-content: center; border-radius: 50%; background: radial-gradient(circle at 32% 24%, #2f3a7d 0%, #1a2252 52%, #0c1130 100%); box-shadow: inset 0 0 0 2px rgba(255, 255, 255, 0.1), inset 0 0 25px rgba(0, 0, 0, 0.5); overflow: hidden; }
.pt-space-patch-top { font-family: 'Courier New', monospace; font-size: 6.5px; letter-spacing: 0.3em; color: #fbbf24; }
.pt-space-patch-inner h1 { margin: 2px 0 0; max-width: 154px; font-size: clamp(27px, 8.3vw, 35px); line-height: 1.02; font-weight: 700; letter-spacing: 0.04em; color: #f8fafc; text-align: center; word-break: break-word; }
.pt-space-patch-sub { font-family: 'Courier New', monospace; font-size: 7px; letter-spacing: 0.13em; color: #a5b4fc; }
.pt-space-rocket { position: relative; width: 31px; height: 44px; margin: 7px 0 6px; animation: pt-space-drift 4s ease-in-out infinite; }
.pt-space-rocket-body { position: absolute; left: 9px; top: 0; width: 13px; height: 30px; border-radius: 50% 50% 28% 28%; background: linear-gradient(90deg, #cbd5e1, #f8fafc 48%, #94a3b8); }
.pt-space-rocket-window { position: absolute; left: 12px; top: 8px; width: 7px; height: 7px; border-radius: 50%; background: #22d3ee; box-shadow: 0 0 6px #22d3ee; }
.pt-space-rocket-fin { position: absolute; top: 23px; width: 9px; height: 13px; background: #fbbf24; }
.pt-space-fin-l { left: 2px; clip-path: polygon(100% 0, 100% 100%, 0 100%); }
.pt-space-rocket-fin-r { right: 2px; clip-path: polygon(0 0, 100% 100%, 0 100%); }
.pt-space-rocket-flame { position: absolute; left: 12px; top: 30px; width: 7px; height: 12px; background: linear-gradient(180deg, #fb923c, transparent); clip-path: polygon(50% 0, 100% 40%, 72% 100%, 28% 100%, 0 40%); }
.pt-space-patch-bottom { font-family: 'Courier New', monospace; font-size: 5.5px; letter-spacing: 0.14em; color: #fbbf24; }
.pt-space-stitch { position: absolute; inset: 4px; border: 1px dashed rgba(255, 247, 215, 0.5); border-radius: 50%; pointer-events: none; }
.pt-space-log { position: relative; z-index: 1; padding: 9px 11px; border: 1px solid rgba(129, 140, 248, 0.27); background: rgba(15, 22, 48, 0.7); }
.pt-space-log-head { display: flex; justify-content: space-between; padding-bottom: 5px; border-bottom: 1px solid rgba(129, 140, 248, 0.2); font-family: 'Courier New', monospace; font-size: 6px; font-weight: 700; letter-spacing: 0.15em; color: #a5b4fc; }
.pt-space-row { display: grid; grid-template-columns: 88px 1fr; gap: 8px; padding: 5px 0; border-bottom: 1px dotted rgba(129, 140, 248, 0.15); font-size: 10px; line-height: 1.45; }
.pt-space-row:last-child { border-bottom: 0; }
.pt-space-row span { font-family: 'Courier New', monospace; font-size: 6px; letter-spacing: 0.08em; color: #7f8bd8; padding-top: 2px; }
.pt-space-row p { margin: 0; color: #dbe2ff; }
.pt-space-row-alert p { color: #fca5a5; font-weight: 600; }
.pt-space-note { position: relative; z-index: 1; display: flex; gap: 7px; align-items: flex-start; margin: 9px 1px 11px; padding: 8px 9px; border-left: 2px solid #fbbf24; background: rgba(251, 191, 36, 0.055); }
.pt-space-note-icon { flex: none; color: #fbbf24; font-size: 9px; line-height: 1.3; }
.pt-space-note p { margin: 0; font-size: 10px; line-height: 1.6; color: #c7d2fe; }
.pt-space-phone { display: block; position: relative; z-index: 1; padding: 10px 12px 11px; border: 1px solid rgba(251, 191, 36, 0.6); background: rgba(251, 191, 36, 0.07); color: #fff8e7; text-decoration: none; transition: background 0.2s ease, box-shadow 0.2s ease; }
.pt-space-phone:hover { background: rgba(251, 191, 36, 0.12); box-shadow: 0 0 16px rgba(251, 191, 36, 0.12); }
.pt-space-phone-top { display: flex; align-items: center; gap: 6px; font-family: 'Courier New', monospace; font-size: 6.5px; letter-spacing: 0.09em; color: #fcd34d; }
.pt-space-signal { width: 5px; height: 5px; border-radius: 50%; background: #fbbf24; box-shadow: 0 0 7px #fbbf24; animation: pt-space-pulse 2s ease-in-out infinite; }
.pt-space-phone strong { display: block; margin: 2px 0; font-family: 'Arial', 'PingFang SC', sans-serif; font-size: clamp(27px, 8.2vw, 34px); font-weight: 700; letter-spacing: 0.02em; color: #fffbeb; }
.pt-space-phone-sub { font-family: 'Courier New', monospace; font-size: 6.5px; letter-spacing: 0.08em; color: #c7b98a; }
.pt-space-foot { position: relative; z-index: 1; display: flex; justify-content: space-between; margin-top: 12px; padding-top: 8px; border-top: 1px solid rgba(129, 140, 248, 0.2); font-family: 'Courier New', monospace; font-size: 6px; letter-spacing: 0.09em; color: #6f7ab8; }
.pt-space-below { margin: 17px 0 0; text-align: center; font-family: 'Courier New', monospace; font-size: 7px; letter-spacing: 0.06em; color: #6672a8; }
@keyframes pt-space-arrive {
  0% { opacity: 0; transform: translateY(26px) scale(0.98); }
  100% { opacity: 1; transform: translateY(0) scale(1); }
}
@keyframes pt-space-float {
  0%, 100% { transform: translateY(0) rotate(0); }
  50% { transform: translateY(-5px) rotate(1.2deg); }
}
@keyframes pt-space-drift {
  0%, 100% { transform: translate(0, 0); }
  50% { transform: translate(1px, -3px); }
}
@keyframes pt-space-pulse {
  0%, 100% { opacity: 1; }
  50% { opacity: 0.25; }
}
@media (prefers-reduced-motion: reduce) {
  .pt-space-content, .pt-space-patch, .pt-space-rocket, .pt-space-signal { animation: none; }
  .pt-space-card, .pt-space-patch, .pt-space-phone { transition: none; }
}
</style>
