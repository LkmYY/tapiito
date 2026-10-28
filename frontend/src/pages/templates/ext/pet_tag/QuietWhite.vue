<template>
  <div class="pt-white-root min-h-screen px-6 py-10 text-stone-900">
    <div class="pt-white-content mx-auto flex min-h-[calc(100vh-5rem)] w-full max-w-md flex-col">
      <header class="pt-white-head">
        <span>PET TAG</span>
        <span>{{ nfcId ? 'NO.' + nfcId : 'EDITION 01' }}</span>
      </header>

      <main class="pt-white-main">
        <div class="pt-white-paw" aria-hidden="true">
          <svg viewBox="0 0 48 48" fill="none" stroke="currentColor" stroke-width="1.4">
            <ellipse cx="24" cy="31" rx="9" ry="7" />
            <ellipse cx="11" cy="21" rx="3.6" ry="4.8" />
            <ellipse cx="21" cy="15" rx="3.6" ry="5.2" />
            <ellipse cx="32" cy="16" rx="3.6" ry="4.8" />
            <ellipse cx="40" cy="24" rx="3.2" ry="4.2" />
          </svg>
        </div>

        <p class="pt-white-eyebrow">This is</p>
        <h1 class="pt-white-name">{{ payload.petName || '未命名' }}</h1>
        <p class="pt-white-desc">
          {{ payload.breed }}<span class="pt-white-sep">·</span>{{ payload.age || '—' }}
        </p>

        <div class="pt-white-contact">
          <span class="pt-white-label">如果你的手边有我，请拨</span>
          <a class="pt-white-phone" :href="'tel:' + payload.phone">{{ payload.phone }}</a>
          <span class="pt-white-sub">电话 24 小时为这只小家伙开着</span>
        </div>

        <div class="pt-white-rows">
          <div class="pt-white-row">
            <span>主人</span>
            <p>{{ payload.ownerName }}</p>
          </div>
          <div class="pt-white-row">
            <span>住址</span>
            <p>{{ payload.address }}</p>
          </div>
          <div class="pt-white-row">
            <span>健康</span>
            <p>{{ payload.medical || '无特殊记录' }}</p>
          </div>
        </div>

        <blockquote class="pt-white-quote">
          {{ payload.note || '怕生，请勿强行抱抱。让我先闻闻你。' }}
        </blockquote>
      </main>

      <footer class="pt-white-foot">
        <span>NFC · TAP TO CONNECT</span>
        <span class="pt-white-dot" aria-hidden="true"></span>
        <span>MADE FOR SAFE RETURNS</span>
      </footer>
    </div>
  </div>
</template>

<script setup>
defineProps(['payload', 'nfcId'])
</script>

<style scoped>
.pt-white-root {
  background: #fbfaf7;
  font-family: 'Helvetica Neue', 'PingFang SC', 'Microsoft YaHei', sans-serif;
  color: #1c1917;
}
.pt-white-content { animation: pt-white-fade 0.9s ease both; }
.pt-white-head { display: flex; justify-content: space-between; padding-bottom: 18px; border-bottom: 1px solid #e4e1da; font-size: 9px; letter-spacing: 0.2em; color: #a8a29e; }
.pt-white-main { flex: 1; padding-top: 40px; }
.pt-white-paw { width: 30px; margin-bottom: 28px; color: #9ca38a; }
.pt-white-paw svg { display: block; width: 100%; }
.pt-white-eyebrow { margin: 0; font-family: 'Times New Roman', 'Songti SC', serif; font-size: 15px; font-style: italic; color: #a8a29e; }
.pt-white-name { margin: 4px 0 0; font-size: clamp(64px, 19vw, 86px); line-height: 0.95; font-weight: 300; letter-spacing: -0.055em; color: #1c1917; }
.pt-white-desc { margin: 17px 0 0; font-size: 11px; letter-spacing: 0.16em; color: #78716c; }
.pt-white-sep { margin: 0 8px; color: #b9b3ab; }
.pt-white-contact { margin-top: 46px; padding: 23px 0 21px; border-top: 1px solid #d9d6ce; border-bottom: 1px solid #d9d6ce; }
.pt-white-label { display: block; font-size: 10px; letter-spacing: 0.13em; color: #78716c; }
.pt-white-phone {
  display: inline-block;
  margin: 7px 0 5px;
  font-family: 'Helvetica Neue', 'Arial', sans-serif;
  font-size: clamp(32px, 9.8vw, 42px);
  font-weight: 400;
  letter-spacing: -0.02em;
  color: #1c1917;
  text-decoration: none;
  border-bottom: 1px solid transparent;
  transition: border-color 0.25s ease, color 0.25s ease;
}
.pt-white-phone:hover { color: #6b7355; border-bottom-color: #6b7355; }
.pt-white-phone:active { color: #3f4433; }
.pt-white-sub { display: block; font-family: 'Times New Roman', 'Songti SC', serif; font-size: 12px; font-style: italic; color: #a8a29e; }
.pt-white-rows { margin-top: 34px; }
.pt-white-row { display: grid; grid-template-columns: 52px 1fr; gap: 14px; padding: 11px 0; border-bottom: 1px solid #eeece6; align-items: baseline; }
.pt-white-row span { font-size: 9px; letter-spacing: 0.14em; color: #a8a29e; }
.pt-white-row p { margin: 0; font-size: 13px; line-height: 1.6; color: #3f3b36; }
.pt-white-quote { margin: 34px 0 0; padding: 0 0 0 17px; border-left: 2px solid #c9cdbd; font-family: 'Times New Roman', 'Songti SC', serif; font-size: 15px; font-style: italic; line-height: 1.85; color: #57534e; }
.pt-white-foot { display: flex; align-items: center; gap: 9px; padding-top: 22px; border-top: 1px solid #e4e1da; font-size: 7px; letter-spacing: 0.17em; color: #b3ada4; }
.pt-white-dot { width: 3px; height: 3px; border-radius: 50%; background: #b3ada4; }
@keyframes pt-white-fade {
  from { opacity: 0; transform: translateY(16px); }
  to { opacity: 1; transform: translateY(0); }
}
@media (prefers-reduced-motion: reduce) {
  .pt-white-content { animation: none; }
  .pt-white-phone { transition: none; }
}
</style>
