<template>
  <div class="pt-lanyard-root min-h-screen px-5 py-0 text-slate-800">
    <div class="pt-lanyard-content mx-auto w-full max-w-md">
      <div class="pt-lanyard-strap" aria-hidden="true">
        <span class="pt-lanyard-strap-text">QINGHE PET CREW · FAMILY MEMBER · QINGHE PET CREW · FAMILY MEMBER ·</span>
      </div>

      <div class="pt-lanyard-clip" aria-hidden="true">
        <span class="pt-lanyard-clip-arm"></span>
        <span class="pt-lanyard-clip-body"></span>
      </div>

      <article class="pt-lanyard-badge">
        <div class="pt-lanyard-slot" aria-hidden="true"></div>

        <header class="pt-lanyard-head">
          <div class="pt-lanyard-logo" aria-hidden="true">
            <svg viewBox="0 0 48 48" fill="currentColor">
              <ellipse cx="24" cy="31" rx="9" ry="7" />
              <ellipse cx="12" cy="22" rx="3.5" ry="4.5" />
              <ellipse cx="21" cy="16" rx="3.5" ry="5" />
              <ellipse cx="31" cy="17" rx="3.5" ry="4.5" />
              <ellipse cx="39" cy="24" rx="3.1" ry="4" />
            </svg>
          </div>
          <div class="pt-lanyard-org">
            <p>青禾宠物救援</p>
            <span>QINGHE PET RESCUE · CREW</span>
          </div>
          <span class="pt-lanyard-year">2026</span>
        </header>

        <div class="pt-lanyard-photo">
          <div class="pt-lanyard-avatar" aria-hidden="true">
            <span class="pt-lanyard-ear pt-lanyard-ear-l"></span>
            <span class="pt-lanyard-ear pt-lanyard-ear-r"></span>
            <span class="pt-lanyard-face"></span>
            <span class="pt-lanyard-eye pt-lanyard-eye-l"></span>
            <span class="pt-lanyard-eye pt-lanyard-eye-r"></span>
            <span class="pt-lanyard-snout"></span>
            <span class="pt-lanyard-nose"></span>
          </div>
          <span class="pt-lanyard-photo-label">MEMBER PHOTO</span>
        </div>

        <div class="pt-lanyard-name">
          <span class="pt-lanyard-role">FAMILY MEMBER / 家庭成员</span>
          <h1>{{ payload.petName || '未命名' }}</h1>
          <p>{{ payload.breed }} · {{ payload.age || '—' }}</p>
        </div>

        <div class="pt-lanyard-info">
          <div class="pt-lanyard-row">
            <span class="pt-lanyard-key">监护人</span>
            <p>{{ payload.ownerName }}</p>
          </div>
          <div class="pt-lanyard-row">
            <span class="pt-lanyard-key">住址</span>
            <p>{{ payload.address }}</p>
          </div>
          <div class="pt-lanyard-row pt-lanyard-row-med">
            <span class="pt-lanyard-key">健康</span>
            <p>{{ payload.medical || '无特殊记录' }}</p>
          </div>
        </div>

        <div class="pt-lanyard-note">
          <span class="pt-lanyard-note-tag">NOTE</span>
          <p>{{ payload.note || '怕生，请勿强行抱抱。' }}</p>
        </div>

        <a class="pt-lanyard-phone" :href="'tel:' + payload.phone">
          <span class="pt-lanyard-phone-head">EMERGENCY CONTACT / 紧急联络</span>
          <strong>{{ payload.phone }}</strong>
          <span class="pt-lanyard-phone-sub">走失请拨 · 24 小时接听</span>
        </a>

        <div class="pt-lanyard-barcode" aria-hidden="true">
          <i v-for="n in 22" :key="n" :style="{ width: (n % 5 === 0 ? 3 : n % 3 === 0 ? 2 : 1) + 'px' }"></i>
        </div>

        <footer class="pt-lanyard-foot">
          <span>ID {{ nfcId || 'PT-0086' }}</span>
          <span>PROPERTY OF A LOVING HOME</span>
        </footer>
      </article>

      <p class="pt-lanyard-below">SCAN · TAP · CALL · TAKE ME HOME</p>
    </div>
  </div>
</template>

<script setup>
defineProps(['payload', 'nfcId'])
</script>

<style scoped>
.pt-lanyard-root {
  background:
    radial-gradient(circle at 50% 0, rgba(45, 212, 191, 0.16), transparent 32%),
    repeating-linear-gradient(135deg, rgba(255, 255, 255, 0.022) 0 1px, transparent 1px 8px),
    #101f26;
  font-family: 'Avenir Next', 'PingFang SC', sans-serif;
}
.pt-lanyard-content { animation: pt-lanyard-drop 0.85s cubic-bezier(0.2, 0.85, 0.3, 1) both; transform-origin: top center; }
.pt-lanyard-strap { position: relative; height: 51px; overflow: hidden; background: repeating-linear-gradient(90deg, #0f766e 0 38px, #115e59 38px 54px); border-bottom: 2px solid #134e4a; box-shadow: 0 3px 8px rgba(0, 0, 0, 0.25); }
.pt-lanyard-strap::after { content: ''; position: absolute; inset: 7px 0; border-top: 1px dashed rgba(204, 251, 241, 0.25); border-bottom: 1px dashed rgba(204, 251, 241, 0.25); }
.pt-lanyard-strap-text { display: block; padding-top: 21px; font-family: 'Courier New', monospace; font-size: 7px; font-weight: 700; letter-spacing: 0.16em; color: #99f6e4; white-space: nowrap; animation: pt-lanyard-slide 26s linear infinite; }
.pt-lanyard-clip { display: flex; flex-direction: column; align-items: center; margin-top: -1px; }
.pt-lanyard-clip-arm { width: 17px; height: 12px; border: 2px solid #94a3b8; border-bottom: 0; border-radius: 9px 9px 0 0; background: #334155; }
.pt-lanyard-clip-body { width: 38px; height: 18px; border-radius: 3px; background: linear-gradient(180deg, #cbd5e1, #7b8794); box-shadow: 0 2px 4px rgba(0, 0, 0, 0.35), inset 0 1px 0 #f8fafc; }
.pt-lanyard-badge { position: relative; margin: -3px 7px 0; padding: 16px 17px 13px; background: #fcfbf7; border-radius: 13px; box-shadow: 0 20px 38px rgba(0, 0, 0, 0.4); transition: transform 0.35s ease; }
.pt-lanyard-badge:hover { transform: rotate(-0.4deg) translateY(-2px); }
.pt-lanyard-slot { width: 49px; height: 7px; margin: -4px auto 13px; border-radius: 5px; background: #eae7df; box-shadow: inset 0 1px 2px rgba(80, 70, 55, 0.25); }
.pt-lanyard-head { display: flex; align-items: center; gap: 9px; padding-bottom: 11px; border-bottom: 2px solid #115e59; }
.pt-lanyard-logo { width: 34px; height: 34px; flex: none; display: grid; place-items: center; border-radius: 8px; background: #115e59; color: #ccfbf1; }
.pt-lanyard-logo svg { width: 21px; height: 21px; }
.pt-lanyard-org { flex: 1; }
.pt-lanyard-org p { margin: 0; font-size: 12px; font-weight: 700; letter-spacing: 0.04em; color: #134e4a; }
.pt-lanyard-org span { font-family: 'Courier New', monospace; font-size: 5.5px; letter-spacing: 0.13em; color: #73948f; }
.pt-lanyard-year { font-family: 'Courier New', monospace; font-size: 8px; color: #8fa5a1; }
.pt-lanyard-photo { position: relative; height: 156px; margin-top: 12px; background: linear-gradient(160deg, #d4e9e4, #a9cec6); overflow: hidden; }
.pt-lanyard-photo::before { content: ''; position: absolute; inset: 0; background: repeating-linear-gradient(0deg, transparent 0 3px, rgba(17, 94, 89, 0.045) 3px 4px); }
.pt-lanyard-avatar { position: absolute; left: 50%; bottom: 9px; width: 105px; height: 110px; margin-left: -52.5px; }
.pt-lanyard-face { position: absolute; left: 12px; top: 21px; width: 81px; height: 85px; border-radius: 46% 46% 43% 43%; background: #d5a676; box-shadow: inset -9px -6px 0 rgba(84, 51, 23, 0.14); }
.pt-lanyard-ear { position: absolute; top: 2px; width: 31px; height: 50px; background: #c08e5d; border-radius: 50% 50% 22% 22%; }
.pt-lanyard-ear-l { left: 7px; transform: rotate(-18deg); }
.pt-lanyard-ear-r { right: 7px; transform: rotate(18deg); }
.pt-lanyard-eye { position: absolute; top: 56px; width: 8px; height: 10px; border-radius: 50%; background: #2b1e12; }
.pt-lanyard-eye-l { left: 32px; }
.pt-lanyard-eye-r { right: 32px; }
.pt-lanyard-snout { position: absolute; left: 50%; top: 73px; width: 37px; height: 27px; margin-left: -18.5px; border-radius: 50%; background: #ecd2ae; }
.pt-lanyard-nose { position: absolute; left: 50%; top: 81px; width: 11px; height: 8px; margin-left: -5.5px; border-radius: 50%; background: #2b1e12; }
.pt-lanyard-photo-label { position: absolute; top: 6px; left: 7px; padding: 2px 4px; background: rgba(252, 251, 247, 0.8); font-family: 'Courier New', monospace; font-size: 5.5px; letter-spacing: 0.12em; color: #47766f; }
.pt-lanyard-name { padding: 14px 1px 11px; border-bottom: 1px solid #e1ddd2; }
.pt-lanyard-role { display: block; font-family: 'Courier New', monospace; font-size: 6.5px; font-weight: 700; letter-spacing: 0.12em; color: #b06b48; }
.pt-lanyard-name h1 { margin: 2px 0 0; font-size: clamp(36px, 10.6vw, 47px); line-height: 1.05; letter-spacing: 0.04em; color: #173f3b; }
.pt-lanyard-name p { margin: 5px 0 0; font-family: 'Courier New', monospace; font-size: 8px; letter-spacing: 0.08em; color: #8b9b96; }
.pt-lanyard-info { padding: 4px 1px 2px; }
.pt-lanyard-row { display: grid; grid-template-columns: 44px 1fr; gap: 8px; padding: 6px 0; border-bottom: 1px dotted #ddd8cc; font-size: 10px; line-height: 1.5; }
.pt-lanyard-key { font-family: 'PingFang SC', 'Microsoft YaHei', sans-serif; font-size: 8px; letter-spacing: 0.08em; color: #98a8a3; }
.pt-lanyard-row p { margin: 0; color: #4e5d59; }
.pt-lanyard-row-med p { color: #a04a34; font-weight: 600; }
.pt-lanyard-note { display: flex; gap: 7px; align-items: flex-start; margin: 7px 0 11px; padding: 8px 9px; background: #f2f5ee; }
.pt-lanyard-note-tag { flex: none; padding: 2px 4px; background: #789378; color: #fafcf7; font-family: 'Courier New', monospace; font-size: 6px; letter-spacing: 0.08em; }
.pt-lanyard-note p { margin: 0; font-family: 'PingFang SC', 'Microsoft YaHei', sans-serif; font-size: 10px; line-height: 1.55; color: #5d6d5c; }
.pt-lanyard-phone { display: block; padding: 10px 12px 11px; background: #115e59; color: #f0fdfa; text-decoration: none; border-radius: 7px; transition: background 0.2s ease, transform 0.18s ease; }
.pt-lanyard-phone:hover { background: #0f766e; }
.pt-lanyard-phone:active { transform: scale(0.99); }
.pt-lanyard-phone-head { display: block; font-family: 'Courier New', monospace; font-size: 6px; letter-spacing: 0.13em; color: #99f6e4; }
.pt-lanyard-phone strong { display: block; margin: 2px 0; font-family: 'Avenir Next', 'PingFang SC', sans-serif; font-size: clamp(26px, 7.9vw, 32px); font-weight: 700; letter-spacing: 0.02em; }
.pt-lanyard-phone-sub { font-family: 'PingFang SC', 'Microsoft YaHei', sans-serif; font-size: 8px; letter-spacing: 0.06em; color: #a7f3d0; }
.pt-lanyard-barcode { display: flex; justify-content: center; align-items: flex-end; gap: 2px; height: 27px; margin-top: 12px; opacity: 0.55; }
.pt-lanyard-barcode i { display: block; height: 100%; background: #3c4a46; }
.pt-lanyard-barcode i:nth-child(4n) { height: 78%; }
.pt-lanyard-foot { display: flex; justify-content: space-between; margin-top: 7px; font-family: 'Courier New', monospace; font-size: 5.5px; letter-spacing: 0.08em; color: #a4a59c; }
.pt-lanyard-below { margin: 16px 0 0; padding-bottom: 22px; text-align: center; font-family: 'Courier New', monospace; font-size: 8px; letter-spacing: 0.17em; color: #5e8b84; }
@keyframes pt-lanyard-drop {
  0% { opacity: 0; transform: translateY(-34px) rotate(-2deg); }
  60% { opacity: 1; transform: translateY(3px) rotate(0.4deg); }
  100% { opacity: 1; transform: translateY(0) rotate(0); }
}
@keyframes pt-lanyard-slide {
  from { transform: translateX(0); }
  to { transform: translateX(-50%); }
}
@media (prefers-reduced-motion: reduce) {
  .pt-lanyard-content, .pt-lanyard-strap-text { animation: none; }
  .pt-lanyard-badge, .pt-lanyard-phone { transition: none; }
}
</style>
