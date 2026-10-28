<template>
  <div class="relative flex min-h-screen items-center justify-center overflow-hidden bg-gray-950 text-white">
    <div class="absolute inset-0 w-full h-full">
      <div class="absolute top-[-10%] left-[-10%] w-[40%] h-[40%] rounded-full bg-purple-500/30 blur-[120px] animate-pulse" />
      <div class="absolute bottom-[-10%] right-[-10%] w-[40%] h-[40%] rounded-full bg-blue-500/30 blur-[120px] animate-pulse delay-1000" />
    </div>
    <div
      v-motion
      :initial="{ opacity: 0, y: 50 }"
      :enter="{ opacity: 1, y: 0, transition: { duration: 800, type: 'spring' } }"
      class="relative z-10 w-full max-w-md p-8 rounded-2xl bg-white/10 backdrop-blur-xl border border-white/20 shadow-2xl"
    >
      <div class="mb-8 text-center">
        <h1 class="text-3xl font-bold bg-clip-text text-transparent bg-gradient-to-r from-blue-400 to-purple-400">
          Tapiito
        </h1>
        <p class="text-gray-400 mt-2 text-sm">{{ subtitle }}</p>
      </div>
      <slot />
    </div>
  </div>
</template>

<script setup>
defineProps({
  subtitle: { type: String, default: '' }
})
</script>
