<template>
  <Transition name="fade">
    <span
        v-if="stateAnswer !== 'hidden'"
        class="status-indicator"
        :class="statusClass"
    >
      <span class="status-icon">
        <svg
            v-if="stateAnswer === 'ok'"
            xmlns="http://www.w3.org/2000/svg"
            width="20"
            height="20"
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
            stroke-width="2"
            stroke-linecap="round"
            stroke-linejoin="round"
        >
          <path d="M20 6L9 17l-5-5"/>
        </svg>

        <svg
            v-else-if="stateAnswer === 'error'"
            xmlns="http://www.w3.org/2000/svg"
            width="20"
            height="20"
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
            stroke-width="2"
            stroke-linecap="round"
            stroke-linejoin="round"
        >
          <circle cx="12" cy="12" r="10"/>
          <line x1="15" y1="9" x2="9" y2="15"/>
          <line x1="9" y1="9" x2="15" y2="15"/>
        </svg>
      </span>

      <span class="status-text">
        {{ statusText }}
      </span>
    </span>
  </Transition>
</template>

<script setup>
import { computed,defineProps } from 'vue'

const props = defineProps({
  stateAnswer: {
    type: String,
    required: true,
    validator: (value) => ['hidden', 'ok', 'error'].includes(value)
  }
})

const statusClass = computed(() => ({
  'status-success': props.stateAnswer === 'ok',
  'status-error': props.stateAnswer === 'error'
}))

const statusText = computed(() => {
  switch (props.stateAnswer) {
    case 'ok':
      return 'Проверка пройдена'
    case 'error':
      return 'Проверка не пройдена'
    default:
      return ''
  }
})
</script>

<style scoped>
.status-indicator {
  display: inline-flex;
  align-items: center;
  gap: 8px;
  padding: 8px 16px;
  border-radius: 8px;
  font-size: 14px;
  font-weight: 500;
  line-height: 1.4;
  transition: all 0.3s ease;
}

.status-success {
  color: #059669;
  background-color: #ecfdf5;
  border: 1px solid #a7f3d0;
}

.status-error {
  color: #dc2626;
  background-color: #fef2f2;
  border: 1px solid #fecaca;
}

.status-icon {
  display: flex;
  align-items: center;
  justify-content: center;
}

.status-text {
  font-family: system-ui, -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
  letter-spacing: 0.01em;
}

/* Анимация появления/исчезновения */
.fade-enter-active,
.fade-leave-active {
  transition: all 0.3s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
  transform: translateY(-10px);
}

/* Hover эффекты */
.status-success:hover {
  background-color: #d1fae5;
}

.status-error:hover {
  background-color: #fee2e2;
}

/* Медиа-запрос для мобильных устройств */
@media (max-width: 640px) {
  .status-indicator {
    padding: 6px 12px;
    font-size: 13px;
  }

  .status-icon svg {
    width: 16px;
    height: 16px;
  }
}
</style>