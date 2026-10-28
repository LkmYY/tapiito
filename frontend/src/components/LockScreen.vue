<template>
  <div class="lock-wrap">
    <div class="lock-card">
      <!-- 暗号模式 -->
      <template v-if="mode === 'passcode'">
        <div class="lock-icon">
          <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8" stroke-linecap="round">
            <rect x="4" y="10" width="16" height="10" rx="2.5" />
            <path d="M8 10V7a4 4 0 0 1 8 0v3" />
          </svg>
        </div>
        <h1 class="lock-title">需要暗号</h1>
        <p v-if="hint" class="lock-hint">{{ hint }}</p>
        <p v-else class="lock-hint">这是一枚私密标签，输入暗号后才能查看</p>

        <form class="lock-form" @submit.prevent="submit">
          <input
            ref="inputRef"
            v-model="passcode"
            type="password"
            inputmode="text"
            autocomplete="off"
            maxlength="64"
            :disabled="submitting || locked"
            placeholder="输入暗号"
            class="lock-input"
          />
          <button type="submit" class="lock-btn" :disabled="!passcode || submitting || locked">
            {{ submitting ? '验证中…' : locked ? '已锁定' : '进入' }}
          </button>
        </form>

        <p v-if="error" class="lock-error">{{ error }}</p>
        <p v-if="locked" class="lock-error">尝试次数过多，{{ lockMinutes }} 分钟后可再试</p>
      </template>

      <!-- 链接失效（ticket 被用掉 / 过期 / 转发） -->
      <template v-else>
        <div class="lock-icon lock-icon-warn">
          <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8" stroke-linecap="round">
            <circle cx="12" cy="12" r="9" />
            <path d="M12 7v6M12 16.5v.5" />
          </svg>
        </div>
        <h1 class="lock-title">链接已失效</h1>
        <p class="lock-hint">这个链接只能本人使用，转发或重复打开都会失效</p>
        <p class="lock-hint lock-hint-dim">如果你是本人，请重新碰一碰标签</p>
      </template>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, nextTick } from 'vue'
import { unlockNfc } from '../api/public'
import { getSessionId } from '../composables/useAccess'

const props = defineProps({
  nfcId: { type: String, required: true },
  mode: { type: String, default: 'passcode' }, // passcode | ticket
  hint: { type: String, default: '' },
})
const emit = defineEmits(['unlocked'])

const passcode = ref('')
const submitting = ref(false)
const error = ref('')
const inputRef = ref(null)
const lockSeconds = ref(0)
let timer = null

const locked = computed(() => lockSeconds.value > 0)
const lockMinutes = computed(() => Math.max(1, Math.ceil(lockSeconds.value / 60)))

const startCountdown = (sec) => {
  lockSeconds.value = Math.max(1, Math.floor(sec || 1800))
  clearInterval(timer)
  timer = setInterval(() => {
    lockSeconds.value -= 1
    if (lockSeconds.value <= 0) {
      clearInterval(timer)
      error.value = ''
    }
  }, 1000)
}

const submit = async () => {
  if (submitting.value || locked.value || !passcode.value) return
  submitting.value = true
  error.value = ''
  try {
    const { data } = await unlockNfc(props.nfcId, passcode.value)
    emit('unlocked', {
      ticket: data.ticket,
      expiresAt: data.expires_at,
      session: getSessionId(),
    })
  } catch (err) {
    const status = err?.response?.status
    const body = err?.response?.data || {}
    if (status === 429) {
      startCountdown(body.retry_after || 1800)
      error.value = body.detail || '尝试次数过多'
    } else if (status === 401) {
      const left = body.remaining
      error.value = left > 0 ? `暗号不正确，还可尝试 ${left} 次` : '暗号不正确'
      passcode.value = ''
      nextTick(() => inputRef.value?.focus())
    } else if (status === 400) {
      error.value = body.detail || '该标签未设置暗号'
    } else {
      error.value = '验证失败，请稍后再试'
    }
  } finally {
    submitting.value = false
  }
}

onMounted(() => {
  nextTick(() => inputRef.value?.focus())
})
</script>

<style scoped>
.lock-wrap {
  min-height: 100vh;
  display: grid;
  place-items: center;
  padding: 24px;
  background: #fbfaf7;
  background-image: radial-gradient(900px 420px at 50% -10%, rgba(255, 107, 53, 0.09), transparent 62%);
}
.lock-card {
  width: 100%;
  max-width: 360px;
  text-align: center;
}
.lock-icon {
  width: 52px;
  height: 52px;
  margin: 0 auto 20px;
  display: grid;
  place-items: center;
  border-radius: 16px;
  background: rgba(20, 24, 29, 0.06);
  color: #14181d;
}
.lock-icon svg { width: 26px; height: 26px; }
.lock-icon-warn {
  background: rgba(239, 68, 68, 0.09);
  color: #dc2626;
}
.lock-title {
  font-size: 22px;
  font-weight: 600;
  letter-spacing: -0.01em;
  color: #14181d;
}
.lock-hint {
  margin-top: 10px;
  font-size: 13.5px;
  line-height: 1.6;
  color: rgba(20, 24, 29, 0.55);
}
.lock-hint-dim {
  margin-top: 6px;
  font-size: 12.5px;
  color: rgba(20, 24, 29, 0.38);
}
.lock-form {
  margin-top: 26px;
  display: flex;
  flex-direction: column;
  gap: 10px;
}
.lock-input {
  width: 100%;
  height: 48px;
  padding: 0 16px;
  border: 1.5px solid rgba(20, 24, 29, 0.12);
  border-radius: 13px;
  background: #fff;
  font-size: 16px;
  letter-spacing: 0.08em;
  text-align: center;
  outline: none;
  transition: border-color 0.2s ease, box-shadow 0.2s ease;
}
.lock-input:focus {
  border-color: rgba(255, 107, 53, 0.6);
  box-shadow: 0 0 0 3px rgba(255, 107, 53, 0.1);
}
.lock-btn {
  height: 48px;
  border-radius: 13px;
  background: #14181d;
  color: #fff;
  font-size: 15px;
  font-weight: 500;
  transition: opacity 0.2s ease, transform 0.15s ease;
}
.lock-btn:disabled {
  opacity: 0.4;
}
.lock-btn:not(:disabled):active {
  transform: scale(0.98);
}
.lock-error {
  margin-top: 14px;
  font-size: 12.5px;
  color: #dc2626;
}
</style>
