<template>
  <div class="pt-zakka-root min-h-screen px-5 py-9 text-stone-700">
    <div class="pt-zakka-content mx-auto w-full max-w-md">
      <div class="pt-zakka-string" aria-hidden="true">
        <span class="pt-zakka-string-line"></span>
        <span class="pt-zakka-bead"></span>
      </div>

      <article class="pt-zakka-tag">
        <span class="pt-zakka-tape pt-zakka-tape-a" aria-hidden="true"></span>
        <span class="pt-zakka-tape pt-zakka-tape-b" aria-hidden="true"></span>

        <div class="pt-zakka-hole" aria-hidden="true"></div>

        <header class="pt-zakka-head">
          <div>
            <p class="pt-zakka-en">HANDMADE PET TAG</p>
            <p class="pt-zakka-jp">ペットの名札</p>
          </div>
          <span class="pt-zakka-stamp">青禾<br />雜貨</span>
        </header>

        <div class="pt-zakka-main">
          <div class="pt-zakka-vertical">どうぞ よろしく</div>
          <div class="pt-zakka-namebox">
            <span class="pt-zakka-label">なまえ / 名字</span>
            <h1>{{ payload.petName || '未命名' }}</h1>
            <p>{{ payload.breed }}<span>／</span>{{ payload.age || '—' }}</p>
          </div>
        </div>

        <div class="pt-zakka-dashes" aria-hidden="true"></div>

        <div class="pt-zakka-list">
          <div class="pt-zakka-item">
            <span class="pt-zakka-key">家人</span>
            <span class="pt-zakka-dots" aria-hidden="true"></span>
            <span class="pt-zakka-val">{{ payload.ownerName }}</span>
          </div>
          <div class="pt-zakka-item">
            <span class="pt-zakka-key">住所</span>
            <span class="pt-zakka-dots" aria-hidden="true"></span>
            <span class="pt-zakka-val">{{ payload.address }}</span>
          </div>
          <div class="pt-zakka-item pt-zakka-item-med">
            <span class="pt-zakka-key">体调</span>
            <span class="pt-zakka-dots" aria-hidden="true"></span>
            <span class="pt-zakka-val">{{ payload.medical || '无特殊记录' }}</span>
          </div>
        </div>

        <div class="pt-zakka-note">
          <span class="pt-zakka-note-mark">メモ</span>
          <p>{{ payload.note || '人見知りです。ゆっくり近づいてね。' }}</p>
        </div>

        <a class="pt-zakka-phone" :href="'tel:' + payload.phone">
          <div class="pt-zakka-phone-head">
            <span>如果迷路了，请联络</span>
            <span class="pt-zakka-phone-stamp">要連絡</span>
          </div>
          <strong>{{ payload.phone }}</strong>
          <span class="pt-zakka-phone-sub">TEL · いつでも お願いします</span>
        </a>

        <footer class="pt-zakka-foot">
          <span>made by hand, kept with love</span>
          <span>NO. {{ nfcId || '086' }}</span>
        </footer>
      </article>

      <p class="pt-zakka-below">— 小さな名札に、大きな安心を —</p>
    </div>
  </div>
</template>

<script setup>
defineProps(['payload', 'nfcId'])
</script>

<style scoped>
.pt-zakka-root {
  background:
    repeating-linear-gradient(0deg, rgba(120, 100, 75, 0.045) 0 1px, transparent 1px 5px),
    repeating-linear-gradient(90deg, rgba(120, 100, 75, 0.035) 0 1px, transparent 1px 5px),
    #f1e8d8;
  font-family: 'Hiragino Mincho ProN', 'Yu Mincho', 'Songti SC', serif;
}
.pt-zakka-content { animation: pt-zakka-sway 0.9s cubic-bezier(0.2, 0.85, 0.3, 1) both; transform-origin: top center; }
.pt-zakka-string { display: flex; flex-direction: column; align-items: center; }
.pt-zakka-string-line { width: 2px; height: 39px; background: repeating-linear-gradient(180deg, #9b7a53 0 4px, #c8ae8b 4px 7px); transform: rotate(-1deg); }
.pt-zakka-bead { width: 12px; height: 12px; margin-top: -1px; border-radius: 50%; background: radial-gradient(circle at 35% 30%, #e8d5b6, #9d7b52); box-shadow: 0 1px 2px rgba(80, 60, 35, 0.3); }
.pt-zakka-tag {
  position: relative;
  margin-top: -2px;
  padding: 31px 21px 17px;
  background:
    radial-gradient(circle at 88% 10%, rgba(150, 120, 80, 0.09), transparent 27%),
    linear-gradient(150deg, #faf4e7, #f3ead8 62%, #ece0ca);
  border: 1px solid #cbb795;
  border-radius: 4px 4px 12px 12px;
  box-shadow: 0 18px 34px rgba(100, 78, 50, 0.17), inset 0 0 0 5px rgba(255, 252, 244, 0.6);
  transition: transform 0.35s ease, box-shadow 0.35s ease;
}
.pt-zakka-tag:hover { transform: translateY(-3px) rotate(0.3deg); box-shadow: 0 24px 42px rgba(100, 78, 50, 0.22); }
.pt-zakka-tape { position: absolute; width: 74px; height: 21px; opacity: 0.8; }
.pt-zakka-tape-a { top: 9px; left: -15px; background: repeating-linear-gradient(90deg, #c8d4bd 0 6px, #dbe3d2 6px 11px); transform: rotate(-24deg); }
.pt-zakka-tape-b { top: 15px; right: -18px; background: repeating-linear-gradient(90deg, #e2c9b7 0 5px, #ecd9c9 5px 10px); transform: rotate(21deg); }
.pt-zakka-hole { position: absolute; top: 12px; left: 50%; width: 13px; height: 13px; margin-left: -6.5px; border-radius: 50%; background: #e6d8bf; box-shadow: inset 0 2px 3px rgba(90, 70, 42, 0.4), 0 1px 0 #fffaf0; }
.pt-zakka-head { display: flex; justify-content: space-between; align-items: flex-start; padding: 3px 2px 10px; border-bottom: 1px solid rgba(130, 104, 67, 0.35); }
.pt-zakka-en { margin: 0; font-family: 'Courier New', monospace; font-size: 6.5px; font-weight: 700; letter-spacing: 0.16em; color: #a0835c; }
.pt-zakka-jp { margin: 3px 0 0; font-size: 12px; letter-spacing: 0.25em; color: #7d6747; }
.pt-zakka-stamp { width: 40px; height: 40px; display: grid; place-items: center; border: 1.5px solid #a8644a; border-radius: 50%; color: #a8644a; font-size: 8px; line-height: 1.25; text-align: center; letter-spacing: 0.05em; transform: rotate(-8deg); opacity: 0.85; }
.pt-zakka-main { display: flex; gap: 12px; padding: 18px 1px 4px; }
.pt-zakka-vertical { writing-mode: vertical-rl; font-size: 9px; letter-spacing: 0.2em; color: #9a8261; padding-top: 3px; border-right: 1px solid rgba(130, 104, 67, 0.25); padding-right: 8px; }
.pt-zakka-namebox { flex: 1; }
.pt-zakka-label { display: block; font-size: 8px; letter-spacing: 0.14em; color: #a0835c; }
.pt-zakka-namebox h1 { margin: 2px 0 0; font-size: clamp(43px, 12.5vw, 57px); line-height: 1.05; font-weight: 400; letter-spacing: 0.1em; color: #4b3b29; }
.pt-zakka-namebox p { margin: 6px 0 0; font-family: 'PingFang SC', 'Microsoft YaHei', sans-serif; font-size: 10px; letter-spacing: 0.12em; color: #8b7556; }
.pt-zakka-namebox p span { margin: 0 5px; color: #b09a7b; }
.pt-zakka-dashes { height: 1px; margin: 14px 0 8px; background: repeating-linear-gradient(90deg, rgba(130, 104, 67, 0.4) 0 4px, transparent 4px 8px); }
.pt-zakka-list { padding: 2px 1px 4px; }
.pt-zakka-item { display: flex; align-items: baseline; gap: 7px; padding: 5px 0; font-family: 'PingFang SC', 'Microsoft YaHei', sans-serif; }
.pt-zakka-key { flex: none; font-size: 9px; letter-spacing: 0.1em; color: #9b8463; }
.pt-zakka-dots { flex: 1; min-width: 10px; border-bottom: 1px dotted #c4b092; transform: translateY(-3px); }
.pt-zakka-val { flex: 0 1 auto; max-width: 70%; font-size: 11px; line-height: 1.5; text-align: right; color: #5f4d37; }
.pt-zakka-item-med .pt-zakka-val { color: #9a4b35; font-weight: 600; }
.pt-zakka-note { display: flex; gap: 9px; align-items: flex-start; margin: 7px 1px 14px; padding: 9px 10px; background: rgba(217, 201, 172, 0.38); border: 1px dashed #bca98a; }
.pt-zakka-note-mark { flex: none; padding: 2px 5px; background: #9d8460; color: #fbf6e9; font-size: 7px; letter-spacing: 0.1em; }
.pt-zakka-note p { margin: 0; font-family: 'PingFang SC', 'Microsoft YaHei', sans-serif; font-size: 11px; line-height: 1.65; color: #64513b; }
.pt-zakka-phone { display: block; position: relative; padding: 11px 13px 12px; background: #fdf8ec; border: 1px solid #bda887; box-shadow: 3px 3px 0 rgba(160, 133, 95, 0.24); color: #493927; text-decoration: none; transition: transform 0.2s ease, box-shadow 0.2s ease; }
.pt-zakka-phone:hover { transform: translate(-1px, -1px); box-shadow: 5px 5px 0 rgba(160, 133, 95, 0.28); }
.pt-zakka-phone:active { transform: translate(2px, 2px); box-shadow: 1px 1px 0 rgba(160, 133, 95, 0.24); }
.pt-zakka-phone-head { display: flex; justify-content: space-between; align-items: center; }
.pt-zakka-phone-head > span:first-child { font-family: 'PingFang SC', 'Microsoft YaHei', sans-serif; font-size: 9px; letter-spacing: 0.09em; color: #8c7656; }
.pt-zakka-phone-stamp { padding: 2px 5px; border: 1px solid #ac5a42; color: #ac5a42; font-size: 8px; letter-spacing: 0.05em; transform: rotate(3deg); }
.pt-zakka-phone strong { display: block; margin: 3px 0 1px; font-family: 'Avenir Next', 'PingFang SC', sans-serif; font-size: clamp(25px, 7.6vw, 31px); font-weight: 700; letter-spacing: 0.015em; }
.pt-zakka-phone-sub { font-family: 'Courier New', monospace; font-size: 7px; letter-spacing: 0.11em; color: #aa9270; }
.pt-zakka-foot { display: flex; justify-content: space-between; margin-top: 14px; padding-top: 9px; border-top: 1px solid rgba(130, 104, 67, 0.28); font-family: 'Courier New', monospace; font-size: 6.5px; letter-spacing: 0.08em; color: #a58e6d; }
.pt-zakka-below { margin: 17px 0 0; text-align: center; font-size: 10px; letter-spacing: 0.16em; color: #9c8767; }
@keyframes pt-zakka-sway {
  from { opacity: 0; transform: translateY(-20px) rotate(-2.2deg); }
  65% { opacity: 1; transform: translateY(2px) rotate(0.5deg); }
  to { opacity: 1; transform: translateY(0) rotate(0); }
}
@media (prefers-reduced-motion: reduce) {
  .pt-zakka-content { animation: none; }
  .pt-zakka-tag, .pt-zakka-phone { transition: none; }
}
</style>
