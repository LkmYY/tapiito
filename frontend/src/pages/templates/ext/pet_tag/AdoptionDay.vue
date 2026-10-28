<template>
  <div class="pt-adopt-root min-h-screen px-5 py-9 text-rose-950">
    <div class="pt-adopt-content mx-auto w-full max-w-md">
      <header class="pt-adopt-head">
        <span>ADOPTION DAY</span>
        <span>{{ nfcId ? 'NO.' + nfcId : 'FOREVER HOME' }}</span>
      </header>

      <article class="pt-adopt-card">
        <div class="pt-adopt-tape" aria-hidden="true"></div>

        <div class="pt-adopt-photo">
          <div class="pt-adopt-photo-inner">
            <span class="pt-adopt-sun" aria-hidden="true"></span>
            <div class="pt-adopt-cat" aria-hidden="true">
              <span class="pt-adopt-ear pt-adopt-ear-l"></span>
              <span class="pt-adopt-ear pt-adopt-ear-r"></span>
              <span class="pt-adopt-eye pt-adopt-eye-l"></span>
              <span class="pt-adopt-eye pt-adopt-eye-r"></span>
              <span class="pt-adopt-muzzle"></span>
              <span class="pt-adopt-nose"></span>
              <span class="pt-adopt-whisker pt-adopt-whisker-l"></span>
              <span class="pt-adopt-whisker pt-adopt-whisker-r"></span>
            </div>
            <span class="pt-adopt-ground" aria-hidden="true"></span>
            <p class="pt-adopt-cap">day one, forever ours</p>
          </div>
        </div>

        <div class="pt-adopt-stamp" aria-hidden="true">已领养<br /><small>ADOPTED</small></div>

        <div class="pt-adopt-intro">
          <p class="pt-adopt-eyebrow">我们家的新成员</p>
          <h1>{{ payload.petName || '未命名' }}</h1>
          <p class="pt-adopt-meta">{{ payload.breed }} · {{ payload.age || '—' }}</p>
        </div>

        <div class="pt-adopt-heartline" aria-hidden="true">
          <span></span>
          <svg viewBox="0 0 24 24" fill="currentColor"><path d="M12 21s-7.5-4.7-9.5-9A5.4 5.4 0 0 1 12 6.6 5.4 5.4 0 0 1 21.5 12c-2 4.3-9.5 9-9.5 9z" /></svg>
          <span></span>
        </div>

        <div class="pt-adopt-details">
          <div class="pt-adopt-person">
            <span class="pt-adopt-label">领养人 / GUARDIAN</span>
            <p>{{ payload.ownerName }}</p>
            <span class="pt-adopt-address">{{ payload.address }}</span>
          </div>
          <div class="pt-adopt-note">
            <span class="pt-adopt-label">小习惯 / NOTE</span>
            <p>{{ payload.note || '喜欢晒太阳，也喜欢被轻轻摸头。' }}</p>
          </div>
        </div>

        <div class="pt-adopt-medical">
          <span class="pt-adopt-medical-icon" aria-hidden="true">✚</span>
          <div>
            <span class="pt-adopt-label">健康备忘 / MEDICAL</span>
            <p>{{ payload.medical || '无特殊记录' }}</p>
          </div>
        </div>

        <a class="pt-adopt-phone" :href="'tel:' + payload.phone">
          <span class="pt-adopt-phone-label">想认识我？联系我的家人</span>
          <strong>{{ payload.phone }}</strong>
          <span class="pt-adopt-phone-arrow" aria-hidden="true">→</span>
        </a>

        <footer class="pt-adopt-foot">
          <span>EST. 2026 · 青禾社区</span>
          <span class="pt-adopt-sign">with love</span>
        </footer>
      </article>

      <p class="pt-adopt-below">领养不是一时心软，是一生的约定。</p>
    </div>
  </div>
</template>

<script setup>
defineProps(['payload', 'nfcId'])
</script>

<style scoped>
.pt-adopt-root {
  background:
    radial-gradient(circle at 82% 6%, rgba(251, 113, 133, 0.18), transparent 34%),
    radial-gradient(circle at 8% 82%, rgba(253, 186, 116, 0.2), transparent 38%),
    #fff4ea;
  font-family: 'Georgia', 'Songti SC', 'PingFang SC', serif;
}
.pt-adopt-content { animation: pt-adopt-fade 0.75s ease both; }
.pt-adopt-head { display: flex; justify-content: space-between; margin: 0 6px 12px; font-family: 'Courier New', monospace; font-size: 8px; letter-spacing: 0.18em; color: #a86f58; }
.pt-adopt-card {
  position: relative;
  padding: 18px 18px 20px;
  background: #fffaf4;
  border: 1px solid #f1d9c8;
  border-radius: 18px;
  box-shadow: 0 22px 44px rgba(146, 79, 53, 0.16);
  transition: transform 0.35s ease, box-shadow 0.35s ease;
}
.pt-adopt-card:hover { transform: translateY(-4px); box-shadow: 0 30px 54px rgba(146, 79, 53, 0.2); }
.pt-adopt-tape { position: absolute; top: -12px; left: 50%; width: 94px; height: 25px; margin-left: -47px; background: rgba(240, 189, 139, 0.5); transform: rotate(-2deg); box-shadow: 0 1px 2px rgba(120, 70, 40, 0.13); }
.pt-adopt-photo {
  padding: 8px 8px 26px;
  background: #fff;
  box-shadow: 0 7px 16px rgba(92, 48, 28, 0.18);
  transform: rotate(-2.4deg);
  transition: transform 0.4s cubic-bezier(0.25, 0.8, 0.3, 1);
  animation: pt-adopt-photo 0.85s 0.1s cubic-bezier(0.2, 0.9, 0.3, 1) both;
}
.pt-adopt-card:hover .pt-adopt-photo { transform: rotate(0.4deg); }
.pt-adopt-photo-inner {
  position: relative;
  height: 214px;
  overflow: hidden;
  background:
    radial-gradient(circle at 76% 18%, #ffe7b8 0 12%, transparent 12.5%),
    linear-gradient(180deg, #ffd9c9 0%, #f8b99f 58%, #e9a384 100%);
}
.pt-adopt-photo-inner::before { content: ''; position: absolute; inset: 0; background: radial-gradient(ellipse at 50% 110%, rgba(121, 57, 33, 0.28), transparent 60%); }
.pt-adopt-sun { position: absolute; top: 22px; right: 26px; width: 44px; height: 44px; border-radius: 50%; background: rgba(255, 250, 229, 0.7); filter: blur(1px); }
.pt-adopt-cat { position: absolute; left: 50%; bottom: 25px; width: 132px; height: 126px; margin-left: -66px; background: #6b4633; border-radius: 48% 48% 42% 42%; box-shadow: inset -15px -10px 0 rgba(41, 22, 14, 0.24); }
.pt-adopt-ear { position: absolute; top: -27px; width: 39px; height: 58px; background: #5f3c2b; border-radius: 50% 50% 12% 12%; }
.pt-adopt-ear::after { content: ''; position: absolute; inset: 12px 8px 16px; background: #c98f7b; border-radius: inherit; opacity: 0.55; }
.pt-adopt-ear-l { left: 3px; transform: rotate(-17deg); }
.pt-adopt-ear-r { right: 3px; transform: rotate(17deg); }
.pt-adopt-eye { position: absolute; top: 44px; width: 14px; height: 18px; border-radius: 50%; background: #1f1710; }
.pt-adopt-eye-l { left: 31px; }
.pt-adopt-eye-r { right: 31px; }
.pt-adopt-eye::after { content: ''; position: absolute; top: 3px; left: 4px; width: 5px; height: 5px; border-radius: 50%; background: #fff; opacity: 0.85; }
.pt-adopt-muzzle { position: absolute; left: 50%; bottom: 16px; width: 55px; height: 36px; margin-left: -27.5px; border-radius: 50%; background: #c79b7f; }
.pt-adopt-nose { position: absolute; left: 50%; bottom: 40px; width: 14px; height: 10px; margin-left: -7px; border-radius: 50% 50% 45% 45%; background: #4b2c1e; }
.pt-adopt-whisker { position: absolute; bottom: 30px; width: 36px; height: 1px; background: rgba(255, 241, 221, 0.7); }
.pt-adopt-whisker-l { left: -23px; transform: rotate(6deg); }
.pt-adopt-whisker-r { right: -23px; transform: rotate(-6deg); }
.pt-adopt-ground { position: absolute; left: 0; right: 0; bottom: 0; height: 30px; background: rgba(112, 57, 37, 0.16); filter: blur(8px); }
.pt-adopt-cap { position: absolute; left: 12px; bottom: 7px; margin: 0; font-family: 'Segoe Script', 'Bradley Hand', cursive; font-size: 12px; color: #7d4a35; transform: rotate(-2deg); }
.pt-adopt-stamp {
  position: absolute;
  top: 262px;
  right: 24px;
  z-index: 2;
  padding: 7px 9px;
  border: 2px solid #c65d55;
  border-radius: 50%;
  color: #c65d55;
  font-size: 10px;
  font-weight: 700;
  line-height: 1.2;
  text-align: center;
  transform: rotate(11deg);
  opacity: 0.75;
}
.pt-adopt-stamp small { font-family: 'Courier New', monospace; font-size: 6px; letter-spacing: 0.08em; }
.pt-adopt-intro { margin-top: 20px; }
.pt-adopt-eyebrow { margin: 0; font-size: 10px; letter-spacing: 0.24em; color: #b87962; }
.pt-adopt-intro h1 { margin: 4px 0 0; font-size: clamp(42px, 12vw, 55px); line-height: 1.02; font-weight: 500; letter-spacing: 0.08em; color: #522e22; }
.pt-adopt-meta { margin: 7px 0 0; font-family: 'Avenir Next', 'PingFang SC', sans-serif; font-size: 12px; letter-spacing: 0.12em; color: #966651; }
.pt-adopt-heartline { display: flex; align-items: center; gap: 8px; margin: 16px 0 14px; color: #e28e7d; }
.pt-adopt-heartline span { flex: 1; height: 1px; background: #eccab9; }
.pt-adopt-heartline svg { width: 15px; height: 15px; }
.pt-adopt-details { display: grid; grid-template-columns: 0.9fr 1.1fr; gap: 14px; }
.pt-adopt-label { display: block; font-family: 'Courier New', monospace; font-size: 7px; font-weight: 700; letter-spacing: 0.12em; color: #b7826c; margin-bottom: 4px; }
.pt-adopt-person p { margin: 0; font-family: 'Avenir Next', 'PingFang SC', sans-serif; font-size: 15px; font-weight: 600; color: #5a3426; }
.pt-adopt-address { display: block; margin-top: 3px; font-family: 'Avenir Next', 'PingFang SC', sans-serif; font-size: 10px; line-height: 1.5; color: #9a6d59; }
.pt-adopt-note p { margin: 0; font-size: 12px; line-height: 1.7; color: #6c4536; }
.pt-adopt-medical { display: flex; gap: 9px; margin-top: 14px; padding: 10px 11px; border: 1px dashed #dfb09a; border-radius: 8px; background: #fff5ed; }
.pt-adopt-medical-icon { flex: none; width: 19px; height: 19px; display: grid; place-items: center; border-radius: 50%; background: #d97963; color: #fff6ee; font-size: 11px; }
.pt-adopt-medical p { margin: 0; font-family: 'Avenir Next', 'PingFang SC', sans-serif; font-size: 11px; line-height: 1.55; color: #7b4434; }
.pt-adopt-phone {
  display: flex;
  flex-direction: column;
  position: relative;
  margin-top: 16px;
  padding: 13px 42px 13px 15px;
  border-radius: 13px;
  background: #5a3426;
  color: #fff6ec;
  text-decoration: none;
  transition: background 0.22s ease, transform 0.2s ease;
}
.pt-adopt-phone:hover { background: #703f2e; }
.pt-adopt-phone:active { transform: scale(0.985); }
.pt-adopt-phone-label { font-family: 'Avenir Next', 'PingFang SC', sans-serif; font-size: 9px; letter-spacing: 0.08em; color: #e6c1ac; }
.pt-adopt-phone strong { margin-top: 1px; font-family: 'Avenir Next', 'PingFang SC', sans-serif; font-size: clamp(24px, 7.5vw, 30px); font-weight: 700; letter-spacing: 0.02em; }
.pt-adopt-phone-arrow { position: absolute; right: 15px; top: 50%; transform: translateY(-50%); font-size: 23px; color: #f1b6a1; transition: transform 0.2s ease; }
.pt-adopt-phone:hover .pt-adopt-phone-arrow { transform: translate(3px, -50%); }
.pt-adopt-foot { display: flex; justify-content: space-between; align-items: baseline; margin-top: 17px; padding-top: 11px; border-top: 1px solid #f0d8c9; font-family: 'Courier New', monospace; font-size: 8px; color: #b18a76; }
.pt-adopt-sign { font-family: 'Segoe Script', 'Bradley Hand', cursive; font-size: 15px; color: #a96049; }
.pt-adopt-below { margin: 20px 0 0; text-align: center; font-size: 10px; letter-spacing: 0.12em; color: #a87a64; }
@keyframes pt-adopt-fade {
  from { opacity: 0; transform: translateY(18px); }
  to { opacity: 1; transform: translateY(0); }
}
@keyframes pt-adopt-photo {
  from { opacity: 0; transform: translateY(-24px) rotate(-7deg) scale(0.96); }
  to { opacity: 1; transform: translateY(0) rotate(-2.4deg) scale(1); }
}
@media (prefers-reduced-motion: reduce) {
  .pt-adopt-content, .pt-adopt-photo { animation: none; }
  .pt-adopt-card, .pt-adopt-photo, .pt-adopt-phone, .pt-adopt-phone-arrow { transition: none; }
}
</style>
