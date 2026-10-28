<template>
  <div class="relative min-h-screen overflow-hidden bg-[#f8f6f0] text-stone-800">
    <!-- Paw watermark -->
    <PawPrint class="pointer-events-none absolute -right-10 -top-8 h-44 w-44 rotate-12 text-[#2f6f4e]/5" />

    <div class="relative px-6 pb-12 pt-10">
      <div class="rise flex items-center justify-between font-mono text-[10px] uppercase tracking-[0.24em] text-stone-400">
        <span>Found Pet</span>
        <span v-if="nfcId" class="max-w-[45%] truncate text-right">NFC · {{ nfcId }}</span>
      </div>

      <!-- Header -->
      <header class="rise rise-1 mt-8">
        <div class="flex items-center gap-3.5">
          <div class="flex h-12 w-12 flex-shrink-0 items-center justify-center rounded-2xl bg-[#e8f1e8] text-[#2f6f4e] ring-1 ring-[#2f6f4e]/10">
            <PawPrint class="h-6 w-6" />
          </div>
          <p class="font-mono text-[10px] uppercase leading-relaxed tracking-[0.2em] text-stone-400">Lost &amp; Found<br />走失招领</p>
        </div>
        <h1 class="mt-6 font-serif text-[36px] leading-tight tracking-tight text-stone-900">{{ payload.petName || '我的宠物' }}</h1>
        <p v-if="payload.breed" class="mt-2 text-[13px] uppercase tracking-[0.18em] text-stone-400">{{ payload.breed }}</p>
        <p class="mt-4 text-sm leading-relaxed text-stone-500">如果发现我独自在外，请联系我的主人接我回家，谢谢。</p>
      </header>

      <!-- Chip -->
      <section v-if="payload.chipId" class="rise rise-2 mt-8 flex items-center gap-4 rounded-2xl border border-stone-200 bg-white px-5 py-4">
        <Fingerprint class="h-5 w-5 flex-shrink-0 text-[#2f6f4e]" />
        <div class="min-w-0 flex-1">
          <p class="font-mono text-[10px] uppercase tracking-[0.2em] text-stone-400">Chip · 芯片号</p>
          <p class="mt-1 truncate font-mono text-[14px] font-medium tracking-wide text-stone-800">{{ payload.chipId }}</p>
        </div>
      </section>

      <!-- Owner contacts -->
      <section class="rise rise-2 mt-8">
        <p class="font-mono text-[10px] uppercase tracking-[0.24em] text-stone-400">Owner · 主人联系方式</p>
        <div v-if="contacts.length" class="mt-3 space-y-3">
          <article
            v-for="(c, i) in contacts"
            :key="i"
            class="rounded-2xl bg-white p-5 shadow-[0_18px_36px_-28px_rgba(47,111,78,0.5)] ring-1 ring-stone-900/[0.04]"
          >
            <div class="flex items-center justify-between gap-4">
              <div class="min-w-0">
                <p class="truncate text-[15px] font-semibold text-stone-900">{{ c.name || '主人' }}</p>
                <p v-if="c.relation" class="mt-0.5 font-mono text-[10px] uppercase tracking-[0.16em] text-[#2f6f4e]">{{ c.relation }}</p>
              </div>
              <a
                :href="'tel:' + c.phone"
                class="flex h-12 w-12 flex-shrink-0 items-center justify-center rounded-full bg-[#2f6f4e] text-white shadow-[0_12px_24px_-10px_rgba(47,111,78,0.65)] transition-all hover:bg-[#255a3f] active:scale-[0.95]"
              >
                <Phone class="h-5 w-5" />
              </a>
            </div>
            <a
              :href="'tel:' + c.phone"
              class="mt-3.5 block border-t border-stone-100 pt-3.5 text-[20px] font-semibold tabular-nums tracking-wide text-stone-900 transition-colors hover:text-[#2f6f4e]"
            >
              {{ c.phone }}
            </a>
          </article>
        </div>
        <div v-else class="mt-3 rounded-2xl border border-dashed border-stone-300 bg-white/70 px-6 py-10 text-center text-sm text-stone-500">暂无主人联系方式</div>
      </section>

      <!-- Vet -->
      <section v-if="payload.vetPhone" class="rise rise-3 mt-6">
        <a
          :href="'tel:' + payload.vetPhone"
          class="flex items-center gap-4 rounded-2xl bg-[#e8f1e8] px-5 py-4 ring-1 ring-[#2f6f4e]/10 transition-colors hover:bg-[#dfebdf]"
        >
          <div class="flex h-10 w-10 flex-shrink-0 items-center justify-center rounded-xl bg-white text-[#2f6f4e] ring-1 ring-[#2f6f4e]/10">
            <Stethoscope class="h-5 w-5" />
          </div>
          <div class="min-w-0 flex-1">
            <p class="font-mono text-[10px] uppercase tracking-[0.2em] text-[#2f6f4e]/70">Vet · 兽医电话</p>
            <p class="mt-0.5 text-[18px] font-semibold tabular-nums tracking-wide text-[#255a3f]">{{ payload.vetPhone }}</p>
          </div>
          <Phone class="h-4 w-4 flex-shrink-0 text-[#2f6f4e]/60" />
        </a>
      </section>

      <!-- Message & photos -->
      <section class="rise rise-3 mt-8 border-t border-stone-200 pt-7">
        <p class="font-mono text-[10px] uppercase tracking-[0.24em] text-stone-400">Message · 留言</p>
        <textarea
          v-model="note.content"
          rows="3"
          class="mt-4 w-full resize-none rounded-xl border border-stone-200 bg-white p-3.5 text-sm text-stone-800 placeholder:text-stone-400 focus:border-[#2f6f4e] focus:outline-none focus:ring-2 focus:ring-[#2f6f4e]/15"
          placeholder="请描述发现宠物的位置..."
        ></textarea>
        <button
          @click="handleSubmitNote"
          :disabled="submitting"
          class="mt-3 flex w-full items-center justify-center gap-2 rounded-xl bg-[#2f6f4e] py-3.5 text-sm font-semibold text-white transition-all hover:bg-[#255a3f] active:scale-[0.98] disabled:opacity-60"
        >
          <Send v-if="!submitting" class="h-4 w-4" />
          <Loader2 v-else class="h-4 w-4 animate-spin" />
          {{ submitting ? '发送中...' : '发送留言' }}
        </button>

        <div class="mt-7 grid grid-cols-3 gap-2">
          <div v-for="p in photos" :key="p.id" class="aspect-square overflow-hidden rounded-xl border border-stone-200 bg-stone-100">
            <img :src="photoUrl(p.oss_key)" class="h-full w-full object-cover" />
          </div>
          <label class="flex aspect-square cursor-pointer flex-col items-center justify-center gap-1 rounded-xl border-2 border-dashed border-[#2f6f4e]/25 text-[#2f6f4e]/50 transition-colors hover:bg-[#e8f1e8]">
            <Camera class="h-5 w-5" />
            <span class="text-[11px]">上传</span>
            <input type="file" accept="image/*" capture="environment" class="hidden" @change="onFile" />
          </label>
        </div>
      </section>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted, computed } from 'vue'
import { PawPrint, Fingerprint, Stethoscope, Phone, Send, Loader2, Camera } from 'lucide-vue-next'
import { ElMessage } from 'element-plus'
import { fetchPhotos, uploadPhoto, submitNote as apiSubmitNote } from '../../api/public'

const props = defineProps(['payload', 'nfcId'])
const note = reactive({ content: '' })
const photos = ref([])
const submitting = ref(false)

const contacts = computed(() =>
  (Array.isArray(props.payload?.contacts) ? props.payload.contacts : []).filter((c) => c && c.phone)
)

const photoUrl = (key) => {
  if (!key) return ''
  return key.startsWith('http') ? key : '/uploads/' + key
}

const loadPhotos = async () => {
  // 预览态（PREVIEW）没有真实标签，跳过请求，避免 404
  if (!props.nfcId || props.nfcId === 'PREVIEW') return
  try {
    const { data } = await fetchPhotos(props.nfcId)
    photos.value = data || []
  } catch {
    // 错误已由 request 拦截器处理
  }
}

const handleSubmitNote = async () => {
  if (!note.content) return
  submitting.value = true
  try {
    await apiSubmitNote(props.nfcId, {
      content: note.content,
      photo_ids: photos.value.map((p) => p.id)
    })
    ElMessage.success('已提交')
    note.content = ''
  } catch {
    ElMessage.error('提交失败')
  } finally {
    submitting.value = false
  }
}

const onFile = async (e) => {
  const file = e.target.files[0]
  if (!file) return
  const formData = new FormData()
  formData.append('file', file)
  try {
    await uploadPhoto(props.nfcId, formData)
    await loadPhotos()
  } catch {
    ElMessage.error('上传失败')
  }
}

onMounted(loadPhotos)
</script>

<style scoped>
@keyframes rise {
  from {
    opacity: 0;
    transform: translateY(14px);
  }
  to {
    opacity: 1;
    transform: none;
  }
}
.rise {
  animation: rise 0.65s cubic-bezier(0.22, 1, 0.36, 1) both;
}
.rise-1 {
  animation-delay: 0.08s;
}
.rise-2 {
  animation-delay: 0.16s;
}
.rise-3 {
  animation-delay: 0.24s;
}
</style>
