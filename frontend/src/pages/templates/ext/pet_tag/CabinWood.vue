<template>
  <div class="pt-wood-root min-h-screen px-5 py-9 text-amber-50">
    <div class="pt-wood-content mx-auto w-full max-w-md">
      <div class="pt-wood-tree" aria-hidden="true">
        <span v-for="n in 5" :key="n" class="pt-wood-pine" :style="{ left: (n * 19 - 4) + '%', transform: 'scale(' + (0.55 + (n % 3) * 0.22) + ')' }"></span>
      </div>

      <div class="pt-wood-rope" aria-hidden="true">
        <span class="pt-wood-rope-line"></span>
        <span class="pt-wood-knot"></span>
      </div>

      <article class="pt-wood-sign">
        <span v-for="n in 4" :key="n" :class="'pt-wood-nail pt-wood-nail-' + n" aria-hidden="true"></span>

        <header class="pt-wood-head">
          <span class="pt-wood-brand">CABIN No.7</span>
          <span class="pt-wood-est">EST. 2021</span>
        </header>

        <div class="pt-wood-title">
          <span class="pt-wood-lead">OUR LITTLE</span>
          <h1>{{ payload.petName || '未命名' }}</h1>
          <p>{{ payload.breed }} · {{ payload.age || '—' }}</p>
        </div>

        <div class="pt-wood-rule" aria-hidden="true">
          <span></span>
          <svg viewBox="0 0 48 48" fill="currentColor">
            <ellipse cx="24" cy="31" rx="9" ry="7" />
            <ellipse cx="12" cy="22" rx="3.4" ry="4.4" />
            <ellipse cx="21" cy="16" rx="3.4" ry="4.8" />
            <ellipse cx="31" cy="17" rx="3.4" ry="4.4" />
            <ellipse cx="39" cy="24" rx="3.1" ry="4" />
          </svg>
          <span></span>
        </div>

        <div class="pt-wood-owner">
          <span class="pt-wood-key">家人</span>
          <p>{{ payload.ownerName }}</p>
          <span class="pt-wood-key">小屋</span>
          <p>{{ payload.address }}</p>
        </div>

        <div class="pt-wood-paper">
          <span class="pt-wood-paper-pin" aria-hidden="true"></span>
          <span class="pt-wood-paper-label">MEDICAL / 健康备忘</span>
          <p>{{ payload.medical || '无特殊记录' }}</p>
        </div>

        <p class="pt-wood-note">“{{ payload.note || '怕生，请勿强行抱抱。让我先闻闻你。' }}”</p>

        <a class="pt-wood-phone" :href="'tel:' + payload.phone">
          <span class="pt-wood-phone-top">IF FOUND, CALL THE CABIN</span>
          <strong>{{ payload.phone }}</strong>
          <span class="pt-wood-phone-bottom">走失请拨打 · 山里信号也为你开着</span>
        </a>

        <footer class="pt-wood-foot">
          <span>HAND CARVED · 手作木牌</span>
          <span>{{ nfcId || 'C-007' }}</span>
        </footer>
      </article>

      <p class="pt-wood-below">— TAKE ME BACK TO THE WOODS —</p>
    </div>
  </div>
</template>

<script setup>
defineProps(['payload', 'nfcId'])
</script>

<style scoped>
.pt-wood-root {
  position: relative;
  overflow: hidden;
  background:
    radial-gradient(circle at 50% 0, rgba(190, 210, 164, 0.2), transparent 37%),
    linear-gradient(180deg, #1d2d23 0%, #16241c 58%, #101b15 100%);
  font-family: 'Rockwell', 'Songti SC', 'Georgia', serif;
}
.pt-wood-tree { position: absolute; inset: 0; pointer-events: none; opacity: 0.25; }
.pt-wood-pine { position: absolute; bottom: 0; display: block; width: 55px; height: 170px; background: linear-gradient(180deg, #6f8d68, #2b4534); clip-path: polygon(50% 0, 88% 35%, 68% 35%, 100% 72%, 62% 72%, 76% 100%, 24% 100%, 38% 72%, 0 72%, 32% 35%, 12% 35%); }
.pt-wood-content { position: relative; z-index: 1; animation: pt-wood-sway 1s cubic-bezier(0.2, 0.85, 0.3, 1) both; transform-origin: top center; }
.pt-wood-rope { display: flex; flex-direction: column; align-items: center; }
.pt-wood-rope-line { width: 6px; height: 45px; background: repeating-linear-gradient(125deg, #b49362 0 3px, #8a6c40 3px 6px); border-radius: 3px; box-shadow: inset 1px 0 0 rgba(255, 236, 199, 0.25); }
.pt-wood-knot { width: 26px; height: 17px; margin-top: -2px; margin-bottom: 4px; border-radius: 50%; border: 4px solid #a28151; box-shadow: inset 0 0 0 2px #6f542f; }
.pt-wood-sign {
  position: relative;
  padding: 21px 20px 17px;
  background:
    repeating-linear-gradient(94deg, rgba(86, 51, 23, 0.12) 0 2px, transparent 2px 9px, rgba(255, 231, 186, 0.06) 9px 11px, transparent 11px 23px),
    radial-gradient(ellipse at 30% 10%, rgba(255, 226, 171, 0.17), transparent 35%),
    linear-gradient(155deg, #9a653a 0%, #784a27 48%, #8b5a31 100%);
  border: 5px solid #5a361b;
  border-radius: 12px 9px 13px 8px;
  box-shadow: 0 24px 40px rgba(0, 0, 0, 0.5), inset 0 0 0 2px rgba(255, 222, 166, 0.2), inset 0 0 30px rgba(53, 29, 12, 0.36);
  transition: transform 0.4s ease, box-shadow 0.4s ease;
}
.pt-wood-sign:hover { transform: rotate(0.35deg) translateY(-3px); box-shadow: 0 32px 50px rgba(0, 0, 0, 0.55), inset 0 0 0 2px rgba(255, 222, 166, 0.24); }
.pt-wood-sign::before { content: ''; position: absolute; inset: 8px; border: 1px solid rgba(255, 226, 176, 0.16); border-radius: 7px; pointer-events: none; }
.pt-wood-nail { position: absolute; width: 8px; height: 8px; border-radius: 50%; background: radial-gradient(circle at 35% 30%, #4b3c2a, #20170d); box-shadow: 0 1px 0 rgba(255, 226, 176, 0.2), inset 0 1px 1px rgba(0, 0, 0, 0.6); }
.pt-wood-nail-1 { top: 11px; left: 11px; }
.pt-wood-nail-2 { top: 11px; right: 11px; }
.pt-wood-nail-3 { bottom: 11px; left: 11px; }
.pt-wood-nail-4 { bottom: 11px; right: 11px; }
.pt-wood-head { display: flex; justify-content: space-between; padding-bottom: 9px; border-bottom: 1px solid rgba(247, 223, 180, 0.28); font-family: 'Courier New', monospace; font-size: 7px; letter-spacing: 0.13em; color: #e2c294; }
.pt-wood-brand { font-weight: 700; }
.pt-wood-title { padding: 18px 3px 8px; text-align: center; }
.pt-wood-lead { display: block; font-family: 'Courier New', monospace; font-size: 7px; letter-spacing: 0.25em; color: #d8b17c; }
.pt-wood-title h1 { margin: 3px 0 0; font-size: clamp(48px, 15vw, 65px); line-height: 1; font-weight: 600; letter-spacing: 0.1em; text-indent: 0.1em; color: #f7e2bb; text-shadow: 2px 3px 0 #56351a, 0 1px 0 rgba(255, 250, 232, 0.4), 0 -1px 0 rgba(57, 33, 14, 0.55); }
.pt-wood-title p { margin: 8px 0 0; font-family: 'Avenir Next', 'PingFang SC', sans-serif; font-size: 10px; letter-spacing: 0.16em; color: #e5c69a; }
.pt-wood-rule { display: flex; align-items: center; gap: 9px; margin: 5px 0 12px; color: #e0bd8d; }
.pt-wood-rule span { flex: 1; height: 1px; background: linear-gradient(90deg, transparent, rgba(240, 210, 163, 0.45), transparent); }
.pt-wood-rule svg { width: 17px; height: 17px; opacity: 0.8; }
.pt-wood-owner { display: grid; grid-template-columns: auto 1fr auto 1fr; gap: 7px; align-items: baseline; padding: 0 2px 6px; }
.pt-wood-key { font-family: 'PingFang SC', 'Microsoft YaHei', sans-serif; font-size: 9px; color: #cfa877; letter-spacing: 0.08em; }
.pt-wood-owner p { margin: 0; font-family: 'PingFang SC', 'Microsoft YaHei', sans-serif; font-size: 11px; line-height: 1.5; color: #f2ddba; word-break: break-word; }
.pt-wood-paper { position: relative; margin: 10px 2px 0; padding: 11px 12px; background: #f2e5c8; color: #51351b; transform: rotate(-0.7deg); box-shadow: 2px 3px 0 rgba(46, 27, 12, 0.25); }
.pt-wood-paper-pin { position: absolute; top: -5px; left: 50%; width: 10px; height: 10px; margin-left: -5px; border-radius: 50%; background: #b98450; box-shadow: inset -1px -1px 0 rgba(83, 51, 25, 0.35); }
.pt-wood-paper-label { display: block; margin-bottom: 3px; font-family: 'Courier New', monospace; font-size: 7px; font-weight: 700; letter-spacing: 0.11em; color: #9d7242; }
.pt-wood-paper p { margin: 0; font-family: 'PingFang SC', 'Microsoft YaHei', sans-serif; font-size: 11px; line-height: 1.55; color: #6b3a26; }
.pt-wood-note { margin: 13px 3px 14px; text-align: center; font-size: 12px; font-style: italic; line-height: 1.7; color: #f0d5ab; }
.pt-wood-phone { display: block; padding: 11px 13px 12px; background: #203827; border: 1px solid rgba(226, 194, 148, 0.65); box-shadow: inset 0 0 0 3px rgba(14, 32, 22, 0.6), 0 4px 0 rgba(39, 25, 12, 0.42); color: #f9e9c8; text-decoration: none; text-align: center; transition: background 0.2s ease, transform 0.18s ease; }
.pt-wood-phone:hover { background: #2a4632; }
.pt-wood-phone:active { transform: translateY(3px); }
.pt-wood-phone-top { display: block; font-family: 'Courier New', monospace; font-size: 7px; letter-spacing: 0.18em; color: #c9aa78; }
.pt-wood-phone strong { display: block; margin: 2px 0; font-family: 'Avenir Next', 'PingFang SC', sans-serif; font-size: clamp(28px, 8.6vw, 36px); font-weight: 700; letter-spacing: 0.02em; color: #fff2d2; text-shadow: 0 1px 0 #0e1c13; }
.pt-wood-phone-bottom { font-family: 'PingFang SC', 'Microsoft YaHei', sans-serif; font-size: 8px; letter-spacing: 0.08em; color: #bda476; }
.pt-wood-foot { display: flex; justify-content: space-between; margin-top: 13px; padding-top: 9px; border-top: 1px solid rgba(240, 210, 163, 0.25); font-family: 'Courier New', monospace; font-size: 6.5px; letter-spacing: 0.1em; color: #c6a575; }
.pt-wood-below { margin: 18px 0 0; text-align: center; font-family: 'Courier New', monospace; font-size: 8px; letter-spacing: 0.18em; color: rgba(226, 194, 148, 0.62); }
@keyframes pt-wood-sway {
  0% { opacity: 0; transform: rotate(-3.5deg) translateY(-22px); }
  55% { opacity: 1; transform: rotate(1.2deg) translateY(2px); }
  100% { opacity: 1; transform: rotate(0) translateY(0); }
}
@media (prefers-reduced-motion: reduce) {
  .pt-wood-content { animation: none; }
  .pt-wood-sign, .pt-wood-phone { transition: none; }
}
</style>
