<template>
  <div class="min-h-screen bg-[#faf6f1] text-stone-900">
    <div class="px-6 pb-12 pt-10">
      <div class="rise flex items-center justify-between font-mono text-[10px] uppercase tracking-[0.24em] text-stone-400">
        <span>Family Contacts</span>
        <span v-if="nfcId" class="max-w-[45%] truncate text-right">NFC · {{ nfcId }}</span>
      </div>

      <!-- Header -->
      <header class="rise rise-1 mt-8">
        <div class="flex items-center gap-3.5">
          <div class="flex h-12 w-12 flex-shrink-0 items-center justify-center rounded-2xl bg-[#fbe4d5] text-[#c2410c] ring-1 ring-[#c2410c]/10">
            <HeartHandshake class="h-6 w-6" />
          </div>
          <h1 class="font-serif text-[28px] leading-tight tracking-tight">家庭联系人</h1>
        </div>
        <p class="mt-4 text-sm leading-relaxed text-stone-500">如果家人需要帮助，请随时联系我们。谢谢您的每一份善意。</p>
      </header>

      <!-- Contacts -->
      <section class="rise rise-2 mt-8">
        <div v-if="contacts.length" class="space-y-4">
          <article
            v-for="(c, i) in contacts"
            :key="i"
            class="rounded-2xl bg-white p-5 shadow-[0_18px_38px_-26px_rgba(120,53,15,0.4)] ring-1 ring-stone-900/[0.04]"
          >
            <div class="flex items-center gap-4">
              <div class="flex h-12 w-12 flex-shrink-0 items-center justify-center rounded-full bg-[#fbe4d5] text-[17px] font-semibold text-[#9a3412]">
                {{ initialOf(c.name) }}
              </div>
              <div class="min-w-0 flex-1">
                <p class="truncate text-[16px] font-semibold tracking-tight">{{ c.name || '家人' }}</p>
                <span
                  v-if="c.relation"
                  class="mt-1 inline-block rounded-full bg-[#fdf1e7] px-2.5 py-0.5 text-[11px] font-medium text-[#c2410c] ring-1 ring-[#c2410c]/15"
                >
                  {{ c.relation }}
                </span>
              </div>
              <a
                :href="'tel:' + c.phone"
                class="flex h-12 w-12 flex-shrink-0 items-center justify-center rounded-full bg-[#c2410c] text-white shadow-[0_12px_24px_-10px_rgba(194,65,12,0.6)] transition-all hover:bg-[#9a3412] active:scale-[0.95]"
              >
                <Phone class="h-5 w-5" />
              </a>
            </div>
            <a
              :href="'tel:' + c.phone"
              class="mt-4 flex items-center justify-between rounded-xl bg-[#faf6f1] px-4 py-3 ring-1 ring-stone-900/[0.04] transition-colors hover:bg-[#f6efe7]"
            >
              <span class="text-[19px] font-semibold tabular-nums tracking-wide text-stone-900">{{ c.phone }}</span>
              <span class="font-mono text-[10px] uppercase tracking-[0.2em] text-[#c2410c]/80">拨打</span>
            </a>
          </article>
        </div>

        <div v-else class="flex flex-col items-center rounded-2xl border border-dashed border-[#c2410c]/25 bg-white/60 px-6 py-10 text-center">
          <Heart class="h-5 w-5 text-[#c2410c]/40" />
          <p class="mt-3 text-sm text-stone-500">暂无联系人</p>
        </div>
      </section>

      <!-- Message & photos -->
      <section class="rise rise-3 mt-8 rounded-2xl bg-white p-5 shadow-[0_18px_38px_-30px_rgba(120,53,15,0.35)] ring-1 ring-stone-900/[0.04]">
        <p class="font-mono text-[10px] uppercase tracking-[0.24em] text-stone-400">Message · 留言</p>
        <textarea
          v-model="note.content"
          rows="3"
          class="mt-4 w-full resize-none rounded-xl border border-stone-200 bg-[#faf6f1] p-3.5 text-sm text-stone-800 placeholder:text-stone-400 focus:border-[#c2410c] focus:outline-none focus:ring-2 focus:ring-[#c2410c]/15"
          placeholder="请输入留言..."
        ></textarea>
        <button
          @click="handleSubmitNote"
          :disabled="submitting"
          class="mt-3 flex w-full items-center justify-center gap-2 rounded-xl bg-stone-800 py-3.5 text-sm font-semibold text-white transition-all hover:bg-stone-700 active:scale-[0.98] disabled:opacity-60"
        >
          <Send v-if="!submitting" class="h-4 w-4" />
          <Loader2 v-else class="h-4 w-4 animate-spin" />
          {{ submitting ? '发送中...' : '发送留言' }}
        </button>

        <div class="mt-7 grid grid-cols-3 gap-2">
          <div v-for="p in photos" :key="p.id" class="aspect-square overflow-hidden rounded-xl border border-stone-200 bg-stone-100">
            <img :src="photoUrl(p.oss_key)" class="h-full w-full object-cover" />
          </div>
          <label class="flex aspect-square cursor-pointer flex-col items-center justify-center gap-1 rounded-xl border-2 border-dashed border-[#c2410c]/25 text-[#c2410c]/50 transition-colors hover:bg-[#fdf1e7]">
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
import { HeartHandshake, Heart, Phone, Send, Loader2, Camera } from 'lucide-vue-next'
import { ElMessage } from 'element-plus'
import { fetchPhotos, uploadPhoto, submitNote as apiSubmitNote } from '../../api/public'

const props = defineProps(['payload', 'nfcId'])
const note = reactive({ content: '' })
const photos = ref([])
const submitting = ref(false)

const contacts = computed(() =>
  (Array.isArray(props.payload?.contacts) ? props.payload.contacts : []).filter((c) => c && c.phone)
)

const initialOf = (name) => {
  const s = (name || '').toString().trim()
  return s ? s[0] : '家'
}

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
