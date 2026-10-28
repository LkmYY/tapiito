<template>
  <div class="pt-pin-root min-h-screen px-5 py-9 text-purple-950">
    <div class="pt-pin-content mx-auto w-full max-w-md">
      <header class="pt-pin-head">
        <span>PINS &amp; PALS</span>
        <span>ENAMEL No. {{ nfcId || '086' }}</span>
      </header>

      <article class="pt-pin-card">
        <div class="pt-pin-grid" aria-hidden="true"></div>

        <div class="pt-pin-stage">
          <span class="pt-pin-hole" aria-hidden="true"></span>
          <div class="pt-pin-badge">
            <span class="pt-pin-shine" aria-hidden="true"></span>
            <div class="pt-pin-enamel">
              <span class="pt-pin-dots" aria-hidden="true"></span>
              <p class="pt-pin-small">HELLO, I AM</p>
              <h1>{{ payload.petName || '未命名' }}</h1>
              <span class="pt-pin-paw" aria-hidden="true">
                <svg viewBox="0 0 48 48" fill="currentColor">
                  <ellipse cx="24" cy="31" rx="9" ry="7" />
                  <ellipse cx="12" cy="22" rx="3.5" ry="4.5" />
                  <ellipse cx="21" cy="16" rx="3.5" ry="5" />
                  <ellipse cx="31" cy="17" rx="3.5" ry="4.5" />
                  <ellipse cx="39" cy="24" rx="3.2" ry="4.1" />
                </svg>
              </span>
              <span class="pt-pin-meta">{{ payload.breed }} · {{ payload.age || '—' }}</span>
            </div>
          </div>
          <div class="pt-pin-backing" aria-hidden="true">
            <span>COLLECTOR'S ITEM</span>
          </div>
        </div>

        <div class="pt-pin-owner">
          <div>
            <span class="pt-pin-label">MY HUMAN</span>
            <p>{{ payload.ownerName }}</p>
          </div>
          <div>
            <span class="pt-pin-label">OUR HOME</span>
            <p>{{ payload.address }}</p>
          </div>
        </div>

        <div class="pt-pin-med">
          <span class="pt-pin-med-badge" aria-hidden="true">✚</span>
          <div>
            <span class="pt-pin-label">MEDICAL NOTE</span>
            <p>{{ payload.medical || '无特殊记录' }}</p>
          </div>
        </div>

        <div class="pt-pin-note">
          <span class="pt-pin-label">A LITTLE NOTE</span>
          <p>{{ payload.note || '怕生，请勿强行抱抱。让我先认识你。' }}</p>
        </div>

        <a class="pt-pin-phone" :href="'tel:' + payload.phone">
          <span class="pt-pin-phone-top">
            <span>IF I'M LOST, PLEASE CALL</span>
            <span class="pt-pin-phone-pin" aria-hidden="true">◉</span>
          </span>
          <strong>{{ payload.phone }}</strong>
          <span class="pt-pin-phone-sub">走失请拨 · 谢谢你好心人</span>
        </a>

        <footer class="pt-pin-foot">
          <span>SOFT ENAMEL · GOLD PLATED</span>
          <span>{{ nfcId || 'PT-086' }}</span>
        </footer>
      </article>

      <p class="pt-pin-below">a little pin, a whole lot of love</p>
    </div>
  </div>
</template>

<script setup>
defineProps(['payload', 'nfcId'])
</script>

<style scoped>
.pt-pin-root {
  background:
    radial-gradient(circle at 15% 12%, rgba(196, 181, 253, 0.45), transparent 30%),
    radial-gradient(circle at 88% 80%, rgba(167, 243, 208, 0.4), transparent 33%),
    #f3efef;
  font-family: 'Avenir Next', 'PingFang SC', sans-serif;
}
.pt-pin-content { animation: pt-pin-pop 0.7s cubic-bezier(0.22, 0.9, 0.3, 1.2) both; }
.pt-pin-head { display: flex; justify-content: space-between; margin: 0 6px 11px; font-family: 'Courier New', monospace; font-size: 8px; letter-spacing: 0.15em; color: #8d7fa8; }
.pt-pin-card { position: relative; padding: 17px 17px 14px; background: #fffdfb; border-radius: 24px; border: 1px solid #e6dfe8; box-shadow: 0 22px 42px rgba(98, 76, 113, 0.15); overflow: hidden; transition: transform 0.35s ease, box-shadow 0.35s ease; }
.pt-pin-card:hover { transform: translateY(-4px); box-shadow: 0 30px 52px rgba(98, 76, 113, 0.2); }
.pt-pin-grid { position: absolute; inset: 0; background-image: radial-gradient(rgba(146, 120, 160, 0.17) 0.8px, transparent 0.8px); background-size: 13px 13px; opacity: 0.6; }
.pt-pin-stage { position: relative; display: flex; justify-content: center; padding: 12px 0 6px; }
.pt-pin-hole { position: absolute; top: 4px; left: 50%; width: 14px; height: 14px; margin-left: -7px; border-radius: 50%; background: #f3efef; box-shadow: inset 0 2px 3px rgba(90, 72, 100, 0.3); }
.pt-pin-badge { position: relative; z-index: 2; width: 206px; height: 206px; padding: 9px; border-radius: 50%; background: conic-gradient(from 30deg, #f7dfa5, #cba86a, #fbeec6, #b08c4f, #f7dfa5); box-shadow: 0 14px 21px rgba(103, 79, 42, 0.28), inset 0 0 0 2px #e9d29c; transition: transform 0.4s cubic-bezier(0.25, 0.8, 0.3, 1); animation: pt-pin-swing 0.9s 0.15s cubic-bezier(0.2, 0.9, 0.3, 1) both; }
.pt-pin-card:hover .pt-pin-badge { transform: rotate(-4deg) scale(1.02); }
.pt-pin-enamel { position: relative; height: 100%; display: flex; flex-direction: column; align-items: center; justify-content: center; border-radius: 50%; overflow: hidden; background: radial-gradient(circle at 32% 25%, #e9e2ff 0%, #c7b8f2 44%, #a78bfa 100%); box-shadow: inset 0 0 0 2px rgba(255, 255, 255, 0.35), inset 0 -12px 20px rgba(87, 57, 138, 0.22); }
.pt-pin-dots { position: absolute; inset: 0; background-image: radial-gradient(rgba(255, 255, 255, 0.38) 1px, transparent 1.4px); background-size: 11px 11px; opacity: 0.5; }
.pt-pin-shine { position: absolute; z-index: 4; top: 20px; left: 39px; width: 80px; height: 31px; border-radius: 50%; background: rgba(255, 255, 255, 0.57); filter: blur(5px); transform: rotate(-26deg); pointer-events: none; }
.pt-pin-small { position: relative; margin: 0; font-family: 'Courier New', monospace; font-size: 6.5px; letter-spacing: 0.19em; color: #59418a; }
.pt-pin-enamel h1 { position: relative; margin: 3px 0 0; max-width: 156px; font-size: clamp(26px, 8vw, 35px); line-height: 1.02; font-weight: 700; letter-spacing: 0.02em; color: #34205e; text-align: center; word-break: break-word; }
.pt-pin-paw { position: relative; width: 28px; margin: 5px 0 2px; color: #fdf6ff; filter: drop-shadow(0 1px 1px rgba(69, 41, 112, 0.35)); }
.pt-pin-paw svg { display: block; width: 100%; }
.pt-pin-meta { position: relative; font-family: 'Courier New', monospace; font-size: 7px; letter-spacing: 0.08em; color: #4c3479; }
.pt-pin-backing { position: absolute; z-index: 1; bottom: -3px; left: 50%; width: 224px; height: 42px; margin-left: -112px; border-radius: 50%; background: #ebe2f7; border: 1px solid #d6c8e9; transform: rotate(1deg); }
.pt-pin-backing span { position: absolute; left: 0; right: 0; bottom: 5px; text-align: center; font-family: 'Courier New', monospace; font-size: 5.5px; letter-spacing: 0.18em; color: #a491bd; }
.pt-pin-owner { display: grid; grid-template-columns: 0.85fr 1.15fr; gap: 11px; margin-top: 11px; padding: 10px 11px; border-radius: 12px; background: #f7f3f8; }
.pt-pin-label { display: block; margin-bottom: 3px; font-family: 'Courier New', monospace; font-size: 6.5px; font-weight: 700; letter-spacing: 0.12em; color: #a18cae; }
.pt-pin-owner p { margin: 0; font-size: 11px; font-weight: 600; line-height: 1.45; color: #5b4868; }
.pt-pin-med { display: flex; align-items: flex-start; gap: 8px; margin-top: 8px; padding: 9px 11px; border-radius: 12px; background: #fff1ea; border: 1px solid #f5d5c7; }
.pt-pin-med-badge { flex: none; width: 20px; height: 20px; display: grid; place-items: center; border-radius: 50%; background: #e7987c; color: #fff9f5; font-size: 11px; }
.pt-pin-med p { margin: 0; font-size: 10px; line-height: 1.55; color: #8a4932; }
.pt-pin-note { margin-top: 8px; padding: 9px 11px; border-radius: 12px; background: #f1f7f2; }
.pt-pin-note p { margin: 0; font-size: 10px; line-height: 1.65; color: #516b55; }
.pt-pin-phone { display: block; margin-top: 10px; padding: 11px 13px 12px; border-radius: 14px; background: #4c3479; color: #fdf9ff; text-decoration: none; box-shadow: 0 4px 0 #38245e; transition: transform 0.18s ease, box-shadow 0.18s ease, background 0.18s ease; }
.pt-pin-phone:hover { background: #5d428d; transform: translateY(-1px); box-shadow: 0 5px 0 #38245e; }
.pt-pin-phone:active { transform: translateY(3px); box-shadow: 0 1px 0 #38245e; }
.pt-pin-phone-top { display: flex; justify-content: space-between; align-items: center; font-family: 'Courier New', monospace; font-size: 6.5px; letter-spacing: 0.13em; color: #cdbdf7; }
.pt-pin-phone-pin { color: #f0abfc; }
.pt-pin-phone strong { display: block; margin: 1px 0; font-family: 'Avenir Next', 'PingFang SC', sans-serif; font-size: clamp(25px, 7.8vw, 32px); font-weight: 700; letter-spacing: 0.015em; }
.pt-pin-phone-sub { font-size: 8px; letter-spacing: 0.05em; color: #c4b5fd; }
.pt-pin-foot { display: flex; justify-content: space-between; margin-top: 12px; font-family: 'Courier New', monospace; font-size: 6.5px; letter-spacing: 0.08em; color: #a99ab5; }
.pt-pin-below { margin: 17px 0 0; text-align: center; font-family: 'Georgia', 'Songti SC', serif; font-size: 11px; font-style: italic; letter-spacing: 0.05em; color: #9b87a9; }
@keyframes pt-pin-pop {
  from { opacity: 0; transform: translateY(-18px) scale(0.97); }
  to { opacity: 1; transform: translateY(0) scale(1); }
}
@keyframes pt-pin-swing {
  0% { opacity: 0; transform: rotate(-9deg) translateY(-14px); }
  65% { opacity: 1; transform: rotate(2deg) translateY(1px); }
  100% { opacity: 1; transform: rotate(0) translateY(0); }
}
@media (prefers-reduced-motion: reduce) {
  .pt-pin-content, .pt-pin-badge { animation: none; }
  .pt-pin-card, .pt-pin-badge, .pt-pin-phone { transition: none; }
}
</style>
