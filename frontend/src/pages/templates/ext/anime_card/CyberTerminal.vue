<template>
  <div class="ac-term-root relative min-h-screen overflow-hidden px-4 py-8">
    <div class="ac-term-scan"></div>
    <div class="ac-term-flicker"></div>
    <div class="relative z-10 mx-auto w-full max-w-md">
      <div class="ac-term-window">
        <header class="ac-term-bar">
          <span class="ac-term-dots"><i></i><i></i><i></i></span>
          <span>root@nfc:~/anime_card</span>
          <span class="ac-term-lock"><Lock :size="10" />SECURE</span>
        </header>

        <div class="ac-term-body">
          <div class="ac-term-boot">
            <p><span>[ OK ]</span> mounting /dev/nfc{{ nfcId || '0000' }}</p>
            <p><span>[ OK ]</span> loading character.dat</p>
            <p><span>[ OK ]</span> decrypting quote.txt ... done</p>
          </div>

          <div class="ac-term-prompt">
            <span class="ac-term-sign">$</span>
            <span>whoami --card</span>
          </div>

          <div class="ac-term-profile">
            <p><i>NAME</i><b>{{ p.name || '未命名角色' }}</b></p>
            <p><i>TITLE</i><b>{{ p.title }}</b></p>
            <p><i>SERIES</i><b>{{ p.series || '未知作品' }}</b></p>
            <p><i>ELEMENT</i><b class="ac-term-elem">{{ p.element }}</b></p>
            <p><i>SOCIAL</i><b class="ac-term-social">@{{ p.social }}</b></p>
          </div>

          <div class="ac-term-prompt">
            <span class="ac-term-sign">$</span>
            <span>cat quote.txt</span>
          </div>

          <blockquote class="ac-term-quote">
            <p>“{{ p.quote || '……connection established, but no words found.' }}”</p>
          </blockquote>

          <div class="ac-term-prompt ac-term-last">
            <span class="ac-term-sign">$</span>
            <span>_</span>
            <i class="ac-term-cursor"></i>
          </div>
        </div>

        <footer class="ac-term-foot">
          <span>MEM 64% · PING 12ms</span>
          <span>NFC://{{ nfcId || '0000' }}</span>
        </footer>
      </div>

      <p class="ac-term-caption"><Terminal :size="10" /> SESSION LOGGED · DO NOT CLOSE</p>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { Lock, Terminal } from 'lucide-vue-next'

const props = defineProps(['payload', 'nfcId'])
const p = computed(() => props.payload || {})
</script>

<style scoped>
.ac-term-root {
  background: #020604;
  font-family: Consolas, 'Courier New', monospace;
}
.ac-term-scan {
  position: absolute;
  inset: 0;
  background: repeating-linear-gradient(180deg, transparent 0 2px, rgba(34, 197, 94, 0.055) 2px 3px);
  pointer-events: none;
}
.ac-term-flicker {
  position: absolute;
  inset: 0;
  background: radial-gradient(ellipse at 50% 40%, rgba(34, 197, 94, 0.07), transparent 60%);
  animation: ac-term-flicker 5s steps(3) infinite;
  pointer-events: none;
}
.ac-term-window {
  border: 1px solid #166534;
  background: rgba(2, 10, 5, 0.9);
  box-shadow:
    0 0 0 1px rgba(34, 197, 94, 0.12),
    0 0 40px rgba(34, 197, 94, 0.14),
    inset 0 0 60px rgba(21, 128, 61, 0.07);
  animation: ac-term-boot 0.7s steps(4) both;
  transition: box-shadow 0.4s ease, border-color 0.4s ease;
}
.ac-term-window:hover {
  border-color: #22c55e;
  box-shadow:
    0 0 0 1px rgba(34, 197, 94, 0.2),
    0 0 60px rgba(34, 197, 94, 0.24),
    inset 0 0 70px rgba(21, 128, 61, 0.1);
}
.ac-term-bar {
  display: flex;
  align-items: center;
  justify-content: space-between;
  border-bottom: 1px solid #14532d;
  background: #071b0d;
  padding: 6px 8px;
  font-size: 7px;
  letter-spacing: 0.08em;
  color: #4ade80;
}
.ac-term-dots { display: flex; gap: 3px; }
.ac-term-dots i { width: 5px; height: 5px; border-radius: 50%; background: #166534; }
.ac-term-dots i:first-child { background: #f87171; }
.ac-term-dots i:nth-child(2) { background: #fbbf24; }
.ac-term-lock { display: inline-flex; align-items: center; gap: 3px; color: #86efac; font-size: 6px; }
.ac-term-body { padding: 14px 12px 10px; }
.ac-term-boot {
  border-bottom: 1px dashed rgba(34, 197, 94, 0.2);
  padding-bottom: 9px;
  margin-bottom: 13px;
}
.ac-term-boot p {
  font-size: 7.5px;
  line-height: 1.7;
  color: rgba(134, 239, 172, 0.5);
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}
.ac-term-boot span { color: #22c55e; }
.ac-term-prompt {
  display: flex;
  align-items: center;
  gap: 7px;
  font-size: 10px;
  color: rgba(134, 239, 172, 0.85);
}
.ac-term-sign { color: #4ade80; text-shadow: 0 0 8px rgba(74, 222, 128, 0.7); }
.ac-term-profile {
  margin: 10px 0 14px;
  border-left: 1px solid rgba(34, 197, 94, 0.3);
  padding-left: 10px;
}
.ac-term-profile p {
  display: flex;
  gap: 8px;
  padding: 2px 0;
  font-size: 9.5px;
  line-height: 1.65;
  overflow-wrap: anywhere;
}
.ac-term-profile i {
  flex: none;
  width: 60px;
  font-style: normal;
  font-size: 7px;
  letter-spacing: 0.14em;
  color: rgba(134, 239, 172, 0.36);
  padding-top: 2px;
}
.ac-term-profile b {
  min-width: 0;
  flex: 1;
  color: #bbf7d0;
  font-weight: 400;
  overflow-wrap: anywhere;
}
.ac-term-profile p:first-child b {
  font-size: 27px;
  font-weight: 700;
  line-height: 1.15;
  letter-spacing: 0.02em;
  color: #4ade80;
  text-shadow: 0 0 16px rgba(34, 197, 94, 0.55);
}
.ac-term-elem { color: #67e8f9 !important; }
.ac-term-social { color: #fde68a !important; }
.ac-term-quote {
  margin-top: 9px;
  border: 1px dashed rgba(34, 197, 94, 0.28);
  background: rgba(34, 197, 94, 0.045);
  padding: 11px 10px;
}
.ac-term-quote p {
  font-family: 'PingFang SC', 'Microsoft YaHei', monospace;
  font-size: 12.5px;
  line-height: 1.85;
  color: #d1fae5;
  overflow-wrap: anywhere;
}
.ac-term-last { margin-top: 14px; }
.ac-term-cursor {
  width: 7px;
  height: 13px;
  background: #4ade80;
  box-shadow: 0 0 10px rgba(74, 222, 128, 0.8);
  animation: ac-term-blink 1s steps(2) infinite;
}
.ac-term-foot {
  display: flex;
  justify-content: space-between;
  border-top: 1px solid #14532d;
  background: #071b0d;
  padding: 5px 8px;
  font-size: 6.5px;
  letter-spacing: 0.12em;
  color: rgba(134, 239, 172, 0.4);
}
.ac-term-caption {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 5px;
  margin-top: 20px;
  font-size: 7px;
  letter-spacing: 0.2em;
  color: rgba(74, 222, 128, 0.3);
  animation: ac-term-fade 1.3s 0.6s both;
}
@keyframes ac-term-boot {
  from { opacity: 0; transform: scaleY(0.6); filter: brightness(2.6); }
  to { opacity: 1; transform: none; filter: none; }
}
@keyframes ac-term-fade { from { opacity: 0; } to { opacity: 1; } }
@keyframes ac-term-blink { 0%, 49% { opacity: 1; } 50%, 100% { opacity: 0; } }
@keyframes ac-term-flicker { 0%, 96%, 100% { opacity: 1; } 97% { opacity: 0.7; } 98% { opacity: 1; } 99% { opacity: 0.85; } }
@media (prefers-reduced-motion: reduce) {
  .ac-term-root *,
  .ac-term-root *::before,
  .ac-term-root *::after {
    animation: none !important;
    transition: none !important;
  }
}
</style>
