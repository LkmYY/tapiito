<template>
  <div class="min-h-screen bg-[#fbfaf7] font-sans text-[#14181d]">
    <SiteHeader />

    <!-- 提交成功 -->
    <section v-if="submitted" class="mx-auto flex min-h-[calc(100vh-64px)] max-w-[640px] items-center px-6 py-16">
      <div class="w-full rounded-2xl border border-[#14181d]/[0.07] bg-white p-10 text-center">
        <div class="mx-auto mb-5 flex h-14 w-14 items-center justify-center rounded-full bg-[#15151B]">
          <Check class="h-6 w-6 text-white" />
        </div>
        <h1 class="text-[22px] font-semibold tracking-tight">已收到，我们会尽快联系你</h1>
        <p class="mt-3 text-[14px] leading-relaxed text-[#14181d]/55">
          工作时间通常 2 小时内回复。急用可以直接打电话给我们。
        </p>
        <div class="mt-8 flex flex-wrap items-center justify-center gap-3">
          <router-link to="/showroom" class="btn-dark">先逛逛展厅</router-link>
          <router-link to="/" class="btn-ghost">回首页</router-link>
        </div>
      </div>
    </section>

    <!-- 表单 -->
    <section v-else class="mx-auto max-w-[640px] px-6 pb-24 pt-14">
      <p class="text-[11px] font-semibold uppercase tracking-[0.2em] text-[#ff6b35]">Consultation</p>
      <h1 class="mt-3 text-[32px] font-semibold leading-tight tracking-[-0.02em] md:text-[38px]">
        先聊需求，再谈报价
      </h1>
      <p class="mt-4 text-[14.5px] leading-relaxed text-[#14181d]/55">
        留个联系方式，我们会主动联系你确认场景、数量和交付时间，确认后再给方案与报价。
        不需要先付款，也不会有推销电话。
      </p>

      <form class="mt-10 space-y-6" @submit.prevent="submit">
        <div class="grid gap-5 sm:grid-cols-2">
          <label class="field">
            <span class="field-label">怎么称呼你 <i>*</i></span>
            <input v-model.trim="form.name" class="field-input" maxlength="32" placeholder="张三" />
          </label>
          <label class="field">
            <span class="field-label">手机号 <i>*</i></span>
            <input
              v-model.trim="form.phone"
              class="field-input"
              inputmode="numeric"
              maxlength="11"
              placeholder="13800138000"
            />
          </label>
        </div>

        <label class="field">
          <span class="field-label">微信号 <span class="opt">选填，方便微信沟通</span></span>
          <input v-model.trim="form.wechat" class="field-input" maxlength="32" placeholder="zhangsan" />
        </label>

        <div class="field">
          <span class="field-label">想做哪种标签 <span class="opt">可多选</span></span>
          <div class="mt-2 flex flex-wrap gap-2">
            <button
              v-for="t in typeOptions"
              :key="t.value"
              type="button"
              class="chip"
              :class="form.types.includes(t.value) ? 'chip-on' : ''"
              @click="toggleType(t.value)"
            >
              {{ t.label }}
            </button>
          </div>
        </div>

        <label class="field">
          <span class="field-label">说说你的场景 <i>*</i></span>
          <textarea
            v-model.trim="form.note"
            class="field-input field-textarea"
            maxlength="600"
            rows="5"
            placeholder="例如：我开了一家咖啡馆，想在每张桌子上放一枚，顾客碰一下看菜单和今日推荐。大概需要 12 枚。"
          />
          <span class="field-count">{{ form.note.length }} / 600</span>
        </label>

        <p v-if="error" class="text-[13px] text-red-600">{{ error }}</p>

        <div class="flex flex-wrap items-center gap-4 pt-2">
          <button type="submit" class="btn-dark" :disabled="submitting">
            {{ submitting ? '提交中…' : '提交，等我们联系' }}
          </button>
          <span class="text-[12.5px] text-[#14181d]/40">信息仅用于本次沟通，不会对外公开</span>
        </div>
      </form>
    </section>
  </div>
</template>

<script setup>
import { reactive, ref, computed } from 'vue'
import { Check } from 'lucide-vue-next'
import SiteHeader from '../components/SiteHeader.vue'
import { createFeedback } from '../api/feedback'

const typeOptions = [
  { value: 'business_card', label: '商务名片' },
  { value: 'menu', label: '门店菜单' },
  { value: 'company', label: '企业官网' },
  { value: 'product', label: '产品展示' },
  { value: 'nightclub', label: '夜店交际' },
  { value: 'emergency', label: '求助信息卡' },
  { value: 'pet', label: '宠物铭牌' },
  { value: 'child', label: '儿童信息卡' },
  { value: 'recipe', label: '菜谱' },
]

const form = reactive({ name: '', phone: '', wechat: '', types: [], note: '' })
const submitting = ref(false)
const submitted = ref(false)
const error = ref('')

const phoneOk = computed(() => /^1[3-9]\d{9}$/.test(form.phone))

const toggleType = (v) => {
  const i = form.types.indexOf(v)
  if (i >= 0) form.types.splice(i, 1)
  else form.types.push(v)
}

const submit = async () => {
  error.value = ''
  if (!form.name) return (error.value = '请填写称呼')
  if (!phoneOk.value) return (error.value = '请填写正确的 11 位手机号')
  if (form.note.length < 5) return (error.value = '请简单描述一下使用场景')

  const labels = form.types
    .map((v) => typeOptions.find((t) => t.value === v)?.label)
    .filter(Boolean)

  // 复用反馈接口投递，后台统一在「用户反馈」里查看
  const content = [
    '【预约咨询】',
    `称呼：${form.name}`,
    `手机：${form.phone}`,
    form.wechat ? `微信：${form.wechat}` : '',
    labels.length ? `意向类型：${labels.join('、')}` : '',
    `场景描述：${form.note}`,
  ]
    .filter(Boolean)
    .join('\n')

  submitting.value = true
  try {
    await createFeedback({ content })
    submitted.value = true
    window.scrollTo({ top: 0, behavior: 'smooth' })
  } catch (e) {
    error.value = e?.response?.data?.detail || '提交失败，请稍后再试或直接联系我们'
  } finally {
    submitting.value = false
  }
}
</script>

<style scoped>
.field { display: block; }
.field-label {
  display: block;
  margin-bottom: 8px;
  font-size: 13px;
  color: rgba(20, 24, 29, 0.62);
}
.field-label i { color: #ff6b35; font-style: normal; }
.opt { font-size: 11.5px; color: rgba(20, 24, 29, 0.32); margin-left: 4px; }
.field-input {
  width: 100%;
  padding: 0.7rem 0.9rem;
  border: 1.5px solid rgba(20, 24, 29, 0.1);
  border-radius: 11px;
  background: #fff;
  font-size: 14.5px;
  color: #14181d;
  outline: none;
  transition: border-color 0.2s ease, box-shadow 0.2s ease;
}
.field-input:focus {
  border-color: rgba(255, 107, 53, 0.55);
  box-shadow: 0 0 0 3px rgba(255, 107, 53, 0.09);
}
.field-textarea { resize: vertical; line-height: 1.7; }
.field-count {
  display: block;
  margin-top: 6px;
  text-align: right;
  font-size: 11.5px;
  color: rgba(20, 24, 29, 0.3);
}
.chip {
  padding: 0.42rem 0.85rem;
  border: 1px solid rgba(20, 24, 29, 0.12);
  border-radius: 999px;
  background: #fff;
  font-size: 13px;
  color: rgba(20, 24, 29, 0.6);
  transition: all 0.18s ease;
}
.chip-on {
  border-color: #15151b;
  background: #15151b;
  color: #fff;
}
.btn-dark {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  padding: 0.72rem 1.4rem;
  border-radius: 11px;
  background: #15151b;
  color: #fff;
  font-size: 14.5px;
  font-weight: 500;
  transition: opacity 0.2s ease, transform 0.15s ease;
}
.btn-dark:disabled { opacity: 0.45; }
.btn-dark:not(:disabled):active { transform: scale(0.985); }
.btn-ghost {
  display: inline-flex;
  align-items: center;
  padding: 0.72rem 1.2rem;
  border-radius: 11px;
  border: 1px solid rgba(20, 24, 29, 0.12);
  font-size: 14px;
  color: #14181d;
}
</style>
