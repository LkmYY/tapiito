<template>
  <div class="pt-lost-root min-h-screen px-4 py-8 text-stone-950">
    <div class="pt-lost-content mx-auto w-full max-w-md">
      <div class="pt-lost-tape pt-lost-tape-l" aria-hidden="true"></div>
      <div class="pt-lost-tape pt-lost-tape-r" aria-hidden="true"></div>

      <article class="pt-lost-poster">
        <header class="pt-lost-top">
          <span class="pt-lost-small">MISSING PET</span>
          <span class="pt-lost-small">走失寻宠</span>
        </header>

        <div class="pt-lost-headline">
          <p class="pt-lost-alert">请帮我回家</p>
          <h1>走失<br /><span>寻宠</span></h1>
          <div class="pt-lost-paw" aria-hidden="true">
            <svg viewBox="0 0 60 60" fill="currentColor">
              <ellipse cx="30" cy="39" rx="14" ry="11" />
              <ellipse cx="13" cy="26" rx="5.5" ry="7" />
              <ellipse cx="26" cy="18" rx="5.5" ry="7.5" />
              <ellipse cx="39" cy="19" rx="5.5" ry="7" />
              <ellipse cx="50" cy="29" rx="4.6" ry="6" />
            </svg>
          </div>
        </div>

        <div class="pt-lost-callout">
          <span class="pt-lost-call-label">见到请立刻拨打</span>
          <a :href="'tel:' + payload.phone">{{ payload.phone }}</a>
          <span class="pt-lost-call-sub">24 小时接听 · 必有重谢</span>
        </div>

        <div class="pt-lost-body">
          <div class="pt-lost-photo">
            <div class="pt-lost-photo-inner">
              <span class="pt-lost-photo-label">LAST SEEN</span>
              <div class="pt-lost-dog" aria-hidden="true">
                <span class="pt-lost-ear pt-lost-ear-l"></span>
                <span class="pt-lost-ear pt-lost-ear-r"></span>
                <span class="pt-lost-eye pt-lost-eye-l"></span>
                <span class="pt-lost-eye pt-lost-eye-r"></span>
                <span class="pt-lost-snout"></span>
                <span class="pt-lost-nose"></span>
              </div>
              <p class="pt-lost-photo-cap">如果你见过我<br />请一定联系她</p>
            </div>
          </div>

          <div class="pt-lost-facts">
            <div class="pt-lost-name">
              <h2>{{ payload.petName || '未命名' }}</h2>
              <p>{{ payload.breed }} / {{ payload.age || '—' }}</p>
            </div>
            <div class="pt-lost-line">
              <span>主人</span>
              <p>{{ payload.ownerName }}</p>
            </div>
            <div class="pt-lost-line">
              <span>住址</span>
              <p>{{ payload.address }}</p>
            </div>
            <div class="pt-lost-med">
              <span class="pt-lost-med-tag">健康</span>
              <p>{{ payload.medical || '无特殊记录' }}</p>
            </div>
          </div>
        </div>

        <p class="pt-lost-note">{{ payload.note || '怕生，请勿强行抱抱。请先联系主人。' }}</p>

        <div class="pt-lost-tabs" aria-hidden="true">
          <span v-for="n in 4" :key="n">{{ payload.phone || '138 0013 8000' }}</span>
        </div>
        <p class="pt-lost-tab-hint">沿虚线撕下 · 随身携带</p>
      </article>

      <p class="pt-lost-foot">EVERY SHARE HELPS · 每一次转发都有意义</p>
    </div>
  </div>
</template>

<script setup>
defineProps(['payload', 'nfcId'])
</script>

<style scoped>
.pt-lost-root {
  background:
    repeating-linear-gradient(90deg, rgba(255, 255, 255, 0.025) 0 2px, transparent 2px 7px),
    radial-gradient(circle at 80% 8%, rgba(255, 214, 10, 0.16), transparent 40%),
    #1c1b19;
  font-family: 'Arial Black', 'PingFang SC', 'Microsoft YaHei', sans-serif;
}
.pt-lost-content { position: relative; animation: pt-lost-slam 0.72s cubic-bezier(0.16, 0.9, 0.28, 1.15) both; }
.pt-lost-tape {
  position: absolute;
  z-index: 3;
  top: -10px;
  width: 78px;
  height: 24px;
  background: rgba(245, 235, 205, 0.6);
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.2);
  backdrop-filter: blur(1px);
}
.pt-lost-tape-l { left: 13px; transform: rotate(-8deg); }
.pt-lost-tape-r { right: 13px; transform: rotate(7deg); }
.pt-lost-poster {
  position: relative;
  padding: 20px 18px 0;
  background:
    radial-gradient(circle at 18% 12%, rgba(120, 90, 30, 0.09), transparent 30%),
    linear-gradient(165deg, #fffdf2 0%, #f6efcf 100%);
  border: 3px solid #14130f;
  box-shadow: 0 24px 44px rgba(0, 0, 0, 0.55);
  transition: transform 0.35s ease;
}
.pt-lost-poster:hover { transform: rotate(-0.35deg) translateY(-3px); }
.pt-lost-top { display: flex; justify-content: space-between; padding-bottom: 8px; border-bottom: 2px solid #14130f; }
.pt-lost-small { font-family: 'Courier New', monospace; font-size: 9px; font-weight: 700; letter-spacing: 0.2em; }
.pt-lost-headline { position: relative; display: flex; align-items: center; gap: 8px; padding: 10px 0 4px; }
.pt-lost-alert {
  align-self: flex-start;
  margin: 12px 0 0;
  writing-mode: vertical-rl;
  font-size: 11px;
  letter-spacing: 0.3em;
  font-weight: 700;
  color: #f7f2df;
  background: #c81e1e;
  padding: 10px 4px;
}
.pt-lost-headline h1 {
  flex: 1;
  margin: 0;
  font-size: clamp(60px, 19vw, 82px);
  line-height: 0.82;
  letter-spacing: -0.04em;
  font-weight: 900;
}
.pt-lost-headline h1 span {
  color: #c81e1e;
  -webkit-text-stroke: 2px #14130f;
  text-shadow: 3px 3px 0 #14130f;
}
.pt-lost-paw { width: 58px; color: #e5b800; transform: rotate(12deg); }
.pt-lost-paw svg { display: block; width: 100%; height: auto; }
.pt-lost-callout {
  margin: 6px -18px 16px;
  padding: 11px 18px 13px;
  background: #ffe100;
  border-top: 3px solid #14130f;
  border-bottom: 3px solid #14130f;
  text-align: center;
}
.pt-lost-call-label { display: block; font-size: 10px; font-weight: 700; letter-spacing: 0.24em; }
.pt-lost-callout a {
  display: block;
  margin: 1px 0 2px;
  font-size: clamp(32px, 9.6vw, 42px);
  line-height: 1.08;
  font-weight: 900;
  letter-spacing: -0.02em;
  color: #14130f;
  text-decoration: none;
  transition: transform 0.18s ease, color 0.18s ease;
}
.pt-lost-callout a:hover { color: #c81e1e; transform: scale(1.015); }
.pt-lost-callout a:active { transform: scale(0.98); }
.pt-lost-call-sub { font-size: 9px; font-weight: 700; letter-spacing: 0.14em; }
.pt-lost-body { display: grid; grid-template-columns: 0.9fr 1.1fr; gap: 14px; }
.pt-lost-photo { background: #14130f; padding: 6px 6px 0; }
.pt-lost-photo-inner { position: relative; min-height: 196px; background: linear-gradient(165deg, #ddd2b2, #b7a987); display: flex; flex-direction: column; align-items: center; justify-content: center; overflow: hidden; }
.pt-lost-photo-inner::after { content: ''; position: absolute; inset: 0; background: repeating-linear-gradient(0deg, transparent 0 3px, rgba(20, 19, 15, 0.05) 3px 4px); }
.pt-lost-photo-label { position: absolute; top: 5px; left: 6px; z-index: 2; font-family: 'Courier New', monospace; font-size: 7px; letter-spacing: 0.15em; background: #14130f; color: #f6efcf; padding: 2px 4px; }
.pt-lost-dog { position: relative; width: 88px; height: 94px; margin-top: 4px; background: #8a5a33; border-radius: 46% 46% 40% 40%; box-shadow: inset -8px -8px 0 rgba(58, 32, 12, 0.22); }
.pt-lost-ear { position: absolute; top: -17px; width: 27px; height: 47px; background: #6c4224; border-radius: 50% 50% 40% 40%; }
.pt-lost-ear-l { left: -4px; transform: rotate(-19deg); }
.pt-lost-ear-r { right: -4px; transform: rotate(19deg); }
.pt-lost-eye { position: absolute; top: 34px; width: 9px; height: 11px; border-radius: 50%; background: #19120c; }
.pt-lost-eye-l { left: 22px; }
.pt-lost-eye-r { right: 22px; }
.pt-lost-snout { position: absolute; left: 50%; bottom: 12px; width: 42px; height: 32px; margin-left: -21px; border-radius: 50%; background: #e9d7b4; }
.pt-lost-nose { position: absolute; left: 50%; bottom: 30px; width: 13px; height: 9px; margin-left: -6.5px; border-radius: 50% 50% 55% 55%; background: #19120c; }
.pt-lost-photo-cap { position: relative; z-index: 2; margin: 8px 0 9px; text-align: center; font-size: 9px; line-height: 1.6; font-weight: 700; letter-spacing: 0.05em; color: #3b2b18; }
.pt-lost-facts { display: flex; flex-direction: column; gap: 8px; }
.pt-lost-name { padding-bottom: 7px; border-bottom: 2px solid #14130f; }
.pt-lost-name h2 { margin: 0; font-size: 33px; line-height: 1; letter-spacing: -0.03em; }
.pt-lost-name p { margin: 4px 0 0; font-size: 10px; font-weight: 700; letter-spacing: 0.08em; }
.pt-lost-line { display: grid; grid-template-columns: 34px 1fr; gap: 6px; padding-bottom: 6px; border-bottom: 1px dashed rgba(20, 19, 15, 0.35); }
.pt-lost-line span { font-size: 9px; font-weight: 700; letter-spacing: 0.08em; color: #6d6250; padding-top: 2px; }
.pt-lost-line p { margin: 0; font-size: 12px; font-weight: 700; line-height: 1.45; }
.pt-lost-med { display: flex; gap: 6px; align-items: flex-start; padding: 7px; background: #f3d2c8; border: 1.5px solid #c81e1e; }
.pt-lost-med-tag { flex: none; background: #c81e1e; color: #fff8e8; font-size: 8px; font-weight: 700; padding: 2px 4px; letter-spacing: 0.1em; }
.pt-lost-med p { margin: 0; font-size: 10px; line-height: 1.5; font-weight: 700; color: #6f1717; }
.pt-lost-note { margin: 14px 0 12px; padding: 9px 10px; border-left: 4px solid #c81e1e; background: rgba(200, 30, 30, 0.06); font-size: 11px; line-height: 1.65; font-weight: 700; }
.pt-lost-tabs {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 0;
  margin: 0 -18px;
  border-top: 2px dashed #14130f;
}
.pt-lost-tabs span {
  padding: 9px 2px 8px;
  text-align: center;
  font-size: 9px;
  font-weight: 700;
  letter-spacing: 0.01em;
  border-right: 1px dashed #14130f;
  background: #ffe100;
  writing-mode: vertical-rl;
  text-orientation: mixed;
  height: 76px;
}
.pt-lost-tabs span:last-child { border-right: 0; }
.pt-lost-tab-hint { margin: 0; padding: 6px 0 10px; text-align: center; font-size: 8px; letter-spacing: 0.18em; font-weight: 700; color: #6e6553; }
.pt-lost-foot { margin: 17px 0 0; text-align: center; font-size: 8px; letter-spacing: 0.2em; color: #b5ac96; }
@keyframes pt-lost-slam {
  0% { opacity: 0; transform: translateY(-40px) rotate(2.4deg) scale(1.04); }
  70% { opacity: 1; transform: translateY(4px) rotate(-0.5deg) scale(1); }
  100% { opacity: 1; transform: translateY(0) rotate(0) scale(1); }
}
@media (max-width: 380px) {
  .pt-lost-body { grid-template-columns: 1fr; }
  .pt-lost-photo-inner { min-height: 150px; }
}
@media (prefers-reduced-motion: reduce) {
  .pt-lost-content { animation: none; }
  .pt-lost-poster, .pt-lost-callout a { transition: none; }
}
</style>
