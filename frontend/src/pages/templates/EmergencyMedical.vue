<template>
  <div class="min-h-screen bg-[#f6f9fc] text-slate-900">
    <!-- Header -->
    <header class="border-b border-slate-200 bg-white px-6 pb-6 pt-9">
      <div class="rise flex items-center justify-between font-mono text-[10px] uppercase tracking-[0.24em] text-slate-400">
        <span>Medical Alert</span>
        <span v-if="nfcId" class="max-w-[45%] truncate text-right">NFC · {{ nfcId }}</span>
      </div>
      <div class="rise rise-1 mt-7 flex items-start gap-4">
        <div class="flex h-12 w-12 flex-shrink-0 items-center justify-center rounded-xl bg-[#005eb8] text-white shadow-[0_12px_26px_-12px_rgba(0,94,184,0.6)]">
          <HeartPulse class="h-6 w-6" />
        </div>
        <div class="min-w-0">
          <h1 class="text-[24px] font-bold leading-tight tracking-tight">医疗紧急联系</h1>
          <p class="mt-1 text-[12px] tracking-wide text-slate-500">Medical Emergency Contact</p>
        </div>
      </div>
    </header>

    <div class="px-6 pb-12">
      <!-- Patient chart -->
      <section v-if="hasMedical" class="rise rise-1 mt-6">
        <div class="overflow-hidden rounded-2xl border border-slate-200 bg-white shadow-[0_18px_40px_-28px_rgba(15,42,74,0.5)]">
          <div class="flex items-center justify-between border-b border-slate-200 px-5 py-3">
            <p class="font-mono text-[10px] uppercase tracking-[0.24em] text-slate-400">Patient Chart · 医疗信息</p>
            <Activity class="h-3.5 w-3.5 text-[#005eb8]" />
          </div>
          <div class="border-l-4 border-[#005eb8]">
            <div v-if="payload.bloodType" class="flex items-center gap-4 border-b border-slate-100 px-5 py-4 last:border-b-0">
              <Droplet class="h-4 w-4 flex-shrink-0 text-[#005eb8]" />
              <div class="min-w-0 flex-1">
                <p class="font-mono text-[10px] uppercase tracking-[0.18em] text-slate-400">Blood Type · 血型</p>
                <p class="mt-1 text-[22px] font-bold leading-none tabular-nums text-slate-900">{{ payload.bloodType }}</p>
              </div>
            </div>
            <div v-if="payload.allergies" class="flex items-start gap-4 border-b border-slate-100 px-5 py-4 last:border-b-0">
              <AlertTriangle class="mt-0.5 h-4 w-4 flex-shrink-0 text-red-600" />
              <div class="min-w-0 flex-1">
                <p class="font-mono text-[10px] uppercase tracking-[0.18em] text-red-500">Allergies · 过敏史</p>
                <p class="mt-1 text-[15px] font-semibold leading-snug text-red-700">{{ payload.allergies }}</p>
              </div>
            </div>
            <div v-if="payload.medications" class="flex items-start gap-4 border-b border-slate-100 px-5 py-4 last:border-b-0">
              <Pill class="mt-0.5 h-4 w-4 flex-shrink-0 text-[#005eb8]" />
              <div class="min-w-0 flex-1">
                <p class="font-mono text-[10px] uppercase tracking-[0.18em] text-slate-400">Medications · 用药</p>
                <p class="mt-1 text-[15px] font-medium leading-snug text-slate-700">{{ payload.medications }}</p>
              </div>
            </div>
            <div v-if="payload.conditions" class="flex items-start gap-4 border-b border-slate-100 px-5 py-4 last:border-b-0">
              <Activity class="mt-0.5 h-4 w-4 flex-shrink-0 text-[#005eb8]" />
              <div class="min-w-0 flex-1">
                <p class="font-mono text-[10px] uppercase tracking-[0.18em] text-slate-400">Conditions · 既往病史</p>
                <p class="mt-1 text-[15px] font-medium leading-snug text-slate-700">{{ payload.conditions }}</p>
              </div>
            </div>
          </div>
        </div>
      </section>

      <!-- Contacts -->
      <section class="rise rise-2 mt-8">
        <p class="font-mono text-[10px] uppercase tracking-[0.24em] text-slate-400">Emergency Contacts · 紧急联系人</p>
        <div v-if="contacts.length" class="mt-3 space-y-3">
          <article v-for="(c, i) in contacts" :key="i" class="rounded-2xl border border-slate-200 bg-white p-4 shadow-[0_14px_30px_-26px_rgba(15,42,74,0.55)]">
            <div class="flex items-center justify-between gap-3">
              <div class="min-w-0">
                <p class="truncate text-[15px] font-semibold text-slate-900">{{ c.name || '联系人' }}</p>
                <p v-if="c.relation" class="mt-0.5 font-mono text-[10px] uppercase tracking-[0.16em] text-[#005eb8]">{{ c.relation }}</p>
              </div>
              <a
                :href="'tel:' + c.phone"
                class="flex h-12 w-12 flex-shrink-0 items-center justify-center rounded-xl bg-[#005eb8] text-white shadow-[0_10px_22px_-10px_rgba(0,94,184,0.65)] transition-all hover:bg-[#004b93] active:scale-[0.95]"
              >
                <Phone class="h-5 w-5" />
              </a>
            </div>
            <a
              :href="'tel:' + c.phone"
              class="mt-3.5 block border-t border-slate-100 pt-3.5 text-[20px] font-semibold tabular-nums tracking-wide text-slate-900 transition-colors hover:text-[#005eb8]"
            >
              {{ c.phone }}
            </a>
          </article>
        </div>
        <div v-else class="mt-3 rounded-2xl border border-dashed border-slate-300 bg-white px-6 py-10 text-center text-sm text-slate-500">暂无紧急联系人</div>
      </section>

      <!-- Message & photos -->
      <section class="rise rise-3 mt-9 border-t border-slate-200 pt-7">
        <p class="font-mono text-[10px] uppercase tracking-[0.24em] text-slate-400">Message · 留言</p>
        <textarea
          v-model="note.content"
          rows="3"
          class="mt-4 w-full resize-none rounded-xl border border-slate-200 bg-white p-3.5 text-sm text-slate-800 placeholder:text-slate-400 focus:border-[#005eb8] focus:outline-none focus:ring-2 focus:ring-[#005eb8]/15"
          placeholder="请输入留言或位置信息..."
        ></textarea>
        <button
          @click="handleSubmitNote"
          :disabled="submitting"
          class="mt-3 flex w-full items-center justify-center gap-2 rounded-xl bg-[#005eb8] py-3.5 text-sm font-semibold text-white transition-all hover:bg-[#004b93] active:scale-[0.98] disabled:opacity-60"
        >
          <Send v-if="!submitting" class="h-4 w-4" />
          <Loader2 v-else class="h-4 w-4 animate-spin" />
          {{ submitting ? '发送中...' : '发送留言' }}
        </button>

        <div class="mt-7 grid grid-cols-3 gap-2">
          <div v-for="p in photos" :key="p.id" class="aspect-square overflow-hidden rounded-xl border border-slate-200 bg-slate-100">
            <img :src="photoUrl(p.oss_key)" class="h-full w-full object-cover" />
          </div>
          <label class="flex aspect-square cursor-pointer flex-col items-center justify-center gap-1 rounded-xl border-2 border-dashed border-slate-300 text-slate-400 transition-colors hover:bg-slate-50">
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
import { HeartPulse, Activity, Droplet, AlertTriangle, Pill, Phone, Send, Loader2, Camera } from 'lucide-vue-next'
import { ElMessage } from 'element-plus'
import { fetchPhotos, uploadPhoto, submitNote as apiSubmitNote } from '../../api/public'

const props = defineProps(['payload', 'nfcId'])
const note = reactive({ content: '' })
const photos = ref([])
const submitting = ref(false)

const contacts = computed(() =>
  (Array.isArray(props.payload?.contacts) ? props.payload.contacts : []).filter((c) => c && c.phone)
)

const hasMedical = computed(() =>
  !!(props.payload?.bloodType || props.payload?.allergies || props.payload?.medications || props.payload?.conditions)
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
