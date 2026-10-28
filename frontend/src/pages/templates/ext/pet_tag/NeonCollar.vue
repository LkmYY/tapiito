<template>
  <div class="pt-neon-root min-h-screen px-5 py-9 text-fuchsia-50">
    <div class="pt-neon-content mx-auto w-full max-w-md">
      <header class="pt-neon-head">
        <span class="pt-neon-live"><i></i> NIGHT WALK</span>
        <span class="pt-neon-id">COLLAR {{ nfcId || '0086' }}</span>
      </header>

      <article class="pt-neon-card">
        <div class="pt-neon-glow pt-neon-glow-a" aria-hidden="true"></div>
        <div class="pt-neon-glow pt-neon-glow-b" aria-hidden="true"></div>

        <div class="pt-neon-collar" aria-hidden="true">
          <span class="pt-neon-band"></span>
          <span class="pt-neon-buckle"></span>
          <span class="pt-neon-tagring"></span>
        </div>

        <div class="pt-neon-sign">
          <p>OPEN 24H · IF FOUND CALL</p>
          <h1>{{ payload.petName || '未命名' }}</h1>
          <div class="pt-neon-tube" aria-hidden="true"></div>
          <span>{{ payload.breed }} / {{ payload.age || '—' }}</span>
        </div>

        <a class="pt-neon-phone" :href="'tel:' + payload.phone">
          <span class="pt-neon-phone-label">紧急联络 · EMERGENCY</span>
          <strong>{{ payload.phone }}</strong>
          <span class="pt-neon-phone-hint">TAP TO CALL · 电话比转发更快</span>
        </a>

        <div class="pt-neon-grid">
          <div class="pt-neon-cell">
            <span>OWNER</span>
            <p>{{ payload.ownerName }}</p>
          </div>
          <div class="pt-neon-cell">
            <span>HOME</span>
            <p>{{ payload.address }}</p>
          </div>
        </div>

        <div class="pt-neon-med">
          <span class="pt-neon-med-key">MED</span>
          <p>{{ payload.medical || '无特殊记录' }}</p>
        </div>

        <div class="pt-neon-note">
          <span class="pt-neon-quote" aria-hidden="true">“</span>
          <p>{{ payload.note || '怕生，请勿强行抱抱。先让我闻闻你。' }}</p>
        </div>

        <footer class="pt-neon-foot">
          <span class="pt-neon-paw" aria-hidden="true">🐾</span>
          <span>NIGHT WALKERS CLUB · KEEP THEM SAFE</span>
        </footer>
      </article>

      <p class="pt-neon-below">LOST &amp; FOUND · 夜色里也有人在等你</p>
    </div>
  </div>
</template>

<script setup>
defineProps(['payload', 'nfcId'])
</script>

<style scoped>
.pt-neon-root {
  background:
    radial-gradient(circle at 18% 4%, rgba(217, 70, 239, 0.24), transparent 34%),
    radial-gradient(circle at 88% 78%, rgba(34, 211, 238, 0.13), transparent 34%),
    linear-gradient(165deg, #0a0616, #110827 52%, #06040d);
  font-family: 'Avenir Next', 'PingFang SC', sans-serif;
}
.pt-neon-content { animation: pt-neon-on 0.9s cubic-bezier(0.2, 0.8, 0.3, 1) both; }
.pt-neon-head { display: flex; justify-content: space-between; align-items: center; margin: 0 5px 12px; font-family: 'Courier New', monospace; font-size: 7px; letter-spacing: 0.13em; color: #c084fc; }
.pt-neon-live { display: flex; align-items: center; gap: 5px; }
.pt-neon-live i { width: 5px; height: 5px; border-radius: 50%; background: #e879f9; box-shadow: 0 0 8px #e879f9; animation: pt-neon-pulse 2.2s ease-in-out infinite; }
.pt-neon-card { position: relative; padding: 19px 18px 16px; background: rgba(17, 9, 35, 0.8); border: 1px solid rgba(217, 70, 239, 0.45); border-radius: 20px; box-shadow: 0 0 0 1px rgba(34, 211, 238, 0.08), 0 22px 44px rgba(0, 0, 0, 0.55), inset 0 0 34px rgba(168, 85, 247, 0.08); overflow: hidden; transition: transform 0.35s ease, border-color 0.35s ease; }
.pt-neon-card:hover { transform: translateY(-4px); border-color: rgba(232, 121, 249, 0.7); }
.pt-neon-glow { position: absolute; border-radius: 50%; filter: blur(38px); pointer-events: none; }
.pt-neon-glow-a { width: 180px; height: 180px; top: -90px; left: -60px; background: rgba(217, 70, 239, 0.2); }
.pt-neon-glow-b { width: 160px; height: 160px; bottom: -80px; right: -60px; background: rgba(34, 211, 238, 0.14); }
.pt-neon-collar { position: relative; height: 53px; margin: 0 6px 6px; }
.pt-neon-band { position: absolute; left: 0; right: 0; top: 9px; height: 23px; border: 2px solid #22d3ee; border-radius: 50%; box-shadow: 0 0 7px rgba(34, 211, 238, 0.65), inset 0 0 7px rgba(34, 211, 238, 0.32); transform: rotate(-2deg); }
.pt-neon-band::after { content: ''; position: absolute; inset: 2px 30% 2px 2px; border-top: 4px solid #0e7490; border-radius: 50%; opacity: 0.55; }
.pt-neon-buckle { position: absolute; right: 22px; top: 12px; width: 21px; height: 17px; border: 2px solid #e879f9; border-radius: 3px; box-shadow: 0 0 6px rgba(232, 121, 249, 0.5); }
.pt-neon-tagring { position: absolute; left: 50%; bottom: 0; width: 17px; height: 17px; margin-left: -8.5px; border: 2px solid #22d3ee; border-radius: 50%; box-shadow: 0 0 7px rgba(34, 211, 238, 0.55); }
.pt-neon-sign { text-align: center; padding: 8px 0 11px; }
.pt-neon-sign p { margin: 0; font-family: 'Courier New', monospace; font-size: 6.5px; letter-spacing: 0.26em; color: #a78bfa; }
.pt-neon-sign h1 { margin: 6px 0 4px; font-size: clamp(48px, 15vw, 67px); line-height: 0.98; font-weight: 700; letter-spacing: 0.06em; color: #fdf4ff; text-shadow: 0 0 4px #fff, 0 0 12px #e879f9, 0 0 28px rgba(217, 70, 239, 0.65); animation: pt-neon-flicker 5s linear infinite; }
.pt-neon-tube { width: 96px; height: 2px; margin: 0 auto 6px; background: #22d3ee; box-shadow: 0 0 4px #22d3ee, 0 0 12px rgba(34, 211, 238, 0.7); }
.pt-neon-sign > span { font-family: 'Courier New', monospace; font-size: 8px; letter-spacing: 0.14em; color: #67e8f9; }
.pt-neon-phone { display: block; padding: 10px 13px 11px; border: 1px solid #e879f9; background: rgba(217, 70, 239, 0.09); box-shadow: 0 0 12px rgba(217, 70, 239, 0.2), inset 0 0 17px rgba(217, 70, 239, 0.06); text-decoration: none; transition: background 0.2s ease, box-shadow 0.2s ease; }
.pt-neon-phone:hover { background: rgba(217, 70, 239, 0.16); box-shadow: 0 0 20px rgba(217, 70, 239, 0.32); }
.pt-neon-phone-label { display: block; font-family: 'Courier New', monospace; font-size: 6.5px; letter-spacing: 0.15em; color: #f0abfc; }
.pt-neon-phone strong { display: block; margin: 2px 0; font-family: 'Arial', 'PingFang SC', sans-serif; font-size: clamp(27px, 8.3vw, 35px); font-weight: 700; letter-spacing: 0.02em; color: #fff; text-shadow: 0 0 9px rgba(232, 121, 249, 0.75); }
.pt-neon-phone-hint { font-family: 'PingFang SC', 'Microsoft YaHei', sans-serif; font-size: 8px; letter-spacing: 0.05em; color: #c4b5fd; }
.pt-neon-grid { display: grid; grid-template-columns: 0.85fr 1.15fr; gap: 8px; margin-top: 10px; }
.pt-neon-cell { padding: 8px 9px; border: 1px solid rgba(34, 211, 238, 0.25); background: rgba(34, 211, 238, 0.035); }
.pt-neon-cell span { display: block; font-family: 'Courier New', monospace; font-size: 6px; letter-spacing: 0.14em; color: #67e8f9; }
.pt-neon-cell p { margin: 2px 0 0; font-size: 10px; line-height: 1.5; color: #e9d5ff; }
.pt-neon-med { display: flex; gap: 8px; align-items: flex-start; margin-top: 8px; padding: 8px 9px; border-left: 2px solid #fb7185; background: rgba(251, 113, 133, 0.07); }
.pt-neon-med-key { flex: none; padding: 1px 4px; border: 1px solid #fb7185; color: #fda4af; font-family: 'Courier New', monospace; font-size: 6px; letter-spacing: 0.1em; }
.pt-neon-med p { margin: 0; font-size: 10px; line-height: 1.55; color: #fecdd3; }
.pt-neon-note { position: relative; margin-top: 9px; padding: 9px 10px 9px 24px; background: rgba(139, 92, 246, 0.07); }
.pt-neon-quote { position: absolute; left: 7px; top: 1px; font-family: Georgia, serif; font-size: 25px; color: #a78bfa; }
.pt-neon-note p { margin: 0; font-size: 10px; line-height: 1.65; color: #ddd6fe; }
.pt-neon-foot { display: flex; align-items: center; gap: 6px; margin-top: 13px; padding-top: 9px; border-top: 1px solid rgba(167, 139, 250, 0.2); font-family: 'Courier New', monospace; font-size: 6px; letter-spacing: 0.09em; color: #8b7ab8; }
.pt-neon-below { margin: 17px 0 0; text-align: center; font-family: 'Courier New', monospace; font-size: 8px; letter-spacing: 0.13em; color: #a78bfa; }
@keyframes pt-neon-on {
  0% { opacity: 0; transform: scale(0.97); filter: brightness(2.2); }
  20% { opacity: 0.4; }
  28% { opacity: 0.1; }
  36% { opacity: 1; }
  44% { opacity: 0.55; }
  100% { opacity: 1; filter: brightness(1); transform: scale(1); }
}
@keyframes pt-neon-flicker {
  0%, 7%, 9%, 12%, 100% { opacity: 1; }
  8% { opacity: 0.55; }
  10% { opacity: 0.8; }
  11% { opacity: 0.35; }
}
@keyframes pt-neon-pulse {
  0%, 100% { opacity: 1; }
  50% { opacity: 0.3; }
}
@media (prefers-reduced-motion: reduce) {
  .pt-neon-content, .pt-neon-sign h1, .pt-neon-live i { animation: none; }
  .pt-neon-card, .pt-neon-phone { transition: none; }
}
</style>
