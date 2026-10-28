<template>
  <footer
    class="relative mt-auto border-t transition-colors duration-500"
    :class="tone === 'dark'
      ? 'border-white/[0.08] bg-[#0a0d12]'
      : 'border-black/[0.07] bg-[#fbfbfa] dark:border-white/[0.08] dark:bg-[#0a0d12]'"
  >
    <a
      href="/"
      class="group flex items-center justify-between gap-3 px-5 py-3.5 transition-colors duration-300"
      :class="tone === 'dark'
        ? 'hover:bg-white/[0.04]'
        : 'hover:bg-black/[0.025] dark:hover:bg-white/[0.04]'"
    >
      <div class="flex min-w-0 items-center gap-2.5">
        <span
          class="flex h-7 w-7 shrink-0 items-center justify-center rounded-[9px] transition-transform duration-300 group-hover:-translate-y-px"
          :class="tone === 'dark'
            ? 'bg-[#15151B] ring-1 ring-inset ring-white/10'
            : 'bg-[#15151B] shadow-[0_1px_2px_rgba(2,6,23,0.24)]'"
        >
          <img src="/tapiito-icon-white.png" alt="Tapiito" class="h-[17px] w-[17px]" draggable="false" />
        </span>
        <span class="min-w-0 leading-tight">
          <span
            class="block truncate text-[11.5px] font-semibold tracking-tight"
            :class="tone === 'dark' ? 'text-white/85' : 'text-[#14181d] dark:text-white/85'"
          >
            Tapiito
          </span>
          <span
            class="block truncate text-[10px] tracking-[0.08em]"
            :class="tone === 'dark' ? 'text-white/35' : 'text-[#14181d]/45 dark:text-white/35'"
          >
            Tap it alive. 碰一下，它就活了
          </span>
        </span>
      </div>

      <span
        class="flex shrink-0 items-center gap-1 rounded-full border px-2.5 py-1 text-[10px] font-medium tracking-wide transition-colors duration-300"
        :class="tone === 'dark'
          ? 'border-white/15 text-white/55 group-hover:border-white/30 group-hover:text-white/85'
          : 'border-black/[0.09] text-[#14181d]/60 group-hover:border-black/[0.18] group-hover:text-[#14181d] dark:border-white/15 dark:text-white/55 dark:group-hover:border-white/30 dark:group-hover:text-white/85'"
      >
        访问首页
        <ArrowUpRight class="h-3 w-3 transition-transform duration-300 group-hover:-translate-y-px group-hover:translate-x-px" />
      </span>
    </a>
  </footer>
</template>

<script setup>
import { Nfc, ArrowUpRight } from 'lucide-vue-next'

defineProps({
  // 'dark' 用于深色模板底部，避免白条割裂画面
  tone: { type: String, default: 'light' }
})
</script>
