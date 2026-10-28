<template>
  <div class="pt-retro-root min-h-screen px-5 py-8 text-amber-950">
    <div class="pt-retro-content mx-auto w-full max-w-md">
      <div class="pt-retro-awning" aria-hidden="true">
        <span v-for="n in 10" :key="n" :class="n % 2 ? 'pt-retro-stripe-a' : 'pt-retro-stripe-b'"></span>
      </div>

      <article class="pt-retro-card">
        <header class="pt-retro-head">
          <div class="pt-retro-badge" aria-hidden="true">
            <svg viewBox="0 0 48 48" fill="currentColor">
              <ellipse cx="24" cy="30" rx="9" ry="7" />
              <ellipse cx="12" cy="21" rx="3.6" ry="4.6" />
              <ellipse cx="21" cy="15" rx="3.6" ry="5" />
              <ellipse cx="31" cy="16" rx="3.6" ry="4.6" />
              <ellipse cx="39" cy="23" rx="3.2" ry="4" />
            </svg>
          </div>
          <div class="pt-retro-shop">
            <p>EST. 1978 · PET SUPPLY</p>
            <h1>Woof &amp; Co.</h1>
          </div>
          <span class="pt-retro-open">OPEN</span>
        </header>

        <div class="pt-retro-oval">
          <p class="pt-retro-oval-small">RESIDENT PET</p>
          <h2>{{ payload.petName || '未命名' }}</h2>
          <p class="pt-retro-oval-meta">{{ payload.breed }} / {{ payload.age || '—' }}</p>
        </div>

        <div class="pt-retro-cols">
          <div class="pt-retro-col">
            <span class="pt-retro-label">OWNER</span>
            <p>{{ payload.ownerName }}</p>
            <span class="pt-retro-label">HOME</span>
            <p>{{ payload.address }}</p>
          </div>
          <div class="pt-retro-col pt-retro-col-line">
            <span class="pt-retro-label">MEDICAL</span>
            <p class="pt-retro-med">{{ payload.medical || '无特殊记录' }}</p>
          </div>
        </div>

        <div class="pt-retro-note">
          <span aria-hidden="true">✳</span>
          <p>{{ payload.note || '喜欢被摸下巴，怕突然的响声。' }}</p>
        </div>

        <a class="pt-retro-phone" :href="'tel:' + payload.phone">
          <span class="pt-retro-phone-top">FOUND ME? CALL THE SHOP</span>
          <span class="pt-retro-phone-num">{{ payload.phone }}</span>
          <span class="pt-retro-phone-bottom">走失请致电 · 随时接听</span>
        </a>

        <footer class="pt-retro-foot">
          <span class="pt-retro-receipt">RECEIPT #{{ nfcId || '1978' }}</span>
          <span>GOOD PET · GOOD DAY</span>
        </footer>
      </article>

      <div class="pt-retro-coins" aria-hidden="true">
        <span>✺</span><span>✳</span><span>✺</span>
      </div>
      <p class="pt-retro-bottom">青禾街 12 号 · 营业中</p>
    </div>
  </div>
</template>

<script setup>
defineProps(['payload', 'nfcId'])
</script>

<style scoped>
.pt-retro-root {
  background:
    radial-gradient(circle at 50% 0, rgba(255, 210, 100, 0.22), transparent 35%),
    repeating-linear-gradient(90deg, rgba(120, 60, 10, 0.035) 0 1px, transparent 1px 9px),
    #f4e2bd;
  font-family: 'Cooper Black', 'Rockwell', 'Songti SC', serif;
}
.pt-retro-content { animation: pt-retro-drop 0.7s cubic-bezier(0.2, 0.9, 0.3, 1) both; }
.pt-retro-awning { display: grid; grid-template-columns: repeat(10, 1fr); height: 24px; overflow: hidden; border-radius: 9px 9px 0 0; box-shadow: 0 4px 6px rgba(80, 40, 10, 0.16); }
.pt-retro-awning span { display: block; }
.pt-retro-stripe-a { background: #c05a2b; }
.pt-retro-stripe-b { background: #f3dfb3; }
.pt-retro-card {
  position: relative;
  padding: 20px 19px 0;
  background:
    radial-gradient(circle at 12% 10%, rgba(255, 255, 255, 0.35), transparent 24%),
    #fbf0d5;
  border: 2px solid #6e3d1c;
  border-top: 0;
  border-radius: 0 0 12px 12px;
  box-shadow: 0 20px 35px rgba(92, 52, 18, 0.2);
  transition: transform 0.3s ease;
}
.pt-retro-card:hover { transform: translateY(-3px); }
.pt-retro-card::after { content: ''; position: absolute; inset: 7px; border: 1px solid rgba(110, 61, 28, 0.25); border-radius: 0 0 8px 8px; pointer-events: none; }
.pt-retro-head { display: flex; align-items: center; gap: 10px; padding-bottom: 12px; border-bottom: 2px solid #6e3d1c; }
.pt-retro-badge { width: 42px; height: 42px; flex: none; display: grid; place-items: center; color: #fbf0d5; background: #6e3d1c; border-radius: 50%; }
.pt-retro-badge svg { width: 26px; height: 26px; }
.pt-retro-shop { flex: 1; }
.pt-retro-shop p { margin: 0; font-family: 'Courier New', monospace; font-size: 7px; font-weight: 700; letter-spacing: 0.12em; }
.pt-retro-shop h1 { margin: 0; font-size: 26px; line-height: 1; color: #6e3d1c; letter-spacing: -0.01em; }
.pt-retro-open { padding: 4px 7px; border: 2px solid #a33f1f; color: #a33f1f; font-family: 'Courier New', monospace; font-size: 9px; font-weight: 700; letter-spacing: 0.1em; transform: rotate(5deg); }
.pt-retro-oval {
  position: relative;
  margin: 17px 6px 14px;
  padding: 16px 12px 13px;
  text-align: center;
  border: 3px solid #6e3d1c;
  border-radius: 50%;
  background: #e8a23a;
  box-shadow: inset 0 0 0 5px #fbf0d5, inset 0 0 0 6px rgba(110, 61, 28, 0.45), 0 5px 0 rgba(110, 61, 28, 0.14);
  transform: rotate(-1.2deg);
}
.pt-retro-oval-small { margin: 0; font-family: 'Courier New', monospace; font-size: 7px; font-weight: 700; letter-spacing: 0.22em; color: #6a3614; }
.pt-retro-oval h2 { margin: 0; font-size: clamp(44px, 13vw, 60px); line-height: 1; letter-spacing: 0.01em; color: #48230b; text-shadow: 2px 2px 0 #f5cd7e; }
.pt-retro-oval-meta { margin: 5px 0 0; font-family: 'Courier New', monospace; font-size: 9px; font-weight: 700; letter-spacing: 0.12em; }
.pt-retro-cols { display: grid; grid-template-columns: 1.15fr 0.85fr; gap: 13px; padding: 4px 2px 12px; }
.pt-retro-col-line { padding-left: 12px; border-left: 1px dashed rgba(110, 61, 28, 0.45); }
.pt-retro-label { display: block; margin-top: 7px; font-family: 'Courier New', monospace; font-size: 7px; font-weight: 700; letter-spacing: 0.16em; color: #a36b33; }
.pt-retro-col p { margin: 2px 0 0; font-family: 'Avenir Next', 'PingFang SC', sans-serif; font-size: 12px; line-height: 1.5; font-weight: 600; color: #5a3418; }
.pt-retro-col .pt-retro-med { color: #9b2c12; }
.pt-retro-note { display: flex; gap: 8px; align-items: flex-start; margin: 0 2px 15px; padding: 9px 10px; background: #f0d39f; border: 1px solid rgba(110, 61, 28, 0.4); border-radius: 3px 9px 3px 9px; }
.pt-retro-note span { flex: none; font-size: 16px; color: #a33f1f; line-height: 1.1; }
.pt-retro-note p { margin: 0; font-family: 'PingFang SC', 'Microsoft YaHei', sans-serif; font-size: 11px; line-height: 1.6; color: #63391a; }
.pt-retro-phone {
  display: block;
  margin: 0 -19px;
  padding: 12px 19px 13px;
  background: #6e3d1c;
  color: #fdf3db;
  text-decoration: none;
  transition: background 0.2s ease;
}
.pt-retro-phone:hover { background: #834c25; }
.pt-retro-phone-top { display: block; font-family: 'Courier New', monospace; font-size: 7px; font-weight: 700; letter-spacing: 0.2em; color: #e9b96c; }
.pt-retro-phone-num { display: block; margin: 1px 0; font-family: 'Avenir Next', 'PingFang SC', sans-serif; font-size: clamp(27px, 8.3vw, 35px); font-weight: 800; letter-spacing: 0.01em; color: #fff4dc; }
.pt-retro-phone-bottom { font-family: 'PingFang SC', 'Microsoft YaHei', sans-serif; font-size: 9px; letter-spacing: 0.1em; color: #d9a95f; }
.pt-retro-foot { display: flex; justify-content: space-between; padding: 12px 0 13px; font-family: 'Courier New', monospace; font-size: 7px; font-weight: 700; letter-spacing: 0.1em; color: #7e4b24; }
.pt-retro-receipt { padding-bottom: 2px; border-bottom: 1px dashed #a3723f; }
.pt-retro-coins { display: flex; justify-content: center; gap: 20px; margin-top: 12px; color: #b56632; font-size: 17px; }
.pt-retro-bottom { margin: 9px 0 0; text-align: center; font-family: 'Courier New', monospace; font-size: 8px; letter-spacing: 0.18em; color: #8e5c2e; }
@keyframes pt-retro-drop {
  from { opacity: 0; transform: translateY(-32px); }
  to { opacity: 1; transform: translateY(0); }
}
@media (prefers-reduced-motion: reduce) {
  .pt-retro-content { animation: none; }
  .pt-retro-card, .pt-retro-phone { transition: none; }
}
</style>
