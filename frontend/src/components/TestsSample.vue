<template>
  <div class="input-container" v-if="!props.hidden">
    <!-- Row 1 -->
    <div class="input-row" v-for="i in arr" :key="i">
      <div class="input-group">
        <input
            type="text"
            :value="fields[i-1].field1"
            placeholder="Введите значение"
            class="custom-input"
            @input = "editFields($event, i-1, 'field1')"
        >
        <label class="floating-label">Входные Данные {{i}}</label>
      </div>
      <div class="connector">
        <span class="connector-line"></span>
      </div>
      <div class="input-group">
        <input
            type="text"
            :value="fields[i-1].field2"
            placeholder="Введите значение"
            class="custom-input"
            @input="editFields($event, i-1, 'field2')"
        >
        <label class="floating-label">Выходные данные {{i}}</label>
      </div>
    </div>
  </div>
</template>

<script setup>

import {defineProps, defineEmits, computed, ref} from 'vue'
const arr = [1,2,3,4,5]

const props = defineProps({
  hidden : Boolean,
  fields : {
    type: Array,
    default: () => []
  }
})

const inputRef = ref(null)
const inputFields = computed({
  get : () => {
    return inputRef.value ?? props.fields
  },
  set : (value) => {
    inputRef.value = value
  }
})

const emit = defineEmits("update:fields")
const editFields = (event, i, field) => {
  let elem = inputFields.value
  elem[i][field] = event.target.value
  inputFields.value = elem
  emit("update:fields", inputFields.value)
}
</script>

<style scoped>
.input-container {
  max-width: 800px;
  margin-top: -5%;
  margin-left: 10%;
  padding: 2rem;
  background: #f8f9fa;
  border-radius: 12px;
  box-shadow: 0 8px 30px rgba(0, 0, 0, 0.05);
}

.input-row {
  display: flex;
  align-items: center;
  margin-bottom: 2rem;
  position: relative;
}

.input-row:last-child {
  margin-bottom: 0;
}

.input-group {
  flex: 1;
  position: relative;
}

.custom-input {
  width: 100%;
  padding: 12px 16px;
  font-size: 16px;
  border: 2px solid #e0e0e0;
  border-radius: 8px;
  background: white;
  transition: all 0.3s ease;
  outline: none;
}

.custom-input:focus {
  border-color: #4a90e2;
  box-shadow: 0 0 0 3px rgba(74, 144, 226, 0.1);
}

.floating-label {
  position: absolute;
  left: 16px;
  top: -10px;
  background: white;
  padding: 0 6px;
  font-size: 14px;
  color: #666;
  transition: all 0.3s ease;
}

.custom-input:focus + .floating-label {
  color: #4a90e2;
}

.connector {
  width: 60px;
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 0 10px;
}

.connector-line {
  height: 2px;
  width: 100%;
  background: linear-gradient(to right, #e0e0e0, #4a90e2, #e0e0e0);
  position: relative;
}

.connector-line::before,
.connector-line::after {
  content: '';
  position: absolute;
  top: 50%;
  width: 6px;
  height: 6px;
  border-radius: 50%;
  background: #4a90e2;
  transform: translateY(-50%);
}

.connector-line::before {
  left: -2px;
}

.connector-line::after {
  right: -2px;
}

/* Hover Effects */
.input-row:hover .connector-line {
  background: linear-gradient(to right, #4a90e2, #5c6bc0, #4a90e2);
  animation: pulse 1.5s infinite;
}

.input-row:hover .custom-input {
  border-color: #4a90e2;
}

/* Animation */
@keyframes pulse {
  0% {
    opacity: 0.6;
  }
  50% {
    opacity: 1;
  }
  100% {
    opacity: 0.6;
  }
}

/* Responsive Design */
@media (max-width: 600px) {
  .input-container {
    padding: 1rem;
  }

  .input-row {
    flex-direction: column;
    gap: 1rem;
  }

  .connector {
    width: 2px;
    height: 30px;
  }

  .connector-line {
    width: 2px;
    height: 100%;
    background: linear-gradient(to bottom, #e0e0e0, #4a90e2, #e0e0e0);
  }

  .connector-line::before,
  .connector-line::after {
    left: 50%;
    transform: translateX(-50%);
  }

  .connector-line::before {
    top: -2px;
  }

  .connector-line::after {
    top: auto;
    bottom: -2px;
  }
}
</style>