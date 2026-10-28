<template>
  <div class="pt-term-root min-h-screen px-4 py-7 text-emerald-300">
    <div class="pt-term-content mx-auto w-full max-w-md">
      <div class="pt-term-window">
        <div class="pt-term-bar">
          <span class="pt-term-bar-dot"></span>
          <span class="pt-term-bar-dot"></span>
          <span class="pt-term-bar-dot"></span>
          <span class="pt-term-path">~/pet_tag/{{ (payload.petName || 'unknown').toLowerCase() }}.cfg</span>
        </div>

        <div class="pt-term-body">
          <div class="pt-term-scan" aria-hidden="true"></div>

          <p class="pt-term-line pt-term-dim"># PET IDENTITY TERMINAL — v2.6.0</p>
          <p class="pt-term-line"><span class="pt-term-prompt">$</span> tag --read --secure</p>
          <p class="pt-term-line pt-term-ok">✓ NFC chip detected · {{ nfcId || '0086' }}</p>

          <div class="pt-term-ascii" aria-hidden="true">
            <span v-for="n in 48" :key="n" :class="{ 'pt-term-ascii-on': [2,3,6,7,10,14,15,18,19,22,26,27,30,31,34,38,39,42,43,46].includes(n) }"></span>
          </div>

          <p class="pt-term-line pt-term-dim">--- IDENTITY ---</p>
          <div class="pt-term-row">
            <span>PET_NAME</span>
            <p>{{ payload.petName }}</p>
          </div>
          <div class="pt-term-row">
            <span>BREED</span>
            <p>{{ payload.breed }}</p>
          </div>
          <div class="pt-term-row">
            <span>AGE</span>
            <p>{{ payload.age || '—' }}</p>
          </div>
          <div class="pt-term-row">
            <span>OWNER</span>
            <p>{{ payload.ownerName }}</p>
          </div>
          <div class="pt-term-row">
            <span>ADDRESS</span>
            <p>{{ payload.address }}</p>
          </div>

          <p class="pt-term-line pt-term-dim">--- MEDICAL FLAGS ---</p>
          <div class="pt-term-alert">
            <span>[!]</span>
            <p>{{ payload.medical || 'NO FLAGS' }}</p>
          </div>

          <p class="pt-term-line pt-term-dim">--- BEHAVIOUR ---</p>
          <p class="pt-term-note"><span class="pt-term-caret">&gt;</span> {{ payload.note || 'APPROACH SLOWLY. DO NOT PICK UP.' }}</p>

          <a class="pt-term-call" :href="'tel:' + payload.phone">
            <span class="pt-term-call-label">$ CALL_OWNER --immediate</span>
            <strong>{{ payload.phone }}<i aria-hidden="true">_</i></strong>
            <span class="pt-term-call-hint">走失时执行此指令 · 优先联络</span>
          </a>

          <p class="pt-term-line pt-term-dim">$ connection ready · awaiting call<span class="pt-term-blink"> ▌</span></p>
        </div>

        <footer class="pt-term-foot">
          <span>SECURE PET TAG</span>
          <span>[ESC] CLOSE</span>
        </footer>
      </div>

      <p class="pt-term-outside">root@qinghe:~$ keep_this_pet_safe</p>
    </div>
  </div>
</template>

<script setup>
defineProps(['payload', 'nfcId'])
</script>

<style scoped>
.pt-term-root {
  background:
    radial-gradient(circle at 50% 0, rgba(16, 185, 129, 0.12), transparent 36%),
    #030b06;
  font-family: 'Courier New', 'Lucida Console', monospace;
}
.pt-term-content { animation: pt-term-boot 0.8s steps(9, end) both; }
.pt-term-window { border: 1px solid #1f6b46; background: #06140c; box-shadow: 0 0 0 1px rgba(16, 185, 129, 0.12), 0 0 28px rgba(16, 185, 129, 0.1), 0 20px 40px rgba(0, 0, 0, 0.6); }
.pt-term-bar { display: flex; align-items: center; gap: 6px; padding: 7px 10px; background: #0b2215; border-bottom: 1px solid #1f6b46; }
.pt-term-bar-dot { width: 7px; height: 7px; border-radius: 50%; background: #2c6e4c; }
.pt-term-bar-dot:nth-child(2) { background: #4f8b46; }
.pt-term-bar-dot:nth-child(3) { background: #69b356; }
.pt-term-path { margin-left: auto; font-size: 7px; letter-spacing: 0.04em; color: #4e8f68; }
.pt-term-body { position: relative; padding: 15px 15px 13px; background-image: linear-gradient(rgba(16, 185, 129, 0.045) 1px, transparent 1px); background-size: 100% 21px; }
.pt-term-scan { position: absolute; inset: 0; pointer-events: none; background: repeating-linear-gradient(0deg, rgba(0, 0, 0, 0.22) 0 2px, transparent 2px 4px); opacity: 0.35; }
.pt-term-line { position: relative; margin: 0 0 6px; font-size: 9px; line-height: 1.55; color: #6ee7b7; }
.pt-term-dim { color: #3f8a64; letter-spacing: 0.03em; }
.pt-term-prompt { margin-right: 4px; color: #a3e635; }
.pt-term-ok { color: #a3e635; }
.pt-term-ascii { position: relative; display: grid; grid-template-columns: repeat(12, 1fr); gap: 2px; width: 78px; margin: 10px 0 12px; }
.pt-term-ascii span { display: block; aspect-ratio: 1; background: #0f2c1b; }
.pt-term-ascii-on { background: #4ade80; box-shadow: 0 0 4px rgba(74, 222, 128, 0.45); }
.pt-term-row { position: relative; display: grid; grid-template-columns: 78px 1fr; gap: 9px; padding: 4px 0; border-bottom: 1px dotted rgba(31, 107, 70, 0.45); font-size: 9px; line-height: 1.5; transition: background 0.2s ease; }
.pt-term-row:hover { background: rgba(16, 185, 129, 0.07); }
.pt-term-row span { font-size: 8px; color: #3f8a64; letter-spacing: 0.05em; }
.pt-term-row p { margin: 0; font-family: 'PingFang SC', 'Microsoft YaHei', monospace; color: #b7f7d8; }
.pt-term-alert { position: relative; display: flex; gap: 7px; align-items: flex-start; margin: 3px 0 9px; padding: 7px 8px; border: 1px solid #b91c1c; background: rgba(185, 28, 28, 0.12); }
.pt-term-alert span { flex: none; color: #f87171; font-size: 9px; font-weight: 700; }
.pt-term-alert p { margin: 0; font-family: 'PingFang SC', 'Microsoft YaHei', monospace; font-size: 10px; line-height: 1.5; color: #fca5a5; }
.pt-term-note { position: relative; margin: 0 0 11px; font-family: 'PingFang SC', 'Microsoft YaHei', monospace; font-size: 10px; line-height: 1.65; color: #9ed6b8; }
.pt-term-caret { color: #a3e635; }
.pt-term-call { position: relative; display: block; padding: 9px 11px 10px; border: 1px solid #a3e635; background: rgba(163, 230, 53, 0.07); color: #ecfccb; text-decoration: none; transition: background 0.18s ease, box-shadow 0.18s ease; }
.pt-term-call:hover { background: rgba(163, 230, 53, 0.14); box-shadow: 0 0 16px rgba(163, 230, 53, 0.18); }
.pt-term-call-label { display: block; font-size: 7px; letter-spacing: 0.08em; color: #84cc16; }
.pt-term-call strong { display: block; margin: 2px 0; font-family: 'Courier New', monospace; font-size: clamp(26px, 8vw, 33px); font-weight: 700; letter-spacing: 0.02em; color: #f7fee7; }
.pt-term-call strong i { font-style: normal; color: #a3e635; animation: pt-term-blink 1.1s steps(2, end) infinite; }
.pt-term-call-hint { font-family: 'PingFang SC', 'Microsoft YaHei', monospace; font-size: 8px; color: #86b98a; }
.pt-term-blink { color: #a3e635; animation: pt-term-blink 1.2s steps(2, end) infinite; }
.pt-term-foot { display: flex; justify-content: space-between; padding: 6px 10px; background: #0b2215; border-top: 1px solid #1f6b46; font-size: 7px; letter-spacing: 0.08em; color: #41835f; }
.pt-term-outside { margin: 15px 0 0; font-size: 8px; letter-spacing: 0.07em; color: #3d7a57; }
@keyframes pt-term-boot {
  from { opacity: 0; clip-path: inset(0 0 100% 0); }
  to { opacity: 1; clip-path: inset(0 0 0 0); }
}
@keyframes pt-term-blink {
  0%, 49% { opacity: 1; }
  50%, 100% { opacity: 0; }
}
@media (prefers-reduced-motion: reduce) {
  .pt-term-content, .pt-term-blink, .pt-term-call strong i { animation: none; }
  .pt-term-row, .pt-term-call { transition: none; }
}
</style>
