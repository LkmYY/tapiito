<template>
  <div class="pt-collage-root min-h-screen px-5 py-9 text-stone-800">
    <div class="pt-collage-content mx-auto w-full max-w-md">
      <p class="pt-collage-top">剪贴簿 / MY LITTLE ARCHIVE</p>

      <div class="pt-collage-board">
        <span class="pt-collage-grid" aria-hidden="true"></span>

        <div class="pt-collage-piece pt-collage-photo">
          <div class="pt-collage-photo-inner">
            <div class="pt-collage-dog" aria-hidden="true">
              <span class="pt-collage-ear pt-collage-ear-l"></span>
              <span class="pt-collage-ear pt-collage-ear-r"></span>
              <span class="pt-collage-eye pt-collage-eye-l"></span>
              <span class="pt-collage-eye pt-collage-eye-r"></span>
              <span class="pt-collage-snout"></span>
              <span class="pt-collage-nose"></span>
            </div>
            <span class="pt-collage-photo-date">04 / 2026</span>
          </div>
          <p>my best friend</p>
        </div>

        <span class="pt-collage-tape pt-collage-tape-a" aria-hidden="true"></span>
        <span class="pt-collage-tape pt-collage-tape-b" aria-hidden="true"></span>

        <div class="pt-collage-titlepiece">
          <span class="pt-collage-eyebrow">FOUND ME? 请联络</span>
          <h1>{{ payload.petName || '未命名' }}</h1>
          <p>{{ payload.breed }}<span>·</span>{{ payload.age || '—' }}</p>
        </div>

        <div class="pt-collage-receipt">
          <span class="pt-collage-receipt-head">HOME RECEIPT</span>
          <div class="pt-collage-row">
            <span>家人</span>
            <b>{{ payload.ownerName }}</b>
          </div>
          <div class="pt-collage-row">
            <span>住址</span>
            <b>{{ payload.address }}</b>
          </div>
          <div class="pt-collage-row pt-collage-row-med">
            <span>健康</span>
            <b>{{ payload.medical || '无特殊记录' }}</b>
          </div>
        </div>

        <div class="pt-collage-sticky">
          <span class="pt-collage-sticky-pin" aria-hidden="true"></span>
          <p>{{ payload.note || '怕生，请勿强行抱抱。' }}</p>
          <span class="pt-collage-doodle" aria-hidden="true">♡</span>
        </div>

        <div class="pt-collage-stamp" aria-hidden="true">
          <span>RETURN<br />TO LOVE</span>
          <small>POSTAGE PAID</small>
        </div>

        <a class="pt-collage-phone" :href="'tel:' + payload.phone">
          <span class="pt-collage-phone-label">☎ CALL THE FAMILY</span>
          <strong>{{ payload.phone }}</strong>
          <span class="pt-collage-phone-sub">走失请拨 · 我们一直在等</span>
        </a>

        <span class="pt-collage-ticket" aria-hidden="true">ADMIT<br />ONE<br /><b>♥</b></span>
      </div>

      <p class="pt-collage-bottom">past·present·forever · 一页一页，都是你</p>
    </div>
  </div>
</template>

<script setup>
defineProps(['payload', 'nfcId'])
</script>

<style scoped>
.pt-collage-root {
  background:
    radial-gradient(circle at 12% 8%, rgba(13, 148, 136, 0.1), transparent 28%),
    radial-gradient(circle at 88% 92%, rgba(244, 114, 182, 0.12), transparent 30%),
    #efe6d6;
  font-family: 'Georgia', 'Songti SC', serif;
}
.pt-collage-content { animation: pt-collage-fade 0.7s ease both; }
.pt-collage-top { margin: 0 3px 13px; font-family: 'Courier New', monospace; font-size: 8px; letter-spacing: 0.2em; color: #8b7f6b; }
.pt-collage-board {
  position: relative;
  padding: 21px 18px 19px;
  min-height: 645px;
  background:
    radial-gradient(circle at 25% 18%, rgba(120, 95, 60, 0.06), transparent 30%),
    #faf4e7;
  border: 1px solid #d7cbb6;
  box-shadow: 0 20px 36px rgba(88, 74, 51, 0.17);
  overflow: hidden;
  transition: transform 0.3s ease;
}
.pt-collage-board:hover { transform: translateY(-3px) rotate(0.15deg); }
.pt-collage-grid { position: absolute; inset: 0; background-image: radial-gradient(rgba(120, 95, 60, 0.16) 0.8px, transparent 0.8px); background-size: 14px 14px; opacity: 0.5; pointer-events: none; }
.pt-collage-photo { position: relative; z-index: 3; width: 164px; padding: 7px 7px 20px; background: #fffdf6; box-shadow: 0 6px 12px rgba(80, 62, 40, 0.2); transform: rotate(-4deg); animation: pt-collage-drop 0.7s 0.1s ease both; }
.pt-collage-photo-inner { position: relative; height: 139px; overflow: hidden; background: linear-gradient(165deg, #b9d5cf, #82a9a1); }
.pt-collage-photo-inner::after { content: ''; position: absolute; inset: 0; background: repeating-linear-gradient(0deg, transparent 0 2px, rgba(255, 255, 255, 0.05) 2px 3px); }
.pt-collage-dog { position: absolute; left: 50%; bottom: 6px; width: 86px; height: 80px; margin-left: -43px; background: #d7a66f; border-radius: 47% 47% 42% 42%; box-shadow: inset -9px -5px 0 rgba(94, 57, 23, 0.16); }
.pt-collage-ear { position: absolute; top: -17px; width: 28px; height: 44px; background: #c18f57; border-radius: 50% 50% 25% 25%; }
.pt-collage-ear-l { left: -2px; transform: rotate(-22deg); }
.pt-collage-ear-r { right: -2px; transform: rotate(22deg); }
.pt-collage-eye { position: absolute; top: 30px; width: 8px; height: 10px; border-radius: 50%; background: #2c2013; }
.pt-collage-eye-l { left: 21px; }
.pt-collage-eye-r { right: 21px; }
.pt-collage-snout { position: absolute; left: 50%; bottom: 8px; width: 34px; height: 25px; margin-left: -17px; border-radius: 50%; background: #efd7b8; }
.pt-collage-nose { position: absolute; left: 50%; bottom: 23px; width: 10px; height: 7px; margin-left: -5px; border-radius: 50%; background: #2c2013; }
.pt-collage-photo-date { position: absolute; right: 6px; top: 6px; z-index: 2; padding: 1px 4px; background: rgba(255, 253, 246, 0.8); font-family: 'Courier New', monospace; font-size: 6px; color: #5d6c67; }
.pt-collage-photo > p { margin: 5px 0 0; text-align: center; font-family: 'Segoe Script', 'Bradley Hand', cursive; font-size: 10px; color: #796752; transform: rotate(-1deg); }
.pt-collage-tape { position: absolute; z-index: 5; width: 65px; height: 19px; background: rgba(214, 196, 157, 0.66); }
.pt-collage-tape-a { top: 3px; left: 92px; transform: rotate(8deg); }
.pt-collage-tape-b { top: 174px; right: 27px; transform: rotate(-11deg); background: rgba(191, 211, 198, 0.7); }
.pt-collage-titlepiece { position: relative; z-index: 2; margin: -4px 0 0 128px; padding: 12px 4px 7px; animation: pt-collage-slide 0.65s 0.2s ease both; }
.pt-collage-eyebrow { display: block; font-family: 'Courier New', monospace; font-size: 6.5px; font-weight: 700; letter-spacing: 0.1em; color: #9c5b3d; }
.pt-collage-titlepiece h1 { margin: 2px 0 0; font-size: clamp(35px, 10vw, 46px); line-height: 1; font-weight: 500; letter-spacing: 0.03em; color: #343c33; }
.pt-collage-titlepiece p { margin: 5px 0 0; font-family: 'Courier New', monospace; font-size: 8px; letter-spacing: 0.05em; color: #887b68; }
.pt-collage-titlepiece p span { margin: 0 4px; color: #b5a794; }
.pt-collage-receipt { position: relative; z-index: 2; width: 57%; margin: 18px 0 0 4px; padding: 10px 11px; background: #f4efdf; border: 1px solid #cfc2a8; box-shadow: 3px 3px 0 rgba(116, 98, 65, 0.13); transform: rotate(1.4deg); animation: pt-collage-rise 0.6s 0.32s ease both; }
.pt-collage-receipt-head { display: block; padding-bottom: 5px; margin-bottom: 4px; border-bottom: 1px dashed #b7a88c; font-family: 'Courier New', monospace; font-size: 7px; font-weight: 700; letter-spacing: 0.15em; color: #7b6a50; }
.pt-collage-row { display: flex; justify-content: space-between; gap: 7px; padding: 2px 0; font-family: 'PingFang SC', 'Microsoft YaHei', sans-serif; font-size: 9px; line-height: 1.45; }
.pt-collage-row span { flex: none; color: #988a74; }
.pt-collage-row b { font-weight: 500; text-align: right; color: #5d5546; }
.pt-collage-row-med b { color: #a04a34; }
.pt-collage-sticky { position: absolute; z-index: 4; top: 249px; right: 15px; width: 150px; padding: 13px 12px 17px; background: #f6dc8e; box-shadow: 2px 5px 9px rgba(100, 76, 34, 0.18); transform: rotate(4.5deg); animation: pt-collage-drop 0.6s 0.42s ease both; }
.pt-collage-sticky-pin { position: absolute; top: -6px; left: 50%; width: 11px; height: 11px; margin-left: -5.5px; border-radius: 50%; background: #c45d4a; box-shadow: inset -2px -2px 0 rgba(90, 32, 24, 0.24); }
.pt-collage-sticky p { margin: 0; font-family: 'PingFang SC', 'Microsoft YaHei', sans-serif; font-size: 10px; line-height: 1.6; color: #665126; }
.pt-collage-doodle { position: absolute; right: 9px; bottom: 1px; color: #b37958; font-size: 17px; }
.pt-collage-stamp { position: relative; z-index: 2; width: 82px; margin: 22px 0 0 172px; padding: 8px; border: 1px dashed #b06f4d; background: #f7ebd8; color: #a65e3e; text-align: center; transform: rotate(-5deg); }
.pt-collage-stamp span { font-family: 'Courier New', monospace; font-size: 8px; font-weight: 700; line-height: 1.2; letter-spacing: 0.06em; }
.pt-collage-stamp small { display: block; margin-top: 3px; font-size: 5.5px; letter-spacing: 0.1em; }
.pt-collage-phone { position: relative; z-index: 3; display: block; margin-top: 17px; padding: 11px 14px 12px; background: #2f4b44; color: #f8f0df; text-decoration: none; transform: rotate(-0.8deg); box-shadow: 4px 4px 0 #cabfa9; transition: transform 0.2s ease, box-shadow 0.2s ease; }
.pt-collage-phone:hover { transform: rotate(0) translateY(-2px); box-shadow: 6px 7px 0 #cabfa9; }
.pt-collage-phone:active { transform: translate(2px, 2px); box-shadow: 1px 1px 0 #cabfa9; }
.pt-collage-phone-label { display: block; font-family: 'Courier New', monospace; font-size: 7px; font-weight: 700; letter-spacing: 0.13em; color: #b7d0c4; }
.pt-collage-phone strong { display: block; margin: 1px 0; font-family: 'Avenir Next', 'PingFang SC', sans-serif; font-size: clamp(26px, 8vw, 33px); font-weight: 800; letter-spacing: 0.015em; }
.pt-collage-phone-sub { font-family: 'PingFang SC', 'Microsoft YaHei', sans-serif; font-size: 8px; letter-spacing: 0.06em; color: #a8c3b5; }
.pt-collage-ticket { position: absolute; z-index: 4; right: -9px; bottom: 19px; padding: 7px 10px; background: #d8a15d; color: #4d2e14; font-family: 'Courier New', monospace; font-size: 7px; font-weight: 700; line-height: 1.15; letter-spacing: 0.08em; transform: rotate(7deg); }
.pt-collage-ticket b { display: block; font-size: 11px; margin-top: 2px; }
.pt-collage-bottom { margin: 17px 0 0; text-align: center; font-family: 'Courier New', monospace; font-size: 8px; letter-spacing: 0.1em; color: #958873; }
@keyframes pt-collage-fade {
  from { opacity: 0; }
  to { opacity: 1; }
}
@keyframes pt-collage-drop {
  from { opacity: 0; transform: translateY(-16px) rotate(-8deg); }
  to { opacity: 1; transform: translateY(0) rotate(-4deg); }
}
@keyframes pt-collage-slide {
  from { opacity: 0; transform: translateX(18px); }
  to { opacity: 1; transform: translateX(0); }
}
@keyframes pt-collage-rise {
  from { opacity: 0; transform: translateY(12px) rotate(1.4deg); }
  to { opacity: 1; transform: translateY(0) rotate(1.4deg); }
}
@media (prefers-reduced-motion: reduce) {
  .pt-collage-content, .pt-collage-photo, .pt-collage-titlepiece, .pt-collage-receipt, .pt-collage-sticky { animation: none; }
  .pt-collage-board, .pt-collage-phone { transition: none; }
}
</style>
