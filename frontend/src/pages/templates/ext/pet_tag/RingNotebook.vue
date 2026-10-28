<template>
  <div class="pt-note-root min-h-screen px-4 py-8 text-stone-800">
    <div class="pt-note-content mx-auto w-full max-w-md">
      <div class="pt-note-tabs" aria-hidden="true">
        <span class="pt-note-tab pt-note-tab-a">健康</span>
        <span class="pt-note-tab pt-note-tab-b">电话</span>
        <span class="pt-note-tab pt-note-tab-c">备忘</span>
      </div>

      <article class="pt-note-book">
        <div class="pt-note-rings" aria-hidden="true">
          <span v-for="n in 9" :key="n"></span>
        </div>

        <div class="pt-note-page">
          <div class="pt-note-margin" aria-hidden="true"></div>

          <header class="pt-note-head">
            <div>
              <p class="pt-note-mono">MY PET NOTES</p>
              <h1>{{ payload.petName || '未命名' }}</h1>
              <p class="pt-note-sub">{{ payload.breed }}<span>·</span>{{ payload.age || '—' }}</p>
            </div>
            <span class="pt-note-date">2026<br /><small>SPRING</small></span>
          </header>

          <div class="pt-note-doodle" aria-hidden="true">
            <svg viewBox="0 0 90 78" fill="none" stroke="currentColor" stroke-width="1.4" stroke-linecap="round">
              <path d="M26 55 C 21 30, 34 17, 46 17 C 60 17, 70 31, 64 53" />
              <path d="M27 38 C 19 28, 22 15, 32 17 C 38 18, 39 27, 35 33" />
              <path d="M62 36 C 69 25, 68 12, 58 15 C 52 17, 52 26, 56 32" />
              <circle cx="40" cy="42" r="1.5" fill="currentColor" />
              <circle cx="54" cy="42" r="1.5" fill="currentColor" />
              <path d="M46 48 C 45 51, 47 52, 49 50" />
              <path d="M34 56 C 44 61, 53 60, 61 55" />
            </svg>
          </div>

          <div class="pt-note-section">
            <span class="pt-note-label">家人 / OWNER</span>
            <p>{{ payload.ownerName }}</p>
            <span class="pt-note-label">住址 / ADDRESS</span>
            <p>{{ payload.address }}</p>
          </div>

          <div class="pt-note-redbox">
            <span class="pt-note-red-label">!! 健康备忘 / MEDICAL</span>
            <p>{{ payload.medical || '无特殊记录' }}</p>
          </div>

          <div class="pt-note-note">
            <span class="pt-note-highlight">它的小脾气：</span>
            <p>{{ payload.note || '怕生，请勿强行抱抱。' }}</p>
          </div>

          <a class="pt-note-phone" :href="'tel:' + payload.phone">
            <span class="pt-note-phone-top">
              <span class="pt-note-phone-mark" aria-hidden="true">☎</span>
              <span>走失时请拨这通电话</span>
            </span>
            <strong>{{ payload.phone }}</strong>
            <span class="pt-note-phone-bottom">— 谢谢每一个愿意停下来的人 —</span>
          </a>

          <div class="pt-note-check" aria-hidden="true">
            <span>✓ 已绝育</span>
            <span>✓ 已打芯片</span>
            <span>✓ 被爱着</span>
          </div>

          <footer class="pt-note-foot">
            <span>NFC {{ nfcId || '0086' }}</span>
            <span class="pt-note-sign">记得回家 ♡</span>
          </footer>
        </div>
      </article>

      <p class="pt-note-below">always write down what matters.</p>
    </div>
  </div>
</template>

<script setup>
defineProps(['payload', 'nfcId'])
</script>

<style scoped>
.pt-note-root {
  background:
    radial-gradient(circle at 85% 7%, rgba(96, 165, 250, 0.13), transparent 28%),
    repeating-linear-gradient(45deg, rgba(120, 100, 70, 0.035) 0 1px, transparent 1px 9px),
    #e9e0d0;
  font-family: 'Avenir Next', 'PingFang SC', sans-serif;
}
.pt-note-content { position: relative; animation: pt-note-open 0.8s cubic-bezier(0.2, 0.85, 0.3, 1) both; transform-origin: left center; }
.pt-note-tabs { display: flex; gap: 3px; padding-left: 35px; margin-bottom: -1px; }
.pt-note-tab { padding: 4px 10px 3px; border-radius: 6px 6px 0 0; font-size: 8px; letter-spacing: 0.08em; color: #5a5142; position: relative; z-index: 1; }
.pt-note-tab-a { background: #c9d8bf; }
.pt-note-tab-b { background: #e8c1a8; }
.pt-note-tab-c { background: #e5d28e; }
.pt-note-book { position: relative; display: flex; padding-left: 24px; filter: drop-shadow(0 20px 26px rgba(91, 75, 47, 0.2)); }
.pt-note-rings { position: absolute; z-index: 4; top: 18px; bottom: 18px; left: 2px; display: flex; flex-direction: column; justify-content: space-around; }
.pt-note-rings span { display: block; width: 35px; height: 9px; border-radius: 6px; border: 1.5px solid #8e8a80; background: linear-gradient(180deg, #f2f0ea, #b9b4a8 55%, #dedbd2); box-shadow: 0 1px 2px rgba(70, 60, 45, 0.3); }
.pt-note-rings span::after { content: ''; display: block; width: 6px; height: 100%; margin-left: 4px; border-right: 1px solid rgba(90, 84, 72, 0.25); }
.pt-note-page {
  position: relative;
  flex: 1;
  padding: 20px 19px 16px 27px;
  background:
    repeating-linear-gradient(0deg, transparent 0 27px, rgba(94, 128, 164, 0.18) 27px 28px),
    #fcfaf3;
  border: 1px solid #d8cfbe;
  border-radius: 2px 9px 9px 2px;
  box-shadow: inset 12px 0 18px -14px rgba(83, 68, 43, 0.35);
  transition: transform 0.35s ease;
}
.pt-note-book:hover .pt-note-page { transform: translateY(-2px); }
.pt-note-margin { position: absolute; top: 0; bottom: 0; left: 20px; width: 1px; background: rgba(215, 120, 105, 0.45); }
.pt-note-head { position: relative; display: flex; justify-content: space-between; gap: 8px; padding-bottom: 8px; }
.pt-note-mono { margin: 0 0 2px; font-family: 'Courier New', monospace; font-size: 6.5px; font-weight: 700; letter-spacing: 0.16em; color: #9a9282; }
.pt-note-head h1 { margin: 0; max-width: 230px; font-family: 'Segoe Script', 'Bradley Hand', 'Kaiti SC', cursive; font-size: clamp(39px, 11.5vw, 51px); line-height: 1.08; font-weight: 400; letter-spacing: 0.01em; color: #2f4b66; word-break: break-word; }
.pt-note-sub { margin: 5px 0 0; font-family: 'Courier New', monospace; font-size: 8px; letter-spacing: 0.06em; color: #8e8676; }
.pt-note-sub span { margin: 0 5px; color: #b8afa0; }
.pt-note-date { flex: none; align-self: flex-start; padding: 3px 6px; border: 1px solid #c9a46d; color: #a37d43; font-family: 'Courier New', monospace; font-size: 8px; line-height: 1.2; text-align: center; transform: rotate(4deg); }
.pt-note-date small { font-size: 5px; letter-spacing: 0.08em; }
.pt-note-doodle { position: absolute; top: 66px; right: 12px; width: 58px; color: #8a9eb0; opacity: 0.72; transform: rotate(3deg); }
.pt-note-doodle svg { display: block; width: 100%; height: auto; }
.pt-note-section { position: relative; padding: 7px 0 5px; }
.pt-note-label { display: block; margin-top: 6px; font-family: 'Courier New', monospace; font-size: 6.5px; font-weight: 700; letter-spacing: 0.1em; color: #a09785; }
.pt-note-section p { margin: 1px 0 0; max-width: 78%; font-size: 11px; font-weight: 500; line-height: 1.4; color: #4b5663; }
.pt-note-redbox { position: relative; margin: 9px 0 8px; padding: 8px 10px; border: 1.5px solid #c46a55; background: rgba(250, 234, 226, 0.75); transform: rotate(-0.4deg); }
.pt-note-red-label { display: block; margin-bottom: 2px; font-family: 'Courier New', monospace; font-size: 6.5px; font-weight: 700; letter-spacing: 0.08em; color: #b14a34; }
.pt-note-redbox p { margin: 0; font-size: 11px; line-height: 1.45; color: #8b3b2b; }
.pt-note-note { position: relative; padding: 7px 0 9px; }
.pt-note-highlight { display: inline-block; font-family: 'Kaiti SC', 'KaiTi', serif; font-size: 12px; color: #5a5142; background: linear-gradient(transparent 52%, rgba(246, 214, 110, 0.7) 52%); padding: 0 2px; }
.pt-note-note p { margin: 3px 0 0; font-family: 'Kaiti SC', 'KaiTi', serif; font-size: 12px; line-height: 1.65; color: #575045; }
.pt-note-phone { display: block; position: relative; margin-top: 2px; padding: 10px 12px 11px; background: #f5e9b6; border: 1px solid #c9b26c; box-shadow: 2px 3px 0 rgba(144, 125, 74, 0.22); color: #3f3a2b; text-decoration: none; transform: rotate(0.5deg); transition: transform 0.2s ease, box-shadow 0.2s ease; }
.pt-note-phone:hover { transform: rotate(0) translateY(-1px); box-shadow: 3px 5px 0 rgba(144, 125, 74, 0.26); }
.pt-note-phone:active { transform: translate(1px, 2px); box-shadow: 0 1px 0 rgba(144, 125, 74, 0.22); }
.pt-note-phone-top { display: flex; align-items: center; gap: 5px; font-family: 'PingFang SC', 'Microsoft YaHei', sans-serif; font-size: 8px; letter-spacing: 0.06em; color: #89763e; }
.pt-note-phone-mark { font-size: 11px; }
.pt-note-phone strong { display: block; margin: 1px 0; font-family: 'Avenir Next', 'PingFang SC', sans-serif; font-size: clamp(26px, 7.9vw, 33px); font-weight: 700; letter-spacing: 0.015em; color: #333025; }
.pt-note-phone-bottom { font-family: 'Kaiti SC', 'KaiTi', serif; font-size: 9px; letter-spacing: 0.04em; color: #8c7a49; }
.pt-note-check { display: flex; flex-wrap: wrap; gap: 5px 11px; margin-top: 11px; padding-top: 7px; border-top: 1px dashed #c9c1b1; }
.pt-note-check span { font-family: 'Kaiti SC', 'KaiTi', serif; font-size: 9px; color: #7e8a77; }
.pt-note-foot { display: flex; justify-content: space-between; align-items: baseline; margin-top: 11px; font-family: 'Courier New', monospace; font-size: 6.5px; letter-spacing: 0.08em; color: #a49b8b; }
.pt-note-sign { font-family: 'Kaiti SC', 'KaiTi', serif; font-size: 11px; color: #728294; }
.pt-note-below { margin: 17px 0 0; text-align: center; font-family: 'Courier New', monospace; font-size: 8px; letter-spacing: 0.09em; color: #958a77; }
@keyframes pt-note-open {
  0% { opacity: 0; transform: perspective(900px) rotateY(-11deg) translateX(-14px); }
  100% { opacity: 1; transform: perspective(900px) rotateY(0) translateX(0); }
}
@media (prefers-reduced-motion: reduce) {
  .pt-note-content { animation: none; }
  .pt-note-page, .pt-note-phone { transition: none; }
}
</style>
