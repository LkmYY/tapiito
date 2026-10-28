<template>
  <div class="min-h-screen bg-white text-stone-900">
    <!-- Warning band -->
    <header class="rise relative overflow-hidden bg-red-600 px-6 pb-7 pt-9 text-white">
      <div
        class="pointer-events-none absolute inset-0 opacity-[0.09]"
        style="background-image: repeating-linear-gradient(-45deg, #fff 0 1px, transparent 1px 16px)"
      ></div>
      <div class="relative">
        <div class="flex items-center justify-between font-mono text-[10px] uppercase tracking-[0.24em] text-red-100/80">
          <span>Emergency Contact</span>
          <span v-if="nfcId" class="max-w-[45%] truncate text-right">NFC · {{ nfcId }}</span>
        </div>
        <div class="mt-7 flex items-center gap-3.5">
          <div class="flex h-11 w-11 flex-shrink-0 items-center justify-center rounded-xl bg-white/15 ring-1 ring-white/25">
            <Siren class="h-6 w-6" />
          </div>
          <h1 class="text-[26px] font-bold leading-tight tracking-tight">紧急联系人</h1>
        </div>
        <p class="mt-4 text-sm leading-relaxed text-red-50/90">如发现持卡人遇到紧急情况，请立即拨打下方电话。</p>
      </div>
    </header>

    <div class="px-6 pb-12">
      <!-- Call list -->
      <section class="rise rise-1 mt-8">
        <div class="flex items-center justify-between">
          <p class="font-mono text-[10px] uppercase tracking-[0.24em] text-stone-400">Call List · 拨打列表</p>
          <span v-if="contacts.length" class="font-mono text-[10px] tabular-nums text-stone-400">{{ String(contacts.length).padStart(2, '0') }}</span>
        </div>

        <div v-if="contacts.length" class="mt-1.5">
          <article v-for="(c, i) in contacts" :key="i" class="border-t border-stone-200 pb-6 pt-5">
            <div class="flex items-baseline gap-3">
              <span class="font-mono text-[11px] tabular-nums text-red-600">{{ String(i + 1).padStart(2, '0') }}</span>
              <h2 class="min-w-0 flex-1 truncate text-[17px] font-semibold tracking-tight">{{ c.name || '联系人' }}</h2>
              <span v-if="c.relation" class="flex-shrink-0 rounded-full border border-stone-200 px-2.5 py-0.5 text-[11px] text-stone-500">{{ c.relation }}</span>
            </div>
            <a
              :href="'tel:' + c.phone"
              class="mt-4 flex items-center justify-between rounded-2xl bg-red-600 px-5 py-4 text-white shadow-[0_14px_30px_-14px_rgba(220,38,38,0.55)] transition-all hover:bg-red-700 active:scale-[0.98]"
            >
              <span class="flex min-w-0 items-center gap-3">
                <Phone class="h-5 w-5 flex-shrink-0 text-red-100" />
                <span class="truncate text-[19px] font-semibold tabular-nums tracking-wide">{{ c.phone }}</span>
              </span>
              <span class="ml-3 flex-shrink-0 font-mono text-[10px] uppercase tracking-[0.2em] text-red-100/85">拨打</span>
            </a>
          </article>
        </div>

        <div v-else class="mt-1.5 flex flex-col items-center rounded-2xl border border-dashed border-stone-300 px-6 py-10 text-center">
          <AlertTriangle class="h-6 w-6 text-stone-300" />
          <p class="mt-3 text-sm text-stone-500">暂无紧急联系人信息</p>
        </div>
      </section>

      <!-- Message & photos -->
      <section class="rise rise-2 mt-2 border-t border-stone-200 pt-7">
        <div class="flex items-center gap-2">
          <MessageSquare class="h-3.5 w-3.5 text-stone-400" />
          <p class="font-mono text-[10px] uppercase tracking-[0.24em] text-stone-400">Leave a Message · 留言</p>
        </div>
        <textarea
          v-model="note.content"
          rows="3"
          class="mt-4 w-full resize-none rounded-xl border border-stone-200 bg-stone-50 p-3.5 text-sm text-stone-800 placeholder:text-stone-400 focus:border-red-500 focus:outline-none focus:ring-2 focus:ring-red-500/15"
          placeholder="请输入您的留言或位置信息..."
        ></textarea>
        <button
          @click="handleSubmitNote"
          :disabled="submitting"
          class="mt-3 flex w-full items-center justify-center gap-2 rounded-xl bg-stone-900 py-3.5 text-sm font-semibold text-white transition-all hover:bg-stone-800 active:scale-[0.98] disabled:opacity-60"
        >
          <Send v-if="!submitting" class="h-4 w-4" />
          <Loader2 v-else class="h-4 w-4 animate-spin" />
          {{ submitting ? '发送中...' : '发送留言' }}
        </button>

        <div class="mt-7 grid grid-cols-3 gap-2">
          <div v-for="p in photos" :key="p.id" class="aspect-square overflow-hidden rounded-xl border border-stone-200 bg-stone-100">
            <img :src="photoUrl(p.oss_key)" class="h-full w-full object-cover" />
          </div>
          <label class="flex aspect-square cursor-pointer flex-col items-center justify-center gap-1 rounded-xl border-2 border-dashed border-stone-300 text-stone-400 transition-colors hover:bg-stone-50">
            <Camera class="h-5 w-5" />
            <span class="text-[11px]">上传</span>
            <input type="file" accept="image/*" capture="environment" class="hidden" @change="onFile" />
          </label>
        </div>
      </section>

      <p class="rise rise-3 mt-10 text-center font-mono text-[10px] uppercase tracking-[0.24em] text-stone-300">Tap number to call · 一键拨打</p>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted, computed } from 'vue'
import { Siren, Phone, AlertTriangle, MessageSquare, Send, Loader2, Camera } from 'lucide-vue-next'
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
    ElMessage.success('已提交并已发送')
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
