<template>
  <div class="pt-glass-root min-h-screen px-5 py-9 text-slate-800">
    <div class="pt-glass-content mx-auto w-full max-w-md">
      <div class="pt-glass-orb pt-glass-orb-a" aria-hidden="true"></div>
      <div class="pt-glass-orb pt-glass-orb-b" aria-hidden="true"></div>
      <div class="pt-glass-orb pt-glass-orb-c" aria-hidden="true"></div>
      <div class="pt-glass-rings" aria-hidden="true"></div>

      <header class="pt-glass-head">
        <span class="pt-glass-logo">
          <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5">
            <path d="M6 15a5 5 0 0 1 1-9.9A6 6 0 0 1 18.5 7 4.5 4.5 0 0 1 18 16H7" />
          </svg>
          CLOUD TAG
        </span>
        <span class="pt-glass-id">{{ nfcId || '0086' }}</span>
      </header>

      <article class="pt-glass-card">
        <div class="pt-glass-sheen" aria-hidden="true"></div>

        <div class="pt-glass-title">
          <span class="pt-glass-hello">a soft place to land</span>
          <h1>{{ payload.petName || '未命名' }}</h1>
          <p>{{ payload.breed }} · {{ payload.age || '—' }}</p>
        </div>

        <div class="pt-glass-divider"></div>

        <div class="pt-glass-grid">
          <div class="pt-glass-tile">
            <span>家人 / OWNER</span>
            <p>{{ payload.ownerName }}</p>
          </div>
          <div class="pt-glass-tile">
            <span>住址 / HOME</span>
            <p>{{ payload.address }}</p>
          </div>
        </div>

        <div class="pt-glass-med">
          <span class="pt-glass-med-badge">MED</span>
          <div>
            <span>健康备注</span>
            <p>{{ payload.medical || '无特殊记录' }}</p>
          </div>
        </div>

        <div class="pt-glass-note">
          <span class="pt-glass-note-icon" aria-hidden="true">❋</span>
          <p>{{ payload.note || '怕生，请勿强行抱抱。让我先靠近你。' }}</p>
        </div>

        <a class="pt-glass-phone" :href="'tel:' + payload.phone">
          <span class="pt-glass-phone-label">走失时，请拨这个号码</span>
          <strong>{{ payload.phone }}</strong>
          <span class="pt-glass-phone-row">
            <span class="pt-glass-phone-hint">TAP TO CALL</span>
            <span class="pt-glass-phone-arrow" aria-hidden="true">→</span>
          </span>
        </a>

        <footer class="pt-glass-foot">
          <span>FLOATING SAFELY · 青禾社区</span>
          <span>NFC · TAP</span>
        </footer>
      </article>

      <p class="pt-glass-below">softly found · safely home</p>
    </div>
  </div>
</template>

<script setup>
defineProps(['payload', 'nfcId'])
</script>

<style scoped>
.pt-glass-root {
  position: relative;
  overflow: hidden;
  background:
    radial-gradient(circle at 15% 8%, #ffd9c0 0%, transparent 32%),
    radial-gradient(circle at 88% 22%, #c7b8ff 0%, transparent 36%),
    radial-gradient(circle at 30% 90%, #9fe3e0 0%, transparent 40%),
    linear-gradient(160deg, #d7e9f4 0%, #efe4f7 48%, #ffe7d6 100%);
  font-family: 'Avenir Next', 'PingFang SC', sans-serif;
}
.pt-glass-orb { position: absolute; border-radius: 50%; filter: blur(2px); pointer-events: none; opacity: 0.55; }
.pt-glass-orb-a { width: 180px; height: 180px; top: 120px; right: -60px; background: radial-gradient(circle at 35% 30%, rgba(255, 255, 255, 0.8), rgba(197, 168, 255, 0.35)); }
.pt-glass-orb-b { width: 120px; height: 120px; top: 390px; left: -45px; background: radial-gradient(circle at 35% 30%, rgba(255, 255, 255, 0.75), rgba(130, 220, 210, 0.4)); }
.pt-glass-orb-c { width: 77px; height: 77px; top: 42px; left: 26px; background: radial-gradient(circle at 35% 30%, rgba(255, 255, 255, 0.8), rgba(255, 190, 154, 0.4)); }
.pt-glass-rings { position: absolute; top: 8%; left: 50%; width: 360px; height: 360px; margin-left: -180px; border: 1px solid rgba(255, 255, 255, 0.4); border-radius: 50%; pointer-events: none; }
.pt-glass-rings::after { content: ''; position: absolute; inset: 45px; border: 1px solid rgba(255, 255, 255, 0.3); border-radius: 50%; }
.pt-glass-content { position: relative; z-index: 1; animation: pt-glass-in 0.95s cubic-bezier(0.2, 0.85, 0.3, 1) both; }
.pt-glass-head { display: flex; justify-content: space-between; align-items: center; margin: 0 6px 12px; color: #51606b; }
.pt-glass-logo { display: flex; align-items: center; gap: 5px; font-family: 'Courier New', monospace; font-size: 8px; font-weight: 700; letter-spacing: 0.13em; }
.pt-glass-logo svg { width: 17px; height: 17px; }
.pt-glass-id { font-family: 'Courier New', monospace; font-size: 8px; letter-spacing: 0.08em; }
.pt-glass-card {
  position: relative;
  padding: 24px 19px 17px;
  background: linear-gradient(135deg, rgba(255, 255, 255, 0.62), rgba(255, 255, 255, 0.3));
  border: 1px solid rgba(255, 255, 255, 0.7);
  border-radius: 25px;
  box-shadow: 0 22px 45px rgba(89, 96, 128, 0.18), inset 0 1px 0 rgba(255, 255, 255, 0.8);
  backdrop-filter: blur(18px) saturate(1.15);
  -webkit-backdrop-filter: blur(18px) saturate(1.15);
  overflow: hidden;
  transition: transform 0.4s ease, box-shadow 0.4s ease;
}
.pt-glass-card:hover { transform: translateY(-4px); box-shadow: 0 30px 55px rgba(89, 96, 128, 0.22), inset 0 1px 0 rgba(255, 255, 255, 0.9); }
.pt-glass-sheen { position: absolute; top: -80px; left: -50px; width: 230px; height: 190px; background: radial-gradient(circle, rgba(255, 255, 255, 0.7), transparent 65%); pointer-events: none; }
.pt-glass-title { position: relative; }
.pt-glass-hello { font-family: 'Georgia', 'Songti SC', serif; font-size: 12px; font-style: italic; color: #71808a; }
.pt-glass-title h1 { margin: 2px 0 0; font-family: 'Georgia', 'Songti SC', serif; font-size: clamp(52px, 15vw, 68px); line-height: 0.98; font-weight: 400; letter-spacing: -0.02em; color: #374a58; text-shadow: 0 1px 0 rgba(255, 255, 255, 0.7); }
.pt-glass-title p { margin: 8px 0 0; font-size: 10px; letter-spacing: 0.16em; color: #6c7f8d; }
.pt-glass-divider { height: 1px; margin: 18px 0 14px; background: linear-gradient(90deg, transparent, rgba(94, 111, 125, 0.28), transparent); }
.pt-glass-grid { display: grid; grid-template-columns: 0.85fr 1.15fr; gap: 8px; }
.pt-glass-tile { padding: 9px 10px; background: rgba(255, 255, 255, 0.33); border: 1px solid rgba(255, 255, 255, 0.6); border-radius: 12px; }
.pt-glass-tile span { display: block; font-family: 'Courier New', monospace; font-size: 6.5px; font-weight: 700; letter-spacing: 0.09em; color: #7b8b96; }
.pt-glass-tile p { margin: 2px 0 0; font-size: 11px; font-weight: 500; line-height: 1.45; color: #42545f; }
.pt-glass-med { display: flex; align-items: flex-start; gap: 8px; margin-top: 9px; padding: 9px 10px; border-radius: 12px; background: rgba(255, 240, 235, 0.46); border: 1px solid rgba(255, 255, 255, 0.62); }
.pt-glass-med-badge { flex: none; padding: 2px 5px; border-radius: 4px; background: rgba(214, 123, 100, 0.85); color: #fffaf6; font-family: 'Courier New', monospace; font-size: 6.5px; font-weight: 700; letter-spacing: 0.06em; }
.pt-glass-med span { font-family: 'PingFang SC', 'Microsoft YaHei', sans-serif; font-size: 7.5px; letter-spacing: 0.06em; color: #a17668; }
.pt-glass-med p { margin: 2px 0 0; font-size: 10px; line-height: 1.5; color: #7f4c3e; }
.pt-glass-note { display: flex; gap: 8px; align-items: flex-start; margin: 10px 1px 13px; }
.pt-glass-note-icon { flex: none; color: #8e9dad; font-size: 12px; line-height: 1.3; }
.pt-glass-note p { margin: 0; font-family: 'Georgia', 'Songti SC', serif; font-size: 12px; font-style: italic; line-height: 1.75; color: #5a6b77; }
.pt-glass-phone {
  display: block;
  padding: 12px 14px 11px;
  border-radius: 16px;
  background: linear-gradient(135deg, rgba(58, 79, 92, 0.84), rgba(76, 96, 112, 0.7));
  border: 1px solid rgba(255, 255, 255, 0.35);
  box-shadow: 0 8px 19px rgba(55, 72, 86, 0.24), inset 0 1px 0 rgba(255, 255, 255, 0.18);
  color: #f7fbfd;
  text-decoration: none;
  transition: background 0.22s ease, transform 0.18s ease;
}
.pt-glass-phone:hover { background: linear-gradient(135deg, rgba(48, 70, 84, 0.92), rgba(68, 89, 105, 0.8)); }
.pt-glass-phone:active { transform: scale(0.99); }
.pt-glass-phone-label { display: block; font-size: 8px; letter-spacing: 0.08em; color: #c8d8e0; }
.pt-glass-phone strong { display: block; margin: 1px 0; font-size: clamp(26px, 7.9vw, 33px); font-weight: 600; letter-spacing: 0.02em; }
.pt-glass-phone-row { display: flex; align-items: center; justify-content: space-between; }
.pt-glass-phone-hint { font-family: 'Courier New', monospace; font-size: 6.5px; letter-spacing: 0.16em; color: #a9c2cf; }
.pt-glass-phone-arrow { font-size: 16px; color: #cfe3eb; transition: transform 0.2s ease; }
.pt-glass-phone:hover .pt-glass-phone-arrow { transform: translateX(3px); }
.pt-glass-foot { display: flex; justify-content: space-between; margin-top: 13px; font-family: 'Courier New', monospace; font-size: 6px; letter-spacing: 0.1em; color: #8496a1; }
.pt-glass-below { margin: 18px 0 0; text-align: center; font-family: 'Georgia', serif; font-size: 11px; font-style: italic; letter-spacing: 0.06em; color: #667b88; }
@keyframes pt-glass-in {
  from { opacity: 0; filter: blur(9px); transform: translateY(18px) scale(0.985); }
  to { opacity: 1; filter: blur(0); transform: translateY(0) scale(1); }
}
@media (prefers-reduced-motion: reduce) {
  .pt-glass-content { animation: none; }
  .pt-glass-card, .pt-glass-phone, .pt-glass-phone-arrow { transition: none; }
}
</style>
