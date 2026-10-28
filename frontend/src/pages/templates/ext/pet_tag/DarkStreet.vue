<template>
  <div class="pt-street-root min-h-screen px-5 py-8 text-zinc-100">
    <div class="pt-street-content mx-auto w-full max-w-md">
      <div class="pt-street-ticker" aria-hidden="true">
        <span>LOST DOG · CALL HOME · REWARD · LOST DOG · CALL HOME · REWARD ·</span>
      </div>

      <article class="pt-street-card">
        <header class="pt-street-head">
          <span class="pt-street-brand">STRAY／CREW</span>
          <span class="pt-street-id">TAG {{ nfcId || '0086' }}</span>
        </header>

        <div class="pt-street-hero">
          <p class="pt-street-kicker">IF YOU SEE THIS ONE</p>
          <h1><span>{{ payload.petName || '未命名' }}</span></h1>
          <p class="pt-street-meta">{{ payload.breed }} / {{ payload.age || '—' }} / KEEP SAFE</p>

          <span class="pt-street-sticker pt-street-sticker-lost">走失<br /><small>LOST</small></span>
          <span class="pt-street-sticker pt-street-sticker-paw" aria-hidden="true">
            <svg viewBox="0 0 48 48" fill="currentColor">
              <ellipse cx="24" cy="31" rx="10" ry="8" />
              <ellipse cx="11" cy="21" rx="4" ry="5" />
              <ellipse cx="21" cy="15" rx="4" ry="5.5" />
              <ellipse cx="32" cy="16" rx="4" ry="5" />
              <ellipse cx="40" cy="24" rx="3.4" ry="4.4" />
            </svg>
          </span>
          <span class="pt-street-tape" aria-hidden="true">REWARD · 必有重谢</span>
        </div>

        <div class="pt-street-grid">
          <div class="pt-street-box">
            <span>OWNER</span>
            <p>{{ payload.ownerName }}</p>
          </div>
          <div class="pt-street-box">
            <span>ADDRESS</span>
            <p>{{ payload.address }}</p>
          </div>
        </div>

        <div class="pt-street-alert">
          <span class="pt-street-alert-mark">!</span>
          <div>
            <span>MEDICAL / 健康备注</span>
            <p>{{ payload.medical || '无特殊记录' }}</p>
          </div>
        </div>

        <div class="pt-street-note">
          <span class="pt-street-note-key">NOTE</span>
          <p>{{ payload.note || '怕生，别追我。蹲下来，先给我闻闻手。' }}</p>
        </div>

        <a class="pt-street-phone" :href="'tel:' + payload.phone">
          <span class="pt-street-phone-top">DON'T POST. JUST CALL.</span>
          <strong>{{ payload.phone }}</strong>
          <span class="pt-street-phone-line" aria-hidden="true"></span>
          <span class="pt-street-phone-bottom">电话比转发更快 · 谢谢你的善意</span>
        </a>

        <footer class="pt-street-foot">
          <span>NO COLLAR? NO PROBLEM.</span>
          <span class="pt-street-chip">NFC</span>
        </footer>
      </article>

      <div class="pt-street-marquee" aria-hidden="true">
        <span>TAKE ME HOME&nbsp;&nbsp;✦&nbsp;&nbsp;TAKE ME HOME&nbsp;&nbsp;✦&nbsp;&nbsp;TAKE ME HOME&nbsp;&nbsp;✦&nbsp;&nbsp;</span>
      </div>
    </div>
  </div>
</template>

<script setup>
defineProps(['payload', 'nfcId'])
</script>

<style scoped>
.pt-street-root {
  background:
    radial-gradient(circle at 88% 6%, rgba(163, 230, 53, 0.15), transparent 33%),
    repeating-linear-gradient(115deg, rgba(255, 255, 255, 0.022) 0 2px, transparent 2px 8px),
    #101010;
  font-family: 'Arial Black', 'PingFang SC', sans-serif;
}
.pt-street-content { animation: pt-street-in 0.62s cubic-bezier(0.2, 0.9, 0.2, 1.15) both; }
.pt-street-ticker { overflow: hidden; padding: 5px 0; background: #a3e635; color: #101010; font-family: 'Courier New', monospace; font-size: 8px; font-weight: 700; letter-spacing: 0.16em; white-space: nowrap; }
.pt-street-ticker span { display: inline-block; animation: pt-street-marquee 18s linear infinite; }
.pt-street-card { position: relative; padding: 18px 17px 15px; background: #161616; border: 1px solid #343434; border-top: 0; box-shadow: 0 20px 36px rgba(0, 0, 0, 0.55); transition: transform 0.3s ease; }
.pt-street-card:hover { transform: translateY(-3px) rotate(-0.2deg); }
.pt-street-head { display: flex; justify-content: space-between; padding-bottom: 12px; border-bottom: 1px solid #383838; font-family: 'Courier New', monospace; font-size: 8px; letter-spacing: 0.13em; color: #8a8a8a; }
.pt-street-brand { color: #a3e635; font-weight: 700; }
.pt-street-hero { position: relative; padding: 24px 0 20px; }
.pt-street-kicker { margin: 0; font-family: 'Courier New', monospace; font-size: 8px; letter-spacing: 0.26em; color: #a3e635; }
.pt-street-hero h1 { margin: 4px 0 0; font-size: clamp(66px, 20vw, 92px); line-height: 0.82; letter-spacing: -0.055em; font-weight: 900; color: transparent; -webkit-text-stroke: 2px #f4f4f5; text-transform: uppercase; }
.pt-street-hero h1 span { display: inline-block; transition: color 0.25s ease, -webkit-text-stroke-color 0.25s ease; }
.pt-street-card:hover .pt-street-hero h1 span { color: #a3e635; -webkit-text-stroke-color: #a3e635; }
.pt-street-meta { margin: 10px 0 0; font-family: 'Courier New', monospace; font-size: 9px; letter-spacing: 0.1em; color: #8f8f8f; }
.pt-street-sticker { position: absolute; display: block; font-weight: 900; line-height: 0.85; }
.pt-street-sticker-lost {
  top: 4px;
  right: 0;
  padding: 9px 10px 8px;
  background: #ef4444;
  color: #fff;
  font-size: 24px;
  text-align: center;
  transform: rotate(9deg);
  box-shadow: 4px 4px 0 #7f1d1d;
  border-radius: 3px;
}
.pt-street-sticker-lost small { font-family: 'Courier New', monospace; font-size: 8px; letter-spacing: 0.16em; }
.pt-street-sticker-paw { top: 72px; right: 7px; width: 43px; color: #a3e635; transform: rotate(-13deg); }
.pt-street-sticker-paw svg { width: 100%; }
.pt-street-tape { position: absolute; left: -17px; bottom: 24px; padding: 5px 11px; background: #fafafa; color: #101010; font-family: 'Courier New', monospace; font-size: 8px; font-weight: 700; letter-spacing: 0.11em; transform: rotate(-2.5deg); box-shadow: 0 2px 4px rgba(0, 0, 0, 0.3); }
.pt-street-grid { display: grid; grid-template-columns: 0.9fr 1.1fr; gap: 8px; }
.pt-street-box { padding: 9px 10px; border: 1px solid #3d3d3d; background: #1c1c1c; }
.pt-street-box span { display: block; font-family: 'Courier New', monospace; font-size: 7px; letter-spacing: 0.17em; color: #868686; }
.pt-street-box p { margin: 3px 0 0; font-family: 'PingFang SC', 'Microsoft YaHei', sans-serif; font-size: 12px; font-weight: 700; line-height: 1.45; color: #e4e4e7; }
.pt-street-alert { display: flex; gap: 9px; margin-top: 9px; padding: 9px 10px; border: 1px solid #ef4444; background: rgba(239, 68, 68, 0.08); }
.pt-street-alert-mark { flex: none; width: 19px; height: 19px; display: grid; place-items: center; border-radius: 50%; background: #ef4444; color: #fff; font-size: 12px; }
.pt-street-alert span { font-family: 'Courier New', monospace; font-size: 7px; letter-spacing: 0.1em; color: #fca5a5; }
.pt-street-alert p { margin: 2px 0 0; font-family: 'PingFang SC', 'Microsoft YaHei', sans-serif; font-size: 11px; line-height: 1.5; color: #fecaca; }
.pt-street-note { display: grid; grid-template-columns: auto 1fr; gap: 8px; margin-top: 9px; padding: 9px 10px; border-left: 3px solid #a3e635; background: rgba(163, 230, 53, 0.055); }
.pt-street-note-key { font-family: 'Courier New', monospace; font-size: 7px; font-weight: 700; letter-spacing: 0.1em; color: #a3e635; padding-top: 2px; }
.pt-street-note p { margin: 0; font-family: 'PingFang SC', 'Microsoft YaHei', sans-serif; font-size: 11px; line-height: 1.6; color: #d4d4d8; }
.pt-street-phone { display: block; position: relative; margin: 15px -17px 0; padding: 13px 17px 14px; background: #a3e635; color: #101010; text-decoration: none; transition: background 0.18s ease; }
.pt-street-phone:hover { background: #bef264; }
.pt-street-phone:active { background: #84cc16; }
.pt-street-phone-top { display: block; font-family: 'Courier New', monospace; font-size: 8px; font-weight: 700; letter-spacing: 0.14em; }
.pt-street-phone strong { display: block; margin: 2px 0 3px; font-size: clamp(31px, 9.5vw, 41px); font-weight: 900; letter-spacing: -0.03em; line-height: 1.05; }
.pt-street-phone-line { display: block; height: 3px; background: #101010; margin-bottom: 4px; }
.pt-street-phone-bottom { font-family: 'PingFang SC', 'Microsoft YaHei', sans-serif; font-size: 9px; font-weight: 700; letter-spacing: 0.06em; }
.pt-street-foot { display: flex; justify-content: space-between; align-items: center; margin-top: 13px; font-family: 'Courier New', monospace; font-size: 7px; letter-spacing: 0.1em; color: #757575; }
.pt-street-chip { padding: 2px 5px; border: 1px solid #a3e635; color: #a3e635; }
.pt-street-marquee { overflow: hidden; margin-top: 14px; padding: 6px 0; border-top: 1px solid #333; border-bottom: 1px solid #333; color: #6b6b6b; font-family: 'Courier New', monospace; font-size: 8px; letter-spacing: 0.2em; white-space: nowrap; }
.pt-street-marquee span { display: inline-block; animation: pt-street-marquee 22s linear infinite; }
@keyframes pt-street-in {
  from { opacity: 0; transform: scale(0.96) rotate(-1deg); }
  to { opacity: 1; transform: scale(1) rotate(0); }
}
@keyframes pt-street-marquee {
  from { transform: translateX(0); }
  to { transform: translateX(-50%); }
}
@media (prefers-reduced-motion: reduce) {
  .pt-street-content, .pt-street-ticker span, .pt-street-marquee span { animation: none; }
  .pt-street-card, .pt-street-hero h1 span, .pt-street-phone { transition: none; }
}
</style>
