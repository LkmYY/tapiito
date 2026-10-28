<template>
  <div class="nc-t3-root relative min-h-screen px-4 py-6 text-[#8dffb0]">
    <div class="nc-t3-noise"></div>

    <div class="relative z-10 mx-auto flex min-h-[calc(100vh-3rem)] w-full max-w-md flex-col">
      <div class="nc-t3-shell">
        <div class="nc-t3-bar">
          <span class="nc-t3-dots"><i></i><i></i><i></i></span>
          <span class="nc-t3-path">guest@neon-room:~/nightclub</span>
          <span class="nc-t3-win">— □ ×</span>
        </div>

        <div class="nc-t3-screen">
          <p class="nc-t3-line nc-t3-dim">NeonOS 3.0.3 · build 1998 · secure shell</p>
          <p class="nc-t3-line nc-t3-dim">Last login: Fri 02:13:47 on ttys003</p>
          <div class="nc-t3-gap"></div>

          <div class="nc-t3-block">
            <p class="nc-t3-cmd"><span class="nc-t3-prompt">$</span> whoami</p>
            <div class="nc-t3-output">
              <p class="nc-t3-name">{{ payload?.name }}</p>
              <p class="nc-t3-alias">{{ payload?.alias || 'NO_ALIAS' }}</p>
            </div>
          </div>

          <div class="nc-t3-block">
            <p class="nc-t3-cmd"><span class="nc-t3-prompt">$</span> cat role.txt</p>
            <p class="nc-t3-output nc-t3-value">{{ payload?.role }}</p>
          </div>

          <div class="nc-t3-block">
            <p class="nc-t3-cmd"><span class="nc-t3-prompt">$</span> locate venue</p>
            <div class="nc-t3-output nc-t3-pathlist">
              <span class="nc-t3-dir">/clubs/</span>
              <span class="nc-t3-venue">{{ payload?.venue }}</span>
              <span class="nc-t3-city">{{ payload?.city || '—' }}</span>
            </div>
          </div>

          <div class="nc-t3-block">
            <p class="nc-t3-cmd"><span class="nc-t3-prompt">$</span> contact --wechat</p>
            <p class="nc-t3-output nc-t3-value">@{{ payload?.wechat }}</p>
          </div>

          <div class="nc-t3-block">
            <p class="nc-t3-cmd"><span class="nc-t3-prompt">$</span> contact --phone</p>
            <a v-if="payload?.phone" :href="'tel:' + payload.phone" class="nc-t3-output nc-t3-phone">
              {{ payload.phone }}
              <span class="nc-t3-hint">[tap to dial]</span>
            </a>
          </div>

          <p class="nc-t3-cmd nc-t3-cursorline">
            <span class="nc-t3-prompt">$</span><span class="nc-t3-cursor"></span>
          </p>
        </div>

        <div class="nc-t3-status">
          <span>CONNECTED</span>
          <span>UPTIME 04:17:33</span>
          <span>NFC:{{ nfcId || '0000' }}</span>
        </div>
      </div>

      <p class="nc-t3-outside">ROOT ACCESS IS TEMPORARY · THE NIGHT IS NOT</p>
    </div>
  </div>
</template>

<script setup>
defineProps(['payload', 'nfcId'])
</script>

<style scoped>
.nc-t3-root {
  background: #020a06;
  font-family: 'SFMono-Regular', Consolas, 'Courier New', monospace;
}
.nc-t3-root::before {
  content: '';
  position: absolute;
  inset: 0;
  background:
    radial-gradient(ellipse at 50% -10%, rgba(57, 255, 136, 0.11), transparent 55%),
    repeating-linear-gradient(180deg, rgba(0, 0, 0, 0) 0 2px, rgba(0, 0, 0, 0.24) 2px 4px);
  pointer-events: none;
}
.nc-t3-noise {
  position: absolute;
  inset: 0;
  opacity: 0.35;
  background-image: radial-gradient(rgba(141, 255, 176, 0.2) 0.5px, transparent 0.5px);
  background-size: 3px 3px;
  pointer-events: none;
}
.nc-t3-shell {
  display: flex;
  flex: 1;
  flex-direction: column;
  overflow: hidden;
  border: 1px solid rgba(57, 255, 136, 0.34);
  background: rgba(2, 17, 10, 0.92);
  box-shadow: 0 0 0 1px rgba(0, 0, 0, 0.8), 0 0 60px -10px rgba(57, 255, 136, 0.24);
  animation: nc-t3-boot 0.7s steps(8, end) both;
}
.nc-t3-bar {
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 8px 11px;
  border-bottom: 1px solid rgba(57, 255, 136, 0.2);
  background: linear-gradient(180deg, rgba(57, 255, 136, 0.13), rgba(57, 255, 136, 0.04));
  font-size: 8px;
  letter-spacing: 0.08em;
}
.nc-t3-dots {
  display: inline-flex;
  gap: 4px;
}
.nc-t3-dots i {
  width: 6px;
  height: 6px;
  border: 1px solid rgba(141, 255, 176, 0.5);
  border-radius: 50%;
}
.nc-t3-dots i:nth-child(2) { opacity: 0.7; }
.nc-t3-dots i:nth-child(3) { opacity: 0.4; }
.nc-t3-path {
  flex: 1;
  text-align: center;
  color: rgba(141, 255, 176, 0.65);
}
.nc-t3-win {
  letter-spacing: 0.2em;
  color: rgba(141, 255, 176, 0.4);
}
.nc-t3-screen {
  flex: 1;
  padding: 22px 18px 18px;
  font-size: 12px;
  line-height: 1.6;
}
.nc-t3-line,
.nc-t3-cmd {
  margin: 0;
}
.nc-t3-dim {
  font-size: 9px;
  color: rgba(141, 255, 176, 0.4);
}
.nc-t3-gap {
  height: 18px;
}
.nc-t3-block {
  margin-bottom: 15px;
  animation: nc-t3-type 0.5s both;
}
.nc-t3-block:nth-of-type(2) { animation-delay: 0.15s; }
.nc-t3-block:nth-of-type(3) { animation-delay: 0.3s; }
.nc-t3-block:nth-of-type(4) { animation-delay: 0.45s; }
.nc-t3-block:nth-of-type(5) { animation-delay: 0.6s; }
.nc-t3-block:nth-of-type(6) { animation-delay: 0.75s; }
.nc-t3-prompt {
  margin-right: 7px;
  color: #d9ffe6;
  font-weight: 700;
  text-shadow: 0 0 8px rgba(57, 255, 136, 0.8);
}
.nc-t3-cmd {
  font-size: 10.5px;
  color: rgba(141, 255, 176, 0.7);
}
.nc-t3-output {
  display: block;
  margin: 5px 0 0 17px;
  padding-left: 11px;
  border-left: 1px solid rgba(57, 255, 136, 0.28);
  text-decoration: none;
}
.nc-t3-name {
  margin: 0;
  font-size: 30px;
  line-height: 1.1;
  letter-spacing: 0.04em;
  color: #dcffe8;
  text-shadow: 0 0 18px rgba(57, 255, 136, 0.5);
  overflow-wrap: anywhere;
}
.nc-t3-alias {
  margin: 2px 0 0;
  font-size: 9px;
  letter-spacing: 0.34em;
  color: #39ff88;
}
.nc-t3-value {
  font-size: 12.5px;
  color: #b5ffcc;
}
.nc-t3-pathlist {
  display: flex;
  flex-wrap: wrap;
  align-items: baseline;
  gap: 5px;
  font-size: 12px;
}
.nc-t3-dir {
  color: rgba(141, 255, 176, 0.45);
}
.nc-t3-venue {
  color: #dcffe8;
  font-weight: 700;
  letter-spacing: 0.06em;
}
.nc-t3-city {
  padding: 0 5px;
  border: 1px solid rgba(57, 255, 136, 0.3);
  font-size: 8px;
  letter-spacing: 0.16em;
  color: rgba(141, 255, 176, 0.7);
}
.nc-t3-phone {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 16px;
  letter-spacing: 0.08em;
  color: #39ff88;
  text-shadow: 0 0 12px rgba(57, 255, 136, 0.45);
  transition: background 0.25s ease, color 0.25s ease, padding 0.25s ease;
}
.nc-t3-phone:hover {
  padding-left: 16px;
  background: rgba(57, 255, 136, 0.1);
  color: #dcffe8;
}
.nc-t3-phone:active {
  background: rgba(57, 255, 136, 0.2);
}
.nc-t3-hint {
  font-size: 7.5px;
  letter-spacing: 0.14em;
  color: rgba(141, 255, 176, 0.42);
}
.nc-t3-muted {
  color: rgba(141, 255, 176, 0.35);
}
.nc-t3-cursorline {
  margin-top: 20px;
  display: flex;
  align-items: center;
}
.nc-t3-cursor {
  display: inline-block;
  width: 8px;
  height: 14px;
  margin-left: 2px;
  background: #39ff88;
  box-shadow: 0 0 10px rgba(57, 255, 136, 0.8);
  animation: nc-t3-blink 1.1s step-end infinite;
}
.nc-t3-status {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 8px;
  padding: 7px 11px;
  border-top: 1px solid rgba(57, 255, 136, 0.2);
  background: rgba(57, 255, 136, 0.07);
  font-size: 7px;
  letter-spacing: 0.14em;
  color: rgba(141, 255, 176, 0.55);
}
.nc-t3-status span:first-child {
  color: #39ff88;
}
.nc-t3-status span:first-child::before {
  content: '● ';
  animation: nc-t3-blink 2s step-end infinite;
}
.nc-t3-outside {
  padding: 18px 2px 6px;
  text-align: center;
  font-size: 7px;
  letter-spacing: 0.22em;
  color: rgba(141, 255, 176, 0.25);
  animation: nc-t3-type 1s 0.9s both;
}
@keyframes nc-t3-boot {
  from { opacity: 0; transform: scaleY(0.85); }
  to { opacity: 1; transform: none; }
}
@keyframes nc-t3-type {
  from { opacity: 0; transform: translateX(-6px); }
  to { opacity: 1; transform: none; }
}
@keyframes nc-t3-blink {
  0%, 49% { opacity: 1; }
  50%, 100% { opacity: 0; }
}
@media (prefers-reduced-motion: reduce) {
  .nc-t3-root *,
  .nc-t3-root *::before,
  .nc-t3-root *::after {
    animation: none !important;
    transition: none !important;
  }
}
</style>
