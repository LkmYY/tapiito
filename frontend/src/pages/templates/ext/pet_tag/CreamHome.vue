<template>
  <div class="pt-home-root min-h-screen px-5 py-9 text-stone-700">
    <div class="pt-home-content mx-auto w-full max-w-md">
      <header class="pt-home-top">
        <span class="pt-home-brand">青禾 · 毛孩子之家</span>
        <span class="pt-home-tag">HOME TAG</span>
      </header>

      <article class="pt-home-card">
        <div class="pt-home-ribbon" aria-hidden="true"><span></span></div>

        <div class="pt-home-illustration" aria-hidden="true">
          <span class="pt-home-cloud pt-home-cloud-a"></span>
          <span class="pt-home-cloud pt-home-cloud-b"></span>
          <span class="pt-home-sun"></span>
          <span class="pt-home-house">
            <span class="pt-home-roof"></span>
            <span class="pt-home-window"></span>
            <span class="pt-home-door"></span>
          </span>
          <span class="pt-home-hill"></span>
          <span class="pt-home-bush pt-home-bush-a"></span>
          <span class="pt-home-bush pt-home-bush-b"></span>
        </div>

        <div class="pt-home-title">
          <span class="pt-home-hello">欢迎回家，</span>
          <h1>{{ payload.petName || '未命名' }}</h1>
          <p>{{ payload.breed }} · {{ payload.age || '—' }}</p>
        </div>

        <div class="pt-home-pills">
          <div class="pt-home-pill">
            <span class="pt-home-pill-icon" aria-hidden="true">⌂</span>
            <div>
              <span class="pt-home-pill-label">家人</span>
              <p>{{ payload.ownerName }}</p>
            </div>
          </div>
          <div class="pt-home-pill">
            <span class="pt-home-pill-icon" aria-hidden="true">◎</span>
            <div>
              <span class="pt-home-pill-label">住址</span>
              <p>{{ payload.address }}</p>
            </div>
          </div>
        </div>

        <div class="pt-home-med">
          <span class="pt-home-med-badge">健康小贴士</span>
          <p>{{ payload.medical || '无特殊记录' }}</p>
        </div>

        <div class="pt-home-note">
          <span class="pt-home-quote-mark" aria-hidden="true">“</span>
          <p>{{ payload.note || '有点怕生，请先蹲下来，让我闻闻你的手。' }}</p>
        </div>

        <a class="pt-home-phone" :href="'tel:' + payload.phone">
          <span class="pt-home-phone-icon" aria-hidden="true">
            <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8">
              <path d="M5 4h4l2 5-2.5 1.5a11 11 0 0 0 5 5L15 13l5 2v4a2 2 0 0 1-2 2A16 16 0 0 1 3 6a2 2 0 0 1 2-2z" />
            </svg>
          </span>
          <div class="pt-home-phone-text">
            <span>如果我不小心走远了，请打给</span>
            <strong>{{ payload.phone }}</strong>
          </div>
          <span class="pt-home-phone-arrow" aria-hidden="true">›</span>
        </a>

        <footer class="pt-home-foot">
          <span class="pt-home-paw" aria-hidden="true">🐾</span>
          <span>NFC 轻轻一碰 · 帮我找到家</span>
        </footer>
      </article>

      <p class="pt-home-below">愿每一次出门，都能平安回家。</p>
    </div>
  </div>
</template>

<script setup>
defineProps(['payload', 'nfcId'])
</script>

<style scoped>
.pt-home-root {
  background:
    radial-gradient(circle at 92% 9%, rgba(190, 202, 173, 0.35), transparent 30%),
    radial-gradient(circle at 4% 74%, rgba(240, 214, 184, 0.42), transparent 34%),
    #f6f1e8;
  font-family: 'Avenir Next', 'PingFang SC', 'Microsoft YaHei', sans-serif;
}
.pt-home-content { animation: pt-home-grow 0.8s cubic-bezier(0.2, 0.85, 0.3, 1) both; }
.pt-home-top { display: flex; justify-content: space-between; margin: 0 8px 13px; font-size: 9px; letter-spacing: 0.13em; color: #9b9483; }
.pt-home-brand { font-weight: 600; color: #7d8a6d; }
.pt-home-tag { font-family: 'Courier New', monospace; }
.pt-home-card { position: relative; padding: 17px 18px 16px; background: #fffdf8; border: 1px solid #e8e1d3; border-radius: 28px; box-shadow: 0 22px 42px rgba(122, 106, 78, 0.14); transition: transform 0.35s ease, box-shadow 0.35s ease; }
.pt-home-card:hover { transform: translateY(-4px); box-shadow: 0 30px 52px rgba(122, 106, 78, 0.18); }
.pt-home-ribbon { position: absolute; top: -13px; left: 31px; display: flex; flex-direction: column; align-items: center; }
.pt-home-ribbon span { width: 18px; height: 33px; background: #aebfa4; border-radius: 0 0 5px 5px; box-shadow: 0 3px 4px rgba(99, 113, 89, 0.22); }
.pt-home-ribbon span::before { content: ''; display: block; width: 18px; height: 18px; margin-top: -9px; border: 4px solid #c7d4bf; border-radius: 50%; box-sizing: border-box; background: #fffdf8; }
.pt-home-illustration { position: relative; height: 166px; overflow: hidden; border-radius: 20px; background: linear-gradient(180deg, #dce8e6 0%, #eef0e3 60%, #e3e6d1 100%); }
.pt-home-sun { position: absolute; top: 18px; right: 27px; width: 34px; height: 34px; border-radius: 50%; background: #f0c97f; box-shadow: 0 0 0 8px rgba(240, 201, 127, 0.16); }
.pt-home-cloud { position: absolute; height: 13px; border-radius: 20px; background: rgba(255, 255, 255, 0.72); }
.pt-home-cloud-a { top: 27px; left: 24px; width: 55px; }
.pt-home-cloud-b { top: 45px; left: 48px; width: 35px; opacity: 0.7; }
.pt-home-house { position: absolute; left: 50%; bottom: 25px; width: 114px; height: 82px; margin-left: -57px; background: #cf9d72; border-radius: 4px; box-shadow: inset -8px 0 0 rgba(122, 82, 45, 0.12); }
.pt-home-roof { position: absolute; left: -17px; top: -43px; width: 148px; height: 58px; background: #9c7350; clip-path: polygon(50% 0, 100% 100%, 0 100%); }
.pt-home-roof::after { content: ''; position: absolute; left: 17px; right: 17px; bottom: 0; height: 43px; clip-path: polygon(50% 0, 100% 100%, 0 100%); background: repeating-linear-gradient(90deg, #8a6545 0 9px, #a67d57 9px 18px); }
.pt-home-window { position: absolute; left: 17px; top: 20px; width: 25px; height: 25px; background: #f7e6bb; border: 4px solid #8a6545; border-radius: 3px; box-shadow: inset 0 0 0 1px #fff5d6; }
.pt-home-window::after { content: ''; position: absolute; left: 7px; top: -4px; bottom: -4px; width: 3px; background: #8a6545; }
.pt-home-door { position: absolute; right: 25px; bottom: 0; width: 29px; height: 48px; background: #7e5a3d; border-radius: 12px 12px 0 0; }
.pt-home-door::after { content: ''; position: absolute; left: 7px; top: 19px; width: 4px; height: 4px; border-radius: 50%; background: #e8c98f; }
.pt-home-hill { position: absolute; left: -25px; right: -25px; bottom: -37px; height: 80px; border-radius: 50%; background: #b9c7a2; }
.pt-home-bush { position: absolute; bottom: 10px; width: 37px; height: 23px; border-radius: 50% 50% 40% 40%; background: #96aa80; }
.pt-home-bush-a { left: 21px; }
.pt-home-bush-b { right: 19px; width: 29px; height: 18px; background: #a5b78d; }
.pt-home-title { padding: 17px 3px 12px; text-align: center; }
.pt-home-hello { font-size: 10px; letter-spacing: 0.2em; color: #a49a86; }
.pt-home-title h1 { margin: 2px 0 0; font-size: clamp(39px, 11.5vw, 52px); line-height: 1.1; font-weight: 500; letter-spacing: 0.09em; color: #4f5a44; }
.pt-home-title p { margin: 6px 0 0; font-size: 10px; letter-spacing: 0.12em; color: #a89d8b; }
.pt-home-pills { display: grid; gap: 7px; }
.pt-home-pill { display: flex; align-items: center; gap: 10px; padding: 9px 12px; border-radius: 14px; background: #f5f2e9; }
.pt-home-pill-icon { width: 26px; height: 26px; flex: none; display: grid; place-items: center; border-radius: 50%; background: #e3e9d8; color: #6d7a5d; font-size: 12px; }
.pt-home-pill-label { display: block; font-size: 8px; letter-spacing: 0.1em; color: #a49b89; }
.pt-home-pill p { margin: 1px 0 0; font-size: 11px; font-weight: 500; line-height: 1.45; color: #5e594d; }
.pt-home-med { display: flex; align-items: flex-start; gap: 8px; margin-top: 9px; padding: 9px 11px; border-radius: 11px; background: #fbf1d9; border: 1px solid #eeddb5; }
.pt-home-med-badge { flex: none; padding: 2px 5px; border-radius: 4px; background: #d8b56e; color: #fffaf0; font-size: 8px; letter-spacing: 0.04em; }
.pt-home-med p { margin: 0; font-size: 10px; line-height: 1.55; color: #7d663d; }
.pt-home-note { position: relative; margin: 10px 2px 12px; padding: 7px 10px 7px 25px; }
.pt-home-quote-mark { position: absolute; left: 4px; top: -2px; font-family: Georgia, serif; font-size: 30px; line-height: 1; color: #c8cdb9; }
.pt-home-note p { margin: 0; font-size: 11px; line-height: 1.7; color: #736d5f; }
.pt-home-phone { display: flex; align-items: center; gap: 10px; padding: 11px 12px; border-radius: 16px; background: #7d8a6d; color: #fdfbf5; text-decoration: none; box-shadow: 0 5px 0 #6a765c; transition: transform 0.18s ease, box-shadow 0.18s ease, background 0.18s ease; }
.pt-home-phone:hover { background: #6f7c60; transform: translateY(-1px); box-shadow: 0 6px 0 #5e6a50; }
.pt-home-phone:active { transform: translateY(3px); box-shadow: 0 2px 0 #6a765c; }
.pt-home-phone-icon { width: 30px; height: 30px; flex: none; display: grid; place-items: center; border-radius: 50%; background: rgba(255, 255, 255, 0.16); }
.pt-home-phone-icon svg { width: 16px; height: 16px; }
.pt-home-phone-text { flex: 1; min-width: 0; }
.pt-home-phone-text span { display: block; font-size: 8px; letter-spacing: 0.06em; color: #e4ead9; }
.pt-home-phone-text strong { display: block; margin-top: 1px; font-size: clamp(23px, 7vw, 28px); font-weight: 700; letter-spacing: 0.01em; }
.pt-home-phone-arrow { font-size: 24px; color: #d2ddc3; }
.pt-home-foot { display: flex; justify-content: center; gap: 6px; margin-top: 13px; font-size: 8px; letter-spacing: 0.1em; color: #b0a795; }
.pt-home-paw { font-size: 10px; }
.pt-home-below { margin: 19px 0 0; text-align: center; font-size: 10px; letter-spacing: 0.12em; color: #a49b89; }
@keyframes pt-home-grow {
  from { opacity: 0; transform: translateY(22px) scale(0.975); }
  to { opacity: 1; transform: translateY(0) scale(1); }
}
@media (prefers-reduced-motion: reduce) {
  .pt-home-content { animation: none; }
  .pt-home-card, .pt-home-phone { transition: none; }
}
</style>
