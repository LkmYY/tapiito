<template>
  <div class="pt-steel-root min-h-screen px-5 py-9 text-zinc-900">
    <div class="pt-steel-content mx-auto w-full max-w-md">
      <div class="pt-steel-hook" aria-hidden="true">
        <span class="pt-steel-hole"></span>
      </div>

      <article class="pt-steel-plate">
        <div class="pt-steel-brush" aria-hidden="true"></div>
        <div class="pt-steel-screw pt-steel-screw-a" aria-hidden="true"></div>
        <div class="pt-steel-screw pt-steel-screw-b" aria-hidden="true"></div>

        <header class="pt-steel-head">
          <div>
            <p class="pt-steel-kicker">PET IDENTIFICATION UNIT</p>
            <p class="pt-steel-code">STL-{{ nfcId || '0417' }} / REV.02</p>
          </div>
          <div class="pt-steel-stamp">REGISTERED</div>
        </header>

        <div class="pt-steel-nameblock">
          <span class="pt-steel-index">01</span>
          <h1>{{ payload.petName || '未命名' }}</h1>
          <p>{{ payload.breed }} · {{ payload.age || '—' }}</p>
        </div>

        <div class="pt-steel-table">
          <div class="pt-steel-row">
            <span>OWNER</span>
            <strong>{{ payload.ownerName }}</strong>
          </div>
          <div class="pt-steel-row">
            <span>ADDRESS</span>
            <strong>{{ payload.address }}</strong>
          </div>
          <div class="pt-steel-row pt-steel-row-alert">
            <span>MEDICAL</span>
            <strong>{{ payload.medical || '无特殊记录' }}</strong>
          </div>
        </div>

        <div class="pt-steel-call">
          <div class="pt-steel-call-label">
            <span class="pt-steel-blip" aria-hidden="true"></span>
            <span>紧急联络 / CALL FIRST</span>
          </div>
          <a :href="'tel:' + payload.phone">
            <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.2" aria-hidden="true">
              <path d="M5 4h4l2 5-2.5 1.5a11 11 0 0 0 5 5L15 13l5 2v4a2 2 0 0 1-2 2A16 16 0 0 1 3 6a2 2 0 0 1 2-2z" />
            </svg>
            {{ payload.phone }}
          </a>
        </div>

        <div class="pt-steel-note">
          <span class="pt-steel-note-key">NOTE</span>
          <p>{{ payload.note || '如发现我独自在外，请先联系主人。' }}</p>
        </div>

        <footer class="pt-steel-foot">
          <span>STAINLESS · 316L</span>
          <span>SCAN / TAP TO VERIFY</span>
        </footer>
      </article>

      <div class="pt-steel-hazard" aria-hidden="true"></div>
      <p class="pt-steel-bottom">REWARD IF FOUND · 送回必有酬谢</p>
    </div>
  </div>
</template>

<script setup>
defineProps(['payload', 'nfcId'])
</script>

<style scoped>
.pt-steel-root {
  background:
    linear-gradient(125deg, rgba(255, 255, 255, 0.75) 0 1px, transparent 1px 11px),
    linear-gradient(165deg, #d6d9dd 0%, #f1f3f5 38%, #b9bec5 100%);
  font-family: 'DIN Alternate', 'Arial Narrow', 'PingFang SC', sans-serif;
}
.pt-steel-content { animation: pt-steel-in 0.68s cubic-bezier(0.22, 0.85, 0.25, 1) both; }
.pt-steel-hook { display: flex; justify-content: center; }
.pt-steel-hole {
  display: block;
  width: 54px;
  height: 54px;
  border-radius: 50%;
  background: linear-gradient(140deg, #f7f8f9, #9ba1a8 60%, #e3e6e9);
  box-shadow: 0 5px 12px rgba(37, 43, 50, 0.35), inset 0 0 0 3px #ced3d8;
  position: relative;
}
.pt-steel-hole::after {
  content: '';
  position: absolute;
  inset: 17px;
  border-radius: 50%;
  background: #33383e;
  box-shadow: inset 0 3px 6px rgba(0, 0, 0, 0.7), 0 1px 0 #fff;
}
.pt-steel-plate {
  position: relative;
  margin-top: -12px;
  padding: 34px 25px 22px;
  background:
    repeating-linear-gradient(94deg, rgba(255, 255, 255, 0.22) 0 1px, rgba(110, 116, 123, 0.08) 1px 3px),
    linear-gradient(160deg, #f4f5f6 0%, #d9dde1 45%, #aeb4bb 100%);
  border: 1px solid #fbfcfd;
  border-radius: 4px;
  box-shadow: 0 22px 38px rgba(47, 54, 62, 0.32), inset 0 0 0 1px rgba(81, 88, 96, 0.25), inset 0 1px 0 #fff;
  clip-path: polygon(0 0, calc(100% - 26px) 0, 100% 26px, 100% 100%, 0 100%);
  overflow: hidden;
  transition: transform 0.3s ease, box-shadow 0.3s ease;
}
.pt-steel-plate:hover { transform: translateY(-3px); box-shadow: 0 30px 48px rgba(47, 54, 62, 0.4); }
.pt-steel-brush {
  position: absolute;
  inset: 0;
  background: linear-gradient(105deg, transparent 30%, rgba(255, 255, 255, 0.55) 38%, transparent 47%);
  pointer-events: none;
}
.pt-steel-screw { position: absolute; width: 9px; height: 9px; border-radius: 50%; background: radial-gradient(circle at 35% 30%, #f7f8f9, #7d848c); box-shadow: 0 1px 1px rgba(0, 0, 0, 0.35); }
.pt-steel-screw::after { content: ''; position: absolute; left: 1px; right: 1px; top: 4px; height: 1px; background: #5e646b; transform: rotate(-25deg); }
.pt-steel-screw-a { top: 12px; left: 12px; }
.pt-steel-screw-b { bottom: 12px; right: 12px; }
.pt-steel-head { position: relative; display: flex; justify-content: space-between; align-items: flex-start; gap: 10px; }
.pt-steel-kicker { margin: 0; font-size: 8px; font-weight: 700; letter-spacing: 0.2em; color: #4d545c; }
.pt-steel-code { margin: 5px 0 0; font-family: 'Courier New', monospace; font-size: 9px; color: #6a7179; }
.pt-steel-stamp {
  padding: 4px 7px;
  border: 1.5px solid #3f464e;
  font-family: 'Courier New', monospace;
  font-size: 8px;
  font-weight: 700;
  letter-spacing: 0.14em;
  transform: rotate(4deg);
  color: #3f464e;
  opacity: 0.7;
}
.pt-steel-nameblock { position: relative; margin: 26px 0 22px; padding-left: 14px; border-left: 4px solid #23282e; }
.pt-steel-index { position: absolute; left: -14px; top: -16px; font-family: 'Courier New', monospace; font-size: 9px; font-weight: 700; background: #23282e; color: #edf0f2; padding: 2px 5px; }
.pt-steel-nameblock h1 { margin: 0; font-size: clamp(42px, 12vw, 58px); line-height: 0.94; font-weight: 900; letter-spacing: -0.04em; color: #24292f; }
.pt-steel-nameblock p { margin: 9px 0 0; font-size: 12px; font-weight: 700; letter-spacing: 0.1em; color: #555c64; text-transform: uppercase; }
.pt-steel-table { border-top: 1px solid rgba(47, 54, 62, 0.35); }
.pt-steel-row { display: grid; grid-template-columns: 74px 1fr; gap: 10px; padding: 10px 0; border-bottom: 1px solid rgba(47, 54, 62, 0.25); font-size: 12px; line-height: 1.45; }
.pt-steel-row span { font-family: 'Courier New', monospace; font-size: 9px; font-weight: 700; letter-spacing: 0.1em; color: #646b73; padding-top: 2px; }
.pt-steel-row strong { font-weight: 600; color: #2d3339; }
.pt-steel-row-alert { background: rgba(255, 255, 255, 0.3); margin: 0 -7px; padding: 10px 7px; border-bottom-color: rgba(47, 54, 62, 0.4); }
.pt-steel-row-alert strong { color: #8a1f1f; }
.pt-steel-call { margin-top: 18px; padding: 13px 14px 12px; background: #24292f; color: #f3f5f6; box-shadow: 0 6px 0 rgba(85, 92, 100, 0.45); transition: transform 0.18s ease, box-shadow 0.18s ease; }
.pt-steel-call:active { transform: translateY(4px); box-shadow: 0 2px 0 rgba(85, 92, 100, 0.45); }
.pt-steel-call-label { display: flex; align-items: center; gap: 7px; font-family: 'Courier New', monospace; font-size: 8px; font-weight: 700; letter-spacing: 0.14em; color: #b9c0c7; }
.pt-steel-blip { width: 6px; height: 6px; border-radius: 50%; background: #e5ff6a; box-shadow: 0 0 8px #e5ff6a; }
.pt-steel-call a { display: flex; align-items: center; gap: 9px; margin-top: 4px; font-size: clamp(27px, 8vw, 35px); font-weight: 900; letter-spacing: 0.01em; line-height: 1.15; color: #fff; text-decoration: none; }
.pt-steel-call a svg { width: 23px; height: 23px; flex: none; color: #e5ff6a; }
.pt-steel-call a:hover { color: #e5ff6a; }
.pt-steel-note { display: flex; gap: 9px; margin-top: 18px; align-items: flex-start; }
.pt-steel-note-key { flex: none; margin-top: 2px; padding: 2px 5px; background: #5b626a; color: #f2f4f6; font-family: 'Courier New', monospace; font-size: 8px; font-weight: 700; letter-spacing: 0.1em; }
.pt-steel-note p { margin: 0; font-size: 12px; line-height: 1.65; color: #444b53; }
.pt-steel-foot { display: flex; justify-content: space-between; margin-top: 20px; padding-top: 10px; border-top: 1px dashed rgba(55, 62, 70, 0.4); font-family: 'Courier New', monospace; font-size: 8px; letter-spacing: 0.08em; color: #656c74; }
.pt-steel-hazard { height: 7px; margin-top: 22px; background: repeating-linear-gradient(135deg, #d2d6da 0 9px, #2e343a 9px 18px); opacity: 0.55; }
.pt-steel-bottom { margin: 12px 0 0; text-align: center; font-family: 'Courier New', monospace; font-size: 9px; letter-spacing: 0.18em; color: #4d545c; }
@keyframes pt-steel-in {
  from { opacity: 0; transform: translateY(26px) scale(0.97); }
  to { opacity: 1; transform: translateY(0) scale(1); }
}
@media (prefers-reduced-motion: reduce) {
  .pt-steel-content { animation: none; }
  .pt-steel-plate, .pt-steel-call { transition: none; }
}
</style>
