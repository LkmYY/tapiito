<template>
  <div class="pt-ink-root min-h-screen px-5 py-9 text-stone-900">
    <div class="pt-ink-content mx-auto w-full max-w-md">
      <header class="pt-ink-head">
        <span>墨 缘 · INK &amp; PAW</span>
        <span>第 {{ nfcId || '八六' }} 号</span>
      </header>

      <article class="pt-ink-paper">
        <div class="pt-ink-wash" aria-hidden="true"></div>
        <span class="pt-ink-brush pt-ink-brush-a" aria-hidden="true"></span>
        <span class="pt-ink-brush pt-ink-brush-b" aria-hidden="true"></span>

        <div class="pt-ink-title">
          <span class="pt-ink-vertical">此心安处是吾家</span>
          <p class="pt-ink-eyebrow">走失时，请替我捎一句话回家</p>
          <h1>{{ payload.petName || '未命名' }}</h1>
          <p class="pt-ink-meta">{{ payload.breed }} · {{ payload.age || '—' }}</p>
        </div>

        <div class="pt-ink-seal" aria-hidden="true">
          <span>平安</span>
          <small>PET</small>
        </div>

        <div class="pt-ink-rule" aria-hidden="true">
          <span></span>
          <i>印</i>
          <span></span>
        </div>

        <div class="pt-ink-info">
          <div class="pt-ink-row">
            <span>主人</span>
            <p>{{ payload.ownerName }}</p>
          </div>
          <div class="pt-ink-row">
            <span>住处</span>
            <p>{{ payload.address }}</p>
          </div>
        </div>

        <div class="pt-ink-med">
          <span class="pt-ink-med-key">医嘱</span>
          <p>{{ payload.medical || '无特殊记录' }}</p>
        </div>

        <div class="pt-ink-note">
          <span class="pt-ink-note-mark" aria-hidden="true">“</span>
          <p>{{ payload.note || '怕生，请勿强行抱抱。先让我闻闻你的手。' }}</p>
        </div>

        <a class="pt-ink-phone" :href="'tel:' + payload.phone">
          <span class="pt-ink-phone-label">若见字，请拨此号</span>
          <strong>{{ payload.phone }}</strong>
          <span class="pt-ink-phone-sub">一通电话，便是一段归途</span>
        </a>

        <footer class="pt-ink-foot">
          <span class="pt-ink-signature">青禾主人 书</span>
          <span class="pt-ink-date">岁在丙午 · 春</span>
        </footer>
      </article>

      <p class="pt-ink-below">一纸薄名 · 万般牵挂</p>
    </div>
  </div>
</template>

<script setup>
defineProps(['payload', 'nfcId'])
</script>

<style scoped>
.pt-ink-root {
  background:
    radial-gradient(circle at 20% 10%, rgba(120, 110, 95, 0.1), transparent 27%),
    radial-gradient(circle at 82% 88%, rgba(120, 110, 95, 0.08), transparent 30%),
    #eeeadf;
  font-family: 'Kaiti SC', 'STKaiti', 'KaiTi', 'Songti SC', serif;
}
.pt-ink-content { animation: pt-ink-bloom 1s ease both; }
.pt-ink-head { display: flex; justify-content: space-between; margin: 0 6px 12px; font-size: 9px; letter-spacing: 0.18em; color: #867d6b; }
.pt-ink-paper {
  position: relative;
  padding: 27px 22px 19px;
  background:
    radial-gradient(circle at 75% 18%, rgba(50, 42, 30, 0.055), transparent 26%),
    repeating-linear-gradient(0deg, transparent 0 33px, rgba(86, 74, 55, 0.05) 33px 34px),
    #f8f4e9;
  border: 1px solid #d6cdb8;
  box-shadow: 0 22px 38px rgba(74, 65, 48, 0.16), inset 0 0 44px rgba(120, 103, 74, 0.08);
  overflow: hidden;
  transition: transform 0.4s ease, box-shadow 0.4s ease;
}
.pt-ink-paper:hover { transform: translateY(-3px); box-shadow: 0 28px 48px rgba(74, 65, 48, 0.2); }
.pt-ink-paper::before { content: ''; position: absolute; inset: 9px; border: 1px solid rgba(112, 98, 73, 0.18); pointer-events: none; }
.pt-ink-wash { position: absolute; top: -42px; right: -55px; width: 195px; height: 175px; border-radius: 44% 56% 62% 38% / 52% 38% 62% 48%; background: radial-gradient(circle at 38% 45%, rgba(44, 39, 31, 0.14), rgba(44, 39, 31, 0.03) 60%, transparent 72%); filter: blur(3px); }
.pt-ink-brush { position: absolute; display: block; background: #2b2620; opacity: 0.55; }
.pt-ink-brush-a { width: 1px; height: 105px; top: 44px; left: 31px; transform: rotate(2deg); }
.pt-ink-brush-b { width: 30px; height: 1px; top: 129px; left: 16px; transform: rotate(-6deg); }
.pt-ink-title { position: relative; padding-left: 45px; }
.pt-ink-vertical { position: absolute; left: 4px; top: 2px; writing-mode: vertical-rl; font-size: 10px; letter-spacing: 0.28em; color: #8b7e66; }
.pt-ink-eyebrow { margin: 0; font-size: 10px; letter-spacing: 0.1em; color: #9c8f76; }
.pt-ink-title h1 { margin: 5px 0 0; font-size: clamp(55px, 16vw, 73px); line-height: 0.98; font-weight: 400; letter-spacing: 0.09em; color: #241f18; }
.pt-ink-meta { margin: 9px 0 0; font-family: 'PingFang SC', 'Microsoft YaHei', sans-serif; font-size: 10px; letter-spacing: 0.14em; color: #8c7d64; }
.pt-ink-seal { position: absolute; top: 32px; right: 27px; width: 48px; height: 48px; display: flex; flex-direction: column; align-items: center; justify-content: center; gap: 1px; border: 2px solid #b4472f; color: #b4472f; transform: rotate(6deg); opacity: 0.78; }
.pt-ink-seal span { font-size: 13px; letter-spacing: 0.08em; line-height: 1; }
.pt-ink-seal small { font-family: 'Courier New', monospace; font-size: 5px; letter-spacing: 0.14em; }
.pt-ink-rule { display: flex; align-items: center; gap: 10px; margin: 21px 0 13px; color: #8e8069; }
.pt-ink-rule span { flex: 1; height: 1px; background: linear-gradient(90deg, transparent, rgba(83, 72, 53, 0.45), transparent); }
.pt-ink-rule i { font-style: normal; font-size: 9px; padding: 2px 4px; border: 1px solid rgba(131, 112, 79, 0.5); }
.pt-ink-info { display: grid; grid-template-columns: 0.8fr 1.2fr; gap: 11px; }
.pt-ink-row span { display: block; margin-bottom: 2px; font-size: 9px; letter-spacing: 0.16em; color: #a1947c; }
.pt-ink-row p { margin: 0; font-family: 'PingFang SC', 'Microsoft YaHei', sans-serif; font-size: 11px; line-height: 1.6; color: #4b4132; }
.pt-ink-med { display: flex; align-items: flex-start; gap: 8px; margin-top: 13px; padding: 8px 10px; border-left: 2px solid #b4472f; background: rgba(180, 71, 47, 0.045); }
.pt-ink-med-key { flex: none; padding: 1px 4px; background: #b4472f; color: #f8f1e2; font-size: 8px; letter-spacing: 0.1em; }
.pt-ink-med p { margin: 0; font-family: 'PingFang SC', 'Microsoft YaHei', sans-serif; font-size: 10px; line-height: 1.55; color: #8a4030; }
.pt-ink-note { position: relative; margin: 12px 1px 14px; padding: 4px 4px 4px 23px; }
.pt-ink-note-mark { position: absolute; left: 2px; top: -4px; font-family: Georgia, serif; font-size: 28px; color: #b3a68d; }
.pt-ink-note p { margin: 0; font-size: 12px; line-height: 1.75; letter-spacing: 0.02em; color: #5c5140; }
.pt-ink-phone { display: block; position: relative; padding: 12px 14px 13px; background: #29241d; color: #f7f1e2; text-decoration: none; transition: background 0.22s ease, transform 0.18s ease; }
.pt-ink-phone:hover { background: #3a3328; }
.pt-ink-phone:active { transform: scale(0.99); }
.pt-ink-phone::after { content: ''; position: absolute; top: 8px; right: 8px; width: 26px; height: 26px; border: 1px solid rgba(247, 241, 226, 0.38); border-radius: 50%; }
.pt-ink-phone-label { display: block; font-size: 9px; letter-spacing: 0.18em; color: #c8b99a; }
.pt-ink-phone strong { display: block; margin: 2px 0; font-family: 'Avenir Next', 'PingFang SC', sans-serif; font-size: clamp(27px, 8.3vw, 35px); font-weight: 700; letter-spacing: 0.02em; color: #fbf6e9; }
.pt-ink-phone-sub { font-size: 8px; letter-spacing: 0.09em; color: #b5a78b; }
.pt-ink-foot { display: flex; justify-content: space-between; align-items: baseline; margin-top: 15px; }
.pt-ink-signature { font-size: 12px; letter-spacing: 0.1em; color: #73664f; border-bottom: 1px solid rgba(111, 96, 70, 0.4); padding-bottom: 2px; }
.pt-ink-date { font-size: 8px; letter-spacing: 0.12em; color: #a39781; }
.pt-ink-below { margin: 18px 0 0; text-align: center; font-size: 10px; letter-spacing: 0.28em; color: #93876f; }
@keyframes pt-ink-bloom {
  0% { opacity: 0; filter: blur(5px); transform: scale(0.99); }
  100% { opacity: 1; filter: blur(0); transform: scale(1); }
}
@media (prefers-reduced-motion: reduce) {
  .pt-ink-content { animation: none; }
  .pt-ink-paper, .pt-ink-phone { transition: none; }
}
</style>
