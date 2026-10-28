<template>
  <div class="relative min-h-screen bg-[#0a0a0a] text-white">
    <!-- Hazard stripe -->
    <div class="h-2.5 w-full" style="background-image: repeating-linear-gradient(-45deg, #dc2626 0 10px, #0a0a0a 10px 20px)"></div>

    <!-- SOS plate -->
    <header class="rise bg-red-600 px-6 pb-7 pt-8">
      <div class="flex items-center justify-between font-mono text-[10px] uppercase tracking-[0.26em] text-red-100/85">
        <span class="flex items-center gap-2">
          <span class="h-1.5 w-1.5 animate-pulse rounded-full bg-white"></span>
          Emergency · 紧急求助
        </span>
        <span v-if="nfcId" class="max-w-[40%] truncate text-right">NFC · {{ nfcId }}</span>
      </div>
      <div class="mt-6 flex items-end justify-between gap-4">
        <h1 class="text-[64px] font-black leading-[0.9] tracking-[0.05em]">SOS</h1>
        <Siren class="mb-1.5 h-9 w-9 flex-shrink-0 text-white/85" />
      </div>
      <p class="mt-4 text-[13px] leading-relaxed text-red-50/90">请立即拨打下方电话，并告知持卡人所处位置。</p>
    </header>

    <div class="px-6 pb-12">
      <section class="rise rise-1 mt-7">
        <template v-if="contacts.length">
          <!-- Primary contact -->
          <a
            :href="'tel:' + contacts[0].phone"
            class="block rounded-2xl bg-red-600 px-6 py-6 text-center shadow-[0_22px_48px_-18px_rgba(220,38,38,0.7)] transition-all hover:bg-red-500 active:scale-[0.98]"
          >
            <p class="font-mono text-[10px] uppercase tracking-[0.3em] text-red-100/85">Primary · 首要联系人</p>
            <p class="mt-3 text-[36px] font-black leading-none tabular-nums tracking-wide">{{ contacts[0].phone }}</p>
            <p class="mt-3 text-[13px] font-medium text-red-50">
              {{ contacts[0].name || '联系人' }}<template v-if="contacts[0].relation"> · {{ contacts[0].relation }}</template>
            </p>
            <span class="mt-4 inline-flex items-center gap-2 rounded-full bg-black/25 px-4 py-1.5 text-[11px] font-bold uppercase tracking-[0.18em] text-white">
              <Phone class="h-3.5 w-3.5" />
              点击立即拨打
            </span>
          </a>

          <!-- Other contacts -->
          <div class="mt-4 space-y-3">
            <a
              v-for="(c, i) in contacts.slice(1)"
              :key="i"
              :href="'tel:' + c.phone"
              class="flex items-center justify-between gap-4 rounded-2xl border-2 border-red-600/50 bg-white/[0.03] px-5 py-4 transition-colors hover:border-red-500 hover:bg-red-600/10 active:scale-[0.98]"
            >
              <div class="min-w-0">
                <p class="truncate text-[22px] font-bold leading-tight tabular-nums tracking-wide">{{ c.phone }}</p>
                <p class="mt-1 truncate text-[12px] text-white/55">
                  {{ c.name || '联系人' }}<template v-if="c.relation"> · {{ c.relation }}</template>
                </p>
              </div>
              <span class="flex h-11 w-11 flex-shrink-0 items-center justify-center rounded-full bg-red-600 text-white">
                <Phone class="h-5 w-5" />
              </span>
            </a>
          </div>
        </template>

        <div v-else class="rounded-2xl border-2 border-dashed border-red-600/40 px-6 py-12 text-center">
          <ShieldAlert class="mx-auto h-7 w-7 text-red-500/70" />
          <p class="mt-3 text-sm text-white/50">暂无紧急联系人 · NO CONTACTS</p>
        </div>
      </section>

      <!-- Message & photos -->
      <section class="rise rise-2 mt-10 border-t border-white/10 pt-7">
        <p class="font-mono text-[10px] uppercase tracking-[0.26em] text-white/35">Leave Info · 留言与现场照片</p>
        <textarea
          v-model="note.content"
          rows="3"
          class="mt-4 w-full resize-none rounded-xl border border-white/15 bg-white/[0.05] p-3.5 text-sm text-white placeholder:text-white/30 focus:border-red-500 focus:outline-none focus:ring-2 focus:ring-red-500/25"
          placeholder="输入留言或位置..."
        ></textarea>
        <button
          @click="handleSubmitNote"
          :disabled="submitting"
          class="mt-3 flex w-full items-center justify-center gap-2 rounded-xl bg-white py-3.5 text-sm font-bold text-[#0a0a0a] transition-all hover:bg-red-50 active:scale-[0.98] disabled:opacity-60"
        >
          <Send v-if="!submitting" class="h-4 w-4" />
          <Loader2 v-else class="h-4 w-4 animate-spin" />
          {{ submitting ? '发送中...' : '发送留言' }}
        </button>

        <div class="mt-7 grid grid-cols-3 gap-2">
          <div v-for="p in photos" :key="p.id" class="aspect-square overflow-hidden rounded-xl border border-white/10 bg-white/[0.05]">
            <img :src="photoUrl(p.oss_key)" class="h-full w-full object-cover" />
          </div>
          <label class="flex aspect-square cursor-pointer flex-col items-center justify-center gap-1 rounded-xl border-2 border-dashed border-white/25 text-white/40 transition-colors hover:bg-white/[0.04]">
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
import { Siren, Phone, ShieldAlert, Send, Loader2, Camera } from 'lucide-vue-next'
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
